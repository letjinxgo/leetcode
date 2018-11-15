package main

import (
	"fmt"
	"sort"
	"time"
)

// type Mux struct{}

// func SayHi(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Hello " + r.URL.Path))
// }
// func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path == "/" {
// 		SayHi(w, r)
// 		return
// 	}
// 	http.NotFound(w, r)
// }
// func main() {
// 	mux := &Mux{}
// 	http.ListenAndServe(":8080", mux)
// }
// func hasAlternatingBits(n int) bool {
// 	fmt.Println(n)
// 	if n <= 2 {
// 		return true
// 	}

// 	curr := (n % 2) ^ 1
// 	for n > 0 {
// 		next := n % 2
// 		if next == curr {
// 			return false
// 		}
// 		curr = next
// 		n /= 2
// 	}
// 	return true
// }
// func kthSmallest(root *TreeNode, k int) int {
// 	result := []int{}
// 	result = append(result, root.va)
// }
// func main() {
// 	fmt.Println(hasAlternatingBits(math.MaxInt64))
// 	fmt.Println(math.MaxInt64)
// }
type haha []byte

func (p haha) Len() int           { return len(p) }
func (p haha) Less(i, j int) bool { return p[i] < p[j] }
func (p haha) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func nextGreatestLetter(letters []byte, target byte) byte {
	ha := haha(letters)
	sort.Sort(ha)
	for _, v := range ha {
		if v > target {
			return v
		}
	}
	return ha[0]
}
func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		fmt.Println(l, r)
		curr := (l + r) / 2
		if nums[curr] == target {
			return curr
		}
		if nums[curr] > target {
			r = curr - 1
		}
		if nums[curr] < target {
			l = curr + 1
		}
	}
	return -1
}
func main() {
	start := time.Now()
	fmt.Println(climbStairs(45))
	fmt.Println(time.Since(start))
}
func climbStairs(n int) int {
	if n <= 0 {
		return 0
	}
	nums := []int{1, 1}
	for n > 0 {
		nums[1], nums[0] = nums[1]+nums[0], nums[1]
		n--
	}
	return nums[0]
}
