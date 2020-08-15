package entity

import (
	"testing"
)

func MiscTests(t *testing.T) {
	ent := New("Test")
	{
		ent.Set("floatTest", 3.141529)
		ent.Set("intTest", 3141529)
		_, err := ent.GetFloat64("floatTest")
		if err != nil {
			t.Fatal(err)
		}

		piVal, err := ent.GetString("floatTest")
		if err != nil {
			t.Fatal(err)
		}

		if piVal != "3.141529" {
			t.Fatal("piVal did not match expectation")
		}

		intval, err := ent.GetInt64("intTest")
		if err != nil {
			t.Fatal(err)
		}
		if intval != 3141529 {
			t.Fatal("intval does not match")
		}
	}

	// had a bug in GetFloat64 due to a bad type cast in switch
	{
		ent.Set("intTest", 3141529)
		_, err := ent.GetFloat64("intTest")
		if err != nil {
			t.Fatal(err)
		}
	}

	{
		_, err := ent.GetInt64("intTest")
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestEntity(t *testing.T) {
	ent := New("person")

	ent.Set("name", "Ryan")
	ent.Set("age", 30)
	ent.Set("date_of_birth", "01-01-1970")
	ent.Set("favorite_number", 3.141529)

	ent.Set("age", 31)
}

func TestEntityMissingKey(t *testing.T) {
	ent := New("person")
	ent.Set("name", "Ryan")

	_, err := ent.GetInt64("age")
	if err != ErrKeyIsMissing {
		t.Fatal("age should be missing")
	}
}

func TestString(t *testing.T) {
	ent := New("person")
	ent.Set("age", 42)

	ageStr, err := ent.GetString("age")
	if err == ErrKeyIsMissing {
		t.Fatal("age shouldn't be missing")
	}

	if ageStr != "42" {
		t.Fatal("age does not match what we expected")
	}
}
