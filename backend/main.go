package main

import (
	"context"
	"fmt"
	"github.com/functionalfoundry/graphqlws"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"net/http"
	"time"
)

func main() {
	defer db.Close()

	go restart(func() {
		var err error
		// main handler

		rootQuery := graphql.ObjectConfig{
			Name: "Query", Fields: queries,
		}
		mutationsQuery := graphql.ObjectConfig{
			Name: "Mutations", Fields: mutations,
		}
		subscribtionsQuery:= graphql.ObjectConfig{
			Name: "Subscriptions", Fields: subscriptions,
		}
		schemaConfig := graphql.SchemaConfig{
			Query:    graphql.NewObject(rootQuery),
			Mutation: graphql.NewObject(mutationsQuery),
			Subscription: graphql.NewObject(subscribtionsQuery),
		}
		schema, err = graphql.NewSchema(schemaConfig)
		if err != nil {
			panic(err)
		}

		h := &AuthByToken{
			handler.New(&handler.Config{
				Schema:   &schema,
				Pretty:   true,
				GraphiQL: true,
			}),
		}

		http.Handle("/api", h)

		// Create a subscription manager
		subscriptionManager = graphqlws.NewSubscriptionManager(&schema)

		// Create a WebSocket/HTTP handler
		graphqlwsHandler := graphqlws.NewHandler(graphqlws.HandlerConfig{
			// Wire up the GraphqL WebSocket handler with the subscription manager
			SubscriptionManager: subscriptionManager,

			// Optional: Add a hook to resolve auth tokens into users that are
			// then stored on the GraphQL WS connections
			Authenticate: func(authToken string) (interface{}, error) {
				// This is just a dumb example
				return authToken, nil
			},
		})

		// The handler integrates seamlessly with existing HTTP servers
		http.Handle("/chat", graphqlwsHandler)

		fmt.Println("Starting GraphQL handler")
		if err := http.ListenAndServe(":8181", nil); err != nil {
			panic(err)
		}
	}, "GraphQL handler")

	go restart(func() {
		// This assumes you have access to the above subscription manager
		if subscriptionManager == nil {
			panic("subscriptionManager is not ready")
		}
		subscriptions := subscriptionManager.Subscriptions()

		for conn, _ := range subscriptions {
			// Things you have access to here:
			//conn.ID()   // The connection ID
			//conn.User() // The user returned from the Authenticate function

			for _, subscription := range subscriptions[conn] {
				// Things you have access to here:
				//subscription.ID            // The subscription ID (unique per conn)
				//subscription.OperationName // The name of the operation
				//subscription.Query         // The subscription query/queries string
				//subscription.Variables     // The subscription variables
				//subscription.Document      // The GraphQL AST for the subscription
				//subscription.Fields        // The names of top-level queries
				//subscription.Connection    // The GraphQL WS connection

				// Prepare an execution context for running the query
				ctx := context.WithValue(context.Background(), "token", conn.User().(string))

				// Re-execute the subscription query
				params := graphql.Params{
					Schema:         schema, // The GraphQL schema
					RequestString:  subscription.Query,
					VariableValues: subscription.Variables,
					OperationName:  subscription.OperationName,
					Context:        ctx,
				}
				result := graphql.Do(params)

				// Send query results back to the subscriber at any point
				data := graphqlws.DataMessagePayload{
					// Data can be anything (interface{})
					Data:   result.Data,
					// Errors is optional ([]error)
					Errors: graphqlws.ErrorsFromGraphQLErrors(result.Errors),
				}
				subscription.SendData(&data)
			}
		}
	}, "GraphQL subscriptions chat")


	go restart(func() {
		cronResolveTopics()
	}, "Cron Resolver")

	go restart(func() {
		cronNotify()
	}, "Cron Notify")

	select {}
}

func restart(f func(), msg string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "recover", r)

			switch t := r.(type) {
			case error:
				fmt.Println("Recover panic!", t.Error())
			case string:
				fmt.Println("Recover panic!", t)
			}
		}
		time.Sleep(time.Second * 5)
		go restart(f, msg)
	}()

	f()
}
