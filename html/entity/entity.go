package entity

const EmptyString = ``

type HTMLEntity int
type Entities map[HTMLEntity]bool

type Entity struct {
	Type HTMLEntity
	// UTF8() -> UTF8 version like \U+000...
	// ASCII() ->
	// HTMLEscape() -> &nbsp;
	//
}

// TODO: Split up ASCII Entities from UTF-8 Entities, then be able to load
// each by charset name.

const (
	// Special
	EmptyEntity HTMLEntity = iota // 127 or 7F? Blank could be ' ' so Empty (and nil?))
	Space                         // (or none-breaking space or NoneBreakingSpace (nbsp))
	// Escaped ('\' prefix)
	NewLine // &NewLine;
	CarriageReturn
	Tab // &Tab;
	// Numbers
	Zero
	One
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	// Alpha Should we just use 1 set? downcase and upcase?
	RuneA
	RuneB
	RuneC
	RuneD
	RuneE
	RuneF
	RuneG
	RuneH
	RuneI
	RuneJ
	RuneK
	RuneL
	RuneM
	RuneN
	RuneO
	RuneP
	RuneQ
	RuneR
	RuneS
	RuneT
	RuneU
	RuneV
	RuneW
	RuneX
	RuneY
	RuneZ
	Runea
	Runeb
	Runec
	Runed
	Runee
	Runef
	Runeg
	Runeh
	Runei
	Runej
	Runek
	Runel
	Runem
	Runen
	Runeo
	Runep
	Runeq
	Runer
	Runes
	Runet
	Runeu
	Runev
	Runew
	Runex
	Runey
	Runez
	// Symbols
	// ASCII
	RegisteredTradeMark
	Copyright
	SoundRecordingCopyright
	ServiceMark
	TwoDotMark
	FourDotMark
	FiveDotMark
	QuadruplePrime
	SwungDash
	DottedCross
	DoublePi
	DoubleCapitalPi
	DoubleSummation
	PropertyLine
	TurnedAmpersand
	PerSign
	Aktieselskab
	SetMinus
	AsterixOperator
	RingOperator   // Is this the same as degree?
	BulletOperator // Is this the same as bullet?
	ProportionalTo
	Infinity
	RightAngle
	Angle
	MeasuredAngle
	SphericalAngle
	Divides
	DoesNotDivide
	Parallel
	NotParallel
	LogicalAnd
	LogicalOr
	Intersection
	Union
	Integral
	DoubleIntegral
	TripleIntegral
	ContourIntegral
	Therefore
	Because
	Ratio
	Proportion
	DotMinus
	Excess
	GeometricProportion
	Homothetic
	ReversedTilde
	InvertedLazyS
	CommercialMinusSign
	TwoVerticalAsterix
	CloseUp
	ReverseSemicolon
	MidlineHorizontalEllipsis
	WhiteDiamondWithBlackDiamond
	GermanPennySymbol
	LibreTournoisSign
	ThereDoesNotExist
	LowAsterix
	TironianSign
	CaretInsertionPoint
	BlackLeftBullet
	BlackRightBullet
	WreathProduct
	NotTilde
	MinusTilde
	AlmostEqual
	NotAlmostEqual
	AlmostEqualOrEqual
	TripleTilde
	AllEqual
	Equivilent
	NotEqual
	NotIdentical
	QuestionedEqual
	MeasuredBy
	EqualToByDefinition
	DeltaEqual
	StarEqual
	Estimates
	CorrespondsTo
	RingInEqual
	StrictlyEquivilent
	Between
	CircledPlus
	CircledMinus
	CircledTimes
	CircledDivision
	CircleDot
	CircleRing
	CircleAsterix
	CircleEqual
	CircleDash
	SquaredPlus
	SquaredMinus
	SquaredTimes
	SquaredDot
	RightTack
	LeftTack
	DownTack
	UpTack
	Assertion
	Models
	TrueSign
	Forces
	OriginalOf
	ImageOf
	Multimap
	HermitianMatrix
	Pitchfork
	DoubleIntersection
	DoubleSuperset
	DoubleUnion
	EqualAndParallel
	VeryMuchLessThan
	VeryMuchGreaterThan
	VerticalEllipsis
	MIdlineHorizontalEllipsis
	LongHorizontalStroke
	Square
	Comet
	FilledStar
	EmptyStar
	Sun
	AscendingNode
	DescendingNode
	Conjunction
	Opposition
	CheckboxSymbol
	CheckboxCheckedSymbol
	CheckboxXSymbol
	Saltire
	Communism
	Caduceus
	Radioactive
	Biohazard
	CautionSign
	SkullAndCrossbones
	SunWithRays
	Cloud
	Umbrella
	RainAndUmbrella
	HotBeverage
	Shamrock
	FloralHeartBullet
	Peace
	WhiteDiamond
	BlackDiamond
	WhiteDiamondWhiteBlackDiamond
	WhiteCircle
	DottedCircle
	CircleWithVerticalFill
	Bullseye
	BlackCircle
	CircleWithBlackLeftHalf
	CircleWithBlackRightHalf
	CircleWithUpperRightQuadrant
	CircleWithUpperLeftQuadrant
	LeftHalfBlackCircle
	RightHalfBlackCircle
	InverseBullet
	UpperLeftQuadrantArc
	UpperRightQuadrantArc
	LowerLeftQuadrantArc
	LowerRightQuadrantArc
	UpperHalfCircle
	LowerHalfCircle
	Recording
	WhiteShogi
	BlackShogi
	Snowman
	Lightning
	Thunderstorm
	Scissors
	UpperBladeScissors
	LowerBladeScissors
	EmptyScissors
	TelephoneLocationSign
	TapeDrive
	Airplane
	Envelope
	RaisedFist
	RaisedHand
	VictorySign
	WritingHand
	LowerRightPencil
	Pencil
	UpperRightPencil
	EmptyNib
	Nib
	NotElementOf
	ContainsMemberOf
	DivisionSign
	Checkmark
	HeavyCheckmark
	HeaviestCheckmark
	MultiplicationSign
	HeavyMultiplicationSign
	BallotMark
	HeavyBallotMark
	OutlinedGreekCross
	GreekCross
	OpenCentreCross
	HeavyOpenCentreCross
	FilledFourPointedStar
	EmptyFourPointedStar
	Sparkles
	StressedOutlinedWhiteStar
	CircledWhiteStar
	HeavyDoubleQuote
	HeavyDoubleTurnedCommaQuote
	HeavySingleTurnedCommaQuote
	HeavySingleQuote
	HeavyComma
	HeavyDoubleComma
	CurvedStemParagraph
	HeavyHeartBang
	Heart
	RotatedHeart
	FloralHeart
	RotatedFloralHeart
	MediumLeftParen
	MediumRightParen
	MediumFlattenedLeftParen
	MediumFlattenedRightParen
	BlackCircleOne
	BlackCircleTwo
	BlackCircleThree
	BlackCircleFour
	BlackCircleFive
	BlackCircleSix
	BlackCircleSeven
	BlackCircleEight
	BlackCircleNine
	BlackCircleTen
	CircleOne
	CircleTwo
	CircleThree
	CircleFour
	CircleFive
	CircleSix
	CircleSeven
	CircleEight
	CircleNine
	CircleTen
	HeavyPlusSign
	HeavyMinusSign
	HeavyDivisionSign
	CurlyLoop
	DoubleCurlyLoop
	Interrobang
	DoubleBang
	Tricolon
	VerticalFourDots
	OpenOutlinedRightArrow
	FilledTelephone
	EmptyTelephone
	SquareRoot
	CubeRoot
	FourthRoot
	FaxSign
	TelephoneSign
	DoubleZ
	OunceSign
	OhmSign
	InvertedOhmSign
	Kelvin
	AngstromSign
	NotSign
	Overline
	Diaeresis
	ColonSign
	CruzeiroSign
	EuroCurrencySign
	FrenchFranc
	LiraSign
	MilSign
	NairaSign
	PesetaSign
	RupeeSign
	WonSign
	NewSheqelSign
	DongSign
	KipSign
	TugrikSign
	DrachmaSign
	GermanyPennySymbol
	PesoSign
	GuaraniSign
	AustralSign
	HryvniaSign
	CediSign
	LibreTournosisSign
	SpesmiloSign
	TengeSign
	IndianRupeeSign
	TurkishLiraSign
	NordicMarkSign
	ManatSign
	RubleSign
	LariSign
	BitcoinSign
	FeminineOrdinal

	// UTF-8
	ReverseTilde
	SineWave
	Prescription
	SmallScript
	SmallScriptE
	CapitalScript
	CapitalScriptB
	EstimatedSign
	Scruple
	Euler
	Eth
	SmallEth
	OE
	SmallOE
	AccountOf
	AddressedToTheSubject
	CareOf
	CadaUna
	SWithCaron
	SmallSWithCaron
	YWithDiaeresis
	SmallYWithAcute
	SmallYWithDiaeresis
	NWithTilde
	OWithGrave
	SmallOWithGrave
	AcuteAccent
	OWithAcute
	SmallOWithAcute
	OWithCircumflex
	SmallOWithCircumflex
	OWithTilde
	SmallOWithTilde
	OWithDiaeresis
	UWithDiaeresis
	UWithCircumflex
	UWithAcute
	UWithGrave
	Cedilla
	Ordinal
	ForAllSign
	Complement
	PartialDifferential
	ThereExists
	ThereDoesNotExists
	SmallElementOf
	SmallContainsAsMember
	EndOfProof
	NArrayProduct
	NArrayCoproduct
	NArraySummation
	DotPlusSign
	EnDash
	EmDash
	GraveAccent
	LowSingleQuote
	RightSingleQuote
	LeftSingleQuote
	LowDoubleQuote
	LeftDoubleQuote
	RightDoubleQuote
	LeftDoubleAngleQuote
	RightDoubleAngleQuote

	// Classic Symbols
	Function
	Sharf
	Thorn
	Paragraph
	PlusMinusSign
	MiddleDot
	SuperscriptOne
	Squared
	Cubed
	Micro
	InvertedQuestionMark
	Multiplication
	ZeroWithSlash
	SmallZeroWithSlash
	OneQuarterSign
	OneHalfSign
	ThreeFourthSign
	EmptySet
	Increment
	Decrement
	ElementOf
	ContainsMemeberOf
	DoesNotContainAsMember

	Yen
	BritishPound
	Currency
	SectionSign
	PerThousandSign
	BrokenVerticalBar
	Dagger
	DoubleDagger
	Euro
	Ellipsis
	TradeMark

	Hash         // # (or Pound or NumberSign)
	CentSign     //
	DoubleQuote  // "
	Quote        // ' (or SingleQuote)
	Period       // .
	Comma        // ,
	At           // @
	Caret        // ^ (or circumflex)
	Semicolon    // ;
	Colon        // :
	Slash        // / (or ForwardSlash)
	Backslash    // \
	Bang         // ! (ExclamationPoint)
	DollarSign   // $
	QuestionMark // ?
	Ampersand    // &
	Underscore   // _
	Tilde        // ~ (or math related as EquivalencySign)
	Bullet       // •
	Degree       // °
	Pipe         // |
	Apostrophe   // ` (or GraveAccent)
	InvertedBang // ¡ (inverted exclamation mark)
	// Math (May be worth doing PlusSign, MinusSign, PercentSign)
	MinusSign   // - (or minus symbol, hypen, (dash might be different length))
	PlusSign    // +
	Asterix     // *
	EqualSign   // =
	PercentSign // %
	// Brackets (May be worth doing LessThanSign, ...)
	OpenSquareBracket  // [ (or just bracket)
	CloseSquareBracket // ]
	OpenGuillemet      // «
	CloseGuillemet     // »
	OpenParen          // ( (remember plural is parentheses)
	CloseParen         // ) (or ClosingParenthesis)
	LessThanSign       // < (or angle bracket)
	GreaterThanSign    // >
	OpenCurlyBrace     // { (HTML spec has it as OpeningBrace)
	CloseCurlyBrace    // }
)

// Aliasing to improve the API by making it more intuitive
// Maybe (If the namespace isnt needed, might as well just rename the
// originals and get rid of this so there is less bloat anyways)
const (
	NonbreakingSpace  = Space
	NBSP              = Space
	Return            = CarriageReturn
	CR                = CarriageReturn
	NL                = NewLine
	OpenParenthesis   = OpenParen
	CloseParenthesis  = CloseParen
	Dash              = MinusSign
	Hyphen            = MinusSign
	Pound             = Hash
	ExclamationMark   = Bang
	SingleQuote       = Quote
	ForwardSlash      = Slash
	OpenBrace         = OpenCurlyBrace
	CloseBrace        = CloseCurlyBrace
	OpenCurlyBracket  = OpenCurlyBrace
	CloseCurlyBracket = CloseCurlyBrace
	OpenBracket       = OpenSquareBracket
	CloseBracket      = CloseSquareBracket
	OpenAngleBracket  = LessThanSign
	CloseAngleBracket = GreaterThanSign
	OpenChevron       = OpenGuillemet
	CloseChevron      = CloseGuillemet
	FullStop          = Period
)

func MarshalSymbolType(inputSymbol string) HTMLEntity {
	// TODO: This will be used for later homoglyph validation, template parsing
	if inputSymbol == EmptyString {
		return EmptyEntity
	}
	// Iterate over all whitelisted HTMLEntities
	// OLD code
	//for symbol, allowed := range AllowedSymbolTypes() {
	//	if allowed && symbol.String() == strings.ToLower(inputSymbol) {
	//		return symbol
	//	}
	//}
	return EmptyEntity
}

// TODO: Lets store the binary :D
// U+232A (UTF-8) for  `` RIGHT-POINTING ANGLE BRACKET (HTML &#9002; · &rang;)
// TODO: Still missing a lot but now it has good coverage for both ASCII and UTF-8 this
// will allow us to pair up homoglyphs to enable our chinese support for good localization
// while still being secure.
func (self HTMLEntity) String() string {
	switch self {
	case Space:
		return ` `
	case NewLine:
		return `\n`
	case CarriageReturn:
		return `\r`
	case Tab:
		return `\t`
	case Degree:
		return `°`
	case RegisteredTradeMark:
		return `®`
	case Copyright:
		return `©`
	case SoundRecordingCopyright:
		return `℗`
	case Prescription:
		return `℞`
	case ServiceMark:
		return `℠`
	case TwoDotMark:
		return `⁚`
	case FourDotMark:
		return `⁛`
	case FiveDotMark:
		return `⁙`
	case QuadruplePrime:
		return `⁗`
	case SwungDash:
		return `⁓`
	case DottedCross:
		return `⁜`
	case DoublePi:
		return `ℼ`
	case DoubleCapitalPi:
		return `ℿ`
	case DoubleSummation:
		return `⅀`
	case PropertyLine:
		return `⅊`
	case TurnedAmpersand:
		return `⅋`
	case PerSign:
		return `⅌`
	case Aktieselskab:
		return `⅍`
	case SetMinus:
		return `∖`
	case AsterixOperator:
		return `∗`
	case RingOperator:
		return `∘`
	case BulletOperator:
		return `∙`
	case ProportionalTo:
		return `∝`
	case Infinity:
		return `∞`
	case RightAngle:
		return `∟`
	case Angle:
		return `∠`
	case MeasuredAngle:
		return `∡`
	case SphericalAngle:
		return `∢`
	case Divides:
		return `∣`
	case DoesNotDivide:
		return `∤`
	case Parallel:
		return `∥`
	case NotParallel:
		return `∦`
	case LogicalAnd:
		return `∧`
	case LogicalOr:
		return `∨`
	case Intersection:
		return `∩`
	case Union:
		return `∪`
	case Integral:
		return `∫`
	case DoubleIntegral:
		return `∬`
	case TripleIntegral:
		return `∭`
	case ContourIntegral:
		return `∮`
	case Therefore:
		return `∴`
	case Because:
		return `∵`
	case Ratio:
		return `∶`
	case Proportion:
		return `∷`
	case DotMinus:
		return `∸`
	case Excess:
		return `∹`
	case GeometricProportion:
		return `∺`
	case Homothetic:
		return `∻`
	case Tilde:
		return `∼`
	case ReversedTilde:
		return `∽`
	case InvertedLazyS:
		return `∾`
	case CommercialMinusSign:
		return `⁒`
	case TwoVerticalAsterix:
		return `⁑`
	case CloseUp:
		return `⁐`
	case ReverseSemicolon:
		return `⁏`
	case LowAsterix:
		return `⁎`
	case TironianSign:
		return `⁊`
	case CaretInsertionPoint:
		return `⁁`
	case BlackLeftBullet:
		return `⁌`
	case BlackRightBullet:
		return `⁍`
	case SineWave:
		return `∿`
	case WreathProduct:
		return `≀`
	case NotTilde:
		return `≁`
	case MinusTilde:
		return `≂`
	case AlmostEqual:
		return `≈`
	case NotAlmostEqual:
		return `≉`
	case AlmostEqualOrEqual:
		return `≊`
	case TripleTilde:
		return `≋`
	case AllEqual:
		return `≌`
	case Equivilent:
		return `≍`
	case NotEqual:
		return `≠`
	case NotIdentical:
		return `≢`
	case QuestionedEqual:
		return `≟`
	case MeasuredBy:
		return `≞`
	case EqualToByDefinition:
		return `≝`
	case DeltaEqual:
		return `≜`
	case StarEqual:
		return `≛`
	case Estimates:
		return `≙`
	case CorrespondsTo:
		return `≘`
	case RingInEqual:
		return `≖`
	case StrictlyEquivilent:
		return `≣`
	case Between:
		return `≬`
	case CircledPlus:
		return `⊕`
	case CircledMinus:
		return `⊖`
	case CircledTimes:
		return `⊗`
	case CircledDivision:
		return `⊘`
	case CircleDot:
		return `⊙`
	case CircleRing:
		return `⊚`
	case CircleAsterix:
		return `⊛`
	case CircleEqual:
		return `⊜`
	case CircleDash:
		return `⊝`
	case SquaredPlus:
		return `⊞`
	case SquaredMinus:
		return `⊟`
	case SquaredTimes:
		return `⊠`
	case SquaredDot:
		return `⊡`
	case RightTack:
		return `⊢`
	case LeftTack:
		return `⊣`
	case DownTack:
		return `⊤`
	case UpTack:
		return `⊥`
	case Assertion:
		return `⊦`
	case Models:
		return `⊧`
	case TrueSign:
		return `⊨`
	case Forces:
		return `⊩`
	case OriginalOf:
		return `⊶`
	case ImageOf:
		return `⊷`
	case Multimap:
		return `⊸`
	case HermitianMatrix:
		return `⊹`
	case Pitchfork:
		return `⋔`
	case DoubleIntersection:
		return `⋒`
	case DoubleSuperset:
		return `⋑`
	case DoubleUnion:
		return `⋓`
	case EqualAndParallel:
		return `⋕`
	case VeryMuchLessThan:
		return `⋘`
	case VeryMuchGreaterThan:
		return `⋙`
	case VerticalEllipsis:
		return `⋮`
	case MidlineHorizontalEllipsis:
		return `⋯`
	case LongHorizontalStroke:
		return `⋲`
	case Square:
		return `□`
	case Comet:
		return `☄`
	case FilledStar:
		return `★`
	case EmptyStar:
		return `☆`
	case Sun:
		return `☉`
	case AscendingNode:
		return `☊`
	case DescendingNode:
		return `☋`
	case Conjunction:
		return `☌`
	case Opposition:
		return `☍`
	case CheckboxSymbol:
		return `☐`
	case CheckboxCheckedSymbol:
		return `☑`
	case CheckboxXSymbol:
		return `☒`
	case Saltire:
		return `☓`
	case Communism:
		return `☭`
	case Caduceus:
		return `☤`
	case Radioactive:
		return `☢`
	case Biohazard:
		return `☣`
	case CautionSign:
		return `☡`
	case SkullAndCrossbones:
		return `☠`
	case SunWithRays:
		return `☀`
	case Cloud:
		return `☁`
	case Umbrella:
		return `☂`
	case RainAndUmbrella:
		return `☔`
	case HotBeverage:
		return `☕`
	case Shamrock:
		return `☘`
	case FloralHeartBullet:
		return `☙`
	case Peace:
		return `☮`
	case WhiteDiamond:
		return `◇`
	case BlackDiamond:
		return `◆`
	case WhiteDiamondWithBlackDiamond:
		return `◈`
	case WhiteCircle:
		return `○`
	case DottedCircle:
		return `◌`
	case CircleWithVerticalFill:
		return `◍`
	case Bullseye:
		return `◎`
	case BlackCircle:
		return `●`
	case CircleWithBlackLeftHalf:
		return `◐`
	case CircleWithBlackRightHalf:
		return `◑`
	case CircleWithUpperRightQuadrant:
		return `◔`
	case CircleWithUpperLeftQuadrant:
		return `◕`
	case LeftHalfBlackCircle:
		return `◖`
	case RightHalfBlackCircle:
		return `◗`
	case InverseBullet:
		return `◘`
	case UpperLeftQuadrantArc:
		return `◜`
	case UpperRightQuadrantArc:
		return `◝`
	case LowerLeftQuadrantArc:
		return `◟`
	case LowerRightQuadrantArc:
		return `◞`
	case UpperHalfCircle:
		return `◠`
	case LowerHalfCircle:
		return `◡`
	case Recording: // Fisheye
		return `◉`
	case WhiteShogi:
		return `☖`
	case BlackShogi:
		return `☗`
	case Snowman:
		return `☃`
	case Lightning:
		return `☇`
	case Thunderstorm:
		return `☈`
	case Scissors:
		return `✂`
	case UpperBladeScissors:
		return `✁`
	case LowerBladeScissors:
		return `✃`
	case EmptyScissors:
		return `✄`
	case TelephoneLocationSign:
		return `✆`
	case TapeDrive:
		return `✇`
	case Airplane:
		return `✈`
	case Envelope:
		return `✉`
	case RaisedFist:
		return `✊`
	case RaisedHand:
		return `✋`
	case VictorySign:
		return `✌`
	case WritingHand:
		return `✍`
	case LowerRightPencil:
		return `✎`
	case Pencil:
		return `✏`
	case UpperRightPencil:
		return `✐`
	case EmptyNib:
		return `✑`
	case Nib:
		return `✒`
	case Checkmark:
		return `✓`
	case HeavyCheckmark:
		return `✔`
	case HeaviestCheckmark:
		return `✅`
	case MultiplicationSign:
		return `✕`
	case HeavyMultiplicationSign:
		return `✖`
	case BallotMark:
		return `✗`
	case HeavyBallotMark:
		return `✘`
	case OutlinedGreekCross:
		return `✙`
	case GreekCross:
		return `✚`
	case OpenCentreCross:
		return `✛`
	case HeavyOpenCentreCross:
		return `✜`
	case FilledFourPointedStar:
		return `✦`
	case EmptyFourPointedStar:
		return `✧`
	case Sparkles:
		return `✨`
	case StressedOutlinedWhiteStar:
		return `✩`
	case CircledWhiteStar:
		return `✪`
	case HeavyDoubleQuote:
		return `❝`
	case HeavyDoubleTurnedCommaQuote:
		return `❞`
	case HeavySingleTurnedCommaQuote:
		return `❛`
	case HeavySingleQuote:
		return `❜`
	case HeavyComma:
		return `❟`
	case HeavyDoubleComma:
		return `❠`
	case CurvedStemParagraph:
		return `❡`
	case HeavyHeartBang:
		return `❣`
	case Heart:
		return `❤`
	case RotatedHeart:
		return `❥`
	case FloralHeart:
		return `❦`
	case RotatedFloralHeart:
		return `❧`
	case MediumLeftParen:
		return `❨`
	case MediumRightParen:
		return `❩`
	case MediumFlattenedLeftParen:
		return `❪`
	case MediumFlattenedRightParen:
		return `❫`
	case BlackCircleOne:
		return `❶`
	case BlackCircleTwo:
		return `❷`
	case BlackCircleThree:
		return `❸`
	case BlackCircleFour:
		return `❹`
	case BlackCircleFive:
		return `❺`
	case BlackCircleSix:
		return `❻`
	case BlackCircleSeven:
		return `❼`
	case BlackCircleEight:
		return `❽`
	case BlackCircleNine:
		return `❾`
	case BlackCircleTen:
		return `❿`
	case CircleOne:
		return `➀`
	case CircleTwo:
		return `➁`
	case CircleThree:
		return `➂`
	case CircleFour:
		return `➃`
	case CircleFive:
		return `➄`
	case CircleSix:
		return `➅`
	case CircleSeven:
		return `➆`
	case CircleEight:
		return `➇`
	case CircleNine:
		return `➈`
	case CircleTen:
		return `➉`
	case HeavyPlusSign:
		return `➕`
	case HeavyMinusSign:
		return `➖`
	case HeavyDivisionSign:
		return `➗`
	case CurlyLoop:
		return `➰`
	case DoubleCurlyLoop:
		return `➿`
	case Interrobang:
		return `‽`
	case DoubleBang:
		return `‼`
	case Tricolon:
		return `⁝`
	case VerticalFourDots:
		return `⁞`
	case OpenOutlinedRightArrow:
		return `➾`
	case FilledTelephone:
		return `☎`
	case EmptyTelephone:
		return `☏`
	case SquareRoot:
		return `√`
	case CubeRoot:
		return `∛`
	case FourthRoot:
		return `∜`
	case FaxSign:
		return `℻`
	case TelephoneSign:
		return `℡`
	case DoubleZ:
		return `ℤ`
	case OunceSign:
		return `℥`
	case OhmSign:
		return `Ω`
	case InvertedOhmSign:
		return `℧`
	case Kelvin:
		return `K`
	case AngstromSign:
		return `Å`
	case NotSign:
		return `¬`
	case Overline:
		return `¯`
	case Diaeresis:
		return `¨`
	case ColonSign:
		return `₡`
	case CruzeiroSign:
		return `₢`
	case EuroCurrencySign:
		return `₠`
	case FrenchFranc:
		return `₣`
	case LiraSign:
		return `₤`
	case MilSign:
		return `₥`
	case NairaSign:
		return `₦`
	case PesetaSign:
		return `₧`
	case RupeeSign:
		return `₨`
	case WonSign:
		return `₩`
	case NewSheqelSign:
		return `₪`
	case DongSign:
		return `₫`
	case KipSign:
		return `₭`
	case TugrikSign:
		return `₮`
	case DrachmaSign:
		return `₯`
	case GermanPennySymbol:
		return `₰`
	case PesoSign:
		return `₱`
	case GuaraniSign:
		return `₲`
	case AustralSign:
		return `₳`
	case HryvniaSign:
		return `₴`
	case CediSign:
		return `₵`
	case LibreTournoisSign:
		return `₶`
	case SpesmiloSign:
		return `₷`
	case TengeSign:
		return `₸`
	case IndianRupeeSign:
		return `₹`
	case TurkishLiraSign:
		return `₺`
	case NordicMarkSign:
		return `₻`
	case ManatSign:
		return `₼`
	case RubleSign:
		return `₽`
	case LariSign:
		return `₾`
	case BitcoinSign:
		return `₿`
	case Yen:
		return `¥`
	case BritishPound:
		return `£`
	case Currency:
		return `¤`
	case SectionSign:
		return `§`
	case CentSign:
		return `¢`
	case InvertedBang:
		return `¡`
	case BrokenVerticalBar:
		return `¦`
	case FeminineOrdinal:
		return `ª`
	case Bullet:
		return `•`
	case PerThousandSign:
		return `‰`
	case Dagger:
		return `†`
	case DoubleDagger:
		return `‡`
	case Euro:
		return `€`
	case Ellipsis:
		return `…`
	case TradeMark:
		return `™`
	case SmallScript:
		return `ℊ`
	case SmallScriptE:
		return `ℯ`
	case CapitalScript:
		return `ℋ`
	case CapitalScriptB:
		return `ℬ`
	case EstimatedSign:
		return `℮`
	case Scruple:
		return `℈`
	case Euler:
		return `ℇ`
	case Eth:
		return `Ð`
	case SmallEth:
		return `ð`
	case OE:
		return `Œ`
	case SmallOE:
		return `œ`
	case AccountOf:
		return `℀`
	case AddressedToTheSubject:
		return `℁`
	case CareOf:
		return `℅`
	case CadaUna:
		return `℆`
	case SWithCaron:
		return `Š`
	case SmallSWithCaron:
		return `š`
	case YWithDiaeresis:
		return `Ÿ`
	case SmallYWithAcute:
		return `ý`
	case SmallYWithDiaeresis:
		return `ÿ`
	case NWithTilde:
		return `ñ`
	case OWithGrave:
		return `Ò`
	case SmallOWithGrave:
		return `ò`
	case AcuteAccent:
		return `´`
	case OWithAcute:
		return `Ó`
	case SmallOWithAcute:
		return `ó`
	case OWithCircumflex:
		return `Ô`
	case SmallOWithCircumflex:
		return `ô`
	case OWithTilde:
		return `Õ`
	case SmallOWithTilde:
		return `õ`
	case OWithDiaeresis:
		return `Ö`
	case UWithDiaeresis:
		return `ü`
	case UWithCircumflex:
		return `û`
	case UWithAcute:
		return `ú`
	case UWithGrave:
		return `ù`
	case Function:
		return `ƒ`
	case Ampersand:
		return `&`
	case Colon:
		return `:`
	case Thorn:
		return `Þ`
	case Sharf:
		return `ß`
	case Paragraph:
		return `¶`
	case PlusMinusSign:
		return `±`
	case MiddleDot:
		return `·`
	case Cedilla:
		return `¸`
	case SuperscriptOne:
		return `¹`
	case Squared:
		return `²`
	case Cubed:
		return `³`
	case Micro:
		return `µ`
	case Ordinal:
		return `º`
	case InvertedQuestionMark:
		return `¿`
	case ForAllSign:
		return `∀`
	case Complement:
		return `∁`
	case PartialDifferential:
		return `∂`
	case Multiplication:
		return `×`
	case ZeroWithSlash:
		return `Ø`
	case SmallZeroWithSlash:
		return `ø`
	case OneQuarterSign:
		return `¼`
	case OneHalfSign:
		return `½`
	case ThreeFourthSign:
		return `¾`
	case ThereExists:
		return `∃`
	case ThereDoesNotExist:
		return `∄`
	case EmptySet:
		return `∅`
	case Increment:
		return `∆`
	case Decrement: // or Nabla
		return `∇`
	case ElementOf:
		return `∈`
	case NotElementOf:
		return `∉`
	case SmallElementOf:
		return `∊`
	case ContainsMemberOf:
		return `∋`
	case DoesNotContainAsMember:
		return `∌`
	case SmallContainsAsMember:
		return `∍`
	case EndOfProof:
		return `∎`
	case NArrayProduct:
		return `∏`
	case NArrayCoproduct:
		return `∐`
	case NArraySummation:
		return `∑`
	case DotPlusSign:
		return `∔`
	case DivisionSign:
		return `÷`
	case EqualSign:
		return `=`
	case EnDash:
		return `–`
	case EmDash:
		return `—`
	case Underscore:
		return `_`
	case GraveAccent:
		return "`"
	case LowSingleQuote:
		return `‚`
	case RightSingleQuote:
		return `’`
	case LeftSingleQuote:
		return `‘`
	case LowDoubleQuote:
		return `„`
	case LeftDoubleQuote:
		return `“`
	case RightDoubleQuote:
		return `”`
	case DoubleQuote:
		return `"`
	case SingleQuote:
		return `'`
	case LeftDoubleAngleQuote:
		return `«`
	case RightDoubleAngleQuote:
		return `»`
	case PercentSign:
		return `%`
	case Semicolon:
		return `;`
	case Caret:
		return `^`
	case LessThanSign:
		return `<`
	case GreaterThanSign:
		return `>`
	case Slash:
		return `/`
	case Backslash:
		return `\`
	case Bang:
		return `!`
	case OpenParen:
		return `(`
	case CloseParen:
		return `)`
	case OpenCurlyBrace:
		return `{`
	case CloseCurlyBrace:
		return `}`
	case OpenSquareBracket:
		return `[`
	case CloseSquareBracket:
		return `]`
	case OpenGuillemet:
		return `«`
	case CloseGuillemet:
		return `»`
	case Zero:
		return "0"
	case One:
		return "1"
	case Two:
		return "2"
	case Three:
		return "3"
	case Four:
		return "4"
	case Five:
		return "5"
	case Six:
		return "6"
	case Seven:
		return "7"
	case Eight:
		return "8"
	case Nine:
		return "9"
	case RuneA:
		return "A"
	case Runea:
		return "a"
	case RuneB:
		return "B"
	case Runeb:
		return "b"
	case RuneC:
		return "C"
	case Runec:
		return "c"
	case RuneD:
		return "D"
	case Runed:
		return "d"
	case RuneE:
		return "E"
	case Runee:
		return "e"
	case RuneF:
		return "F"
	case Runef:
		return "f"
	case RuneG:
		return "G"
	case Runeg:
		return "g"
	case RuneH:
		return "H"
	case Runeh:
		return "h"
	case RuneI:
		return "I"
	case Runei:
		return "i"
	case RuneJ:
		return "J"
	case Runej:
		return "j"
	case RuneK:
		return "K"
	case Runek:
		return "k"
	case RuneL:
		return "L"
	case Runel:
		return "l"
	case RuneM:
		return "M"
	case Runem:
		return "m"
	case RuneN:
		return "N"
	case Runen:
		return "n"
	case RuneO:
		return "O"
	case Runeo:
		return "o"
	case RuneP:
		return "P"
	case Runep:
		return "p"
	case RuneQ:
		return "Q"
	case Runeq:
		return "q"
	case RuneR:
		return "R"
	case Runer:
		return "r"
	case RuneS:
		return "S"
	case Runes:
		return "s"
	case RuneT:
		return "T"
	case Runet:
		return "t"
	case RuneU:
		return "U"
	case Runeu:
		return "u"
	case RuneV:
		return "V"
	case Runev:
		return "v"
	case RuneW:
		return "W"
	case Runew:
		return "w"
	case RuneX:
		return "X"
	case Runex:
		return "x"
	case RuneY:
		return "Y"
	case Runey:
		return "y"
	case RuneZ:
		return "Z"
	case Runez:
		return "z"
	default: // "" or Empty
		return EmptyString
	}
}
