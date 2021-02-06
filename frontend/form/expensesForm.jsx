import React from "react";

function ExpensesForm(props) {
  const expenseType = "Mortgage";

  return (
    <div className="budget-saving">
      <h1>Expenses</h1>

      <label for="expense-source">Type</label>
      <input type="text" value={expenseType} name="expense-source" />
      <label for="expense-amount">Amount</label>
      <input
        onChange={(event) => {
          props.setExpensesAmount(event.target.value);
        }}
        type="numeric"
        value={props.expenseAmount}
        name="expense-amount"
      />
      <label for="expense-occurrence">Occurrence</label>
      <select name="expense-occurrence" id="expense-occurrence">
        <option value="monthly">Monthly</option>
        <option value="weekly">Weekly</option>
        <option value="yearly">Yearly</option>
        <option value="one-off">One-Off</option>
      </select>
    </div>
  );
}

export default ExpensesForm;
