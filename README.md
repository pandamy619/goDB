# goDB
This app for work to databases

pre-install postgreSQL

#### init struct
```go
database.Database{Host string, Port string, User string, Password string, conn *sql.DB}
```
------------------

#### open SQL connect

```go
func (db *Database) InitDatabase()
```
Initialize database connection

------------------
#### Create database
```go
func (db *Database) CreateDatabase(dbname string)
```
Used to create a new SQL database. If database exists return error

------------------
#### Create user
```go
func (db *Database) CreateUser(username string, password string)
```
Used to create a new SQL user. If role exists return error

------------------