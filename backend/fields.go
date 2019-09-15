package main

import "github.com/graphql-go/graphql"

var queries = graphql.Fields{
	"user":   fieldUser,
	"userAll":   fieldUserAll,
	"topics": fieldTopic,
	"questions": fieldQuestion,
	"messages": fieldMessage,
}

var mutations = graphql.Fields{
	"createUser": fieldUserCreate,
	"editUser":   fieldUserEdit,

	"createTopic": fieldTopicCreate,
	"createQuestion": fieldQuestionCreate,

	"vote": fieldVote,
}

var subscriptions = graphql.Fields{
	"addMessage": fieldMessageAdd,
	//"listMessage": fieldMessageList,
}
