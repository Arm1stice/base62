package base62

import "testing"

func TestToB62(t *testing.T) {
	answer := ToB62(61)
	if answer != "Z" {
		t.Errorf("Base 62 representation of 61 was incorrect. Expected 'Z', got '%s'", answer)
	}
	answer = ToB62(62)
	if answer != "10" {
		t.Errorf("Base 62 representation of 62 was incorrect. Expected '10', got '%s'", answer)
	}
	answer = ToB62(3845)
	if answer != "101" {
		t.Errorf("Base 62 representation of 3845 was incorrect. Expected '101', got '%s'", answer)
	}
}

func TestFromB62(t *testing.T) {
	answer, err := FromB62("Z")
	if err != nil {
		t.Errorf("Base 62 conversion of 'Z' got an error: %s", err.Error())
	} else {
		if answer != 61 {
			t.Errorf("Base 62 conversion of 'Z' was incorrect. Expected 61, got '%d'", answer)
		}
	}
	answer, err = FromB62("10")
	if err != nil {
		t.Errorf("Base 62 conversion of '10' got an error: %s", err.Error())
	} else {
		if answer != 62 {
			t.Errorf("Base 62 conversion of '10' was incorrect. Expected 62, got '%d'", answer)
		}
	}
	answer, err = FromB62("101")
	if err != nil {
		t.Errorf("Base 62 conversion of '101' got an error: %s", err.Error())
	} else {
		if answer != 3845 {
			t.Errorf("Base 62 conversion of '101' was incorrect. Expected 3845, got '%d'", answer)
		}
	}
}
