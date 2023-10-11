// TODO: This is the beginning sketch of a active record like ORM that can be
// paired with the database agnostic manager that was completed earlier to
// provide powerful and rails-like data modeling with a lot of flexibility in
// choosing type, backend, mix-and-match, and keep them all open and working at
// the same time.

// With even the basic versions of each of these concepts we are have a good
// foundation on the database and model end of the maglev system and will
// hopefully let us start projects that have us writing the unique part of the
// code, not the same shit over and over.

// TODO: Maybe we can devise a single looping function for pulling by a field of
// the struct perhaps based on the type of comparison and feed in the field
// names but maybe in a way that generates go code automatically (jane), the
// reflect library, or enumerating each struct in its own package treating it
// like a model

//type CommandField int

// TODO: THIS IS A PRETTY GOOD IDEA BUT TOTALLY OUT OF THE FUCKING SCOPE OF THIS
// CLI FRAMEWORK. This is more of an active record replacement but that is needed
// for maglev yard so I'm gonna copy the code into that project and continue
// forth.
// NOTE: We only select the ones we are going to use as selectors. Minimizing
// things. This is all experimental on the fly kinda high level design; judge
// accordingly
//const (
//  CategoryField CommandField = iota
//  NameField
//  AliasField
//  HiddenField
//)

//type DataType int

//const (
//  IntegerType DataType = iota
//  StringType
//  BooleanType
//  DecimalType
//)

// Objects.Where(Height, Decimal).LessThan(5.41)
// Objects.Where(Hidden, Boolean).IsTrue()
// Objects.Where(Name, String).Contains("ello")
// Objects.Where(Alias, String).Is("a")
// Objects.Where(Books, Integer).MoreThan(5)
// Objects.Where(People, Object).== ability to disect single object

//func (self Commands) Select(field CommandField, dataType DataType, fieldValue interface{}) (commands Commands) {
//  for _, command := range self {
//    if CategoryField
//
//    if !command.Hidden {
//      commands = append(commands, command)
//    }
//  }
//}
