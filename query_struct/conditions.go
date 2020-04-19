package query_struct

/*
 * Imposes conditions on the execution of a SQL statement.
 * The SQL statement that follows an IF keyword and its condition is executed if the condition is satisfied:
 * the Boolean expression returns TRUE.
 * The optional ELSE keyword introduces another SQL statement that is executed when the IF condition is not satisfied:
 * the Boolean expression returns FALSE.
 */
const IF = "IF"
const ELSE = "ELSE"

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

