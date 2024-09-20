include .env

.PHONY: migrateup migratedown createdb removedb genquery

migrateup:
	echo "Running migrations up..." 
	cd sql/models && goose postgres $(DB_STRING) up

migratedown:
	echo "Running migrations down..."
	cd sql/models && goose postgres $(DB_STRING) down

createdb:
	echo "Creating database..."
	docker compose up -d

removedb:
	echo "Removing database..."
	docker compose down

genquery:
	sqlc generate

