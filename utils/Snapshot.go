package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

// Show info snapshot: name, size, created
func ShowSnapshotsInfo(vmName string) {
	cmd := exec.Command("virsh", "snapshot-list", vmName)
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(lines) < 3 {
		fmt.Println("No snapshots found for VM", vmName)
		return
	}
	fmt.Printf("%-20s %-10s %-25s\n", "Name", "Size", "Created")
	for _, line := range lines[2:] { // Skip header
		fields := strings.Fields(line)
		if len(fields) < 3 {
			continue
		}
		name := fields[0]
		created := fields[2]

		// Get disk path for VM
		diskCmd := exec.Command("virsh", "domblklist", vmName)
		diskOut, err := diskCmd.Output()
		size := "-"
		if err == nil {
			diskLines := strings.Split(strings.TrimSpace(string(diskOut)), "\n")
			for _, diskLine := range diskLines[2:] { // Skip header
				diskFields := strings.Fields(diskLine)
				if len(diskFields) == 2 {
					diskPath := diskFields[1]
					duCmd := exec.Command("du", "-h", diskPath)
					duOut, err := duCmd.Output()
					if err == nil {
						duFields := strings.Fields(string(duOut))
						if len(duFields) > 0 {
							size = duFields[0]
						}
					}
					break // Only show first disk size
				}
			}
		}
		fmt.Printf("%-20s %-10s %-25s\n", name, size, created)
	}
}

func CreateSnapshot(vmName, snapshotName string) {
	cmd := exec.Command("virsh", "snapshot-create-as", vmName, snapshotName)
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("Snapshot created:", string(output))
}

func GetAllSnapshotsforVM(vmName string) {
	cmd := exec.Command("virsh", "snapshot-list", vmName)
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	for _, line := range lines[2:] { // Skip header
		fmt.Println(line)
	}
	fmt.Println("Snapshots for VM", vmName, "retrieved successfully.")
}

func GetInfoAboutSpecificSnapshot(vmName, snapshotName string) {
	cmd := exec.Command("virsh", "snapshot-info", vmName, snapshotName)
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("Snapshot info for", snapshotName, "on VM", vmName, ":\n", string(output))
}

func GetSnapshotDiskSize(vmName string) {
	cmd := exec.Command("virsh", "domblklist", vmName)
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	found := false
	for _, line := range lines[2:] { // Skip header
		fields := strings.Fields(line)
		if len(fields) == 2 {
			diskPath := fields[1]
			sizeCmd := exec.Command("du", "-h", diskPath)
			sizeOut, err := sizeCmd.Output()
			if err == nil {
				fmt.Printf("Disk %s size: %s", diskPath, string(sizeOut))
				found = true
			}
		}
	}
	if !found {
		fmt.Println("No disk found for VM", vmName)
	}
}

func RevertSnapshot(vmName, snapshotName string) {
	cmd := exec.Command("virsh", "snapshot-revert", vmName, snapshotName)
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("Reverted to snapshot:", snapshotName, "for VM:", vmName)
	fmt.Println(string(output))
}
