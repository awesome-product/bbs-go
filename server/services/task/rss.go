package task

import (
	"github.com/mlogclub/bbs-go/services"
)

// 生成rss
func RssTask() {
	services.ArticleService.GenerateRss()
	services.TopicService.GenerateRss()
	services.ProjectService.GenerateRss()
}
