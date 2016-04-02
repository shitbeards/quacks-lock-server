package main

type counter struct {
    count int
}

func (c *counter) increment() {
    c.count = c.count + 1
}

var quacks = counter{}
