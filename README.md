# go

Repository that holds exercises related to the Go programming language.

* [Configure Visual Studio Code for Go development](https://learn.microsoft.com/en-us/azure/developer/go/configure-visual-studio-code)


## Writing tests
Writing a test is just like writing a function, with a few rules:

* It needs to be in a file with a name like `xxx_test.go`
* The test function must start with the word `Test`
* The test function takes one argument only `t *testing.T`
* In order to use the `*testing.T` type, you need to `import "testing"`, like we did with `fmt` in the other file

`t` of type `*testing.T` is your "hook" into the testing framework so you can do things like `t.Fail()` when you want to fail.