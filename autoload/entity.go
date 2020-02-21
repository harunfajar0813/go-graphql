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

	err = db.Debug().DropTableIfExists(&model.Event{}, &model.User{}).Error
	if err != nil {
		log.Fatal(err)
	}

	err = db.Debug().DropTableIfExists(&model.UserRole{}).Error
	if err != nil {
		log.Fatal(err)
	}

	err = db.Debug().AutoMigrate(&model.User{}, &model.Event{}, &model.UserRole{}).Error
	if err != nil {
		log.Fatal(err)
	}

	err = db.Debug().Model(&model.Event{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatal(err)
	}

	err = db.Debug().Model(&model.User{}).AddForeignKey("user_role_id", "user_roles(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatal(err)
	}
}
