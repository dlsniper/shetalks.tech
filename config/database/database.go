package database

/*import (
  "fmt"

  "github.com/gocraft/dbr"
)

func InitDB() {

  fmt.Println("HELLO\n")

  conn, err := dbr.Open("mysql", "...", nil)

  if err == nil {
    fmt.Println("\n YOU'RE NOT INSANE \n")
  }

  // create a session for each business unit of execution (e.g. a web request or goworkers job)
  conn.NewSession(nil)
} */

// probably stick with dbr

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitDB() {

	fmt.Println("HELLO\n")

	db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")

	if err == nil {
		fmt.Println("\n YOU ARE INSANE \n")
	}

	if err != nil {
		fmt.Println("YOU ARE NOT INSANE")
	}

	db.HasTable("users")

}
