package logic

import (
	"encoding/json"
	"go.uber.org/zap"
	"pinkacg/logger"
	"pinkacg/models"
	"time"
)

// CreateLog 创建日志
func CreateLog(log *models.Log) error {
	var param models.LogParam
	err := json.Unmarshal([]byte(log.Param), &param)
	if err != nil {
		return err
	}
	logger.Logger.Info("recommend", zap.String("actionTime", time.Now().Format("2006-01-02 15:04:05")), zap.String("readTime", log.ReadTime), zap.Int("category_id", int(log.CategoryId)), zap.Any("param", param))
	return nil
}
