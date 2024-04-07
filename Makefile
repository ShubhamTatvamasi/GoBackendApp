postgres: 
		docker run -d -p 5432:5432 \
  		--name postgres \
  		-e POSTGRES_PASSWORD=postgres \
  		-e POSTGRES_DB=simple_bank \
  		postgres:16-alpine

createdb:
		docker exec -it postgres createdb --username=postgres --owner=postgres simple_bank

dropdb:
		docker exec -it postgres dropdb --username=postgres simple_bank

removepostgres:
		docker rm -f postgres

.PHONY: postgres createdb dropdb removepostgres
