package middleware

import (
	"net/http"

	"github.com/CJN-Team/examanager-server/database"
)

//DatabaseVerify se encarga de verificar que la conexion de la base de datos se encuentra funcionando antes de realizar cualquier operacion
func DatabaseVerify(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if database.CheckConnection() == 0 {
			http.Error(w, "Conexion perdida con la Base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
