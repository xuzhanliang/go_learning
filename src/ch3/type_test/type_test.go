package type_test

import "testing"

type MyInt int64

/**
go语言不支持隐式类型转换，要强类型转换
cannot use a (type int32) as type int64 in assignment
*/
func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	b = int64(a)
	t.Log(a, b)
}

/**
不支持指针运算
invalid operation: aPtr + 1 (mismatched types *int and int)
*/
func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	//aPtr = aPtr + 1
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
	if s == "" {
		t.Log("空")
	}
}
