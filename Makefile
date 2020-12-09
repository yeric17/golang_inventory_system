postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
createdb:
	docker exec -it postgres12  createdb --username=root --owner=root inventory
dropdb:
	docker exec -it postgres12  dropdb inventory
migrateup:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/inventory?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/inventory?sslmode=disable" -verbose down
buildapp:
	docker build -t inventory-system .
startapp:
	docker run -it -p 9090:9090 inventory-system

.PHONY: postgres createdb dropdb migrateup migratedown