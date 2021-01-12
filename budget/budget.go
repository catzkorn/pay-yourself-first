package budget

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgtype"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/shopspring/decimal"
)

// DataStore defines the interface required for a budget
type DataStore interface {
	RecordIncome(income *Income) error
}

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
		return nil, fmt.Errorf("unexpected connection error: %w", err)
	}
	return &Database{database: db}, nil
}

// RecordIncome inserts a type of income into the temporary data store
func (d *Database) RecordIncome(income Income) (*Income, error) {

	var id uint32
	var date time.Time
	var source string
	var amount pgtype.Numeric

	insertQuery := `
	INSERT INTO income (date, source, amount)
	VALUES ($1, $2, $3)
	RETURNING id, date, source, amount
	`

	err := d.database.QueryRowContext(
		context.Background(),
		insertQuery,
		income.Date,
		income.Source,
		income.Amount,
	).Scan(
		&id,
		&date,
		&source,
		&amount,
	)

	if err != nil {
		return nil, fmt.Errorf("unexpected insert error: %w", err)
	}

	returnedIncome := Income{
		ID:     id,
		Date:   date,
		Source: source,
		Amount: decimal.NewFromBigInt(amount.Int, amount.Exp),
	}

	return &returnedIncome, nil
}

// GetAllIncome returns every entry of income ordered by descending date
// func (d *Database) GetAllIncome() ([]Income, error) {

// }
