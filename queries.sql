use fantasy_products;
-- 1. Montos totales redondeados a 2 decimales por condition del customer
SELECT c.condition, TRUNCATE(SUM(i.total), 2)
FROM customers c 
INNER JOIN invoices i
ON i.customer_id = c.id
GROUP BY  c.condition;

-- 2. Top 5 de los products más vendidos y sus cantidades vendidas
SELECT p.description as Description, SUM(s.quantity) as Total
FROM products p 
INNER JOIN sales s
ON p.id = s.product_id
GROUP BY p.id 
ORDER BY Total DESC
LIMIT 5;

-- 3. El Top 5 de los customers activos que gastaron la mayor cantidad de dinero
SELECT c.last_name, c.first_name, TRUNCATE(SUM(i.total), 3) Amount 
FROM customers c 
INNER JOIN invoices i
ON i.customer_id = c.id
GROUP BY i.customer_id
ORDER BY Amount DESC
LIMIT 5;