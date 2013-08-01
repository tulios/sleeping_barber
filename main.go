package main

import (
  "fmt"
  "time"
)

type Client struct {
	name string
}

func barber() chan<- Client{
  channel := make(chan Client)
  go func() {
    for client := range channel {
      fmt.Printf("cutting %s...\n", client.name)
      time.Sleep(time.Millisecond)
    }
  }()
  return channel
}

func chair(barber chan<- Client) chan<- Client {
  channel := make(chan Client, 1)
  go func() {
    for client := range channel {
      fmt.Printf("%s waiting...\n", client.name)
      barber <- client
    }
  }()
  return channel
}

func main() {
  clients := []Client{
    {name: "A"}, {name: "B"}, {name: "C"}, {name: "D"}, {name: "E"}, {name: "F"}, {name: "G"}, {name: "H"},
  }

  barber := barber()
  chair1 := chair(barber)
  chair2 := chair(barber)
  chair3 := chair(barber)

  for _, client := range clients {
    fmt.Printf("%s arrives\n", client.name)

    select {
    case chair1 <- client:
    case chair2 <- client:
    case chair3 <- client:
    default:
      fmt.Printf("%s leaving...\n", client.name)
    }
  }

  time.Sleep(time.Second)
}
