package main

import "fmt"

// 定义角色结构体
type Person struct {
	Name       string // 名字
	Level      int    // 等级
	Experience int    // 经验值
	Health     int    // 血量
	Attack     int    // 攻击力
}

// 定义角色攻击接口
type Attacker interface {
	Attacki(target *Person)
}

// 角色攻击方法的实现
func (p *Person) Attacki(target *Person) {
	fmt.Printf("%s 攻击 %s，造成 %d 点伤害！\n", p.Name, target.Name, p.Attack)
	target.Health -= p.Attack
	fmt.Printf("%s 剩余血量为 %d\n", target.Name, target.Health)
}

func main() {
	// 创建两个角色
	player1 := &Person{Name: "player1", Level: 1, Experience: 0, Health: 80, Attack: 10}
	player2 := &Person{Name: "player2", Level: 1, Experience: 0, Health: 100, Attack: 8}

	// 开始攻击
	for player2.Health != 0 && player1.Health != 0 {
		player1.Attacki(player2)
		if player2.Health == 0 {
			break
		}
		player2.Attacki(player1)
	}
	fmt.Printf("Game is over!")
}
