import * as React from "react";

interface IncomeFormProps {
  incomeAmount: number;
  setIncomeAmount: (incomeAmount: number) => void;
  incomeType: string;
  setIncomeType: (incomeType: string) => void;
}

function IncomeForm(props: IncomeFormProps): JSX.Element {
  return (
    <div className="budget-income">
      <h1>Income</h1>
      <label htmlFor="income-source">Type</label>
      <input
        type="text"
        value={props.incomeType}
        name="income-source"
        onChange={(event) => {
          if (event.target.value === "") {
            props.setIncomeType("");
          } else {
            props.setIncomeType(event.target.value);
          }
        }}
      />
      <label htmlFor="income-amount">Amount</label>
      <input
        onChange={(event) => {
          if (event.target.value === "") {
            props.setIncomeAmount(0);
          } else {
            props.setIncomeAmount(parseInt(event.target.value));
          }
        }}
        type="numeric"
        value={props.incomeAmount}
        name="income-amount"
      />
    </div>
  );
}

export default IncomeForm;
