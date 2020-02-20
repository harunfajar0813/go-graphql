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

	err = db.Debug().DropTableIfExists(&model.Event{}, &model.Balance{}).Error
	if err != nil {
		log.Fatal(err)
	}

	err = db.Debug().DropTableIfExists(&model.User{}).Error
	if err != nil {
		log.Fatal(err)
	}

	err = db.Debug().AutoMigrate(&model.User{}, &model.Event{}, &model.Balance{}).Error
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
}
