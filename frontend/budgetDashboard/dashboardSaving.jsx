import React from "react";

function BudgetSaving(props) {
  const monthlySavings = 1000;

  return (
    <div className="dashboard-saving">
      <h4>Monthly Savings</h4>
      <h5>{monthlySavings}</h5>
      <h4>% Savings of Income</h4>
      <h5>{props.savingPercent}%</h5>
    </div>
  );
}

export default BudgetSaving;
