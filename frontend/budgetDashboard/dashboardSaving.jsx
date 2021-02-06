import React from "react";

function BudgetSaving(props) {
  return (
    <div className="dashboard-saving">
      <h4>Monthly Savings</h4>
      <h5>£{props.savingTotal}</h5>
      <h4>% Savings of Income</h4>
      <h5>{props.savingPercent}%</h5>
    </div>
  );
}

export default BudgetSaving;
