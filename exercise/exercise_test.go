package exercise

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

func TestName1(t *testing.T) {
	num := 2
	height := []int{100, 95}

	stack := make([]int, num)
	fIndexs := make([]int, num)
	for i, v := range height {
		if len(stack) != 0 {
			for _, index := range stack {
				if v > height[index] {
					fIndexs[index] = i
					stack = stack[:len(stack)-1]
				}
			}
		}
		stack = append(stack, i)
	}

	for _, fi := range fIndexs {
		println(fi)
	}
}

func TestName(t *testing.T) {

	carStr := "1,1,0,0,1,1,1,0,1"
	cars := strings.Split(carStr, ",")
	total, sum := 0, 0
	for i, car := range cars {
		if car == "1" {
			sum++
		}

		if car == "0" && i != 0 && sum != 0 {
			total++
			sum = 0
		} else if sum > 3 {
			total++
			sum = 1
		}
	}
	if cars[len(cars)-1] == "1" {
		total++
	}

	fmt.Println(total)
}

func TestName2(t *testing.T) {
	//cite := 0
	num := 6
	route := make(map[int][]int)
	route2 := [][]int{
		{1, 2},
		{2, 3},
		{2, 5},
		{3, 4},
		{3, 6},
	}

	for i := 0; i < num-1; i++ {
		city1 := route2[i][0]
		if _, ok := route[city1]; !ok {
			route[city1] = []int{}
		}
		route[city1] = append(route[city1], route2[i][1])
		//fmt.Scan(&route[i][0], &route[i][1])
	}

	minDP := math.MaxInt
	DPSlice := []int{}
	for city, cCity := range route {
		cityDp := math.MinInt
		cCityCount := 0
		for _, i2 := range cCity {
			cCityDp := Deep(i2, route)
			if cCityDp > cityDp {
				cityDp = cCityDp
			}
			cCityCount += cCityDp
		}
		otherDP := num - cCityCount - 1
		if otherDP > cityDp && otherDP != 0 {
			cityDp = otherDP
		}

		if minDP > cityDp {
			DPSlice = []int{city}
			minDP = cityDp
		} else if minDP == cityDp {
			DPSlice = append(DPSlice, city)
		}
	}

	for _, i2 := range DPSlice {
		fmt.Println(i2)
	}

}

func Deep(city int, route map[int][]int) (coun int) {
	if _, ok := route[city]; !ok {
		return 1
	}
	for _, i2 := range route[city] {
		coun += Deep(i2, route)
	}
	return coun + 1
}
