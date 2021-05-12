package pubsub

import (
    "fmt"
    "os"

    "gitlab.com/EndTechCompany/iod-internet-open-device/raspgpsGolang/src/github.com/raspgps/code_gps_gsm_go/gpio"
)

func main() {
    gpio.PubPanicButton = func(buttonMapJSON interface{}) {
        fmt.Println(buttonMapJson)
    }
    gpio.PubEngineButton = func(buttonMapJSON interface{}) {
        fmt.Println(buttonMapJson)
    }
}
