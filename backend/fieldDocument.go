package main

import (
	"github.com/graphql-go/graphql"
)

//id        int8
//file      text
//name      varchar
//topic_id  int8

type Document struct {
	Id      uint64 `json:"id"`
	File    string `json:"file"`
	Name    string `json:"name"`
	TopicId uint64 `json:"topicId"`
}

var documentType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Message",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"file": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"topicId": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)
