package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/catzkorn/pay-yourself-first/internal/income"
	"github.com/jackc/pgtype"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/shopspring/decimal"
)

// Database allows the user to store and read back information from the DB
type Database struct {
	database *sql.DB
}

// NewDatabaseConnection starts connection with database
func NewDatabaseConnection() (*Database, error) {

	db, err := sql.Open("pgx", os.Getenv("DATABASE_CONN_STRING"))
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

// GetMonthIncome retrieves all income for a specific month
// and year
func (d *Database) GetMonthIncome(ctx context.Context, date time.Time) (*income.Income, error) {

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
		return nil, fmt.Errorf("no income for specified month")
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
