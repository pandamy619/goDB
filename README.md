# goDB
This app for work to databases

pre-install postgreSQL

#### The example
```
InitDatabase(host string, port int, database string, username string, password string)
```

```bash
2020/04/09 15:30:27 Connection host: localhost:5432
2020/04/09 15:30:27 create database testdbgo successfully
```
##### if db exists
```bash
2020/04/09 15:31:29 Connection host: localhost:5432
2020/04/09 15:31:29 database testdbgo is exists
exit status 1
```