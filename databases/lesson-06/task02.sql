/*
Создайте представление, которое выводит название (name) товарной позиции 
из таблицы products и соответствующее название (name) каталога из таблицы 
catalogs.
*/

-- Выберем базу:
USE gsy_shop;

-- Создадим представление:
CREATE OR REPLACE VIEW view_products (`product`, `catalog`) 
    AS SELECT prd.name, cat.name
        FROM products AS prd
        JOIN catalogs AS cat ON cat.id = prd.catalog_id;

-- Результат:
SELECT * FROM view_products;
