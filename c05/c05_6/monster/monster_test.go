package monster

import (
	"testing"
)

func TestStore(t *testing.T) {
	monster := &Monster{
		Name: "红孩儿",
		Age: 10,
		Skill: "吐火",
	}

	res := monster.Store()

	if !res {
		t.Fatalf("monster.Store()错误，期望=%v, 实际=%v", true, res)
	}

	t.Logf("monster.Store()成功")
}

func TestUnStore(t *testing.T) {
	var monster = &Monster{}
	res := monster.ReStore()

	if !res {
		t.Fatalf("monster.ReStore() 错误，希望为=%v 实际为=%v", true, res)
	}

	//进一步判断
	if monster.Name != "红孩儿" {
		t.Fatalf("monster.ReStore() 错误，希望为=%v 实际为=%v", "红孩儿", monster.Name)
	}

	t.Logf("monster.ReStore() 测试成功!")
}