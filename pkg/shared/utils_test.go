package shared

import "testing"

func TestRun(t *testing.T) {
	got, _ := Run()
	want := len(RunShared + "\n")

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
