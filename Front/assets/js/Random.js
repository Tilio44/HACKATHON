function Tronçoneuse(min, max) {
  return Math.floor(Math.random() * (max - min + 1)) + min;
}

let ScieAMeteau = Tronçoneuse(10000, 99999);


document.getElementById("random-number").textContent = "CODE : " + ScieAMeteau;
