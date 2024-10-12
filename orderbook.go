package orderbook

import (
	"container/heap"

	"github.com/google/uuid"
)

type Price float64
type Quantity uint
type Side uint8

const (
	Buy Side = iota
	Sell
)

type Limit struct {
	LimitPrice Price
}

type Order struct {
	OrderId           uuid.UUID
	Price             Price
	Side              Side
	Quantity          Quantity
	RemainingQuantity Quantity
	Next              *Order
	Prev              *Order
}

type OrderChain struct {
	start *Order
	end   *Order
}

type OrderBook struct {
	buyHeap  Limit_Max_PQ
	sellHeap Limit_Min_PQ
	buyMap   map[Price]OrderChain
	sellMap  map[Price]OrderChain
}

func (ob *OrderBook) addBuyOrder(price Price, quantity Quantity) *Order {
	order := &Order{uuid.New(), price, Side(Buy), quantity, quantity, nil, nil}

	// If the order already exists in the buy map, link it to the existing chain.
	if chain, exists := ob.buyMap[price]; exists {
		// Update order pointers: link the new order to the current chain's end.
		temp := chain.end
		chain.end.Next = order
		order.Prev = temp
		chain.end = order

	} else { // Otherwise, add the order to the book and insert the limit into the heap.
		// Add limit to heap
		limit := Limit{price}
		heap.Push(&ob.buyHeap, limit)

		// Start chain in orderbook
		chain = OrderChain{order, order}
		ob.buyMap[price] = chain
	}

	return order
}
