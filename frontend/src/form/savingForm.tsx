import * as React from "react";

interface SavingFormProps {
  savingPercent: number;
  setSavingPercent: (savingPercent: number) => void;
}

function SavingForm(props: SavingFormProps): JSX.Element {
  return (
    <>
      <h1>Savings</h1>
      <div
        className="border flex flex-row items-center space-y-2"
        id="budget-saving"
      >
        <table>
          <thead>
            <tr>
              <th scope="col">Saving Percent</th>
            </tr>
          </thead>
          <tbody>
            <th scope="row">
              <input
                className=""
                id="saving-percent"
                type="numeric"
                value={props.savingPercent}
                onChange={(event) => {
                  if (event.target.value === "") {
                    props.setSavingPercent(0);
                  } else {
                    props.setSavingPercent(parseInt(event.target.value));
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

export default SavingForm;
