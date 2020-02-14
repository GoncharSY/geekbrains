"use strict";

let catalog = {
  settings: {
    selector: '.myCatalog',
  },

  products: [{
    name: "Белая кружка",
    price: 100,
    photo: "images/white.png",
  }],

  /**
   * Инициализирует каталог.
   * @param {Object} userSettings Объект настроек для каталога.
   */
  init(userSettings = {}) {
    Object.assign(this.settings, userSettings);
  },



  /**
   * Создать карточку товара.
   * @param {object} product 
   * @param {string} product.name 
   * @param {int} product.price
   * @param {string} product.photo
   */
  createProductCard(product) {
    const photo = new Image();

    photo.src = product.photo;
    photo.classList.add("photo");

    const name = document.createElement("div");
    const nameLabel = document.createElement("span");
    const nameValue = document.createElement("span");

    name.classList.add("dataItem", "name")
    nameLabel.classList.add("label");
    nameLabel.classList.add("value");

    const price = document.createElement("div");
    const priceLabel = document.createElement("span");
    const priceValue = document.createElement("span");

    price.classList.add("dataItem", "price")
    priceLabel.classList.add("label");
    priceLabel.classList.add("value");

    const button = document.createElement("button");

    button.classList.add("buy");
    button.innerHTML = "Купить";

    const details = document.createElement("div");

    details.classList.add("details");
    details.append(name);
    details.append(price);

    const productCard = document.createElement("div");

    productCard.classList.add("productCart");
    productCard.append(photo);
    productCard.append(details);
  }
}

/*
<div class="productCard">
  <img class="photo" src="images/green.png"/>
  <div class="details">
    <div class="dataItem name">
      <span class="label">Название:</span>
      <span class="value">Какой-то товар</span>
    </div>
    <div class="dataItem price">
      <span class="label">Цена:</span>
      <span class="value">0</span>
      <span class="currency">руб.</span>
    </div>
    <button>Купить</button>
  </div>
</div>

*/