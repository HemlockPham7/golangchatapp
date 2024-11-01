# golangchatapp
### set up migrate on Window
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser
Invoke-RestMethod -Uri https://get.scoop.sh | Invoke-Expression 
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

### pull
docker pull postgres:15-alpine

### postgresinit:
docker run --name postgres15 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15-alpine

### createdb:
docker exec -it postgres15 createdb --username=root --owner=root go-chat

### postgres:
docker exec -it postgres15 psql

### dropdb:
docker exec -it postgres15 dropdb go-chat

### migrate create table
migrate create -ext sql -dir db/migrations add_users_table

### migrateup:
migrate -path db/migrations -database "postgresql://root:password@localhost:5433/go-chat?sslmode=disable" -verbose up

### migratedown:
migrate -path db/migrations -database "postgresql://root:password@localhost:5433/go-chat?sslmode=disable" -verbose down