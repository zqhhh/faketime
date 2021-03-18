package faketime

import (
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	ftime := Add(1 * time.Hour)
	now := time.Now()
	ftime.Hook()
	defer ftime.UnHook()
	hookTime := time.Now()
	got := hookTime.Sub(now)
	want := 1 * time.Hour
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
	t.Logf("before:  %q, after:%q", now, hookTime)
}

func TestUnix(t *testing.T) {

}

func TestGetRealTime(t *testing.T) {

}
