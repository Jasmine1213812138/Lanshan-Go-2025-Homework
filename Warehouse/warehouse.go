package Warehouse

import "fmt"

type Warehouse struct {
	Parsnip    int
	Greenbeans int
	Potato     int
	Tulip      int
	Garlic     int
}
type Price struct {
	Parsnip    int
	Greenbeans int
	Potato     int
	Tulip      int
	Garlic     int
}

// 记账系统
func Income(x string, in int) map[string]int {
	income := map[string]int{
		"Parsnip":    0,
		"Greenbeans": 0,
		"Potato":     0,
		"Tulip":      0,
		"Garlic":     0,
	}
	income[x] += in
	return income
}

// 加减仓库
func (m *Warehouse) Add(product string, account int) int {
	switch product {
	case "Parsnip":
		m.Parsnip += account
		return m.Parsnip
	case "Potato":
		m.Potato += account
		return m.Potato
	case "Tulip":
		m.Tulip += account
		return m.Tulip
	case "Garlic":
		m.Garlic += account
		return m.Garlic
	default:
		return 0
	}
}
func (m *Warehouse) Delete(product string, account int) int {
	switch product {
	case "Parsnip":
		m.Parsnip -= account
		return m.Parsnip
	case "Potato":
		m.Potato -= account
		return m.Potato
	case "Tulip":
		m.Tulip -= account
		return m.Tulip
	case "Garlic":
		m.Garlic -= account
		return m.Garlic
	default:
		return 0
	}
}

// 判断库存
func Judge(want string, account int, left Warehouse, price Price, judge int) (int, int) {
	switch want {
	case "Parsnip":
		if account > left.Parsnip {
			fmt.Println("不好意思小姐，没有那么多")
			judge = 1
			return judge, 0
		} else {
			fmt.Printf("总共是%d元", account*price.Parsnip)
			return judge, price.Parsnip * account
		}
	case "Greenbeans":
		if account > left.Greenbeans {
			fmt.Println("不好意思小姐，没有那么多")
			judge = 1
			return judge, 0
		} else {
			fmt.Printf("总共是%d元", account*price.Greenbeans)
			return judge, price.Greenbeans * account
		}
	case "Potato":
		if account > left.Potato {
			fmt.Println("不好意思小姐，没有那么多")
			judge = 1
			return judge, 0
		} else {
			fmt.Printf("总共是%d元", account*price.Potato)
			return judge, price.Potato * account
		}
	case "Tulip":
		if account > left.Tulip {
			fmt.Println("不好意思小姐，没有那么多")
			judge = 1
			return judge, 0
		} else {
			fmt.Printf("总共是%d元", account*price.Tulip)
			return judge, price.Tulip * account
		}
	case "Garlic":
		if account > left.Garlic {
			fmt.Println("不好意思小姐，没有那么多")
			judge = 1
			return judge, 0
		} else {
			fmt.Printf("总共是%d元", account*price.Garlic)
			return judge, price.Garlic * account
		}
	default:
		fmt.Println("不好意思我们店里没有这个东西")
		return 1, 0
	}
}
