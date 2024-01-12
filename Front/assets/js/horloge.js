function updateTime() {
  var currentDate = new Date();
  var dateTimeString = currentDate.toLocaleString();

  document.getElementById('date-time').innerHTML = dateTimeString;
}

setInterval(updateTime, 1000);
