package goedcmeg

type EdcMega struct {
	Config
}

func NewEdc(cfg Config) Edc {
	edc := &EdcMega{Config: cfg}
	return edc
}

func (e EdcMega) Sale(tx SaleTransaction) (SaleResult, error) {
	return SaleResult{}, nil
}
