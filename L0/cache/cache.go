package cache

import (
	"database/sql"
	"fmt"
	"sync"

	"L0/model"
	"L0/common"
	"L0/dbpart"
)

type My_cache struct {
	rwmu sync.RWMutex
	data map[string]model.Order_t
}

func Init() *My_cache {
	dat := make(map[string]model.Order_t)
	out := My_cache{data: dat}
	fmt.Println("Init cache!")
	return &out
}

func (c *My_cache) Write(order model.Order_t) {
	c.rwmu.Lock()
	defer c.rwmu.Unlock()
	c.data[order.Order_uid] = order
	fmt.Println("Write data in cache!")
}

func (c *My_cache) Read(order_uid string) (model.Order_t, bool) {

	c.rwmu.RLock()
	defer c.rwmu.RUnlock()

	order, is_ok := c.data[order_uid]

	if is_ok {
		fmt.Println("Read data from cache!")
	} else {
		fmt.Println("No such data in cache!")
	}

	return order, is_ok
}

func (c *My_cache) RestoreCache(db *sql.DB) {
	orders, err := dbpart.GetAllOrders(db)
	common.CheckFatal(err)

	for _, order := range orders {
		c.Write(order)
		//fmt.Println(order)
	}

	fmt.Println("Cache restored!")
}
