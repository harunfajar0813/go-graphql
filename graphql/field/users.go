package field

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"

	"graphi/domain/model"
)

var user = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id":        &graphql.Field{Type: graphql.ID},
			"firstName": &graphql.Field{Type: graphql.String},
			"lastName":  &graphql.Field{Type: graphql.String},
			"email":     &graphql.Field{Type: graphql.String},
			"password":  &graphql.Field{Type: graphql.String},
		},
		Description: "Users data",
	},
)

// query
func GetUsers(db *gorm.DB) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(user),
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			var u []*model.User
			if err := db.Find(&u).Error; err != nil {
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
			var u []*model.User

			id, ok := p.Args["id"].(int)
			if ok {
				if err := db.First(&u, id).Error; err != nil {
					log.Fatal(err)
					return nil, err
				}
			}
			return u[0], nil
		},
		Description: "get user by id",
	}
}

// mutation
func CreateUser(db *gorm.DB) *graphql.Field {
	return &graphql.Field{
		Type: user,
		Args: graphql.FieldConfigArgument{
			"firstName": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"lastName": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"email": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (i interface{}, e error) {

			firstName, _ := params.Args["firstName"].(string)
			lastName, _ := params.Args["lastName"].(string)
			email, _ := params.Args["email"].(string)
			hashedPass, _ := bcrypt.GenerateFromPassword([]byte(params.Args["password"].(string)), bcrypt.DefaultCost)

			newUser := &model.User{
				FirstName: firstName,
				LastName:  lastName,
				Email:     email,
				Password:  string(hashedPass),
			}

			err := db.Debug().Model(&model.User{}).Create(&newUser).Error
			if err != nil {
				log.Fatal(err)
			}

			return newUser, nil
		},
	}
}
