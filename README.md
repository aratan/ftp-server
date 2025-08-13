# ftp-server
pruebas de penetracion por FB


Servidor/Host: 127.0.0.1
Puerto: 2121
Usuario: admin
Contraseña: password

Desde el cliente sería:

```bash
ftp open 127.0.0.1 2121

```

### Comandos de Sesión:

USER: Especifica el nombre de usuario.

PASS: Especifica la contraseña.

QUIT: Termina la sesión.

SYST: Pide información del sistema operativo del servidor.

TYPE: Establece el modo de transferencia (ASCII/Binario).

PORT / PASV: Comandos para establecer la conexión de datos (modos activo y pasivo).

### Comandos de Navegación y Listado:

PWD: (Print Working Directory) Muestra el directorio actual.

CWD: (Change Working Directory) Cambia a un directorio diferente.

CDUP: Cambia al directorio padre.

LIST / NLST: Lista los archivos y directorios.

### Comandos de Transferencia y Gestión de Archivos:

RETR: (Retrieve) Descarga un archivo desde el servidor al cliente.

STOR: (Store) Sube un archivo desde el cliente al servidor.

DELE: (Delete) Borra un archivo en el servidor.

RNFR / RNTO: (Rename From / Rename To) Renombra un archivo.

MKD: (Make Directory) Crea un nuevo directorio.

RMD: (Remove Directory) Borra un directorio.






