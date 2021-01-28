

window.addEventListener('load', function() {
  loadMonthIncome();
});


// Record monthly incomes

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


// Load month incomes

function loadMonthIncome() {
  _getMonthIncome(_showMonthIncome);
}

function _getMonthIncome(callback) {

  let month = document.getElementById('date-month').value;
  let year = document.getElementById('date-year').value;

  let xhttp = new XMLHttpRequest();
  let path = '/api/v1/budget?date=' + _formatDateForQuery(month, year);
  xhttp.onreadystatechange = function() {
    if (xhttp.readyState === 4 && xhttp.status === 200) {
      let income = _convertToIncome(xhttp.responseText);
      callback(income);
    }
  };
  xhttp.open("GET", path, true);
  xhttp.send();
}

function _convertToIncome(response) {
  let resIncome = JSON.parse(response);

  return resIncome;
}



function _showMonthIncome(income) {
  document.getElementById('income-source').value = income.Source;
  document.getElementById('income-amount').value = income.Amount;
}




// Helper functions

function _formatDateForJSON(month, year) {
  return year + "-" + month + "-" + "01" + "T00:00:00Z";
}

function _formatDateForQuery(month, year) {
  return year + "-" + month + "-" + "01";
}
