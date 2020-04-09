package query_struct

/*
 * Оператор EXISTS используется для проверки существования любой записи в подзапросе.
 * Example
 * SELECT column_name(s)
 * FROM table_name
 * WHERE EXISTS
 * (SELECT column_name FROM table_name WHERE condition);
 */
const EXISTS = "EXISTS"

const WITH = "WITH"

