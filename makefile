migration:
	migrate create -ext sql -seq -dir ./db/migrations $(name)

migrate:
	migrate -database sqlite3://db/database.db -path ./db/migrations $(type)

run:
	air
