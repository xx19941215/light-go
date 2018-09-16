package main

import (
	"light-go/c04/c04_12/service"
	"light-go/c04/c04_12/model"
	"fmt"
)

type customerView struct {
	key string
	loop bool
	customerService *service.CustomerService
}

func (this *customerView) list() {
	customers := this.customerService.List()
	fmt.Println("---------------------------客户列表---------------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		//fmt.Println(customers[i].Id,"\t", customers[i].Name...)
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Printf("\n-------------------------客户列表完成-------------------------\n\n")
}

func (this *customerView) add() {
	fmt.Println("---------------------添加客户---------------------")
	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("电邮:")
	email := ""
	fmt.Scanln(&email)
	customer := model.NewCustomer2(name, gender, age, phone, email)
	if this.customerService.Add(customer) {
		fmt.Println("---------------------添加完成---------------------")
	} else {
		fmt.Println("---------------------添加失败---------------------")
	}
}

func (this *customerView) delete() {
	fmt.Println("---------------------删除客户---------------------")
	fmt.Println("请选择待删除客户编号(-1退出)：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return //放弃删除操作
	}
	fmt.Println("确认是否删除(Y/N)：")
	//这里可以加入一个循环判断，直到用户输入 y 或者 n,才退出..
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" || choice == "Y" {
		//调用customerService 的 Delete方法
		if this.customerService.Delete(id) {
			fmt.Println("---------------------删除完成---------------------")
		} else {
			fmt.Println("---------------------删除失败，输入的id号不存在----")
		}
	}
}

//退出软件 
func (this *customerView) exit() {

	fmt.Println("确认是否退出(Y/N)：")
	for {
		fmt.Scanln(&this.key)
		if this.key == "Y" || this.key == "y" || this.key == "N" || this.key == "n" {
			break
		}

		fmt.Println("你的输入有误，确认是否退出(Y/N)：")
	}

	if this.key == "Y" || this.key == "y" {
		this.loop = false
	}

}

//显示主菜单
func (this *customerView) mainMenu() {

	for {
		
		fmt.Println("-----------------客户信息管理软件-----------------")
		fmt.Println("                 1 添 加 客 户")
		fmt.Println("                 2 修 改 客 户")
		fmt.Println("                 3 删 除 客 户")
		fmt.Println("                 4 客 户 列 表")
		fmt.Println("                 5 退       出")
		fmt.Print("请选择(1-5)：")

		fmt.Scanln(&this.key)
		switch this.key {
			case "1" :
				this.add()
			case "2" :
				this.update()
			case "3" :
				this.delete()
			case "4" :
				this.list()
			case "5" :
				this.exit()
			default :
				fmt.Println("你的输入有误，请重新输入...")
		}

		if !this.loop {
			break
		}

	}
	fmt.Println("你退出了客户关系管理系统...")
}

func main() {
	//在main函数中，创建一个customerView,并运行显示主菜单..
	customerView := customerView{
		key : "",
		loop : true,
	}
	//这里完成对customerView结构体的customerService字段的初始化
	customerView.customerService = service.NewCustomerService()
	//显示主菜单.
	customerView.mainMenu()

}

func (this *customerView) update() {
	var (
		id int
		updateIndex int
	)
	fmt.Println("请输入需要修改的用户ID(无需请输入-1)")
	if id == -1 {
		return
	}

	fmt.Scanln(&id)

	updateIndex = this.customerService.FindById(id)

	if updateIndex == -1 {
		fmt.Println("no this users")
		return 
	}

	_, name, gender, age, phone, email := this.customerService.List()[updateIndex].GetInfo()

	fmt.Printf("姓名(%v):", name)
	fmt.Scanln(&name)
	fmt.Printf("性别(%v):", gender)
	fmt.Scanln(&gender)
	fmt.Printf("年龄(%v):", age)
	fmt.Scanln(&age)
	fmt.Printf("电话(%v):", phone)
	fmt.Scanln(&phone)
	fmt.Printf("邮箱(%v):", email)
	fmt.Scanln(&email)
	this.customerService.List()[updateIndex].Update(name, gender, age, phone, email)
}