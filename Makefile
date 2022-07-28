current_dir = $(shell pwd)

# Start the zero and alpha containers
up:
	docker-compose up

# Stop the containers
stop:
	docker-compose stop

# Load/update the GraphQL schema
schema:
	curl -v  --data-binary '@./schema.graphql' --header 'content-type: application/octet-stream' http://localhost:8080/admin/schema

# Drops all data (but not the schema)
drop-data:
	curl -X POST localhost:8080/alter -d '{"drop_op": "DATA"}'

# Drops data and schema
drop-all:
	curl -X POST localhost:8080/alter -d '{"drop_all": true}'

load-data:
	docker run -it -v $(current_dir):/export dgraph/dgraph:v21.03.0 dgraph live -a host.docker.internal:9080 -z host.docker.internal:5080 -f /export/data.json
	