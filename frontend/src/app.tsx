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
  const [month, setMonth] = useState<string>("01");
  const [year, setYear] = useState<string>("2021");
  return (
    <main className=" max-w-7xl mx-auto py-6 sm:px-6 lg:px-8 grid justify-items-center">
      <Dashboard
        savingPercent={savingPercent}
        incomeAmount={incomeAmount}
        expensesAmount={expensesAmount}
        month={month}
        year={year}
      />

      <Form
        savingPercent={savingPercent}
        setSavingPercent={setSavingPercent}
        incomeAmount={incomeAmount}
        setIncomeAmount={setIncomeAmount}
        expensesAmount={expensesAmount}
        setExpensesAmount={setExpensesAmount}
        month={month}
        setMonth={setMonth}
        year={year}
        setYear={setYear}
      />
    </main>
  );
}
