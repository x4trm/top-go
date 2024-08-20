package main

import (
	"time"
	"top-go/internal/display"
	"top-go/internal/monitor"
)

func main() {
	monitor := monitor.NewMonitor()
	renderer := display.NewRenderer()

	for {
		sysInfo := monitor.GetSystemInfo()
		if sysInfo == nil {
			continue
		}

		procInfo := monitor.GetTopProcesses(20)
		if procInfo == nil {
			continue
		}

		renderer.Render(sysInfo, procInfo)
		time.Sleep(1 * time.Second)
	}
}
