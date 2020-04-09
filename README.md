# goDB
This app for work to databases

pre-install postgreSQL

#### The example
```
InitDatabase(host string, port int, database string, username string, password string)
```


### Create database
```bash
create database testdbgo successfully
```
##### if database exists
```bash
database testdbgo is exists
exit status 1
```


### Create user
```bash
create user: test
```
##### if user exists
```bash
pq: role "test" already exists
exit status 1
```