package sortingAlgo

import (
	"bytes"
	"encoding/binary"
	"math/rand"
	"time"
)

func BubbleSort(arr []int) ([]int, time.Duration) {
	startTime := time.Now()
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	endTime := time.Now()
	exeTime := endTime.Sub(startTime)
	return arr, exeTime
}

func QuickSort(arr []int) ([]int, time.Duration) {
	startTime := time.Now()
	if len(arr) < 2 {
		endTime := time.Now()
		exeTime := endTime.Sub(startTime)
		return arr, exeTime
	}
	left, right := 0, len(arr)-1
	pivot := rand.Int() % len(arr)
	arr[pivot], arr[right] = arr[right], arr[pivot]

	for i, _ := range arr {
		if arr[i] < arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}
	arr[left], arr[right] = arr[right], arr[left]
	QuickSort(arr[:left])
	QuickSort(arr[left+1:])
	endTime := time.Now()
	exeTime := endTime.Sub(startTime)
	return arr, exeTime
}

func SelectionSort(arr []int) ([]int, time.Duration) {
	startTime := time.Now()
	for i := 0; i < len(arr)-1; i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	endTime := time.Now()
	exeTime := endTime.Sub(startTime)
	return arr, exeTime
}

func InsertionSort(arr []int) ([]int, time.Duration) {
	startTime := time.Now()
	var n = len(arr)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
			j = j - 1
		}
	}
	endTime := time.Now()
	exeTime := endTime.Sub(startTime)
	return arr, exeTime
}

func ShellSort(arr []int) ([]int, time.Duration) {
	startTime := time.Now()
	n := len(arr)
	gap := n / 2
	for gap > 0 {
		for i := gap; i < n; i++ {
			if arr[i] < arr[i-gap] {
				arr[i], arr[i-gap] = arr[i-gap], arr[i]
			}
		}
		gap = gap / 2
	}
	endTime := time.Now()
	exeTime := endTime.Sub(startTime)
	return arr, exeTime
}

func CombSort(arr []int) ([]int, time.Duration) {
	startTime := time.Now()
	var (
		n       = len(arr)
		gap     = len(arr)
		shrink  = 1.2
		swapped = true
	)

	for swapped {
		swapped = false
		gap = int(float64(gap) / shrink)
		if gap < 1 {
			gap = 1
		}
		for i := 0; i+gap < n; i++ {
			if arr[i] > arr[i+gap] {
				arr[i+gap], arr[i] = arr[i], arr[i+gap]
				swapped = true
			}
		}
	}
	endTime := time.Now()
	exeTime := endTime.Sub(startTime)
	return arr, exeTime
}

func Merge(left, right []int) (result []int) {
	result = make([]int, len(left)+len(right))
	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}
	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}
	return
}

func MergeSort(arr []int) ([]int, time.Duration) {
	startTime := time.Now()
	var num = len(arr)

	if num == 1 {
		endTime := time.Now()
		exeTime := endTime.Sub(startTime)
		return arr, exeTime
	}

	middle := int(num / 2)
	var (
		left  = make([]int, middle)
		right = make([]int, num-middle)
	)
	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = arr[i]
		} else {
			right[i-middle] = arr[i]
		}
	}
	leftArr, _ := MergeSort(left)
	rightArr, _ := MergeSort(right)
	returnVal := Merge(leftArr, rightArr)
	endTime := time.Now()
	exeTime := endTime.Sub(startTime)
	return returnVal, exeTime
}

func RadixSort(arr []int) ([]int, time.Duration) {
	startTime := time.Now()
	const digit = 4
	const maxbit = -1 << 31
	buf := bytes.NewBuffer(nil)
	ds := make([][]byte, len(arr))
	for i, e := range arr {
		binary.Write(buf, binary.LittleEndian, e^maxbit)
		b := make([]byte, digit)
		buf.Read(b)
		ds[i] = b
	}
	countingSort := make([][][]byte, 256)
	for i := 0; i < digit; i++ {
		for _, b := range ds {
			countingSort[b[i]] = append(countingSort[b[i]], b)
		}
		j := 0
		for k, bs := range countingSort {
			copy(ds[j:], bs)
			j += len(bs)
			countingSort[k] = bs[:0]
		}
	}
	var w int
	for i, b := range ds {
		buf.Write(b)
		binary.Read(buf, binary.LittleEndian, &w)
		arr[i] = w ^ maxbit
	}
	endTime := time.Now()
	exeTime := endTime.Sub(startTime)
	return arr, exeTime
}

type arrList []int

func (arr arrList) Flip(r int) {
	for l := 0; l < r; l, r = l+1, r-1 {
		arr[l], arr[r] = arr[r], arr[l]
	}
}

func PancakeSort(arr arrList) ([]int, time.Duration) {
	startTime := time.Now()
	for uns := len(arr) - 1; uns > 0; uns-- {
		// find largest in unsorted range
		lx, lg := 0, arr[0]
		for i := 1; i <= uns; i++ {
			if arr[i] > lg {
				lx, lg = i, arr[i]
			}
		}
		// move to final position in two flips
		arr.Flip(lx)
		arr.Flip(uns)
	}
	endTime := time.Now()
	exeTime := endTime.Sub(startTime)
	return arr, exeTime
}
