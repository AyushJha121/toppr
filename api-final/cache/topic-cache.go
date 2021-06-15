package cache

import "api-final/entity"

type TopicCache interface {
	Set(key string, value entity.Topic)
	Get(key string) entity.Topic
}
