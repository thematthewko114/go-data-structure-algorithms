package main

import (
	"bufio"
	"fmt"
	"go-dsa/utils"
	"os"
	"strconv"
	"strings"

	queueADT "go-dsa/queues"
	// searchingAlgo "go-dsa/searching-algo"
	// sortingAlgo "go-dsa/sorting-algo"
	stackADT "go-dsa/stacks"
	// "os"
)

func main() {
	modeArr := []string{"home", "stack", "queue", "hash table", "linked list", "binary tree", "linear search", "binary search", "interpolation search", "bubble sort", "quick sort", "selection sort", "insertion sort", "shell sort", "comb sort", "merge sort", "radix sort", "pancake sort", "rabin karp", "levenshtein distance", "longest common subsequence", "kmp algorithm"}
	stack := stackADT.NewStack()
	queue := queueADT.NewQueue()
	mode := 0
	fmt.Println(mode)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell for data structures")
	utils.PrintSeparator()
	utils.PrintHelp()
	fmt.Println("Enter the mode you want to enter:")
	for {
		fmt.Print(modeArr[mode])
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if text == "exit" {
			break
		} else if text == "--help" {
			utils.PrintHelp()
		}
		checkNum := text[1:]
		checkModeNum, _ := strconv.Atoi(checkNum)
		modeSwitch := false
		if strings.HasPrefix(text, "%") && checkModeNum <= 20 && checkModeNum >= 0 {
			mode, _ = strconv.Atoi(checkNum)
			fmt.Printf("Entered %s mode\n", modeArr[mode])
			modeSwitch = true
		}

		switch mode {
		case 1:
			if text == "--help" {
				utils.PrintStackHelp()
			} else if strings.HasPrefix(text, "push ") {
				stack.Push(text[5:])
			} else if text == "pop" {
				fmt.Println("Removed stack item:", stack.Pop())
			} else if modeSwitch == false {
				fmt.Println("Invalid command")
			}
			stack.PrintStack()
		case 2:
			if text == "--help" {
				utils.PrintQueueHelp()
			} else if strings.HasPrefix(text, "enqueue") {
				queue.Enqueue(text[8:])
			} else if text == "dequeue" {
				fmt.Println("Removed queue item:", queue.Dequeue())
			} else if modeSwitch == false {
				fmt.Println("Invalid command")
			}
			queue.PrintQueue()
		case 3:
			fmt.Println("Entering hash table mode")
		case 4:
			fmt.Println("Entering linked list mode")
		case 5:
			fmt.Println("Entering binary tree mode")
		case 6:
			fmt.Println("Entering linear search mode")
		case 7:
			fmt.Println("Entering binary search mode")
		case 8:
			fmt.Println("Entering interpolation search mode")
		case 9:
			fmt.Println("Entering bubble sort mode")
		case 10:
			fmt.Println("Entering quick sort mode")
		case 11:
			fmt.Println("Entering selection sort mode")
		case 12:
			fmt.Println("Entering insertion sort mode")
		case 13:
			fmt.Println("Entering shell sort mode")
		case 14:
			fmt.Println("Entering comb sort mode")
		case 15:
			fmt.Println("Entering merge sort mode")
		case 16:
			fmt.Println("Entering radix sort mode")
		case 17:
			fmt.Println("Entering pancake sort mode")
		case 18:
			fmt.Println("Entering Rabin Karp algorithm mode")
		case 19:
			fmt.Println("Entering Levenshtein distance mode")
		case 20:
			fmt.Println("Entering Longest common subsequence mode")
		case 21:
			fmt.Println("Entering KMP algorithm mode")
		default:
			fmt.Println("Unknown command")
		}
		utils.PrintSeparator()
	}
}
