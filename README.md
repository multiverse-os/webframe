[<img src="https://avatars2.githubusercontent.com/u/24763891?s=400&u=c1150e7da5667f47159d433d8e49dad99a364f5f&v=4"  width="256px" height="256px" align="right" alt="Multiverse OS Logo">](https://github.com/multiverse-os)

## Multiverse: Webframe Web-Application Framework 
**URL** [multiverse-os.org](https://multiverse-os.org)

#### Introduction
Starshipyard is a web application framework heavily inspired by Ruby's Rails
framework, by developers with more than a decade of experience using Rails in
production environments.

This has translated into more than just taking on its design philosophy, and 
focus on human-readability; but additionally the importance put on testing, 
code generation for accelerated development, baking in features that are 
common to medium to large size web application, and providing a command-line 
interface tool that allows access to these tools.

The `starship` command-line includes the familiar: [(`new` not decided if
this will be included or opt for a model that includes the framework in the
application for overrides), `generate`,
`server`, `build`]. Those scratching their head about the last one, the `build`
command is meant to replace the Rail's command `rake`. The decision to use
`build` over `sake OR ssake` was picked because it is more descriptive, and 
while this project is heavily inspired by rails, it is not religious adherence, 
and the API and user interface will not be a direct translation. 

The `new` command will still create a directory filled with project files making
up a basic skeleton to accelerate getting started. However, it will provide a 
`rails-composer`-esque CLI prompt asking questions about the project to allow 
the developer to  pick defaults that best suite their project. 

Another important deviation from Rails is that `starshipyard` primary focus is on 
providing cutting-edge security over everything else. 

### Scramble Identity base session and authorization


### Templating 



### Development
**Tasks** include forward push to cover MVC, basic object example, code
generation for each of the MVC componenets, and completing each CLI command and
subcommand currently outlined in the `cmd/starship/main.go` file. 

**Additionally, the before/after action hooks on controllers need to be
implemented to provide the basic structure most Rails developers use (along with
the basic MVC, and routes in the root path).** 

Once the alpha is built with the above requirements, it will be at that time, it
is decided if the software will be released as is for applications to be built
over it, maintaining the Rails structure is the base project but allowing deep,
low-level customization by including the framework with the web application. Or
if, like Rails, this tool eventually is used to build a new project directory
and it is distributed as a binary, that builds, manages and launches starship
projects. Which the metaphor works well with that at the very least.



### Starship Yard Design and Security Focus
Starshipyard will leverage functionality provided by Multiverse OS's
`portal-gun` to provide a isolated and deterministic operating environment.

It will be deployed in a fully virtualized binary, using secure containers to
divide up the subcomponents and provide further isolation and security, by
limiting any damage breakouts are able to do, and allowing the subcomponents
to be ephemeral while the outside hardware virtualized VM maintains the data
to rebuild the secure containers. 

Additionally, this system always provides an isolated procses running in
isolation providing full router functionality. This allows for the machine
to be transparently routed, and essnetially impossible to leak its actual
location when using an inverted proxy (functionaly many people sacrafice
all their actual security and privacy for through cloudflare; when they
should be doing this themselves, since it is easy and doesn't require giving
up their SSL certificates, running iframes from cloduflare, and giving up all
their user privacy.

A quick sketch of the structure can possibly be elucidated with the following
illustration:

```

   Starship Yard
  ________________                       
  |Full HW Virt VM|   Ephemeral 
  |____ __________|    Secure                        
  |    |     |    |    Containers        _________________
  | DB | Web |Router      |             |_VM_or_Container_|     
  | || | App | || |       |             |                 |
  | \/ | ||  | \/ | /_____/             |  Gatway Router  |            
  |    | \/  |    | \                   |                 |         
  |    |     |    |                     | (Public WAN IP) |             
  |____|_____|__ _|                     |______________ __|                  
               |_______Inverting_Proxy_________________|

```

### MVC


##### Controllers

Should determine if controllers should have the routes merged in or if the
routes shoudl remain separate like with Rails. The routes will not likely end up
in the config path as with Rails because that location is tedious so itll either
be near controllers or in the root path. 

##### Templates
The controller is where we will select the template used and the views is where
the specifics of the template with be defined. Supporting multiple templates
should be a lot easier than it was in Rails and it should be very easy to use
specific templates for specific controller actions.


### Dependencies
Currently the only dependency is the `xid` library for handling universal ids,
but if we incorproate this library into a bson and UUID library as planned we
can remove all dependencies from this web framework. This would be ideal, using
only standard libraries and requiring no outside libraries beyond the Multiverse
maintained libraries (like scramble keys). 

*The remaining documentation needs to be updated, merged, or removed.*

-------------------------------------------------------------------------------
### Project Layout
Like a Rails project `starshipyard` will build a project folder that will be 
used by the core library and command-line tool to start and run a web
application. The core idea being that only the unique aspects of a web
application are stored in the project folder and so that developers of web
applications can focus just on the aspect they need too. Anything not related to
security should be customizable through overrides; following a design principle
of Multiverse OS, we do not want to make decisions for developers, we only
enforce security related design decisions but anything that is a matter of
preference like the data type used to hold configurations for example (YAML vs
TOML) should be easily configurable and set to the developers preferred option.

Below is the start of the structure, one major change is moving away from the
"app" folder, in favor a root folder containing all the major folders with less
nesting. 

```
├── app
│   ├── config
│   ├── controllers
│   ├── db
│   ├── jobs
│   ├── locales
│   ├── models
│   ├── routes.go
│   ├── tmp
│   └── views
├── cmd
│   └── starship
├── framework
│   ├── cmd
│   ├── codec
│   ├── config
│   ├── config.go
│   ├── controller
│   ├── database
│   ├── framework.go
│   ├── framework_test.go
│   ├── helpers
│   ├── html
│   ├── job
│   ├── model
│   ├── os
│   ├── router
│   ├── server
│   ├── service.go
│   ├── session
│   ├── version
│   └── view
├── install-deps.apt.sh
├── Makefile
└── README.md

```

One significant difference between `starshipyard` and `rails` is that the
library dependencies for your application are included. In addition, the
included files are organized into a single framework folder structure, 
instead of many different plugins. This allows for specific customizations 
to be made as needed while the framework remains in pre-alpha development. 

The folder one would generate with the `rails` app, is included as the `app`
folder. This template application will include a more sophisticated default
program than included with vanilla rails. It will include basic structure,
examples of all the code needed for a user system. 

The motivation behind to cover the almost ubiqitious problems, like users,
and login, but more importantly ensure that there is a secure baseline
for a feature that is almost included in every web application. 

The plan is to eventually turn these into easily installable and removeable
plugin-like modules. 

______
## Legacy Documentation  

Webframe is a security focused web application and SSE powered real-time REST
API framework designed for both use as described, for developing web
applications, but starshipyard is specifically being designed for use as a GUI when
combined with either a rendering engine component in an existing UI framework 
like wx or qt, or rendering the HTML another way served by a locally run
web-server. 

Webframe is heavily rails-inspired, but has entirely different priorities,
design, and naming; which results in primarily carrying over functionality
expectations: 

  * Conviences as `current_user` should be available to abstract away
  complexities beyond the web application scope.
    * All methods use `self` the same name used by default in Ruby. The reason
      for this is because it allows changing variable names freely without
      needing to change name in all assocated methods and I like it better than
      this. 

  * Expectations to avoid needless abbreviation in favor of readable code (we
    are compiling this code anyway, its not like its interpreted).

  * Will avoid being opinionated outside of matters of security. Want to use
    YAML or JSON or XML to load configurations? It should not matter, it should
    be built in a way that the config logic is agnostic to these matters of
    opinion. By now its probably better to be using something like CBOR for API
    data transfer and YAML (essentially a slightly slimmed down JSON) for
    config but it should not matter, so it is easier to interlope with
    existing systems. 

  * Provide a `rails` like command-line utility that will be able to:
      * control the webserver (start/stop/restart/...)
      * generate Go MVC code from command-line arguments or configuration file
      * generate Go test files associated with MVC code and provide tools and
        templates needed for testing models and controllers. 
      * run registered commands (rake/make) to easily script cleanup, update,
        etc

  * Will work without javascript enabled, too much of the internet is becoming
    impossible to use without javascript, which ignores the massive attack
    vector javascript is and failing to acknolwedge the client resource use is
    compounded if things that could be done on the server are moved to the
    client-- which may save large media companies money, the downside is its passed
    to the client, draining the battery faster battery, and generally wasting
   computer resources caused by unnecessary redundant execution of code. 
 


