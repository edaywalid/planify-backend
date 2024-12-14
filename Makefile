start-db:
	docker-compose up -d

stop-db:
	docker-compose down

create-db:
	docker-compose exec psql_database createdb --username=root --owner=root planify

drop-db:
	docker-compose exec psql_database dropdb planify

db-shell:
	docker-compose  exec -it psql_database psql --username=root -d planify


db-script:
	@echo "Enter the psql script you want to run:"
	@read script; \
	echo "Script: $$script"; \
	docker-compose exec -T psql_database psql --username=root -d planify -c "$$script"

mongosh:
	docker exec -it mongo_database bash -c "mongosh --username root --password password  --authenticationDatabase admin"

tidy:
	go mod tidy

build: tidy
	go build -o bin/main cmd/server/main.go

run: build
	./bin/main

clean:
	rm -rf bin/*

.PHONY: start-db stop-db create-db drop-db db-shell db-script tidy build run clean