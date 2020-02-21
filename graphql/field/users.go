package field

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"

	"graphi/domain/model"
)

var user = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id":          &graphql.Field{Type: graphql.ID},
		"name":        &graphql.Field{Type: graphql.String},
		"description": &graphql.Field{Type: graphql.String},
		"email":       &graphql.Field{Type: graphql.String},
		"phone":       &graphql.Field{Type: graphql.String},
		"password":    &graphql.Field{Type: graphql.String},
		"events":      &graphql.Field{Type: graphql.NewList(event)},
	},
	Description: "Users data",
})

// query
func GetUsers(db *gorm.DB) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(user),
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			var u []*model.User
			if err := db.Where("user_role_id = ?", 1).Preload("Events").Find(&u).Error; err != nil {
				log.Fatal(err)
			}
			return u, nil
		},
		Description: "get users",
	}
}

func GetUser(db *gorm.DB) *graphql.Field {
	return &graphql.Field{
		Type: user,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			var u model.User

			id, ok := p.Args["id"].(int)
			if ok {
				if err := db.Where("user_role_id = ?", 1).Preload("Events").First(&u, id).Error; err != nil {
					log.Fatal(err)
					return nil, err
				}
			}
			return u, nil
		},
		Description: "get user by id",
	}
}

// mutation
func CreateUser(db *gorm.DB) *graphql.Field {
	return &graphql.Field{
		Type: user,
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"description": &graphql.ArgumentConfig{
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
			description, _ := params.Args["description"].(string)
			email, _ := params.Args["email"].(string)
			phone, _ := params.Args["phone"].(string)
			hashedPass, _ := bcrypt.GenerateFromPassword([]byte(params.Args["password"].(string)), bcrypt.DefaultCost)

			newUser := &model.User{
				Name:        name,
				Description: description,
				Email:       email,
				Phone:       phone,
				Password:    string(hashedPass),
				UserRoleID:  1,
			}

			err := db.Debug().Model(&model.User{}).Create(&newUser).Error
			if err != nil {
				log.Fatal(err)
			}

			return newUser, nil
		},
	}
}
