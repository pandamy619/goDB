package query_struct

/*
 * Структура sql-запросов
 * Общая структура запроса выглядит следующим образом:
 * SELECT ('столбцы или * для выбора всех столбцов; обязательно')
 * FROM ('таблица; обязательно')
 * WHERE ('условие/фильтрация, например, city = 'Moscow'; необязательно')
 * GROUP BY ('столбец, по которому хотим сгруппировать данные; необязательно')
 * HAVING ('условие/фильтрация на уровне сгруппированных данных; необязательно')
 * ORDER BY ('столбец, по которому хотим отсортировать вывод; необязательно')
 */
/*
 * SELECT, FROM — элементы запроса,
 * которые определяют выбранные столбцы, их порядок и источник данных.
 * Example
 * SELECT * FROM Customers
*/
const SELECT = "SELECT"
const FROM = "FROM"
/*
 * WHERE — элемент запроса,
 * который используется, когда нужно отфильтровать данные по нужному условию
 * Example
 * SELECT * FROM Customers
 * WHERE City = 'London'
 */
const WHERE = "WHERE"
/*
 * GROUP BY — элемент запроса,
 * с помощью которого можно задать агрегацию по нужному столбцу
 * Example
 * SELECT City, count(CustomerID) FROM Customers
 * GROUP BY City
 */
const GROUP_BY = "GROUP BY"
/*
 * HAVING — элемент запроса,
 * который отвечает за фильтрацию на уровне сгруппированных данных
 * Example
 * SELECT City, count(CustomerID) FROM Customers
 * GROUP BY City
 * HAVING count(CustomerID) >= 5
*/
const HAVING = "HAVING"
/*
 * ORDER BY — элемент запроса,
 * который отвечает за сортировку таблицы.
 * Example
 * SELECT * FROM Customers
 * ORDER BY City
 */
const ORDER_BY = "ORDER BY"
