package main

import "fmt"

/*
你有50枚金币，需要分配给队下几个人：Matthew，Sarah，Augustus，Heidi，Emilie，Peter，Giana，Adriano，Aaron，Elizabeth。
分配规则如下：
a.名字中每包含1个'e'或E'分1枚金币
b.名字中每包含1个‘i’或I'分2枚金币
c.名字中每包含1个‘o’或‘0'分3枚金币
d：名字中每包含1个'u‘或‘u“分4枚金币

写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现‘dispatchcoin’函数
*/

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() {
	//1.依次拿到名字
	for _, name := range users {
		//2.根据规则给人分金币
		for _, c := range name {
			switch c {
			case 'e', 'E':
				//3.记录分发金币数  与金币剩余总数
				distribution[name]++
				coins--
			case 'i', 'I':
				distribution[name] += 2
				coins -= 2
			case 'o', 'O':
				distribution[name] += 3
				coins -= 3
			case 'u', 'U':
				distribution[name] += 4
				coins -= 4
			}
		}
	}
}

func main() {
	dispatchCoin()
	for k, v := range distribution {
		fmt.Printf("%s:%d\n",k,v)
	}
	fmt.Println("剩余金币：", coins)
	//4.向屏幕输出每个人获取金币数量，与剩余金币数量
}
