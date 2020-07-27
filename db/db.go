package db

import (
	"fmt"
	"net"
	"net/url"
	"os/exec"
	"strings"

	"go-jwt/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
	"gopkg.in/gormigrate.v1"
)

type dbParams struct {
	Type     string
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

var DB *gorm.DB
var params dbParams

func init() {
	u, err := url.Parse(viper.Get("DATABASE_URL").(string))
	if err != nil {
		panic(err.Error())
	}
	host, port, _ := net.SplitHostPort(u.Host)
	p, _ := u.User.Password()
	params = dbParams{u.Scheme, host, port, u.User.Username(), p, u.Path[1:]}
}

func CreateDB() error {
	cmd := exec.Command("createdb", "-p", params.Port, "-h", params.Host, "-U", params.User, "-e", params.DBName)
	return cmd.Run()
}

func openDB() (err error) {
	DB, err = gorm.Open(params.Type,
		fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v",
			params.Host,
			params.Port,
			params.User,
			params.DBName,
			params.Password,
		),
	)
	return
}

func ConnectDB() {
	if err := openDB(); err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf("\"%v\" does not exist", params.DBName)) {
			if err = CreateDB(); err != nil {
				panic(err.Error())
			}
			if err = openDB(); err != nil {
				panic(err.Error())
			}
		} else {
			panic(err.Error())
		}
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

	err := m.Migrate()
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
