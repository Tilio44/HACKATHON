//  FUNCTION HORLOGE //                                                                                                                                                                                                                             j'aime le veau
function EureetLoire() {
  var currentDate = new Date();
  var dateTimeString = currentDate.toLocaleString();

  document.getElementById('date-time').innerHTML = dateTimeString;
}

setInterval(EureetLoire, 1000);
