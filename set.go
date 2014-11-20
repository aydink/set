package set

import "fmt"
import "math/rand"
import "time"

type Posting struct {
	docId             int
	documentFrequency float32
	boost             float32
}

func UnionInt(arr1, arr2 []int) []int {
	m := len(arr1)
	n := len(arr2)

	// Uninon of the two array at least will be the size of the longest. For performance we preallocate
	max := 0
	if m < n {
		max = n
	} else {
		max = m
	}

	fmt.Println(max)

	p := make([]int, 0, max)

	i, j := 0, 0

	for i < m && j < n {
		if arr1[i] < arr2[j] {
			p = append(p, arr1[i])
			//fmt.Printf(" %d ", arr1[i])
			i++
		} else if arr2[j] < arr1[i] {
			p = append(p, arr2[j])
			//fmt.Printf(" %d ", arr2[j])
			j++
		} else {
			p = append(p, arr2[j])
			//fmt.Printf(" %d ", arr2[j])
			j++
			i++
		}
	}

	/* Print remaining elements of the larger array */
	for i < m {
		p = append(p, arr1[i])
		//fmt.Printf(" %d ", arr1[i])
		i++
	}
	for j < n {
		p = append(p, arr2[j])
		//fmt.Printf(" %d ", arr2[j])
		j++
	}

	return p
}

func Union(arr1, arr2 []Posting) []Posting {
	m := len(arr1)
	n := len(arr2)

	// Uninon of the two array at least will be the size of the longest. For performance we preallocate
	max := 0
	if m < n {
		max = n
	} else {
		max = m
	}

	fmt.Println(max)

	p := make([]Posting, 0, max)

	i, j := 0, 0

	for i < m && j < n {
		if arr1[i].docId < arr2[j].docId {
			p = append(p, arr1[i])
			//fmt.Printf(" %d ", arr1[i])
			i++
		} else if arr2[j].docId < arr1[i].docId {
			p = append(p, arr2[j])
			//fmt.Printf(" %d ", arr2[j])
			j++
		} else {
			p = append(p, arr2[j])
			//fmt.Printf(" %d ", arr2[j])
			j++
			i++
		}
	}

	/* Print remaining elements of the larger array */
	for i < m {
		p = append(p, arr1[i])
		//fmt.Printf(" %d ", arr1[i])
		i++
	}
	for j < n {
		p = append(p, arr2[j])
		//fmt.Printf(" %d ", arr2[j])
		j++
	}

	return p
}

func IntersectionInt(arr1, arr2 []int) []int {
	m := len(arr1)
	n := len(arr2)

	min := 0
	if m < n {
		min = m
	} else {
		min = n
	}

	p := make([]int, 0, min/4)

	i, j := 0, 0

	for i < m && j < n {
		if arr1[i] < arr2[j] {
			i++
		} else if arr2[j] < arr1[i] {
			j++
		} else { /* if arr1[i] == arr2[j] */
			//fmt.Printf(" %d ", arr2[j])
			p = append(p, arr2[j])
			j++
			i++
		}
	}

	return p
}

func Intersection(arr1, arr2 []Posting) []Posting {
	m := len(arr1)
	n := len(arr2)

	min := 0
	if m < n {
		min = m
	} else {
		min = n
	}

	p := make([]Posting, 0, min/4)

	i, j := 0, 0

	for i < m && j < n {
		if arr1[i].docId < arr2[j].docId {
			i++
		} else if arr2[j].docId < arr1[i].docId {
			j++
		} else { /* if arr1[i] == arr2[j] */
			//fmt.Printf(" %d ", arr2[j])
			p = append(p, arr2[j])
			j++
			i++
		}
	}

	return p
}

func DifferenceInt(arr1, arr2 []int) []int {
	m := len(arr1)
	n := len(arr2)

	p := make([]int, 0, m/4)

	i, j := 0, 0

	for i < m && j < n {
		if arr1[i] < arr2[j] {
			p = append(p, arr1[i])
			i++
		} else if arr2[j] < arr1[i] {
			//p = append(p, arr2[j])
			j++
		} else { /* if arr1[i] == arr2[j] */
			//fmt.Printf(" %d ", arr2[j])
			//p = append(p, arr2[j])
			j++
			i++
		}
	}

	/* Print remaining elements of the larger array */
	for i < m {
		p = append(p, arr1[i])
		//fmt.Printf(" %d ", arr1[i])
		i++
	}

	return p
}

func Difference(arr1, arr2 []Posting) []Posting {
	m := len(arr1)
	n := len(arr2)

	p := make([]Posting, 0, m/4)

	i, j := 0, 0

	for i < m && j < n {
		if arr1[i].docId < arr2[j].docId {
			p = append(p, arr1[i])
			i++
		} else if arr2[j].docId < arr1[i].docId {
			//p = append(p, arr2[j])
			j++
		} else { /* if arr1[i] == arr2[j] */
			//fmt.Printf(" %d ", arr2[j])
			//p = append(p, arr2[j])
			j++
			i++
		}
	}

	/* Print remaining elements of the larger array */
	for i < m {
		p = append(p, arr1[i])
		//fmt.Printf(" %d ", arr1[i])
		i++
	}

	return p
}

func BinarySearch(p []int, value int) int {

	start_index := 0
	end_index := len(p) - 1

	for start_index <= end_index {

		median := (start_index + end_index) / 2

		if p[median] < value {
			start_index = median + 1
		} else {
			end_index = median - 1
		}

	}

	if start_index == len(p) || p[start_index] != value {
		return -1
	} else {
		return start_index
	}
}

func IntersectionBinarySearch(arr1, arr2 []int) []int {
	m := len(arr1)
	n := len(arr2)

	min := 0
	if m < n {
		min = m
	} else {
		min = n
	}

	p := make([]int, 0, min/4)

	for i := 0; i < m; i++ {
		val := BinarySearch(arr2, arr1[i])
		if val > -1 {
			p = append(p, val)
		}
	}

	return p
}

/* Driver program to test above function */
func main() {

	n := 1000000
	number := 0

	s1 := make([]int, 0, n)
	s2 := make([]int, 0, n)

	p1 := make([]Posting, 0, n)
	p2 := make([]Posting, 0, n)

	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < n; i++ {
		number += rand.Intn(1000) + 1
		s1 = append(s1, number)
		p1 = append(p1, Posting{number, 1, 1})
	}

	number = 0

	for i := 0; i < n; i++ {
		number += rand.Intn(100) + 1
		s2 = append(s2, number)
		p2 = append(p2, Posting{number, 1, 1})
	}

	//fmt.Println(s1)
	//fmt.Println(s2)

	difference := DifferenceInt(s1, s2)
	intersection := IntersectionInt(s1, s2)

	if len(difference)+len(intersection) == len(s1) {
		fmt.Println("difference works!")
	}

	start := time.Now()
	p := UnionInt(s1, s2)
	elapsed := time.Since(start)
	fmt.Printf("Union took %s\n", elapsed)

	fmt.Printf("len:%v\n", len(p))

	start = time.Now()
	p = IntersectionInt(s1, s2)
	elapsed = time.Since(start)
	fmt.Printf("Intersection took %s\n", elapsed)
	fmt.Printf("len:%v\n", len(p))

	start = time.Now()
	p = IntersectionBinarySearch(s1, s2)
	elapsed = time.Since(start)
	fmt.Printf("IntersectionBinarySearch took %s\n", elapsed)
	fmt.Printf("len:%v\n", len(p))

	start = time.Now()
	p = DifferenceInt(s1, s2)
	elapsed = time.Since(start)
	fmt.Printf("Difference took %s\n", elapsed)

	fmt.Printf("len:%v\n", len(p))
	//fmt.Println(s1)
	//fmt.Println(s2)
	//fmt.Println(p)

	/*	start = time.Now()

		set1 := set.NewNonTS() // thread safe version
		for _, v := range s1 {
			set1.Add(v)
		}

		set2 := set.NewNonTS() // thread safe version
		for _, v := range s2 {
			set2.Add(v)
		}

		elapsed = time.Since(start)
		fmt.Printf("Adding s1 to fatih.set took %s\n", elapsed)

		start = time.Now()
		set.Intersection(set1, set2)
		elapsed = time.Since(start)
		fmt.Printf("Set Union fatih.set took %s\n", elapsed)*/

	start = time.Now()
	res := Union(p1, p2)
	elapsed = time.Since(start)
	fmt.Printf("Intersection2 took %s\n", elapsed)
	fmt.Printf("len:%v\n", len(res))

}
