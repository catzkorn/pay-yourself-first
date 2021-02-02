package saving

import "time"

// Saving defines a saving level for a specific month
type Saving struct {
	ID      uint32
	Percent int
	Date    time.Time
}
