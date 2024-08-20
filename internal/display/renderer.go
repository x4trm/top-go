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

func (r *Renderer) Render(sysInfo *monitor.SystemInfo, procInfo []*monitor.ProcessInfo) {
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

	fmt.Printf("%-8s %-12s %-5s %-8s %-8s %-10s\n", "PID", "USER", "PR", "CPU (%)", "MEM (%)", "TIME+")
	for _, proc := range procInfo {
		fmt.Printf("%-8d %-12s %-5d %-8.2f %-8.2f %-10s\n", proc.PID, proc.User, proc.Priority, proc.CPUUsage, proc.MemUsage, proc.CPUTime)
	}
}
