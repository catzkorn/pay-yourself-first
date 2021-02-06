import React from "react";
import DateForm from "./dateForm";
import ExpensesForm from "./expensesForm";
import IncomeForm from "./incomeForm";
import SavingForm from "./savingForm";
function Form(props) {
  return (
    <div className="input-forms">
      <form>
        <DateForm />

        <IncomeForm
          incomeAmount={props.incomeAmount}
          setIncomeAmount={props.setIncomeAmount}
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
