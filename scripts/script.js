async function newAuthor({args, graphql}) {
    // lets give every new author a reputation of 3 by default
    const results = await graphql(`mutation ($name: String!) {
        addAuthor(input: [{name: $name, reputation: 3.0 }]) {
            author {
                id
                reputation
            }
        }
    }`, {"name": args.name})
    console.log('results--------------\n', results)
    return results.data.addAuthor.author[0].id
}

async function authorsByName({args, dql}) {
    const results = await dql.query(`query queryAuthor($name: string) {
        queryAuthor(func: type(Author)) @filter(eq(Author.name, $name)) {
            name: Author.name
            reputation: Author.reputation
        }
    }`, {"$name": args.name})
    console.log('results--------------\n', results)
    return results.data.queryAuthor
}

const authorBio = ({parent: {name, reputation}}) => `My name is ${name} and my reputation is ${reputation}.`

self.addGraphQLResolvers({
    "Author.bio": authorBio,
    "Mutation.newAuthor": newAuthor,
    "Query.authorsByName": authorsByName,
})
