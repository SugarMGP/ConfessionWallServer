package blockController

import (
	"ConfessionWall/app/services/blockService"
	"ConfessionWall/app/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type DeleteBlockData struct {
	TargetID uint `json:"target_id" binding:"required"`
}

func DeleteBlock(c *gin.Context) {
	id := c.GetUint("user_id")

	var data DeleteBlockData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		zap.L().Error("请求数据绑定失败", zap.Error(err))
		utils.JsonErrorResponse(c, 200506, "参数错误")
		return
	}

	_, err = blockService.GetBlockByID(id, data.TargetID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			zap.L().Debug("拉黑关系不存在", zap.Uint("user_id", id), zap.Uint("target_id", data.TargetID))
			utils.JsonErrorResponse(c, 200508, "拉黑不存在")
		} else {
			zap.L().Error("获取拉黑信息失败", zap.Uint("user_id", id), zap.Uint("target_id", data.TargetID), zap.Error(err))
			utils.JsonInternalServerErrorResponse(c)
		}
		return
	}

	err = blockService.DeleteBlock(id, data.TargetID)
	if err != nil {
		zap.L().Error("删除拉黑失败", zap.Uint("user_id", id), zap.Uint("target_id", data.TargetID), zap.Error(err))
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	// 成功删除拉黑
	zap.L().Info("删除拉黑成功", zap.Uint("user_id", id), zap.Uint("target_id", data.TargetID))
	utils.JsonSuccessResponse(c, nil)

}
