package main

import (
    "fmt"
    "math"
)

const s string = "constant"

func main() {
    fmt.Println(s)

    const n = 500000000

    const d = 3e20 / n
    fmt.Println(d)

    fmt.Println(int64(d))       //数值型常量没有确定的类型，直到被给定某个类型，比如显式类型转化

    fmt.Println(math.Sin(n))    //一个数字可以根据上下文的需要（比如变量赋值、函数调用）自动确定类型
}
