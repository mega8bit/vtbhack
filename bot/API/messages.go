package API

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/machinebox/graphql"
	"log"
)

type Message struct {
	Id         uint64 `json:"id"`
	Body       string `json:"body"`
	QuestionId uint64 `json:"questionId"`
	QuoteId    sql.NullInt64 `json:"quoteId"`
	UserName   string `json:"userName"`
}
type ResponseMessageStruct struct {
	Messages []Message
}

func GetMessages(token string, topicId string) ResponseMessageStruct {
	// make a request
	req := graphql.NewRequest(`
query getmessages($topicId:Int!){
messages(
    topicId:$topicId,
  ){
    id
    body
    questionId
    quoteId
    userName
  }
}
`)
	req.Var("topicId", topicId)
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Authorization", "Bearer "+token)
	fmt.Println(token)

	ctx := context.Background()
	var respData ResponseMessageStruct
	if err := client.Run(ctx, req, &respData); err != nil {
		log.Fatal(err)
	}
	return respData
}
