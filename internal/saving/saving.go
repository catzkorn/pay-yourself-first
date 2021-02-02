package saving

import "fmt"

// Validate checks that the savings data submitted is within expected limits
func (s *Saving) Validate() error {

	switch {
	case s.Percent > 100:
		return fmt.Errorf("saving percent exceeds 100")
	case s.Percent < 0:
		return fmt.Errorf("saving percent is under 0")
	case s.Date.IsZero():
		return fmt.Errorf("date was not provided")
	default:
		return nil
	}
}
