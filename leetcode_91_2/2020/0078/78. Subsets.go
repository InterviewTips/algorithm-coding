package _078

import "fmt"

//public List<List<Integer>> subsets(int[] nums) {
//    List<List<Integer>> res = new ArrayList<>();
//    backtrack(res, new ArrayList<>(), nums, 0);
//    return res;
//}
//
//void backtrack(List<List<Integer>> res, List<Integer> item, int[] nums, int index) {
//    res.add(new ArrayList<>(item));
//    for (int i = index; i < nums.length; i++) {
//        item.add(nums[i]);
//        backtrack(res, item, nums, i + 1);
//        item.remove(item.size() - 1);
//    }
//}
func subsets(nums []int) [][]int {
	var res [][]int
	backTrack(&res, []int{}, nums, 0)
	return res
}

func backTrack(res *[][]int, item []int, nums []int, index int) {
	var s []int
	for i := 0; i < len(item); i++ {
		s = append(s, item[i])
	}
	*res = append(*res, s)
	for i := index; i < len(nums); i++ {
		item = append(item, nums[i])
		fmt.Printf("item is %v\n", item)
		backTrack(res, item, nums, i+1)
		item = item[:len(item)-1]
	}
}
