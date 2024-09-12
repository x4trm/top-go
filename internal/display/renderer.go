package display

import (
	"fmt"
	"time"
	"top-go/internal/monitor"
)

type Renderer struct{}

func NewRenderer() *Renderer {
	return &Renderer{}
}

func (r *Renderer) ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func (r *Renderer) Render(sysInfo *monitor.SystemInfo, procInfo []*monitor.ProcessInfo) {
	r.ClearScreen()

	fmt.Printf("top - %s up %s,  %d user(s),  load average: %.2f, %.2f, %.2f\n", time.Now().Format("15:04:05"),
		sysInfo.Uptime, sysInfo.Users, sysInfo.LoadAverage[0], sysInfo.LoadAverage[1], sysInfo.LoadAverage[2])

	fmt.Printf("Tasks: total: %d, running: %d, sleeping: %d, stopped: %d, zombie: %d\n",
		sysInfo.TotalTasks, sysInfo.RunningTasks, sysInfo.SleepingTasks, 0, sysInfo.ZombieTasks)

	fmt.Printf("%%CPU: %.1f us, %.1f sy, %.1f ni, %.1f id, %.1f wa, %.1f hi, %.1f si, %.1f st\n",
		sysInfo.CPUUsage[0], sysInfo.CPUUsage[1], sysInfo.CPUUsage[2], sysInfo.CPUUsage[3],
		sysInfo.CPUUsage[4], sysInfo.CPUUsage[5], sysInfo.CPUUsage[6], sysInfo.CPUUsage[7])

	fmt.Printf("MiB RAM : %8.1f total, %8.1f free, %8.1f used, %8.1f buff/cache\n",
		float64(sysInfo.MemTotal)/1024/1024, float64(sysInfo.MemFree)/1024/1024,
		float64(sysInfo.MemUsed)/1024/1024, float64(sysInfo.MemTotal-sysInfo.MemUsed-sysInfo.MemFree)/1024/1024)

	fmt.Printf("MiB Swap: %8.1f total, %8.1f free, %8.1f used. %8.1f avail Mem\n",
		float64(sysInfo.SwapTotal)/1024/1024, float64(sysInfo.SwapFree)/1024/1024,
		float64(sysInfo.SwapUsed)/1024/1024, float64(sysInfo.MemTotal-sysInfo.MemUsed)/1024/1024)

	fmt.Printf("%-8s %-12s %-5s %-8s %-8s %-10s %-s\n", "PID", "USER", "PR", "CPU (%)", "MEM (%)", "TIME+", "NAME")
	for _, proc := range procInfo {
		cpuColor := r.getColorForProcUsage(proc.CPUUsage)
		memColor := r.getColorForMemUsage(proc.MemUsage)
		fmt.Printf("%-8d %-12s %-5d %s%-8.2f\033[0m %s%-8.2f\033[0m %-10s %-s\n", proc.PID, proc.User, proc.Priority, cpuColor, proc.CPUUsage, memColor, proc.MemUsage, proc.CPUTime, proc.Name)
	}
}

func (r *Renderer) getColorForProcUsage(cpuUsage float64) string {
	switch {
	case cpuUsage < 50:
		return "\033[32m" // green
	case cpuUsage >= 50 && cpuUsage < 70:
		return "\033[33m" // yellow
	case cpuUsage >= 70 && cpuUsage < 90:
		return "\033[35m" // orange
	default:
		return "\033[31m" // red
	}
}

func (r *Renderer) getColorForMemUsage(memUsage float32) string {
	switch {
	case memUsage < 50:
		return "\033[32m" // green
	case memUsage >= 50 && memUsage < 70:
		return "\033[33m" // yellow
	case memUsage >= 70 && memUsage < 90:
		return "\033[35m" // orange
	default:
		return "\033[31m" // red
	}
}
