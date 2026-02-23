package operate

import (
	"bufio"
	"github.com/compfaculty/dismap/internal/parse"
	"github.com/compfaculty/dismap/pkg/logger"
	"io"
	"net/url"
	"os"
	"strings"
	"sync"
)

func FlagFile(op *os.File, wg *sync.WaitGroup, lock *sync.Mutex, file string, Args map[string]interface{}) {
	thread := Args["FlagThread"].(int)
	f, err := os.Open(file)
	if err != nil {
		logger.Error("There is no " + logger.LightRed(file) + " file or the directory does not exist")
		return
	}
	defer f.Close()

	logger.Info(logger.LightGreen("Batch scan the targets in " + logger.Yellow(file) + logger.LightGreen(", priority network segment")))
	buf := bufio.NewReader(f)

	intSyncThread := 0
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if line != "" {
			if parse.NetJudgeParse(line) {
				FlagNetwork(op, wg, lock, line, Args)
			} else {
				_, urlErr := url.Parse(line)
				if logger.DebugError(urlErr) {
					logger.Error(logger.Red("Unable to parse: " + line))
				} else {
					wg.Add(1)
					intSyncThread++
					go func(line string, Args map[string]interface{}) {
						lock.Lock()
						FlagUrl(op, line, Args)
						lock.Unlock()
						wg.Done()
					}(line, Args)
					if intSyncThread >= thread {
						intSyncThread = 0
						wg.Wait()
					}
				}
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			logger.DebugError(err)
			break
		}
	}
	wg.Wait()
}