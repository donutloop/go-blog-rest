package middelware


import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/dancannon/gorethink"
	"github.com/donutloop/go-blog-rest/library"
)

type RethinkDatabaseSessionMiddleware struct {
}

func NewRethinkDatabaseSessionMiddleware () *RethinkDatabaseSessionMiddleware {
	return &RethinkDatabaseSessionMiddleware{}
}

// MiddlewareFunc returns a HandlerFunc that implements the middleware.
func (mw *RethinkDatabaseSessionMiddleware) MiddlewareFunc(h rest.HandlerFunc) rest.HandlerFunc {
	return func(w rest.ResponseWriter, r *rest.Request) {

		session := initDatabaseSession()

		context.Set(r, "dbSession", session)

		defer func() {
			session.Close()
			context.Delete(r, "dbSession")
		}()

		h(w, r)
	}
}

func initDatabaseSession() *gorethink.Session {
	session, err := gorethink.Connect(gorethink.ConnectOpts{
		Address: "127.0.1.1",
	})

	if err != nil {
		panic(err)
	}

	return session
}