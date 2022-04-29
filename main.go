package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2

	fmt.Print("Write ", i, " as ")

	switch i {

	case 1:

		fmt.Println("one")

	case 2:

		fmt.Println("two")

	case 3:

		fmt.Println("three")

	}

	// for another program :

	// we can use commas to separate multiple expressions in the same case statement.
	// We use the optional 'default' case in this example as well.
	// a switch statement can have an optional default case, which must appear at the end of the switch.
	// The default case can be used for performing a task when none of the cases is true. No break is needed in the default case.

	switch time.Now().Weekday() {

	case time.Saturday, time.Sunday:

		fmt.Println("It's the weekend")

	default:

		fmt.Println("It's a weekday")

	}

	// switch without an expression is an alternate way to express if/else logic.
	// Here we also show how the case expressions can be non-constants.

	// We called the time package above to use its standard library

	// Package time provides functionality for measuring and displaying time.
	// The calendrical calculations always assume a Gregorian calendar, with no leap seconds.

	t := time.Now()

	switch {

	case t.Hour() < 12:

		fmt.Println("It's before noon")

	default:

		fmt.Println("It's after noon")

	}

	// A type switch compares types instead of values. You can use this to discover the type of an interface value.
	// In this example, the variable `t` will have the type corresponding to its clause.

	// A brief explanation about interface :
	// Interface is a concept similar to class, except that no method can be implemented in it.
	// In fact, an interface is just a set of contracts in which different methods are defined without defining body and function,
	// and the class that implements this interface must use that function within itself and define that function.
	// Each class can implement more than one interface.

	whatAmI := func(i interface{}) {

		switch t := i.(type) {

		case bool:

			fmt.Println("I'm a bool")

		case int:

			fmt.Println("I'm an int")

		default:

			fmt.Printf("Don't know type %T\n", t)

		}

	}
	whatAmI(true)

	whatAmI(1)

	whatAmI("hey")

}

// output :
// Write 2 as two
// It's a weekday
// It's after noon
// I'm a bool
//I'm an int
//Don't know type string
