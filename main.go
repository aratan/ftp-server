package main

import (
	"log"
	"os"

	"goftp.io/server/core"
	"goftp.io/server/driver/file"
)

func main() {
	// Define el directorio raíz para el servidor FTP.
	// Asegúrate de que esta carpeta exista.
	ftpRoot := "./ftp-files"
	if _, err := os.Stat(ftpRoot); os.IsNotExist(err) {
		log.Printf("El directorio raíz %s no existe, creándolo...", ftpRoot)
		os.Mkdir(ftpRoot, 0755)
	}

	// Configura el "driver" del sistema de archivos.
	// Esto le dice al servidor cómo interactuar con los archivos locales.
	factory := &file.DriverFactory{
		// Ruta de la carpeta que se compartirá.
		RootPath: ftpRoot,
		// Define un sistema de permisos simple.
		Perm: core.NewSimplePerm("user", "group"),
	}

	// Configura las opciones del servidor.
	opts := &core.ServerOpts{
		// Define el nombre del servidor que se mostrará a los clientes.
		Name:    "Go FTP Server",
		// Asocia el driver del sistema de archivos al servidor.
		Factory: factory,
		// Define el nombre de usuario y la contraseña.
		// Para un uso real, considera un sistema de autenticación más seguro.
		Auth: &core.SimpleAuth{
			Name:     "admin",
			Password: "password",
		},
		// El puerto en el que escuchará el servidor.
		Port:           2121,
		// Mensaje de bienvenida que verán los usuarios al conectarse.
		WelcomeMessage: "Bienvenido al servidor FTP de Go.",
	}

	// --- CORRECCIÓN AQUÍ ---
	// Crea una nueva instancia del servidor con las opciones definidas.
	// NewServer ahora solo devuelve una instancia del servidor y entra en pánico en caso de error.
	s := core.NewServer(opts)

	// Inicia el servidor y espera a que los clientes se conecten.
	log.Printf("Iniciando servidor FTP en el puerto %d...", opts.Port)
	log.Printf("Usuario: %s | Contraseña: %s", opts.Auth.(*core.SimpleAuth).Name, opts.Auth.(*core.SimpleAuth).Password)
	log.Printf("Directorio compartido: %s", ftpRoot)

	err := s.ListenAndServe()
	if err != nil {
		log.Fatal("Error al iniciar el servidor: ", err)
	}
}
