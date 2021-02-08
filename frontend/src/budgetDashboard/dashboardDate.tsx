import * as React from "react";

interface DashboardDateProps {
  month: string;
  year: string;
}

function DashboardDate(props: DashboardDateProps): JSX.Element {
  return (
    <div
      className="border m-2 flex flex-col items-center space-y-2"
      id="budget-dashboard-totals"
    >
      <h4>Date</h4>
      <h5>
        {props.month}/{props.year}
      </h5>
    </div>
  );
}

export default DashboardDate;
