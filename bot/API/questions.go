package API

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/machinebox/graphql"
	"log"
)

type Question struct {
	Id      uint64 `json:"id"`
	Title   string `json:"title"`
	Status  int16  `json:"status"`
	TopicId uint64 `json:"topicId"`
}
type ResponseQuestionsStruct struct {
	Questions []Question
}

func FindQuestions(token string, topicId string) ResponseQuestionsStruct {
	// make a request
	req := graphql.NewRequest(`
query FindQuestions($topic_id:Int!) {
  questions(topic_id:$topic_id) {
    id,
    title,
    status,
    topicId
  }
}
`)
	req.Var("topic_id", topicId)
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Authorization", "Bearer "+token)
	fmt.Println(token)

	ctx := context.Background()
	var respData ResponseQuestionsStruct
	if err := client.Run(ctx, req, &respData); err != nil {
		log.Fatal(err)
	}
	spew.Dump(respData)
	return respData
}
