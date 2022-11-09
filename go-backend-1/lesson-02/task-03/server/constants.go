package server

const DefaultHost = "localhost"
const DefaultPort = "9002"

type Operator uint8

const (
	Inc Operator = iota + 1
	Dec
	Mul
	Div
)
