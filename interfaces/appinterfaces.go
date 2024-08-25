package interfaces

//interfaz para reigstrar usuarios
type Registrar struct {
	Name     string `json:"name" binding:"required"`
	LastName string `json:"lastname" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Rol      int    `json:"rol" binding:"required"`
}

type Autenticar struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type ActualizarDatosEmpresa struct {
	NombreEmpresa string `json:"nombreEmpresa" binding:"required"`
}

type PublicarVacante struct {
	IdEmpresa      int    `json:"idEmpresa"`
	NombreVacante  string `json:"nombreVacante"`
	OcultarEmpresa int    `json:"ocultarEmpresa"`
	Contratacion   int    `json:"contratacion"`
	Horario        int    `json:"horario"`
	Modalidad      int    `json:"modalidad"`
	Municipio      int    `json:"municipio"`
	Estado         int    `json:"estado"`
	RangoMin       int    `json:"rangoMin"`
	RangoMax       int    `json:"rangoMax"`
	OcultarRango   int    `json:"ocultarRango"`
	Descripcion    string `json:"descripcion"`
}

type Vacante struct {
	IdVacante        int    `json:"idVacante"`
	NombreEmpresa    string `json:"nombreEmpresa"`
	IdEmpresa        int    `json:"idEmpresa"`
	NombreVacante    string `json:"nombreVacante"`
	TipoContratacion string `json:"tipoContratacion"`
	Horario          string `json:"horario"`
	Modalidad        string `json:"modalidad"`
	Municipio        string `json:"municipio"`
	Estado           string `json:"estado"`
	RangoMin         int    `json:"rangoMin"`
	RangoMax         int    `json:"rangoMax"`
	Descripcion      string `json:"descripcion"`
}
