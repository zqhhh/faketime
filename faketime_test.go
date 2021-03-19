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
	now := time.Now()
	ftime := Add(1 * time.Hour)
	ftime.Hook()
	defer ftime.UnHook()
	realTime := ftime.GetRealTime()
	got := realTime.Unix()
	want := now.Unix()
	if got != want {
		t.Errorf("got %d want  %d", got, want)
	}
	t.Logf("realTime:  %s", realTime)
}

func TestDate(t *testing.T) {
	ftime := Date(1970, 1, 1, 0, 0, 0, 0, time.Local)
	ftime.Hook()
	defer ftime.UnHook()
	hookTime := time.Now()
	if hookTime.Year() != 1970 || hookTime.Month() != 1 || hookTime.Day() != 1 {
		t.Errorf("hookTime: %s", hookTime)
	}
	t.Logf("hookTime: year=%d,month=%d,day=%d", hookTime.Year(), hookTime.Month(), hookTime.Day())
}
