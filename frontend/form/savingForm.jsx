import React from "react";

function SavingForm(props) {
  return (
    <div className="budget-saving">
      <h1>Savings</h1>

      <label for="saving-percent">Percent</label>
      <input
        onChange={(event) => {
          props.setSavingPercent(event.target.value);
        }}
        type="numeric"
        value={props.savingPercent}
        name="saving-percent"
      />
    </div>
  );
}

export default SavingForm;
