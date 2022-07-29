package goretry

import (
	"time"
)

/*Do is basic retry function. If action return error, it will retry after constant backoff duration.
If backoff = 0, no waiting duration between action retries, same with NoBackoff().*/
func Do(backoff time.Duration, action func() error) {
	std.Do(backoff, action)
}

/*Do is basic retry function. If action return error, it will retry after constant backoff duration.
If backoff = 0, no waiting duration between action retries, same with NoBackoff().*/
func (i *Instance) Do(backoff time.Duration, action func() error) {
	i.Linear(backoff, time.Duration(0), action)
}

// NoBackoff is a shortcut to call Do(0, action). It will retry immediately and don't wait at all.
func NoBackoff(action func() error) {
	std.NoBackoff(action)
}

// NoBackoff is a shortcut to call Do(0, action). It will retry immediately and don't wait at all.
func (i *Instance) NoBackoff(action func() error) {
	i.Do(0, action)
}
