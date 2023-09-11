package main

func InsertionSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}

	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && key < arr[j] {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = key
	}

  return arr
}

func main() {
	// for i in 0..n-1
	//   j = i
	//   while j > 0 and a[j-1] < a[j]
	//     swap(a[j-1], a[j])
	//     j--
}
