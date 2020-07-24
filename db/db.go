package db

import (
	"fmt"
	"net"
	"net/url"

	"go-jwt/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
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
	fmt.Println(u.Scheme, host, port, u.User.Username(), u.Path[1:], p)
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
	defer DB.Close()
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
