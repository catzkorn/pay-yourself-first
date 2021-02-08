import * as React from "react";

interface DateFormProps {
  setMonth: (month: string) => void;
  setYear: (year: string) => void;
  month: string;
  year: string;
}

function DateForm(props: DateFormProps): JSX.Element {
  function handleMonthChange(event: React.ChangeEvent<HTMLSelectElement>) {
    props.setMonth(event.target.value);
  }

  function handleYearChange(event: React.ChangeEvent<HTMLSelectElement>) {
    props.setYear(event.target.value);
  }

  return (
    <>
      <h2>Budget Date</h2>
      <div className="border m-6 flex flex-row" id="budget-date-selection">
        <table>
          <thead>
            <tr>
              <th scope="col">Month</th>
              <th scope="col">Year</th>
            </tr>
          </thead>
          <tbody>
            <th scope="col">
              <select
                className=""
                id="date-month"
                value={props.month}
                onChange={(event) => handleMonthChange(event)}
              >
                <option value="01">January</option>
                <option value="02">February</option>
                <option value="03">March</option>
                <option value="04">April</option>
                <option value="05">May</option>
                <option value="06">June</option>
                <option value="07">July</option>
                <option value="08">August</option>
                <option value="09">September</option>
                <option value="10">October</option>
                <option value="11">November</option>
                <option value="12">December</option>
              </select>
            </th>
            <th scope="col">
              <select
                className=""
                id="date-year"
                value={props.year}
                onChange={(event) => handleYearChange(event)}
              >
                <option value="2020">2020</option>
                <option value="2021">2021</option>
                <option value="2022">2022</option>
                <option value="2023">2023</option>
              </select>
            </th>
          </tbody>
        </table>
      </div>
    </>
  );
}

export default DateForm;
