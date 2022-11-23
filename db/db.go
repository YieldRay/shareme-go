package db

type IDB interface {
	Get(namespace string) (string, error)
	Set(namespace, content string) bool
}
