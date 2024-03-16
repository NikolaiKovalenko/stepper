package main

import (
	"context"
	"fmt"
	"log"

	"github.com/NikolaiKovalenko/stepper"
	mongoEngine "github.com/NikolaiKovalenko/stepper/engines/mongo"
	"github.com/NikolaiKovalenko/stepper/examples"
)

func main() {
	db, err := examples.CreateMongoDatabase("stepepr")

	if err != nil {
		log.Fatal(err)
	}

	e := mongoEngine.NewMongoWithDb(db)
	s := stepper.NewService(e)

	s.RegisterJob(context.Background(), &stepper.JobConfig{
		Name:    "log-job",
		Pattern: "@every 10s",
	}, func(ctx stepper.Context) error {
		fmt.Println("wake up the log-job")

		ctx.CreateSubtask(stepper.CreateTask{
			Name: "log-subtask",
			Data: []byte("Hello 1 subtask"),
		})

		ctx.CreateSubtask(stepper.CreateTask{
			Name: "log-subtask",
			Data: []byte("Hello 2 subtask"),
		})

		return nil
	}).OnFinish(func(ctx stepper.Context, data []byte) error {
		fmt.Println("success job log-job")

		return nil
	})

	s.TaskHandler("log-subtask", func(ctx stepper.Context, data []byte) error {
		fmt.Println("message from subtask:", string(data))
		return nil
	})

	if err := s.Listen(context.Background()); err != nil {
		log.Fatal(err)
	}
}
