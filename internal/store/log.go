package store

type LogCmd interface {
	PaasCheck() string
	CmdbCheck() string
}
