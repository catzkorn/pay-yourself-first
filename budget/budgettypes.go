package budget

import (
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
