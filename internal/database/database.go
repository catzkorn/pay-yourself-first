package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/catzkorn/pay-yourself-first/internal/income"
	"github.com/catzkorn/pay-yourself-first/internal/saving"
	"github.com/jackc/pgtype"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/shopspring/decimal"
)

// Database allows the user to store and read back information from the DB
type Database struct {
	database *sql.DB
}

// NewDatabaseConnection starts connection with database
func NewDatabaseConnection(databaseString string) (*Database, error) {

	db, err := sql.Open("pgx", os.Getenv(databaseString))
	if err != nil {
		return nil, fmt.Errorf("unexpected connection error: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("could not ping database: %w", err)
	}
	return &Database{database: db}, nil
}

// RecordIncome inserts a type of income into the temporary data store
func (d *Database) RecordIncome(ctx context.Context, i income.Income) (*income.Income, error) {

	var id uint32
	var date time.Time
	var source string
	var amount pgtype.Numeric

	insertQuery := `
	INSERT INTO income (date, source, amount)
	VALUES ($1, $2, $3)
	ON CONFLICT(date)
	DO UPDATE SET source=EXCLUDED.source, amount=EXCLUDED.amount
	RETURNING id, date, source, amount;
	`

	err := d.database.QueryRowContext(
		ctx,
		insertQuery,
		i.Date,
		i.Source,
		i.Amount,
	).Scan(
		&id,
		&date,
		&source,
		&amount,
	)

	if err != nil {
		return nil, fmt.Errorf("unexpected insert error: %w", err)
	}

	returnedIncome := income.Income{
		ID:     id,
		Date:   date,
		Source: source,
		Amount: decimal.NewFromBigInt(amount.Int, amount.Exp),
	}

	return &returnedIncome, nil
}

// ListIncomes returns every entry of income ordered by descending date
func (d *Database) ListIncomes(ctx context.Context) ([]income.Income, error) {

	selectQuery := `
	SELECT id, date, source, amount
	FROM income
	ORDER BY date;
	`

	rows, err := d.database.QueryContext(
		ctx,
		selectQuery,
	)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve all income data: %w", err)
	}

	var incomes []income.Income

	for rows.Next() {
		var id uint32
		var date time.Time
		var source string
		var amount pgtype.Numeric

		err = rows.Scan(
			&id,
			&date,
			&source,
			&amount,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		incomes = append(incomes, income.Income{
			ID:     id,
			Date:   date,
			Source: source,
			Amount: decimal.NewFromBigInt(amount.Int, amount.Exp),
		})
	}

	return incomes, nil
}

// RetrieveMonthIncome retrieves all income for a specific month
// and year
func (d *Database) RetrieveMonthIncome(ctx context.Context, date time.Time) (*income.Income, error) {

	var id uint32
	var returnedDate time.Time
	var source string
	var amount pgtype.Numeric

	selectQuery := `
	SELECT id, date, source, amount
	FROM income
	WHERE date = $1;
	`

	err := d.database.QueryRowContext(
		ctx,
		selectQuery,
		date,
	).Scan(
		&id,
		&returnedDate,
		&source,
		&amount,
	)

	switch {
	case err == sql.ErrNoRows:
		return nil, income.ErrNoIncomeForMonth
	case err != nil:
		return nil, fmt.Errorf("unexpected database error: %w", err)
	default:
		monthIncome := income.Income{
			ID:     id,
			Date:   returnedDate,
			Source: source,
			Amount: decimal.NewFromBigInt(amount.Int, amount.Exp),
		}
		return &monthIncome, nil
	}
}

// DeleteIncome deletes an income based on the id
func (d *Database) DeleteIncome(ctx context.Context, id uint32) error {

	deleteQuery := `
	DELETE FROM income
	WHERE id = $1
	`

	result, err := d.database.ExecContext(
		ctx,
		deleteQuery,
		id,
	)
	if err != nil {
		return fmt.Errorf("unexpected database error: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected by deletion: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows were affected by deletion request")
	}
	return nil
}

// Savings

// RecordMonthSavingPercent inserts savings data for a specific month
func (d *Database) RecordMonthSavingPercent(ctx context.Context, s saving.Saving) (*saving.Saving, error) {

	var returnedSaving saving.Saving

	insertQuery := `
	INSERT INTO saving (percent, date)
	VALUES ($1, $2)
	ON CONFLICT(date)
	DO UPDATE SET percent=EXCLUDED.percent
	RETURNING id, percent, date;
	`

	err := d.database.QueryRowContext(
		ctx,
		insertQuery,
		s.Percent,
		s.Date,
	).Scan(
		&returnedSaving.ID,
		&returnedSaving.Percent,
		&returnedSaving.Date,
	)

	if err != nil {
		return nil, fmt.Errorf("unexpected insert error: %w", err)
	}

	return &returnedSaving, nil
}

// RetrieveMonthSavingPercent returns the stored saving percent of a specific month
func (d *Database) RetrieveMonthSavingPercent(ctx context.Context, date time.Time) (*saving.Saving, error) {

	var returnedSaving saving.Saving

	selectQuery := `
	SELECT id, percent, date
	FROM saving
	WHERE date = $1;
	`

	err := d.database.QueryRowContext(
		ctx,
		selectQuery,
		date,
	).Scan(
		&returnedSaving.ID,
		&returnedSaving.Percent,
		&returnedSaving.Date,
	)

	switch {
	case err == sql.ErrNoRows:
		return nil, saving.ErrNoSavingForMonth
	case err != nil:
		return nil, fmt.Errorf("unexpected database error: %w", err)
	default:
		return &returnedSaving, nil
	}
}
