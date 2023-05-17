package internal

import (
	"os"
	"runtime"
	"strconv"
)

func SetNumCPU() {
	var n = runtime.NumCPU()
	if s := os.Getenv("GOMAXPROCS"); s != "" {
		n, _ = strconv.Atoi(s)
	}
	runtime.GOMAXPROCS(n)
}
