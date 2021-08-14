package helpers

import "testing"

func AssertStatusCode(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("expected response code to be %d, got %d", want, got)
	}
}

func AssertError(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %s want %s for error message", got, want)
	}
}

func AssertEqualLength(t *testing.T, got, want int, desc string) {
	t.Helper()

	if got != want {
		t.Errorf("got %d want %d for number of %s", got, want, desc)
	}
}
