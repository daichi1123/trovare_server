package Helper

import "testing"

func AssertCorrectMsg(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func AssertCorrectBool(t testing.TB, got, want bool) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
