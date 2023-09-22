package main

import "fmt"

func main() {
    word := "selamat malam";
    count := map[string]int{};  // key -> string character; value -> total count char in word

    for _, asciiCode := range word {
        char := string(asciiCode);  // will convert ascii representation to a string

        mCount, doesExist := count[char];

        if(!doesExist) {
            mCount = 0;
        }

        count[char] = mCount + 1;

        fmt.Println(char);
    }

    fmt.Println(count);
}
