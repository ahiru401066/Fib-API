package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"fib_api/internal/utils"
	"fib_api/internal/service"
	"fib_api/internal/model"
)

func FibHandler(w http.ResponseWriter, r *http.Request) {
	// リクエストメゾットを確認
	if r.Method != http.MethodGet {
		utils.ResponseError(w, http.StatusBadRequest, "BadRequest")
		return 
	}
	// クエリパラメータの取得
	query := r.URL.Query()
	nStr := query.Get("n")
	if nStr == "" {
		utils.ResponseError(w, http.StatusBadRequest, "BadRequest")
		return 
	}
	// バリデーション
	n, err := strconv.Atoi(nStr)
	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, "BadRequest")
		return 
	}
	
	// 数値処理
	ans,err := service.Fib(n)
	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, "BadRequest")
		return 
	}

	// リクエスト処理
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.ResultResponse{
		Result: ans.String(),
	})
}