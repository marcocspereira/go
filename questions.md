# Questions

## Intro

**What is the purpose of a package in Go?**

To group together code with similar purpose.

**What is the special name we use for a package to tell Go that we want it to be turned into a file that can be executed?**

main

**The one requirement of packages named "main" is that we...**

Define a function named `main`, which is ran automatically when the program runs.

**Why do we use "import" statements?**

To give our package access to code written in another package.

## Variable Assignement

**Is the following a valid way of initializing and assigning a value to a variable?** `var bookTitle string = "Harry Potter"`

Yes

**Is the following a valid way of initializing an assigning a value to a variable?** `fruitCount := 5`

Yes

**After running the following code, Go will assume that the variable `quizQuestionCount` is of what type?** `quizQuestionCount := 10`

Integer

**Will the following code compile?  Why or why not?**

```go
paperColor := "Green"
paperColor := "Blue"
````

No, because a variable can only be initialized one time. In this case, the `:=` operator is being used to initialize `paperColor` twho times

**Are the two lines following ways of initializing the variable `pi` equivalent?**

```go
pi := 3.14 
var pi float64 = 3.14
````

Yes

**Is the following code valid?**

    
```go
package main
 
import "fmt"
 
deckSize := 20
 
func main() {
  fmt.Println(deckSize)
}
```

No, because there is a `syntax error: non-declaration statement outside function body`. This is because the `:=` operator can only be used to initialize a variable inside a function.

**We might be able to initialize a variable and then later assign a variable to it. Is the following code valid?**

```go
package main
 
import "fmt"
 
func main() {
  var deckSize int
  deckSize = 52
  fmt.Println(deckSize)
}
```

Yes. The variable `deckSize` is initialized with the `var` keyword, and then assigned a value with the `=` operator.

**Is the following code valid?**

```go
package main
 
import "fmt"
 
var deckSize int
 
func main() {
  deckSize = 50
  fmt.Println(deckSize)
}
```

Yes. We can initialize a variable outside of a function, we just can't assign a value to it.

**Is the following code valid?  Why or why not?**

```go
package main
 
import "fmt"
 
func main() {
  deckSize = 52
  fmt.Println(deckSize)
}
```

No, because `deckSize`is assigned a value before it is initialized. We can't assign a value to a variable that hasn't been initialized. Variables must first be initialized with the `:=` operator or the `var variableName type` syntax.