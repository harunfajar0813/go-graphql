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
			rows, err := db.Table("events").
				Select("events.id, "+
					"events.name, "+
					"events.description, "+
					"events.address,"+
					"events.start_event, "+
					"events.price, "+
					"events.stock, "+
					"users.id, "+
					"users.name, "+
					"users.email, "+
					"users.phone").
				Joins("join users on users.id = events.user_id").
				Where("users.user_role_id = ?", 1).
				Rows()
			defer rows.Close()

			var event model.Event
			for rows.Next() {
				err := rows.Scan(&event.ID,
					&event.Name,
					&event.Description,
					&event.Address,
					&event.StartEvent,
					&event.Price,
					&event.Stock,
					&event.User.ID,
					&event.User.Name,
					&event.User.Email,
					&event.User.Phone)
				if err != nil {
					log.Fatal(err)
				}

				var totalStock int
				if err := db.Model(&model.Invoice{}).Where("event_id = ?", event.ID).Count(&totalStock).Error; err != nil {
					log.Fatal(err)
				} else {
					event.Stock = event.Stock - totalStock
				}

				e = append(e, &event)
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
			var e model.Event

			id, ok := p.Args["id"].(int)
			if ok {
				if err := db.First(&e, id).Error; err != nil {
					log.Fatal(err)
					return nil, err
				}
				if err := db.Table("events").
					Select("events.id, "+
						"events.name, "+
						"events.description, "+
						"events.address, "+
						"events.start_event, "+
						"events.price, "+
						"events.stock, "+
						"users.id, "+
						"users.name, "+
						"users.email, "+
						"users.phone").
					Joins("join users on users.id = events.user_id").
					Where("users.user_role_id = ?", 1).
					Row().Scan(&e.ID, &e.Name, &e.Description, &e.Address, &e.StartEvent, &e.Price, &e.Stock,
					&e.User.ID, &e.User.Name, &e.User.Email, &e.User.Phone); err != nil {
					log.Fatal(err)
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
				log.Fatal(errors.New("cannot set price is under 0"))
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
					log.Fatal("not found")
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
