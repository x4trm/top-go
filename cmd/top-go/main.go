package main

import (
	"flag"
	"time"
	"top-go/internal/display"
	"top-go/internal/monitor"
)

func main() {
	numProcesses := flag.Int("n", 20, "Number of processes to display")
	flag.Parse()

	monitor := monitor.NewMonitor()
	renderer := display.NewRenderer()

	for {
		sysInfo := monitor.GetSystemInfo()
		if sysInfo == nil {
			continue
		}

		procInfo := monitor.GetTopProcesses(*numProcesses)
		if procInfo == nil {
			continue
		}

		renderer.Render(sysInfo, procInfo)
		time.Sleep(1 * time.Second)
	}
}
