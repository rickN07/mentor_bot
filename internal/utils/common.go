package utils

import "time"

func DoWithTries(fn func() error, attempts int, delay time.Duration) (err error) {
	for attempts > 0 {
		if err = fn(); err == nil {
			return
		}
println(attempts)
		attempts--
		if attempts > 0 {
			time.Sleep(delay)
			continue
		}

		return
	}

	return
}
