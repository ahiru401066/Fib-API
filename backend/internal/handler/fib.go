package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"fib_api/internal/service"
)

func responseError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	// json形式でレスポンス
	json.NewEncoder(w).Encode(ErrorResponse{
		Status: http.StatusBadRequest,
		Message: message,
	})
} 

func FibHandler(w http.ResponseWriter, r *http.Request) {
	// リクエストメゾットを確認
	if r.Method != http.MethodGet {
		responseError(w, http.StatusBadRequest, "BadRequest")
		return 
	}
	// クエリパラメータの取得
	query := r.URL.Query()
	nStr := query.Get("n")
	if nStr == "" {
		responseError(w, http.StatusBadRequest, "BadRequest")
		return 
	}
	// バリデーション
	n, err := strconv.Atoi(nStr)
	if err != nil {
		responseError(w, http.StatusBadRequest, "BadRequest")
		return 
	}
	if n <= 0 {
		responseError(w, http.StatusBadRequest, "BadRequest")
		return 
	}
	
	// 数値処理
	ans,err := service.Fib(n)
	if err != nil {
		fmt.Fprint(w, "err")
	}
	fmt.Fprint(w, ans)
}