package monster

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name string
	Age int
	Skill string
}

func (this *Monster) Store() bool {
	data, err := json.Marshal(this)
	if err != nil {
		fmt.Println("marshal err=", err)
		return false
	}

	filePath := "/tmp/monster.ser";
	err = ioutil.WriteFile(filePath, data, 0666)

	if err != nil {
		fmt.Println("write file err=", err)
		return false
	}

	return true
}

func (this *Monster) ReStore() bool {
	filePath := "/tmp/monster.ser"
	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		fmt.Println("Readfile err=", err)
		return false
	}

	err = json.Unmarshal(data, this)

	if err != nil {
		fmt.Println("unmarshal err=", err)
		return false
	}

	return true
}

