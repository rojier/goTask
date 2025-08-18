package main

func main() {

	nums := []int{1, 2, 3, 1, 3}
	println(singleNumber(nums))

}

// 给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素
func singleNumber(nums []int) int {

	times := make(map[int]int)
	for _, v := range nums {
		i := times[v]
		times[v] = i + 1

	}
	for k, v := range times {
		println("k==", k)
		println("v==", v)
		if v == 1 {
			return k
		}
	}

	return -1
}
