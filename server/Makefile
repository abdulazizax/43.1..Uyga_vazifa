CURRENT_DIR=$(shell pwd)
APP=template
APP_CMD_DIR=./cmd

DB_URL=postgres://ahrorbek:ahrorbek@localhost:5432/bookstore_db

run:
	go run cmd/main.go

	
migrate_up:
	migrate -path migrations -database ${DB_URL} -verbose up

migrate_down:
	migrate -path migrations -database ${DB_URL} -verbose down

migrate_force:
	migrate -path migrations -database ${DB_URL} -verbose force 1

migrate_file:
	migrate create -ext sql -dir migrations -seq ideal_cleaning_db_table

swag-gen:
	~/go/bin/swag init -g ./api/router.go -o api/docs force 1