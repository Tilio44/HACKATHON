function randomInt(min, max) {
  return Math.floor(Math.random() * (max - min + 1)) + min;
}

let randomNumber = randomInt(10000, 99999);


document.getElementById("random-number").textContent = "CODE : " + randomNumber;