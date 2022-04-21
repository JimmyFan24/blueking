package store

//存储层抽象接口

type StoreFactory interface {
	Health() HealthCmd
	Saas() SaasCmd
	OsHealth() OsHealthCmd
}
