# File Reading Examples in Go

## Importance of Reader and Writer in go
Go makes use of `io.Reader` and `io.Writer` interface in golang. For this demonstration, our examples focuses on Reader interfaces (as input argument) which simply reads the data from the file. With that data, we use various methods and some logic to process and return the data. Moving ahead, we will also be doing some tests and benchmarks to check our function performance.

#### <ins>Examples</ins>
1. File Reading using `bufio` package
2. File Reading using `io` package 
    1. Using `io.Copy` 
    2. Using `io.Read`
    3. Using `io.ReadAll`
    4. Using `io.TeeReader`


#### <ins>File Reading using `bufio` package</ins>
The bufio package in Go helps make reading and writing data more efficient by using buffered I/O. Instead of handling data byte by byte, it processes data in chunks, which reduces the number of I/O operations and improves performance. You can use `bufio.NewReader()` to read a file line by line, or the `bufio.Scanner` to read data based on lines or a custom delimiter, making it useful for working with text files or user input. The package also provides handy methods like `ReadString()` to read data until a specific character, such as a newline, simplifying the process of managing input and output.

#### <ins>File Reading using internal `io` package</ins>
The `io` package in Go provides essential interfaces for handling input and output operations, it is a core part of reading and writing data. The most common interfaces are **Reader** and **Writer**. The Reader interface uses the `Read()` method to read data into a slice of bytes, while the Writer interface uses the `Write()` method to write data from a slice of bytes. Additionally, the package includes useful functions like `Copy()`, which copies data from a Reader to a Writer, and `ReadAll()`, which reads all the data from a Reader at once. This makes the io package both flexible and powerful for handling _I/O tasks_ in Go.

#### <ins>Using `io.Read` function</ins>
The `io.Read` function in Go is used to read data from a source into a slice of bytes. It’s part of the _Reader_ interface, which defines how data is read in Go. The data in read in chunks(buffer) and write the buffered data to bytes. This method of reading and processing data is effective for handling large amounts of data or when we want to efficiently handle **I/O operations**.

#### <ins>Using `io.ReadAll` function</ins>
The `io.ReadAll` function is used to read the entire content of the data into memory. This method is effective when you know that the data size is manageable in memory. For very large inputs, it may not be suitable due to memory constraints.

#### <ins>Using `io.TeeReader` function</ins>
`io.TeeReader` in Go allows you to read data from a source while simultaneously write the same data to a log or buffer. This is especially useful when you need to process or analyze the data and also keep record of it. By splitting the data flow, `io.TeeReader` lets you handle the data in real-time while also maintaining a copy for logging, debugging, or duplication purposes. This dual functionality makes it an effective solution for tracking and managing data, ensuring you have a comprehensive record for various applications.

## Testing and Benchmarks
Testing and benchmarking are crucial in Go as they ensure that your code is **efficient**, **correct** and **reliable**. Go has a strong built-in testing framework that makes writing and running tests. Go also provides built-in support for benchmarking, allowing you to measure the performance of your code. This is particularly important in performance-sensitive applications such as cloud-native services, microservices, and systems programming.

#### <ins>Testing</ins>
Use the following command for testing various package. 
```go
git clone https://github.com/SmartByteLabs/filereadingexamples-go.git
cd filereadingexamples ; cd bufio_read
go test 
```
For verbose output, you can use `-v` flag.
```go
go test -v
```

#### <ins>Benchmarks</ins>
For benchmarks, use the `-bench` flag
```go
go test -bench=.
```
This will perform the tests but will ignore the test case and prints only the benchmarks results.

#### For running only benchmarks tests
Using `run=^$` flag will ignore the tests and only perform benchmarks test. You can also use `-v` to get verbose output.

```go
cd bufio_read
go test -bench=. -benchmem -run=^$ .
```

#### Running All benchmarks
You can run all of the benchmarks at onnce using `run_benchmarks.sh` (make it executable). This will perform benchmarks of all the function line by line.
```bash
cd scripts
chmod +x run_benchmarks.sh
./run_benchmarks.sh
```
This command is same as `go test -bench=. -benchmem -run=^$ .`,the same process is automated with the use of script.

## Benchmarks Result
After running benchmark tests on machine, we received the following output. The machine specs includes:

#### Machine Specs
```
OS : Ubuntu 22.04.4 LTS WSL2 on Windows 10
CPU : Intel(R) Core(TM) i5-7200U CPU @ 2.50GHz (4 Core)
Arch : amd64
Ram : 16 GB
```


## Ouput

#### <ins>Program Execution Time</ins>
| **Filename** | **File Size** | **bufio_read** | **bytessbuffer_read** | **io_read** | **io_readall** | **tee_read** |
| :------------: | :-------------: | :--------------: | :---------------------: | :-----------: | :--------------: | :------------: |
| words_1.txt  | 4.7 MB        | 0.039s         | 0.004s                | 0.003s      | 0.010s         | 0.003s       |
| words_2.txt  | 101 MB        | 0.725s         | 0.049s                | 0.052s      | 0.105s         | 0.005s       |
| words_3.txt  | 1.1 GB        | 7.85s          | 17.68s                | 3.42s       | 2.20s          | 18.88s       |

#### <ins>Memory Consumption</ins>
| **Filename** | **FileSize** | **bufio_read** | **bytesbuffer_read** | **io_read** | **io_readall** | **tee_read** |
| :------------: | :------------: | :-----------------------: | :-----------------------------: | :--------------------: | :-----------------------: | :---------------------: |
| words_1.txt  | 4.7 MB       | 33.04 MB                | 21.64 MB                      | 21.64 MB             | 31.37 MB                | 21.64 MB              |
| words_2.txt  | 101 MB       | 756.18 MB               | 373.30 MB                     | 373.30 MB            | 720 Mb                  | 53.60 MB              |
| words_3.txt  | 1.1 GB       | 7179.86 MB              | 5368.71 MB                    | 5368.71 MB           | 6810 MB                 | 5368.71 MB            |

#### <ins>Memory Allocation Count</ins>
| **File Name** | **File Size** | **bufio_read** | **bytesbuffer_read** | **io_read** | **io_readall** | **tee_read** |
| :-------------: | :-------------: | :-------------------------: | :-------------------------------: | :----------------------: | :-------------------------: | :-----------------------: |
| words_1.txt   | 4.7 MB        | 46.6k                     | 17                              | 16                     | 35 <br>                   | 19                      |
| words_2.txt   | 101 MB        | 10.06 Million             | 21 Million                      | 20                     | 50                        | 23                      |
| words_3.txt   | 1.1 GB        | 103.2 Million             | 26 Million                      | 25                     | 59                        | 28                      |

### Note
The file operations done on our local environment are larger than _100 MB_. You can generate large txt files using `words_generater.go` present in `scripts` directory.

```go
cd scripts
go run words_generator.go 200 MB
```





