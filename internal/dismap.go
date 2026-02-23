package internal

import (
	"github.com/compfaculty/dismap/internal/flag"
	"github.com/compfaculty/dismap/internal/operate"
	"github.com/compfaculty/dismap/internal/output"
	"sync"

	"github.com/compfaculty/dismap/pkg/logger"
)


func which(Args map[string]interface{}, wg *sync.WaitGroup, lock *sync.Mutex) {
	op, err := output.Open(Args)
	if err != nil || op == nil {
		logger.Error("Failed to open output file")
		return
	}
	defer output.Close(op)

	address := Args["FlagNetwork"].(string)
	if address != "" {
		operate.FlagNetwork(op, wg, lock, address, Args)
		return
	}

	uri := Args["FlagUrl"].(string)
	if uri != "" {
		operate.FlagUrl(op, uri, Args)
		return
	}

	file := Args["FlagFile"].(string)
	if file != "" {
		operate.FlagFile(op, wg, lock, file, Args)
		return
	}
}

func DisMap() {
	Args := flag.Flags()
	wg := &sync.WaitGroup{}
	lock := &sync.Mutex{}

	information()
	which(Args, wg, lock)
	logger.Info("Identification completed and ended")
}
