psql:
	docker exec -it postgres psql

create_db:
	docker exec -it postgres createdb --username=root --owner=root chat

drop_db:
	docker exec -it postgres dropdb chat

migrate_up:
	migrate -path ./migrations -database 'postgresql://root:postgres@localhost:5432/chat?sslmode=disable' -verbose up

migrate_down:
	migrate -path ./migrations -database 'postgresql://root:postgres@localhost:5432/chat?sslmode=disable' -verbose down

.PHONY: postgres create_db drop_db migrate_up migrate_down