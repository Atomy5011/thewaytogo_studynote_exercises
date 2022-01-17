package main

import "fmt"

func init() {
   fmt.Println("init 1")
}

func init() {
   fmt.Println("init 2")
}

func InitExample() {
   fmt.Println("main")
}