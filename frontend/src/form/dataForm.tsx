import * as React from "react";
import { useEffect, useState } from "react";
import DateForm from "./dateForm";
import ExpensesForm from "./expensesForm";
import IncomeForm from "./incomeForm";
import SavingForm from "./savingForm";

interface FormProps {
  setIncomeAmount: (incomeAmount: number) => void;
  setSavingPercent: (savingPercent: number) => void;
  setExpensesAmount: (expensesAmount: number) => void;
  setYear: (year: string) => void;
  setMonth: (month: string) => void;
  month: string;
  year: string;
  savingPercent: number;
  incomeAmount: number;
  expensesAmount: number;
}

function Form(props: FormProps): JSX.Element {
  const [incomeType, setIncomeType] = useState<string>("");
  const [expensesType, setExpensesType] = useState<string>("");
  const [expensesOccurrence, setExpensesOccurrence] = useState<string>(
    "monthly"
  );

  useEffect(() => {
    fetch(
      "/api/v1/budget/dashboard?date=" +
        _formatDateForQuery(props.month, props.year)
    )
      .then((response) => {
        return response.json();
      })
      .then((payload) => {
        props.setIncomeAmount(parseInt(payload.Income.Amount));
        setIncomeType(payload.Income.Source);
        props.setSavingPercent(parseInt(payload.Saving.Percent));
        props.setExpensesAmount(parseInt(payload.Expense.Amount));
        setExpensesType(payload.Expense.Source);
        setExpensesOccurrence(payload.Expense.Occurrence);
      });
  }, [props.month, props.year]);

  function handleSavingSubmit() {
    const url = "/api/v1/budget/saving";
    const date = _formatDateForJSON(props.month, props.year);
    const options = {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ percent: props.savingPercent, date: date }),
    };

    fetch(url, options).then((response) => {
      if (response.status !== 200) {
        console.log("There was an error with the submitted data", response);
      }
    });
  }

  function handleIncomeSubmit() {
    const url = "/api/v1/budget/income";
    const date = _formatDateForJSON(props.month, props.year);
    const options = {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        date: date,
        source: incomeType,
        amount: props.incomeAmount,
      }),
    };

    fetch(url, options).then((response) => {
      if (response.status !== 200) {
        console.log("There was an error with the submitted data", response);
      }
    });
  }

  function handleExpensesSubmit() {
    const url = "/api/v1/budget/expenses";
    const date = _formatDateForJSON(props.month, props.year);
    const options = {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        date: date,
        source: expensesType,
        amount: props.expensesAmount,
        occurrence: expensesOccurrence,
      }),
    };

    fetch(url, options).then((response) => {
      if (response.status !== 200) {
        console.log("There was an error with the submitted data", response);
      }
    });
  }

  function _formatDateForQuery(month: string, year: string): string {
    return year + "-" + month + "-" + "01";
  }

  function _formatDateForJSON(month: string, year: string): string {
    return year + "-" + month + "-" + "01" + "T00:00:00Z";
  }

  return (
    <div
      className="px-4 py-6 sm:px-0 border-4 border-dashed border-gray-200 rounded-lg m-2 flex flex-row grid justify-items-center"
      id="budget-input-forms"
    >
      <form className="">
        <DateForm
          month={props.month}
          setMonth={props.setMonth}
          year={props.year}
          setYear={props.setYear}
        />

        <IncomeForm
          incomeAmount={props.incomeAmount}
          setIncomeAmount={props.setIncomeAmount}
          incomeType={incomeType}
          setIncomeType={setIncomeType}
        />

        <SavingForm
          savingPercent={props.savingPercent}
          setSavingPercent={props.setSavingPercent}
        />

        <ExpensesForm
          expensesAmount={props.expensesAmount}
          setExpensesAmount={props.setExpensesAmount}
          expensesType={expensesType}
          setExpensesType={setExpensesType}
          expensesOccurrence={expensesOccurrence}
          setExpensesOccurrence={setExpensesOccurrence}
        />

        <button
          type="button"
          onClick={(event) => {
            event.preventDefault();
            handleSavingSubmit();
            handleIncomeSubmit();
            handleExpensesSubmit();
          }}
        >
          Save
        </button>
      </form>
    </div>
  );
}

export default Form;
