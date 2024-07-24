package interfaces

//interfaz para reigstrar usuarios
type Registrar struct {
	Name     string `json:"name" binding:"required"`
	LastName string `json:"lastname" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Rol      int    `json:"rol" binding:"required"`
}

type PostAutenticar struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
