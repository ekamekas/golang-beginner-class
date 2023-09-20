package main

import "fmt";

func main() {
    fizzbuzz(15);
}

/**
  will run for n-times. If ith is multiple of 3, then print "Fizz".
  If ith is multiple of 5, then print "Buzz".
  If ith is multiple both 3 and 5, then print "FizzBuzz".
  Otherwise print ith.
*/
func fizzbuzz(n int) {
    for i := 1; i <= n; i++ {
        if(0 == i % 3 && 0 == i % 5) {
            fmt.Println("FizzBuzz");
        } else if(0 == i % 3) {
            fmt.Println("Fizz");
        } else if(0 == i % 5) {
            fmt.Println("Buzz");
        } else {
            fmt.Println(i);
        }
    }
}
