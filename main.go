package main

import (
	fileservice "fileService/kitex_gen/fileService/fileservice"
	"log"
)

func main() {
	svr := fileservice.NewServer(new(FileServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
