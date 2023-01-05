package main

import (
	"awesomeProject/db"
	"awesomeProject/easy"
	"fmt"
)

func main() {

	//fmt.Println("hello world")
	//
	//struct_demo.TestStruct()
	//nums := make([]int, 3)
	//nums[0] = 3
	//nums[1] = 2
	//nums[2] = 4
	//var target = 6
	//sum := easy.TwoSum(nums, target)
	//
	//fmt.Println(sum)

	//rewardList := getRewardList()
	//fmt.Println(len(rewardList))
	//fmt.Println(cap(rewardList))
	//var sum = 0
	//for i := range rewardList {
	//	sum += rewardList[i]
	//}
	//fmt.Println(sum)

	// 计算乘风破浪的分数
	//getTotalScore(27, 184)
	//getRealGradeAndScore(6334)
	//getRealGradeAndScore(8000)

	//var s = "{0,199|427008},{200,399|427009},{400,600|427010},{600,999|427011},{1000,9999999|427012}"
	//re := regexp.MustCompile("[{},\\|]")

	//split := re.Split(s, -1)
	//set := []string{}
	//
	//for i := range split {
	//	set = append(set, split[i])
	//}

	//fmt.Println(set)
	//arr := re.Split(s, -1)
	//for i := 0; i < len(arr); i++ {
	//	fmt.Println(arr[i])
	//}

	//dateStr := utils.GetDateStr(1581292800000)
	//fmt.Println(time.Now())
	//fmt.Println(time.Now().Unix())
	//fmt.Println(time.Now().UnixMilli())
	//fmt.Println(time.Now().UTC())
	//fmt.Println(time.Now().Format(time.Layout))
	//fmt.Println(time.Now().Format(time.ANSIC))
	//fmt.Println(time.Now().Format(time.RFC3339))
	//fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	//var l1 easy.ListNode
	//l1.Val = 1
	//var l2 easy.ListNode
	//l2.Val = 2
	//l1.Next = &l2
	//var l3 easy.ListNode
	//l3.Val = 3
	//l2.Next = &l3
	//
	//
	//var j1 easy.ListNode
	////头部结构体
	//j1.Val = 1
	//var j2 easy.ListNode
	//j2.Val = 3
	//j1.Next = &j2
	//var j3 easy.ListNode
	//j3.Val = 4
	//j2.Next = &j3
	//
	//Req(&l1)
	//fmt.Println("--------------------------")
	//Req(&j1)
	//
	//lists := easy.MergeTwoLists(&l1, &j1)
	//fmt.Println("--------------------------")
	//Req(lists)

	//db.TestConnection()
	db.TestMysql()
}

func Req(tmp *easy.ListNode) {        //tmp指针是指向下一个结构体的地址，加*就是下一个结构体
	for tmp != nil {            //遍历输出链表中每个结构体，判断是否为空
		fmt.Println(*tmp)
		tmp = tmp.Next          //tmp变更为下一个结构体地址
	}
}

func getRewardList() []int {
	rewardList := make([]int, 30)
	for i := 1; i <= 30; i++ {
		idx := i - 1
		if i < 11 {
			rewardList[idx] = i * 10
		} else if i < 21 {
			rewardList[idx] = (i - 5) * 20
		} else {
			rewardList[idx] = (i - 14) * 50
		}
	}
	return rewardList
}

func getTotalScore(grade int, score int) int {
	rewardList := getRewardList()
	var totalScore = score
	for i := 1; i <= grade; i++ {
		if i >= len(rewardList) {
			totalScore += rewardList[len(rewardList)-1]
		} else {
			totalScore += rewardList[i-1]
		}
	}
	fmt.Println(totalScore)
	return totalScore
}

func getRealGradeAndScore(totalScore int) {
	var score = totalScore
	rewardList := getRewardList()
	var grade = 0

	var sum = 0
	for i2 := range rewardList {
		sum += i2
	}

	for i := 1; i <= len(rewardList); i++ {
		var point = 0
		if grade >= i {
			point = rewardList[len(rewardList)-1]
		} else {
			point = rewardList[i-1]
		}
		if score > point {
			score -= point
			grade++
		}
	}
	var point = rewardList[len(rewardList)-1]

	if score > point {
		var n = score / point
		if score%point > 0 {
			n = n + 1
		}
		for j := 1; j < n; j++ {
			score -= point
			grade++
		}
	}

	fmt.Println("grade - score :", grade, score)
}
