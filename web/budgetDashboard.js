window.addEventListener('load', function() {
  loadDashboardData();
});

function loadDashboardData() {
  _getDashboardData(_showDashboardData);
}

function _getDashboardData(callback) {

  let month = document.getElementById('date-month').value;
  let year = document.getElementById('date-year').value;

  let xhttp = new XMLHttpRequest();
  let path = '/api/v1/budget/dashboard?date=' + _formatDateForQuery(month, year);
  xhttp.onreadystatechange = function() {
    if (xhttp.readyState === 4 && xhttp.status === 200) {
      let dashboard = _convertToDashboardData(xhttp.responseText);
      callback(dashboard);
    }
  };
  xhttp.open("GET", path, true);
  xhttp.send();

};

function _showDashboardData(dashboard) {
  // Saving
  document.getElementById('saving-percent').value = dashboard.Saving.Percent;
  document.getElementById('budget-dashboard-monthly-saving-percent').innerHTML = dashboard.Saving.Percent + "%";
  document.getElementById('budget-dashboard-monthly-saving').innerHTML = "£" + dashboard.SavingTotal;

  // Income
  document.getElementById('income-source').value = dashboard.Income.Source;
  document.getElementById('income-amount').value = dashboard.Income.Amount;
  document.getElementById('budget-dashboard-monthly-income').innerHTML = "£" + dashboard.Income.Amount;
}


function _convertToDashboardData(response) {
  let resDashboard = JSON.parse(response);

  return resDashboard;
}


function _formatDateForJSON(month, year) {
  return year + "-" + month + "-" + "01" + "T00:00:00Z";
}

function _formatDateForQuery(month, year) {
  return year + "-" + month + "-" + "01";
}