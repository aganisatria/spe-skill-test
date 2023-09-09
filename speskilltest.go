package main

import (
	"fmt"
	"math"
)

type SpeSkillTest struct {
	
}

func (s *SpeSkillTest) Narcissistic(inputNumb int) bool{
	
	strNumb := fmt.Sprint(inputNumb)
	resNumb := 0

	for _, v := range strNumb{
		digit := int(v - '0')
		resNumb += int(math.Pow(float64(digit), float64(len(strNumb))))
	}

	return resNumb == inputNumb
}

func (s *SpeSkillTest) ParityOutlier(inputArr []int) interface{}{
	arrOdd := []int{}
	arrEven := []int{}

	for _, v := range inputArr{
		if len(arrEven) > 1 && len(arrOdd) > 1 {
			return false
		}
		if v % 2 == 0{
			if len(arrEven) > 1 {
				continue
			}
			arrEven = append(arrEven, v)
		} else {
			if len(arrOdd) > 1{
				continue
			}
			arrOdd = append(arrOdd, v)
		}
	}

	if (len(arrEven) > 1 && len(arrOdd) > 1)  {
		return false
	} else if len(arrOdd) > 1 {
		return arrEven[0]
	} else if len(arrEven) > 1{
		return arrOdd[0]
	} else{
		return false
	}

}

func (s *SpeSkillTest) FindNeedle(haystack []string, needle string) int{
	for i, v := range haystack{
		if v == needle{
			return i
		}
	}

	return 9999
}

func (s *SpeSkillTest) BlueOcean(arrInt1 []int, arrInt2 []int) (res []int) {
    arr2Map := make(map[int]bool)
    for _, v := range arrInt2 {
        arr2Map[v] = true
    }

    for _, v := range arrInt1 {
        if !arr2Map[v] {
            res = append(res, v)
        }
    }
	return res
}

func main() {
	spe := SpeSkillTest{
	}

	fmt.Println(spe.Narcissistic(153))
	fmt.Println(spe.Narcissistic(1634))
	fmt.Println(spe.Narcissistic(111))

	
	fmt.Println(spe.ParityOutlier([]int{2, 4, 0, 100, 4, 11, 2602, 36}))
	fmt.Println(spe.ParityOutlier([]int{160, 3, 1719, 19, 11, 13, -21}))
	// fmt.Println(spe.ParityOutlier([]int{11, 13, 15, 19, 9, 13, -21}))

	fmt.Println(spe.FindNeedle([]string{"red", "blue", "yellow", "black", "grey"}, "blue"))

	x1 := []int{1, 2, 3, 4, 6, 10}
    x2 := []int{1}
    blueoc := spe.BlueOcean(x1, x2)
    fmt.Println(blueoc) 
}