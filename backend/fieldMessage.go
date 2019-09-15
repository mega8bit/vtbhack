package main

import (
	"database/sql"
	"errors"
	"github.com/graphql-go/graphql"
)

//id            int8
//body          text
//question_id   int8
//quote_id      int8
//user_id       int8

type Message struct {
	Id         uint64 `json:"id"`
	Body       string `json:"body"`
	QuestionId uint64 `json:"questionId"`
	QuoteId    uint64 `json:"quoteId"`
	UserId     uint64 `json:"userId"`
}

type MessageWithName struct {
	Id         uint64 `json:"id"`
	Body       string `json:"body"`
	QuestionId uint64 `json:"questionId"`
	QuoteId    sql.NullInt64 `json:"quoteId"`
	UserName   string `json:"userName"`
}

var messageType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Message",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"body": &graphql.Field{
				Type: graphql.String,
			},
			"questionId": &graphql.Field{
				Type: graphql.Int,
			},
			"quoteId": &graphql.Field{
				Type: graphql.Int,
			},
			"userId": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var messageTypeWithName = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "MessageTypeWithName",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"body": &graphql.Field{
				Type: graphql.String,
			},
			"questionId": &graphql.Field{
				Type: graphql.Int,
			},
			"quoteId": &graphql.Field{
				Type: graphql.Int,
			},
			"userName": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)


var fieldMessageAdd = &graphql.Field{
	Type: messageType,
	Args: graphql.FieldConfigArgument{
		"questionId": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"body": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"quoteId": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {

		token, _ := p.Context.Value("token").(string)
		var userId uint64
		err := db.QueryRow(`
			SELECT 
				id
			FROM user
			WHERE token = $1
        `, token).
		Scan(&userId)

		if err != nil {
			return false, err
		}

		if userId == 0 {
			return false, errors.New("user not found")
		}

		questionId, _ := p.Args["questionId"].(uint64)
		body, _ := p.Args["body"].(string)
		quoteId, _ := p.Args["quoteId"].(uint64)

		if quoteId == 0 {
			_, err := db.Exec(`
                INSERT INTO "message"(
                    question_id, 
                    body, 
                    user_id,
                ) VALUES ($1, $2, $3)
            `,
				questionId,
				body,
				userId,
			)
			if err != nil {
				return false, err
			}
			return true, nil
		}

		_, err = db.Exec(`
                INSERT INTO "message"(
                    question_id, 
                    body, 
                    quote_id, 
                    user_id,
                ) VALUES ($1, $2, $3, $4)
            `,
			questionId,
			body,
			quoteId,
			userId,
		)
		if err != nil {
			return false, err
		}
		return true, nil

	},
}

var fieldMessage = &graphql.Field{
	Type: graphql.NewList(messageTypeWithName),
	Args: graphql.FieldConfigArgument{
		"topicId": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		//token, _ := p.Context.Value("token").(string)
		topicId, _ := p.Args["topicId"].(int)
		rows, err := db.Query(`
			SELECT 
				m.id,
				m.body,
				m.question_id,
				m.quote_id,
				u.name
			FROM message m
			JOIN "user" u ON u.id = m.user_id
			JOIN question q ON q.id = m.question_id
			JOIN topic t ON t.id = q.topic_id
			WHERE t.id = $1
			ORDER BY m.id
        `, topicId)
		var messages []*MessageWithName
		if err == sql.ErrNoRows {
			return messages, nil
		}
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			model := new(MessageWithName)
			err = rows.Scan(&model.Id, &model.Body, &model.QuestionId, &model.QuoteId, &model.UserName)
			if err != nil {
				continue
			}
			messages = append(messages, model)
		}

		return messages, nil
	},
}

