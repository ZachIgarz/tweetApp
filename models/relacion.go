package models

/*Relacion de usuarios */
type Relacion struct {
	UsuarioID         string `bson:"usuarioid" json:"usuarioID"`
	UsuarioRelacionID string `bson:"usuariorelacionid" json:"usuarioRelacionID"`
}
