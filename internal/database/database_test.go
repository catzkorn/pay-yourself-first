package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/catzkorn/pay-yourself-first/internal/income"
	"github.com/catzkorn/pay-yourself-first/internal/saving"
	"github.com/shopspring/decimal"
)

func TestRecordIncome(t *testing.T) {

	t.Run("user can provide a salary income", func(t *testing.T) {

		amount, _ := decimal.NewFromString("1550.55")

		income := income.Income{
			Date:   time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			Source: "Salary",
			Amount: amount,
		}

		store, err := NewDatabaseConnection("DATABASE_CONN_TEST_STRING")
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

	t.Run("user can update an income for a previously submitted month", func(t *testing.T) {

		amount, _ := decimal.NewFromString("5000.00")

		julyIncome := income.Income{
			Date:   time.Date(2020, time.July, 1, 0, 0, 0, 0, time.UTC),
			Source: "Salary",
			Amount: amount,
		}

		ctx := context.Background()

		store, err := NewDatabaseConnection("DATABASE_CONN_TEST_STRING")
		assertDatabaseError(t, err)

		returnedJulyIncome, err := store.RecordIncome(ctx, julyIncome)
		assertDatabaseError(t, err)

		julyIncome.ID = returnedJulyIncome.ID
		updatedAmount, _ := decimal.NewFromString("5500.00")
		julyIncome.Amount = updatedAmount

		updatedJulyIncome, err := store.RecordIncome(ctx, julyIncome)
		assertDatabaseError(t, err)

		if updatedJulyIncome.ID != returnedJulyIncome.ID {
			t.Errorf("incorrect id returned got %v want %v", updatedJulyIncome.ID, returnedJulyIncome.ID)
		}

		if updatedJulyIncome.Date != (returnedJulyIncome.Date) {
			t.Errorf("incorrect date returned got %v want %v", updatedJulyIncome.Date, returnedJulyIncome.Date)
		}

		if !updatedJulyIncome.Amount.Equal(updatedAmount) {
			t.Errorf("amount was not updated to the specified amount got %v want %v", updatedJulyIncome.Amount, updatedAmount)
		}

		if updatedJulyIncome.Amount.Equal(returnedJulyIncome.Amount) {
			t.Errorf("amount was not updated got %v want %v", updatedJulyIncome.Amount, returnedJulyIncome.Amount)
		}

		err = clearIncomeTable()
		assertDatabaseError(t, err)

	})
}

func TestGetAllIncome(t *testing.T) {

	t.Run("retrieves all income in decending order", func(t *testing.T) {

		store, err := NewDatabaseConnection("DATABASE_CONN_TEST_STRING")
		assertDatabaseError(t, err)

		aprilAmount, _ := decimal.NewFromString("2000.00")
		aprilIncome := income.Income{
			Date:   time.Date(2021, time.April, 1, 0, 0, 0, 0, time.UTC),
			Source: "Salary",
			Amount: aprilAmount,
		}

		marchAmount, _ := decimal.NewFromString("1550.55")
		marchIncome := income.Income{
			Date:   time.Date(2021, time.March, 1, 0, 0, 0, 0, time.UTC),
			Source: "Salary",
			Amount: marchAmount,
		}

		// TODO: Removed due to MVP change. Only one income stream per month - 14/01/21
		// marchSecondAmount, _ := decimal.NewFromString("2000.00")
		// marchSecondIncome := Income{
		// 	Date:   time.Date(2021, time.March, 1, 0, 0, 0, 0, time.UTC),
		// 	Source: "Ebay",
		// 	Amount: marchSecondAmount,
		// }

		juneAmount, _ := decimal.NewFromString("2000.00")
		juneIncome := income.Income{
			Date:   time.Date(2021, time.June, 1, 0, 0, 0, 0, time.UTC),
			Source: "Furlough Pay",
			Amount: juneAmount,
		}

		incomes := []income.Income{aprilIncome, marchIncome, juneIncome}

		for _, income := range incomes {
			_, err = store.RecordIncome(context.Background(), income)
			assertDatabaseError(t, err)
		}

		retrievedIncome, err := store.ListIncomes(context.Background())
		assertDatabaseError(t, err)

		if len(retrievedIncome) != len(incomes) {
			t.Errorf("incorrect number of entries retrieved got %v want %v", len(retrievedIncome), len(incomes))
		}

		err = clearIncomeTable()
		assertDatabaseError(t, err)
	})
}

func TestRetrieveMonthIncome(t *testing.T) {

	t.Run("retrieves all income for a specific month", func(t *testing.T) {

		store, err := NewDatabaseConnection("DATABASE_CONN_TEST_STRING")
		assertDatabaseError(t, err)

		aprilAmount, _ := decimal.NewFromString("2000.00")
		aprilIncome := income.Income{
			Date:   time.Date(2021, time.April, 1, 0, 0, 0, 0, time.UTC),
			Source: "Salary",
			Amount: aprilAmount,
		}

		marchAmount, _ := decimal.NewFromString("1550.55")
		marchIncome := income.Income{
			Date:   time.Date(2021, time.March, 1, 0, 0, 0, 0, time.UTC),
			Source: "Salary",
			Amount: marchAmount,
		}

		juneAmount, _ := decimal.NewFromString("2000.00")
		juneIncome := income.Income{
			Date:   time.Date(2021, time.June, 1, 0, 0, 0, 0, time.UTC),
			Source: "Furlough Pay",
			Amount: juneAmount,
		}

		incomes := []income.Income{aprilIncome, marchIncome, juneIncome}

		for _, income := range incomes {
			_, err = store.RecordIncome(context.Background(), income)
			assertDatabaseError(t, err)
		}

		monthOfDate := time.Date(2021, time.March, 1, 0, 0, 0, 0, time.UTC)

		result, err := store.RetrieveMonthIncome(context.Background(), monthOfDate)
		assertDatabaseError(t, err)

		// if len(marchIncomes) != 2 {
		// 	t.Errorf("did not retrieve expected number of income results got %v want %v", len(marchIncomes), 2)
		// }

		if result.Date != monthOfDate {
			t.Errorf("did not retrieve income from the correct month got %v want %v", result.Date, monthOfDate)
		}

		err = clearIncomeTable()
		assertDatabaseError(t, err)
	})
}

func TestDeleteIncome(t *testing.T) {

	t.Run("deletes a specific income entry", func(t *testing.T) {

		amount, _ := decimal.NewFromString("1550.55")

		income := income.Income{
			Date:   time.Date(2020, time.Now().Month()+1, 12, 0, 0, 0, 0, time.UTC),
			Source: "Salary",
			Amount: amount,
		}

		store, err := NewDatabaseConnection("DATABASE_CONN_TEST_STRING")
		assertDatabaseError(t, err)

		returnIncome, err := store.RecordIncome(context.Background(), income)
		assertDatabaseError(t, err)

		err = store.DeleteIncome(context.Background(), returnIncome.ID)

		if err != nil {
			t.Errorf("income entry was not deleted")
		}

		err = clearIncomeTable()
		assertDatabaseError(t, err)
	})

	t.Run("attempts to delete an entry that does not exist", func(t *testing.T) {
		store, err := NewDatabaseConnection("DATABASE_CONN_TEST_STRING")
		assertDatabaseError(t, err)

		err = store.DeleteIncome(context.Background(), 0)

		if err == nil {
			t.Errorf("an entry that does not exist did not trigger error")
		}

	})
}

func TestRecordMonthSavingPercent(t *testing.T) {

	t.Run("records a integer that represents a 45 percent saving level for a specific month", func(t *testing.T) {

		store, err := NewDatabaseConnection("DATABASE_CONN_TEST_STRING")
		assertDatabaseError(t, err)

		saving := saving.Saving{
			Percent: 45,
			Date:    time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC),
		}

		returnMonthSaving, err := store.RecordMonthSavingPercent(context.Background(), saving)
		assertDatabaseError(t, err)

		if returnMonthSaving.ID == 0 {
			t.Errorf("id returned is not valid")
		}

		if returnMonthSaving.Percent != saving.Percent {
			t.Errorf("incorrect saving percent returned for month got %v want %v", returnMonthSaving.Percent, saving.Percent)
		}

		if returnMonthSaving.Date != saving.Date {
			t.Errorf("incorrect month of saving returned got %v want %v", returnMonthSaving.Date, saving.Date)
		}

		err = clearSavingTable()
		assertDatabaseError(t, err)
	})

	t.Run("records a integer that represents a 73 percent saving level for a specific month", func(t *testing.T) {

		store, err := NewDatabaseConnection("DATABASE_CONN_TEST_STRING")
		assertDatabaseError(t, err)

		saving := saving.Saving{
			Percent: 73,
			Date:    time.Date(2021, time.February, 1, 0, 0, 0, 0, time.UTC),
		}

		returnMonthSaving, err := store.RecordMonthSavingPercent(context.Background(), saving)
		assertDatabaseError(t, err)

		if returnMonthSaving.ID == 0 {
			t.Errorf("id returned is not valid")
		}

		if returnMonthSaving.Percent != saving.Percent {
			t.Errorf("incorrect saving percent returned for month got %v want %v", returnMonthSaving.Percent, saving.Percent)
		}

		if returnMonthSaving.Date != saving.Date {
			t.Errorf("incorrect month of saving returned got %v want %v", returnMonthSaving.Date, saving.Date)
		}
		err = clearSavingTable()
		assertDatabaseError(t, err)
	})

	t.Run("records a integer that represents a 73 percent saving level for a specific month", func(t *testing.T) {

		store, err := NewDatabaseConnection("DATABASE_CONN_TEST_STRING")
		assertDatabaseError(t, err)

		initialSaving := saving.Saving{
			Percent: 25,
			Date:    time.Date(2021, time.February, 1, 0, 0, 0, 0, time.UTC),
		}

		returnMonthSaving, err := store.RecordMonthSavingPercent(context.Background(), initialSaving)
		assertDatabaseError(t, err)

		updatedSaving := saving.Saving{
			Percent: 35,
			Date:    time.Date(2021, time.February, 1, 0, 0, 0, 0, time.UTC),
		}

		returnUpdatedMonthSaving, err := store.RecordMonthSavingPercent(context.Background(), updatedSaving)
		assertDatabaseError(t, err)

		if returnMonthSaving.Date != initialSaving.Date {
			t.Fatalf("incorrect month of saving returned got %v want %v", returnMonthSaving.Date, initialSaving.Date)
		}

		if returnUpdatedMonthSaving.ID != returnMonthSaving.ID {
			t.Errorf("id returned is not as expected got %v want %v", returnUpdatedMonthSaving.ID, returnMonthSaving.ID)
		}

		if returnUpdatedMonthSaving.Percent == returnMonthSaving.Percent {
			t.Errorf("incorrect saving percent returned for month got %v want %v", returnUpdatedMonthSaving.Percent, updatedSaving.Percent)
		}

		err = clearSavingTable()
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
	db, err := sql.Open("pgx", os.Getenv("DATABASE_CONN_TEST_STRING"))
	if err != nil {
		return fmt.Errorf("unexpected connection error: %w", err)
	}
	_, err = db.ExecContext(context.Background(), "TRUNCATE TABLE income;")

	return err
}

func clearSavingTable() error {
	db, err := sql.Open("pgx", os.Getenv("DATABASE_CONN_TEST_STRING"))
	if err != nil {
		return fmt.Errorf("unexpected connection error: %w", err)
	}
	_, err = db.ExecContext(context.Background(), "TRUNCATE TABLE saving;")

	return err
}
