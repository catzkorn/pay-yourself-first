import * as React from "react";
import { useEffect, useState } from "react";
import formatDateAsDay from "../util/formatDate";
import formatAmountTwoDecimals from "../util/formatDecimalAmount";
import IncomeModal from "./incomeModal";
import Income from "./incomeTypes";

function ListIncome(): JSX.Element {
  const [incomes, setIncomes] = useState<Income[]>([]);

  useEffect(() => {
    loadIncomes(setIncomes);
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
            return (
              <Income
                income={income}
                key={income.ID}
                incomes={incomes}
                setIncomes={setIncomes}
              />
            );
          }
        )}
      </tbody>
    </table>
  );
}

interface IncomeProps {
  income: Income;
  incomes: Income[];
  setIncomes: (incomes: Income[]) => void;
}

function Income(props: IncomeProps) {
  const [showModal, setShowModal] = useState<boolean>(false);

  function handleDeleteIncome(id: number) {
    const url = "/api/v1/income/" + String(id);
    console.log(id);
    console.log(url);
    const options = {
      method: "DELETE",
    };
    fetch(url, options).then((response) => {
      if (response.status !== 200) {
        console.log("There was an error deleting the request", response);
      }
      let newIncomes = props.incomes.filter((income) => {
        return income.ID !== id;
      });
      props.setIncomes(newIncomes);
    });
  }

  return (
    <tr>
      <td>{formatDateAsDay(props.income.Date)}</td>
      <td>{props.income.Source}</td>
      <td>£ {formatAmountTwoDecimals(props.income.Amount)}</td>
      <td>
        <button
          className="border p-1 rounded hover:shadow-md"
          id="delete-income-button"
          type="button"
          onClick={(event) => {
            event.preventDefault();
            setShowModal(true);
          }}
        >
          Delete
        </button>

        <IncomeModal
          income={props.income}
          deleteIncome={() => handleDeleteIncome(props.income.ID)}
          setShowModal={setShowModal}
          showModal={showModal}
        />
      </td>
    </tr>
  );
}

function loadIncomes(setIncomes: (incomes: Income[]) => void) {
  const url = "/api/v1/income/";
  fetch(url)
    .then((response) => {
      return response.json();
    })
    .then((payload) => {
      setIncomes(payload);
    });
}

export default ListIncome;
