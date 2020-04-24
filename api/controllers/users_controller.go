package controllers

import "log"

// this function will process data and save it in either/redis or database
func ProcessRequest(data []byte) {
	log.Printf("proccessed data : %s", data)
}
