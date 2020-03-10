package arithmetic

/**
* @ Description:
* @Author:
* @Date: 2020/2/23 17:43
 */
func TwoSum(nums []int, target int) []int {
	var (
		hash=map[int]int{}
	)
	for i,v:=range(nums){
		if _,ok:=hash[target-v];ok{
			return []int{hash[target-v],i}
		}else{
			hash[v]=i
		}
	}
	return []int{}
}
