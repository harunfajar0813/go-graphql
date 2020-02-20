package field

import (
	"errors"
	"log"

	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"

	"graphi/domain/model"
)

var event = graphql.NewObject(graphql.ObjectConfig{
	Name: "Event",
	Fields: graphql.Fields{
		"id":          &graphql.Field{Type: graphql.ID},
		"name":        &graphql.Field{Type: graphql.String},
		"description": &graphql.Field{Type: graphql.String},
		"address":     &graphql.Field{Type: graphql.String},
		"startEvent":  &graphql.Field{Type: graphql.String},
		"price":       &graphql.Field{Type: graphql.Int},
		"stock":       &graphql.Field{Type: graphql.Int},
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

// query
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
			"address": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"startEvent": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"price": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"stock": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"userId": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, err error) {
			name, _ := p.Args["name"].(string)
			description, _ := p.Args["description"].(string)
			address, _ := p.Args["address"].(string)
			startEvent, _ := p.Args["startEvent"].(string)
			price, _ := p.Args["price"].(int)
			stock, _ := p.Args["stock"].(int)
			userId, _ := p.Args["userId"].(int)

			if price < 0 {
				log.Fatal(errors.New("cannot set price is under 0"))
				return nil, errors.New("cannot set price is under 0")
			} else {
				newEvent := &model.Event{
					Name:        name,
					Description: description,
					Address:     address,
					StartEvent:  startEvent,
					Price:       price,
					Stock:       stock,
					UserID:      userId,
				}
				err = db.Debug().Model(&model.Event{}).Create(newEvent).Error
				if err != nil {
					log.Fatal(err)
				}
				return newEvent, nil
			}
		},
	}
}
