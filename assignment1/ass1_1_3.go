// Question 1 Part 3
package main

import (
  "fmt"
  "time"
  "math/rand"
)

// send integer array to server periodically (i.e. at constant time intervals)
// each client updates the value at its own index
// for example, client 2 increments the value of index 2

// every client increments the same message
// and there is a delay between the increments
// so a causality violation does not happen
// a causality violation happens when there is no/minimal delay between the increments
// this could be because messages are being broadcasted
// before they are synchronized in the server
// it could also be because clients are sending at different rates
// due to latency issues
func sender(id int, clock chan []int, msg []int) {
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
      // fmt.Println("Received ", msg0, " from Client ", 0)
      go fan(server1, msg0)
    case msg1 := <- server0[1]:
      // fmt.Println("Received ", msg1, " from Client ", 1)
      go fan(server1, msg1)
    case msg2 := <- server0[2]:
      // fmt.Println("Received ", msg2, " from Client ", 2)
      go fan(server1, msg2)
    case msg3 := <- server0[3]:
      // fmt.Println("Received ", msg3, " from Client ", 3)
      go fan(server1, msg3)
    case msg4 := <- server0[4]:
      // fmt.Println("Received ", msg4, " from Client ", 4)
      go fan(server1, msg4)
    }
  }
}

// randomly delay before broadcasting message to all clients
func fan(server [5]chan []int, msg []int) {
  amt := time.Duration(rand.Intn(1000))
  time.Sleep(time.Millisecond * amt)
  for i := range server {
    server[i] <- msg
    fmt.Println("Broadcasting ", msg, " to Client ", i)
  }
}

// print message received from server

// we can check for a causality violation here
// if one client receive different messages from the server,
// that means that a causality violation has occurred
// in addition, because the increments are periodic
// all clients should increment at the same rate
// i.e. the value in every index of the message array should be the same
// a causality violation occurs if the values in the message array are different
func clientReceiver(id int, clock chan []int) {
  for {
    msg := <- clock
    fmt.Println("Client ", id, " received ", msg, " from server")
  }
}

func main() {
  // server0 is for clients to send messages to the server
  // server1 is for the server to broadcast messages to the clients
  var server0 [5]chan []int
  var server1 [5]chan []int
  msg := []int{0,0,0,0,0}
  for i := range server0 {
    server0[i] = make(chan []int)
    go sender(i, server0[i], msg)
    server1[i] = make(chan []int)
    go clientReceiver(i, server1[i])
  }
  go serverReceiver(server0, server1)

  var input string
  fmt.Scanln(&input)
}
