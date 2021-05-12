package gpio

import (
    "fmt"
    "log"
    "encoding/json"
    "time"
    "math/rand"

    "periph.io/x/periph/conn/gpio"
    "periph.io/x/periph/conn/gpio/gpioreg"
    "periph.io/x/periph/host"
)

const (
    // Low represents 0v.
    Low= false
    // High represents Vin, generally 3.3v or 5v.
    High= true
)

type CallBackPanic func(interface{})
type CallBackEngine func(interface{})

var PubPanicButton CallBackPanic
var PubEngineButton CallBackEngine

func blockForever() {
    c := make(chan struct{})
    <-c
}

func Start() {
    rand.Seed(time.Now().UTC().UnixNano())

    fmt.Println("started GPIO")
    //Make sure periph is initialized.
    if _, err := host.Init(); err != nil {
        log.Fatal(err)
    }


    go goPanicButton()
    go goEngineButton()

    blockForever()
}

func goPanicButton() {
    gpioPanic := gpioreg.ByName("GPIO23")
    if err := gpioPanic.In(gpio.PullDown, gpio.RisingEdge); err != nil {
        log.Fatal(err)
    }

    valuePanic := gpioPanic.Read()

    for {
        panicMap := make(map[string]string)

        actualPanic := gpioPanic.Read()
        if valuePanic != actualPanic {
            valuePanic = actualPanic
            if valuePanic == High {
                panicMap["Panic"] = "ON"
            } else {
                panicMap["Panic"] = "OFF"
            }
        }
        panicJSON, _ := json.Marshal(panicMap)
        PubPanicButton(panicJSON)

        time.Sleep(time.Second * 7)
    }
}

func goEngineButton() {
    gpioEngine := gpioreg.ByName("GPIO18")
    if err := gpioEngine.In(gpio.PullDown, gpio.RisingEdge); err != nil {
        log.Fatal(err)
    }

    valueEngine := gpioEngine.Read()

    for {
        engineMap := make(map[string]string)

        actualEngine := gpioEngine.Read()
        if valueEngine != actualEngine {
            valueEngine = actualEngine
            if valueEngine == High {
                engineMap["Engine"] = "ON"
            } else {
                engineMap["engine"] = "OFF"
            }
        }
        engineJSON, _ := json.Marshal(engineMap)
        PubEngineButton(engineJSON)

        time.Sleep(time.Second * 7)
    }
}
