package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// Read the database names from the txt file
	databases := readDatabaseNames("database_list.txt")

	// Exclude certain databases
	excludedDatabases := []string{"information_schema", "mysql", "performance_schema", "sys", "dba_meta", "mysql_screening_db"}
	filteredDatabases := excludeDatabases(databases, excludedDatabases)

	// Generate the mysqldump command
	mysqldumpCommand := generateMysqldumpCommand(filteredDatabases)

	fmt.Println(mysqldumpCommand)
}

// Read the database names from the txt file
func readDatabaseNames(filename string) []string {
	var databases []string

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		database := scanner.Text()
		databases = append(databases, database)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	return databases
}

// Exclude certain databases from the list
func excludeDatabases(databases []string, excludedDatabases []string) []string {
	var filteredDatabases []string

	for _, database := range databases {
		// Check if the database is in the excluded list
		if !contains(excludedDatabases, database) {
			filteredDatabases = append(filteredDatabases, database)
		}
	}

	return filteredDatabases
}

// Check if a string exists in a slice
func contains(slice []string, str string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}

// Generate the mysqldump command for the included databases with dates and host IP
func generateMysqldumpCommand(databases []string) string {
	// Get the current date and time
	currentTime := time.Now().Format("20060102_150405")

	// Get the host IP
	hostname, _ := os.Hostname()

	// Generate the backup file name with dates and host IP
	backupFileName := fmt.Sprintf("backup_%s_%s.sql", currentTime, hostname)

	cmd := fmt.Sprintf("mysqldump --compact --single-transaction --skip-add-drop-table --set-gtid-purged=off --user=ld-bryan --password --port=6606 --databases %s > %s", strings.Join(databases, " "), backupFileName)
	return cmd
}
