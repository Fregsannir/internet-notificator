package main

import (
	"log"
	"net"
	"runtime"
	"time"

	"github.com/gen2brain/beeep"
)

func main() {
	beeep.Beep(60, 200)
	var (
		errNotification     bool = false
		successNotification bool = false
	)
	log.Printf("Your operation system: %s\n", runtime.GOOS)
	for {
		_, err := net.Dial("tcp", "google.com:https")
		if err != nil {
			log.Println("Connection refused...")
			if !errNotification {
				beeep.Alert("Internet Notification", "Connection refused", "assets/warning.png")
				errNotification = true
				successNotification = false
			}
			time.Sleep(1 * time.Second)
		} else {
			log.Println("Connection is stable")
			if !successNotification {
				beeep.Alert("Internet Notification", "Connection is stable", "assets/success.png")
				successNotification = true
				errNotification = false
			}
			time.Sleep(1 * time.Second)
		}
	}
}
