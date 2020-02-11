"use strict";

window.addEventListener("load", () => {
  // let square = createChessBoard("abc", "#000000", "");

  // document.body.append(square);

  console.log(getSquareColor(0, 0));
  console.log(getSquareColor(0, 1));
  console.log(getSquareColor(1, 0));
  console.log(getSquareColor(1, 1));

  console.log("_____________________");

  console.log(getSquareColor(9, 9));
  console.log(getSquareColor(9, 8));
  console.log(getSquareColor(8, 9));
  console.log(getSquareColor(8, 8));

  console.log("_____________________");

  console.log(getSquareColor(1, 1));
  console.log(getSquareColor(1, 2));
  console.log(getSquareColor(1, 3));
  console.log(getSquareColor(1, 4));

  console.log("_____________________");

  console.log(getSquareColor(1, 1));
  console.log(getSquareColor(2, 1));
  console.log(getSquareColor(3, 1));
  console.log(getSquareColor(4, 1));
});



/**
 * Создать шахматную доску.
 */
function createChessBoard() {
  const board = document.createElement("div");

  for (let x = 0; x < 11; x++) {
    document.body.append(createChessLine(x));
  }
}



/**
 * Создать линию на шахматной доске.
 * @param {int} idx 
 * @returns {HTMLDivElement}
 */
function createBoardLine(idx) {
  

  for (let x = 0; x < 11; x++) {
  }
}



/**
 * Создать шахматную ячейку.
 * @param {int} x
 * @param {int} y
 * @returns {HTMLDivElement}
 */
function createLineSquare(x, y) {
  const size = "40px";
  const letters = ["a", "b", "c", "d", "e", "f", "g", "h"];
  const square = document.createElement("div");

  square.style.width = size;
  square.style.height = size;
  square.style.backgroundColor = getSquareColor(x, y);
  square.innerHTML = text;

  return square;
}



/**
 * Получить цвет шахматной ячейки.
 * @param {int} x
 * @param {int} y
 * @returns {string}
 */
function getSquareColor(x, y) {
  const wbColors = ["#fff", "#000"];
  const outerColor = "#aaa";

  // Не игровые клетки (с буквами)
  if (x < 1 || y < 1 || x > 8 || y > 8) {
    return outerColor;
  }

  // Нечетные игровые линии.
  if (y % 2) {
    return wbColors[(x + 1) % 2];
  }

  // Четные игровые линии.
  return wbColors[(x) % 2];
}