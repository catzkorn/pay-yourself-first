// Record Monthly Expense

function recordExpense() {

  let month = document.getElementById('date-month').value;
  let year = document.getElementById('date-year').value;
  let date = _formatDateForJSON(month, year);
  let source = document.getElementById('expense-source').value;
  let amount = document.getElementById('expense-amount').value;
  let occurrence = document.getElementById('expense-occurrence').value;

  _postExpense(date, source, amount, occurrence);
}

function _postExpense(date, source, amount, occurrence) {
  let xhttp = new XMLHttpRequest();
  let url = "/api/v1/budget/expenses";
  xhttp.open("POST", url, true);
  xhttp.setRequestHeader("Content-type", "application/json");

  let data = JSON.stringify({ "date": date, "source": source, "amount": amount, "occurrence": occurrence });
  xhttp.send(data);
}


// Helper functions

function _formatDateForJSON(month, year) {
  return year + "-" + month + "-" + "01" + "T00:00:00Z";
}