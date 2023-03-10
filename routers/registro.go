package routers

import (
	"encoding/json"
	"net/http"

	"github.com/JosueOblitas/twittor/bd"
	"github.com/JosueOblitas/twittor/models"
)
func Registro(w http.ResponseWriter, r *http.Request){
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w,"Error en los datos recibidos "+err.Error(),400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w,"El email de usuario es requerido"+err.Error(),400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w,"Debe espeificar una contraseña de al menos 6 caracteres"+err.Error(),400)
		return
	}
	_,encontrado,_ := bd.ChequeoYaExisteUsuario(t.Email)

	if encontrado == true {
		http.Error(w,"Ya existe un usuario registrado con ese email"+err.Error(),400)
		return
	}
	_,status,err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w,"Ocurrio un error al intentar realizar el registro de usuarios :"+err.Error(),400)
		return
	}
	if status == false {
		http.Error(w,"No se ha logrado insertar el registro del usuario :"+err.Error(),400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}