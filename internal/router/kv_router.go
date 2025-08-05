package router

func (r *Router) kvRouter() {
	_ = r.app.Group("/kv")

}
