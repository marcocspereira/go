# go

## Intro

Repository that holds exercises related to the Go programming language.

Some sources of information:

* [Configure Visual Studio Code for Go development](https://learn.microsoft.com/en-us/azure/developer/go/configure-visual-studio-code)
* [Go by Example](https://gobyexample.com/)
* [Learng Go with Tests](https://quii.gitbook.io/learn-go-with-tests/)
* [Questions](questions.md)

## Basics

**Type of language:**

* Python, Ruby or JavaScript are all examples of dynamically typed languages, one which the developers do not care what values are assigned to variables
* With **Go**, whenever a variable is defined, a type is also assigned to it, like Java or C.

**Types of variables**

| Type | Example |
| --- | --- |
|boolean|`true` or `false`|
|string |`"Hello World"`|
|int|`0` `-100` `999`|
|float|`0.0` `-100.5` `3.141592`|

## Writing tests
Writing a test is just like writing a function, with a few rules:

* It needs to be in a file with a name like `xxx_test.go`
* The test function must start with the word `Test`
* The test function takes one argument only `t *testing.T`
* In order to use the `*testing.T` type, you need to `import "testing"`, like we did with `fmt` in the other file

`t` of type `*testing.T` is your "hook" into the testing framework so you can do things like `t.Fail()` when you want to fail.

## Go CLI

| Command | Description |
| --- | --- |
| `go run` | Compile and execute one or two or more files |
| `go build` | Compile a bunch of go source code files and does not actually execute and creating an executable file |
| `go fmt` | Format all the code in each file in the current directory |
| `go install` | Compile and install a package |
| `go get` | Download the raw source code of someone else's package |
| `go test` | Run any tests associated with the current project |

## Go Packages

Package == Workspace == Project

A package is:
* a collection of common source code files that are compiled together;
* a directory that contains one or more Go source code files.

If we are working on a project that has multiple files:
* we can group them into a single package;
* the only requirement for every file inside of a package is that the very first line of the file must be a package declaration.

### Types of packages

* **Executable package:** generates a file that we can run directly. Is one when compiled, spits out an executable file that we can run. Useful for when we want to create a program that we can run and to accomplish tasks.
* **Reusable package:** code used as 'helpers'. Good place to put reusable logic. They're not to execute directly, but to be used as helpers inside of other code.

**How do you know if a package is executable or reusable?**

**executale:** package main -> go build -> executable

* the package named **main** is a special package;
* it tells the Go compiler that the package should compile as an executable program instead of a shared library;
* must have a function called **main**.

**reusable package:** package blablabla -> go build -> nothing

* any other package name will be compiled as a shared library
* defines a package that can be imported by other code, used as dependency.

## Meaning of `import "fmt"` statement

The `import` statement is how we include code from other packages to use with our program.

* `import "fmt"` specifically means "*give my package main access to all the code and all functionallity inside of the package called fmt*";
* `fmt` is the name of a standard library package that is included with the go programming language by default;
* is a shortened form of the word format;
* it is used to print out things to the terminal;
* [official documentation](https://golang.org/pkg/fmt/) for the fmt package.

## Workspaces

From Go 1.18 onwards there is native support for multi-module workspaces. This is done by having a go.work file present in your parent directory.

```
$ tree /my/parent/dir
/my/parent/dir
├── project-one
│   ├── go.mod
│   ├── project-one.go
│   └── project-one_test.go
└── project-two
    ├── go.mod
    ├── project-two.go
    └── project-two_test.go
```

Create and populate the file by executing go work:

```
cd /my/parent/dir
go work init
go work use project-one
go work use project-two
```

This will add a go.work file in your parent directory that contains a list of directories you marked for usage:

```
go 1.18

use (
    ./project-one
    ./project-two
)
```