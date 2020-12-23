package main

import "fmt"

func main() {

    var a = "initial"
    fmt.Println(a)

    var b, c int = 1, 2     //var 声明 1 个或者多个变量
    fmt.Println(b, c)

    var d = true            //Go 会自动推断已经有初始值的变量的类型
    fmt.Println(d)

    var e int               //声明后却没有给出对应的初始值时，变量将会初始化为零值 
    fmt.Println(e)

    f := "short"            //:= 语法是声明并初始化变量的简写， 例如 var f string = "short" 可以简写为右边这样
    fmt.Println(f)
}
