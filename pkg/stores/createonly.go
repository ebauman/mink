package stores

import (
	"github.com/acorn-io/mink/pkg/strategy"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/rest"
)

var (
	_ rest.Creater  = (*CreateOnlyStore)(nil)
	_ strategy.Base = (*CreateOnlyStore)(nil)
)

type CreateOnlyStore struct {
	*strategy.CreateAdapter
	*strategy.DestroyAdapter
	*strategy.TableAdapter
}

func NewCreateOnly(scheme *runtime.Scheme, s strategy.Creater) *CreateOnlyStore {
	return &CreateOnlyStore{
		TableAdapter:  strategy.NewTable(s),
		CreateAdapter: strategy.NewCreate(scheme, s),
	}
}
