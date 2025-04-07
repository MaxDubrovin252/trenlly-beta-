# Makefile

.PHONY: run docker migrate

docker:
	sudo docker run --name=trsesdn -e POSTGRES_PASSWORD='qwerty' -p 5439:5432 -d --rm postgres
# Запускает приложение
migrate:
	
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5439/postgres?sslmode=disable' up

run:
	go run cmd/main.go

# Запускает Docker-контейнер




# Выполняет миграции



# Запускает приложение, Docker и миграции последовательно
all:  docker migrate run