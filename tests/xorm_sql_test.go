package tests

import (
	"fmt"
	"os"
	"time"

	"testing"

	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	_ "github.com/lunny/godbc"
)

type _UC_Student struct {
	StudentID   int    `xorm:"pk not null 'StudentID'"`
	ParentsID   int    `xorm:"not null 'ParentsID'"`
	FamilyID    int    `xorm:"not null FamilyID"`
	StudentName string `xorm:"StudentName"`

	CreateTime time.Time `xorm:"CreateTime"`
	UpdateTime time.Time `xorm:"UpdateTime"`
}

func Test_xorm_01(t *testing.T) {

	File, _ := os.Create("result")
	defer File.Close()
	Engine, err := xorm.NewEngine("odbc", "driver={SQL Server};Server=NING-PC\\NINGZI2014;Database=BBZMv2016.3;uid=sa;pwd=system;")
	if err != nil {
		fmt.Println("新建引擎", err)
		return
	}

	if err := Engine.Ping(); err != nil {
		fmt.Println(err)
		return
	}

	Engine.SetTableMapper(core.SameMapper{})
	Engine.ShowSQL(true)
	// Engine.SetMaxConns(5)
	Engine.SetMaxIdleConns(5)
	//result := new(_UC_Student)
	//lines, _ := Engine.Rows(result)

    st := &_UC_Student{ StudentID:60 }
    has,_ := Engine.Get(st)

    fmt.Println(has)
    fmt.Println(st.StudentName)


	//defer lines.Close()
	//lines.Next()

	// r := new(_UC_Student)
	// for {
	// 	err = lines.Scan(r)
	// 	if err != nil {
	// 		return
	// 	}
	// 	fmt.Println(*r)
	// 	File.WriteString(fmt.Sprintln(*r))
	// 	if !lines.Next() {
	// 		break
	// 	}
	// }
}
