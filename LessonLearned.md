# Lesson Learned
This document contains the things that I actually pickup when following through this `Learn Go with tests` curriculum

## The Concept of Test
- It is important for the test files to have the naming like `some_name_test.go`. The `_test.go` suffix is the key. Because it will be automatically picked up by the testing sytem to be executed
- We can think of your typical testing structure when using any testing tool in JS/TS project. You have **test suites** and **test cases**
    - The **test suites** are the functions in the testing file. Example: `TestingHello`, `TestingAddition`, `TestingSubstraction`
    - The **test cases** are defined by invoking the `t.Run()` function from the `testing` package. With `t.Run` you can supply the test case title and the callback function that actually tests the thing you need to test. *Think of it like test cases with Jest or Vitest*
- It is important to invoke `t.Helper()` in any test helper functions, so when an error happens in the test helper, the test runner would still print the filename and line number of the calling test function, instead of the line number in the helper function itself.
