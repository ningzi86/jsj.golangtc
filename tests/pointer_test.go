package tests

import "testing"
import "fmt"
import "unsafe"

func _Test_Pointer01(t *testing.T) {

	// 检查指针对象的零值为 nil
	var p *int
	fmt.Println("The zero value of a pointer is: ", p)

	// 指向指针的指针
	pp := &p
	fmt.Printf("The type of a pointer points another pointer is: %T\n", pp)

	// 指针对象赋值
	intVar := 100000000
	p = &intVar
	fmt.Println("After assignment, p is: ", p)
	fmt.Println("The value pointer p points is: ", *p)

	// 使用 unsafe.Pointer 方法可以将一个类型的指针转化为 Pointer
	// Pointer 可以被转化为任意类型的指针。
	// 注意由于 int 为 int32 的别名，占 4 个字节，所以我们将其转化为含有 4 个字节元素的 `byte` 数组指针
	var strP *[4]byte
	strP = (*[4]byte)(unsafe.Pointer(p))
	fmt.Println("After \"(*[4]byte)(unsafe.Pointer(p))\", *[4]byte pointer strP is: ", strP)
	fmt.Println("After \"(*[4]byte)(unsafe.Pointer(p))\", *[4]byte pointer strP points to: ", *strP)

	// 指针指向的对象内容使用 `.` 而不是 `->` 来进行访问
	type User struct {
		name string
	}
	userP := &User{
		"Xiaohui",
	}
	fmt.Println("Before change, The value userP points to is: ", userP.name)
	userP.name = "Ross"
	fmt.Println("After change,  The value userP points to is: ", *userP)

}

func _Test_Pointer02(t *testing.T) {

	var i int
	i = 1
	var p *int
	p = &i

	fmt.Printf("i=%d;p=%d;*p=%d\n", i, p, *p)

	*p = 2
	fmt.Printf("i=%d;p=%d;*p=%d\n", i, p, *p)

	i = 3 // 验证想法
	fmt.Printf("i=%d;p=%d;*p=%d\n", i, p, *p)

}

func Test_Pointer03(t *testing.T)  {
	var p = pointermd{}
	p.aaaa()
	p.bbb()
	p.cccc()

	fmt.Println(p)
}


type pointermd struct{
	v int
}


func (p pointermd) aaaa()  {
	p.v = 1
	fmt.Printf("1:%d\n", p.v)
}

func (p *pointermd) bbb()  {
	fmt.Printf("2:%d\n", p.v)

	p.v = 2008
	fmt.Printf("3:%d\n", p.v)
}

func (p pointermd) cccc()  { 
	fmt.Printf("4:%d\n", p.v)
}
