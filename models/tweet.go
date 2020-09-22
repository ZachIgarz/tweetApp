package models

/*Tweet captura del modi los mensajes que llegan*/
type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`
}
