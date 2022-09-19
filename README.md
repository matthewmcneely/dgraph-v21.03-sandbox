# Dgraph Sandbox for version 23.01

This repo serves as a quick start for spinning up a dgraph cluster, updating a schema and loading data. Everything's 
done with `make`, the only other requirement is Docker and optionally `jq`.

#### Requirements
- Docker
- make
- curl (optional, for queries from the command line)
- gql (optional, for graphql queries, download from [here](https://github.com/matthewmcneely/gql/tree/feature/add-query-and-variables-from-file/builds))
- jq (optional, for queries from the command line)

## Branch-Specific Steps

```
make schema-dql
make load-rdf-file
COLLECTION=BAR make upsert-query
make query-dql | jq                           #everything's good here
COLLECTION=BAR make upsert-query
make query-dql | jq                           #counts are off
```


## Steps

1. Clone this repo. It's possible I've created a branch for some issue we're collaborating on. If so, check out the branch for the issue.

2. Spin up the cluster
```
make up
```

To spin one up with lambda support: `make up-with-lambda`.

3. Then in another terminal, load the schema
```
make schema-dql
```

or `make schema-gql` for an SDL-based one.

## Make targets

### `make help`
Lists all available make targets and a short description.

### `make up`
Brings up a simple alpha and zero node using docker compose in your local Docker environment.

### `make up-with-lambda`
Brings up the alpha and zero containers along with the dgraph lambda container. Note this lambda container is based on `dgraph/dgraph-lambda:1.4.0`. 

### `make down` and `make down-with-lambda`
Stops the containers.

### `make schema-dql`
Updates dgraph with the schema defined in `schema.dql`.

Example schema.dql:
```
type Person {
    name
    boss_of
    works_for
}

type Company {
    name
    industry
    work_here
}

industry: string @index(term) .
boss_of: [uid] .
name: string @index(exact, term) .
work_here: [uid] .
boss_of: [uid] @reverse .
works_for: [uid] @reverse .
```

### `make schema-gql`
Updates dgraph with the schema defined in `schema.graphql`

Example schema.gql:
```graphql
type Post {
    id: ID!
    title: String!
    text: String
    datePublished: DateTime
    author: Author!
}

type Author {
    id: ID!
    name: String!
    posts: [Post!] @hasInverse(field: author)
}
```

### `make drop-data`
Drops all data from the cluster, but not the schema.

### `make drop-all`
Drops all data and the schema from the cluster.

### `make load-data-gql`
Loads JSON data defined in `gql-data.json`. This target is useful for loading data into schemas defined with GraphQL SDL.

Example gql-data.json:
```json
[
    {
        "uid": "_:katie_howgate",
        "dgraph.type": "Author",
        "Author.name": "Katie Howgate",
        "Author.posts": [
            {
                "uid": "_:katie_howgate_1"
            },
            {
                "uid": "_:katie_howgate_2"
            }
        ]
    },
    {
        "uid": "_:timo_denk",
        "dgraph.type": "Author",
        "Author.name": "Timo Denk",
        "Author.posts": [
            {
                "uid": "_:timo_denk_1"
            },
            {
                "uid": "_:timo_denk_2"
            }
        ]
    },
    {
        "uid": "_:katie_howgate_1",
        "dgraph.type": "Post",
        "Post.title": "Graph Theory 101",
        "Post.text": "https://www.lancaster.ac.uk/stor-i-student-sites/katie-howgate/2021/04/27/graph-theory-101/",
        "Post.datePublished": "2021-04-27",
        "Post.author": {
            "uid": "_:katie_howgate"
        }
    },
    {
        "uid": "_:katie_howgate_2",
        "dgraph.type": "Post",
        "Post.title": "Hypergraphs – not just a cool name!",
        "Post.text": "https://www.lancaster.ac.uk/stor-i-student-sites/katie-howgate/2021/04/29/hypergraphs-not-just-a-cool-name/",
        "Post.datePublished": "2021-04-29",
        "Post.author": {
            "uid": "_:katie_howgate"
        }
    },
    {
        "uid": "_:timo_denk_1",
        "dgraph.type": "Post",
        "Post.title": "Polynomial-time Approximation Schemes",
        "Post.text": "https://timodenk.com/blog/ptas/",
        "Post.datePublished": "2019-04-12",
        "Post.author": {
            "uid": "_:timo_denk"
        }
    },
    {
        "uid": "_:timo_denk_2",
        "dgraph.type": "Post",
        "Post.title": "Graph Theory Overview",
        "Post.text": "https://timodenk.com/blog/graph-theory-overview/",
        "Post.datePublished": "2017-08-03",
        "Post.author": {
            "uid": "_:timo_denk"
        }
    }
]
```

### `make load-data-dql-json`
Loads JSON data defined in `dql-data.json`. This target is useful for loading data into schemas defined with base dgraph types.

Example dql-data.json:
```json
{
    "set": [
        {
            "uid": "_:company1",
            "industry": "Machinery",
            "dgraph.type": "Company",
            "name": "CompanyABC"
        },
        {
            "uid": "_:company2",
            "industry": "High Tech",
            "dgraph.type": "Company",
            "name": "The other company"
        },
        {
            "uid": "_:jack",
            "works_for": { "uid": "_:company1"},
            "dgraph.type": "Person",
            "name": "Jack"
        },
        {
            "uid": "_:ivy",
            "works_for": { "uid": "_:company1"},
            "boss_of": { "uid": "_:jack"},
            "dgraph.type": "Person",
            "name": "Ivy"
        },
        {
            "uid": "_:zoe",
            "works_for": { "uid": "_:company1"},
            "dgraph.type": "Person",
            "name": "Zoe"
        },
        {
            "uid": "_:jose",
            "works_for": { "uid": "_:company2"},
            "dgraph.type": "Person",
            "name": "Jose"
        },
        {
            "uid": "_:alexei",
            "works_for": { "uid": "_:company2"},
            "boss_of": { "uid": "_:jose"},
            "dgraph.type": "Person",
            "name": "Alexei"
        }
    ]
}
```

### `make load-data-dql-rdf`
Loads RDF data defined in `dql-data.rdf`. This target is useful for loading data into schemas defined with base dgraph types.

Example dql-data.rdf:
```rdf
{
  set {
    _:company1 <name> "CompanyABC" .
    _:company1 <dgraph.type> "Company" .
    _:company2 <name> "The other company" .
    _:company2 <dgraph.type> "Company" .

    _:company1 <industry> "Machinery" .

    _:company2 <industry> "High Tech" .

    _:jack <works_for> _:company1 .
    _:jack <dgraph.type> "Person" .

    _:ivy <works_for> _:company1 .
    _:ivy <dgraph.type> "Person" .

    _:zoe <works_for> _:company1 .
    _:zoe <dgraph.type> "Person" .

    _:jack <name> "Jack" .
    _:ivy <name> "Ivy" .
    _:zoe <name> "Zoe" .
    _:jose <name> "Jose" .
    _:alexei <name> "Alexei" .

    _:jose <works_for> _:company2 .
    _:jose <dgraph.type> "Person" .
    _:alexei <works_for> _:company2 .
    _:alexei <dgraph.type> "Person" .

    _:ivy <boss_of> _:jack .

    _:alexei <boss_of> _:jose .
  }
}
```

### `make query-dql`
Runs the query defined in query.dql.

Example query.dql:
```
{
  q(func: eq(name, "CompanyABC")) {
    name
    works_here : ~works_for {
        uid
        name
    }
  }
}
```

### `make query-gql`
Runs the query defined in query.gql and optional variables defined in variables.json.

Example query-gql:
```graphql
query QueryAuthor($order: PostOrder) {
  queryAuthor {
    id
    name
    posts(order: $order) {
      id
      datePublished
      title
      text
    }
  }
}
```

Example variables.json:
```json
{
    "order": {
      "desc": "datePublished"
    }
}
```

