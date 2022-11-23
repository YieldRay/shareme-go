package db

type tmp struct {
	m map[string]string
}

func (this tmp) Get(namespace string) (content string, err error) {
	content = this.m[namespace]
	err = nil
	return

}

func (this tmp) Set(namespace, content string) bool {
	this.m[namespace] = content
	return true
}

func TmpDB() IDB {
	return tmp{make(map[string]string)}
}
