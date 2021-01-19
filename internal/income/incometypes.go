package income

import (
	"errors"
	"time"

	"github.com/shopspring/decimal"
)

// Income defines an income by type and amount
type Income struct {
	ID     uint32
	Date   time.Time
	Source string
	Amount decimal.Decimal
}

// ErrNoIncomeForMonth is the error returned when no data could be
// found for the month requested
var ErrNoIncomeForMonth = errors.New("no income for selected month")
