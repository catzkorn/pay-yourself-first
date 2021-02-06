import React, { useState } from "react";
import ReactDOM from "react-dom";
import Dashboard from "./budgetDashboard/dashboard";
import Form from "./form/dataForm";

ReactDOM.render(<App />, document.getElementById("root"));

function App(props) {
  const [savingPercent, setSavingPercent] = useState(0);
  const [incomeAmount, setIncomeAmount] = useState(0);
  const [expensesAmount, setExpensesAmount] = useState(0);
  return (
    <>
      <Dashboard
        savingPercent={savingPercent}
        incomeAmount={incomeAmount}
        expensesAmount={expensesAmount}
      />

      <Form
        savingPercent={savingPercent}
        setSavingPercent={setSavingPercent}
        incomeAmount={incomeAmount}
        setIncomeAmount={setIncomeAmount}
        expensesAmount={expensesAmount}
        setExpensesAmount={setExpensesAmount}
      />
    </>
  );
}
