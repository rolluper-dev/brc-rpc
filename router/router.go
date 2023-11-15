package router

import (
	"github.com/gin-gonic/gin"

	"github.com/rolluper-dev/brc-rpc/controller"
)

const (
	V1BTCGroup = "/v1/btc"
)

func Register(engine *gin.Engine) {
	v1btc := engine.Group(V1BTCGroup)
	// inscription
	v1btc.GET("/inscriptionList", controller.InscriptionList)
	v1btc.GET("/inscription", controller.Inscription)
	v1btc.GET("/inscriptionHistory", controller.InscriptionHistory)

	// token
	v1btc.GET("/token", controller.GetToken)
	v1btc.GET("/tokenList", controller.TokenList)
	v1btc.GET("/tokenHolders", controller.TokenHolders)
	v1btc.GET("/tokenHistory", controller.TokenHistory)

	// address
	v1btc.GET("/addressTokenBalance", controller.AddressTokenBalance)
	v1btc.GET("/addressTransferList", controller.AddressTransferList)
	v1btc.GET("/addressHistory", controller.AddressHistory)

}
