import React from "react";

function ExpensesForm(props) {
  return (
    <div className="budget-saving">
      <h1>Expenses</h1>

      <label for="expense-source">Type</label>
      <input type="text" value={props.expensesType} name="expense-source" />
      <label for="expense-amount">Amount</label>
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
      <label for="expense-occurrence">Occurrence</label>
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
