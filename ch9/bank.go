package ch9

var desposit = make(chan int)
var balances = make(chan int)

// Disposit 存取款
func Disposit(account int) {
	desposit <- account
}

// Balance 查询余额
func Balance() int { return <-balances }

func teller() {
	var banlance int
	for {
		select {
		case account := <-desposit:
			banlance += account
		case balances <- banlance:
		}
	}
}
