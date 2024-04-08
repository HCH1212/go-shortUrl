package main

import (
	example "gocode/kitex_gen/kitex/example/shorturlservice"
	"log"
)

func main() {
	svr := example.NewServer(new(ShortUrlServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
