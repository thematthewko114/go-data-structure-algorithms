package searchingAlgo

import (
	"time"
)

func LinearSearch(arr []int, key int) (bool, time.Duration) {
	startTime := time.Now()
	for _, item := range arr {
		if item == key {
			endTime := time.Now()
			exeTime := endTime.Sub(startTime)
			return true, exeTime
		}
	}
	endTime := time.Now()
	exeTime := endTime.Sub(startTime)
	return false, exeTime
}

func BinarySearch(arr []int, key int) (bool, time.Duration) {
	startTime := time.Now()
	low := 0
	high := len(arr) - 1

	for low <= high {
		median := (low + high) / 2

		if arr[median] < key {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(arr) || arr[low] != key {
		endTime := time.Now()
		exeTime := endTime.Sub(startTime)
		return false, exeTime
	}
	endTime := time.Now()
	exeTime := endTime.Sub(startTime)
	return true, exeTime
}

func InterpolationSearch(array []int, key int) (bool, time.Duration) {
	startTime := time.Now()
	min, max := array[0], array[len(array)-1]

	low, high := 0, len(array)-1

	for {
		if key < min || key > max {
			endTime := time.Now()
			exeTime := endTime.Sub(startTime)
			return false, exeTime
		}

		// make a guess of the location
		var guess int
		if high == low {
			guess = high
		} else {
			size := high - low
			offset := int(float64(size-1) * (float64(key-min) / float64(max-min)))
			guess = low + offset
		}

		// maybe we found it?
		if array[guess] == key {
			// scan backwards for start of value range
			for guess > 0 && array[guess-1] == key {
				guess--
			}
			endTime := time.Now()
			exeTime := endTime.Sub(startTime)
			return true, exeTime
		}

		// if we guessed to high, guess lower or vice versa
		if array[guess] > key {
			high = guess - 1
			max = array[high]
		} else {
			low = guess + 1
			min = array[low]
		}
	}
}
