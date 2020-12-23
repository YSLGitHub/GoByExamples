package main

import (
    "fmt"
    "time"
)

func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done")

    done <- true
}

func main() {

    done := make(chan bool, 1)
    go worker(done)

    <-done		//<-done 将会等待从 channel 中读取数据，直到 say 中执行了 done <- struct{}{} ，即子协程完成任务，才会退出。
}
