import React from "react";
import BudgetSaving from "./dashboardSaving";

function BudgetIncome(props) {
  const monthlyFlexibleSpending = 1000;
  const monthlyExpensesOfIncome = 60;

  return (
    <div className="dashboard-income">
      <h4>Monthly Income</h4>
      <h5>£{props.incomeAmount}</h5>
      <h4>Monthly Expenses</h4>
      <h5>{props.expensesAmount}</h5>
      <h4>Monthly Flexible Spending</h4>
      <h5>{monthlyFlexibleSpending}</h5>
      <h4>% Expenses of Income</h4>
      <h5>{monthlyExpensesOfIncome}%</h5>
    </div>
  );
}

export default BudgetIncome;
