// Question 1 Part 2
package main

import (
  "fmt"
  "time"
  "math/rand"
)

// send integer array to server periodically (i.e. at constant time intervals)
// each client updates the value at its own index
// for example, client 2 increments the value of index 2
func sender(id int, clock chan []int) {
  msg := []int{0,0,0,0,0}
  for i := 0 ; ; i++ {
    fmt.Println("Sending ", msg, " from Client ", id)
    clock <- msg
    time.Sleep(time.Millisecond * 1000)
    msg[id] = msg[id] + 1
  }
}

// receives message from client
func serverReceiver(server0 [5]chan []int, server1 [5]chan []int) {
  for {
    select {
    case msg0 := <- server0[0]:
      go fan(server1, msg0)
      server_clock[0] = server_clock[0] + 1
      fmt.Println("Received ", msg0, " from Client ", 0, "| Order: ", server_clock)
    case msg1 := <- server0[1]:
      go fan(server1, msg1)
      server_clock[1] = server_clock[1] + 1
      fmt.Println("Received ", msg1, " from Client ", 1, "| Order: ", server_clock)
    case msg2 := <- server0[2]:
      go fan(server1, msg2)
      server_clock[2] = server_clock[2] + 1
      fmt.Println("Received ", msg2, " from Client ", 2, "| Order: ", server_clock)
    case msg3 := <- server0[3]:
      go fan(server1, msg3)
      server_clock[3] = server_clock[3] + 1
      fmt.Println("Received ", msg3, " from Client ", 3, "| Order: ", server_clock)
    case msg4 := <- server0[4]:
      go fan(server1, msg4)
      server_clock[4] = server_clock[4] + 1
      fmt.Println("Received ", msg4, " from Client ", 4, "| Order: ", server_clock)
    }
  }
}

// randomly delay before broadcasting message to all clients
func fan(server [5]chan []int, msg []int) {
  amt := time.Duration(rand.Intn(1000))
  time.Sleep(time.Millisecond * amt)
  for i := range server {
    // fmt.Println("Broadcasting ", msg, " to Client ", i)
    server[i] <- msg
  }
}

// print message received from server
func clientReceiver(id int, clock chan []int) {
  for {
    msg := <- clock
    fmt.Println("Client ", id, " received ", msg, " from server")
  }
}

// server_clock gives the total order of the messages receive
// each index corresponds to a client
// the value of each index corresponds to the order
var server_clock [5]int
func main() {
  // server0 is for clients to send messages to the server
  // server1 is for the server to broadcast messages to the clients
  var server0 [5]chan []int
  var server1 [5]chan []int
  for i := range server0 {
    server0[i] = make(chan []int)
    go sender(i, server0[i])
    server1[i] = make(chan []int)
    go clientReceiver(i, server1[i])
  }
  go serverReceiver(server0, server1)

  var input string
  fmt.Scanln(&input)
}
