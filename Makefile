postgres:
	docker run --name shaffra -p 5432:5432  -e POSTGRES_USER=postgres  -e POSTGRES_PASSWORD=postgres -d postgres:latest

createdb:
	docker exec -it shaffra createdb --username=postgres --owner=postgres user_data

migratecreate:
	migrate create -ext sql -dir internal/database/migrations/ -seq init_schema

migrateup:
	 migrate -path internal/database/migrations/ -database "postgresql://postgres:postgres@localhost:5432/user_data?sslmode=disable" -verbose up

dropdb:
	docker exec -it schedule dropdb user_data

migratedown:
	migrate -path internal/database/migrations/ -database "postgresql://postgres:postgres@$5432:5432/user_data?sslmode=disable" -verbose down

.PHONY: postgres createdb createtestdb dropdb migrateup migratedown migratecreate