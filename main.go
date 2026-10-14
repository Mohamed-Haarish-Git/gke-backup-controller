package main

import "fmt"




func main() {
    const MAX_RETRIES int = 5 // final in Java = const in Go
    var bucketName string = "my-gke-backups"
    var failureCount int = 3

    fmt.Println(bucketName, failureCount, MAX_RETRIES)


	var retries [3]int = [3]int{1,5,15}
	failedBackups := []string{}
	failedBackups = append(failedBackups,"backup-123")
	failedBackups = append(failedBackups, "backup-456")
	fmt.Println("Retries array : ", retries)
	fmt.Println("Failed backups data :",failedBackups)
	fmt.Println("First failed :",failedBackups[0])



	  backupName := "backup-123"
    reason := "storage full"

    // This is your structured logging
    log := fmt.Sprintf("Backup %s failed: %s", backupName, reason)
    fmt.Println(log)
}

	
	

