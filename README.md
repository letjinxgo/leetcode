# leetcode
go solutions of leetcode problems

---

#### 1. Two Sum

Given an array of integers, return **indices** of the two numbers such that they add up to a specific target.

You may assume that each input would have **exactly** one solution, and you may not use the *same* element twice.

**Example:**

```
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
```

**[Code:](1.TwoSum.go)**

```go
/**
 * 1. TwoSum.go
 */
 func twoSum(nums []int, target int) []int {
	for i, x := range nums {
		for j, y := range nums {
			if x+y == target && j > i {
				nums = append([]int{}, i, j)
			}
		}
	}
	return nums
}
```

-----

#### 26. Remove Duplicates from Sorted Array

Given a sorted array *nums*, remove the duplicates [**in-place**](https://en.wikipedia.org/wiki/In-place_algorithm) such that each element appear only *once* and return the new length.

Do not allocate extra space for another array, you must do this by **modifying the input array in-place** with O(1) extra memory.

**Example 1:**

```
Given nums = [1,1,2],

Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively.

It doesn't matter what you leave beyond the returned length.
```

**Example 2:**

```
Given nums = [0,0,1,1,1,2,2,3,3,4],

Your function should return length = 5, with the first five elements of nums being modified to 0, 1, 2, 3, and 4 respectively.

It doesn't matter what values are set beyond the returned length.
```

**Clarification:**

Confused why the returned value is an integer but your answer is an array?

Note that the input array is passed in by **reference**, which means modification to the input array will be known to the caller as well.

Internally you can think of this:

```
// nums is passed in by reference. (i.e., without making a copy)
int len = removeDuplicates(nums);

// any modification to nums in your function would be known by the caller.
// using the length returned by your function, it prints the first len elements.
for (int i = 0; i < len; i++) {
    print(nums[i]);
}
```

**[Code:](26.RemoveDuplicatesfromSortedArray.go)**

```go
/**
 * 26.RemoveDuplicatesfromSortedArray.go
 */
func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return len(nums)
	}
	for i := 1; i <= len(nums)-1; i++ {
		if nums[i] == nums[i-1] {
			nums = append(nums[:i-1], nums[i:]...)
			if i > 0 {
				i--
			} else {
				i = 0
			}
		}
	}
	return len(nums)
}
```

**[Code:](26.RemoveDuplicatesfromNotSortedArray.go)**

```go
/**
 * 26.RemoveDuplicatesfromNotSortedArray.go
 */
func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return len(nums)
	}
	for i := 0; i <= len(nums)-1; i++ {
		for j := len(nums) - 1; j > i; j-- {
			if nums[i] == nums[j] {
				nums = append(nums[:i], nums[i+1:]...)
				if i > 0 {
					i--
				} else {
					i = 0
				}

			}
		}
	}
	return len(nums)
}
```

----

#### 66. Plus One

Given a **non-empty** array of digits representing a non-negative integer, plus one to the integer.

The digits are stored such that the most significant digit is at the head of the list, and each element in the array contain a single digit.

You may assume the integer does not contain any leading zero, except the number 0 itself.

**Example 1:**

```
Input: [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
```

**Example 2:**

```
Input: [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
```

**[Code:](66.PlusOne.go)**

```go
/**
 * 66.PlusOne.go
 */
func plusOne(digits []int) []int {
	first := 0
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] < 10 {
			break
		}
		digits[i] %= 10
		if i == 0 {
			first = 1
		}
	}
	if first == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
```

---

#### 118. Pascal's Triangle

Given a non-negative integer *numRows*, generate the first *numRows* of Pascal's triangle.

![img](PascalTriangleAnimated2.gif)
In Pascal's triangle, each number is the sum of the two numbers directly above it.

**Example:**

```
Input: 5
Output:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]
```

---

**[Code:](118.Pascal'sTriangle.go)**

```go
/**
 * 118.Pascal'sTriangle.go
 */
func generate(numRows int) [][]int {
	var result [][]int
	if numRows == 1 {
		result = [][]int{{1}}
	}
	if numRows == 2 {
		result = [][]int{{1}, {1, 1}}
	}
	if numRows >= 3 {
		pre := generate(numRows - 1)
		prelast := pre[len(pre)-1]
		var last []int = []int{1}
		for i, _ := range prelast {
			if i <= len(prelast)-2 {
				temp := prelast[i] + prelast[i+1]
				last = append(last, temp)
			}
		}
		last = append(last, 1)
		result = append(pre, last)
	}
	return result
}

```

---

#### 136. Single Number

Given a **non-empty** array of integers, every element appears *twice* except for one. Find that single one.

**Note:**

Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

**Example 1:**

```
Input: [2,2,1]
Output: 1
```

**Example 2:**

```
Input: [4,1,2,1,2]
Output: 4
```

**[Code:](136.SingleNumber.go)**

```go
/**
 * 136.SingleNumber.go
 */
 func singleNumber(nums []int) int {
	for _, v := range nums {
		counts := 0
		for _, m := range nums {
			if v != m {
				counts++
			}
		}
		if counts == len(nums)-1 {
			return v
		}
	}
	return 0
}
```

---

#### 137. Single Number II

Given a **non-empty** array of integers, every element appears *three* times except for one, which appears exactly once. Find that single one.

**Note:**

Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

**Example 1:**

```
Input: [2,2,3,2]
Output: 3
```

**Example 2:**

```
Input: [0,1,0,1,0,1,99]
Output: 99
```

**[Code:](137.SingleNumberII.go)**

```go
/**
 * 137.SingleNumberII.go
 */
 func singleNumber(nums []int) int {
	for _, v := range nums {
		counts := 0
		for _, m := range nums {
			if v != m {
				counts++
			}
		}
		if counts == len(nums)-1 {
			return v
		}
	}
	return 0
}
```

---

#### 226. Invert Binary Tree

Invert a binary tree.

**Example:**

Input:

```
     4
   /   \
  2     7
 / \   / \
1   3 6   9
```

Output:

```
     4
   /   \
  7     2
 / \   / \
9   6 3   1
```

**Trivia:**
This problem was inspired by [this original tweet](https://twitter.com/mxcl/status/608682016205344768) by [Max Howell](https://twitter.com/mxcl):

> Google: 90% of our engineers use the software you wrote (Homebrew), but you can’t invert a binary tree on a whiteboard so f*** off.

**[Code:](226.InvertBinaryTree.go)**

```go
/**
 * InvertBinaryTree.go
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root = &TreeNode{root.Val, invertTree(root.Right), invertTree(root.Left)}
	return root
}
```

---

#### 260. Single Number III

Given an array of numbers `nums`, in which exactly two elements appear only once and all the other elements appear exactly twice. Find the two elements that appear only once.

**Example:**

```
Input:  [1,2,1,3,2,5]
Output: [3,5]
```

**Note**:

1. The order of the result is not important. So in the above example, `[5, 3]` is also correct.
2. Your algorithm should run in linear runtime complexity. Could you implement it using only constant space complexity?

**[Code:](260.SingleNumberIII.go)**

```go
/**
 * 260.SingleNumberIII.go
 */
func singleNumber(nums []int) []int {
	var numss []int
	for _, v := range nums {
		counts := 0
		for _, m := range nums {
			if v != m {
				counts++
			}
		}
		if counts == len(nums)-1 {
			numss = append(numss, v)
		}
		if len(numss) == 2 {
			return numss
		}
	}
	return nil
}

```

---

#### 268. Missing Number

Given an array containing *n* distinct numbers taken from `0, 1, 2, ..., n`, find the one that is missing from the array.

**Example 1:**

```
Input: [3,0,1]
Output: 2
```

**Example 2:**

```
Input: [9,6,4,2,3,5,7,0,1]
Output: 8
```

**Note**:
Your algorithm should run in linear runtime complexity. Could you implement it using only constant extra space complexity?

**[Code:](./268.MissingNumber.go)**

```go
/**
 * 268.MissingNumber.go
 */
func missingNumber(nums []int) int {
	sort.Ints(nums)
	fmt.Println(nums)
	for i, v := range nums {
		if i != v {
			return i
		}
		if i == len(nums)-1 {
			return i + 1
		}
	}
	return 0
}
```

---

#### 326. Power of Three

Given an integer, write a function to determine if it is a power of three.

**Example 1:**

```
Input: 27
Output: true
```

**Example 2:**

```
Input: 0
Output: false
```

**Example 3:**

```
Input: 9
Output: true
```

**Example 4:**

```
Input: 45
Output: false
```

**Follow up:**
Could you do it without using any loop / recursion?

**[Code:](326.PowerofThree.go)**

```go
/**
 * 326.PowerofThree.go
 */
 func isPowerOfThree(n int) bool {
	if n > 0 && 1162261467%n == 0 {
		return true
	}
	return false
}
```

---

#### 344. Reverse String

Write a function that takes a string as input and returns the string reversed.

**Example 1:**

```
Input: "hello"
Output: "olleh"
```

**Example 2:**

```
Input: "A man, a plan, a canal: Panama"
Output: "amanaP :lanac a ,nalp a ,nam A"
```

**[Code:](344.ReverseString.go)**

```go
func reverseString(s string) string {
	runes := []byte(s)
	for from, to := 0, len(s)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}
```

---

#### 347. Top K Frequent Elements

Given a non-empty array of integers, return the **k** most frequent elements.

**Example 1:**

```
Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]
```

**Example 2:**

```
Input: nums = [1], k = 1
Output: [1]
```

**Note:**

- You may assume *k* is always valid, 1 ≤ *k* ≤ number of unique elements.
- Your algorithm's time complexity **must be** better than O(*n* log *n*), where *n* is the array's size.

**[Code:](./347.TopKFrequentElements.go)**

```go
/**
 * 347.TopKFrequentElements.go
 */
func topKFrequent(nums []int, k int) []int {
	var numcounts = make(map[int]int)
	var counts []int
	var result []int
	for _, v := range nums {
		numcounts[v]++
	}
	for _, v := range numcounts {
		counts = append(counts, v)
	}
	sort.Ints(counts)
	c := counts[len(counts)-k]
	for k, v := range numcounts {
		if v >= c {
			result = append(result, k)
		}
	}
	return result
}
```

---

#### 389. Find the Difference

Given two strings **s** and **t** which consist of only lowercase letters.

String **t** is generated by random shuffling string **s** and then add one more letter at a random position.

Find the letter that was added in **t**.

**Example:**

```
Input:
s = "abcd"
t = "abcde"

Output:
e

Explanation:
'e' is the letter that was added.
```

**[Code:](389.FindtheDifference.go)**

```go
/**
 * 389.FindtheDifference.go
 */
func findTheDifference(s string, t string) byte {
	var ss = make([]byte, len(s))
	var tt = make([]byte, len(t))
	for i := 0; i < len(s); i++ {
		ss[i] = s[i]
	}
	for i := 0; i < len(t); i++ {
		tt[i] = t[i]
	}
	for _, t := range tt {
		countss := 0
		counttt := 0
		for _, s := range tt {
			if t == s {
				counttt++
			}
		}
		for _, s := range ss {
			if t == s {
				countss++
			}
		}
		if countss < counttt {
			return t
		}
	}
	return ' '
}
```



#### 397. Integer Replacement

Given a positive integer *n* and you can do operations as follow:



1. If *n* is even, replace *n* with `*n*/2`.
2. If *n* is odd, you can replace *n* with either `*n* + 1` or `*n* - 1`.



What is the minimum number of replacements needed for *n* to become 1?



**Example 1:**

```
Input:
8

Output:
3

Explanation:
8 -> 4 -> 2 -> 1
```



**Example 2:**

```
Input:
7

Output:
4

Explanation:
7 -> 8 -> 4 -> 2 -> 1
or
7 -> 6 -> 3 -> 2 -> 1
```

**[Code:](./397.IntegerReplacement.go)**

```go
/**
 * 397.IntegerReplacement.go
 */
func integerReplacement(n int) int {
	fmt.Println(n)
	if n == 1 {
		return 0
	}
	if n%2 == 0 {
		return integerReplacement(n/2) + 1
	} else {
		if n == math.MaxInt32 {
			return integerReplacement(n - 1)
		}
		return int(math.Min(float64(integerReplacement(n-1)+1), float64(integerReplacement(n+1)+1)))
	}
}
```

---

