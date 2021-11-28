## Go Race Detector
- Go provides a race detector tool for finding race conditions in Go code.
    1. `go test -race <pkg>` - test the package
    2. `go run -race <src file>` - compile nd run the program
    3. `go build -race <mycmd>` - build the command
    4. `go install -race <pkg>` - install the package
- The binary build needs to be race enabled
- When racy behavior is detected a warning is printed.
- Race enabled binary will be 10 times slower and consume 10 times more memory.
---