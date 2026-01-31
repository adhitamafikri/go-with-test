# Lesson Learned
This document contains the things that I actually pickup when following through this `Learn Go with tests` curriculum

## The Concepts of Test
- It is important for the test files to have the naming like `some_name_test.go`. The `_test.go` suffix is the key. Because it will be automatically picked up by the testing sytem to be executed
- We can think of your typical testing structure when using any testing tool in JS/TS project. You have **test suites** and **test cases**
    - The **test suites** are the functions in the testing file. Example: `TestingHello`, `TestingAddition`, `TestingSubstraction`
    - The **test cases** are defined by invoking the `t.Run()` function from the `testing` package. With `t.Run` you can supply the test case title and the callback function that actually tests the thing you need to test. *Think of it like test cases with Jest or Vitest*
- It is important to invoke `t.Helper()` in any test helper functions, so when an error happens in the test helper, the test runner would still print the filename and line number of the calling test function, instead of the line number in the helper function itself.
- We can add `Testable Examples` into our test files
    - Keep in mind that the naming convention a `Testable Examples` is "Example" + "method name from the package". Example:
    ```go:some_package.go
    package some_package
    import "fmt"

    func ConcatenateName(n1, n2 string) string {
        // do concatenation
    }
    ```

    ```go:some_package_test.go
    package some_package
    import "testing"

    func TestSomePackage(t *testing.T) {
        t.Run("case", func (t *testing T) { ... })
    }

    // testable example
    func ExampleConcatenateName() {
        result := ConcatenateName("Santo", "Hara")
        fmt.Println(result)
        // Output: "Santo Hara"
    }
    ```
- We can have multiple `Testable Examples` in our test file for any functions in the package by doing this:
    - Having functions with the name like `ExampleFunctionName_someIdentifier`.
        - The `_someIdentifier` here could be anything like `_bigNumber`, `_maskingEmail`, `_negativeAndPositiveNumOps`, .etc
        - The `_someIdentifier` could not be started with non alphabet chars, it will raise an error. Example: `ExampleFunctionName_1`, `ExampleFunctionName_#something`

## Useful Testing Command Flags
- `-v` -> produce verbose output. Prefer to always use this so we can easily spot which cases are succeeded or failed
- `-bench` -> run benchmark on a test file to figure out the time it takes for any operations to be done per nanosecond
- `-benchmem` -> run benchmark on a test file to figure out the size being allocated by Go to the memory when executing any operations
- `-cover` -> display percentage of lines of code being covered in the test
- `-race` -> a race detector, that can be used to detect race conditions that might happen in our program

## Benchmarking
Read: [Benchmarking](https://pkg.go.dev/testing#hdr-Benchmarks)

- We can do *benchmarking* alongside testing the actual behavior of the function
- The benchmark function naming convention is `Benchmark` + `FunctionNameInPackage`
- The benchmark could be run by executing `go test -bench=.` command in your terminal. It could also be used in conjunction with the `-v` flag so it will produce verbose test result plus the benchmark result. Example `go run test -v -bench=.`

## Dependency Injection
> Learned how to implement DI without using any framework
Directory: 008_dependency_injection

## Printf and Fprintf quirks
- `Printf` will send the output to the `stdout`
- `Fprintf` allows us to print the output into any writer that implements `io.Writer` interface like:
    - Writing to files
    - Writing to buffers
    - Network responses
    - Any custom destinations

## Concurrency
> The primary goal of implementing concurrency is to have several process running concurrently.
Making the process execution non-blocking and faster

- You have to leverage `channels` and `goroutines`
- `goroutine` is a separate process in Go that allows us to execute without blocking
    - You can tell Go to start a new `goroutine` by turning a function call into `go` statement in front of a function call like this. (Basically all you need to do is to wrap the statements that will be run concurrently with an immediately invoked `go func() {}` anonymous function)
    ```go
    func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
        results := make(map[string]bool)

        for _, url := range urls {
            // run goroutines here
            go func() {
                results[url] = wc(url)
            }()
        }

        return results
    }
    ```
    - It is hard to predict what's going to happen on the concurrently running processes, so we have carefully handle them
- `channels` are a Go data structure that could both receive and send values. This also allows communication between different processes
    - `channels` are used to coordinate our `goroutines` to prevent race conditions
    - Send statement -> `channelName <- doSomething()`
    - Receive statement -> `result := <-channelName`
