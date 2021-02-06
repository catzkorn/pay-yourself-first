import React, { useEffect, useState } from "react";
import DateForm from "./dateForm";
import ExpensesForm from "./expensesForm";
import IncomeForm from "./incomeForm";
import SavingForm from "./savingForm";

function Form(props) {
  const [month, setMonth] = useState("01");
  const [year, setYear] = useState("2021");
  const [incomeType, setIncomeType] = useState("");
  const [expensesType, setExpensesType] = useState("");
  const [expensesOccurrence, setExpensesOccurrence] = useState("monthly");

  useEffect(() => {
    fetch("/api/v1/budget/dashboard?date=" + _formatDateForQuery(month, year))
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
  }, [month, year]);

  function handleSavingSubmit() {
    const url = "/api/v1/budget/saving";
    const date = _formatDateForJSON(month, year);
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
    const date = _formatDateForJSON(month, year);
    const options = {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        date: date,
        source: props.incomeType,
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
    const date = _formatDateForJSON(month, year);
    const options = {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        date: date,
        source: props.expensesType,
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

  function _formatDateForQuery(month, year) {
    return year + "-" + month + "-" + "01";
  }

  function _formatDateForJSON(month, year) {
    return year + "-" + month + "-" + "01" + "T00:00:00Z";
  }

  return (
    <div className="input-forms">
      <form>
        <DateForm
          month={month}
          setMonth={setMonth}
          year={year}
          setYear={setYear}
        />

        <IncomeForm
          incomeAmount={props.incomeAmount}
          setIncomeAmount={props.setIncomeAmount}
          incomeType={incomeType}
          setIncomeType={setIncomeType}
        />
        <br />
        <SavingForm
          savingPercent={props.savingPercent}
          setSavingPercent={props.setSavingPercent}
        />

        <br />

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
