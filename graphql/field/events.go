package field

import (
	"errors"
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
		"balance":     &graphql.Field{Type: graphql.String},
		"timeEvent":   &graphql.Field{Type: graphql.String},
		"user": &graphql.Field{Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "EventOrganizer",
			Fields: graphql.Fields{
				"id":    &graphql.Field{Type: graphql.ID},
				"name":  &graphql.Field{Type: graphql.String},
				"email": &graphql.Field{Type: graphql.String},
				"phone": &graphql.Field{Type: graphql.String},
			},
		})},
	},
	Description: "Events data",
})

// Query
func GetEvents(db *gorm.DB) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(event),
		Resolve: func(p graphql.ResolveParams) (i interface{}, err error) {
			var e []*model.Event

			if err := db.Debug().Find(&e).Error; err != nil {
				return err, err
			}

			for _, event := range e {
				if err := db.Debug().First(&event.User, event.UserID).Error; err != nil {
					return err, err
				}
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
			var e model.Event

			id, ok := p.Args["id"].(int)
			if ok {
				if err := db.First(&e, id).Error; err != nil {
					return err, err
				}
				if err := db.Debug().First(&e.User, e.UserID).Error; err != nil {
					return err, err
				}

				var totalStock int
				if err := db.Model(&model.Invoice{}).Where("event_id = ?", e.ID).Count(&totalStock).Error; err != nil {
					return err, err
				} else {
					e.Stock = e.Stock - totalStock
				}
			}
			return e, nil
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
				return nil, errors.New("cannot set price is under 0")
			} else {
				var result int64
				db.Table("users").
					Where("users.user_role_id = ?", 1).
					Where("users.id = ?", userId).
					Count(&result)
				newEvent := &model.Event{
					Name:        name,
					Description: description,
					Address:     address,
					StartEvent:  startEvent,
					Price:       price,
					Stock:       stock,
					UserID:      userId,
				}
				if result == 0 {
					return errors.New("not found"), errors.New("not found")
				}
				err = db.Debug().Model(&model.Event{}).Create(newEvent).Error
				if err != nil {
					return err, err
				}
				return newEvent, nil
			}
		},
	}
}
