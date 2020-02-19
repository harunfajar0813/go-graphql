package autoload

import (
	"log"

	"graphi/domain/model"
	"graphi/datastore"
)

func Load() {
	db, err := datastore.NewMyqlDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Debug().DropTableIfExists(&model.Event{}, &model.User{}, &model.BalanceStatus{}, &model.Balance{}).Error
	if err != nil {
		log.Fatal(err)
	}
	err = db.Debug().AutoMigrate(&model.User{}, &model.Event{}, &model.BalanceStatus{}, &model.Balance{}).Error
	if err != nil {
		log.Fatal(err)
	}

	err = db.Debug().Model(&model.Event{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatal(err)
	}

	err = db.Debug().Model(&model.Balance{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatal(err)
	}
	err = db.Debug().Model(&model.Balance{}).AddForeignKey("balance_status_id", "balances_status(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatal(err)
	}
}
