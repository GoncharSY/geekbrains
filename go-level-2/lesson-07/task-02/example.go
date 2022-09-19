package main

var var1 = example0()
var var2 = example1()

func example0() bool {
	return true
}

func example1() int {
	go example0()
	go example0()
	return 1
}

func example2() {
	go func() {}()
	example0()
	example1()
}
