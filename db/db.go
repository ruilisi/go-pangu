package db

import (
	"fmt"
	"net"
	"net/url"

	"go-jwt/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
	"gopkg.in/gormigrate.v1"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	//replace your database
	u, err := url.Parse(viper.Get("DATABASE_URL").(string))
	if err != nil {
		panic(err.Error())
	}
	host, port, _ := net.SplitHostPort(u.Host)
	p, _ := u.User.Password()
	DB, err = gorm.Open(u.Scheme,
		fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v",
			host,
			port,
			u.User.Username(),
			u.Path[1:],
			p,
		),
	)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	m := gormigrate.New(DB, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "initial",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.User{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("user").Error
			},
		},
	})

	m.InitSchema(func(tx *gorm.DB) error {
		err := tx.AutoMigrate(
			&models.User{},
			// all other tables of you app
		).Error
		if err != nil {
			return err
		}

		return nil
	})

	err = m.Migrate()
	if err != nil {
		panic("Could not migrate: " + err.Error())
	}
}

func FindUserByEmail(email string) models.User {
	var user models.User
	DB.Where("email = ?", email).First(&user)
	return user
}

func FindUserById(id string) models.User {
	var user models.User
	DB.Where("id = ?", id).First(&user)
	return user
}
