package demo

import (
	"fmt"
	binarytreeADT "go-dsa/binarytrees"
	hashtableADT "go-dsa/hashtables"
	linkedlistADT "go-dsa/linkedlists"
	patternSeachingAlgo "go-dsa/pattern-seaching-algo"
	queueADT "go-dsa/queues"
	searchingAlgo "go-dsa/searching-algo"
	sortingAlgo "go-dsa/sorting-algo"
	stackADT "go-dsa/stacks"
	utils "go-dsa/utils"
	"os"
)

func main() {
	// stacks
	fmt.Println("--------------")
	fmt.Println("Stacks")
	stack := stackADT.NewStack()
	stack.Push("one")
	stack.Push("two")
	stack.PrintStack()

	// queues
	fmt.Println("--------------")
	fmt.Println("Queues")
	queue := queueADT.NewQueue()
	queue.Enqueue("one")
	queue.Enqueue("two")
	queue.Enqueue("three")
	val := queue.Dequeue()
	queue.Enqueue(val)
	queue.PrintQueue()

	// hash tables
	fmt.Println("--------------")
	fmt.Println("Hash tables")
	hashtable := hashtableADT.NewHashTable()
	hashtable.Insert(1, "one")
	fmt.Println(hashtable.Get(1))

	// linked lists
	fmt.Println("--------------")
	fmt.Println("Linked lists")
	linkedlist := linkedlistADT.NewLinkedList()
	linkedlist.Add(1)
	linkedlist.Add(2)
	linkedlist.Add(3)
	linkedlist.Remove(2)
	linkedlist.PrintList()

	// binary trees
	fmt.Println("--------------")
	fmt.Println("Binary trees")
	binaryTree := binarytreeADT.NewBinaryTree()
	binaryTree.Insert(5).Insert(1).Insert(2).Insert(3).Insert(4).Insert(6).Insert(7).Insert(8)
	binaryTree.GetRoot().PrintTree(os.Stdout, 0, 'M')

	arr := utils.GetArray(200000, true)

	// linear search
	fmt.Println("--------------")
	fmt.Println("Linear search")
	linExists, linExetime := searchingAlgo.LinearSearch(arr, 1000)
	fmt.Printf("Linear Search result: %v\n", linExists)
	fmt.Print("Execution time: ")
	fmt.Print(linExetime)
	fmt.Printf("\n")
	fmt.Println("Time complexity: O(n), space complexity: O(n)")

	// binary search
	fmt.Println("--------------")
	fmt.Println("Binary search")
	binExists, binExetime := searchingAlgo.BinarySearch(arr, 1000)
	fmt.Printf("Binary Search result: %v\n", binExists)
	fmt.Print("Execution time: ")
	fmt.Print(binExetime)
	fmt.Printf("\n")
	fmt.Println("Time complexity: O(log n), space complexity: O(1)")

	// interpolation search
	fmt.Println("--------------")
	fmt.Println("Interpolation search")
	intExists, intExetime := searchingAlgo.InterpolationSearch(arr, 1000)
	fmt.Printf("Interpolation Search result: %v\n", intExists)
	fmt.Print("Execution time: ")
	fmt.Print(intExetime)
	fmt.Printf("\n")
	fmt.Println("Time complexity: O(log log n), space complexity: O(1)")

	arr = utils.GetArray(200000, false)
	// bubble sort
	fmt.Println("--------------")
	fmt.Println("Bubble sort")
	_, bubbleExetime := sortingAlgo.BubbleSort(arr)
	fmt.Print("Execution time: ")
	fmt.Print(bubbleExetime)
	fmt.Printf("\n")
	fmt.Println("Time complexity: O(n^2), space complexity: O(1)")

	// quick sort
	fmt.Println("--------------")
	fmt.Println("Quick sort")
	_, quickExetime := sortingAlgo.QuickSort(arr)
	fmt.Print("Execution time: ")
	fmt.Print(quickExetime)
	fmt.Printf("\n")
	fmt.Println("Time complexity: O(n log n), space complexity: O(1)")

	// selection sort
	fmt.Println("--------------")
	fmt.Println("Selection sort")
	_, selectionExetime := sortingAlgo.SelectionSort(arr)
	fmt.Print("Execution time: ")
	fmt.Print(selectionExetime)
	fmt.Printf("\n")
	fmt.Println("Time complexity: O(n^2), space complexity: O(1)")

	// insertion sort
	fmt.Println("--------------")
	fmt.Println("Insertion sort")
	_, insertionExetime := sortingAlgo.InsertionSort(arr)
	fmt.Print("Execution time: ")
	fmt.Print(insertionExetime)
	fmt.Printf("\n")
	fmt.Println("Time complexity: O(n^2), space complexity: O(1)")

	// shell sort
	fmt.Println("--------------")
	fmt.Println("Shell sort")
	_, shellExetime := sortingAlgo.ShellSort(arr)
	fmt.Print("Execution time: ")
	fmt.Print(shellExetime)
	fmt.Printf("\n")
	fmt.Println("Time complexity: O(n log n), space complexity: O(1)")

	// Comb sort
	fmt.Println("--------------")
	fmt.Println("Comb sort")
	_, combExetime := sortingAlgo.CombSort(arr)
	fmt.Print("Execution time: ")
	fmt.Print(combExetime)
	fmt.Printf("\n")
	fmt.Println("Time complexity: O(n log n), space complexity: O(1)")

	// Merge sort
	fmt.Println("--------------")
	fmt.Println("Merge sort")
	_, mergeExetime := sortingAlgo.MergeSort(arr)
	fmt.Print("Execution time: ")
	fmt.Print(mergeExetime)
	fmt.Printf("\n")
	fmt.Println("Time complexity: O(n log n), space complexity: O(1)")

	// Radix sort
	fmt.Println("--------------")
	fmt.Println("Radix sort")
	_, radixExetime := sortingAlgo.RadixSort(arr)
	fmt.Print("Execution time: ")
	fmt.Print(radixExetime)
	fmt.Printf("\n")
	fmt.Println("Time complexity: O(n log n), space complexity: O(1)")

	// Pancake sort
	fmt.Println("--------------")
	fmt.Println("Pancake sort")
	_, pancakeExetime := sortingAlgo.PancakeSort(arr)
	fmt.Print("Execution time: ")
	fmt.Print(pancakeExetime)
	fmt.Printf("\n")
	fmt.Println("Time complexity: O(n^2), space complexity: O(1)")

	// Rabin Karp algorithm
	fmt.Println("--------------")
	fmt.Println("Rabin Karp algorithm")
	txt := "Help me find a substring"
	patterns := []string{"you", "string"}
	matches := patternSeachingAlgo.RabinKarpAlgo(txt, patterns)
	fmt.Println(matches)
	fmt.Println("Time complexity: O(mn), space complexity: O(1)")

	// Levenshtein distance
	fmt.Println("--------------")
	fmt.Println("Levenshtein distance")
	txt1 := "Help me find a substring"
	txt2 := "Help me find a substring"
	score := patternSeachingAlgo.LevenshteinDistance(txt1, txt2)
	fmt.Println("String distance: ", score)
	fmt.Println("Time complexity: O(mn), space complexity: O(n)")

	// Longest common subsequence
	fmt.Println("--------------")
	fmt.Println("Longest common subsequence")
	lcsLength := patternSeachingAlgo.LCS("AABAACAADAABAABAAABAACAADAABAABA", "CABABBDD", 32, 8)
	fmt.Println("LCS: ", lcsLength)
	fmt.Println("Time complexity: O(2^mn), space complexity: O(1)")

	// KMP algorithm
	fmt.Println("--------------")
	fmt.Println("KMP algorithm")
	// txt3 := "Help me find a substring"
	// txt4 := "Help you find a substring"
	kmpLength := patternSeachingAlgo.KMP_SearchString("AABAACAADAABAABAAABAACAADAABAABA", "AABA")
	fmt.Println("KMP index: ", kmpLength)
	fmt.Println("Time complexity: O(m+n), space complexity: O(m)")
}
