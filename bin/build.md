Kgen dosen't depend on any syscalls so it's should be theoretically fully crossplatform.

To build a debug version run from the project directory:

> `go build -C . -o ./bin`

Or if you have task installer:

> `task build`

To build a release version run from the project directory:

> `go build -C . -o ./bin -ldflags "-s -w"`

Or if you have task installer:

> `task build-r`