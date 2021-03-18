package faketime

import (
	"time"
	_ "unsafe"

	"bou.ke/monkey"
)

//go:linkname unixTime time.unixTime
func unixTime(sec int64, nsec int32) time.Time

//go:linkname now time.now
func now() (sec int64, nsec int32, mono int64)

type Time struct {
	d time.Duration
}

func (t Time) GetRealTime() time.Time {
	sec, nsec, _ := now()
	return unixTime(sec, nsec)
}

func (t Time) Hook() {
	monkey.Patch(time.Now, func() time.Time {
		return t.GetRealTime().Add(t.d)
	})
}

func (t Time) UnHook() {
	monkey.Unpatch(time.Now)
}

func Date() Time {
	return Time{}
}

func Unix(timestamp int64) Time {
	t := time.Unix(timestamp, 0)
	return Time{d: time.Until(t)}
}

func Add(d time.Duration) Time {
	return Time{d: d}
}
