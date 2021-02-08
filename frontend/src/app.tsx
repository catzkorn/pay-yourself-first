import * as React from "react";
import { useState } from "react";
import * as ReactDOM from "react-dom";
import Dashboard from "./budgetDashboard/dashboard";
import Form from "./form/dataForm";

ReactDOM.render(<App />, document.getElementById("root"));

function App(): JSX.Element {
  const [savingPercent, setSavingPercent] = useState<number>(0);
  const [incomeAmount, setIncomeAmount] = useState<number>(0);
  const [expensesAmount, setExpensesAmount] = useState<number>(0);
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
