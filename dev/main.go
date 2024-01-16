package main

import (
    "log"
)

func main() {
    chana := make(chan int)
    chanb := make(chan int)

    go func() {
        for i := 0; i < 1000; i++ {
            chana <- 100 * i
        }
    }()

    go func() {
        for i := 0; i < 1000; i++ {
            chanb <- i
        }
    }()

    acount := 0
    bcount := 0

    for {
        select {
        case <-chana:
            acount++
            log.Println("acount", acount)
        default:
        }

        select {
        case <-chanb:
            bcount++
            log.Println("bcount", bcount)
        default:
        }

        if acount == 1000 && bcount == 1000 {
            break
        }
    }
}
