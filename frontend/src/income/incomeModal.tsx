import * as React from "react";
import { useEffect, useState } from "react";
import formatDateAsDay from "../util/formatDate";
import Income from "./incomeTypes";

interface IncomeModalProps {
  income: Income;
  deleteIncome: () => void;
  setShowModal: (showModal: boolean) => void;
  showModal: boolean;
}

function IncomeModal(props: IncomeModalProps): JSX.Element | null {
  if (!props.showModal) {
    return null;
  }

  return (
    <div className="fixed inset-0 z-50 overflow-auto bg-gray-500 flex">
      <div className="relative p-8 bg-white w-full max-w-md m-auto flex-col flex rounded-lg">
        <span className="absolute top-0 right-0 p-4">
          <button onClick={() => props.setShowModal(false)}>x</button>
        </span>

        <div className="bg-white px-4 pt-5 pb-4 sm:p-6 sm:pb-4">
          Are you sure you want to delete the income from "{props.income.Source}
          " for {formatDateAsDay(props.income.Date)}?
        </div>
        <div
          className="bg-gray-50 px-4 py-3 sm:px-6 sm:flex sm:flex-row-reverse"
          id="delete-or-cancel-buttons"
        >
          <button
            className="w-full inline-flex justify-center rounded-md border border-transparent shadow-sm px-4 py-2 bg-red-600 text-base font-medium text-white hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500 sm:ml-3 sm:w-auto sm:text-sm"
            onClick={() => props.deleteIncome()}
          >
            Delete
          </button>
          <button
            className="mt-3 w-full inline-flex justify-center rounded-md border border-gray-300 shadow-sm px-4 py-2 bg-white text-base font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 sm:mt-0 sm:ml-3 sm:w-auto sm:text-sm"
            onClick={() => props.setShowModal(false)}
          >
            Cancel
          </button>
        </div>
      </div>
    </div>
  );
}

export default IncomeModal;
