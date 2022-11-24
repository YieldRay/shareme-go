package db

type tmp_db struct {
	m map[string]string
}

func (this tmp_db) Get(namespace string) (content string, err error) {
	content = this.m[namespace]
	err = nil
	return

}

func (this tmp_db) Set(namespace, content string) bool {
	this.m[namespace] = content
	return true
}

func TmpDB() IDB {
	return tmp_db{make(map[string]string)}
}
