package main

import (
	"fmt"
)

func main() {
	key := ""
	loop := true
	balance := 0.00
	money := 0.00
	note := ""
	details := "收支\t账户金额\t收支金额\t说    明"


	for {
		fmt.Println("\n============家庭收支软件===========")
		fmt.Println("1.收支明细")
		fmt.Println("2.登记收入")
		fmt.Println("3.登记支出")
		fmt.Println("4.退出软件")

		fmt.Println("请选择（1-4）")
		fmt.Scanln(&key)

		switch key {
			case "1":
				fmt.Println(details);
			case "2":
				fmt.Println("请输入本次收入金额")
				fmt.Scanln(&money)
				fmt.Println("请输入本次收入备注")
				fmt.Scanln(&note)

				balance += money
				details += fmt.Sprintf("\n收入\t%v\t%v\t%v", balance, money, note)
				fmt.Print(details)
			case "3":
				fmt.Println("请输入本次支出金额")
				fmt.Scanln(&money)
				if money > balance {
					fmt.Println("余额不足")
					break
				}

				fmt.Println("请输入本次支出备注")
				fmt.Scanln(&note)

				balance -= money
				details += fmt.Sprintf("\n支出\t%v\t%v\t%v", balance, money, note)
				fmt.Print(details)
			case "4":
				choice := ""
				fmt.Println("请输入y/n确定")
				for {
					fmt.Scanln(&choice)
					if (choice == "y" || choice == "n") {
						break
					}

					fmt.Println("请输入正确的选项")
				}

				if choice == "y" {
					loop = false
				}
			default:
				fmt.Println("请选择正确的选项")
		}


		if !loop {
			fmt.Println("你已经退出")
			break
		}
	}
}
