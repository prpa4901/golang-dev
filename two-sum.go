func twoSum(nums []int, target int) []int {
    hmap := make(map[int]int)
    for index, elements := range nums {
        diff := target - elements
        if value, exists := hmap[diff]; exists {
            return []int{value, index}
        } else {
            hmap[elements] = index
        }
    }  
}