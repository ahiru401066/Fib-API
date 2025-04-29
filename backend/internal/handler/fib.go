package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"fib_api/internal/utils"
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
	if n <= 0 {
		utils.ResponseError(w, http.StatusBadRequest, "BadRequest")
		return 
	}
	
	// 数値処理
	ans,err := utils.Fib(n)
	if err != nil {
		fmt.Fprint(w, "err")
	}
	fmt.Fprint(w, ans)
}