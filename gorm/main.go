package main
import (
	"github.com/google/uuid"
	"github.com/spf13/pflag"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"fmt"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"size:64"`
	Password string `json:"password" gorm:"size:64"`
	UUID string `json:"uuid" gorm:"column:uuid"`
	Email string `json:"email" gorm:"size:32"`
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	u.UUID = uuid.NewString()
	return tx.Save(u).Error
}


var (
	host = pflag.StringP("host","H","127.0.0.1","mysql server addr")
	username = pflag.StringP("username","u","root","mysql username")
	password = pflag.StringP("password","p","root","mysql password")
	database = pflag.StringP("database","d","gtest","mysql database")
	help = pflag.BoolP("help","h",false,"help message")
)

func main()  {

	pflag.CommandLine.SortFlags = false
	pflag.Usage = func() {
		pflag.PrintDefaults()
	}
	pflag.Parse()
	if *help {
		pflag.Usage()
		return
	}

	//dsn := "root:root@tcp(127.0.0.1:3306)/gtest?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
		*username,
		*password,
		*host,
		*database,
		true,
		"Local")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(User{})
	PrintList(db)
}

func PrintList(db *gorm.DB)  {
	users := []*User{}
	db.Find(&users)

	for _, user := range  users {
		fmt.Printf("用户:%s\n",user.Username)
	}

}