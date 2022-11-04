package main

import (
	"cities-1/pkg/etc"
	"cities-1/pkg/http"
	"cities-1/pkg/store"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const (
	filename = "cities.csv"
)

func main() {
	st := store.NewStore()
	err := st.LoadFromCsv(filename)
	if err != nil {
		log.Fatal("Unable to parse file CSV", err)
		return
	}

	//ctrl-c handler
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("CTRL-C were pressed, save data to file...")
		err := st.SaveToCsv(filename)
		if err != nil {
			os.Exit(1)
		} else {
			log.Println("CTRL-C were pressed, data has been saved")
			os.Exit(0)
		}
	}()
	http.Router(etc.HostPortResolver(), st)
}
