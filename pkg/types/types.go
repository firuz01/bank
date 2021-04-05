package types

type Money int64
//Category is the category on which the payment was provided
type Category string

type Status string

const (
	StatusOk Status = "OK"
	StatusFail Status = "FAIL"
	StatusInProgress Status ="INPROGRESS"
)

type Payment struct {
	ID int
	Amount Money
	Category Category
	Status Status
}
