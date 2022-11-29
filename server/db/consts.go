package db


//consts operation type in adding task
const (
	AddPrice    = 0
	ReducePrice = 1
)


//consts condition task
const (
	TaskNew     = 0
	TaskInQueue = 1
	TaskInJob   = 2
	TaskSuccess = 3
	TaskError   = 4
)
