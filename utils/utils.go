package utils

import (
	"fmt"
	"math/rand"
	"sort"
)

func PrintSeparator() {
	fmt.Println("================================")
}

func PrintHelp() {
	fmt.Println("=====================")
	fmt.Println("Guide to use this shell:")
	fmt.Println("Enter a mode first to use different data structures")
	fmt.Println("Type %0 to go back to home mode")
	fmt.Println("Type '%1' to use stacks")
	fmt.Println("Type '%2' to use queues")
	fmt.Println("Type '%3' to use hash tables")
	fmt.Println("Type '%4' to use linked lists")
	fmt.Println("Type '%5' to use binary trees")
	fmt.Println("Type '%6' to use linear search")
	fmt.Println("Type '%7' to use binary search")
	fmt.Println("Type '%8' to use interpolation search")
	fmt.Println("Type '%9' to use bubble sort")
	fmt.Println("Type '%10' to use quick sort")
	fmt.Println("Type '%11' to use selection sort")
	fmt.Println("Type '%12' to use insertion sort")
	fmt.Println("Type '%13' to use shell sort")
	fmt.Println("Type '%14' to use comb sort")
	fmt.Println("Type '%15' to use merge sort")
	fmt.Println("Type '%16' to use radix sort")
	fmt.Println("Type '%17' to use pancake sort")
	fmt.Println("Type '%18' to use Rabin Karp algorithm")
	fmt.Println("Type '%19' to use Levenshtein distance")
	fmt.Println("Type '%20' to use Longest common subsequence")
	fmt.Println("Type '%21' to use KMP algorithm")
	fmt.Println("Type --help to read manual inside the mode")
}

func PrintStackHelp() {
	fmt.Println("=====================")
	fmt.Println("Guide to use this stack:")
	fmt.Println("You are given a stack, you can perform push and pop operations")
	fmt.Println("Type 'push <string>' to push an element to the stack, e.g. 'push hello'")
	fmt.Println("Type 'pop' to pop an element from the stack, and the return value will be printed, e.g. 'pop'")
}

func PrintQueueHelp() {
	fmt.Println("=====================")
	fmt.Println("Guide to use this queue:")
	fmt.Println("You are given a queue, you can perform enqueue and dequeue operations")
	fmt.Println("Type 'enqueue <string>' to enqueue an element to the queue, e.g. 'enqueue hello'")
	fmt.Println("Type 'dequeue' to dequeue an element from the queue, and the return value will be printed, e.g. 'dequeue'")
}

func GetArray(n int, sorted bool) []int {
	array := make([]int, n)
	for i := 0; i < n; i++ {
		array[i] = rand.Intn(10000000) // Generate a random number between 0 and 99
	}
	if sorted == true {
		sort.Ints(array)
	}
	return array
}
