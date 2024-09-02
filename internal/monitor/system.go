package monitor

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/process"
)

type SystemInfo struct {
	Uptime        string
	Users         int
	LoadAverage   [3]float64
	CPUUsage      [8]float64
	TotalTasks    int
	RunningTasks  int
	SleepingTasks int
	ZombieTasks   int
	MemTotal      uint64
	MemUsed       uint64
	MemFree       uint64
	SwapTotal     uint64
	SwapUsed      uint64
	SwapFree      uint64
}

type Monitor struct{}

func NewMonitor() *Monitor {
	return &Monitor{}
}

func (m *Monitor) GetSystemInfo() *SystemInfo {
	uptime, err := host.Uptime()
	if err != nil {
		fmt.Println("Error retrieving uptime:", err)
		return nil
	}

	uptimeStr := (time.Duration(uptime) * time.Second).String()

	users, err := host.Users()
	if err != nil {
		fmt.Println("Error retrieving users:", err)
		return nil
	}

	avg, err := load.Avg()
	if err != nil {
		fmt.Println("Error retrieving load average:", err)
		return nil
	}

	cpuUsage, err := cpu.Percent(0, false)
	if err != nil {
		fmt.Println("Error retrieving CPU usage:", err)
		return nil
	}

	vmStat, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Error retrieving memory stats:", err)
		return nil
	}

	swapStat, err := mem.SwapMemory()
	if err != nil {
		fmt.Println("Error retrieving swap stats:", err)
		return nil
	}

	processes, err := process.Processes()
	if err != nil {
		fmt.Println("Error retrieving processes:", err)
		return nil
	}

	var runningTasks, sleepingTasks, zombieTasks int
	for _, proc := range processes {
		status, err := proc.Status()
		if err != nil {
			continue
		}

		switch status {
		case "R":
			runningTasks++
		case "S", "D":
			sleepingTasks++
		case "Z":
			zombieTasks++
		}
	}

	return &SystemInfo{
		Uptime:        uptimeStr,
		Users:         len(users),
		LoadAverage:   [3]float64{avg.Load1, avg.Load5, avg.Load15},
		CPUUsage:      [8]float64{cpuUsage[0], 0, 0, 100 - cpuUsage[0], 0, 0, 0, 0},
		TotalTasks:    len(processes),
		RunningTasks:  runningTasks,
		SleepingTasks: sleepingTasks,
		ZombieTasks:   zombieTasks,
		MemTotal:      vmStat.Total,
		MemUsed:       vmStat.Used,
		MemFree:       vmStat.Free,
		SwapTotal:     swapStat.Total,
		SwapUsed:      swapStat.Used,
		SwapFree:      swapStat.Free,
	}
}
