package gcache

import (
	"github.com/linthan/gcache/internal"
	"github.com/linthan/gcache/option"
	"github.com/jinzhu/gorm"
)

// Plugin .
type Plugin interface {
	//清库
	FlushDB() error
	//删除模型缓存
	DeleteModel(model interface{}, primarys ...interface{}) error
	//删除查询缓存
	DeleteSearch(model interface{}) error
	//insert select update delete 都会跳过缓存处理
	SkipCache() *gorm.DB

	//join 和 子查询， 需要传入模型。
	CreateRelative(...interface{}) *gorm.DB
	SetRelative(*gorm.DB, ...interface{}) *gorm.DB

	//tag
	CreateTag(...interface{}) *gorm.DB
	SetTag(*gorm.DB, ...interface{}) *gorm.DB
	Debug()
}

// AttachDB .
func AttachDB(db *gorm.DB, opt *option.DefaultOption, redisOption *option.RedisOption) Plugin {
	return internal.InjectGorm(db, opt, redisOption)
}
