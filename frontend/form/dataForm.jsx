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
      });
  }, [month, year]);

  function _formatDateForQuery(month, year) {
    return year + "-" + month + "-" + "01";
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
        />

        <button
          type="button"
          onclick="recordIncome(); recordSaving(); recordExpense();loadDashboardData();"
        >
          Save
        </button>
      </form>
    </div>
  );
}

export default Form;
