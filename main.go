package main

import (
	"fmt"

	"github.com/libvirt/libvirt-go"
	"github.com/novatechteam25/libvirt-snapshot/pkg"
	"github.com/novatechteam25/libvirt-snapshot/utils"
)

func main() {
	conn, err := libvirt.NewConnect(pkg.DefaultQemuURI)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	for {
		fmt.Println("===========================")
		fmt.Println("What do you want to do?")
		fmt.Println("1. List all VMs")
		fmt.Println("2. List VMs without snapshots")
		fmt.Println("3. List VMs with snapshots")
		fmt.Println("4. Get all snapshots for a VM")
		fmt.Println("5. Create a snapshot for a VM")
		fmt.Println("6. Get info about a specific snapshot")
		fmt.Println("7. Get snapshot size")
		fmt.Println("8. Revert snapshot")
		fmt.Println("9. Delete snapshot")
		fmt.Println("===========================")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Listing all VMs:")
			utils.GetVmList()
			return
		case 2:
			fmt.Println("Listing VMs without snapshots:")
			utils.GetVmListWithoutSnapshot()
			return
		case 3:
			fmt.Println("Listing VMs with snapshots:")
			utils.GetVmListWithSnapshot()
			return
		case 4:
			var vmName string
			fmt.Print("Enter VM name to get snapshots: ")
			fmt.Scan(&vmName)
			fmt.Println("Getting all snapshots for VM:", vmName)
			utils.GetAllSnapshotsforVM(vmName)
			return
		case 5:
			var vmName, snapshotName string
			fmt.Print("Enter VM name to create snapshot: ")
			fmt.Scan(&vmName)
			fmt.Print("Enter snapshot name: ")
			fmt.Scan(&snapshotName)
			fmt.Println("Creating snapshot for VM:", vmName, "with name:", snapshotName)
			utils.CreateSnapshot(vmName, snapshotName)
			return
		case 6:
			var vmName string
			fmt.Print("Enter VM name to get snapshot info: ")
			fmt.Scan(&vmName)
			// fmt.Print("Enter snapshot name: ")
			// fmt.Scan(&snapshotName)
			// fmt.Println("Getting info about snapshot:", snapshotName, "for VM:", vmName)
			utils.ShowSnapshotsInfo(vmName)
			return
		case 7:
			var vmName string
			fmt.Print("Enter VM name to get snapshot size: ")
			fmt.Scan(&vmName)
			fmt.Println("Getting snapshot size for:", vmName)
			utils.GetSnapshotDiskSize(vmName)
			return
		case 8:
			var vmName, snapshotName string
			fmt.Print("Enter VM name to revert snapshot: ")
			fmt.Scan(&vmName)
			fmt.Print("Enter snapshot name to revert: ")
			fmt.Scan(&snapshotName)
			fmt.Println("Reverting snapshot:", snapshotName, "for VM:", vmName)
			utils.RevertSnapshot(vmName, snapshotName)
			return
		case 9:
			var vmName, snapshotName string
			fmt.Print("Enter VM name to delete snapshot: ")
			fmt.Scan(&vmName)
			fmt.Print("Enter snapshot name to delete: ")
			fmt.Scan(&snapshotName)
			fmt.Println("Deleting snapshot:", snapshotName, "for VM:", vmName)
			utils.DeleteSnapshot(vmName, snapshotName)
			return
		default:
			fmt.Println("Invalid choice, please try again.")
			return
		}
	}
}
