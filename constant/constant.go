package constant

const TokenTypeBRC20 = "brc20"

const (
	TokenStatusNon = iota
	TokenStatusMinting
	TokenStatusComplete
)

const TimeFormat = "2006-01-02 15:04:05"

const BlockHandlerHeight = "block_handler_height"

const (
	BRC20_HISTORY_TYPE_INSCRIBE_DEPLOY   = iota + 1 // "inscribe-deploy"
	BRC20_HISTORY_TYPE_INSCRIBE_MINT                // "inscribe-mint"
	BRC20_HISTORY_TYPE_INSCRIBE_TRANSFER            // "inscribe-transfer"
	BRC20_HISTORY_TYPE_TRANSFER                     //"transfer"
	BRC20_HISTORY_TYPE_SEND                         // "send"
	BRC20_HISTORY_TYPE_RECEIVE                      // "receive"
)

func HistoryType(str string) uint8 {
	switch str {
	case "deploy":
		return BRC20_HISTORY_TYPE_INSCRIBE_DEPLOY
	case "mint":
		return BRC20_HISTORY_TYPE_INSCRIBE_MINT
	case "transfer":
		return BRC20_HISTORY_TYPE_TRANSFER
	case "inscribeTransfer":
		return BRC20_HISTORY_TYPE_INSCRIBE_TRANSFER
	case "send":
		return BRC20_HISTORY_TYPE_SEND
	case "receive":
		return BRC20_HISTORY_TYPE_RECEIVE
	default:
		return 0
	}
}

func StrHistoryType(tp uint8) string {
	switch tp {
	case BRC20_HISTORY_TYPE_INSCRIBE_DEPLOY:
		return "deploy"
	case BRC20_HISTORY_TYPE_INSCRIBE_MINT:
		return "mint"
	case BRC20_HISTORY_TYPE_TRANSFER:
		return "transfer"
	case BRC20_HISTORY_TYPE_INSCRIBE_TRANSFER:
		return "inscribeTransfer"
	case BRC20_HISTORY_TYPE_SEND:
		return "send"
	case BRC20_HISTORY_TYPE_RECEIVE:
		return "receive"
	default:
		return ""
	}
}

func StrBrc201Extend(tp string) uint8 {
	switch tp {
	case "bridge-in":
		return 1
	case "bridge-out":
		return 2
	default:
		return 0
	}
}
