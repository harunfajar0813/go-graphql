package field

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"

	"graphi/domain/model"
)

var balance = graphql.NewObject(graphql.ObjectConfig{
	Name: "Balance",
	Fields: graphql.Fields{
		"id":        &graphql.Field{Type: graphql.ID},
		"amount":    &graphql.Field{Type: graphql.String},
		"createdAt": &graphql.Field{Type: graphql.String},
	},
	Description: "Balance data",
})

func TopUpBalance(db *gorm.DB) *graphql.Field {
	return &graphql.Field{
		Type: balance,
		Args: graphql.FieldConfigArgument{
			"amount": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"userId": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(params graphql.ResolveParams) (i interface{}, e error) {
			amount, _ := params.Args["amount"].(int)
			userId, _ := params.Args["userId"].(int)

			var result int64
			db.Table("users").
				Where("users.user_role_id = ?", 2).
				Where("users.id = ?", userId).
				Count(&result)

			if result == 0 {
				return errors.New("not found"), errors.New("not found")
			}

			addBalance := &model.Balance{
				Amount: amount,
				UserID: userId,
			}

			err := db.Debug().Model(&model.Balance{}).Create(&addBalance).Error
			if err != nil {
				return err, err
			}

			return addBalance, nil
		},
	}
}
