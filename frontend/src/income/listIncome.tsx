import * as React from "react";
import { useEffect, useState } from "react";
import formatDateAsDay from "../util/formatDate";
import formatAmountTwoDecimals from "../util/formatDecimalAmount";
import Income from "./incomeTypes";

function ListIncome(): JSX.Element {
  const [incomes, setIncomes] = useState<Income[]>([]);

  useEffect(() => {
    const url = "/api/v1/income";
    fetch(url)
      .then((response) => {
        return response.json();
      })
      .then((payload) => {
        setIncomes(payload);
      });
  }, []);

  return (
    <table className="" id="list-income-table">
      <thead>
        <tr>
          <th scope="col">Date</th>
          <th scope="col">Source</th>
          <th scope="col">Amount</th>
        </tr>
      </thead>
      <tbody>
        {incomes.map(
          (income): JSX.Element => {
            return <Income income={income} key={income.ID} />;
          }
        )}
      </tbody>
    </table>
  );
}

interface IncomeProps {
  income: Income;
}

function Income(props: IncomeProps) {
  return (
    <tr>
      <td>{formatDateAsDay(props.income.Date)}</td>
      <td>{props.income.Source}</td>
      <td>£ {formatAmountTwoDecimals(props.income.Amount)}</td>
    </tr>
  );
}

export default ListIncome;
