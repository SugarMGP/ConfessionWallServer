package blockService

import (
	"ConfessionWall/app/models"

	"gorm.io/gorm"
)

type BlockRepository struct {
	db *gorm.DB
}

// DeleteBlock 删除某条拉黑记录
func (r *BlockRepository) DeleteBlock(userID uint, blockedBy uint) error {
	return r.db.Where("user_id = ? AND blocked_by = ?", userID, blockedBy).Delete(&models.Block{}).Error
}
