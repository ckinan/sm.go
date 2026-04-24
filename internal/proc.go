package internal

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

type Process struct {
	Pid     int // process id
	Ppid    int // parent process id
	Name    string
	State   string // process state (R=running, S=sleeping, Z=zombie, etc)
	Threads int
	RssKB   int // actual RAM used (in kB)
}

func ListProcess() ([]Process, error) {
	var processes []Process
	procDirs, err := os.ReadDir("/proc")
	if err != nil {
		return []Process{}, err
	}

	for _, entry := range procDirs {
		pid, err := strconv.Atoi(entry.Name())

		if err != nil {
			slog.Debug("not a proc", "dir name", entry.Name())
		} else if entry.IsDir() {
			process := Process{}
			process.Pid = pid

			slog.Debug("reading dir", "name", pid)
			file, err := os.Open(fmt.Sprintf("/proc/%d/status", pid))
			if err != nil {
				return []Process{}, err
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)

			for scanner.Scan() {
				line := scanner.Text()
				if strings.HasPrefix(line, "PPid:") {
					var ppidStr string
					ppidStr, err = extractFieldFromLine(line)
					if err != nil {
						return []Process{}, err
					}
					process.Ppid, err = strconv.Atoi(ppidStr)
					if err != nil {
						return []Process{}, err
					}
				}
				if strings.HasPrefix(line, "Name:") {
					process.Name, err = extractFieldFromLine(line)
					if err != nil {
						return []Process{}, err
					}
				}
				if strings.HasPrefix(line, "State:") {
					process.State, err = extractFieldFromLine(line)
					if err != nil {
						return []Process{}, err
					}
				}
				if strings.HasPrefix(line, "Threads:") {
					var threadsStr string
					threadsStr, err = extractFieldFromLine(line)
					if err != nil {
						return []Process{}, err
					}
					process.Threads, err = strconv.Atoi(threadsStr)
					if err != nil {
						return []Process{}, err
					}
				}
				if strings.HasPrefix(line, "VmRSS:") {
					var rsskbStr string
					rsskbStr, err = extractFieldFromLine(line)
					if err != nil {
						return []Process{}, err
					}
					process.RssKB, err = strconv.Atoi(rsskbStr)
					if err != nil {
						return []Process{}, err
					}
				}
			}
			processes = append(processes, process)
		}
	}
	return processes, nil
}
