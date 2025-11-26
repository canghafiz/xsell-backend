package services

type RedisService interface {
	SetData(key string, value interface{}, durationMinute int) error
	GetData(key string) (interface{}, error)
	DeleteData(key string) error
	DeleteAllWithPrefix(prefix string) error
	ClearAllData() error
}
