package main

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	halfIndex := len(arr) / 2
	a := MergeSort(arr[:halfIndex])
	b := MergeSort(arr[halfIndex:])

	return Merge(a, b)
}

func Merge(a []int, b []int) []int {
	n := len(a)
	m := len(b)
	mergedArray := make([]int, n+m)

	i := 0
	j := 0
	k := 0
	for i < n || j < m {
		if j == m || (i < n && a[i] < b[j]) {
			mergedArray[k] = a[i]
			k++
			i++
		} else {
			mergedArray[k] = b[j]
			k++
			j++
		}
	}

	return mergedArray
}

func main() {
	// merge(a,b)
	//  a_n = len(a)
	//  b_m = len(b)
	//  i = 0, j = 0, k = 0, c = [a_n + b_m]int
	//  while i < a_n or j < b_m:
	//    if j == b_m or (i < a_n and a[i] < b[j]):
	//      c[k++] = a[i++]
	//    else:
	//      c[k++] = b[j++]
	//  return c

	// sort(c):
	//  if len(c) < 2: return c
	//  a = c[0..n/2-1]
	//  b = c[n/2..n-1]
	//  a = sort(a)
	//  b = sort(b)
	//  return merge(a, b)
}
