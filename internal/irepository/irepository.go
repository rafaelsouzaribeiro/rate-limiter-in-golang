package irepository

type IRedisRepository interface {
	InsertIp(ip string) error
	IncreaseIPRequest(ip string, seconds int) (int64, error)
}
