package tests

import "testing"
import "fmt"

type DataCommission int32

const (
	DataCommission_DataCommissionNoSetting DataCommission = 0
	DataCommission_Personal                DataCommission = 1
	DataCommission_Department              DataCommission = 2
	DataCommission_DepartmentAndSub        DataCommission = 3
	DataCommission_All                     DataCommission = 4
)

var DataCommission_desc = map[int32]string{
	0: "未设置",
	1: "个人",
	2: "部门",
	3: "部门及子部门",
	4: "所有",
}

var DataCommission_name = map[int32]string{
	0: "DataCommissionNoSetting",
	1: "Personal",
	2: "Department",
	3: "DepartmentAndSub",
	4: "All",
}
var DataCommission_value = map[string]int32{
	"DataCommissionNoSetting": 0,
	"Personal":                1,
	"Department":              2,
	"DepartmentAndSub":        3,
	"All":                     4,
}

func Test_Enum01(t *testing.T) {

   var desc =  DataCommission_desc[int32(DataCommission_Personal)]
   fmt.Println(desc)
   
}
