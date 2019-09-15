package API

import (
	"github.com/machinebox/graphql"
)

var (
	client *graphql.Client
)

func init() {
	client = graphql.NewClient("https://api.vtbhack.farwydi.ru/api")
}
