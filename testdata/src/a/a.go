package a

import (
	"time"
)

const (
	timeout = 10 * time.Second
	foo     = 10
)

func validCases() {
	y := 10

	_ = time.Second * 30

	_ = time.Duration(10) * time.Second

	_ = time.Second * time.Duration(10)

	_ = time.Duration(10+20*5) * time.Second

	_ = time.Second * time.Duration(10+20*5)

	_ = 2 * 24 * time.Hour

	_ = time.Hour * 2 * 24

	_ = -1 * time.Hour

	_ = time.Hour * -1

	_ = time.Duration(y) * time.Second

	_ = time.Second * time.Duration(y)

	_ = time.Duration(someDurationMillis()) * time.Millisecond

	_ = time.Millisecond * time.Duration(someDurationMillis())

	_ = timeout / time.Millisecond

	_ = foo * time.Second

	_ = time.Second * foo
}

func invalidCases() {
	x := 30 * time.Second

	_ = x * time.Second // want `Multiplication of durations`

	_ = time.Second * x // want `Multiplication of durations`

	_ = timeout * time.Millisecond // want `Multiplication of durations`

	_ = someDuration() * time.Second // want `Multiplication of durations`

	_ = time.Millisecond * someDuration() // want `Multiplication of durations`

	_ = (30 * time.Second) * time.Millisecond // want `Multiplication of durations`

	_ = time.Millisecond * (30 * time.Second) // want `Multiplication of durations`

	_ = time.Millisecond * time.Second * 1 // want `Multiplication of durations`

	_ = 1 * time.Second * (time.Second) // want `Multiplication of durations`
}

func someDuration() time.Duration {
	return 10 * time.Second
}

func someDurationMillis() int {
	return 10
}
