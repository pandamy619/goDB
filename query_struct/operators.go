package query_struct


const CREATE = "CREATE"
/*
 * CREATE DATABASE this request create new database in SQL
 * Example
 * CREATE DATABASE my_database;
 */
const CREATE_DB = CREATE + " DATABASE"
/*
 * CREATE USER this request create new user in SQL
 * Example
 * CREATE USER my_database WITH password 'password';
 */
const CREATE_USER = CREATE + " USER"
const PASSWORD = "PASSWORD"

