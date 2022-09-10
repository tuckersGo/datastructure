package main

import "fmt"

func main() {
	students := []struct {
		Name   string
		Height float64
	}{
		{Name: "Kyle", Height: 173.4},
		{Name: "Ken", Height: 164.5},
		{Name: "Ryu", Height: 178.8},
		{Name: "Honda", Height: 154.2},
		{Name: "Hwarang", Height: 188.8},
		{Name: "Lebron", Height: 209.8},
		{Name: "Hodong", Height: 197.7},
		{Name: "Tom", Height: 164.8},
		{Name: "Kevin", Height: 164.8},
	}

	var heightMap [3000][]string
	for i := 0; i < len(students); i++ {
		idx := int(students[i].Height * 10)
		heightMap[idx] = append(heightMap[idx], students[i].Name)
	}

	// 140 ~ 170
	for i := 1400; i < 1700; i++ {
		for _, name := range heightMap[i] {
			fmt.Println("name:", name, "height:", float64(i)/10)
		}
	}

	// 180 ~ 210
	for i := 1800; i < 2100; i++ {
		for _, name := range heightMap[i] {
			fmt.Println("name:", name, "height:", float64(i)/10)
		}
	}
}
