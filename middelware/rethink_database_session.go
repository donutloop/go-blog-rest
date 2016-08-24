package middelware


import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/dancannon/gorethink"
	"github.com/donutloop/go-blog-rest/library"
)

type RethinkDatabaseSessionMiddleware struct {
	port string
	hostname string
}

func NewRethinkDatabaseSessionMiddleware (hostname string, port string) *RethinkDatabaseSessionMiddleware {
	return &RethinkDatabaseSessionMiddleware{
		hostname: hostname,
		port: port}
}

// MiddlewareFunc returns a HandlerFunc that implements the middleware.
func (self *RethinkDatabaseSessionMiddleware) MiddlewareFunc(h rest.HandlerFunc) rest.HandlerFunc {
	return func(w rest.ResponseWriter, r *rest.Request) {

		session := self.initDatabaseSession()
		context.Set(r, "dbSession", session)

		defer func() {

			session.Close()
			context.Delete(r, "dbSession")
		}()

		h(w, r)
	}
}

func (self *RethinkDatabaseSessionMiddleware) initDatabaseSession() *gorethink.Session {
	session, err := gorethink.Connect(gorethink.ConnectOpts{
		Address: self.hostname + ":" + self.port,
	})

	if err != nil {
		panic(err)
	}

	return session
}