package handler

import (
	"demo/api/internal/logic"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"demo/api/internal/svc"
)

func OrderNotifHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		l := logic.NewOrderNotifLogic(r.Context(), svcCtx)
		fmt.Println(r.Form, "=========r.Form")
		err = l.OrderNotif(r.Form)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			w.Write([]byte("成功"))
			httpx.Ok(w)
		}
	}
}
