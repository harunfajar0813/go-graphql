package field

import (
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"

	"graphi/domain/model"
)

var client = graphql.NewObject(graphql.ObjectConfig{
	Name: "Client",
	Fields: graphql.Fields{
		"id":           &graphql.Field{Type: graphql.ID},
		"name":         &graphql.Field{Type: graphql.String},
		"email":        &graphql.Field{Type: graphql.String},
		"phone":        &graphql.Field{Type: graphql.String},
		"password":     &graphql.Field{Type: graphql.String},
		"topUpHistory": &graphql.Field{Type: graphql.NewList(balance)},
		"balanceNow":   &graphql.Field{Type: graphql.String},
	},
	Description: "Users data",
})

// query
func GetClients(db *gorm.DB) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(client),
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			var c []*model.User
			if err := db.Where("user_role_id = ?", 2).
				Preload("TopUpHistory").
				Find(&c).Error; err != nil {
				return err, err
			}
			for _, data := range c {
				var balanceNow int
				db.Table("balances").
					Where("balances.user_id = ?", data.ID).
					Select("sum(amount) as n").
					Row().
					Scan(&balanceNow)

				var alreadyBuy int
				db.Table("invoices").Joins("join events on invoices.event_id=events.id").
					Where("invoices.user_id = ?", data.ID).
					Select("sum(price) as p").
					Row().
					Scan(&alreadyBuy)

				data.BalanceNow = balanceNow - alreadyBuy
			}
			return c, nil
		},
		Description: "get clients",
	}
}

func GetClient(db *gorm.DB) *graphql.Field {
	return &graphql.Field{
		Type: client,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			var c model.User

			id, ok := p.Args["id"].(int)
			if ok {
				if err := db.Where("user_role_id = ?", 2).
					Preload("TopUpHistory").
					First(&c, id).Error; err != nil {
					return err, err
				}

				var balanceNow int
				db.Table("balances").
					Where("balances.user_id = ?", id).
					Select("sum(amount) as n").
					Row().
					Scan(&balanceNow)

				var alreadyBuy int
				db.Table("invoices").Joins("join events on invoices.event_id=events.id").
					Where("invoices.user_id = ?", id).
					Select("sum(price) as p").
					Row().
					Scan(&alreadyBuy)

				c.BalanceNow = balanceNow - alreadyBuy
			}
			return c, nil
		},
		Description: "get client by id",
	}
}

// mutation
func CreateClient(db *gorm.DB) *graphql.Field {
	return &graphql.Field{
		Type: client,
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"email": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"phone": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (i interface{}, e error) {
			name, _ := params.Args["name"].(string)
			email, _ := params.Args["email"].(string)
			phone, _ := params.Args["phone"].(string)
			hashedPass, _ := bcrypt.GenerateFromPassword([]byte(params.Args["password"].(string)), bcrypt.DefaultCost)

			newUser := &model.User{
				Name:       name,
				Email:      email,
				Phone:      phone,
				Password:   string(hashedPass),
				UserRoleID: 2,
			}

			err := db.Debug().Model(&model.User{}).Create(&newUser).Error
			if err != nil {
				log.Fatal(err)
			}

			return newUser, nil
		},
	}
}
