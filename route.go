package maglev

// TODO: This isnt used yet; its a sketch for how we would like to eventually
// implement the individual routes that make up our router (which we would
// prefier a radix trie method)
type Route struct {
	Controller

	Name   string
	Path   string
	Action func() error
}

// TODO: Perhaps a default UseRoutes() function that is called if the app level
// one is undefined to provide a sensible default
//func (f *Framework) UseRoutes(r *router.Mux) {
//	f.Router = f.Routes()
//}
