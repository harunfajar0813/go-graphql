package field

import (
	"log"

	"github.com/jinzhu/gorm"

	"github.com/graphql-go/graphql"
	"graphi/domain/model"
)

var event = graphql.NewObject(graphql.ObjectConfig{
	Name: "Event",
	Fields: graphql.Fields{
		"id":          &graphql.Field{Type: graphql.ID},
		"name":        &graphql.Field{Type: graphql.String},
		"description": &graphql.Field{Type: graphql.String},
		"address":     &graphql.Field{Type: graphql.String},
		"startEvent":  &graphql.Field{Type: graphql.DateTime},
		"price":       &graphql.Field{Type: graphql.Int},
		"userId":      &graphql.Field{Type: graphql.Int},
	},
	Description: "Events data",
})

// Query
func GetEvents(db *gorm.DB) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(event),
		Resolve: func(p graphql.ResolveParams) (i interface{}, err error) {
			var e []*model.Event
			if err := db.Find(&e).Error; err != nil {
				log.Fatal(err)
			}
			return e, nil
		},
		Description: "get events",
	}
}

func GetEvent(db *gorm.DB) *graphql.Field {
	return &graphql.Field{
		Type: event,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, err error) {
			var e []*model.Event

			id, ok := p.Args["id"].(int)
			if ok {
				if err := db.First(&e, id).Error; err != nil {
					log.Fatal(err)
					return nil, err
				}
			}
			return e[0], nil
		},
		Description: "get event by id",
	}
}

// mutation
func CreateEvent(db *gorm.DB) *graphql.Field {
	return &graphql.Field{
		Type: event,
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"description": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
	}
}
