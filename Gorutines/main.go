package Gorutines

import (
	"fmt"
	"time"
)

func HolaMundo() {
	// creamos un canal de tipo string
	ch := make(chan string)

	// generamos una funcion anonima con una gorutine
	go func() {
		time.Sleep(time.Second)
		// enviamos hola mundo a
		ch <- "Hola mundo"
	}()
	// leemos los datos del canal y los guardamos en holamundo
	holamundo := <-ch
	fmt.Printf(holamundo)
}
