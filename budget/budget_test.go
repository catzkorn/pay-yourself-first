package budget

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/shopspring/decimal"
)

func TestRecordIncome(t *testing.T) {

	t.Run("user can provide a salary income", func(t *testing.T) {

		amount, _ := decimal.NewFromString("1550.55")

		income := Income{
			Date:   time.Date(2020, time.Now().Month()+1, 12, 0, 0, 0, 0, time.UTC),
			Source: "Salary",
			Amount: amount,
		}

		store, err := NewDatabaseConnection()
		assertDatabaseError(t, err)

		returnIncome, err := store.RecordIncome(context.Background(), income)
		assertDatabaseError(t, err)

		if returnIncome.ID == 0 {
			t.Errorf("no id was returned")
		}

		if returnIncome.Date != income.Date {
			t.Errorf("income date is not as expected got %v want %v", returnIncome.Date, income.Date)
		}

		if returnIncome.Source != income.Source {
			t.Errorf("income source is not as expected got %v want %v", returnIncome.Source, income.Source)
		}

		if !returnIncome.Amount.Equal(income.Amount) {
			t.Errorf("income Amount is not as expected got %v want %v", returnIncome.Amount, income.Amount)
		}

		err = clearIncomeTable()
		assertDatabaseError(t, err)
	})
}

func TestGetAllIncome(t *testing.T) {

	t.Run("retrieves all income in decending order", func(t *testing.T) {

		store, err := NewDatabaseConnection()
		assertDatabaseError(t, err)

		aprilAmount, _ := decimal.NewFromString("2000.00")
		aprilIncome := Income{
			Date:   time.Date(2021, time.April, 1, 0, 0, 0, 0, time.UTC),
			Source: "Salary",
			Amount: aprilAmount,
		}

		marchAmount, _ := decimal.NewFromString("1550.55")
		marchIncome := Income{
			Date:   time.Date(2021, time.March, 1, 0, 0, 0, 0, time.UTC),
			Source: "Salary",
			Amount: marchAmount,
		}

		marchSecondAmount, _ := decimal.NewFromString("2000.00")
		marchSecondIncome := Income{
			Date:   time.Date(2021, time.March, 1, 0, 0, 0, 0, time.UTC),
			Source: "Ebay",
			Amount: marchSecondAmount,
		}

		juneAmount, _ := decimal.NewFromString("2000.00")
		juneIncome := Income{
			Date:   time.Date(2021, time.June, 1, 0, 0, 0, 0, time.UTC),
			Source: "Furlough Pay",
			Amount: juneAmount,
		}

		incomes := []Income{aprilIncome, marchIncome, marchSecondIncome, juneIncome}

		for _, income := range incomes {
			_, err = store.RecordIncome(context.Background(), income)
			assertDatabaseError(t, err)
		}

		retrievedIncome, err := store.ListIncomes(context.Background())
		assertDatabaseError(t, err)

		if len(retrievedIncome) != len(incomes) {
			t.Errorf("incorrect number of entries retrieved got %v want %v", len(retrievedIncome), len(incomes))
		}

		t.Logf("%v", retrievedIncome)

		if retrievedIncome[0].Date != marchIncome.Date {
			t.Errorf("list of incomes not returned in correct order")
		}

		err = clearIncomeTable()
		assertDatabaseError(t, err)

	})
}

func TestGetMonthlyIncome(t *testing.T) {

	t.Run("retrieves all income for a specific month", func(t *testing.T) {

		store, err := NewDatabaseConnection()
		assertDatabaseError(t, err)

		aprilAmount, _ := decimal.NewFromString("2000.00")
		aprilIncome := Income{
			Date:   time.Date(2021, time.April, 1, 0, 0, 0, 0, time.UTC),
			Source: "Salary",
			Amount: aprilAmount,
		}

		marchAmount, _ := decimal.NewFromString("1550.55")
		marchIncome := Income{
			Date:   time.Date(2021, time.March, 1, 0, 0, 0, 0, time.UTC),
			Source: "Salary",
			Amount: marchAmount,
		}

		marchSecondAmount, _ := decimal.NewFromString("2000.00")
		marchSecondIncome := Income{
			Date:   time.Date(2021, time.March, 1, 0, 0, 0, 0, time.UTC),
			Source: "Ebay",
			Amount: marchSecondAmount,
		}

		juneAmount, _ := decimal.NewFromString("2000.00")
		juneIncome := Income{
			Date:   time.Date(2021, time.June, 1, 0, 0, 0, 0, time.UTC),
			Source: "Furlough Pay",
			Amount: juneAmount,
		}

		incomes := []Income{aprilIncome, marchIncome, marchSecondIncome, juneIncome}

		for _, income := range incomes {
			_, err = store.RecordIncome(context.Background(), income)
			assertDatabaseError(t, err)
		}

		monthOfDate := time.Date(2021, time.March, 1, 0, 0, 0, 0, time.UTC)

		marchIncomes, err := store.GetMonthIncome(context.Background(), monthOfDate)
		assertDatabaseError(t, err)

		if len(marchIncomes) != 2 {
			t.Errorf("did not retrieve expected number of income results got %v want %v", len(marchIncomes), 2)
		}

		if marchIncomes[0].Date != monthOfDate && marchIncomes[1].Date != monthOfDate {
			t.Errorf("did not retrieve incomes from the correct month got %v & %v want %v", marchIncomes[0].Date, marchIncomes[1].Date, monthOfDate)
		}

		err = clearIncomeTable()
		assertDatabaseError(t, err)

	})

}

func assertDatabaseError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected database error: %v", err)
	}
}

func clearIncomeTable() error {
	db, err := sql.Open("pgx", os.Getenv("DATABASE_CONN_STRING"))
	if err != nil {
		return fmt.Errorf("unexpected connection error: %w", err)
	}
	_, err = db.ExecContext(context.Background(), "TRUNCATE TABLE income;")

	return err
}
