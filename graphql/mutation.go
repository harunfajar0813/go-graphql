package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"

	"graphi/graphql/field"
)

func mutation(db *gorm.DB) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createUser":     field.CreateUser(db),
			"createEvent":    field.CreateEvent(db),
			"createUserRole": field.CreateUserRole(db),
			"createClient":   field.CreateClient(db),
			"topUpBalance":   field.TopUpBalance(db),
			"buyTicket":      field.CreateInvoice(db),
		},
	})
}
