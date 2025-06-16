package cmd

import (
	"fmt"
	"gin-project/internal/config"
)

func initCommonPart(confFilePath string) error {
	_, err := config.InitConfig(confFilePath)
	if err != nil {
		return fmt.Errorf("[ERROR] load config error: %v\n", err)
	}

	// todo 日志
	// todo 调用链

	return nil
}
