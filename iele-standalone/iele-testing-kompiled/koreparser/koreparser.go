// Code generated by goyacc -o koreparser.go -p kore koreparser.y. DO NOT EDIT.

//line koreparser.y:8

package koreparser

import __yyfmt__ "fmt"

//line koreparser.y:9

//line koreparser.y:13
type koreSymType struct {
	yys   int
	str   []byte
	k     K
	kseq  KSequence
	klist []K
}

const KSEQ = 57346
const DOTK = 57347
const DOTKLIST = 57348
const TOKENLABEL = 57349
const KLABELLABEL = 57350
const KVARIABLE = 57351
const KLABEL = 57352
const STRING = 57353

var koreToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"KSEQ",
	"DOTK",
	"'('",
	"')'",
	"','",
	"DOTKLIST",
	"TOKENLABEL",
	"KLABELLABEL",
	"KVARIABLE",
	"KLABEL",
	"STRING",
}
var koreStatenames = [...]string{}

const koreEofCode = 1
const koreErrCode = 2
const koreInitialStackSize = 16

//line koreparser.y:96

var lastResult K

func Parse(kast []byte) K {
	lexer := koreLexerImpl{line: kast}
	koreParse(&lexer)
	return lastResult
}

//line yacctab:1
var koreExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const korePrivate = 57344

const koreLast = 44

var koreAct = [...]int{

	2, 9, 28, 21, 27, 20, 7, 6, 8, 5,
	9, 22, 23, 18, 19, 7, 6, 8, 5, 7,
	6, 8, 5, 3, 26, 9, 25, 14, 29, 30,
	7, 6, 8, 5, 15, 16, 24, 13, 12, 11,
	10, 1, 4, 17,
}
var korePact = [...]int{

	20, -1000, -1000, 36, 35, 32, 31, 21, -1000, -1000,
	9, 9, 5, -8, -11, -1000, -1000, 4, -1000, -1000,
	29, 18, -1000, -4, -1000, -12, -1000, 20, 22, -1000,
	-1000,
}
var korePgo = [...]int{

	0, 0, 23, 43, 42, 41,
}
var koreR1 = [...]int{

	0, 5, 4, 4, 4, 1, 1, 2, 2, 2,
	2, 3, 3, 3, 3,
}
var koreR2 = [...]int{

	0, 1, 3, 3, 1, 1, 1, 4, 4, 6,
	1, 1, 3, 4, 1,
}
var koreChk = [...]int{

	-1000, -5, -1, -2, -4, 13, 11, 10, 12, 5,
	4, 4, 6, 6, 6, -2, -2, -3, -1, 9,
	13, 14, 7, 8, 7, 8, -1, 8, 14, -1,
	7,
}
var koreDef = [...]int{

	0, -2, 1, 5, 6, 0, 0, 0, 10, 4,
	0, 0, 0, 0, 0, 2, 3, 0, 11, 14,
	0, 0, 7, 0, 8, 0, 12, 0, 0, 13,
	9,
}
var koreTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	6, 7, 3, 3, 8,
}
var koreTok2 = [...]int{

	2, 3, 4, 5, 9, 10, 11, 12, 13, 14,
}
var koreTok3 = [...]int{
	0,
}

var koreErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	koreDebug        = 0
	koreErrorVerbose = false
)

type koreLexer interface {
	Lex(lval *koreSymType) int
	Error(s string)
}

type koreParser interface {
	Parse(koreLexer) int
	Lookahead() int
}

type koreParserImpl struct {
	lval  koreSymType
	stack [koreInitialStackSize]koreSymType
	char  int
}

func (p *koreParserImpl) Lookahead() int {
	return p.char
}

func koreNewParser() koreParser {
	return &koreParserImpl{}
}

const koreFlag = -1000

func koreTokname(c int) string {
	if c >= 1 && c-1 < len(koreToknames) {
		if koreToknames[c-1] != "" {
			return koreToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func koreStatname(s int) string {
	if s >= 0 && s < len(koreStatenames) {
		if koreStatenames[s] != "" {
			return koreStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func koreErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !koreErrorVerbose {
		return "syntax error"
	}

	for _, e := range koreErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + koreTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := korePact[state]
	for tok := TOKSTART; tok-1 < len(koreToknames); tok++ {
		if n := base + tok; n >= 0 && n < koreLast && koreChk[koreAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if koreDef[state] == -2 {
		i := 0
		for koreExca[i] != -1 || koreExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; koreExca[i] >= 0; i += 2 {
			tok := koreExca[i]
			if tok < TOKSTART || koreExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if koreExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += koreTokname(tok)
	}
	return res
}

func korelex1(lex koreLexer, lval *koreSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = koreTok1[0]
		goto out
	}
	if char < len(koreTok1) {
		token = koreTok1[char]
		goto out
	}
	if char >= korePrivate {
		if char < korePrivate+len(koreTok2) {
			token = koreTok2[char-korePrivate]
			goto out
		}
	}
	for i := 0; i < len(koreTok3); i += 2 {
		token = koreTok3[i+0]
		if token == char {
			token = koreTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = koreTok2[1] /* unknown char */
	}
	if koreDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", koreTokname(token), uint(char))
	}
	return char, token
}

func koreParse(korelex koreLexer) int {
	return koreNewParser().Parse(korelex)
}

func (korercvr *koreParserImpl) Parse(korelex koreLexer) int {
	var koren int
	var koreVAL koreSymType
	var koreDollar []koreSymType
	_ = koreDollar // silence set and not used
	koreS := korercvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	korestate := 0
	korercvr.char = -1
	koretoken := -1 // korercvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		korestate = -1
		korercvr.char = -1
		koretoken = -1
	}()
	korep := -1
	goto korestack

ret0:
	return 0

ret1:
	return 1

korestack:
	/* put a state and value onto the stack */
	if koreDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", koreTokname(koretoken), koreStatname(korestate))
	}

	korep++
	if korep >= len(koreS) {
		nyys := make([]koreSymType, len(koreS)*2)
		copy(nyys, koreS)
		koreS = nyys
	}
	koreS[korep] = koreVAL
	koreS[korep].yys = korestate

korenewstate:
	koren = korePact[korestate]
	if koren <= koreFlag {
		goto koredefault /* simple state */
	}
	if korercvr.char < 0 {
		korercvr.char, koretoken = korelex1(korelex, &korercvr.lval)
	}
	koren += koretoken
	if koren < 0 || koren >= koreLast {
		goto koredefault
	}
	koren = koreAct[koren]
	if koreChk[koren] == koretoken { /* valid shift */
		korercvr.char = -1
		koretoken = -1
		koreVAL = korercvr.lval
		korestate = koren
		if Errflag > 0 {
			Errflag--
		}
		goto korestack
	}

koredefault:
	/* default state action */
	koren = koreDef[korestate]
	if koren == -2 {
		if korercvr.char < 0 {
			korercvr.char, koretoken = korelex1(korelex, &korercvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if koreExca[xi+0] == -1 && koreExca[xi+1] == korestate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			koren = koreExca[xi+0]
			if koren < 0 || koren == koretoken {
				break
			}
		}
		koren = koreExca[xi+1]
		if koren < 0 {
			goto ret0
		}
	}
	if koren == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			korelex.Error(koreErrorMessage(korestate, koretoken))
			Nerrs++
			if koreDebug >= 1 {
				__yyfmt__.Printf("%s", koreStatname(korestate))
				__yyfmt__.Printf(" saw %s\n", koreTokname(koretoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for korep >= 0 {
				koren = korePact[koreS[korep].yys] + koreErrCode
				if koren >= 0 && koren < koreLast {
					korestate = koreAct[koren] /* simulate a shift of "error" */
					if koreChk[korestate] == koreErrCode {
						goto korestack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if koreDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", koreS[korep].yys)
				}
				korep--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if koreDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", koreTokname(koretoken))
			}
			if koretoken == koreEofCode {
				goto ret1
			}
			korercvr.char = -1
			koretoken = -1
			goto korenewstate /* try again in the same state */
		}
	}

	/* reduction by production koren */
	if koreDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", koren, koreStatname(korestate))
	}

	korent := koren
	korept := korep
	_ = korept // guard against "declared and not used"

	korep -= koreR2[koren]
	// korep is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if korep+1 >= len(koreS) {
		nyys := make([]koreSymType, len(koreS)*2)
		copy(nyys, koreS)
		koreS = nyys
	}
	koreVAL = koreS[korep+1]

	/* consult goto table to find next state */
	koren = koreR1[koren]
	koreg := korePgo[koren]
	korej := koreg + koreS[korep].yys + 1

	if korej >= koreLast {
		korestate = koreAct[koreg]
	} else {
		korestate = koreAct[korej]
		if koreChk[korestate] != -koren {
			korestate = koreAct[koreg]
		}
	}
	// dummy call; replaced with literal code
	switch korent {

	case 1:
		koreDollar = koreS[korept-1 : korept+1]
//line koreparser.y:32
		{
			lastResult = koreDollar[1].k
		}
	case 2:
		koreDollar = koreS[korept-3 : korept+1]
//line koreparser.y:38
		{
			koreVAL.kseq = KSequence([]K{koreDollar[1].k, koreDollar[3].k})
		}
	case 3:
		koreDollar = koreS[korept-3 : korept+1]
//line koreparser.y:42
		{
			koreVAL.kseq = append(koreDollar[1].kseq, koreDollar[3].k)
		}
	case 4:
		koreDollar = koreS[korept-1 : korept+1]
//line koreparser.y:46
		{
			koreVAL.kseq = nil
		}
	case 5:
		koreDollar = koreS[korept-1 : korept+1]
//line koreparser.y:52
		{
			koreVAL.k = koreDollar[1].k
		}
	case 6:
		koreDollar = koreS[korept-1 : korept+1]
//line koreparser.y:56
		{
			koreVAL.k = koreDollar[1].kseq
		}
	case 7:
		koreDollar = koreS[korept-4 : korept+1]
//line koreparser.y:62
		{
			koreVAL.k = KApply{Label: string(koreDollar[1].str), List: koreDollar[3].klist}
		}
	case 8:
		koreDollar = koreS[korept-4 : korept+1]
//line koreparser.y:66
		{
			koreVAL.k = InjectedKLabel{Label: string(koreDollar[3].str)}
		}
	case 9:
		koreDollar = koreS[korept-6 : korept+1]
//line koreparser.y:70
		{
			koreVAL.k = KToken{Value: string(koreDollar[3].str), Sort: string(koreDollar[5].str)}
		}
	case 10:
		koreDollar = koreS[korept-1 : korept+1]
//line koreparser.y:74
		{
			koreVAL.k = KVariable{Name: string(koreDollar[1].str)}
		}
	case 11:
		koreDollar = koreS[korept-1 : korept+1]
//line koreparser.y:80
		{
			koreVAL.klist = []K{koreDollar[1].k}
		}
	case 12:
		koreDollar = koreS[korept-3 : korept+1]
//line koreparser.y:84
		{
			koreVAL.klist = append(koreDollar[1].klist, koreDollar[3].k)
		}
	case 13:
		koreDollar = koreS[korept-4 : korept+1]
//line koreparser.y:88
		{
			koreVAL.klist = append(koreDollar[1].klist, koreDollar[4].k)
		}
	case 14:
		koreDollar = koreS[korept-1 : korept+1]
//line koreparser.y:92
		{
			koreVAL.klist = nil
		}
	}
	goto korestack /* stack new state and value */
}
