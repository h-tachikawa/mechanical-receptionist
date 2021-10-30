package adapter

type NotifierAdapter interface {
	Notify(message string) error
}
