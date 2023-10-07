package main

import (
	"errors"
	"fmt"
)

func main() {
	list := []string{"brendo", "hermanoteu", "isie", "marina", "micalateia", "monalisa"}
	name, err := binarySearch(list, "brendo")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(name)
}

func binarySearch(list []string, name string) (string, error) {
	low := 0
	high := len(list)
	for low <= high {
		middle := (low + high) / 2
		current := list[middle]
		if current == name {
			return current, nil
		} else if current > name {
			high = middle - 1
		} else {
			low = middle + 1
		}
	}
	return "", errors.New("not found")
}
