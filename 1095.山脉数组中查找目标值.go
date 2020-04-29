package main

import "fmt"

/*
 * @lc app=leetcode.cn id=1095 lang=golang
 *
 * [1095] 山脉数组中查找目标值
 */

type MountainArray []int

func (this *MountainArray) get(index int) int {
	// fmt.Println("call index", index)
	return (*this)[index]
}
func (this *MountainArray) length() int {
	return len(*this)
}

// @lc code=start
/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

func findInMountainArray(target int, mountainArr *MountainArray) int {
	leftind, rightind := 0, mountainArr.length()-1
	left, right := mountainArr.get(leftind), mountainArr.get(rightind)
	if left == target {
		return leftind
	}
	tmpind := binarySearchBoth(target, mountainArr, leftind, rightind, left, right)
	if tmpind != -1 {
		return tmpind
	} else if right == target {
		return rightind
	} else {
		return -1
	}
}

func binarySearchBoth(target int, mountainArr *MountainArray, leftind, rightind int, left, right int) int {
	//fmt.Println("both", leftind, rightind)
	if rightind-leftind <= 1 {
		return -1
	}
	midind := leftind + (rightind-leftind)/2
	mid := mountainArr.get(midind)
	leftres, rightres := -1, -1
	// left indexs
	if mid < right {
		if mid > target {
			leftres = binarySearchAsc(target, mountainArr, leftind, midind, left, mid)
		}
	} else {
		leftres = binarySearchBoth(target, mountainArr, leftind, midind, left, mid)
	}
	if leftres != -1 {
		return leftres
	}
	// mid
	if mid == target {
		return midind
	}
	// right
	if mid < left {
		if mid > target {
			rightres = binarySearchDesc(target, mountainArr, midind, rightind, mid, right)
		}
	} else {
		rightres = binarySearchBoth(target, mountainArr, midind, rightind, mid, right)
	}
	return rightres
}

func binarySearchAsc(target int, mountainArr *MountainArray, leftind, rightind int, left, right int) int {
	//fmt.Println("asc", leftind, rightind)
	if right-left <= 1 {
		return -1
	}
	midind := leftind + (rightind-leftind)/2
	mid := mountainArr.get(midind)
	if mid < target {
		return binarySearchAsc(target, mountainArr, midind, rightind, mid, right)
	} else if mid == target {
		return midind
	} else {
		return binarySearchAsc(target, mountainArr, leftind, midind, left, mid)
	}
}

func binarySearchDesc(target int, mountainArr *MountainArray, leftind, rightind int, left, right int) int {
	//fmt.Println("desc", leftind, rightind)
	if right-left <= 1 {
		return -1
	}
	midind := leftind + (rightind-leftind)/2
	mid := mountainArr.get(midind)
	if mid > target {
		return binarySearchDesc(target, mountainArr, midind, rightind, mid, right)
	} else if mid == target {
		return midind
	} else {
		return binarySearchDesc(target, mountainArr, leftind, midind, left, mid)
	}
}

// @lc code=end
func main() {
	fmt.Println(findInMountainArray(3, &MountainArray{1, 2, 3, 4, 5, 3, 1}))
	fmt.Println(findInMountainArray(3, &MountainArray{0, 1, 2, 4, 2, 1}))
	fmt.Println(findInMountainArray(3, &MountainArray{1, 2, 3, 5, 3}))
	// fmt.Println(findInMountainArray(21, &MountainArray{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2}))
	// fmt.Println(findInMountainArray(81, &MountainArray{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 100, 99, 98, 97, 96, 95, 94, 93, 92, 91, 90, 89, 88, 87, 86, 85, 84, 83, 82}))
	// fmt.Println(findInMountainArray(481, &MountainArray{1, 6, 11, 16, 21, 26, 31, 36, 41, 46, 51, 56, 61, 66, 71, 76, 81, 86, 91, 96, 101, 106, 111, 116, 121, 126, 131, 136, 141, 146, 151, 156, 161, 166, 171, 176, 181, 186, 191, 196, 201, 206, 211, 216, 221, 226, 231, 236, 241, 246, 251, 256, 261, 266, 271, 276, 281, 286, 291, 296, 301, 306, 311, 316, 321, 326, 331, 336, 341, 346, 351, 356, 361, 366, 371, 376, 381, 386, 391, 396, 401, 406, 411, 416, 421, 426, 431, 436, 441, 446, 451, 456, 461, 466, 471, 476, 481, 486, 491, 496, 501, 496, 491, 486, 481, 476, 471, 466, 461, 456, 451, 446, 441, 436, 431, 426, 421, 416, 411, 406})) //96
	fmt.Println(findInMountainArray(101, &MountainArray{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 100, 99, 98, 97, 96, 95, 94, 93, 92, 91, 90, 89, 88, 87, 86, 85, 84, 83, 82}))

}
