package expenses

import (
	"errors"
	"time"

	"github.com/shopspring/decimal"
)

// Expense defines information linked with an expense
type Expense struct {
	ID         uint32
	Date       time.Time
	Source     string
	Amount     decimal.Decimal
	Occurrence string
}

// ErrNoSavingForMonth is the error returned when no data could be
// found for the month requested
var ErrNoExpensesForMonth = errors.New("no expenses data for selected month")
