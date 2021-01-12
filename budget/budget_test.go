package budget

import (
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

		returnIncome, err := store.RecordIncome(income)
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

	})

}

func assertDatabaseError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected database error: %v", err)
	}
}
