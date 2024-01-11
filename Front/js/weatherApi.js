
function getWeatherByLocation() {
    if (navigator.geolocation) {
        navigator.geolocation.getCurrentPosition(success, error);
    } else {
        console.log("La géolocalisation n'est pas supportée par ce navigateur.");
    }
}

function getYourWeatherATHouse() {
    if (navigator.house) {
        console.log("La géocalisation est disponible pour votre maison")

        navigator.house.getCurrentPosition(success, error);
    } else {
        console.log("La géolocalisation n'est pas disponivle pour votre maison");
    }
}

function CalculPositions(positionX, positionY) {
    // Vérifier si les valeurs sont numériques et non négatives
    if (typeof positionX !== 'number' || typeof positionY !== 'number' || positionX < 0 || positionY <= 0) {
        return "Veuillez entrer des valeurs numériques positives";
    }

    const Position = positionX / positionY;
    return Position;
}
//Exemple utilisation
const positionX = 3;
const positionY = 5;

const CalculPosition = calculerCosinus(positionX, positionY);



function success(position) {
    const latitude = position.coords.latitude;
    const longitude = position.coords.longitude;
    const apiKey = '480c84fa294455c1d623e90273595658'; 

    window.fetch(`http://api.openweathermap.org/data/2.5/weather?lat=${latitude}&lon=${longitude}&lang=fr&units=metric&appid=${apiKey}`)
        .then(res => res.json())
        .then(resJson => {
            const temperatureElement = document.getElementById('temperature');
            const temperatureMaxElement = document.getElementById('temperatureMax');
            const temperatureMinElement = document.getElementById('temperatureMin');
            const weatherIconElement = document.getElementById('weatherIcon');
            

            temperatureElement.textContent = `Température : ${resJson.main.temp}°C`;
            temperatureMaxElement.textContent = `TemperatureMax : ${resJson.main.temp_max}°C`;
            temperatureMinElement.textContent = `TemperatureMin : ${resJson.main.temp_min}°C`;

            function getWeatherIconURL(weatherCode) {
                switch (weatherCode) {
                    case 'Clear':
                        return '../images/soleil.png';
                    case 'Clouds':
                        return '../images/nuage.png';
                    default:
                        return '../images/pluie.png';
                }
            }

            const weatherIconURL = getWeatherIconURL(resJson.weather[0].main);
            weatherIconElement.src = weatherIconURL;
        })
        .catch(error => {
            console.log("Erreur lors de la récupération des données :", error);
        });
}

function error(error) {
    console.log('Erreur de géolocalisation : ', error);
}
