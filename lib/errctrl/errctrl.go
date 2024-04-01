package errctrl

import "github.com/morikuni/failure"

const (
	NotFound        failure.StringCode = "NotFound"
	InvalidArgument failure.StringCode = "InvalidArgument"
	Internal        failure.StringCode = "Internal"
)

func Must[T any](ret T, err error) T {
	if err != nil {
		panic(err)
	}
	return ret
}

func MustExec(err error) {
	if err != nil {
		panic(err)
	}
}
