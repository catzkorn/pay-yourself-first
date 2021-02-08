import * as React from "react";

interface SavingFormProps {
  savingPercent: number;
  setSavingPercent: (savingPercent: number) => void;
}

function SavingForm(props: SavingFormProps): JSX.Element {
  return (
    <div className="budget-saving">
      <h1>Savings</h1>

      <label htmlFor="saving-percent">Percent</label>
      <input
        onChange={(event) => {
          if (event.target.value === "") {
            props.setSavingPercent(0);
          } else {
            props.setSavingPercent(parseInt(event.target.value));
          }
        }}
        type="numeric"
        value={props.savingPercent}
        name="saving-percent"
      />
    </div>
  );
}

export default SavingForm;
