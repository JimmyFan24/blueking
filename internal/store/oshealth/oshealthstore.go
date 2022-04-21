package oshealth

import "bluekinghealth/internal/store"

var (
	cf store.StoreFactory
)

type osHealth struct {
	name string
}

func (o osHealth) OsHealth() store.OsHealthCmd {
	return newOsHealthStore("oshealthstore")
}

func (o osHealth) Health() store.HealthCmd {
	return nil
}

func (o osHealth) Saas() store.SaasCmd {
	return nil
}

func GetOsHealthFactory(name string) (store.StoreFactory, error) {
	cf = &osHealth{name: name}
	return cf, nil
}
