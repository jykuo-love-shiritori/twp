package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jykuo-love-shiritori/twp/minio"
)

func main() {
	mc, err := minio.NewMINIO()
	if err != nil {
		log.Fatal(err)
	}
	fileName := "test.jpg"
	path := "./test/minio/" + fileName
	_, err = mc.PutFileByPath(context.Background(), path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MINIO insert success")
	err = mc.RemoveFile(context.Background(), fileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MINIO delete success")

}
