// Question 2 Part 1
package main

import (
  "fmt"
  "time"
  "math/rand"
)

func node(id int, faulty_id int, coordinators []int) (int){
  // the node calls an election
  // by sending its id to all the other nodes
  // except for the faulty node
  fmt.Println("Node ", id, " calls an election")
  for i := 0 ; i < 5 ; i++ {
    if (i != faulty_id) {
      fmt.Println("Node ", id, " broadcasts to Node ", i)
      // if the ID of the node receiving the broadcast is higher
      // it does not respond
      if id >= i {
        coordinators[i] = id
      } else {
         _ = node(i, faulty_id, coordinators)
      }
    }
  }

  if check_coordinators(faulty_id, coordinators) {
    return coordinators[0]
  }
  return -1
}

func check_coordinators(faulty_id int, coordinators []int) (bool) {
  coordinators = append(coordinators[:faulty_id], coordinators[faulty_id+1:]...)
  for i := 1 ; i < 4 ; i++ {
    if coordinators[i] != coordinators[i-1] {
      return false
    }
  }
  return true
}

func main() {
  // use Seed to get different random values for each run
  rand.Seed(time.Now().UnixNano())

  // assume we have 5 nodes
  // randomly select a faulty node
  faulty_node := rand.Intn(5)
  fmt.Println("Faulty node: ", faulty_node)

  // coordinators give the coordinator of each node
  // the indices correspond to the nodes
  coordinators := []int{-1,-1,-1,-1,-1}

  // best-case: largest node calls an election
  // if faulty_node == 4 {
  //   elected := node(3, faulty_node, coordinators)
  //   fmt.Println("Elected: Node ", elected)
  // } else {
  //   elected := node(4, faulty_node, coordinators)
  //   fmt.Println("Elected: ", elected)
  // }

  // worst-case: smallest node calls an election
  if faulty_node == 0 {
    elected := node(1, faulty_node, coordinators)
    fmt.Println("Elected: ", elected)
  } else {
    elected := node(0, faulty_node, coordinators)
    fmt.Println("Elected: ", elected)
  }

  // the print output has some repeats
  // this is because some of the calls for elections run concurrently

  var input string
  fmt.Scanln(&input)
}
