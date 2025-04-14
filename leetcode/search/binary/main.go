package main

import "log"

func binary_search(arr []int, key, imin, imax int) int {
	if imax < imin {
		return 0
	}
	imid := imin + (imax - imin) / 2
	if arr[imid] > key {
		return binary_search(arr, key, imin, imid-1)
	} else if arr[imid] < key {
		return binary_search(arr, key, imid+1, imax)
	} else {
		return imid
	}
}

func main() {
	log.Printf("%d", binary_search([]int{1,2,3,20,30,46}, 20, 0, 5))
}