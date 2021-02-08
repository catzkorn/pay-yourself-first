import * as React from "react";
import { useEffect } from "react";

function ListIncome(): JSX.Element {
  useEffect(() => {
    fetch("/api/v1/income")
      .then((response) => {
        return response.json();
      })
      .then((payload) => {
        console.log(payload);
      });
  });

  return <p>Hello I'm a cow</p>;
}

export default ListIncome;
