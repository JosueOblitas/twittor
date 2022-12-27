package middlew

import (
	"net/http"

	"github.com/JosueOblitas/twittor/bd"
)
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter,r *http.Request){
		if bd.ChequeoConnection() == 0{
			http.Error(w, "Conexión Perdida con la Base de datos",500)
			return
		}
		next.ServeHTTP(w,r)
	}
}