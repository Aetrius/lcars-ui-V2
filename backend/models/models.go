package models

import (
	"fmt"
	"math"
	"os"
	"time"

	"github.com/mackerelio/go-osstat/cpu"
)

type model struct {
}

type Model interface {
	// Return read only channel that has information about % current cpu
	GetLiveCpuUsage() (<-chan Cpu, error)
}

func NewModel() Model {
	return &model{}
}

type Cpu struct {
	PercentageUsage int `json:"cpuPercentageUsage"`
}

// Consider this model responsible for puling data from backend and serving it to controller
func (m *model) GetLiveCpuUsage() (<-chan Cpu, error) {
	ch := make(chan Cpu)
	go WriteValues(ch)

	return ch, nil
}

// Generates random values and writes to passed on channel
func WriteValues(ch chan<- Cpu) error {
	ticker := time.NewTicker(1 * time.Second)

	for {
		before, err := cpu.Get()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)

		}
		time.Sleep(time.Duration(1) * time.Second)
		after, err := cpu.Get()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)

		}
		total := float64(after.Total - before.Total)

		<-ticker.C
		//newVal := Cpu{PercentageUsage: rand.Intn(100)}
		newVal := Cpu{PercentageUsage: int(math.Round(float64(after.System-before.System) / total * 100))}
		ch <- newVal
	}
}
