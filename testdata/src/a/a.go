package a

import (
	"b"
	"time"
)

const (
	timeout = 10 * time.Second
	foo     = 10
)

type myStruct struct {
	fieldA int
	fieldB time.Duration
	fieldC *int
}

func validCases() {
	y := 10
	ms := myStruct{fieldA: 10, fieldB: 10 * time.Second, fieldC: func(v int) *int { return &v }(10)}
	intArr := []int{1}

	_ = time.Second * 30

	_ = time.Duration(10) * time.Second

	_ = time.Second * time.Duration(10)

	_ = time.Duration(10+20*5) * time.Second

	_ = time.Second * time.Duration(10+20*5)

	_ = time.Duration((foo + 20)) * time.Second

	_ = time.Second * time.Duration((foo + 20))

	_ = 2 * 24 * time.Hour

	_ = time.Hour * 2 * 24

	_ = -1 * time.Hour

	_ = time.Hour * -1

	_ = time.Duration(y) * time.Second

	_ = time.Second * time.Duration(y)

	_ = time.Duration(someDurationMillis()) * time.Millisecond

	_ = time.Millisecond * time.Duration(someDurationMillis())

	_ = time.Duration(*somePointerDurationMillis()) * time.Millisecond

	_ = time.Millisecond * time.Duration(*somePointerDurationMillis())

	_ = timeout / time.Millisecond

	_ = foo * time.Second

	_ = time.Second * foo

	_ = time.Duration(ms.fieldA) * time.Second

	_ = time.Second * time.Duration(ms.fieldA)

	_ = time.Duration(*ms.fieldC) * time.Second

	_ = time.Second * time.Duration(*ms.fieldC)

	_ = b.SomeInt * time.Second

	_ = time.Second * b.SomeInt

	_ = time.Duration(intArr[0]) * time.Second

	_ = time.Duration(y) * 24 * time.Hour
}

func invalidCases() {
	x := 30 * time.Second
	ms := myStruct{fieldA: 10, fieldB: 10 * time.Second}
	tdArr := []time.Duration{1}

	_ = x * time.Second // want `Multiplication of durations`

	_ = time.Second * x // want `Multiplication of durations`

	_ = timeout * time.Millisecond // want `Multiplication of durations`

	_ = someDuration() * time.Second // want `Multiplication of durations`

	_ = time.Millisecond * someDuration() // want `Multiplication of durations`

	_ = *somePointerDuration() * time.Second // want `Multiplication of durations`

	_ = time.Millisecond * *somePointerDuration() // want `Multiplication of durations`

	_ = (30 * time.Second) * time.Millisecond // want `Multiplication of durations`

	_ = time.Millisecond * (30 * time.Second) // want `Multiplication of durations`

	_ = time.Millisecond * time.Second * 1 // want `Multiplication of durations`

	_ = 1 * time.Second * (time.Second) // want `Multiplication of durations`

	_ = ms.fieldB * time.Second // want `Multiplication of durations`

	_ = time.Second * ms.fieldB // want `Multiplication of durations`

	_ = b.SomeDuration * time.Second // want `Multiplication of durations`

	_ = time.Second * b.SomeDuration // want `Multiplication of durations`

	_ = time.Duration(tdArr[0]) * time.Second // want `Multiplication of durations`
}

func someDuration() time.Duration {
	return 10 * time.Second
}

func someDurationMillis() int {
	return 10
}

func somePointerDuration() *time.Duration {
	v := 10 * time.Second
	return &v
}

func somePointerDurationMillis() *int {
	v := 10
	return &v
}
