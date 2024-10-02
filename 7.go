// -- Write your PostgreSQL query statement below
Select Product.product_name,Sales.year,Sales.price from Product
JOIN Sales on Product.product_id = Sales.product_id;