package main

import (
	"fmt"
	"runtime"
	"time"
)

const (
	chunkSizeKB = 1024
	chunkSizeMB = 1024 * 1024
	chunkSizeGB = 1024 * 1024 * 1024

	intervalTime = time.Second * 1
	checkStep    = 50
	maxStep      = 1024
	stepSize     = chunkSizeGB
)

var (
	totalSize = int64(0)
	memCol    = make([][]byte, 0)
)

const (
	sizeKB = 1 << ((iota + 1) * 10)
	sizeMB
	sizeGB
	sizeTB
	sizePB
)

func formatSize(size uint64) string {
	switch sf := float64(size); {
	case sf < sizeKB:
		return fmt.Sprintf("%d B", size)
	case sf < sizeMB:
		return fmt.Sprintf("%.2f KiB", sf/sizeKB)
	case sf < sizeGB:
		return fmt.Sprintf("%.2f MiB", sf/sizeMB)
	case sf < sizeTB:
		return fmt.Sprintf("%.2f GiB", sf/sizeGB)
	case sf < sizePB:
		return fmt.Sprintf("%.2f TiB", sf/sizeTB)
	default:
		return fmt.Sprintf("%.2f PiB", sf/sizePB)
	}
}

func allocMemory(size int64) {
	memCol = append(memCol, make([]byte, 0, size))
	totalSize += size
}

func getMemoryStatus() {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	fmt.Printf("GC #%d, Sys: %s, Alloc: %s, Total: %s\n",
		mem.NumGC,
		formatSize(mem.Sys),
		formatSize(mem.Alloc),
		formatSize(mem.TotalAlloc),
	)
}

func main() {
	for i := 1; i <= maxStep; i++ {
		allocMemory(stepSize)
		fmt.Printf("done #%d - total: %d\n", i, totalSize)

		if i%checkStep == 0 {
			getMemoryStatus()
			time.Sleep(intervalTime)
		}
	}

	getMemoryStatus()
	fmt.Println("done!")
}
