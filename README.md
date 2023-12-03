# photogo
A photo album in Go.


### Mod

```bash
go mod init github.com/prantoran/photogo
go mod tidy
```

### Live preloading

#### Install
```bash
go install github.com/cosmtrek/air@latest
go get github.com/gorilla/mux
go get github.com/gorilla/schema
```
#### Run directly
```
air --build.cmd "go build ." --build.bin "./photogo"
```
#### Run using config
```
air init
air
```

## Postgres
### Start
```
docker-compose up -d
```
#### Stop
```
docker-compose down
```
#### Commands
```
docker exec -it PostgresCont bash
psql -U postgres
psql
> CREATE DATABASE photogo_db;
```
- Connect to database
```
postgres=# \c photogo_db
```
output:
```
You are now connected to database "photogo_db" as user "postgres".
photogo_db=#
```
- Crate tables
```
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT,
    email TEXT NOT NULL 
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    amount INT,
    description TEXT
);
```
- Selects
```
Select * from users;
```
- Drop tables
```
DROP TABLE users;
DROP TABLE orders;
```


#### Debugging
> `postgres` role does not exists
1. Kill and remove the docker containers
2. Configure `docker-compose.yml`
3. Restart docker compose


### GORM
- Install
```
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
```