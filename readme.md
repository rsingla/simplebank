

Golang migrate library:
go get -tags 'postgres' -u github.com/golang-migrate/migrate/v4/cmd/migrate

mkdir migrations

migrate create -ext sql -dir migrations <migration-name>


UP means apply the script changes  
migrate -path migrations -database "postgres://root:simplebankpass@localhost:5432/postgres?sslmode=disable" up

Down means rollback

migrate -path migrations -database "postgres://user:password@localhost:5432/database_name?sslmode=disable" down


Get driver for Postgres 

go get github.com/lib/pq
