package maglev

//type CollectionRouter interface {
//	Index(b http.ResponseWriter, request *http.Request)
//	Show(b http.ResponseWriter, request *http.Request)
//
//	New(b http.ResponseWriter, request *http.Request)
//	Create(b http.ResponseWriter, request *http.Request)
//
//	Edit(b http.ResponseWriter, request *http.Request)
//	Update(b http.ResponseWriter, request *http.Request)
//	Delete(b http.ResponseWriter, request *http.Request)
//}

// TODO: We need to incldue a way to specify a model, and a way to whitelist
// field interaction from the controller.
type Controller struct {
	Framework *Framework

	Name string

	Views []*View

	BeforeAction []func()
	AfterAction  []func()
}

// TODO: Concept if we go with a gin like solution
//func (c Controller) Root(c framework.Context) {
//	p := c.Param("test")
//
//	c.HTML(views.Root())
//	c.JSON(someJSON)
//	c.Render()
//}

//func ServeHTTP() {
//
//}

// TODO: The controller could store these views to establish our cache system
// too which will need to be fairly complex for big applications
func (c Controller) View(name string) *View {
	for _, view := range c.Views {
		if len(view.Name) == len(name) && view.Name == name {
			return view
		}
	}
	return nil
}

func (c Controller) NewView(name string) *View {
	if c.View(name) == nil {
		c.Views = append(c.Views, &View{
			Name: name,
		})
		return c.Views[len(c.Views)-1]
	}
	return nil
}
