/*
    Выведите список товаров products и разделов catalogs, 
    который соответствует товару.
*/

SELECT
	prd.name AS 'Product',
	cat.name AS 'Catalog'
FROM products AS prd
LEFT JOIN catalogs AS cat ON cat.id = prd.catalog_id;