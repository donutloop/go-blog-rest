package middelware

import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/dancannon/gorethink"
	"github.com/donutloop/go-blog-rest/utils/context"
)

type RethinkDatabaseSessionMiddleware struct {
	port int
	hostname string
}

func NewRethinkDatabaseSessionMiddleware (hostname string, port int) *RethinkDatabaseSessionMiddleware {
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

	host := gorethink.NewHost(self.hostname, self.port)

	session, err := gorethink.Connect(gorethink.ConnectOpts{
		Address: host.String(),
	})


	if err != nil {
		panic(err)
	}

	return session
}