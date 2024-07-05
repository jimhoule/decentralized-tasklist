package controllers

import (
	"gateway/presenters/http/utils"
	"net/http"
	"net/rpc"
)

type WalletsController struct{
	WalletsRpcClient *rpc.Client
}

func (wc *WalletsController) Create(writer http.ResponseWriter, request *http.Request) {	
		var wallet any;
		// NOTE: WalletsRpcHandler is the registered Rpc Server in wallets microservice
		err := wc.WalletsRpcClient.Call("WalletsRpcHandler.Create", nil, &wallet);
		if err != nil {
			utils.WriteHttpError(writer, http.StatusInternalServerError, err)
			return
		}

		utils.WriteHttpResponse(writer, http.StatusCreated, wallet)
}