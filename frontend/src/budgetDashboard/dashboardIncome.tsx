import * as React from "react";

interface BudgetIncomeProps {
  incomeAmount: number;
  expensesAmount: number;
  flexibleSpending: number;
}

function BudgetIncome(props: BudgetIncomeProps): JSX.Element {
  let expensePercent = 0;
  if (props.incomeAmount > 0) {
    expensePercent = (props.expensesAmount / props.incomeAmount) * 100;
  }
  return (
    <div className="border m-2" id="budget-dashboard-income">
      <h4>Monthly Income</h4>
      <h5>£{props.incomeAmount}</h5>
      <h4>Monthly Expenses</h4>
      <h5>£{props.expensesAmount}</h5>
      <h4>Monthly Flexible Spending</h4>
      <h5>£{props.flexibleSpending}</h5>
      <h4>% Expenses of Income</h4>
      <h5>{expensePercent.toFixed(2)}%</h5>
    </div>
  );
}

export default BudgetIncome;
