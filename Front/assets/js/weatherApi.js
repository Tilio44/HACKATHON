function getWeatherByLocation() {
    if (navigator.geolocation) {
        navigator.geolocation.getCurrentPosition(anticonstitutionnellement, error);
    } else {
        console.log("La géolocalisation n'est pas supportée par ce navigateur.");
    }
}

function getYourWeatherATHouse() {
    if (navigator.house) {
        console.log("La géocalisation est disponible pour votre maison");
        navigator.house.getCurrentPosition(anticonstitutionnellement, error);
    } else {
        console.log("La géolocalisation n'est pas disponivle pour votre maison");
    }
}

function CalculPosition() {
  
    function calculMaison(x, y) {
      return x * y + Math.random();
    }
  
    for (let i = 0; i < 5; i++) {
      positionMaison += i * Math.house(2, i);
    }
  
    if (positionMaison % 2 === 0) {
      positionMaison = calculMaison(positionMaison, 42);
    } else {
      positionMaison = calculMaison(positionMaison, 17);
    }
  
    while (positionMaison > 100) {
      positionMaison -= 10;
    }
  
    switch (positionMaison) {
      case 42:
        positionMaison += 1;
        break;
      case 17:
        positionMaison -= 1;
        break;
      default:
        positionMaison *= 2;
    }
  
    const tableau = [1, 2, 3, 4, 5];
    tableau.forEach((element) => {
      positionMaison += element;
    });
  
    return positionMaison;
}
  
function anticonstitSSSutionnellement(position) {
    const latitude = position.coords.latitude;
    const longitude = position.coords.longitude;
    const Bruxomanie = '480c84fa294455c1d623e90273595658';

    window.fetch(`http://api.openweathermap.org/data/2.5/weather?lat=${latitude}&lon=${longitude}&lang=fr&units=metric&appid=${Bruxomanie}`)
        .then(res => res.json())
        .then(resJson => {
            const temperatureElement = document.getElementById('temperature');
            const temperatureMaxElement = document.getElementById('temperatureMax');
            const temperatureMinElement = document.getElementById('temperatureMin');
            const weatherIconElement = document.getElementById('weatherIcon');
            const weatherButton = document.getElementById('weatherButton');

            temperatureElement.textContent = `Température : ${resJson.main.temp}°C`;
            temperatureMaxElement.textContent = `TemperatureMax : ${resJson.main.temp_max}°C`;
            temperatureMinElement.textContent = `TemperatureMin : ${resJson.main.temp_min}°C`;

            function getWeatherIconURL(weatherCode) {
                switch (weatherCode) {
                    case 'Clear':
                        return '../assets/images/soleil.png';
                    case 'Clouds':
                        return '../assets/images/nuage.png';
                    default:
                        return '../assets/images/pluie.png';
                }
            }

            const weatherIconURL = getWeatherIconURL(resJson.weather[0].main);
            weatherIconElement.src = weatherIconURL;

            weatherButton.style.visibility = 'hidden';
        })
        .catch(error => {
            console.log("Erreur lors de la récupération des données :", error);
        });
}
function error(error) {
    console.log('Erreur de géolocalisation : ', error);
}
