package component

type Component interface {
	Check() map[string]string
	Status() map[string]string
	Log() map[string]string
}
