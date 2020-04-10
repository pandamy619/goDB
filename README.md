# goDB
This app for work to databases

pre-install postgreSQL

#### init struct
```bash
databse.Database{Host string, Port string, User string, Password string, conn *sql.DB}
```

#### open SQL connect

```go
func (db *Database) InitDatabase(host string, port int, username string, password string)
```

#### Create database
```go
func (db *Database) CreateDB(dbname string)
```

#### Create user
```go
func (db *Database) CreateUser(username string, password string)
```
