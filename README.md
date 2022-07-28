# Dgraph Sandbox for version 23.01

This repo serves as a quick start for spinning up a dgraph cluster, updating a schema and loading data. Everything's 
done with `make`, the only other requirement is Docker and optionally `jq`.

#### Requirements
- Docker
- make
- curl (for queries from the command line)
- jq (for queries from the command line)

## Steps

1. Clone this repo. It's possible I've created a branch for some issue we're collaborating on. If so, check out the branch for the issue.

2. Spin up the cluster
```
make up
```

3. Then in another terminal, load the schema
```
make schema
```

4. Load the sample data
```
make load-data
```

5. If there's some DQL query or mutation that needs to be applied for debugging/testing
```
make query-dql
```

If there's a graphql query to be applied, see the query.graphql file for the query and variables (you'll have to pop those into your favorite GraphQL client, use the http://localhost:8080/graphql endpoint)
