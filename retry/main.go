package retry

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var (
	DefaultTiming Timing = Static{
		Timeout: 10 * time.Second,
		Delay:   500 * time.Millisecond,
	}
)

type Timing interface {
	GetTimeout() time.Duration
	Before(retry int)
}

type Static struct {
	Timeout time.Duration
	Delay   time.Duration
}

func (d Static) GetTimeout() time.Duration {
	return d.Timeout
}

func (d Static) Before(retry int) {
	if retry > 0 {
		time.Sleep(d.Delay)
	}
}

type Backoff struct {
	Timeout time.Duration
}

func (d Backoff) GetTimeout() time.Duration {
	return d.Timeout
}

// 0 = 0s, 1 = 1s, 2 = 2s, 3 = 4s, 4 = 8s, 5 = 17s,
// 6 = 32s, 7 = 1m5s, 8 = 2m9s, 9 = 4m23s, 10 = 8m58s
func (d Backoff) Before(retry int) {
	backoff := float64(uint(1) << (uint(retry) - 1))
	backoff += backoff * (0.1 * rand.Float64())
	time.Sleep(time.Second * time.Duration(backoff))
}

func RecoverFunc(function func()) func() error {
	return func() (err error) {
		defer func() {
			if value := recover(); value != nil {
				switch errorValue := value.(type) {
				case error:
					err = errorValue
				default:
					err = errors.New(fmt.Sprint(value))
				}
			}
		}()
		function()
		return
	}
}

func FuncPanic(t Timing, function func()) error {
	return BaseRerty(t, RecoverFunc(function))
}

func Func(t Timing, function func() error) error {
	return BaseRerty(t, function)
}

func BaseRerty(t Timing, function func() error) error {
	var (
		err      error
		retry    = 0
		start    = time.Now()
		deadline = t.GetTimeout()
	)
	for time.Since(start) < deadline {
		t.Before(retry)
		if err = function(); err == nil {
			return nil
		}
		retry++
	}
	return err
}
