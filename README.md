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
  
  -  The code can be run with"
    ```
     go run dynamicMySQLEnvironments.go -h server1.example.com,server2.example.com,server3.example.com,10.12.12.1
    ```
