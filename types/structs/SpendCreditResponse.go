package structs

type SpendCreditResponse struct {
	Amount     uint
	UsageLogId uint
	Err        *CustomError
}
