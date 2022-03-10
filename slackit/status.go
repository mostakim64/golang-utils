package slackit

const (
	Success = iota + 1
	Warning
	Error

)

var StatusMap = map[int]string{
	Success : "#00FF00",
	Warning : "#FFFF00",
	Error : "#FF0000",
}