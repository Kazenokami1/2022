package main

import (
	"fmt"
	"os"
	"reflect"
	"time"
)

func main() {
	start := time.Now()
	d := Days{}
	dMethod := reflect.ValueOf(d).MethodByName(os.Getenv("DAYTORUN"))
	dMethod.Call(nil)
	duration := time.Since(start)
	fmt.Print("Time Since Start: ")
	fmt.Println(duration)
}
