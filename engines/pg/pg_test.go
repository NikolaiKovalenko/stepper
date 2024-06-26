package pg

import (
	"context"
	"log"
	"testing"

	"github.com/NikolaiKovalenko/stepper"
	"github.com/NikolaiKovalenko/stepper/tests"
)

func TestPG(t *testing.T) {
	tests.Run(t, func() stepper.Stepper {
		pgEngine, err := NewPG("postgres://postgres:test@localhost:5432/postgres")
		if err != nil {
			log.Fatal(err)
		}

		if err := pgEngine.Init(context.Background()); err != nil {
			log.Fatal(err)
		}

		return stepper.NewService(pgEngine)
	})
}
