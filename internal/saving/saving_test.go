package saving

import (
	"testing"
	"time"
)

func TestSavingValidation(t *testing.T) {

	t.Run("errors if value over 100 is submitted", func(t *testing.T) {
		testSaving := Saving{
			Percent: 101,
			Date:    time.Date(2021, time.February, 1, 0, 0, 0, 0, time.UTC),
		}

		err := testSaving.Validate()
		if err == nil {
			t.Fatalf("expected validation to error as saving percent is over 100")
		}
	})

	t.Run("errors if value under 0 is submitted", func(t *testing.T) {
		testSaving := Saving{
			Percent: -1,
			Date:    time.Date(2021, time.February, 1, 0, 0, 0, 0, time.UTC),
		}

		err := testSaving.Validate()
		if err == nil {
			t.Fatalf("expected validation to error as saving percent is under 0")
		}
	})

	t.Run("errors if date is zero value", func(t *testing.T) {
		testSaving := Saving{
			Percent: -1,
		}

		err := testSaving.Validate()
		if err == nil {
			t.Fatalf("expected validation to error as date is zero value")
		}
	})

	t.Run("expects validation to pass", func(t *testing.T) {
		testSaving := Saving{
			Percent: 40,
			Date:    time.Date(2021, time.June, 1, 0, 0, 0, 0, time.UTC),
		}

		err := testSaving.Validate()
		if err != nil {
			t.Fatalf("validation failed when expected to pass with data within constraints")
		}
	})
}
