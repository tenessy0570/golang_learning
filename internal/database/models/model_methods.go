package models

import "fmt"

type Printable interface {
	ToString() string
}

func ToString(obj any) string {
	return fmt.Sprintf("%+v", obj)
}

func (a *Account) ToString() string {
	return ToString(a)
}

func (c *Cart) ToString() string {
	return ToString(c)
}

func (i *ShopItem) ToString() string {
	return ToString(i)
}
