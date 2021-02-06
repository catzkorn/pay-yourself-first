import React, { useState } from "react";

function DateForm(props) {
  const [month, setMonth] = useState("01");
  const [year, setYear] = useState("2021");

  function handleMonthChange(event) {
    setMonth(event.target.value);
  }

  function handleYearChange(event) {
    setYear(event.target.value);
  }

  return (
    <>
      <h2>Select Date:</h2>
      <form onclick="loadDashboardData()">
        <label for="date-month">Month:</label>
        <select
          onChange={(event) => handleMonthChange(event)}
          name="date-month"
          value={month}
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

        <label for="date-year">Year:</label>
        <select
          onChange={(event) => handleYearChange(event)}
          name="date-year"
          value={year}
        >
          <option value="2020">2020</option>
          <option value="2021">2021</option>
          <option value="2022">2022</option>
          <option value="2023">2023</option>
        </select>
      </form>
    </>
  );
}

export default DateForm;
