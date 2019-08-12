package slices

import "fmt"

// Tests equality of string arrays regardless of element order.
func Equal(slice1 []string, slice2 []string) bool {
  sliceMap := make(map[string]int)
  for _, value := range slice1 {
    sliceMap[value] = sliceMap[value] + 1
  }
  for _, value := range slice2 {
    if sliceMap[value] == 0 {
      return false
    } else {
      sliceMap[value] = sliceMap[value] - 1
    }
    if sliceMap[value] == 0 { 
      delete(sliceMap, value)
    }
  }
  if len(sliceMap) != 0 {
    return false
  }
  return true
}

// Non-Functional
func Intersection(slice1 []string, slice2 []string) []string {
 map1 := make(map[string]int)
 map2 := make(map[string]int)
 for _, value := range slice1 {
   map1[value] = map1[value] + 1
 }
 for _, value := range slice2 {
   map2[value] = map2[value] + 1
 }
 fmt.Println("HEllO")
 fmt.Println(map1)
 fmt.Println(map2)
 return []string{"hi"}
}