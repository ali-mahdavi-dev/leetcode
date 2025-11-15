func twoSum(nums []int, target int) []int {
    sumMap:=make(map[int]int)
    
    for i,v:=range nums{
        if ei, ok:=sumMap[target-v]; ok{
            return []int{ei, i}
        }
        sumMap[v]=i
    }

    return []int{}
}