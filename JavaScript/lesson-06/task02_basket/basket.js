"use strict";

let basket = {
  settings: {
    selector: '.myBasket',
  },

  productCount: 2,
  totalAmount: 100,

  /**
   * Инициализирует каталог.
   * @param {Object} userSettings Объект настроек для корзины.
   */
  init(userSettings = {}) {
    Object.assign(this.settings, userSettings);

    const basket = document.querySelector(this.settings.selector);
    //const prd = basket.querySelector(".productCount .value");
    //const amt = basket.querySelector(".totalAmount .value");
  },



  addProduct(product) {

  }

};