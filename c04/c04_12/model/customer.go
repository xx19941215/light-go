package model

import (
)

type Customer struct {
	Id int
	Name string
	Gender string
	Age int
	Phone string
	Email string
}

func NewCustomer(id int, name string, gender string, age int, phone string, email string) Customer {
	return Customer{
		Id: id,
		Name : name,
		Gender : gender,
		Age : age,
		Phone : phone,
		Email : email,
	}
}

func NewCustomer2(name string, gender string, age int, phone string, email string) Customer {
	return Customer{
		Name : name,
		Gender : gender,
		Age : age,
		Phone : phone,
		Email : email,
	}
}

func (c *Customer) GetInfo() (int, string, string, int, string, string) {
	 return c.Id, c.Name, c.Gender, c.Age, c.Phone, c.Email
}

func (c *Customer) Update(name string, gender string, age int, phone string, email string) {
	if name != "" {
		c.Name = name
	}

	if gender != "" {
		c.Gender = gender
	}

	if age != 0 {
		c.Age = age
	}

	if phone != "" {
		c.Phone = phone
	}

	if email != "" {
		c.Email = email
	}
}