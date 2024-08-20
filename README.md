# Go System Monitor

## Overview

This project is a simple system monitor tool implemented in Go inspired by the `top` command in Linux. It provides real-time monitoring of system resources, displaying the 20 most CPU-intensive processes along with various system statistics such as CPU usage, memory usage, and load average.

## Features

- **Top Processes:** Displays the top 20 processes by CPU usage.
- **System Information:** Shows uptime, load average, CPU usage breakdown, memory usage, and swap usage.
- **Process Details:** Each process is shown with details including PID, user, priority, CPU usage, memory usage, command line, and CPU time.
- **Real-Time Update:** The monitor refreshes every second to provide real-time updates.

## Installation

To install the Go System Monitor, make sure you have Go installed, then clone this repository and build the project:

```bash
$ git clone https://github.com/x4trm/top-go.git
$ cd top-go/cmd/top-go
$ go build
```

## Usage

Run the compiled binary

```bash
$ ./top-go
```
