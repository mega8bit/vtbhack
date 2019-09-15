package main

import (
	"database/sql"
	"github.com/graphql-go/graphql"
	"time"
)

//id                int8
//title             varchar
//type_id           int2
//start_datetime    timestamp
//end_datetime      timestamp
//status            int2

type Topic struct {
	Id            uint64    `json:"id"`
	Title         string    `json:"title"`
	TypeId        int16     `json:"typeId"`
	StartDateTime time.Time `json:"string"`
	EndDateTime   time.Time `json:"string"`
	Status        int16     `json:"status"`
	ChairmanId    int16     `json:"chairman_id"`
}

var topicType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Topic",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"typeId": &graphql.Field{
				Type: graphql.Int,
			},
			"startDateTime": &graphql.Field{
				Type: graphql.DateTime,
			},
			"endDateTime": &graphql.Field{
				Type: graphql.DateTime,
			},
			"status": &graphql.Field{
				Type: graphql.Int,
			},
			"chairman_id": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var fieldTopic = &graphql.Field{
	Type: graphql.NewList(topicType),
	Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
		token, _ := p.Context.Value("token").(string)
		var topics []*Topic
		rows, err := db.Query(`
          SELECT 
            t.id,
            t.title,
            t.type_id,
            t.start_datetime,
            t.end_datetime,
            t.status,
            t.chairman_id
          FROM topic t
          JOIN "user" u ON u.token = $1
          JOIN topic_user tu ON tu.user_id = u.id
          WHERE t.id = tu.topic_id
          ORDER BY end_datetime, status
        `, token)

		if err == sql.ErrNoRows {
			return topics, nil
		}

		if err != nil {
			return nil, err
		}
		for rows.Next() {
			model := &Topic{}
			err = rows.Scan(
				&model.Id,
				&model.Title,
				&model.TypeId,
				&model.StartDateTime,
				&model.EndDateTime,
				&model.Status,
				&model.ChairmanId,
			)
			if err != nil {
				continue
			}
			topics = append(topics, model)
		}

		return topics, nil
	},
}

var fieldTopicCreate = &graphql.Field{
    Type: graphql.Int,
    Args: graphql.FieldConfigArgument{
        "title": &graphql.ArgumentConfig{
            Type: graphql.NewNonNull(graphql.String),
        },
        "typeId": &graphql.ArgumentConfig{
            Type: graphql.NewNonNull(graphql.Int),
        },
        "startDateTime": &graphql.ArgumentConfig{
            Type: graphql.NewNonNull(graphql.String),
        },
        "endDateTime": &graphql.ArgumentConfig{
            Type: graphql.NewNonNull(graphql.String),
        },
        "chairmanId": &graphql.ArgumentConfig{
            Type: graphql.NewNonNull(graphql.Int),
        },
    },
    Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
        token, _ := p.Context.Value("token").(string)
        title, _ := p.Args["title"].(string)
        typeId, _ := p.Args["typeId"].(int)
        startDateTime, _ := p.Args["startDateTime"].(string)
        endDateTime, _ := p.Args["endDateTime"].(string)
        chairmanId, _ := p.Args["chairmanId"].(int)
        var newTopicId int
        err := db.QueryRow(`
          INSERT INTO topic(
            title,
            type_id,
            start_datetime,
            end_datetime,
            status,
            chairman_id
            )
          VALUES ($1, $2, $3, $4, $5, $6)
          RETURNING id
        `, title, typeId, startDateTime, endDateTime, 0, chairmanId).Scan(&newTopicId)

        if err != nil {
            return nil, err
        }

        _, err = db.Exec(`
          INSERT INTO topic_user (topic_id, user_id)
          VALUES ($1, (SELECT id FROM "user" WHERE token = $2 LIMIT 1))
        `, newTopicId, token)

        if err != nil {
            return nil, err
        }

        return newTopicId, nil
    },
}