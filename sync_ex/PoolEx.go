package sync_ex

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

var bufPool = sync.Pool{
	// function to execute when no instance of a buffer is not found
	New: func() interface{} {
		fmt.Println("allocating new bytes.Buffer")
		return new(bytes.Buffer)
	},
}

func log(w io.Writer, val string) {
	// type cast the return to the *bytes.Buffer
	b := bufPool.Get().(*bytes.Buffer)

	b.Reset()

	b.WriteString(time.Now().Format("15:04:05"))
	b.WriteString(" : ")
	b.WriteString(val)
	b.WriteString("\n")

	w.Write(b.Bytes())

	// put buffer back into the buffer pool
	bufPool.Put(b)
}

func PoolEx() {
	log(os.Stdout, "debug-string1")
	log(os.Stdout, "debug-string2")
}
