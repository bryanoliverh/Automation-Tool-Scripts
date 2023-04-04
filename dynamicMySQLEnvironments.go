package main

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "os/exec"
    "regexp"
    "strconv"
    "strings"
)

func main() {
    // Define the list of servers to SSH to
    servers := []string{"server1.example.com", "server2.example.com", "server3.example.com"}

    totalMem := getTotalMem()
    for _, server := range servers {
        err := sshExec(server, "/etc/my.cnf", 0.8, totalMem)
        if err != nil {
            fmt.Println(err)
            continue
        }
        fmt.Printf("innodb_buffer_pool_size updated on %s\n", server)

        err = changeReadOnly(server, "/etc/my.cnf")
        if err != nil {
            fmt.Println(err)
            continue
        }
        fmt.Printf("read_only updated on %s\n", server)
    }
}

func sshExec(server, filePath string, factor float64, totalMem int) error {
    // Define the SSH command
    sshCmd := fmt.Sprintf("ssh -i /path/to/rsa/key -F /path/to/config/file %s", server)

    // Define the commands to execute
    freeCmd := "free -g"
    sedCmd := fmt.Sprintf(`sed -i '/^innodb_buffer_pool_size/d' %s && sed -i '$ a\innodb_buffer_pool_size = %dG' %s`, filePath, int(factor*float64(totalMem)), filePath)

    // Create the SSH command and connect the commands
    sshArgs := strings.Split(sshCmd, " ")
    ssh := exec.Command(sshArgs[0], sshArgs[1:]...)
    ssh.Stdin = bytes.NewBufferString(fmt.Sprintf("%s && %s\n", freeCmd, sedCmd))

    // Run the command and get the output
    output, err := ssh.CombinedOutput()
    if err != nil {
        return fmt.Errorf("failed to execute commands on %s: %v - %s", server, err, output)
    }

    return nil
}

func getTotalMem() int {
    cmd := exec.Command("free", "-g")
    out, err := cmd.Output()
    if err != nil {
        fmt.Println(err)
        return 0
    }

    // Extract the total memory using a regular expression
    re := regexp.MustCompile(`Mem:\s+(\d+)\s+.*`)
    match := re.FindStringSubmatch(string(out))
    if len(match) != 2 {
        fmt.Println("Error: unable to extract total memory")
        return 0
    }

    // Convert the total memory from string to int
    totalMem, err := strconv.Atoi(match[1])
    if err != nil {
        fmt.Println("Error: unable to convert total memory to int")
        return 0
    }

    return totalMem
}

func changeReadOnly(server, filePath string) error {
    // Define the SSH command
    sshCmd := fmt.Sprintf("ssh -i /path/to/rsa/key -F /path/to/config/file %s", server)

    // Define the command to execute based on hostname
    hostnameCmd := "hostname"
    sedCmd := ""

    cmd := exec.Command("ssh", "-i", "/path/to/rsa/key", "-F", "/path/to/config/file", server, hostnameCmd)
    out, err := cmd.Output()
    if err != nil {
        return fmt.Errorf("failed to execute command on %s: %v", server, err)
    }

    hostname := strings.TrimSpace(string(out))
    if strings.Contains(hostname, "slave") {
        sedCmd = "sed -i 's/^read-only = 1/read-only = 0/' " + filePath
    } else if strings.Contains(hostname, "master") {
        sedCmd = "sed -i 's/^read-only = 0/read-only = 1/' " + filePath
    } else {
        return fmt.Errorf("unable to determine hostname type for %s", server)
    }

    // Create the SSH command and connect the commands
    commands := []string{sedCmd, "echo read_only updated"}
    sshArgs := strings.Split(sshCmd, " ")
    ssh := exec.Command(sshArgs[0], sshArgs[1:]...)
    ssh.Stdin = bytes.NewBufferString(strings.Join(commands, " && ") + "\n")

    // Run the command and get the output
    output, err := ssh.CombinedOutput()
    if err != nil {
        return fmt.Errorf("failed to execute command on %s: %v - %s", server, err, output)
    }

    return nil
}
