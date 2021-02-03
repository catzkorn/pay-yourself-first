package expenses

import (
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
