function Eure() {
  var currentDate = new Date();
  var dateTimeString = currentDate.toLocaleString();

  document.getElementById('date-time').innerHTML = dateTimeString;
}

setInterval(Eure, 1000);
