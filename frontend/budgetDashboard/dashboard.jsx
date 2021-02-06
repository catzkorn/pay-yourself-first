import React from "react";
import BudgetIncome from "./dashboardIncome";
import BudgetSaving from "./dashboardSaving";
import DashboardTotals from "./dashboardTotals";
import DateForm from "../form/dateForm";

function Dashboard(props) {
  const savingTotal = props.incomeAmount * (props.savingPercent / 100);
  const totalExpensesAndSavings = savingTotal + props.expensesAmount;
  const flexibleSpending = props.incomeAmount - totalExpensesAndSavings;

  return (
    <div className="budget-dashboard">
      <span>Budget Dashboard</span>

      <BudgetIncome
        incomeAmount={props.incomeAmount}
        expensesAmount={props.expensesAmount}
        flexibleSpending={flexibleSpending}
      />

      <BudgetSaving
        savingPercent={props.savingPercent}
        savingTotal={savingTotal}
      />
      <DashboardTotals
        totalExpensesAndSavings={totalExpensesAndSavings}
        incomeAmount={props.incomeAmount}
      />
    </div>
  );
}

export default Dashboard;
