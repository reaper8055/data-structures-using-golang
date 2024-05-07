package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sync"
)

type order struct {
	ProductCode int
	Quantity    float64
	status      orderStatus
}

type orderStatus int

type invalidOrder struct {
	order order
	err   error
}

const (
	none orderStatus = iota
	new
	received
	reserved
	filled
)

func (o order) String() string {
	return fmt.Sprintf("Product code: %v, Quantity: %v, Status: %v\n",
		o.ProductCode, o.Quantity, orderStatusToText(o.status))
}

func orderStatusToText(o orderStatus) string {
	switch o {
	case none:
		return "none"
	case new:
		return "new"
	case received:
		return "received"
	case filled:
		return "filled"
	default:
		return "unknown status"
	}
}

var rawOrders = []string{
	`{"productCode": 1111, "quantity": 5, "status": 1}`,
	`{"productCode": 2222, "quantity": 42.5, "status": 1}`,
	`{"productCode": 3333, "quantity": 19, "status": 1}`,
	`{"productCode": 4444, "quantity": 8, "status": 1}`,
	`{"productCode": 5555, "quantity": -10, "status": 1}`,
}

func main() {
	var wg sync.WaitGroup
	receivedOrderCh := receiveOrders()
	validOrderCh, invalidOrderCh := validateOrders(receivedOrderCh)
	wg.Add(1)
	go func(validOrder <-chan order, invalidOrder <-chan invalidOrder) {
	loop:
		for {
			select {
			case validOrder, ok := <-validOrderCh:
				if !ok {
					break loop
				}
				fmt.Println(validOrder)
			case invalidOrder, ok := <-invalidOrderCh:
				if !ok {
					break loop
				}
				fmt.Println(invalidOrder.order)
				fmt.Println(invalidOrder.err)
			}
		}
		wg.Done()
	}(validOrderCh, invalidOrderCh)
	wg.Wait()
}

func validateOrders(in <-chan order) (<-chan order, <-chan invalidOrder) {
	out := make(chan order)
	errCh := make(chan invalidOrder, 1)
	go func() {
		for order := range in {
			if order.Quantity < 0 {
				errCh <- invalidOrder{
					order: order,
					err:   errors.New("Quantity should be positive"),
				}
			} else {
				out <- order
			}
		}
		close(out)
		close(errCh)
	}()
	return out, errCh
}

func receiveOrders() <-chan order {
	out := make(chan order)
	go func() {
		for _, rawOrder := range rawOrders {
			var newOrder order
			err := json.Unmarshal([]byte(rawOrder), &newOrder)
			if err != nil {
				log.Print(err)
				continue
			}
			out <- newOrder
		}
		close(out)
	}()
	return out
}
