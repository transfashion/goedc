package goedcmeg

type Edc interface {
	Sale(tx SaleTransaction) (SaleResult, error)
}
