package field

import (
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
	"log"

	"graphi/domain/model"
)

var balances = graphql.NewObject(graphql.ObjectConfig{
	Name: "Balances",
	Fields: graphql.Fields{
		"id":      &graphql.Field{Type: graphql.ID},
		"amount":  &graphql.Field{Type: graphql.Int},
		"user_id": &graphql.Field{Type: graphql.Int},
	},
	Description: "Balances status data",
})

// query
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
		Description: "get balances",
	}
}

// Mutation
func TopUpBalance(db *gorm.DB) *graphql.Field {
	return &graphql.Field{
		Type: balances,
		Args: graphql.FieldConfigArgument{
			"amount": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"user_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(params graphql.ResolveParams) (i interface{}, e error) {
			amount, _ := params.Args["amount"].(int)
			userId, _ := params.Args["user_id"].(int)

			newBalance := &model.Balance{
				Amount: amount,
				UserID: userId,
			}

			err := db.Debug().Model(&model.Balance{}).Create(&newBalance).Error
			if err != nil {
				log.Fatal(err)
			}

			return newBalance, nil
		},
	}
}
