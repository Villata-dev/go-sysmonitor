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

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
)

func getSystemInfo() string {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "N/A"
	}
	info := fmt.Sprintf("System: %s / %s\nHostname: %s", runtime.GOOS, runtime.GOARCH, hostname)
	return info
}

func Colorize(value, total uint64) string {
	usage := float64(value) / float64(total) * 100
	color := Green
	if usage > 90 {
		color = Red
	} else if usage > 70 {
		color = Yellow
	}
	return fmt.Sprintf("%s%.2f%%%s", color, usage, Reset)
}

func getMemoryInfo() string {
	if runtime.GOOS == "linux" {
		data, err := os.ReadFile("/proc/meminfo")
		if err == nil {
			totalMem, freeMem := parseMemInfo(string(data))
			usedMem := totalMem - freeMem
			return fmt.Sprintf("Memory: %d MiB free / %d MiB Total (Usage: %s)",
				freeMem/1024, totalMem/1024, Colorize(usedMem, totalMem))
		}
	}

	// Fallback for non-Linux systems or if /proc/meminfo is unreadable
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	totalMem := m.Sys / 1024 / 1024
	usedMem := m.Alloc / 1024 / 1024
	freeMem := totalMem - usedMem
	return fmt.Sprintf("Memory: %d MiB free / %d MiB Total (Usage: %s)",
		freeMem, totalMem, Colorize(usedMem, totalMem))
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
	used := total - free

	return fmt.Sprintf("Disk: %dGB free / %dGB Total (Usage: %s)",
		free/1024/1024/1024, total/1024/1024/1024, Colorize(used, total))
}

func parseMemInfo(data string) (uint64, uint64) {
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
	return totalMem, freeMem
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
