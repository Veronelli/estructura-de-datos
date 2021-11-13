package miPackage

// CarPublic Car con acceso publico;
type CarPublic struct {
	Brand string
	Year  uint
}

// carPrivate de acceso privado
type carPrivate struct {
	brand string
	year  uint
}

func Saludar(nombre string) string {
	return "Hola " + nombre

}
