package service

import (
	"light-go/c04/c04_12/model"
)

type CustomerService struct {
	customers []model.Customer
	customerNum int
}

func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "xiaoxiao", "男", 24, "17676767676", "8888@qq.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}


func (this *CustomerService) Add(customer model.Customer) bool {
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers, customer)
	return true
}

func (this *CustomerService) Delete(id int) bool {
	index := this.FindById(id)
	if index == -1 {
		return false
	}

	//从切片中删除一个元素
	this.customers = append(this.customers[:index], this.customers[index+1:]...)
	return true
}

func (this *CustomerService) List() []model.Customer {
	return this.customers
}

func (this *CustomerService) FindById(id int) int {
	index := -1
	for i := 0; i < len(this.customers); i++ {
		if this.customers[i].Id == id {
			index = i
		}
	}

	return index
}