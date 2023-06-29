/* 
# Para probar la aplicación, puedes guardar el código en un archivo llamado app.go y compilarlo con el comando go build app.go. 
# Luego puedes ejecutar el archivo generado (app en Unix o app.exe en Windows) para iniciar el servidor. 
# Si no se especifica el puerto en la variable de entorno PUERTO, el servidor escuchará en un puerto aleatorio y lo imprimirá en la consola al inicio.
# Si se especifica un puerto, el servidor escuchará en ese puerto en su lugar. 
# Los pedidos de conexión que reciba el servidor se loguearán en la consola y el archivo de inicio, y se responderán con el contenido del archivo /config/respuesta.txt, o con un mensaje de texto genérico si ese archivo no existe.
# El HandleFunc que hemos creado utiliza la función anónima para manejar todas las solicitudes a la raíz / y devuelve la respuesta definida en la variable respuesta.
# Si la variable de entorno PUERTO no está definida, generamos un número aleatorio entre 1001 y 10000 utilizando math/rand y lo convertimos en una cadena de texto para utilizarlo como puerto.
#Hemos utilizado la función rand.Seed para inicializar el generador de números aleatorios con la hora actual, lo que ayuda a obtener una secuencia de números aleatorios más diversa.
# Con estas modificaciones, la aplicación escuchará en un puerto aleatorio entre 1001 y 10000 si la variable PUERTO no está definida.

# Hemos agregado una nueva variable de entorno llamada LOGFILE que se utiliza para especificar la ubicación del archivo de log.
# Si la variable LOGFILE no está definida, generamos un nombre de archivo aleatorio en la ruta /tmp/ utilizando math/rand y lo utilizamos como archivo de log.
# Hemos utilizado la función log.Printf para indicar en consola la ubicación del archivo de log utilizado.
# Con estas modificaciones, la aplicación utilizará el archivo de log especificado en la variable LOGFILE si está definida, o bien generará un archivo de log aleatorio en /tmp/ si la variable no está definida. Además, indicará en consola la ubicación del archivo de log utilizado.

*/ 

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {

	log.Printf("App Version 1.0\n")
	// Obtener puerto
	port := os.Getenv("PUERTO")
	if port == "" {
		rand.Seed(time.Now().UnixNano())
		port = fmt.Sprintf("%d", rand.Intn(9000)+1001) // Puerto aleatorio entre 1001 y 10000
	}

	// Obtener ubicación del logfile
	logfile := os.Getenv("LOGFILE")
	if logfile == "" {
		rand.Seed(time.Now().UnixNano())
		logfile = fmt.Sprintf("/tmp/logfile_%d.log", rand.Intn(100000)) // Logfile aleatorio en /tmp/
	}

	// Utilizar el archivo de log
	logFile, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	// Agregar linea con hostname al archivo de inicio
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(logFile, "%s started at %s\n", hostname, time.Now().Format("2006-01-02 15:04:05"))

	// Indicar ubicación del logfile en consola
	log.Printf("Utilizando logfile en %s\n", logfile)

	// Leer archivo de respuesta
	respuesta, err := ioutil.ReadFile("./config/respuesta.txt")
	if err != nil {
		respuesta = []byte("Mensaje de texto generico")
	}

	// Manejador HTTP
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Loggear pedidos de conexion
		log.Printf("%s - Conexion desde %s \n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr )
		w.Write(respuesta)
	})

	// Escuchar en el puerto especificado
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	log.Printf("Escuchando en el puerto %s\n", port)

	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal(err)
	}
}

