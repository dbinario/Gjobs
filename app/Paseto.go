package app

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/joho/godotenv"
)

func GenerarLlaves() error {

	key := paseto.NewV4AsymmetricSecretKey()

	archivo := ".env"

	// Línea que quieres agregar al archivo

	linea1 := fmt.Sprintf("PASETO_PRIVATE_KEY=%s", key.ExportHex())
	linea2 := fmt.Sprintf("PASETO_PUBLIC_KEY=%s", key.Public().ExportHex())

	// Abrir el archivo en modo de escritura con la bandera de añadir ('a' = append)
	file, err := os.OpenFile(archivo, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	// Crear un escritor bufio
	writer := bufio.NewWriter(file)

	// Escribir la nueva línea en el archivo
	_, err = writer.WriteString(linea1 + "\n" + linea2 + "\n")

	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	// Asegurarse de que todos los datos se escriban en el archivo
	err = writer.Flush()
	if err != nil {
		return fmt.Errorf("error flushing writer: %w", err)
	}

	// Volver a cargar las variables de entorno
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("error reloading environment variables: %w", err)
	}

	return nil

}

func GenerateTokenPaseto(id_usuario string, rol string) string {

	token := paseto.NewToken()

	//creamos el token y le damos una duracion de 2 horas
	token.SetIssuedAt(time.Now())
	token.SetNotBefore(time.Now())
	token.SetExpiration(time.Now().Add(2 * time.Hour))

	//campos personalizados que agregados al token
	token.SetString("id_usuario", id_usuario)
	token.SetString("rol", rol)

	privatekey, _ := paseto.NewV4AsymmetricSecretKeyFromHex(os.Getenv("PASETO_PRIVATE_KEY"))
	signed := token.V4Sign(privatekey, nil)

	return signed

}
