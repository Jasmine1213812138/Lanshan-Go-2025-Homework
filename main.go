package main

import (
	"fmt"
	"lesson03/Warehouse"
	"os"
)

func main() {
	//定义变量
	var want string
	var account int
	judge := 0
	var income int
	Pierre := Warehouse.Warehouse{
		Parsnip:    5,
		Greenbeans: 5,
		Potato:     5,
		Garlic:     8,
		Tulip:      10,
	}
	price := Warehouse.Price{
		Parsnip:    20,
		Greenbeans: 60,
		Potato:     50,
		Garlic:     40,
		Tulip:      70,
	}
	//前台欢迎
	fmt.Printf("欢迎来到皮埃尔商店，你想要点什么种子（格式：物品 数量）")
	for {
		_, err := fmt.Scanf("%s %d", &want, &account)
		if err == nil {
			break
		}
		fmt.Println("抱歉，你说啥勒,恁要这样跟我说：物品 数量，麻烦回车后重新输入")
		var disgard string
		fmt.Scanln(&disgard)
	}
	fmt.Println("好嘞，您稍等")
	//判断是否有库存
	judge, income = Warehouse.Judge(want, account, Pierre, price, judge)
	for judge == 1 {
		fmt.Println("你还要买点啥不，不要就说no，要就说yes ")
		var again string
		fmt.Scanln(&again)
		if again == "no" {
			fmt.Printf("拜拜下次光临")
			os.Exit(0)
		} else if again == "yes" {
			fmt.Println("好了，说说你要啥吧")
			_, err1 := fmt.Scanf("%s %d", &want, &account)
			if err1 != nil {
				fmt.Println("老妹儿你又说错了,麻烦回车一下")
				var discard string
				fmt.Scanln(&discard)
				continue
			}
			judge = 0
			judge, income = Warehouse.Judge(want, account, Pierre, price, judge)
		} else {
			fmt.Println("老妹你说啥呢")

		}
	}
	fmt.Println("请输入你的付款")
	//进行收银
	var pay int
	for {
		_, err := fmt.Scanf("%d", &pay)
		if err == nil {
			break
		}
		fmt.Println("抱歉，你给的是钱么，麻烦回车后重新输入")
		var discard string
		fmt.Scanln(&discard)
	}
	charge := income - pay
	for income > pay {
		fmt.Printf("小姐不够,还差%d元", charge)
		fmt.Scanf("%d", &pay)
		if pay >= charge {
			fmt.Printf("谢谢惠顾，找零%d元", pay-charge)
			break
		} else {
			fmt.Printf("别玩我，钱还你，你再数数")
			fmt.Scanf("%d", &pay)
		}
	}
	fmt.Printf("谢谢惠顾，找零%d元", pay-income)

	//今日账单
	list := make(map[string]int)
	list = Warehouse.Income(want, income)
	fmt.Println(list)
}
