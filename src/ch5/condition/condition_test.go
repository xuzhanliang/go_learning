package condition

import "testing"

func TestConditionIf(t *testing.T) {
	if a := 1 == 1; a {
		t.Log("1==1")
	}
}

func TestIfMultiSec(t *testing.T) {
	//if v,err := someFun();err==nil{
	//	t.Log("1==1")
	//}else{
	//	t.Log("2==2")
	//}
}

func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("偶数")
		case 1, 3:
			t.Log("奇数")
		default:
			t.Log("it is not 0-3")
		}
	}
}

func TestSwitchCaseCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("偶数")
		case i%2 == 1:
			t.Log("奇数")
		default:
			t.Log("unknow")
		}
	}
}
