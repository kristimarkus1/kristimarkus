package piscine

func SortIntegerTable(table []int) {
	n := len(table)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if table[j] > table[j+1] {
				table[j], table[j+1] = table[j+1], table[j]
			}
		}
	}
}

// Write a function that reorders a slice of int in ascending order.

// s := []int{5,4,3,2,1,0}

// [0 1 2 3 4 5]
