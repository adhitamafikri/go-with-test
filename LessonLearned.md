# Lesson Learned

Site: [https://quii.gitbook.io/learn-go-with-tests](https://quii.gitbook.io/learn-go-with-tests)

---

This document contains the things that I actually pickup when following through this `Learn Go with tests` curriculum

## The Concepts of Test

- It is important for the test files to have the naming like `some_name_test.go`. The `_test.go` suffix is the key. Because it will be automatically picked up by the testing sytem to be executed
- We can think of your typical testing structure when using any testing tool in JS/TS project. You have **test suites** and **test cases**
  - The **test suites** are the functions in the testing file. Example: `TestingHello`, `TestingAddition`, `TestingSubstraction`
  - The **test cases** are defined by invoking the `t.Run()` function from the `testing` package. With `t.Run` you can supply the test case title and the callback function that actually tests the thing you need to test. _Think of it like test cases with Jest or Vitest_
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
- It is important not to rely on external services such as online API placeholder, dev/staging/prod API services when doing tests. Avoiding these things would make our test faster, more predictable, and allows us to test the possible edge cases
- We can use `net/http/httptest` to test our functions that have an actual call to the external services

## Useful Testing Command Flags

- `-v` -> produce verbose output. Prefer to always use this so we can easily spot which cases are succeeded or failed
- `-bench` -> run benchmark on a test file to figure out the time it takes for any operations to be done per nanosecond
- `-benchmem` -> run benchmark on a test file to figure out the size being allocated by Go to the memory when executing any operations
- `-cover` -> display percentage of lines of code being covered in the test
- `-race` -> a race detector, that can be used to detect race conditions that might happen in our program

## Benchmarking

Read: [Benchmarking](https://pkg.go.dev/testing#hdr-Benchmarks)

- We can do _benchmarking_ alongside testing the actual behavior of the function
- The benchmark function naming convention is `Benchmark` + `FunctionNameInPackage`
- The benchmark could be run by executing `go test -bench=.` command in your terminal. It could also be used in conjunction with the `-v` flag so it will produce verbose test result plus the benchmark result. Example `go run test -v -bench=.`

## Dependency Injection

> Learned how to implement DI without using any framework
> Directory: 008_dependency_injection

## Printf and Fprintf quirks

- `Printf` will send the output to the `stdout`
- `Fprintf` allows us to print the output into any writer that implements `io.Writer` interface like:
  - Writing to files
  - Writing to buffers
  - Network responses
  - Any custom destinations

## Concurrency

> The primary goal of implementing concurrency is to have several process running concurrently.
> Making the process execution non-blocking and faster

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
- `chan struct{}` is the smallest data type available from a memory perspective, virtually no allocation. Perfect for a use case when we need to create a channel that does not send anything
- Always use `make()` when creating channels, using `var channelName chan struct{}` or similar will initialize the variable with the "zero" value of the type, which is `nil`. When we try to send `nil` value, it will block forever because basically we cannot send `nil` to channels

## Select

> `select` in Go is used to synchronize multiple channel operations

- This allows us to wait on multiple `channel`. The first one to send a value "wins"
- Sometimes it might be needed to have a `case` with the `time.After`, typically the last `case` statement inside the `select`. Preventing the system from blocking forever

More on [https://gobyexample.com/select](https://gobyexample.com/select)

## Defer

> Delay the execution of a function at the end of the containing function. Example `defer someFunction()`

- Typically used for executing anything that can be delayed until later or after the other important statements of a function are executed. Example usages:
  - Closing connection to database
  - Closing server
  - Closing files
  - Test Teardown, Cleaning up test mocks


## Reflection

> This can be used to inspect the information/metadata of a variable. It can be used purely for inspection or manipulation

- Only use `reflection` when we really need to. Typically used for "polymorphic functions", which are the functions having parameters that can be of any type/multiple types (`interface{}` type of parameters)
- The most important function from the `reflect` package are `reflect.ValueOf` and `reflect.TypeOf`
- Need to be careful when using `val.Field()`. A value **might not have fields at all** or we are trying to access non-existent fields (example: the value has 3 fields, we are trying to access the 4th field), this will cause a panic
- We can also use `field.Kind()` to get the kind of type data of the given value
- We can use `reflect.<TypeData>` to compare the value produced by `field.Kind()`, useful to check the exact data type of a value
- We can use `field.Interface()` to get the interface value. Useful when we want to get the value of nested interfaces for inspection or manipulation purposes
- If the function receives **pointer** as parameter value, we need to get the actual value element by using `val.Elem()`. Without doing this the program will panic
- We cannot access `NumField` on a slice, they will raise error. The safest way to work with slice is to check if the value kind is of `reflect.Slice`, iterate each value, and extract the value interface by doing `val.Index(i).interface()`
- Arrays can be handled the same way as slices
- We have to be careful with `maps`, because Go can't guarantee order of the items in the map
- Handling `channel` is a bit tricky. You need to use `val.Recv()` which receives and returns value from the channel
- You might need to do recursive traversal if complex reflection operations are needed

More on:
- [https://dasarpemrogramangolang.novalagung.com/A-reflect.html](https://dasarpemrogramangolang.novalagung.com/A-reflect.html)
- [https://pkg.go.dev/reflect](https://pkg.go.dev/reflect)
- [https://pkg.go.dev/reflect#Value.Recv](https://pkg.go.dev/reflect#Value.Recv)
