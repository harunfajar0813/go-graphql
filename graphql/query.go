package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"

	"graphi/graphql/field"
)

func userQueries(db *gorm.DB) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"getUsers": field.GetUsers(db),
			"getUser" : field.GetUser(db),
		},
	})
}

func userMutation(db *gorm.DB) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createUser": field.CreateUser(db),
		},
	})
}
