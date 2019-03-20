package string

import (
	"strconv"
	"strings"
	"testing"
)

/**
len()字符数或者byte数
一个中文三个byte

string是不可变的string slice
Cannot assign to 's[1]'

常用字符串处理的函数
strings包 https://golang.org/pkg/strings/
strconv包 https://golang.org/pkg/strconv/
*/
func TestString(t *testing.T) {
	var s string
	t.Log(s)
	s = "hello"
	t.Log(len(s))
	s = "\xE4\xB8\xA5"
	t.Log(s)
	s = "中"
	t.Log(len(s))

	//s[1] = '3'
	c := []rune(s)
	t.Log(len(c))
	t.Logf("%s unicode %x", s, c[0])
	t.Logf("%s UTF8 %x", s, s)
}

func TestStringToRune(t *testing.T) {
	s := "中华人名共和国"
	for _, c := range s {
		t.Logf("%[1]c %[1]d", c) //[1]的意思是，都是以第一个参数格式化
	}
}

func TestStringFunc(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}
	t.Log(strings.Join(parts, "-"))
}

func TestConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str" + s)
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}
}
