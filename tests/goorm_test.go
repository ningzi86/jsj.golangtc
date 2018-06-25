package tests

import (
         "fmt"
         _ "github.com/go-sql-driver/mysql"
         "github.com/jinzhu/gorm"
        //  "time"
        "testing"
 )
 

 type User struct {
         Id          int    `sql:"AUTO_INCREMENT"`
         Name    string `sql:"varchar(50);unique"`
         //Created_On  time.Time
         Email     string `sql:"type:varchar(50)"`
         Age int `sql:"type:int"`
         Is_active  string `sql:"type:bit"`
 }

 func Test_Gorm(t *testing.T) {

         dbConn, err := gorm.Open("mysql", "root:zaq123wsx@tcp(127.0.0.1:53306)/beego_golangtc?charset=utf8&parseTime=true")

         if err != nil {
                 fmt.Println(err)
         }

        dbConn.LogMode(true)

         // init
         dbConn.DB()
         dbConn.DB().Ping()
         dbConn.DB().SetMaxIdleConns(10)
         dbConn.DB().SetMaxOpenConns(100)

         user := User{} // a record

         // get first record
         dbConn.Table("user").First(&user)

         fmt.Println(user)

        //  fmt.Println("------------------------------")
        //  // get all records
        //  users := []User{} // a slice

        //  //dbConn.Find(&users)
        //  //fmt.Println(users)


        //  // get records from record offset 1 to 10
        //  // useful for pagination purpose! - just need to calculate the offset and limit dynamically

        //  dbConn.Limit(10).Offset(0).Order("id asc").Find(&users)

        //  for _, v := range users {
        //          fmt.Println("Id : ", v.Id)
        //          fmt.Println("Username : ", v.Name)
        //         //  fmt.Println("Created On : ", v.Created_On)
        //          fmt.Println("Email : ", v.Email)
        //          fmt.Println("Is_active : ", v.Is_active)
        //          fmt.Println("Age : ", v.Age)
        //  }

 }