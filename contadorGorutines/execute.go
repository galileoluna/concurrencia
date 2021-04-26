package contadorGorutines

import (
	"fmt"
	"os"
	"sync"
	"time"

	"log"
)

func Execute() {

	log.SetFlags(log.Ldate | log.Lmicroseconds)

	if len(os.Args) < 2 {
		log.Output(2, fmt.Sprintf("Se esperaba al menos un nombre de archivo."))
		os.Exit(1)
	}

	filenames := os.Args[1:]

	timeStart := time.Now()

	// **** WaitGroup ****
	// Un WaitGroup es útil cuando tienes un solo canal compartido entre
	// goroutines. ¿Cómo se sabe cuando se terminan todas las gorutines?
	//
	// Cuando inicias una goroutine que usa un canal compartido, agregas () al WaitGroup.
	// Una vez finalizado el procesamiento, llama a Done (). En algún lugar tendrás que esperar () hasta
	// todos los Add () tienen un Done () correspondiente. Luego, Wait () se desbloqueará y podrá
	// cierra tu canal compartido.
	wg := sync.WaitGroup{}

	// Iterate through files from the command line arguments
	for i := 0; i < len(filenames); i++ {

		// **** WaitGroups ***
		// ¿Cómo sabemos cuando los gorutines terminan de contar palabras? Debemos usar un WaitGroup.
		// Agregamos () las rutinas que comenzamos, que luego llaman a Done () cuando se completan.
		wg.Add(1)

		// Execute wordCounter. We pass it he filename we're processing and a WaitGroup to send Done() to
		go ContadorGorutines(filenames[i], &wg)
	}

	// Esta llamada se bloquea hasta que todas las goroutines estén Done ()
	wg.Wait()

	timeEnd := time.Now()
	log.Output(2, fmt.Sprintf("Tiempo total %s", timeEnd.Sub(timeStart)))

}
