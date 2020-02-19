package autoload

import (
	"log"

	"graphi/datastore"
	"graphi/domain/model"
)

func Load() {
	db, err := datastore.NewMyqlDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Debug().DropTableIfExists(&model.Transaction{}, &model.Event{}, &model.Balance{}).Error
	if err != nil {
		log.Fatal(err)
	}

	err = db.Debug().DropTableIfExists(&model.User{}, &model.Client{}).Error
	if err != nil {
		log.Fatal(err)
	}

	err = db.Debug().AutoMigrate(&model.Client{}, &model.User{}, &model.Event{}, &model.Balance{}, &model.Transaction{}).Error
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

	err = db.Debug().Model(&model.Transaction{}).AddForeignKey("client_id", "clients(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatal(err)
	}

	err = db.Debug().Model(&model.Transaction{}).AddForeignKey("event_id", "events(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatal(err)
	}
}
