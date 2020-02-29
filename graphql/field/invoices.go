package field

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
	"graphi/domain/model"
	"log"
)

var invoice = graphql.NewObject(graphql.ObjectConfig{
	Name: "Invoice",
	Fields: graphql.Fields{
		"userId":  &graphql.Field{Type: graphql.ID},
		"eventId": &graphql.Field{Type: graphql.ID},
	},
	Description: "Invoice data",
})

func CreateInvoice(db *gorm.DB) *graphql.Field {
	return &graphql.Field{
		Type: invoice,
		Args: graphql.FieldConfigArgument{
			"userId": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"eventId": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(params graphql.ResolveParams) (i interface{}, e error) {
			userId, _ := params.Args["userId"].(int)
			eventId, _ := params.Args["eventId"].(int)

			var stockTicketNow int
			db.Table("invoices").
				Where("invoices.event_id = ?", eventId).
				Count(&stockTicketNow)

			var saldoUser int
			var userRole int
			db.Table("balances").
				Joins("join users on balances.user_id=users.id").
				Where("balances.user_id = ?", userId).
				Select("sum(amount) as n, users.user_role_id").
				Row().
				Scan(&saldoUser, &userRole)

			var riwayatTransactionUserId int
			db.Table("invoices").
				Joins("join events on invoices.event_id=events.id").
				Where("invoices.user_id = ?", userId).
				Select("sum(events.price)").
				Row().
				Scan(&riwayatTransactionUserId)

			var hargaTiket int
			var stockTiket int
			db.Table("events").
				Where("events.id = ?", eventId).
				Select("events.price, events.stock").
				Row().
				Scan(&hargaTiket, &stockTiket)

			// 1 orang adalah 1 ticket
			var invoicesCountByUserIDAndByEventID int
			db.Table("invoices").
				Select("sum(invoices.id) as n").
				Where("invoices.user_id = ?", userId).
				Where("invoices.event_id = ?", eventId).
				Row().
				Scan(&invoicesCountByUserIDAndByEventID)

			if stockTiket == stockTicketNow {
				return errors.New("ticket is empty"), errors.New("ticket is empty")
			}

			if invoicesCountByUserIDAndByEventID == 1 {
				return errors.New("ticket is already bought"), errors.New("ticket is already bought")
			}

			if (saldoUser-riwayatTransactionUserId)-hargaTiket < 0 && userRole != 2 {
				return errors.New("buy ticket is denied"), errors.New("buy ticket is denied")
			}

			newInvoice := &model.Invoice{
				EventID: eventId,
				UserID:  userId,
			}

			err := db.Debug().Model(&model.Invoice{}).Create(&newInvoice).Error
			if err != nil {
				log.Fatal(err)
			}

			return newInvoice, nil
		},
	}
}
