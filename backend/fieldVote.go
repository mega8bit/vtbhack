package main

import (
	"database/sql"
	"github.com/graphql-go/graphql"
)

//id            int8
//question_id   int8
//user_id       int8
//result        int2

type Vote struct {
	Id         uint64 `json:"id"`
	QuestionId uint64 `json:"questionId"`
	UserId     uint64 `json:"userId"`
	Result     int16  `json:"result"`
}

var voteType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Vote",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"questionId": &graphql.Field{
				Type: graphql.Int,
			},
			"userId": &graphql.Field{
				Type: graphql.Int,
			},
			"result": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var fieldVote = &graphql.Field{
	Type: graphql.Boolean,
	Args: graphql.FieldConfigArgument{
		"questionId": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"action": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
		token, _ := p.Context.Value("token").(string)
		questionId, _ := p.Args["questionId"].(int)
		action, _ := p.Args["action"].(int)

		var lastVoteId int
		err := db.QueryRow(`
          SELECT id FROM vote
          WHERE question_id = $1
          AND user_id = (SELECT id FROM "user" WHERE token = $2)
        `, questionId, token).Scan(&lastVoteId)

		if err == sql.ErrNoRows {
			db.Exec(`
				INSERT INTO vote (question_id, user_id, result)
				VALUES ($1, (SELECT id FROM "user" WHERE token = $2), $3)
			`, questionId, token, action)
			return true, nil
		}

		if err != nil {
			return nil, err
		}

			db.Exec(`
				UPDATE vote SET result = $3
				WHERE question_id = $1
				AND user_id = (SELECT id FROM "user" WHERE token = $2)
			`, questionId, token, action)

		return true, nil
	},
}
