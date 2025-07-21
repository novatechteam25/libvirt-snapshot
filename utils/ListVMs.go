package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

func GetVmList() {
	cmd := exec.Command("virsh", "list", "--all")
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	// Skip header (first 2 lines)
	for _, line := range lines[2:] {
		// Split line into fields
		fields := strings.Fields(line)
		// Check if there are enough fields
		if len(fields) >= 3 {
			id := fields[0]
			name := fields[1]
			state := fields[2]
			fmt.Printf("%s %s %s\n", id, name, state)
		}
	}
}

func GetVmListWithoutSnapshot() {
	cmd := exec.Command("virsh", "list", "--all", "--title", "--without-snapshot")
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	// Skip header (first 2 lines)
	for _, line := range lines[2:] {
		// Split line into fields
		fields := strings.Fields(line)
		// Check if there are enough fields
		if len(fields) >= 3 {
			id := fields[0]
			name := fields[1]
			state := fields[2]
			fmt.Printf("%s %s %s\n", id, name, state)
		}
	}
}

func GetVmListWithSnapshot() {
	cmd := exec.Command("virsh", "list", "--all", "--title", "--with-snapshot")
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	// Skip header (first 2 lines)
	for _, line := range lines[2:] {
		// Split line into fields
		fields := strings.Fields(line)
		// Check if there are enough fields
		if len(fields) >= 3 {
			id := fields[0]
			name := fields[1]
			state := fields[2]
			fmt.Printf("%s %s %s\n", id, name, state)
		}
	}
}
