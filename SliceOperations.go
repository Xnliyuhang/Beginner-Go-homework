package main

import (
	"fmt"
)

func main() {
	src := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //定义切片
	idx := 2                                    // 要删除的元素的索引
	vals, val, _ := SliceOperations(idx, src)
	fmt.Printf("vals:%v,val:%d", vals, val)
}

func SliceOperations[T any](idx int, src []T) ([]T, T, error) {
	length := len(src)
	//判断idx是否超出切片长度
	if idx < 0 || idx > length {
		var zero T
		return nil, zero, newErrIndexOutOfRange(length, idx)
	}

	//从切片中获取到idx的元素
	res := src[idx]

	//从idx位置开始,后面的元素依次往前挪1个位置
	for i := idx; i+1 < length; i++ {
		src[i] = src[i+1]
	}

	//去掉最后一个重复元素
	src = src[:length-1]
	src = Shrink(src)
	return src, res, nil
}

func newErrIndexOutOfRange(length int, idx int) error {
	return fmt.Errorf("ekit:下标超出范围，长度 %d，下标%d", length, idx)
}

func Shrink[T any](src []T) []T {
	c := cap(src)
	len := len(src)
	n, changed := calCapacity(c, len)
	if !changed {
		return src
	}
	s := make([]T, 0, n)
	s = append(s, src...)
	return s

}

func calCapacity(c int, len int) (int, bool) {
	//容量小于64,不用缩
	if c <= 64 {
		return c, false
	}

	//容量小于大于2048并且容量是长度的2倍以上,容量按照0.625的比例缩
	if c > 2048 && (c/len >= 2) {
		factor := 0.625
		return int(float32(c) * float32(factor)), true
	}

	//容量小于大于2048并且容量是长度的2倍以上,容量按照0.625的比例缩
	if c < 2048 && (c/len >= 4) {
		return int(c / 2), true
	}
	return c, false
}
