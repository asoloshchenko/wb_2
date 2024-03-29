package main

import (
	"testing"
	//"wb2/develop/dev02/unpack"
)

func TestUnpackBase(t *testing.T) {
	args := [...]string{
		"a4bc2d5e",
		"abcd",
		"Привет2",
		"🖐2",
		"a10",
		"a0",
	}
	wanted := []string{
		"aaaabccddddde",
		"abcd",
		"Приветт",
		"🖐🖐",
		"aaaaaaaaaa",
		"",
	}

	for i := 0; i < len(args); i++ {
		got, err := Unpack(args[i])
		if err != nil {
			t.Errorf("Unpack(%q) error: %v", args[i], err)
		}
		if got != wanted[i] {
			t.Errorf("Unpack(%q) = %q; want %q", args[i], got, wanted[i])
		}
	}
}

func TestUnpackError(t *testing.T) {
	args := [...]string{
		"45",
		"",
		"45c",
	}

	for i := 0; i < len(args); i++ {
		_, err := Unpack(args[i])
		if err == nil {
			t.Errorf("Unpack(%q); want error", args[i])
		}

	}
}

func TestUnpackEscape(t *testing.T) {
	args := [...]string{
		"qwe\\4\\5",
		"qwe\\45",
		"qwe\\\\5",
	}
	wanted := [...]string{
		"qwe45",
		"qwe44444",
		"qwe\\\\\\\\\\",
	}

	for i := 0; i < len(args); i++ {
		got, err := Unpack(args[i])
		if err != nil {
			t.Errorf("Unpack(%q) error: %v", args[i], err)
		}
		if got != wanted[i] {
			t.Errorf("Unpack(%q) = %q; want %q", args[i], got, wanted[i])
		}
	}
}
