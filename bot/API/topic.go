package API

import (
	"context"
	"github.com/machinebox/graphql"
	"log"
	"time"
)

type Topic struct {
	Id            uint64    `json:"id"`
	Title         string    `json:"title"`
	TypeId        int16     `json:"typeId"`
	StartDateTime time.Time `json:"string"`
	EndDateTime   time.Time `json:"string"`
	Status        int16     `json:"status"`
	ChairmanId    int16     `json:"chairman_id"`
}

type ResponseTopicsStruct struct {
	Topics []Topic
}

func GetAllTopics(token string) ResponseTopicsStruct {
	// make a request
	req := graphql.NewRequest(`
	{
		topics {
			id
			status
			title
			typeId
		}
	}
`)
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Authorization", "Bearer "+token)
	ctx := context.Background()
	var respData ResponseTopicsStruct
	if err := client.Run(ctx, req, &respData); err != nil {
		log.Fatal(err)
	}
	return respData
}
