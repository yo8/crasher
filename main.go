package main

import (
	"fmt"
	"runtime"
	"time"
)

const (
	chunkSizeKB  = 1024
	chunkSizeMB  = 1024 * 1024
	chunkSizeGB  = 1024 * 1024 * 1024
	intervalTime = time.Second * 1
	checkStep    = 50
)

var (
	totalSize = int64(0)
	memCol    = make([][]byte, 0)
)

func allocMemory(size int64) {
	memCol = append(memCol, make([]byte, 0, size))
	totalSize += size
}

func getMemoryStatus() {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	fmt.Printf("GC #%d, Sys: %.2f MiB, Alloc: %.2f MiB, Total: %.2f MiB\n",
		mem.NumGC,
		float64(mem.Sys)/1024/1024,
		float64(mem.Alloc)/1024/1024,
		float64(mem.TotalAlloc)/1024/1024,
	)
}

func main() {
	fmt.Println("Hello ❤ World!")
	for i := 1; i <= 1024; i++ {
		allocMemory(chunkSizeMB)
		fmt.Printf("done #%d - total: %d\n", i, totalSize)

		if i%checkStep == 0 {
			getMemoryStatus()
			time.Sleep(intervalTime)
		}
	}

	getMemoryStatus()
	fmt.Println("done!")
}
