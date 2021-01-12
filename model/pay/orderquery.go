package pay

import (
	"time"
)

// OrderQueryResponse OrderQueryResponse struct
type OrderQueryResponse struct {
	OpenID             string    // The unique identifier of the user under the merchant appid
	ProductID          string    // Order ID
	TrandeState        string    // Trading status
	TransactionID      string    // WeChat Pay order number
	OutTradeNo         string    // Merchant system order number
	TradeType          string    // Transaction type submitted by calling the interface: NATIVE
	BankType           string    // Bank type
	TotalFee           int64     // Total order amount: unit is cent
	CashFee            int64     // Cash payment amount
	TimeEnd            time.Time // Order payment time
	DeviceInfo         string    // Terminal device number assigned by WeChat Pay
	SettlementTotalFee int64     // Settlement total fee = Total order amount - Non-rechargeable vouchers
	FeeType            string    // Currency type：CNY
	Detail             string    // Product details
}
