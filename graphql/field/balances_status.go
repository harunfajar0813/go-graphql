package field

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"

	"graphi/domain/model"
)

var balanceStatus = graphql.NewObject(graphql.ObjectConfig{
	Name: "Balance Status",
	Fields: graphql.Fields{
		"id":   &graphql.Field{Type: graphql.ID},
		"name": &graphql.Field{Type: graphql.String},
	},
	Description: "Balance status data",
})

func GetBalancesStatus(db *gorm.DB) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(balanceStatus),
		Resolve: func(p graphql.ResolveParams) (i interface{}, err error) {
			var bs []*model.BalanceStatus
			if err := db.Find(&bs).Error; err != nil {
				log.Fatal(err)
			}
			return bs, nil
		},
		Description: "get balance status",
	}
}
