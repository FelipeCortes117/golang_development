package pubsub

import (
    "fmt"
    "os"

    "github.com/FelipeCortes117/golang_development/gpio"
)

func main() {
    gpio.PubPanicButton = func(buttonMapJSON interface{}) {
        fmt.Println(buttonMapJson)
    }
    gpio.PubEngineButton = func(buttonMapJSON interface{}) {
        fmt.Println(buttonMapJson)
    }
}
