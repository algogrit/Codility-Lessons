package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	carsGoingEast := 0
	// carsGoingWest := 0
	carsPassingEachOther := 0

	for _, el := range A {
		if el == 0 {
			carsGoingEast += 1

			// carsPassingEachOther += carsGoingWest
		} else {
			// carsGoingWest += 1

			carsPassingEachOther += carsGoingEast
		}

		if carsPassingEachOther > 1_000_000_000 {
			return -1
		}
	}

	return carsPassingEachOther
}
