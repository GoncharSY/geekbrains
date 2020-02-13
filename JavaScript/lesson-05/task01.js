/*
  1. Создать функцию, генерирующую шахматную доску. При этом можно 
  использовать любые html-теги по своему желанию.

  Доска должна быть разлинована соответствующим образом, т.е. чередовать черные 
  и белые ячейки.
  
  Строки должны нумероваться числами от 1 до 8, столбцы – латинскими буквами 
  A, B, C, D, E, F, G, H.
*/

"use strict";



/**
 * Создать шахматную доску.
 * @returns {HTMLDivElement}
 */
function createChessBoard() {
  const board = document.createElement("div");

  board.style.display = "flex";
  board.style.flexDirection = "column";
  board.style.borderTop = "1px solid #868686";
  board.style.borderLeft = board.style.borderTop;

  for (let y = 9; y >= 0; y--) {
    board.append(createBoardLine(y));
  }

  return board;
}



/**
 * Создать линию на шахматной доске.
 * @param {int} y 
 * @returns {HTMLDivElement}
 */
function createBoardLine(y) {
  const line = document.createElement("div");

  line.style.display = "flex";
  line.style.flexDirection = "row";

  for (let x = 0; x < 10; x++) {
    line.append(createSquare(x, y));
  }

  return line;
}



/**
 * Создать шахматную ячейку.
 * @param {int} x
 * @param {int} y
 * @returns {HTMLDivElement}
 */
function createSquare(x, y) {
  const square = document.createElement("div");

  square.style.width = "40px";
  square.style.height = square.style.width;
  square.style.backgroundColor = getSquareColor(x, y);
  square.style.borderBottom = "1px solid #868686";
  square.style.borderRight = square.style.borderBottom;
  square.style.display = "flex";
  square.style.alignItems = "center";
  square.style.justifyContent = "center";
  square.style.fontWeight = "bold";

  square.innerHTML = getSquareText(x, y);
  square.dataset.position = `x${x}_y${y}`;
  square.classList.add("square");

  return square;
}



/**
 * Получить текст для шахматной ячейки.
 * @param {int} x 
 * @param {int} y
 * @returns {string}
 */
function getSquareText(x, y) {
  const letters = ["a", "b", "c", "d", "e", "f", "g", "h"];

  // Игровые клетки.
  if (x > 0 && x < 9 && y > 0 && y < 9 ) {
    return "";
  }

  // Вертикальные полосы.
  if (x <= 0 || x >= 9) {
    // Углы
    if (y <= 0 || y >= 9) {
      return "";
    }

    return String(y);
  }

  // Горизонтальные полосы.
  if (y <= 0 || y >= 9) {
    return letters[x - 1];
  }
}



/**
 * Получить цвет шахматной ячейки.
 * @param {int} x
 * @param {int} y
 * @returns {string}
 */
function getSquareColor(x, y) {
  const wbColors = ["#fff", "#888"];
  const outerColor = "#ccc";

  // Не игровые клетки (с буквами).
  if (x < 1 || y < 1 || x > 8 || y > 8) {
    return outerColor;
  }

  // Нечетные игровые линии.
  if (y % 2) {
    return wbColors[x % 2];
  }

  // Четные игровые линии.
  return wbColors[(x + 1) % 2];
}
