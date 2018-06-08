package squash

import (
	"testing"
)

func Test_squash(t *testing.T) {
	s := []byte("R 	äk	s	mörgås中  	 国")
	got := squash(s)
	want := "Räksmörgås中国"
	if got != want {
		t.Errorf("got %v, want %v", string(got), want)
	}
}
