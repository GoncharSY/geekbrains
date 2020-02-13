/*
  2*. Заполнить созданную таблицу фигурами, фигуры должны выглядеть как 
  картинки, либо можно использовать символы (существуют символы шахматных 
  фигур) причем все фигуры должны стоять на своих местах и быть соответственно 
  черными и белыми.
*/

"use strict";

const pieces = [{
  caracter: "♔",
  name: "Белый король",
  startPositions: [{x: 5, y: 1}]
},{
  caracter: "♕",
  name: "Белый ферзь",
  startPositions: [{x: 4, y: 1}]
},{
  caracter: "♗",
  name: "Белый офицер",
  startPositions: [
    {x: 3, y: 1},
    {x: 6, y: 1}
  ]
},{
  caracter: "♘",
  name: "Белый конь",
  startPositions: [
    {x: 2, y: 1},
    {x: 7, y: 1}
  ]
},{
  caracter: "♖",
  name: "Белая ладья",
  startPositions: [
    {x: 1, y: 1},
    {x: 8, y: 1}
  ]
},{
  caracter: "♙",
  name: "Белая пешка",
  startPositions: [
    {x: 1, y: 2},
    {x: 2, y: 2},
    {x: 3, y: 2},
    {x: 4, y: 2},
    {x: 5, y: 2},
    {x: 6, y: 2},
    {x: 7, y: 2},
    {x: 8, y: 2}
  ]
},{
  caracter: "♚",
  name: "Черный король",
  startPositions: [{x: 4, y: 8}]
},{
  caracter: "♛",
  name: "Черный ферзь",
  startPositions: [{x: 5, y: 8}]
},{
  caracter: "♝",
  name: "Черный офицер",
  startPositions: [
    {x: 3, y: 8},
    {x: 6, y: 8}
  ]
},{
  caracter: "♞",
  name: "Черный конь",
  startPositions: [
    {x: 2, y: 8},
    {x: 7, y: 8}
  ]
},{
  caracter: "♜",
  name: "Черная ладья",
  startPositions: [
    {x: 1, y: 8},
    {x: 8, y: 8}
  ]
},{
  caracter: "♟",
  name: "Черная пешка",
  startPositions: [
    {x: 1, y: 7},
    {x: 2, y: 7},
    {x: 3, y: 7},
    {x: 4, y: 7},
    {x: 5, y: 7},
    {x: 6, y: 7},
    {x: 7, y: 7},
    {x: 8, y: 7}
  ]
}];



/**
 * Расставить шахматные фигуры в исходных позициях.
 * @param {HTMLDivElement} board 
 */
function setInitialPositions(board) {
  pieces.forEach(piece => {
    piece.startPositions.forEach(startPos => {
      let position = `x${startPos.x}_y${startPos.y}`;
      let square = board.querySelector(`.square[data-position='${position}']`);

      if (!square) {
        throw `Для фигуры '${piece.name}' не найдена позиция '${position}' на доске`;
      }

      setPieceToSquare(piece, square);
    });
  });
}



/**
 * Установить фигуру в квадрат на шахматной доске.
 * @param {object} piece 
 * @param {HTMLDivElement} square 
 */
function setPieceToSquare(piece, square) {
  const pieceTag = document.createElement("span");

  pieceTag.style.fontSize = "40px";
  pieceTag.innerHTML = piece.caracter;
  square.append(pieceTag);
}
