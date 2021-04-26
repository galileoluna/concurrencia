package contadorGorutines

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"
	"unicode"
)

// Nuestro tamaño de búfer
const chunkSizeGourintes int32 = 32 * 1024

// Esta función abre un archivo de texto, lee los datos en trozos y extrae palabras de él.
// Genera el recuento de palabras una vez hecho.
func ContadorGorutines(filename string, wg *sync.WaitGroup) {

	defer wg.Done()

	// Iniciar un temporizador
	timeStart := time.Now()

	log.Output(2, fmt.Sprintf("comenzando a leer el archivo %s", filename))

	var totalBytesRead, wordsRead int64 // Bytes leídos
	var str, lastWord string            // str es un carácter individual. lastWord es una palabra.

	//  Intenta abrir el archivo
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		log.Output(2, fmt.Sprintf("Error al abrir el archivo %s", filename))
		return
	}

	// Comienza a procesar datos del archivo
	for {

		buf := make([]byte, chunkSizeGourintes) // Crea un búfer
		bytes, err := file.Read(buf)            // Leer datos en el búfer

		// Contamos palabras solo si leemos algunos datos.
		if bytes > 0 {
			totalBytesRead += int64(bytes) // Suma los bytes que hemos leído
			str = string(buf)              // Convertir bytes en un string

			// Bool que nos dice si estamos en una región de espacio en blanco.
			var inSpace bool = true

			// Convierte la cadena en una matriz de runas (unicode) e itera a través de ella un carácter a la vez
			for _, r := range []rune(str) {

				if unicode.IsSpace(r) {
					// Si acabamos de entrar en una región de espacios en blanco y no lo estábamos antes, entonces debemos tener
					// encontró una palabra.
					if !inSpace {
						inSpace = true
						wordsRead += 1 // Incrementar el contador de palabras
						lastWord = ""  // En blanco la palabra lista para la siguiente
					}
					// Si la ejecución no es un espacio en blanco, agréguela a la cadena lastWord. No estamos en una región de espacios en blanco.
				} else {
					inSpace = false
					lastWord += string(r)
				}
			}
		}

		// Termino el archivo
		if err == io.EOF {
			break
		}

		// Hubo otro error
		if err != nil {
			log.Output(2, fmt.Sprintf("Error de lectura %s: %s", filename, err))
			break
		}
	}
	timeEnd := time.Now()
	log.Output(2, fmt.Sprintf("%d\t palabras en total en el archivo %s (%s)", wordsRead, filename, timeEnd.Sub(timeStart)))
}
