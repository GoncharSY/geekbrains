"use strict";

let app = {
  settings: {
    appSelector: ".myApp",
    catalogClass: "productCatalog",
    basketClass: "productBasket",
  },

  basketItems: [],

  products: [{
    name: "Белая кружка",
    price: 100,
    photo: "images/white.png",
  },{
    name: "Черная кружка",
    price: 110,
    photo: "images/black.png",
  },{
    name: "Синяя кружка",
    price: 120,
    photo: "images/blue.png",
  },{
    name: "Зеленая кружка",
    price: 130,
    photo: "images/green.png",
  },{
    name: "Красная кружка",
    price: 140,
    photo: "images/red.png",
  }],



  /**
   * Инициализирует приложение интерней магазина.
   * Создает панель корзины и каталог товаров.
   * @param {Object} userSettings Объект настроек для каталога.
   */
  init(userSettings = {}) {
    Object.assign(this.settings, userSettings);

    const app = this.getAppContainer();
    const basket = this.createNewBasket();
    const catalog = document.createElement("div");

    if (!app) {
      throw "Контейнер для приложения не найден";
    }

    basket.classList.add(this.settings.basketClass);
    catalog.classList.add(this.settings.catalogClass);
    this.fillCatalogWithProducts(catalog, this.products);
    app.append(basket);
    app.append(catalog);
  },


  
  /**
   * Получить контейнер приложения.
   */
  getAppContainer() {
    const container = document.querySelector(this.settings.appSelector);

    if (!container) {
      throw "Контейнер для приложения не найден";
    }

    return container;
  },



  /**
   * Получить корзину с товарами.
   */
  getBasket() {
    const container = this.getAppContainer();
    const basket = container.querySelector(`.${this.settings.basketClass}`);

    if (!basket) {
      throw "Корзина товаров не найдена";
    }

    return basket;
  },



  /**
   * Создать новый экранный элемент для корзины.
   * Элемент представляет из себя панель с информацией о количестве
   * выбранного товара и его общей стоимости.
   * @returns {HTMLDivElement}
   */
  createNewBasket() {
    // Количество товаров.
    const count = document.createElement("div");
    const countLabel = document.createElement("span");
    const countValue = document.createElement("span");
    const countUnits = document.createElement("span");

    countLabel.classList.add("label");
    countLabel.innerHTML = "Всего товаров:";
    countValue.classList.add("value");
    countValue.innerHTML = "0";
    countUnits.classList.add("units");
    countUnits.innerHTML = "шт.";
    count.classList.add("dataItem", "productCount");
    count.append(countLabel, countValue, countUnits);

    // Cумма покупки.
    const amount = document.createElement("div");
    const amountLabel = document.createElement("span");
    const amountValue = document.createElement("span");
    const amountUnits = document.createElement("span");

    amountLabel.classList.add("label");
    amountLabel.innerHTML = "Общая цена:";
    amountValue.classList.add("value");
    amountValue.innerHTML = "0";
    amountUnits.classList.add("units");
    amountUnits.innerHTML = "руб.";
    amount.classList.add("dataItem", "totalAmount");
    amount.append(amountLabel, amountValue, amountUnits);

    // Кнопка "Очистить".
    const button = document.createElement("button");

    button.classList.add("clear", "dataItem");
    button.innerHTML = "Очистить";
    button.addEventListener("click", () => this.clearBasket());

    // Инф. панель/контейнер корзины.
    const basket = document.createElement("div");

    basket.classList.add(this.settings.basketClass);
    basket.append(count);
    basket.append(amount);
    basket.append(button);

    return basket;
  },


  /**
   * Очистить корзину.
   */
  clearBasket() {
    this.basketItems = [];
    this.updateBasket();
  },



  /**
   * Добавить товар в корзину.
   * @param {object} product 
   */
  addBasketItem(product) {
    this.basketItems.push(product);
  },



  /**
   * Обновить значения на панели корзины.
   */
  updateBasket() {
    const basket = this.getBasket();
    const itmCnt = this.basketItems.length;
    let amount = 0;

    this.basketItems.forEach(itm => amount += itm.price);
    basket.querySelector(".productCount .value").innerHTML = itmCnt;
    basket.querySelector(".totalAmount .value").innerHTML = amount;
  },



  /**
   * Купить товар.
   * @param {object} product 
   */
  buyProduct(product) {
    this.addBasketItem(product);
    this.updateBasket();
  },



  /**
   * Получить каталог товаров.
   */
  getCatalog() {
    const container = this.getAppContainer();
    const catalog = container.querySelector(`.${this.settings.catalogClass}`);

    if (!catalog) {
      throw "Каталог товаров не найден";
    }

    return catalog;
  },



  /**
   * Наполнить каталог товарами.
   * @param {HTMLDivElement} catalog
   * @param {object[]} products
   * @returns {HTMLDivElement}
   */
  fillCatalogWithProducts(catalog, products) {
    products.forEach(prod => {
      catalog.append(this.createProductCard(prod));
    });
  },



  /**
   * Создать карточку товара.
   * @param {object} product 
   * @param {string} product.name 
   * @param {int} product.price
   * @param {string} product.photo
   */
  createProductCard(product) {
    // Фото продукта.
    const photo = new Image();

    photo.src = product.photo;
    photo.classList.add("photo");

    // Название товара.
    const name = document.createElement("div");
    const nameLabel = document.createElement("span");
    const nameValue = document.createElement("span");

    nameLabel.classList.add("label");
    nameLabel.innerHTML = "Название:";
    nameValue.classList.add("value");
    nameValue.innerHTML = product.name;
    name.classList.add("dataItem", "name");
    name.append(nameLabel, nameValue);

    // Цена товара.
    const price = document.createElement("div");
    const priceLabel = document.createElement("span");
    const priceValue = document.createElement("span");

    priceLabel.classList.add("label");
    priceLabel.innerHTML = "Цена:";
    priceValue.classList.add("value");
    priceValue.innerHTML = product.price;
    price.classList.add("dataItem", "price");
    price.append(priceLabel, priceValue);

    // Кнопка "Купить".
    const button = document.createElement("button");

    button.classList.add("buy");
    button.innerHTML = "Купить";
    button.addEventListener("click", () => this.buyProduct(product));

    // Описание товара (название, цена...).
    const details = document.createElement("div");

    details.classList.add("details");
    details.append(name);
    details.append(price);
    details.append(button);

    // Карточка товара.
    const productCard = document.createElement("div");

    productCard.classList.add("productCard");
    productCard.append(photo);
    productCard.append(details);

    return productCard;
  },
};
