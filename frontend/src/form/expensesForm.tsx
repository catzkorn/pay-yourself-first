import * as React from "react";

interface ExpensesFormProps {
  expensesType: string;
  setExpensesType: (expensesType: string) => void;
  setExpensesAmount: (expensesAmount: number) => void;
  expensesAmount: number;
  setExpensesOccurrence: (expensesOccurrence: string) => void;
  expensesOccurrence: string;
}

function ExpensesForm(props: ExpensesFormProps): JSX.Element {
  return (
    <>
      <h1>Expenses</h1>
      <div
        className="border m-6 flex flex-row items-center space-y-2"
        id="budget-saving"
      >
        <table className="" id="expenses-table">
          <thead>
            <tr>
              <th scope="col">Type</th>
              <th scope="col">Amount</th>
              <th scope="col">Occurrence</th>
            </tr>
          </thead>
          <tbody>
            <th scope="row">
              {" "}
              <input
                type="text"
                value={props.expensesType}
                name="expense-source"
              />
            </th>
            <th scope="row">
              <input
                onChange={(event) => {
                  if (event.target.value === "") {
                    props.setExpensesAmount(0);
                  } else {
                    props.setExpensesAmount(parseInt(event.target.value));
                  }
                }}
                type="numeric"
                value={props.expensesAmount}
                name="expense-amount"
              />
            </th>
            <th scope="row">
              <select
                onChange={(event) => {
                  props.setExpensesOccurrence(event.target.value);
                }}
                value={props.expensesOccurrence}
                name="expense-occurrence"
              >
                <option value="monthly">Monthly</option>
                <option value="weekly">Weekly</option>
                <option value="yearly">Yearly</option>
                <option value="one-off">One-Off</option>
              </select>
            </th>
          </tbody>
        </table>
      </div>
    </>
  );
}

export default ExpensesForm;
