
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

function _formatDateForJSON(month, year) {
  return year + "-" + month + "-" + "01" + "T00:00:00Z";
}

