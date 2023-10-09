# Webframe: HTML Package

**Notes**
Since UTF-8 is a functional requirement for Webframe localization, lets just
jump ahead, only think about UTF-8 and HTML5. Ignore breaking out HTML4 and
ASCII. 

Lets make this modern, and slim because we avoid worrying about legacy stuff
that we simply are moving past. 


Each sub-component of webframe is inteded to when combined provide complete
acceess to HTML5 via Go. The generlic topic of HTML5 is broad but specific
features include programmtic HTML genheration via Go to avoid both the
insecurity of mucking around with HTML elements directly.

Primary functionality includes but is not neccessarily limited to HTML (and
Markdown (for userinput)) parsing and  HTML code geenration enabling developers
to easily marshal between HTML strings and Go objects easily providing a solid
framework for building allow list (or "whitelisting" based validation), 
templating, rails helper-style re-usable HTML elements, "simple_form gem"-like
form generation functionality to simplify and increase the security of every
form element, and a collection of other features that may eventually warrant
thier own package if it grows too big. Or the simpler version of the html
package found in early iteration of html.go (found in the git history) may
also be desirable as its own package.

Importantly, rendered HTML will only be successfully rendered if it has been
explictedly added to a constant DB style by allowing it or whitelisting it. This
provides far greater security than the currently conventional "blacklisting" 

### HTML Package Terminology
*(Should it be sign? We have moved to PlusSign and EqualSign to avoid namespace
collisions already)*

**Symbol**
ASCII support has not ended because of the newly available UTF-8, and Webframe
responds by mapping all the ASCII symbols and a user-defined set of UTF-8 symbols
to a type which is based on int for fast validation based on allow lists.

These **Symbol(s)** include the `lessThanSign` (`<`) as used in the `<html>`
**Tag**.

Mapping, and creating allow list of ASCII/UTF8 symbols allows for
incredibly fine grain control over the HTML rendering. Not only are these 
providing the HTML numbers and `&nbsp;` like conversions but also per tag, or
per attribute allow list of symbols using immutable constants stored in an 
optimized form (not yet, but eventually should store as binary).

>> "To strip newlines from a string, remove any U+000A LF and U+000D CR code
>> points from the string." - [HTML5 Spec:
>> ASCII](https://infra.spec.whatwg.org/#ascii-digit)


**HTML Element**
**Tag** is a subcomponent of **HTML Element** (which is the building block of
HTML). The **Tag** has two (2) types: (1) *Open* **Tag** such as `<html>`, and
(2) *Close* **Tag** as in `</html>`. 

**Attribute**
Both the opening and close tag, and everything contained
inbetween is the **HTML Element**. An **HTML Element** is customized using
**Attribute**, these are key/value data stored in the *Open* **Tag** of the
**HTML Element**. 


### Core Functionality
Below is a evolving list of functional requirements of the HTML package used by
the Webframe Go web application framework. 

  * Parse HTML, Markdown, and (HAML?) and build corresponding Go objects which fully represent parsed data

  * Render only allowed strings for an allowed set of attributes defined per tag or globally, return byte sequence for inserting dirrectly into byte streams, and strings for writing HTML files
  
  * Ability to filter/transform (minify, find/replace, remove XSS/SQLi, remove or replace substrings in disallowed list

  * (May eventually move to own package) Templating
    * Easily define remote location of CSS framework, and it will download and place contents inline before developer CSS overrides

  * Conviences for creating forms heavily inspired by the Ruby (specifically for Rails) gem "simple_form" providing modern
way to handle forms, present multiple errors, highlight fields that have errors and display each error message below the input. A collection of input types done in pure CSS or JS with CSS fallback for date input, date/time input, color input, time input, email input, phone input, currency input, and so on. 

  * Should have conviences for essentially all the things in a standard CSS framework. Use the documentation of a few different CSS frameworks to get a good list of expected helpers (the term Rails uses for re-usable code widgets). 

  * Easy to use and intuitive HTML creation from Go, below is a quick example of creating unordered list with a few items and class defintions, arbitrary attribute defintion on an element, and all the content betweent he tags ends up between the parenthesis:

```
UL(
  LI(Link("Home", "/")).Class("active", "headitem").Attribute("is-active", "true"),
	LI(Link("Projects", "/projects"),
  ...
).Class("headmenu").Id("mainmenu")
```

A full example including the `<html>` and `<body>` tags required to create a docuemnt can be accomplished very simply without needing to concat any strings:

```
HTML(
	Body(
		H1("Hello World"),
		Br().Class("spacer"),
		P(
			"A simple paragraph about the ",
			Span("Application.").Class("text-warning"),
		),
	),
)
```

This allows for easy single binary executables without needing to have a packer which is convienent. 


### Model for code generation for later "scaffolding"-like functionality
One of the easily overlooked features that makes Ruby's web application
framework Rails so powerful is the ability to generate or scaffold common MVC
design patterns based on easily definable templates. This powerful feature is a
large part responsible for the reputation Rails has for being a development
framework that allows users to increase the project development speed.

At least a portion of this scaffolding code will need to be able to generate Go
that is rended by the HTML package into HTML. For example, when scaffolding a
User model, the HTML package will need to be used to generate Go code that will
provide the various CRUD pages including the form for creation, editing, etc.
These forms should support the newest features (multierror messages, error
messages below the input field, highlighting the field red, etc and not require
javascript. 

The HTML package will be a rough-draft for one type of code generation, but
patterns experimented with in the HTML package will be valuable during the
development of CLI functionality in the command-line tool. This will allow
creation of HTML widgets/elements/reusable-code section, layout/template,
routes, controllers, models, individually, and full CRUD at once for a single
object.

### Development
Multiverse OS and associated projects are open source community iniatitives that
rely on volunteer work, donated resources and community driven design and
development.

**Breaking off sub-packages**
The templating (flash messages, favicon, template, login form built with forms)
all of this could be broken off into a templating package using the HTML
library. This could be combined with the portion in the server code, and made
into a server "plugin" type package. That way the HTML package is as generic 
and therefore usable in many projects as possible. 

The resulting HTML package should be useful for parsing HTML documents,
rendering, filtering/transforming (minifying, find/replace, etc). Rebuilding all
documents using whitelist and hardcoded symbols to attempt to guarantee no
tampering with resulting HTML by user input. 

Ideally this will be used in both our mechanize-like HTML client like software
and our web server for preparing HTML, parsing user input, and
providing a basic pipeline of fitlering/transforming before serving. 

#### References

[HTML5 Spec: Microreferences](https://html.spec.whatwg.org/multipage/common-microsyntaxes.html#common-microsyntaxes)


#### Contribute
Pull requests are welcome, and accepted after peer-review, (collaborative)
eediting,   and eventually pulling into the next iterative relesae version
sorted using coventional semantic versions Major.Minor.Patch. 

