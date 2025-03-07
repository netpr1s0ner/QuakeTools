package utils

import (
	"fmt"
	"os"
)
import "QuakeAPI/log"

func WriteOutput(content string, filename string) {
	_, err := os.Stat(filename)
	if err == nil {
		err := os.Remove(filename)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0777)
	if file != nil {
		defer file.Close()
	}
	if err != nil {
		log.Log("Create File Error:"+err.Error(), log.ERROR)
		return
	}
	_, err = file.Write([]byte(content))
	if err != nil {
		log.Log("Write File Error:"+err.Error(), log.ERROR)
		return
	}
}
