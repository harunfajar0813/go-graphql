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
			"getUsers":   field.GetUsers(db),
			"getEvents":  field.GetEvents(db),
			"getUser":    field.GetUser(db),
			"getEvent":   field.GetEvent(db),
			"getClients": field.GetClients(db),
			"getClient":  field.GetClient(db),
		},
	})
}
