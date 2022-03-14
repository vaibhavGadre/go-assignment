package main

import (
	"testing"
)

func TestIsAnagram(t *testing.T)  {
	if isAnagram("eat", "tea") != true {
		t.Error("Angrams. Expected false to be equal true")
	}
}

func Test_isAnagram(t *testing.T) {
	type args struct {
		ip1 string
		ip2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Anagram case",  args{"elbow",  "below"}, true},
		{"Length mismatch case", args{ "boycott", "cotton"}, false},
		{"Not Anagram case", args{ "state", "toast"}, false},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := isAnagram(tt.args.ip1, tt.args.ip2); got != tt.want {
				t.Errorf("isAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
