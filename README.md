Duration Check
===============

A Go linter to detect cases where two `time.Duration` values are being multiplied in possibly erroneous ways.

For example, consider the following snippet:

```go
func waitFor(someDuration time.Duration) {
    timeToWait := someDuration * time.Second
    time.Sleep(timeToWait)
}
```

Although the above code would compile without any errors, the behaviour is most likely to be incorrect. A caller would
reasonably expect `waitFor(5 * time.Seconds)` to wait for ~5 seconds but they would end up waiting for ~1,388,889 hours.

A majority of these problems would be spotted almost immediately but some could still slip through unnoticed. Hopefully
this linter will help catch those rare cases before they cause a production issue.

See the [test cases](testdata/src/a/a.go) for more examples of the types of errors detected by the linter.


Installation
-------------

Requires Go 1.11 or above.

```
go get -u github.com/charithe/durationcheck/cmd/durationcheck
```

Usage
-----

Invoke `durationcheck` with your package name

```
durationcheck ./...
# or
durationcheck github.com/you/yourproject/...
```
