package field

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"

	"graphi/domain/model"
)

var client = graphql.NewObject(graphql.ObjectConfig{
	Name: "Client",
	Fields: graphql.Fields{
		"id":          &graphql.Field{Type: graphql.ID},
		"name":        &graphql.Field{Type: graphql.String},
		"email":       &graphql.Field{Type: graphql.String},
		"phone":       &graphql.Field{Type: graphql.String},
		"password":    &graphql.Field{Type: graphql.String},
	},
	Description: "Users data",
})

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
				Name:        name,
				Email:       email,
				Phone:       phone,
				Password:    string(hashedPass),
				UserRoleID:  2,
			}

			err := db.Debug().Model(&model.User{}).Create(&newUser).Error
			if err != nil {
				log.Fatal(err)
			}

			return newUser, nil
		},
	}
}

