package main

import "fmt"

type exchangeStater interface {
	balance() string
}

type moscowExchange struct {
	balance	int
	state exchangeStater
}

func new(balance int, state exchangeStater) moscowExchange {
	return moscowExchange{balance: balance, state: state}
}

type openedExchange struct {
	exchange moscowExchange
}

func (op *openedExchange) balance() string {
	return fmt.Sprintf("Exchange is opened. Current balance is %d", op.exchange.balance)
}

type closedExchange struct {
	exchange moscowExchange
}

func (cl *closedExchange) balance() string {
	return fmt.Sprintf("Exchange is closed. Balance is fixed at %d", cl.exchange.balance)
}

func main() {
	ex := new(228, nil)
	op := &openedExchange{exchange: ex}
	cl := &closedExchange{exchange: ex}

	ex.state = op

	fmt.Println(ex.state.balance())

	ex.state = cl
	fmt.Println(ex.state.balance())
}
