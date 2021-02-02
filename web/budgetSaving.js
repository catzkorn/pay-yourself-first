window.addEventListener('load', function() {
  loadMonthSaving();
});
function recordSaving() {

  let month = document.getElementById('date-month').value;
  let year = document.getElementById('date-year').value;
  let date = _formatDateForJSON(month, year);

  let percent = parseInt(document.getElementById('saving-percent').value);

  _postSaving(percent, date);
}

function _postSaving(percent, date) {

  let xhttp = new XMLHttpRequest();
  let url = "/api/v1/budget/saving";
  xhttp.open("POST", url, true);
  xhttp.setRequestHeader("Content-type", "application/json");

  let data = JSON.stringify({ "percent": percent, "date": date });
  xhttp.send(data);
}

function loadMonthSaving() {
  _getMonthSaving(_showMonthSaving);
}

function _getMonthSaving(callback) {

  let month = document.getElementById('date-month').value;
  let year = document.getElementById('date-year').value;

  let xhttp = new XMLHttpRequest();
  let path = '/api/v1/budget/saving?date=' + _formatDateForQuery(month, year);
  xhttp.onreadystatechange = function() {
    if (xhttp.readyState === 4 && xhttp.status === 200) {
      let saving = _convertToSaving(xhttp.responseText);
      callback(saving);
    }
  };
  xhttp.open("GET", path, true);
  xhttp.send();

};

function _showMonthSaving(saving) {
  document.getElementById('saving-percent').value = saving.Percent;
}


function _convertToSaving(response) {
  let resSaving = JSON.parse(response);

  return resSaving;
}


function _formatDateForJSON(month, year) {
  return year + "-" + month + "-" + "01" + "T00:00:00Z";
}

function _formatDateForQuery(month, year) {
  return year + "-" + month + "-" + "01";
}