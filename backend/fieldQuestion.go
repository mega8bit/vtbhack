package main

import (
	"database/sql"
	"github.com/graphql-go/graphql"
)

//id        int8
//title	    varchar
//status    int2
//topic_id  int8

type Question struct {
	Id      uint64 `json:"id"`
	Title   string `json:"title"`
	Status  int16  `json:"status"`
	TopicId uint64 `json:"topicId"`
}

var questionType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Question",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"status": &graphql.Field{
				Type: graphql.Int,
			},
			"topicId": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var fieldQuestion = &graphql.Field{
	Type: graphql.NewList(questionType),
	Args: graphql.FieldConfigArgument{
		"topic_id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
		token, _ := p.Context.Value("token").(string)
		topicId, _ := p.Args["topic_id"].(int)
		var questions []*Question
		rows, err := db.Query(`
          SELECT 
			q.id,
			q.title,
			q.status,
			q.topic_id
          FROM question q
          JOIN "user" u ON u.token = $1
          JOIN topic_user tu ON tu.user_id = u.id
          JOIN topic t on q.topic_id = t.id
          WHERE t.id = tu.topic_id AND t.id=$2
        `, token, topicId)

		if err == sql.ErrNoRows {
			return questions, nil
		}

		if err != nil {
			return nil, err
		}
		for rows.Next() {
			model := &Question{}
			err = rows.Scan(
				&model.Id,
				&model.Title,
				&model.Status,
				&model.TopicId,
			)
			if err != nil {
				break
			}
			questions = append(questions, model)
		}
		if err := rows.Close(); err != nil {
			return nil, err
		}
		if err != nil {
			return nil, err
		}
		if err := rows.Err(); err != nil {
			return nil, err
		}

		return questions, nil
	},
}

var fieldQuestionCreate = &graphql.Field{
    Type: graphql.ID,
    Args: graphql.FieldConfigArgument{
        "topicId": &graphql.ArgumentConfig{
            Type: graphql.NewNonNull(graphql.Int),
        },
        "title": &graphql.ArgumentConfig{
            Type: graphql.NewNonNull(graphql.String),
        },
    },
    Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
        //token, _ := p.Context.Value("token").(string)
        topicId, _ := p.Args["topicId"].(int)
        title, _ := p.Args["title"].(string)
        var newQuestionId int
        err := db.QueryRow(`
          INSERT INTO question (title, topic_id)
          VALUES ($1, $2)
          RETURNING id
        `, title, topicId).Scan(&newQuestionId)

        if err != nil {
            return nil, err
        }

        return newQuestionId, nil
    },
}
