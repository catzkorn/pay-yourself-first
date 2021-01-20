

function recordIncome() {

  let month = document.getElementById('date-month').value;
  let year = document.getElementById('date-year').value;
  let date = _formatDateForJSON(month, year);
  let source = document.getElementById('income-source').value;
  let amount = document.getElementById('income-amount').value;

  _postIncome(date, source, amount);
}

function _postIncome(date, source, amount) {
  let xhttp = new XMLHttpRequest();
  let url = "/api/v1/budget";
  xhttp.open("POST", url, true);
  xhttp.setRequestHeader("Content-type", "application/json");

  let data = JSON.stringify({ "date": date, "source": source, "amount": amount });
  xhttp.send(data);
}

function _formatDateForJSON(month, year) {
  return year + "-" + month + "-" + "01" + "T00:00:00Z";
}



