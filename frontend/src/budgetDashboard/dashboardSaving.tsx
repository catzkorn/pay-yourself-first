import * as React from "react";

interface BudgetSavingProps {
  savingTotal: number;
  savingPercent: number;
}

function BudgetSaving(props: BudgetSavingProps): JSX.Element {
  return (
    <div
      className="border m-2 flex flex-col items-center space-y-2"
      id="budget-dashboard-saving"
    >
      <h4>Savings</h4>
      <h5>£{props.savingTotal}</h5>
      <h4>% Savings of Income</h4>
      <h5>{props.savingPercent}%</h5>
    </div>
  );
}

export default BudgetSaving;
