package core

import (
	"path/filepath"
	"time"

	"github.com/JsonLee12138/json-server/pkg/configs"
	"github.com/JsonLee12138/json-server/pkg/utils"
	"gopkg.in/natefinch/lumberjack.v2"
)

type writeSyncer struct {
	config configs.LogConfig
	level  string
}

func (s *writeSyncer) Write(p []byte) (n int, err error) {
	dirPath := filepath.Join(s.config.Director, time.Now().Format("2006-01-02"))
	fileName := filepath.Join(dirPath, s.level+".log")
	utils.CreateDir(dirPath)
	lumberjackWriter := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    s.config.MaxSize,    // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: s.config.MaxBackups, // 日志文件最多保存多少个备份
		MaxAge:     s.config.MaxAge,     // 文件最多保存多少天
		Compress:   s.config.Compress,   // 是否压缩
	}
	return lumberjackWriter.Write(p)
}

func (s *writeSyncer) Sync() error {
	return nil
}
