package irepository

type IRedisRepository interface {
	InsertIp(ip string) error
	GetIp() (string, error)
}
