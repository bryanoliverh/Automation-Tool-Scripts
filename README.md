# AutomationTool-Golang

1.  [findAndReplace.go](https://github.com/bryanoliverh/AutomationTool-Golang/blob/main/findAndReplace.go)
  -  Time Complexity: O(nm), where n is the length of the input file and m is the length of the search string. Therefore, the overall time complexity of the program is O(nm).
   
     Space Complexity: O(n) 

  -  This is a command-line tool written in Go that finds and replaces text in a file. It provides various options to customize the find and replace operation.

  - Usage:
    ```  
    findreplace [options] <file>
    Options:

    -T, --t TEXT : Text to find.
    -R, --r TEXT : Text to replace with.
    -file FILE : File to modify.
    -f FILE : Alternative file to modify (shorthand).
    -C, --case CASE : Specify case sensitivity (ci for case insensitive, cs for case sensitive).
    -O, --out FILE : Output file.
    ```
    Examples
    ```
    findreplace -T "foo" -R "bar" input.txt
    findreplace --t "foo" --r "bar" --out output.txt --case cs --file input.txt
    ```
   - Implementation Details: The tool uses the flag package to define and parse command-line arguments. The ioutil package is used to read and write files.
  
  - The tool checks if both -T and -R options are provided. It also checks if either -file, -f, or input file path are provided. The tool uses the strings.ReplaceAll function to find and replace text. The case sensitivity option can be specified with -C or --case. The default case sensitivity is case-insensitive. The output is written to the file specified with -O or --out.

  - If an error occurs during the operation, the tool prints an error message and exits with a non-zero status code. Otherwise, it prints a success message and the path of the output file.

2.  [binarySearch.go](https://github.com/bryanoliverh/AutomationTool-Golang/blob/main/binarySearch.go)
  -  Time Complexity: O(log n), where where n is the number of elements in the input array.
   
     Space Complexity: O(1) 

  -  This is a command-line tool written in Go that finds the index of the given list by binary search.

3.  [dynamicMySQLEnvironments.go](https://github.com/bryanoliverh/AutomationTool-Golang/blob/main/dynamicMySQLEnvironments.go)


  -  This is a tool to adjust MySQL Environments dynamically using multiple parameters such as the read only and the innodb_buffer_pool_size, as well as to print the DB server host information into a table.
  
  -  The code can be run with
    ``` 
        go run dynamicMySQLEnvironments.go -h server1.example.com,server2.example.com,server3.example.com,10.12.12.1
    ```
4.  [dynamicAuditLogMySQL.go](https://github.com/bryanoliverh/AutomationTool-Golang/blob/main/dynamicAuditLogMySQL.go)


  -  This tool will change MySQL global variables to help you manage your MySQL audit logs.
  -  This script will SSH to DB Servers and change the audit_log_include_accounts to include all of the accounts with the MySQL 'CREATE' privilege.
  -  Usage:
        ```
        To use this tool, you need to provide the following parameters:

        -h: The MySQL server host.
        -P: The MySQL server port.
        -u: The MySQL server username.
        -p: The MySQL server password.
        -t: The audit log policy to set (optional, default: ALL).
        ```
   - Here is an example command to set the audit log policy to NONE for a MySQL server running on localhost:
    ```
      go run dynamicAuditLogMySQL -h {yourhost} -u {youruser} -p {yourpassword} -t LOGINS
    ```
  

5.  [TerraformGolang](https://github.com/bryanoliverh/AutomationTool-Golang/tree/main/TerraformGolang)


  -  This is a practice project to explore Terraform using Golang for the GCP Services such as the Compute Engine, GCP Buckets, and Hypervisor.
  -  Terraform basics: https://github.com/shuaibiyy/awesome-terraform.
  
6.  [mysqldumpCommandCreator](https://github.com/bryanoliverh/AutomationTool-Golang/tree/main/mysqldumpCommandCreator)

  -  mysqldump command creator based on the database list inside your db server.
  -  this tool will automatically generate mysqldump commands based on the .txt file that must be filled with all of the databases in the server.

7.  [get_serverid.sh](https://github.com/bryanoliverh/Automation-Tool-Scripts/blob/main/get_serverid.sh)

  -  get the server IP and turn it into a server id for MySQL Configuration
    
    
