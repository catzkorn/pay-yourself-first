package saving

import (
	"errors"
	"time"
)

// Saving defines a saving level for a specific month
type Saving struct {
	ID      uint32
	Percent int
	Date    time.Time
}

// ErrNoSavingForMonth is the error returned when no data could be
// found for the month requested
var ErrNoSavingForMonth = errors.New("no saving data for selected month")
