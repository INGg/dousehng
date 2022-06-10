package repository

import (
	"demo1/model/entity"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"sync"
)

type RelationDao struct {
}

var (
	relationDao  *RelationDao
	relationOnce sync.Once
)

func NewRelationDAO() *RelationDao {
	relationOnce.Do(func() {
		relationDao = &RelationDao{}
	})
	return relationDao
}

//	AddRelation 将 用户id 和被其关注人的id 插入表中 relation
func (r *RelationDao) AddRelation(FollowerId uint, AuthorId uint) error {
	res := db.Create(&entity.Relation{
		UserID:   AuthorId,
		FollowID: FollowerId,
	})
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		zap.L().Error("Add Relation error")
		return res.Error
	}
	zap.L().Info("insert relation success")
	return nil
}

// DeleteRelation 根据 userid followerId 删除对应记录
func (r *RelationDao) DeleteRelation(FollowerId uint, AuthorId uint) error {
	res := db.Where(&entity.Relation{
		UserID:   AuthorId,
		FollowID: FollowerId,
	}).Delete(&entity.Relation{})

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		zap.L().Error("Delete relation error")
		return res.Error
	}
	zap.L().Info("Delete relation success")
	return nil
}

/* 查询当前用户的粉丝(id)
 */
//func (r *RelationDao) QueryFollowIdByAuthorId(AuthorId uint, FollowerIdList *[]Relation) error {
//	res := db.Model(&Relation{}).Where("author_id = ?", AuthorId).Find(FollowerIdList)
//	// QueryFollowIdByUserID 查询当前用户的关注列表(id)
//}
func (r *RelationDao) QueryFollowIdByUserID(uid uint, RelationList *[]entity.Relation) error {
	res := db.Model(&entity.Relation{}).Where("user_id = ?", uid).Find(RelationList)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

/*查询当前用户的关注(id)
 */
//func (r *RelationDao) QueryAuthorIdByFollowId(FollowerId uint, AuthorIdList *[]Relation) error {
//
//	res := db.Model(&Relation{}).Where("follower_id = ?", FollowerId).Find(AuthorIdList)
//}

// QueryUsersIDByFollowId 查询当前用户的粉丝(id)
func (r *RelationDao) QueryUsersIDByFollowId(FollowerId uint, relationList *[]entity.Relation) error {
	res := db.Model(&entity.Relation{}).Where("follow_id = ?", FollowerId).Find(relationList)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

/*
判断relation库中是否已经存在数据
*/
//func (r *RelationDao) IsFollow(UserA uint, UserB uint) (bool, error) {
//	var FollowList []Relation
//	res := db.Model(&Relation{}).Where("follower_id=", UserA).Where("author_id", UserB).Find(&FollowList)
//	if res.Error != nil {
//		return false, res.Error
//	}
//	if len(FollowList) == 0 {
//		return false, nil
//	} else {
//		return true, nil
//	}
//}
func (r *RelationDao) QueryAFollowB(Auid uint, Buid uint) bool {
	res := db.Model(&entity.Relation{}).Where("user_id = ?", Auid).Where("follow_id = ?", Buid)
	if res.Error != nil {
		return false
	}
	return true
}
