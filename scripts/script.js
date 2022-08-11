async function newUserWithScores({args, graphql}) {
    console.log("ARGS", args)
    let highScore = {
        record_date: new Date(),
        score: Math.max(...args.scores)
    }
    let scores = new Array()
    args.scores.forEach(score => {
        scores.push({
            score: score,
            record_date: new Date()
        })
    })    
    const results = await graphql(`mutation ($username: String!, $scores: [GameScoreRef]!, $highScore: GameScoreRef!) {
        addUser(input: [{username: $username, scores: $scores}]) {
            user {
                id
            }
        }
    }`, {
        "username": args.username, 
        "scores": scores,
    })
    return results.data.addUser.user[0].id
}


async function calculateHighScore({parent, graphql}) {
    const results = await graphql(`
    query {
        getUser(id: "${parent.id}") {
          id
          scores(order: {desc: score}, first: 1) {
            score
            record_date
          }
        }
    }`)
    return {
        user: {id: parent.id},
        score: results.data.getUser.scores[0].score,
        record_date: results.data.getUser.scores[0].record_date
    }
}

async function visitCount( { parent } ) {
    console.log(parent);
    return parent.guest_visit_dates.length;
}

self.addGraphQLResolvers({
    "User.high_score": calculateHighScore,
    "Guest.guest_visit_count": visitCount
})


