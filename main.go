package main

import (
	"bufio"
	"fmt"
	binaryTreeADT "go-dsa/binarytrees"
	hashtableADT "go-dsa/hashtables"
	linkedlistADT "go-dsa/linkedlists"
	patternSearchingAlgo "go-dsa/pattern-seaching-algo"
	queueADT "go-dsa/queues"
	searchingAlgo "go-dsa/searching-algo"
	sortingAlgo "go-dsa/sorting-algo"
	stackADT "go-dsa/stacks"
	treeAlgo "go-dsa/tree-algo"
	"go-dsa/utils"
	"os"
	"strconv"
	"strings"
)

func main() {
	modeArr := []string{"home", "stack", "queue", "hash table", "linked list", "binary tree", "searching", "sorting", "pattern searching", "tree traversals"}
	stack := stackADT.NewStack()
	queue := queueADT.NewQueue()
	hashtable := hashtableADT.NewHashTable()
	linkedList := linkedlistADT.NewLinkedList()
	binaryTree := binaryTreeADT.NewBinaryTree()
	treeForTraversals := treeAlgo.NewBinaryTree()
	arrForSearching := utils.GetArray(5, true)
	arrForSorting := utils.GetArray(5, false)
	searchingInt := 0
	mode := 0
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
		if text == "--exit" {
			break
		} else if text == "--help" {
			utils.PrintHelp()
			continue
		} else if text == "--clean" {
			stack = stackADT.NewStack()
			queue = queueADT.NewQueue()
			hashtable = hashtableADT.NewHashTable()
			linkedList = linkedlistADT.NewLinkedList()
			binaryTree = binaryTreeADT.NewBinaryTree()
			treeForTraversals = treeAlgo.NewBinaryTree()
			arrForSearching = utils.GetArray(5, true)
			arrForSorting = utils.GetArray(5, false)
			continue
		}
		checkNum := text[1:]
		checkModeNum, _ := strconv.Atoi(checkNum)
		modeSwitch := false
		if strings.HasPrefix(text, "%") && checkModeNum <= 9 && checkModeNum >= 0 {
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
			} else if text == "print" {
				stack.PrintStack()
			} else if modeSwitch == false {
				fmt.Println("Invalid command")
			}
			stack.PrintStack()
		case 2:
			if text == "--help" {
				utils.PrintQueueHelp()
			} else if strings.HasPrefix(text, "enqueue ") {
				queue.Enqueue(text[8:])
			} else if text == "dequeue" {
				fmt.Println("Removed queue item:", queue.Dequeue())
			} else if text == "print" {
				queue.PrintQueue()
			} else if modeSwitch == false {
				fmt.Println("Invalid command")
			}
			queue.PrintQueue()
		case 3:
			if text == "--help" {
				utils.PrintHashTableHelp()
			} else if strings.HasPrefix(text, "insert ") {
				splicedInput := strings.Split(text, " ")
				hashtable.Insert(splicedInput[1], splicedInput[2])
				fmt.Println("Inserted", splicedInput[1], "to hash table to key", splicedInput[2])
				hashtable.PrintHashTable()
			} else if strings.HasPrefix(text, "get") {
				key := text[4:]
				fmt.Println("Retrieved from hash table:", hashtable.Get(key))
			} else if text == "print" {
				hashtable.PrintHashTable()
			} else if modeSwitch == false {
				fmt.Println("Invalid command")
			}
		case 4:
			if text == "--help" {
				utils.PrintLinkedListHelp()
			} else if strings.HasPrefix(text, "add ") {
				splicedInput := strings.Split(text, " ")
				linkedList.Add(splicedInput[1])
				fmt.Println("Added", splicedInput[1], "to linked list")
				linkedList.PrintList()
			} else if strings.HasPrefix(text, "remove ") {
				splicedInput := strings.Split(text, " ")
				linkedList.Remove(splicedInput[1])
				fmt.Println("Removed", splicedInput[1], "from linked list")
				linkedList.PrintList()
			} else if text == "print" {
				linkedList.PrintList()
			} else if modeSwitch == false {
				fmt.Println("Invalid command")
			}
		case 5:
			if text == "--help" {
				utils.PrintBinaryTreeHelp()
			} else if strings.HasPrefix(text, "insert ") {
				splicedInput := strings.Split(text, " ")
				value, _ := strconv.Atoi(splicedInput[1])
				binaryTree.Insert(value)
				fmt.Println("Inserted", splicedInput[1], "to binary tree")
				binaryTree.GetRoot().PrintTree(os.Stdout, 0, 'R')
			} else if strings.HasPrefix(text, "print") {
				binaryTree.GetRoot().PrintTree(os.Stdout, 0, 'R')
			} else if modeSwitch == false {
				fmt.Println("Invalid command")
			}
		case 6:
			if text == "--help" {
				utils.PrintSearchAlgoHelp()
			} else if strings.HasPrefix(text, "size ") {
				size, _ := strconv.Atoi(strings.Split(text, " ")[1])
				arrForSearching = utils.GetArray(size, true)
				fmt.Println(arrForSearching)
			} else if strings.HasPrefix(text, "find ") {
				searchingInt, _ = strconv.Atoi(strings.Split(text, " ")[1])
			} else if text == "search" {
				fmt.Println("--------------")
				fmt.Println("Linear search")
				linExists, linExetime := searchingAlgo.LinearSearch(arrForSearching, searchingInt)
				fmt.Printf("Linear Search result: %v\n", linExists)
				fmt.Print("Execution time: ")
				fmt.Print(linExetime)
				fmt.Printf("\n")
				fmt.Println("Time complexity: O(n), space complexity: O(n)")

				fmt.Println("--------------")
				fmt.Println("Binary search")
				binExists, binExetime := searchingAlgo.BinarySearch(arrForSearching, searchingInt)
				fmt.Printf("Binary Search result: %v\n", binExists)
				fmt.Print("Execution time: ")
				fmt.Print(binExetime)
				fmt.Printf("\n")
				fmt.Println("Time complexity: O(log n), space complexity: O(1)")

				fmt.Println("--------------")
				fmt.Println("Interpolation search")
				intExists, intExetime := searchingAlgo.InterpolationSearch(arrForSearching, searchingInt)
				fmt.Printf("Interpolation Search result: %v\n", intExists)
				fmt.Print("Execution time: ")
				fmt.Print(intExetime)
				fmt.Printf("\n")
				fmt.Println("Time complexity: O(log log n), space complexity: O(1)")
			} else if modeSwitch == false {
				fmt.Println("Invalid command")
			}
		case 7:
			if text == "--help" {
				utils.PrintSortingAlgoHelp()
			} else if strings.HasPrefix(text, "size ") {
				size, _ := strconv.Atoi(strings.Split(text, " ")[1])
				arrForSorting = utils.GetArray(size, false)
				fmt.Println(arrForSorting)
			} else if text == "sort" {
				fmt.Println("--------------")
				fmt.Println("Bubble sort")
				_, bubbleExetime := sortingAlgo.BubbleSort(arrForSorting)
				fmt.Print("Execution time: ")
				fmt.Print(bubbleExetime)
				fmt.Printf("\n")
				fmt.Println("Time complexity: O(n^2), space complexity: O(1)")

				// quick sort
				fmt.Println("--------------")
				fmt.Println("Quick sort")
				_, quickExetime := sortingAlgo.QuickSort(arrForSorting)
				fmt.Print("Execution time: ")
				fmt.Print(quickExetime)
				fmt.Printf("\n")
				fmt.Println("Time complexity: O(n log n), space complexity: O(1)")

				// selection sort
				fmt.Println("--------------")
				fmt.Println("Selection sort")
				_, selectionExetime := sortingAlgo.SelectionSort(arrForSorting)
				fmt.Print("Execution time: ")
				fmt.Print(selectionExetime)
				fmt.Printf("\n")
				fmt.Println("Time complexity: O(n^2), space complexity: O(1)")

				// insertion sort
				fmt.Println("--------------")
				fmt.Println("Insertion sort")
				_, insertionExetime := sortingAlgo.InsertionSort(arrForSorting)
				fmt.Print("Execution time: ")
				fmt.Print(insertionExetime)
				fmt.Printf("\n")
				fmt.Println("Time complexity: O(n^2), space complexity: O(1)")

				// shell sort
				fmt.Println("--------------")
				fmt.Println("Shell sort")
				_, shellExetime := sortingAlgo.ShellSort(arrForSorting)
				fmt.Print("Execution time: ")
				fmt.Print(shellExetime)
				fmt.Printf("\n")
				fmt.Println("Time complexity: O(n log n), space complexity: O(1)")

				// Comb sort
				fmt.Println("--------------")
				fmt.Println("Comb sort")
				_, combExetime := sortingAlgo.CombSort(arrForSorting)
				fmt.Print("Execution time: ")
				fmt.Print(combExetime)
				fmt.Printf("\n")
				fmt.Println("Time complexity: O(n log n), space complexity: O(1)")

				// Merge sort
				fmt.Println("--------------")
				fmt.Println("Merge sort")
				_, mergeExetime := sortingAlgo.MergeSort(arrForSorting)
				fmt.Print("Execution time: ")
				fmt.Print(mergeExetime)
				fmt.Printf("\n")
				fmt.Println("Time complexity: O(n log n), space complexity: O(1)")

				// Radix sort
				fmt.Println("--------------")
				fmt.Println("Radix sort")
				_, radixExetime := sortingAlgo.RadixSort(arrForSorting)
				fmt.Print("Execution time: ")
				fmt.Print(radixExetime)
				fmt.Printf("\n")
				fmt.Println("Time complexity: O(n log n), space complexity: O(1)")

				// Pancake sort
				fmt.Println("--------------")
				fmt.Println("Pancake sort")
				_, pancakeExetime := sortingAlgo.PancakeSort(arrForSorting)
				fmt.Print("Execution time: ")
				fmt.Print(pancakeExetime)
				fmt.Printf("\n")
				fmt.Println("Time complexity: O(n^2), space complexity: O(1)")
			} else if modeSwitch == false {
				fmt.Println("Invalid command")
			}
		case 8:
			if text == "--help" {
				utils.PatternSearchAlgoHelp()
			} else if strings.HasPrefix(text, "rabin-karp ") {
				splicedInput := strings.Split(text[11:], " | ")
				string1 := splicedInput[0]
				string2 := strings.Split(splicedInput[1], " ")
				fmt.Println("--------------")
				fmt.Println("Rabin Karp algorithm")
				matches := patternSearchingAlgo.RabinKarpAlgo(string1, string2)
				fmt.Println("Rabin Karp matches", matches)
				fmt.Println("Time complexity: O(mn), space complexity: O(1)")
			} else if strings.HasPrefix(text, "levenshtein ") {
				splicedInput := strings.Split(text[12:], " | ")
				string1 := splicedInput[0]
				string2 := splicedInput[1]
				fmt.Println("--------------")
				fmt.Println("Levenshtein distance")
				score := patternSearchingAlgo.LevenshteinDistance(string1, string2)
				fmt.Println("String distance: ", score)
				fmt.Println("Time complexity: O(mn), space complexity: O(n)")
			} else if strings.HasPrefix(text, "lcs") {
				splicedInput := strings.Split(text[3:], " | ")
				string1 := splicedInput[0]
				string2 := splicedInput[1]
				fmt.Println("--------------")
				fmt.Println("Longest common subsequence")
				lcs := patternSearchingAlgo.LCS(string1, string2, len(string1), len(string2))
				fmt.Println("Length of longest common subsequence", lcs)
				fmt.Println("Time complexity: O(2^mn), space complexity: O(1)")
			} else if strings.HasPrefix(text, "kmp ") {
				splicedInput := strings.Split(text[4:], " | ")
				string1 := splicedInput[0]
				string2 := splicedInput[1]
				fmt.Println("--------------")
				fmt.Println("KMP algorithm")
				index := patternSearchingAlgo.KMP_SearchString(string1, string2)
				fmt.Println("KMP index", index)
				fmt.Println("Time complexity: O(m+n), space complexity: O(m)")
			} else if modeSwitch == false {
				fmt.Println("Invalid command")
			}
		case 9:
			if text == "--help" {
				utils.TreeAlgoHelp()
			} else if strings.HasPrefix(text, "insert ") {
				splicedInput := strings.Split(text[7:], " ")
				weight, err := strconv.Atoi(splicedInput[1])
				if err != nil {
					fmt.Println(err.Error())
				}
				treeForTraversals.Insert(splicedInput[0], weight)
				treeForTraversals.GetRoot().PrintTree(os.Stdout, 0, 'R')
			} else if text == "traverse" {
				fmt.Println("Pre-order traversal:", treeAlgo.PreorderTraversal(treeForTraversals.GetRoot()))
				fmt.Println("In-order traversal:", treeAlgo.InorderTraversal(treeForTraversals.GetRoot()))
				fmt.Println("Post-order traversal:", treeAlgo.PostorderTraversal(treeForTraversals.GetRoot()))
			} else if modeSwitch == false {
				fmt.Println("Invalid command")
			}
		default:
			fmt.Println("Unknown command")
		}
		utils.PrintSeparator()
	}
}
