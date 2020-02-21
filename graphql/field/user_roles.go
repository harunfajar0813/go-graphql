package field

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
	"graphi/domain/model"
)

var userRole = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserRole",
	Fields: graphql.Fields{
		"id":   &graphql.Field{Type: graphql.ID},
		"name": &graphql.Field{Type: graphql.String},
	},
	Description: "Users data",
})

func CreateUserRole(db *gorm.DB) *graphql.Field {
	return &graphql.Field{
		Type: userRole,
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (i interface{}, e error) {
			name, _ := params.Args["name"].(string)

			newUserRole := &model.UserRole{
				Name: name,
			}

			err := db.Debug().Model(&model.UserRole{}).Create(&newUserRole).Error
			if err != nil {
				log.Fatal(err)
			}

			return newUserRole, nil
		},
	}
}
