# Dgraph Sandbox for version 23.01

This repo serves as a quick start for spinning up a dgraph cluster, updating a schema and loading data. Everything's 
done with `make`, the only other requirement is Docker and optionally `jq`.

#### Requirements
- Docker
- make
- curl (optional, for queries from the command line)
- gql (optional, for graphql queries, download from [here](https://github.com/matthewmcneely/gql/tree/feature/add-query-and-variables-from-file/builds))
- jq (optional, for queries from the command line)

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

This schema contains @auth directives for the sole Consumer type to prevent queries from non-authenticated users.

4. Load the sample data
```
make load-data
```

5. Check out the `jwt.json` file. These are the claims that will get encoded into a JWT token and send to the query via headers.

6. Try issuing the query (in query.graphql), there should be no results.

```
make query-gql
```

7. Try issuing the query by encoding a JWT Token (requires the jwt cli tool available from https://github.com/mike-engel/jwt-cli)

```
make query-gql-auth
```

The results should only show user "Matthew".
