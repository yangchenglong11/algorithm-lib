/*
 * MIT License
 *
 * Copyright (c) 2017 SmartestEE Co., Ltd..
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

/*
 * Revision History:
 *     Initial: 2017/11/02        Yang Chenglong
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	arr01 := []int{34, 45, 3, 6, 76, 34, 46, 809, 92, 82432,3344, 4543, 33, 6432, 76432, 3234, 426, 80439, 4392, 84}

	fmt.Print("排序前")
	fmt.Println(arr01)

	now:=time.Now()

	mergeSort(arr01, 0, len(arr01)-1)
	finish:=time.Since(now)
	fmt.Print("排序后")
	fmt.Println(arr01)
	fmt.Println(finish)
}

func merge1(arr []int, low, mid, high int) {
	leftLen := mid - low + 1
	rightLen := high - mid

	arrLeft := make([]int, leftLen+1)
	for i := 0; i < leftLen; i++ {
		arrLeft[i] = arr[low+i]
	}
	arrLeft[leftLen] = 99999 //哨兵牌

	arrRight := make([]int, rightLen+1)
	for j := 0; j < rightLen; j++ {
		arrRight[j] = arr[mid+j+1]
	}
	arrRight[rightLen] = 99999 //哨兵牌

	i, j := 0, 0
	for k := low; k <= high; k++ {
		if arrLeft[i] <= arrRight[j] {
			arr[k] = arrLeft[i]
			i++
		} else {
			arr[k] = arrRight[j]
			j++
		}
	}
}

func mergeSort(arr []int, low, high int) {
	if low < high {
		mid := (low + high) / 2
		mergeSort(arr, low, mid)
		mergeSort(arr, mid+1, high)
		merge2(arr, low, mid, high)
	}
}

func merge2(arr []int, low, mid, high int) {
	leftLen := mid - low + 1
	rightLen := high - mid

	arrLeft := make([]int, leftLen)
	for i := 0; i < leftLen; i++ {
		arrLeft[i] = arr[low+i]
	}

	arrRight := make([]int, rightLen)
	for j := 0; j < rightLen; j++ {
		arrRight[j] = arr[mid+j+1]
	}

	i, j, k := 0, 0, low
	for ; k <= high && i < leftLen && j < rightLen; k++ {
		if arrLeft[i] <= arrRight[j] {
			arr[k] = arrLeft[i]
			i++
		} else {
			arr[k] = arrRight[j]
			j++
		}
	}

	for ; i < leftLen && k <= high; k++ {
		arr[k] = arrLeft[i]
		i++
	}

	for ; j < rightLen && k <= high; k++ {
		arr[k] = arrRight[j]
		j++
	}
}
