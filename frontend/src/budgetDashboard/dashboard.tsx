import * as React from "react";
import DashboardDate from "./dashboardDate";
import BudgetIncome from "./dashboardIncome";
import BudgetSaving from "./dashboardSaving";
import DashboardTotals from "./dashboardTotals";

interface DashboardProps {
  incomeAmount: number;
  savingPercent: number;
  expensesAmount: number;
  month: string;
  year: string;
}

function Dashboard(props: DashboardProps): JSX.Element {
  const savingTotal = props.incomeAmount * (props.savingPercent / 100);
  const totalExpensesAndSavings = savingTotal + props.expensesAmount;
  const flexibleSpending = props.incomeAmount - totalExpensesAndSavings;

  return (
    <>
      <span>Budget Dashboard</span>
      <div
        className="px-4 py-6 sm:px-0 border-4 border-dashed border-gray-200 rounded-lg m-2 flex flex-row"
        id="budget-dashboard"
      >
        <BudgetIncome
          incomeAmount={props.incomeAmount}
          expensesAmount={props.expensesAmount}
          flexibleSpending={flexibleSpending}
        />

        <div className="flex flex-col" id="budget-dashboard-col-align">
          <DashboardDate month={props.month} year={props.year} />

          <DashboardTotals
            totalExpensesAndSavings={totalExpensesAndSavings}
            incomeAmount={props.incomeAmount}
          />
        </div>

        <BudgetSaving
          savingPercent={props.savingPercent}
          savingTotal={savingTotal}
        />
      </div>
    </>
  );
}

export default Dashboard;
