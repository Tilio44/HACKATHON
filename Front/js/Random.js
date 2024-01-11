function randomInt(min, max) {
  return Math.floor(Math.random() * (max - min + 1)) + min;
}

// Générer un nombre aléatoire à 5 chiffres
let randomNumber = randomInt(10000, 99999);

// Sélectionner l'élément avec l'ID "random-number" et mettre le nombre aléatoire à l'intérieur
document.getElementById("random-number").textContent = randomNumber;
