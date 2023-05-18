package internal

import (
	"os"
	"runtime"
	"strconv"
)

func SetNumCPU() {
	var n = 2
	if s := os.Getenv("GOMAXPROCS"); s != "" {
		n, _ = strconv.Atoi(s)
	}
	runtime.GOMAXPROCS(n)
}
