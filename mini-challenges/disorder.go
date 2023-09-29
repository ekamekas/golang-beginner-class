package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup

    for i := 0; i < 4; i++ {
        wg.Add(2);
        go print(&wg, []string{ "coba1", "coba2", "coba3" }, i + 1); 
        go print(&wg, []string{ "bisa1", "bisa2", "bisa3" }, i + 1); 
    }

    wg.Wait();
}

func print(wg *sync.WaitGroup, data ...interface{}) {
    defer wg.Done();

    fmt.Printf("%+v\n", data);
}
