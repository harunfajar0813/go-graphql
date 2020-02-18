package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"

	"graphi/graphql/field"
)

func queries(db *gorm.DB) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"getUsers":  field.GetUsers(db),
			"getUser":   field.GetUser(db),
			"getEvents": field.GetEvents(db),
			"getEvent":  field.GetEvent(db),
		},
	})
}

func mutation(db *gorm.DB) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createUser":  field.CreateUser(db),
			"createEvent": field.CreateEvent(db),
		},
	})
}
