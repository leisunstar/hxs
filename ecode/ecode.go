package ecode

const (
	Success           = 0
	UserIdLengthError = iota
	MinerMoneyNotEnough
	WithDrawUserIdError
	WithDrawStatusError
	WithDrawMoneyError
	AddMoneyAmountError
	ChargeAmountError
	LimitEndNumError
	AppOrderIdDuplicateError
	WithDrawApplyError
	ServerError = -1
)

var msg = map[int]string{
	Success:                  "成功", // 0
	UserIdLengthError:        "用户ID长度不对",
	MinerMoneyNotEnough:      "矿工钱包钱数不够",
	WithDrawUserIdError:      "提现userid 不符",
	WithDrawStatusError:      "提现状态码不符",
	WithDrawMoneyError:       "提现钱数错误",
	AddMoneyAmountError:      "加钱钱数",
	ChargeAmountError:        "充值钱数错误",
	LimitEndNumError:         "查询limit条数错误",
	AppOrderIdDuplicateError: "app和orderId联合唯一错误",
	WithDrawApplyError:       "提现账号，实名错误",
	// server error
	ServerError: "未知错误", // -1
}

func GetMessage(code int) (m string) {
	m, ok := msg[code]
	if !ok {
		m = msg[ServerError]
	}
	return
}
