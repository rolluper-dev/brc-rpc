package entity

// ============================== request ==============================

type TokenListRequest struct {
	Page   int   `form:"page"`
	Size   int   `form:"size"`
	Status uint8 `form:"status"`
}

type TokenHoldersRequest struct {
	Page  int    `form:"page"`
	Size  int    `form:"size"`
	Token string `form:"token" binding:"required"`
}

type TokenHistoryRequest struct {
	Token        string `form:"token" binding:"required"`
	Type         string `form:"type"`
	From         string `form:"from"`
	To           string `form:"to"`
	Page         int    `form:"page"`
	Size         int    `form:"size"`
	Brc201Extend string `form:"brc201_extend"`
	Brc201Chain  string `form:"brc201_chain"`
}

type GetTokenRequest struct {
	Token string `form:"token"`
}

// ============================== response ==============================

type TokenListResponse struct {
	TokenType    string `json:"tokenType"`
	Token        string `json:"token"`
	TotalSupply  uint64 `json:"totalSupply"`
	Limit        uint64 `json:"limit"`
	Decimals     uint8  `json:"decimals"`
	Minted       uint64 `json:"minted"`
	Creator      string `json:"creator"`
	DeployTime   int64  `json:"deployTime"`
	DeployHeight int64  `json:"deployHeight"`
	Holders      int64  `json:"holders"`
}

type TokenHoldersResponse struct {
	Address string `json:"address"`
	Amount  string `json:"amount"`
}

type TokenHistoryResponse struct {
	TxId              string `json:"txid"`
	Block             string `json:"block"`
	TokenType         string `json:"tokenType"`
	Token             string `json:"token"`
	TxType            string `json:"txType"` // transaction type: all, deploy, mint, transfer, inscribeTransfer
	From              string `json:"from"`
	To                string `json:"to"`
	Amount            string `json:"amount"`
	InscriptionNumber string `json:"inscriptionNumber"`
	InscriptionId     string `json:"inscriptionId"`
	Time              string `json:"time"`
	Brc201Extend      int    `json:"brc201Extend"`
	Brc201Chain       string `json:"brc201_chain"`
}

type GetTokenResponse struct {
	TokenType              string `json:"tokenType"`
	Token                  string `json:"token"`
	TotalSupply            uint64 `json:"totalSupply"`
	Limit                  uint64 `json:"limit"`
	Decimals               uint8  `json:"decimals"`
	Minted                 uint64 `json:"minted"`
	Creator                string `json:"creator"`
	DeployTime             int64  `json:"deployTime"`
	DeployHeight           int64  `json:"deployHeight"`
	Holders                int64  `json:"holders"`
	TxID                   string `json:"txid"`
	InscriptionNumberStart uint64 `json:"inscriptionNumberStart"`
	InscriptionNumberEnd   uint64 `json:"inscriptionNumberEnd"`
}
