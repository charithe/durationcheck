Duration Check
===============

A Go linter that detects cases where two `time.Duration` values are being multiplied in possibly erroneous ways.

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


Installation
-------------

Grab the sources with go-get

```
go get -d github.com/charithe/durationcheck
```

To install a standalone binary in `$GOPATH/bin`:

```
cd $GOPATH/src/github.com/charithe/durationcheck && make install
```


Usage
-----

Invoke `durationcheck` with your package name

```
durationcheck ./...
# or
durationcheck github.com/you/yourproject/...
```
