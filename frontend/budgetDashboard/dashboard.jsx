import React from "react";
import BudgetIncome from "./dashboardIncome";
import BudgetSaving from "./dashboardSaving";
import DashboardTotals from "./dashboardTotals";
import DateForm from "./dateForm";

function Dashboard(props) {
  return (
    <div className="budget-dashboard">
      <span>Budget Dashboard</span>

      <BudgetIncome
        incomeAmount={props.incomeAmount}
        expensesAmount={props.expensesAmount}
      />
      <DateForm />
      <BudgetSaving savingPercent={props.savingPercent} />
      <DashboardTotals />
    </div>
  );
}

export default Dashboard;
