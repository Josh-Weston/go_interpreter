package object

import "testing"

func TestStringHashKey(t *testing.T) {
	hello1 := &String{Value: "Hello World"}
	hello2 := &String{Value: "Hello World"}
	diff1 := &String{Value: "My name is johnny"}
	diff2 := &String{Value: "My name is johnny"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("strings with same contnet have different hash keys")
	}

	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("strings with different content have same hash keys")
	}
}

func TestBooleanHashKey(t *testing.T) {
	bool1 := &Boolean{Value: true}
	bool2 := &Boolean{Value: true}
	diff1 := &Boolean{Value: false}
	diff2 := &Boolean{Value: false}

	if bool1.HashKey() != bool2.HashKey() {
		t.Errorf("booleans with same content have different hash keys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("booleans with same contnet have different hash keys")
	}

	if bool1.HashKey() == diff1.HashKey() {
		t.Errorf("booleans with different content have same hash keys")
	}
}

func TestInterHashKey(t *testing.T) {
	hello1 := &Integer{Value: 100}
	hello2 := &Integer{Value: 100}
	diff1 := &Integer{Value: 200}
	diff2 := &Integer{Value: 200}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("integers with same content have different hash keys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("integers with same contnet have different hash keys")
	}

	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("integers with different content have same hash keys")
	}
}
