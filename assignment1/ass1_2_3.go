// Question 2 Part 3
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

  // multiple nodes start electing simultaneously
  // assume 2 simultaneous elections
  all_nodes := []int{0,1,2,3,4}
  all_nodes = append(all_nodes[:faulty_node], all_nodes[faulty_node+1:]...)
  start_node_0 := all_nodes[rand.Intn(4)]
  if (start_node_0 != 4) {
    all_nodes = append(all_nodes[:start_node_0], all_nodes[start_node_0+1:]...)
  } else {
    all_nodes = all_nodes[:4]
  }
  start_node_1 := all_nodes[rand.Intn(3)]
  elected_0 := node(start_node_0, faulty_node, coordinators)
  fmt.Println("Elected 0: ", elected_0)
  elected_1 := node(start_node_1, faulty_node, coordinators)
  fmt.Println("Elected 1: ", elected_1)
  // from the print statements, we can see that there is no conflict
  // the elected node is still the same

  var input string
  fmt.Scanln(&input)
}
