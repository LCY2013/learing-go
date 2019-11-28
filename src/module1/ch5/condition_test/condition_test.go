package condition_test

import "testing"

func TestIfMultiSec(t *testing.T) {
	if a := 0 == 0; a {
		t.Log("0==0 is ", a)
	}
}

func TestSwitchMultiSec(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2, 4:
			t.Log("偶数")
		case 1, 3:
			t.Log("奇数")
		default:
			t.Log("未知")
		}
	}

	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("Event")
		case i%2 == 1:
			t.Log("odd")
		default:
			t.Log("unknown")
		}
	}
}
