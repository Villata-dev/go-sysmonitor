package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func getSystemInfo() string {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "N/A"
	}
	info := fmt.Sprintf("System: %s / %s\nHostname: %s", runtime.GOOS, runtime.GOARCH, hostname)
	return info
}

func getMemoryInfo() string {
	if runtime.GOOS == "linux" {
		data, err := os.ReadFile("/proc/meminfo")
		if err == nil {
			return parseMemInfo(string(data))
		}
	}

	// Fallback for non-Linux systems or if /proc/meminfo is unreadable
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return fmt.Sprintf("Memory (Go runtime):\nAlloc = %v MiB\nTotalAlloc = %v MiB\nSys = %v MiB\nNumGC = %v",
		m.Alloc/1024/1024, m.TotalAlloc/1024/1024, m.Sys/1024/1024, m.NumGC)
}

func getDiskInfo() string {
	path := "/"
	if runtime.GOOS == "windows" {
		path = "C:"
	}

	var stat syscall.Statfs_t
	err := syscall.Statfs(path, &stat)
	if err != nil {
		return "Disk: N/A"
	}

	total := stat.Blocks * uint64(stat.Bsize)
	free := stat.Bfree * uint64(stat.Bsize)

	return fmt.Sprintf("Disk: %dGB free / %dGB Total", free/1024/1024/1024, total/1024/1024/1024)
}

func parseMemInfo(data string) string {
	lines := strings.Split(data, "\n")
	var totalMem, freeMem uint64
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}
		key := parts[0]
		value := parts[1]

		var err error
		switch key {
		case "MemTotal:":
			totalMem, err = strconv.ParseUint(value, 10, 64)
			if err != nil {
				totalMem = 0
			}
		case "MemFree:":
			freeMem, err = strconv.ParseUint(value, 10, 64)
			if err != nil {
				freeMem = 0
			}
		}
	}
	return fmt.Sprintf("Memory (System):\nTotal: %d MiB\nFree: %d MiB", totalMem/1024, freeMem/1024)
}

func main() {
	for {
		// Clear screen
		fmt.Print("\033[H\033[2J")

		// Print dashboard
		fmt.Println("--- Go System Monitor ---")
		fmt.Println("Time:", time.Now().Format("15:04:05"))
		fmt.Println("-------------------------")
		fmt.Println(getSystemInfo())
		fmt.Println("-------------------------")
		fmt.Println(getMemoryInfo())
		fmt.Println("-------------------------")
		fmt.Println(getDiskInfo())
		fmt.Println("-------------------------")

		time.Sleep(2 * time.Second)
	}
}
