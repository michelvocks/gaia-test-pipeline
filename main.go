package main

import (
	"log"
	"time"

	sdk "github.com/gaia-pipeline/gosdk"
)

func CreateUser() error {
	log.Println("CreateUser has been executed!")

	// lets sleep to simulate that we do something
	time.Sleep(15 * time.Second)

	return nil
}

func MigrateDB() error {
	log.Println("MigrateDB has been executed!")

	// lets sleep to simulate that we do something
	time.Sleep(15 * time.Second)

	return nil
}

func CreateNamespace() error {
	log.Println("CreateNamespace has been executed!")

	// lets sleep to simulate that we do something
	time.Sleep(15 * time.Second)

	return nil
}

func CreateDeployment() error {
	log.Println("CreateDeployment has been executed!")

	// lets sleep to simulate that we do something
	time.Sleep(15 * time.Second)

	return nil
}

func CreateService() error {
	log.Println("CreateService has been executed!")

	// lets sleep to simulate that we do something
	time.Sleep(15 * time.Second)

	return nil
}

func CreateIngress() error {
	log.Println("CreateIngress has been executed!")

	// lets sleep to simulate that we do something
	time.Sleep(15 * time.Second)

	return nil
}

func Cleanup() error {
	log.Println("Cleanup has been executed!")

	// lets sleep to simulate that we do something
	time.Sleep(15 * time.Second)

	return nil
}

func main() {
	// Serve
	sdk.Serve(jobs)
}
