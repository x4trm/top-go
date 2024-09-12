package monitor

import (
	"fmt"
	"sort"

	"github.com/shirou/gopsutil/process"
)

type ProcessInfo struct {
	PID      int32
	Name     string
	User     string
	Priority int32
	CPUUsage float64
	MemUsage float32
	Command  string
	CPUTime  string
}

func (m *Monitor) GetTopProcesses(limit int) []*ProcessInfo {
	processes, err := process.Processes()
	if err != nil {
		fmt.Println("Error retrieving processes:", err)
		return nil
	}

	var processList []*ProcessInfo

	for _, proc := range processes {
		name, err := proc.Name()
		if err != nil {
			continue
		}

		user, err := proc.Username()
		if err != nil {
			continue
		}

		priority, err := proc.Nice()
		if err != nil {
			continue
		}

		cpuPercent, err := proc.CPUPercent()
		if err != nil {
			continue
		}

		memPercent, err := proc.MemoryPercent()
		if err != nil {
			continue
		}

		cpuTimes, err := proc.Times()
		if err != nil {
			continue
		}

		cpuTime := fmt.Sprintf("%.2fs", cpuTimes.Total())

		cmdline, err := proc.Cmdline()
		if err != nil || cmdline == "" {
			cmdline = name
		}

		processList = append(processList, &ProcessInfo{
			PID:      proc.Pid,
			Name:     name,
			User:     user,
			Priority: priority,
			CPUUsage: cpuPercent,
			MemUsage: memPercent,
			Command:  cmdline,
			CPUTime:  cpuTime,
		})
	}

	sort.Slice(processList, func(i, j int) bool {
		return processList[i].CPUUsage > processList[j].CPUUsage
	})

	if limit < len(processList) {
		processList = processList[:limit]
	}

	return processList
}
