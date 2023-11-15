package entity

// ============================== request ==============================

type InscriptionListRequest struct {
	Page    int    `form:"page"`
	Size    int    `form:"size"`
	Address string `form:"address"`
}

type InscriptionRequest struct {
	InscriptionId     string `form:"inscriptionId"`
	InscriptionNumber string `form:"inscriptionNumber"`
}

type InscriptionHistoryRequest struct {
	InscriptionId string `form:"inscriptionId" binding:"required"`
	Page          int    `form:"page"`
	Size          int    `form:"size"`
}

// ============================== response ==============================

type InscriptionListResponse struct {
	Owner             string `json:"owner"`
	Output            string `json:"output"`
	InscriptionID     string `json:"inscriptionId"`
	InscriptionNumber int64  `json:"inscriptionNumber"`
}

type InscriptionResponse struct {
	InscriptionID     string `json:"inscriptionId"`
	InscriptionNumber int64  `json:"inscriptionNumber"`
	TxID              string `json:"txid"`
	Holder            string `json:"holder"`
	Creator           string `json:"creator"`
	DeployHeight      string `json:"deployHeight"`
	DeployTime        string `json:"deploy_time"`
	Output            string `json:"output"`
}

type InscriptionHistoryResponse struct {
	TxID   string `json:"txid"`
	Block  string `json:"block"`
	TxType string `json:"txType"`
	From   string `json:"from"`
	To     string `json:"to"`
	Time   string `json:"time"`
}
