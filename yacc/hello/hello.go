// Code generated by goyacc -o hello.go -p hello hello.y. DO NOT EDIT.

//line hello.y:2
package hello

import __yyfmt__ "fmt"

//line hello.y:2

import "hellogo/yacc/hello/sem/tree"

//line hello.y:7
type helloSymType struct {
	yys         int
	type_string string
}

const PRINT = 57346
const HELLO = 57347

var helloToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"PRINT",
	"HELLO",
}

var helloStatenames = [...]string{}

const helloEofCode = 1
const helloErrCode = 2
const helloInitialStackSize = 16

//line hello.y:23

//line yacctab:1
var helloExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
}

const helloPrivate = 57344

const helloLast = 2

var helloAct = [...]int8{
	2, 1,
}

var helloPact = [...]int16{
	-5, -1000, -1000,
}

var helloPgo = [...]int8{
	0, 1,
}

var helloR1 = [...]int8{
	0, 1,
}

var helloR2 = [...]int8{
	0, 1,
}

var helloChk = [...]int16{
	-1000, -1, 5,
}

var helloDef = [...]int8{
	0, -2, 1,
}

var helloTok1 = [...]int8{
	1,
}

var helloTok2 = [...]int8{
	2, 3, 4, 5,
}

var helloTok3 = [...]int8{
	0,
}

var helloErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	helloDebug        = 0
	helloErrorVerbose = false
)

type helloLexer interface {
	Lex(lval *helloSymType) int
	Error(s string)
}

type helloParser interface {
	Parse(helloLexer) int
	Lookahead() int
}

type helloParserImpl struct {
	lval  helloSymType
	stack [helloInitialStackSize]helloSymType
	char  int
}

func (p *helloParserImpl) Lookahead() int {
	return p.char
}

func helloNewParser() helloParser {
	return &helloParserImpl{}
}

const helloFlag = -1000

func helloTokname(c int) string {
	if c >= 1 && c-1 < len(helloToknames) {
		if helloToknames[c-1] != "" {
			return helloToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func helloStatname(s int) string {
	if s >= 0 && s < len(helloStatenames) {
		if helloStatenames[s] != "" {
			return helloStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func helloErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !helloErrorVerbose {
		return "syntax error"
	}

	for _, e := range helloErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + helloTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := int(helloPact[state])
	for tok := TOKSTART; tok-1 < len(helloToknames); tok++ {
		if n := base + tok; n >= 0 && n < helloLast && int(helloChk[int(helloAct[n])]) == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if helloDef[state] == -2 {
		i := 0
		for helloExca[i] != -1 || int(helloExca[i+1]) != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; helloExca[i] >= 0; i += 2 {
			tok := int(helloExca[i])
			if tok < TOKSTART || helloExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if helloExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += helloTokname(tok)
	}
	return res
}

func hellolex1(lex helloLexer, lval *helloSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = int(helloTok1[0])
		goto out
	}
	if char < len(helloTok1) {
		token = int(helloTok1[char])
		goto out
	}
	if char >= helloPrivate {
		if char < helloPrivate+len(helloTok2) {
			token = int(helloTok2[char-helloPrivate])
			goto out
		}
	}
	for i := 0; i < len(helloTok3); i += 2 {
		token = int(helloTok3[i+0])
		if token == char {
			token = int(helloTok3[i+1])
			goto out
		}
	}

out:
	if token == 0 {
		token = int(helloTok2[1]) /* unknown char */
	}
	if helloDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", helloTokname(token), uint(char))
	}
	return char, token
}

func helloParse(hellolex helloLexer) int {
	return helloNewParser().Parse(hellolex)
}

func (hellorcvr *helloParserImpl) Parse(hellolex helloLexer) int {
	var hellon int
	var helloVAL helloSymType
	var helloDollar []helloSymType
	_ = helloDollar // silence set and not used
	helloS := hellorcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	hellostate := 0
	hellorcvr.char = -1
	hellotoken := -1 // hellorcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		hellostate = -1
		hellorcvr.char = -1
		hellotoken = -1
	}()
	hellop := -1
	goto hellostack

ret0:
	return 0

ret1:
	return 1

hellostack:
	/* put a state and value onto the stack */
	if helloDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", helloTokname(hellotoken), helloStatname(hellostate))
	}

	hellop++
	if hellop >= len(helloS) {
		nyys := make([]helloSymType, len(helloS)*2)
		copy(nyys, helloS)
		helloS = nyys
	}
	helloS[hellop] = helloVAL
	helloS[hellop].yys = hellostate

hellonewstate:
	hellon = int(helloPact[hellostate])
	if hellon <= helloFlag {
		goto hellodefault /* simple state */
	}
	if hellorcvr.char < 0 {
		hellorcvr.char, hellotoken = hellolex1(hellolex, &hellorcvr.lval)
	}
	hellon += hellotoken
	if hellon < 0 || hellon >= helloLast {
		goto hellodefault
	}
	hellon = int(helloAct[hellon])
	if int(helloChk[hellon]) == hellotoken { /* valid shift */
		hellorcvr.char = -1
		hellotoken = -1
		helloVAL = hellorcvr.lval
		hellostate = hellon
		if Errflag > 0 {
			Errflag--
		}
		goto hellostack
	}

hellodefault:
	/* default state action */
	hellon = int(helloDef[hellostate])
	if hellon == -2 {
		if hellorcvr.char < 0 {
			hellorcvr.char, hellotoken = hellolex1(hellolex, &hellorcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if helloExca[xi+0] == -1 && int(helloExca[xi+1]) == hellostate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			hellon = int(helloExca[xi+0])
			if hellon < 0 || hellon == hellotoken {
				break
			}
		}
		hellon = int(helloExca[xi+1])
		if hellon < 0 {
			goto ret0
		}
	}
	if hellon == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			hellolex.Error(helloErrorMessage(hellostate, hellotoken))
			Nerrs++
			if helloDebug >= 1 {
				__yyfmt__.Printf("%s", helloStatname(hellostate))
				__yyfmt__.Printf(" saw %s\n", helloTokname(hellotoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for hellop >= 0 {
				hellon = int(helloPact[helloS[hellop].yys]) + helloErrCode
				if hellon >= 0 && hellon < helloLast {
					hellostate = int(helloAct[hellon]) /* simulate a shift of "error" */
					if int(helloChk[hellostate]) == helloErrCode {
						goto hellostack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if helloDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", helloS[hellop].yys)
				}
				hellop--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if helloDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", helloTokname(hellotoken))
			}
			if hellotoken == helloEofCode {
				goto ret1
			}
			hellorcvr.char = -1
			hellotoken = -1
			goto hellonewstate /* try again in the same state */
		}
	}

	/* reduction by production hellon */
	if helloDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", hellon, helloStatname(hellostate))
	}

	hellont := hellon
	hellopt := hellop
	_ = hellopt // guard against "declared and not used"

	hellop -= int(helloR2[hellon])
	// hellop is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if hellop+1 >= len(helloS) {
		nyys := make([]helloSymType, len(helloS)*2)
		copy(nyys, helloS)
		helloS = nyys
	}
	helloVAL = helloS[hellop+1]

	/* consult goto table to find next state */
	hellon = int(helloR1[hellon])
	hellog := int(helloPgo[hellon])
	helloj := hellog + helloS[hellop].yys + 1

	if helloj >= helloLast {
		hellostate = int(helloAct[hellog])
	} else {
		hellostate = int(helloAct[helloj])
		if int(helloChk[hellostate]) != -hellon {
			hellostate = int(helloAct[hellog])
		}
	}
	// dummy call; replaced with literal code
	switch hellont {

	case 1:
		helloDollar = helloS[hellopt-1 : hellopt+1]
//line hello.y:16
		{
			st := tree.HelloStatement{
				Name: "自定义参数",
			}

			st.Exec()
		}
	}
	goto hellostack /* stack new state and value */
}
