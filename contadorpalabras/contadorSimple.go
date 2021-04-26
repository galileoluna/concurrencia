package contadorpalabras

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
	"unicode"
)

const chunkSize int32 = 32 * 1024

// Esta función abre un archivo de texto, lee los datos en trozos y extrae palabras de él.
// Genera el recuento de palabras una vez hecho.
func Contador(filename string) (bytes int64, words int64) {

	timeStart := time.Now()

	var totalBytesRead, wordsRead int64
	var str, lastWord string

	log.Output(2, fmt.Sprintf("comenzando a leer el archivo %s", filename))

	// Intenta abrir el archivo
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		log.Output(2, fmt.Sprintf("Error abriendo %s", filename))
		return
	}

	// Comienza a procesar datos del archivo
	for {

		buf := make([]byte, chunkSize) // Crea un búfer
		bytes, err := file.Read(buf)   // Leer datos en el búfer

		// Contamos palabras solo si leemos algunos datos.
		if bytes > 0 {
			totalBytesRead += int64(bytes) // Suma los bytes que hemos leído
			str = string(buf)              // Convertir bytes en una cadena

			// Booleano que nos dice si estamos en una región de espacio en blanco.
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

		// ¿Llegamos al final del archivo?
		if err == io.EOF {
			break
		}
		// ¿Encontramos un error diferente?
		if err != nil {
			log.Output(2, fmt.Sprintf("Error de lectura %s: %s", filename, err))
			break
		}

	}

	timeEnd := time.Now()
	log.Output(2, fmt.Sprintf("El contador termino  %s segundos leyendo el archivo %s", timeEnd.Sub(timeStart), filename))
	return totalBytesRead, wordsRead

}
