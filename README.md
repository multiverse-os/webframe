[<img src="https://avatars2.githubusercontent.com/u/24763891?s=400&u=c1150e7da5667f47159d433d8e49dad99a364f5f&v=4"  width="256px" height="256px" align="right" alt="Multiverse OS Logo">](https://github.com/multiverse-os)

## Multiverse: Maglev Application Framework
**URL** [multiverse-os.org](https://multiverse-os.org)

#### Introduction
Maglev is primarily a web application framework heavily inspired by Ruby's Rails
framework, by developers with more than a decade of experience using Rails in
production environments.

Rails for carts in a mine to move `Ruby`, Maglev is both the name for the 
bullet train and the tracks, which `Go` very fast. Silly, but it is meant
to signify how similar the two projects are, and not like so many other 
projects which claim to be Rails inspired, but provide less than 5% of 
feature parity beyond some organizational similarities. As Ruby was to Pearl,
Maglev is to Rails. 

This has translated into more than just taking on its design philosophy, and 
focus on human-readability; but additionally the importance put on testing, 
code generation for accelerated development, baking in features that are 
common to medium to large size web application, and providing a command-line 
interface tool that allows access to these tools.

The `maglev` command-line includes the familiar: [(`new` not decided if
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

Another important deviation from Rails is that `maglev` primary focus is on 
providing cutting-edge security over everything else. 

*The documentation needs to be updated, merged, or removed.*

### Development
**Tasks** include forward push to cover MVC, basic object example, code
generation for each of the MVC components, and completing each CLI command and
subcommand currently outlined in the `cmd/maglev/main.go` file. 

**Additionally, the before/after action hooks on controllers need to be
implemented to provide the basic structure most Rails developers use (along with
the basic MVC, and routes in the root path).** 

Once the alpha is built with the above requirements, it will be at that time, it
is decided if the software will be released as is for applications to be built
over it, maintaining the Rails structure is the base project but allowing deep,
low-level customization by including the framework with the web application. Or
if, like Rails, this tool eventually is used to build a new project directory
and it is distributed as a binary, that builds, manages and launches `maglev` 
projects. Which the metaphor works well with that at the very least.


### Design and Security Focus
`maglev` will leverage functionality provided by Multiverse OS's
`portal-gun` to provide a isolated and deterministic operating environment.

It will be deployed in a fully virtualized binary, using secure containers to
divide up the sub-components and provide further isolation and security, by
limiting any damage breakouts are able to do, and allowing the subcomponents
to be ephemeral while the outside hardware virtualized VM maintains the data
to rebuild the secure containers. 

Additionally, this system always provides an isolated processes running in
isolation providing full router functionality. This allows for the machine
to be transparently routed, and essentially impossible to leak its actual
location when using an inverted proxy (functionality many people sacrifice
all their actual security and privacy for through cloudflare; when they
should be doing this themselves, since it is easy and doesn't require giving
up their SSL certificates, running iframes from cloudflare, and giving up all
their user privacy.

A quick sketch of the structure can possibly be elucidated with the following
illustration:

```

   `maglev` Yard
   __________________________                      
  |     Full HW Virt VM      |   Ephemeral 
  |__________________________|     Secure                        
  |        |        |        |       VMs        _________________
  |   DB   |  Web   | Router |       |         |_VM_(VPS)________|     
  |   ||   |   App  |   ||   |       |         |                 |
  |   \/   |   ||   |   \/   | .:|___/         |  Gatway Router  |            
  |        |   \/   |        | `\|             |                 |         
  |        |        |        |                 | (Public WAN IP) |             
  |________|________|___ ____|                 |______________ __|                  
                        |_______Inverting_Proxy_______________|

```


##### Controllers
Should determine if controllers should have the routes merged in or if the
routes should remain separate like with Rails. The routes will not likely end up
in the config path as with Rails because that location is tedious so it'll either
be near controllers or in the root path. 

##### Templates
The controller is where we will select the template used and the views is where
the specifics of the template with be defined. Supporting multiple templates
should be a lot easier than it was in Rails and it should be very easy to use
specific templates for specific controller actions.


## Legacy Documentation  
`maglev` is a security focused web application and SSE powered real-time REST
API framework designed for both use as described, for developing web
applications, but `maglev` is specifically being designed for use as a GUI when
combined with either a rendering engine component in an existing UI framework.

`maglev` is heavily rails-inspired, but has entirely different priorities,
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
 


