package contadorpalabras

import (
	"fmt"
	"log"
	"os"
	"time"
)

var logger *log.Logger

func Execute() {

	log.SetFlags(log.Ldate | log.Lmicroseconds)

	if len(os.Args) < 2 {
		log.Output(2, fmt.Sprintf("Se esperaba al menos un nombre de archivo.."))
		os.Exit(1)
	}

	filenames := os.Args[1:]

	timeStart := time.Now()

	for i := 0; i < len(filenames); i++ {

		_, w := Contador(filenames[i])
		log.Output(2, fmt.Sprintf("%d\tpalabras en  %s", w, filenames[i]))
	}

	timeEnd := time.Now()
	log.Output(2, fmt.Sprintf("Tiempo total %s", timeEnd.Sub(timeStart)))

}
