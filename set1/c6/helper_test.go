package c6

import "testing"

func TestHammingDistance(t *testing.T) {
	want := 37
	input1 := "this is a test"
	input2 := "wokka wokka!!!"
	got := HammingDistance(input1, input2)
	if got != want {
		t.Fatalf("HammingDistance(%q, %q) %d; want = %d", input1, input2, got, want)

	}
}
