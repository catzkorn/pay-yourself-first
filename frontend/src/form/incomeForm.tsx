import * as React from "react";

interface IncomeFormProps {
  incomeAmount: number;
  setIncomeAmount: (incomeAmount: number) => void;
  incomeType: string;
  setIncomeType: (incomeType: string) => void;
}

function IncomeForm(props: IncomeFormProps): JSX.Element {
  return (
    <>
      <h1>Income</h1>
      <div
        className="border flex flex-row items-center space-y-2"
        id="budget-income"
      >
        <table className="" id="budget-income-table">
          <thead>
            <tr>
              <th scope="col">Type</th>
              <th scope="col">Amount</th>
            </tr>
          </thead>
          <tbody>
            <th scope="col">
              <input
                className="focus:ring-indigo-500 focus:border-indigo-500 flex-1 block w-full rounded-none rounded-r-md sm:text-sm border-gray-300"
                id="income-source"
                type="text"
                value={props.incomeType}
                onChange={(event) => {
                  if (event.target.value === "") {
                    props.setIncomeType("");
                  } else {
                    props.setIncomeType(event.target.value);
                  }
                }}
              />
            </th>
            <th>
              <input
                className="focus:ring-indigo-500 focus:border-indigo-500 flex-1 block w-full rounded-none rounded-r-md sm:text-sm border-gray-300"
                id="income-amount"
                type="numeric"
                value={props.incomeAmount}
                onChange={(event) => {
                  if (event.target.value === "") {
                    props.setIncomeAmount(0);
                  } else {
                    props.setIncomeAmount(parseInt(event.target.value));
                  }
                }}
              />
            </th>
          </tbody>
        </table>
      </div>
    </>
  );
}

export default IncomeForm;
