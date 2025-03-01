package imps

import (
	xx "github.com/blevesearch/bleve-mapping-ui"
	"github.com/gorilla/mux"
	"net/http"
)

type GorillaMuxImp struct {
	router *mux.Router
	path   string
	f      func(http.ResponseWriter, *http.Request)
}

func NewGorillaMuxImp(r *mux.Router) *GorillaMuxImp {
	return &GorillaMuxImp{router: r}
}
func (r *GorillaMuxImp) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) xx.MethodsFace {
	r.path = path
	r.f = f
	var k xx.MethodsFace = r
	return k
}

func (r *GorillaMuxImp) Methods(methods ...string) {
	r.router.HandleFunc(r.path, r.f).Methods(methods...)
}
