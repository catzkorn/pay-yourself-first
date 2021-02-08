import * as React from "react";

interface DashboardTotalsProps {
  totalExpensesAndSavings: number;
  incomeAmount: number;
}

function DashboardTotals(props: DashboardTotalsProps): JSX.Element {
  let percentIncomeRemaining = 0;
  if (props.incomeAmount > 0) {
    percentIncomeRemaining =
      (1 - props.totalExpensesAndSavings / props.incomeAmount) * 100;
  }

  return (
    <div
      className="border m-2 flex flex-col items-center space-y-2"
      id="budget-dashboard-totals"
    >
      <h3>Totals</h3>
      <h4>Expenses & Saving</h4>
      <h5>£{props.totalExpensesAndSavings}</h5>
      <h4>% Income Remaining</h4>
      <h5>{percentIncomeRemaining.toFixed(2)}%</h5>
    </div>
  );
}

export default DashboardTotals;
