package store

type StatusCmd interface {
	PaasStatusCmd() string
	CmdbStatusCmd() string
}
