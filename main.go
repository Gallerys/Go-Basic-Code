package main

import "fmt"

func main() {
    // Declaring a variable
    myVariable := 10

    // Declaring a constant
    const myConstant = 5

    // Printing to the console
    fmt.Println("Hello, Go!")

    // Conditional statement
    if myVariable > myConstant {
        fmt.Println("myVariable is greater than myConstant")
    } else if myVariable == myConstant {
        fmt.Println("myVariable is equal to myConstant")
    } else {
        fmt.Println("myVariable is smaller than myConstant")
    }

    // Loop
    for i := 1; i <= 5; i++ {
        fmt.Printf("Iteration %d\n", i)
    }

    // Slice (dynamic array)
    mySlice := []int{1, 2, 3, 4, 5}

    // Loop through slice
    for _, number := range mySlice {
        fmt.Println(number)
    }

    // Function
    greeting := greet("Alice")
    fmt.Println(greeting)
}

// Function definition
func greet(name string) string {
    return fmt.Sprintf("Hello, %s!", name)
}
