package module

import (
	"database/sql"
	"time"
)

type Articles struct {
	Id            int       `orm:"id" json:"id"`
	Title         sql.NullString    `orm:"title" json:"title"`
	Content       sql.NullString    `orm:"content" json:"content"`
	PublishedAt   time.Time `orm:"published_at" json:"published_at"`
	Status        int       `orm:"status" json:"status"`
	Summary       string    `orm:"summary" json:"summary"`
	StafferId     int       `orm:"staffer_id" json:"staffer_id"`
	Cover         string    `orm:"cover" json:"cover"`
	Sort          int       `orm:"sort" json:"sort"`
	CreatedAt     time.Time `orm:"created_at" json:"created_at"`
	UpdatedAt     time.Time `orm:"updated_at" json:"updated_at"`
	IsMaterial    int       `orm:"is_material" json:"is_material"`
	CleanContent  string    `orm:"clean_content" json:"clean_content"`
	MvUrl         string    `orm:"mv_url" json:"mv_url"`
	VUrl          string    `orm:"v_url" json:"v_url"`
	Audio         string    `orm:"audio" json:"audio"`
	SourceMedia   string    `orm:"source_media" json:"source_media"`
	AuthorId      int       `orm:"author_id" json:"author_id"`
	LikeNum       int       `orm:"like_num" json:"like_num"`
	CommentsCount int       `orm:"comments_count" json:"comments_count"`
	Type          string    `orm:"type" json:"type"`
	IsDeleted     int       `orm:"is_deleted" json:"is_deleted"`
	HasImg        int       `orm:"has_img" json:"has_img"`
	LikesCount    int       `orm:"likes_count" json:"likes_count"`
	DefaultCover  int       `orm:"default_cover" json:"default_cover"`
	ReadCount     int       `orm:"read_count" json:"read_count"`
	IsHidden      int       `orm:"is_hidden" json:"is_hidden"` // 是否在列表中隐藏
	CrawlerId     int       `orm:"crawler_id" json:"crawler_id"`
	VideoId       int       `orm:"video_id" json:"video_id"`             // 视频Id
	IsOriginal    int       `orm:"is_original" json:"is_original"`       // 是否是原创
	OriginalLevel int       `orm:"original_level" json:"original_level"` // 原创类别
	PenLecturerId int       `orm:"pen_lecturer_id" json:"pen_lecturer_id"`
	OpinionTagId  int       `orm:"opinion_tag_id" json:"opinion_tag_id"`
	AgentId       int       `orm:"agent_id" json:"agent_id"`
	IsLocked      int       `orm:"is_locked" json:"is_locked"`     // 是否加锁
	AuditState    int       `orm:"audit_state" json:"audit_state"` // 审核状态
	UpParse       string    `orm:"up_parse" json:"up_parse"`       // 主题风向标的涨停解析
	ExpertId      int       `orm:"expert_id" json:"expert_id"`     // 关联专家表id
}

func (*Articles) getTableName() string {
	return "articles"
}
