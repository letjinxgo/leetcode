package main

import (
	"fmt"
	"math"
	"net/http"
	"sort"
)

type Mux struct{}

func SayHi(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello " + r.URL.Path))
}
func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		SayHi(w, r)
		return
	}
	http.NotFound(w, r)
}
func hasAlternatingBits(n int) bool {
	fmt.Println(n)
	if n <= 2 {
		return true
	}

	curr := (n % 2) ^ 1
	for n > 0 {
		next := n % 2
		if next == curr {
			return false
		}
		curr = next
		n /= 2
	}
	return true
}

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

func findNumberOfLIS(nums []int) int {
	counts := []int{}
	count := 0
	Count := 0
	i := 0
	for i <= len(nums)-2 {
		if nums[i] < nums[i+1] {
			count++
		}
		if nums[i] >= nums[i+1] || i == len(nums)-2 {
			counts = append(counts, count)
			Count++
			count = 0
		}
		i++
	}
	sort.Ints(counts)
	return Count
}

// func main() {
// 	// nums := []int{2, 3, 4, 4, 1, 2}
// 	// fmt.Println(findNumberOfLIS(nums))
// 	fmt.Println(countNumbersWithUniqueDigits(3))
// }
func countNumbersWithUniqueDigits(n int) int {
	counts := 1
	for n > 0 {
		counts *= (10 - n + 1)
		n--
	}
	return counts + 1
}
func largeGroupPositions(S string) [][]int {
	result := [][]int{}
	record := []int{}
	pos := 0
	count := 0
	for i := 0; i < len(S)-1; i++ {
		if S[i+1] != S[i] {
			result = append(result, []int{pos, pos + count})
			record = append(record, count)
			pos = i
			count = 0
		}
		count++
	}
	fmt.Println(result, record)
	return result
}

func thirdMax(nums []int) int {
	sort.Ints(nums)
	if len(nums) < 3 {
		return nums[len(nums)-1]
	}
	max := nums[len(nums)-1]
	counts := 1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < max {
			max = nums[i]
			counts++
		}
		if counts == 3 {
			return max
		}
	}
	fmt.Println(counts)
	if counts < 3 {
		return nums[len(nums)-1]
	}
	return 0
}
func setZeroes(matrix [][]int) {
	records := [][]int{}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				records = append(records, []int{i, j})
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			for _, r := range records {
				if i == r[0] {
					matrix[i][j] = 0
				}
				if j == r[1] {
					matrix[i][j] = 0
				}
			}
		}
	}
}

// func main() {
// 	matrix := [][]int{[]int{0, 1, 2, 0}, []int{3, 4, 5, 2}, []int{1, 3, 1, 5}}
// 	setZeroes(matrix)
// }
// /**
//  * Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	record := head
	vals := []int{}
	for head != nil {
		vals = append(vals, head.Val)
		head = head.Next
	}
	sort.Ints(vals)
	head = record
	counts := 0
	for head != nil {
		head.Val = vals[counts]
		counts++
		head = head.Next
	}
	return record
}
func trailingZeroes(n int) int {
	c := 0
	for n/5 != 0 {
		n /= 5
		c += n
	}
	return c
}

func preimageSizeFZF(K int) int {
	counts := 0
	for i := 0; i < math.MaxInt64; i++ {
		if trailingZeroes(i) == K {
			fmt.Println(i)
			counts++
		}
		if trailingZeroes(i) > K {
			break
		}
	}
	return counts
}
func findPeakElement(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	for index := 1; index < len(nums)-1; index++ {
		if nums[index] > nums[index+1] && nums[index] > nums[index-1] {
			return index
		}
	}
	return -1
}

func nextGreaterElement(findNums []int, nums []int) []int {
	result := []int{}
label:
	for _, f := range findNums {
		pos := math.MaxInt64
		for i := 0; i < len(nums); i++ {
			if f == nums[i] {
				pos = i
			}
			if nums[i] > f && i > pos {
				result = append(result, nums[i])
				continue label
			}
		}
		result = append(result, -1)
	}
	return result
}
func main() {
	nums1 := []int{1, 3, 4, 5}
	nums2 := []int{1, 2, 3, 4}
	fmt.Println(nextGreaterElement(nums1, nums2))
}
