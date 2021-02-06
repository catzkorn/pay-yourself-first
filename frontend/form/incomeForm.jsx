import React from "react";

function IncomeForm(props) {
  const incomeSource = "Salary";

  return (
    <div className="budget-income">
      <h1>Income</h1>
      <label for="income-source">Type</label>
      <input type="text" value={incomeSource} name="income-source" />
      <label for="income-amount">Amount</label>
      <input
        onChange={(event) => {
          props.setIncomeAmount(event.target.value);
        }}
        type="numeric"
        value={props.incomeAmount}
        name="income-amount"
      />
    </div>
  );
}

export default IncomeForm;
