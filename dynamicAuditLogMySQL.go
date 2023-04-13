package main

import (
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/ssh"
)

func main() {
	// Parse command-line arguments
	mysqlHost := flag.String("h", "localhost", "MySQL host")
	mysqlPort := flag.String("P", "3306", "MySQL port")
	mysqlUser := flag.String("u", "root", "MySQL username")
	mysqlPass := flag.String("p", "", "MySQL password")
	auditLogPolicy := flag.String("t", "ALL", "Audit log policy to set")
	flag.Parse()

	// SSH configuration
	sshConfig := &ssh.ClientConfig{
		User: "your-ssh-username",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(getPrivateKey("path/to/id_rsa")),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect to SSH server
	sshClient, err := ssh.Dial("tcp", "your-db-server:22", sshConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer sshClient.Close()

	// Connect to MySQL server
	mysqlConfig := mysql.NewConfig()
	mysqlConfig.User = *mysqlUser
	mysqlConfig.Passwd = *mysqlPass
	mysqlConfig.Addr = fmt.Sprintf("%s:%s", *mysqlHost, *mysqlPort)
	mysqlConfig.Net = "tcp"
	mysqlConfig.AllowNativePasswords = true
	mysqlConfig.ParseTime = true

	sshConn, err := sshClient.Dial("tcp", mysqlConfig.Addr)
	if err != nil {
		log.Fatal(err)
	}
	defer sshConn.Close()

	db, err := sql.Open("mysql", fmt.Sprintf("%s@%s/", mysqlConfig.FormatDSN(), mysqlConfig.Addr))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Query the user list
	rows, err := db.Query("SELECT CONCAT(\"'\", USER, \"'@'\", HOST, \"',\") FROM mysql.user WHERE USER NOT IN ('root', 'mysql.sys')")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Build the final command to set audit_log_include_accounts for all users and hosts
	var userHosts string
	for rows.Next() {
		var userHost string
		err := rows.Scan(&userHost)
		if err != nil {
			log.Fatal(err)
		}
		userHosts += userHost
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	// Remove the trailing comma from the userHosts string
	userHosts = strings.TrimSuffix(userHosts, ",")

	// Set audit_log_include_accounts for all users and hosts
	_, err = db.Exec(fmt.Sprintf("SET GLOBAL audit_log_include_accounts=%s", userHosts))
	if err != nil {
		log.Fatal(err)
	}

	// Set audit_log_policy if provided
	if *auditLogPolicy != "" {
		_, err = db.Exec(fmt.Sprintf("SET GLOBAL audit_log_policy=%s", *auditLogPolicy))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Set audit_log_policy=%s\n", *auditLogPolicy)
	}
}

func getPrivateKey(path string) (ssh.Signer, error) {
    privateKey, err := ioutil.ReadFile(path)
    if err != nil {
        return nil, err
    }
    signer, err := ssh.ParsePrivateKey(privateKey)
    if err != nil {
        return nil, err
    }
    return signer, nil
}