import * as React from "react";

interface ExpensesFormProps {
  expensesType: string;
  setExpensesType: (expensesType: string) => void;
  setExpensesAmount: (expensesAmount: number) => void;
  expensesAmount: number;
  setExpensesOccurrence: (expensesOccurrence: string) => void;
  expensesOccurrence: string;
}

function ExpensesForm(props: ExpensesFormProps): JSX.Element {
  return (
    <div className="budget-saving">
      <h1>Expenses</h1>

      <label htmlFor="expense-source">Type</label>
      <input type="text" value={props.expensesType} name="expense-source" />
      <label htmlFor="expense-amount">Amount</label>
      <input
        onChange={(event) => {
          if (event.target.value === "") {
            props.setExpensesAmount(0);
          } else {
            props.setExpensesAmount(parseInt(event.target.value));
          }
        }}
        type="numeric"
        value={props.expensesAmount}
        name="expense-amount"
      />
      <label htmlFor="expense-occurrence">Occurrence</label>
      <select
        onChange={(event) => {
          props.setExpensesOccurrence(event.target.value);
        }}
        value={props.expensesOccurrence}
        name="expense-occurrence"
      >
        <option value="monthly">Monthly</option>
        <option value="weekly">Weekly</option>
        <option value="yearly">Yearly</option>
        <option value="one-off">One-Off</option>
      </select>
    </div>
  );
}

export default ExpensesForm;
