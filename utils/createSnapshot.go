package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

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
	