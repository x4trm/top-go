package display

import (
	"fmt"
	"time"
	"top-go/internal/monitor"
)

type Renderer struct {}

func NewRenderer() *Renderer {
    return &Renderer{}
}

func (r *Renderer) Render(sysInfo *monitor.SystemInfo, procInfo []*monitor.ProcessInfo) {
    fmt.Printf("top - %s up %s,  %d user,  load average: %.2f, %.2f, %.2f\n", time.Now().Format("15:04:05"),
        sysInfo.Uptime, sysInfo.Users, sysInfo.LoadAverage[0], sysInfo.LoadAverage[1], sysInfo.LoadAverage[2])

    fmt.Printf("Zadania: razem: %d, działających: %d, śpiących: %d, zatrzymanych: %d, zombie: %d\n",
        sysInfo.TotalTasks, sysInfo.RunningTasks, sysInfo.SleepingTasks, 0, sysInfo.ZombieTasks)

    fmt.Printf("%%CPU: %.1f uż, %.1f sy, %.1f ni, %.1f be, %.1f io, %.1f hi, %.1f si, %.1f sk\n",
        sysInfo.CPUUsage[0], sysInfo.CPUUsage[1], sysInfo.CPUUsage[2], sysInfo.CPUUsage[3],
        sysInfo.CPUUsage[4], sysInfo.CPUUsage[5], sysInfo.CPUUsage[6], sysInfo.CPUUsage[7])

    fmt.Printf("MiB RAM : %8.1f razem, %8.1f wolne, %8.1f użyte, %8.1f buf/cache\n",
        float64(sysInfo.MemTotal)/1024/1024, float64(sysInfo.MemFree)/1024/1024,
        float64(sysInfo.MemUsed)/1024/1024, float64(sysInfo.MemTotal-sysInfo.MemUsed-sysInfo.MemFree)/1024/1024)

    fmt.Printf("MiB Swap: %8.1f razem, %8.1f wolne, %8.1f użyte. %8.1f dost. RAM\n",
        float64(sysInfo.SwapTotal)/1024/1024, float64(sysInfo.SwapFree)/1024/1024,
        float64(sysInfo.SwapUsed)/1024/1024, float64(sysInfo.MemTotal-sysInfo.MemUsed)/1024/1024)

		fmt.Printf("%-8s %-12s %-5s %-8s %-8s %-10s\n", "PID", "USER", "PR", "CPU (%)", "MEM (%)", "TIME+")
		for _, proc := range procInfo {
			fmt.Printf("%-8d %-12s %-5d %-8.2f %-8.2f %-10s\n", proc.PID, proc.User, proc.Priority, proc.CPUUsage, proc.MemUsage, proc.CPUTime)
		}
}
