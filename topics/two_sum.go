package main
import "fmt"

// 两数之和，暴力破解版，O(n^2)
func twoSum(nums []int, target int) []int {
	for i, x := range nums {
		for j := i + 1; j < len(nums); j++ {
			if x+nums[j] == target {
				return []int{i, j}
			}
		}

	}
	return nil

}

// hash表去重，O(1)
func twoSum2(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}

		}
		hashTable[x] = i

	}
	return nil

}

func main() {
	[]int ret:= twoSum([]int{1, 2, 34, 5}, 6)
	fmt.
}

main()