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
	got := hookTime.Sub(now) / 1e9
	want := 1 * time.Hour / 1e9
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
	t.Logf("before:  %q, after:%q", now, hookTime)
}

func TestUnix(t *testing.T) {
	var timestamp int64 = 1616165198
	ftime := Unix(timestamp)
	ftime.Hook()
	defer ftime.UnHook()
	hookTime := time.Now()
	if hookTime.Unix() != timestamp {
		t.Errorf("hook error,got:%d want  %d", hookTime.Unix(), timestamp)
	}
	t.Logf("hook time %s", hookTime)
}

func TestGetRealTime(t *testing.T) {

}
