import React from "react";

function DashboardTotals(props) {
  let percentIncomeRemaining = 0;
  if (props.incomeAmount > 0) {
    percentIncomeRemaining =
      (1 - props.totalExpensesAndSavings / props.incomeAmount) * 100;
  }

  return (
    <div className="dashboard-totals">
      <h4>Total Expenses & Saving</h4>
      <h5>£{props.totalExpensesAndSavings}</h5>
      <h4>% of Monthly Income Remaining</h4>
      <h5>{percentIncomeRemaining.toFixed(2)}%</h5>
    </div>
  );
}

export default DashboardTotals;
