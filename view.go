package webframe

// TODO: This is the default functionality we provide the View object, we
// 'inherit' it through the embedding functionality
// NOTE: One major benefit by having this object, is that we can build them
// beforehand to make caching easier and more logical. A view is built then it
// is rendered, this is a good system. We add to the view as we go through the
// controller and then render it at the end of the process.
type View struct {
	Name string
	// NOTE: This is where we will get all of our model access
	Controller *Controller
	// TODO: Maybe the view itself stores the byte data that is later rendered,
	// and possibly into different formats so we are getting something like:
	//       view.Controller.Model("users").First().Name
	//         or since we have access to controller
	//         view.Model("users").First().Name (but access is mitaged by
	//         controller
	//
	//
	//
	//       view.Render(format.JSON) or
	// todo; well it should be html per say, it should be the bytes to render; and
	// then we can have json.Element too, for example for JSON API
	//HTML []html.Element

	Data []byte
	//bytes []byte
}

// In controller
//RootController() {
//
//	self.NewView(views.Root).

// Writer loop and that

// TODO: Termporary name for experiment
//func (v View) Writer(pipeline chan []byte) {
//	for {
//		data := <-pipeline
//		fmt.Println("writing data from asset pipeline to open connection")
//		//activity.messages++
//
//		// when pipeline feeds down some special end byte message or we have a
//		// separate channel or maybe we have a viarable to check?
//
//	}
//}

// NOTE: These provide useful ways of outputting the data; but we will also
// desire a way to specify if we want the output string or bytes for JSON, HTML,
// XML, etc.
//func (v View) Bytes() (output []byte) {
//	for _, element := range v.HTML {
//		output = append(output, element.Bytes()...)
//	}
//	return output
//}
//
//func (v View) String() (output string) {
//	for _, element := range v.HTML {
//		output += element.String()
//	}
//	return output
//}
