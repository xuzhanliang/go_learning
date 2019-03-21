package encapsulation

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

func TestCreate(t *testing.T) {
	e := Employee{"0", "Bob", 25}
	e1 := Employee{Id: "1", Name: "Mike", Age: 30}
	e2 := new(Employee)
	e2.Id = "2"
	e2.Name = "Rose"
	e2.Age = 35
	t.Log(e)
	t.Log(e1)
	t.Log(e2)
	t.Logf("e is %T", &e)
	t.Logf("e2 is %T", e2)
	t.Log(e2.String1())
	t.Log(e2.String2())
}

//第一种定义方式在实例对应方法被调用时，实例的成员会进行值复制
func (e Employee) String1() string {
	fmt.Printf("Address is %x", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

//通常情况下为了避免内存拷贝我们使用第二种定义方式
func (e *Employee) String2() string {
	fmt.Printf("Address is %x", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

func TestStructOperations(t *testing.T) {
	e := Employee{"0", "Bob", 25}
	fmt.Printf("Address is %x", unsafe.Pointer(&e.Name))
	t.Log(e.String1())
	//t.Log(e.String2())
}
