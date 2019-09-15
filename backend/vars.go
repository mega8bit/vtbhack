package main

import (
	"database/sql"
	"github.com/functionalfoundry/graphqlws"
	"github.com/graphql-go/graphql"
	_ "github.com/lib/pq"
)

const DbConnectString = "user=vtbhack password=Koo6ahghok3cahGu9cae dbname=vtbhack host=195.181.245.183 port=5432 sslmode=disable"

var (
	db                  *sql.DB
	subscriptionManager graphqlws.SubscriptionManager
	schema              graphql.Schema
)
