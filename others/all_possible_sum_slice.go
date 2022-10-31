package main

import (
	"fmt"
	"sort"
	"strconv"
)

func sum(nums ...float64) float64 {
	var out float64

	for _, el := range nums {
		out += el
	}

	return out
}

func getSliceCopy(currentSlice []float64) []float64 {
	newSlice := make([]float64, len(currentSlice))
	copy(newSlice, currentSlice)

	return newSlice
}

func getSliceCopyWithEl(currentSlice []float64, el float64) []float64 {
	newSlice := make([]float64, len(currentSlice), len(currentSlice)+1)
	copy(newSlice, currentSlice)
	newSlice = append(newSlice, el)

	return newSlice
}

func getSliceCopyWithoutEl(in []float64, elToRemove float64) []float64 {
	out := make([]float64, 0, len(in)-1)

	var isDeleted bool

	for _, el := range in {
		if !isDeleted && el == elToRemove {
			isDeleted = true
			continue
		}

		out = append(out, el)
	}

	return out
}

func sortSlice(in []float64) []float64 {
	sortableIn := getSliceCopy(in)
	sort.SliceStable(sortableIn, func(i, j int) bool { return sortableIn[i] < sortableIn[j] })

	return sortableIn
}

func getKey(in []float64) string {
	sortedIn := sortSlice(in)

	key := ""

	for _, el := range sortedIn {
		key += ":" + strconv.Itoa(int(el))
	}

	return key
}

type Memo map[string][][]float64

func getSubSlice(memo Memo, currentSlice, remainingInputs []float64, expectedSum float64) [][]float64 {
	if val, ok := memo[getKey(currentSlice)]; ok {
		return val
	}

	if sum(currentSlice...) == expectedSum {
		out := [][]float64{currentSlice}

		fmt.Println(currentSlice)

		memo[getKey(currentSlice)] = out

		return out
	}

	if (sum(currentSlice...) > expectedSum) && len(remainingInputs) == 0 {
		memo[getKey(currentSlice)] = nil

		return nil
	}

	var out [][]float64

	for _, el := range remainingInputs {
		newSlice := append(currentSlice, el)

		subSlice := getSubSlice(memo, newSlice, getSliceCopyWithoutEl(remainingInputs, el), expectedSum)

		if subSlice != nil {
			out = append(out, subSlice...)
		}
	}

	memo[getKey(currentSlice)] = out

	return out
}

func getSumSlices(inputs []float64, expectedSum float64) [][]float64 {
	memo := Memo{}
	return getSubSlice(memo, make([]float64, 0, len(inputs)), inputs, expectedSum)
}

func main() {
	inputs := []float64{91.0, 63.0, 1.0, 1.0, 1.0, 2.0, 2.0, 2.0, 14.0, 13.0, 13.0, 44.0, 123.0, -2.0, -63.0, -68.0, -200.0, -122.0}
	expectedSum := 35.0

	fmt.Println(getSumSlices(inputs, expectedSum))
}
