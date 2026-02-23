package output

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/compfaculty/dismap/internal/flag"
	"github.com/compfaculty/dismap/pkg/logger"
	"os"
	"time"
)

func Open(Args map[string]interface{}) (*os.File, error) {
	if len(Args["FlagOutJson"].(string)) != 0 {
		return openFile(Args["FlagOutJson"].(string))
	}
	return openFile(Args["FlagOutput"].(string))
}

func Write(result map[string]interface{}, output *os.File) {
	if output == nil {
		return
	}
	if result["status"].(string) == "close" {
		return
	}
	if len(flag.OutJson) != 0 {
		result["banner.byte"] = hex.EncodeToString(result["banner.byte"].([]byte))
		result["date"] = time.Now().Unix()
		byteR, err := json.Marshal(result)
		if err != nil {
			logger.DebugError(err)
			return
		}
		writeContent(output, string(byteR))
	} else {
		content := fmt.Sprintf("%s, %s, %s, %s, %s, %s",
			logger.GetTime(),
			result["type"],
			result["protocol"],
			logger.Clean(result["identify.string"].(string)),
			result["uri"],
			result["banner.string"])
		writeContent(output, content)
	}
}

func Close(file *os.File) {
	if file == nil {
		return
	}
	err := file.Close()
	if logger.DebugError(err) {
		logger.Error(fmt.Sprintf("Close file %s exception", logger.Red(file.Name())))
	} else {
		logger.Info("The identification results are saved in " + logger.Yellow(file.Name()))
	}
}

func openFile(name string) (*os.File, error) {
	osFile, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to open file %s", logger.Red(name)))
		return nil, err
	}
	return osFile, nil
}

func writeContent(file *os.File, content string) {
	if file == nil {
		return
	}
	_, err := file.Write([]byte(content + "\n"))
	if logger.DebugError(err) {
		logger.Error(fmt.Sprintf("Write failed: %s", logger.Red(content)))
	}
}
