package entity

// ============================== request ==============================

type AddressTokenBalanceRequest struct {
	Page    int    `form:"page"`
	Size    int    `form:"size"`
	Token   string `form:"token"`
	Address string `form:"address" binding:"required"`
}

type AddressTransferListRequest struct {
	Page        int    `form:"page"`
	Size        int    `form:"size"`
	Token       string `form:"token"`
	Address     string `form:"address" binding:"required"`
	Brc201Chain string `form:"brc201_chain"`
}

type AddressHistoryRequest struct {
	Page         int    `form:"page"`
	Size         int    `form:"size"`
	Type         string `form:"type"`
	Token        string `form:"token" binding:"required"`
	Address      string `form:"address" binding:"required"`
	Brc201Extend string `form:"brc201Extend"`
	Brc201Chain  string `form:"brc201_chain"`
}

// ============================== response ==============================

type AddressTokenBalanceResponse struct {
	TokenType           string `json:"tokenType"`
	TokenTick           string `json:"tokenTick"`
	Balance             string `json:"balance"`
	TransferableBalance string `json:"transferableBalance"`
	AvaliableBalance    string `json:"avaliableBalance"`
}

type AddressTransferListResponse struct {
	TokenType         string `json:"tokenType"`
	TokenTick         string `json:"tokenTick"`
	Amount            string `json:"amount"`
	InscriptionId     string `json:"inscriptionId"`
	InscriptionNumber string `json:"inscriptionNumber"`
	Brc201Extend      int    `json:"brc201Extend"`
	Brc201Chain       string `json:"brc201_chain"`
	Brc201Ref         string `json:"brc201_ref"`
}

type AddressHistoryResponse struct {
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
