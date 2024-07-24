package main

import (
  "machine"
  "time"
)

func main () {
  led := machine.LED
  led.Configure(machine.PinConfig{Mode: machine.PinOutput})

  for {
    println("Hello world", "\r")
    led.Toggle()
    time.Sleep(100 * time.Millisecond)
  }
}
