package main

import (
	"fmt"
	"os"
	"reflect"
	"time"
)

func main() {
	start := time.Now()
	duration := time.Since(start)
	d := Days{}
	dMethod := reflect.ValueOf(d).MethodByName(os.Getenv("DAYTORUN"))
	dMethod.Call(nil)
	fmt.Print("Time Since Start: ")
	fmt.Println(duration)
}
