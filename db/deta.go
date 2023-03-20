package db

import (
	"fmt"
	"os"

	"github.com/deta/deta-go/deta"
	"github.com/deta/deta-go/service/base"
)

type deta_base struct {
	base *base.Base
}

type record struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (this deta_base) Get(namespace string) (content string, err error) {
	r := record{}
	if err := this.base.Get(namespace, &r); err != nil {
		if err == deta.ErrNotFound {
			return "", nil
		} else {
			return "", err
		}
	}
	return r.Value, nil
}

func (this deta_base) Set(namespace, content string) bool {
	r := &record{
		Key:   namespace,
		Value: content,
	}
	_, err := this.base.Put(r)
	return err == nil
}

func Deta(projectKey, baseName string) IDB {
	d, err := deta.New(deta.WithProjectKey(projectKey))
	if err != nil {
		fmt.Fprint(os.Stderr, "failed to init new Deta instance")
		panic(err)
	}

	db, err := base.New(d, baseName)
	if err != nil {
		fmt.Fprint(os.Stderr, "failed to init new Base instance")
		panic(err)
	}
	return deta_base{db}
}
