package model

import (
	"demo1/model/entity"
	"mime/multipart"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	entity.Video
}

// -- feed --
type FeedRequest struct {
	FromUserID uint
	LatestTime int64  `json:"latest_time"`
	Token      string `json:"token"`
}

type FeedResponse struct {
	Response
	VideoList *[]Video `json:"video_list"`
	NextTime  int64    `json:"next_time"`
}

// -- publish --
type PublishActionRequest struct {
	Token    string                `json:"token" form:"token"`
	Data     *multipart.FileHeader `json:"data" form:"data"`
	Title    string                `json:"title" form:"title"`
	UserName string
	UserID   uint
}

type PublishActionResponse struct {
	Response
}

type PublishListRequest struct {
	Token      string `json:"token" form:"token"`
	UserID     uint   `json:"user_id" form:"user_id"`
	FromUserID uint
}

type PublishListResponse struct {
	Response
	VideoList *[]Video `json:"video_list"`
}

// -- user --
type UserInfoRequest struct {
	Token    string `json:"token" form:"token"`
	UserID   uint   `json:"user_id" form:"user_id" `
	UserName string
}

type UserInfoResponse struct {
	Response
	User entity.User `json:"user"`
}

type UserLoginRequest struct {
	UserName string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type UserLoginResponse struct {
	Response
	UserID uint   `json:"user_id"`
	Token  string `json:"token"`
}

type UserRegisterRequest struct {
	UserName string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type UserResisterResponse struct {
	Response
	UserID uint   `json:"user_id"`
	Token  string `json:"token"`
}

// -- comment --
// Comment 这里的comment和repository里的comment本质是没有太大的区别，这里只是为了请求好返回
type Comment struct {
	ID         uint        `json:"id"`
	User       entity.User `json:"user"`
	Content    string      `json:"content"`
	CreateDate string      `json:"create_date"` // 评论发布日期，格式 mm-dd
}

type CommentActionRequest struct {
	UserID      uint
	Token       string `json:"token" form:"token"`
	VideoID     uint   `json:"video_id" form:"video_id"`
	ActionType  uint8  `json:"action_type" form:"action_type"`
	CommentText string `json:"comment_text,omitempty" form:"comment_text"` // 用户填写的评论内容，在action_type=1的时候使用
	CommentID   uint   `json:"comment_id,omitempty" form:"comment_id"`     // 要删除的评论id，在action_type=2的时候使用
	UserName    string
}

type CommentActionResponse struct {
	Response
	Comment
}

type CommentListRequest struct {
	Token      string `json:"token" form:"token"`
	VideoID    uint   `json:"video_id" form:"video_id"`
	FromUserID uint
}

type CommentListResponse struct {
	Response
	CommentList *[]Comment `json:"comment_list"`
}

// --Relation--
type FollowActionRequest struct {
	UserID     uint   `json:"user_id" form:"user_id" `
	Token      string `json:"token" form:"token"`
	ToUserID   uint   `json:"to_user_id" form:"to_user_id"`
	ActionType uint   `json:"action_type" form:"action_type"`
}

type FollowActionResponse struct {
	Response
}

type UserFollowListRequest struct {
	UserID     uint   `json:"user_id" form:"user_id" `
	Token      string `json:"token" form:"token"`
	FromUserID uint
}

type UserFollowListResponse struct {
	Response
	UserList *[]entity.User `json:"user_list"`
}

type UserFollowerListRequest struct {
	UserID     uint   `json:"user_id" form:"user_id" `
	Token      string `json:"token" form:"token"`
	FromUserID uint
}

type UserFollowerListResponse struct {
	Response
	UserList *[]entity.User `json:"user_list"`
}

type UserFavoriteRequest struct {
	UserID     uint
	Token      string `json:"token" form:"token"`
	VideoID    uint   `json:"video_id" form:"video_id"`
	ActionType int32  `json:"action_type" form:"action_type"`
}

type UserFavoriteResponse struct {
	Response
}

type UserFavoriteListRequest struct {
	UserID     uint   `json:"user_id" form:"user_id"`
	Token      string `json:"token" form:"token"`
	FromUserID uint
}

type UserFavoriteListResponse struct {
	Response
	VideoList *[]entity.Video `json:"video_list"`
}
