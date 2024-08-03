package iterates

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("exptected '%q', but actual is '%q'", repeated, expected)
	}

	actual := Repeat("b", 2)
	expected2 := "bb"

	if actual != expected2 {
		t.Errorf("expected '%q', but actual is '%q'", repeated, expected2)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 3)
	}
}