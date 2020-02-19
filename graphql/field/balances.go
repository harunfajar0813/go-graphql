package field

import (
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
	"graphi/domain/model"
	"log"
)

var balances = graphql.NewObject(graphql.ObjectConfig{
	Name: "Balances",
	Fields: graphql.Fields{
		"id":                &graphql.Field{Type: graphql.ID},
		"name":              &graphql.Field{Type: graphql.String},
		"user_id":           &graphql.Field{Type: graphql.Int},
		"balance_status_id": &graphql.Field{Type: graphql.Int},
	},
	Description: "Balances status data",
})

func GetBalances(db *gorm.DB) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(balances),
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			var b []*model.Balance
			if err := db.Find(&b).Error; err != nil {
				log.Fatal(err)
			}
			return b, nil
		},
		Description: "get users",
	}
}
