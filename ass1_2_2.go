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
  // let's kill a randomly selected node
  kill_node := rand.Intn(5)
  if coordinators[kill_node] != -1 {
    fmt.Println("Kill Node ", kill_node)
    coordinators[kill_node] = -1
  }
  // sometimes the killed node still appears to broadcast in the prints
  // this is because an election could have been called before the node
  // was determined to be faulty
  for i := 0 ; i < 5 ; i++ {
    if (coordinators[i] != -1) {
      fmt.Println("Node ", id, " broadcasts to Node ", i)
      // if the ID of the node receiving the broadcast is higher
      // it does not respond
      if id >= i {
        coordinators[i] = id
      } else {
         _ = node(i, faulty_id, coordinators)
      }
    } else {
      continue
    }
  }

  if check_coordinators(faulty_id, coordinators) {
    var n int
    for i := 1 ; i < 5 ; i++ {
      if coordinators[i] != -1 {
        n = coordinators[i]
      }
    }
    if coordinators[coordinators[n]] == -1 {
      return -1
    } else {
      return coordinators[n]
    }
  }
  return -1
}

func check_coordinators(faulty_id int, coordinators []int) (bool) {
  // fmt.Println("Coordinators: ", coordinators)
  for i := 1 ; i < 5 ; i++ {
    if (coordinators[i] != coordinators[i-1]) && (coordinators[i] != -1) && (coordinators[i-1] != -1){
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
  coordinators := []int{-2,-2,-2,-2,-2}
  coordinators[faulty_node] = -1

  all_nodes := []int{0,1,2,3,4}
  all_nodes = append(all_nodes[:faulty_node], all_nodes[faulty_node+1:]...)
  start_node := all_nodes[rand.Intn(4)]
  elected := node(start_node, faulty_node, coordinators)
  if elected != -1 {
    fmt.Println("Elected: Node ", elected)
  } else {
    fmt.Println("Deadlock")
  }

  // the asynchronous processes make it challenging to tell
  // which process happened first
  // but from the print results
  // we can see that having nodes fail during the election
  // can result in a deadlock
  // this could happen when no other nodes can call an election
  // and the largest node so far fails

  var input string
  fmt.Scanln(&input)
}
