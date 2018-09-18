# leetcode
go solutions of leetcode problems



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

