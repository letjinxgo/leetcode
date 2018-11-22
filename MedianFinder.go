package main

import "fmt"

type MedianFinder struct {
	Nums []int
}

func Constructor() MedianFinder {
	return MedianFinder{[]int{}}
}

func (this *MedianFinder) AddNum(num int) {
	fmt.Println(this)
	if len(this.Nums) == 0 {
		this.Nums = append(this.Nums, num)
		return
	}
	if this.Nums[0] >= num {
		this.Nums = append([]int{num}, this.Nums...)
		return
	}
	if this.Nums[len(this.Nums)-1] <= num {
		this.Nums = append(this.Nums, num)
		return
	}
	l := 0
	r := len(this.Nums)
	var p int
	for l <= r {
		p = (l + r) / 2
		if this.Nums[p] >= num && this.Nums[p-1] <= num {
			record := this.Nums[p]
			temp := this.Nums[p:]
			this.Nums = append(this.Nums[:p], num)
			this.Nums = append(this.Nums, temp...)
			this.Nums[p+1] = record
			return
		}
		if this.Nums[p] > num {
			r = p - 1
		}
		if this.Nums[p] < num {
			l = p + 1
		}
	}

}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.Nums) == 0 {
		return 0
	}
	if len(this.Nums) == 1 {
		return float64(this.Nums[0])
	}
	l := len(this.Nums)
	if l%2 == 0 {
		return float64(this.Nums[l/2]+this.Nums[l/2-1]) / float64(2)
	}
	if l%2 == 1 {
		return float64(this.Nums[l/2])
	}
	return 0
}

func main() {
	MF := Constructor()
	MF.AddNum(6)
	fmt.Println(MF.FindMedian())
	MF.AddNum(10)
	fmt.Println(MF.FindMedian())
	MF.AddNum(2)
	fmt.Println(MF.FindMedian())
	MF.AddNum(6)
	fmt.Println(MF.FindMedian())
	MF.AddNum(5)
	fmt.Println(MF.FindMedian())
	MF.AddNum(0)
	fmt.Println(MF.FindMedian())
	MF.AddNum(6)
	fmt.Println(MF.FindMedian())
	MF.AddNum(3)
	fmt.Println(MF.FindMedian())
	MF.AddNum(1)
	fmt.Println(MF.FindMedian())
	MF.AddNum(0)
	fmt.Println(MF.FindMedian())
	MF.AddNum(0)
	fmt.Println(MF.FindMedian())
	MF.AddNum(5)
	fmt.Println(MF)
	MF.AddNum(4)
	fmt.Println(MF)
	MF.AddNum(3)
	fmt.Println(MF)
	MF.AddNum(2)
	fmt.Println(MF)
	MF.AddNum(1)
	MF.AddNum(2)
	MF.AddNum(8)
	MF.AddNum(7)
	MF.AddNum(6)
	MF.AddNum(5)
	fmt.Println(MF)

}
