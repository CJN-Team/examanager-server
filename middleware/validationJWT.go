package middleware

import (
	"net/http"

	"github.com/CJN-Team/examanager-server/routers"
)

func ValidationJWT(next http.HandlerFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){
		_,_,_, error := routers.GetToken(r.Header.Get("Authorization"))
		if error!= nil{
			http.Error(w,"Error en el token"+error.Error(),http.StatusBadRequest)
		}
		next.ServeHTTP(w,r)
	}
}