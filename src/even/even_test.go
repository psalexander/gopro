package even

import "testing"

//Entering the twilight zone
func TestEven(t *testing.T) {
	if !Even(2) {
		t.Log("2 should be even!")
		t.Fail()
	}
}
