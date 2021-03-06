

















module IELE-SYNTAX
  imports IELE-COMMON

  syntax IeleName ::= r"(?<![A-Za-z0-9\\_\\.\\-\\$])[a-zA-Z\\.\\_\\$][0-9a-zA-Z\\.\\_\\-\\$]*" [token, notInRules, prec(3)]

  syntax Keyword ::= "load" | "store" | "sload"  | "sstore" | "iszero" | "not"  | "add"  | "mul"  | "sub" | "div"
                   | "exp"  | "mod"   | "addmod" | "mulmod" | "expmod" | "byte" | "sext" | "twos" | "and" | "or"
                   | "xor"  | "shift" | "lt"     | "le"     | "gt"     | "ge"   | "eq"   | "ne"   | "cmp" | "sha3"
  
  syntax Keyword ::= "br"   | "call"   | "staticcall" | "at" | "send"  | "gaslimit" | "ret"      | "void"   | "revert"
                   | "log"  | "create" | "copycreate" | "selfdestruct" | "contract" | "external" | "define" | "public"
                   | "log2" | "bswap"  | "calladdress"
  syntax IeleName ::= Keyword [token]

  syntax NumericIeleName ::= r"[0-9]+" [token]

  syntax StringIeleName ::= r"\\\"(([^\\\"\\\\])|(\\\\[0-9a-fA-F]{2}))*\\\"" [token]
endmodule





module IELE-COMMON
  imports DOMAINS-SYNTAX
  imports INT-SYNTAX
  syntax NumericIeleName ::= Int
  syntax StringIeleName
  syntax IeleName ::= NumericIeleName
  syntax IeleName ::= StringIeleName [klabel(StringIeleName), avoid, symbol, function]










  syntax GlobalName ::= "@" IeleName

  syntax LocalName ::= "%" IeleName

  syntax LocalNames ::= List{LocalName, ","} [klabel(localNameList)]








  syntax Constant ::= IntConstant | GlobalName
  syntax IntConstant ::= Int | Bool | HexConstant
  syntax HexConstant ::= r"0x[0-9a-fA-F]+" [token]















  syntax LValue ::= LocalName
  
  syntax LValues ::= NeList{LValue, ","} [klabel(lvalueList)]

  syntax Operand ::= LValue | Constant

  syntax Operands ::= List{Operand, ","} [klabel(operandList)]

  syntax NonEmptyOperands ::= NeList{Operand, ","} [klabel(operandList)]

  syntax Ints ::= List{Int, ","} [klabel(operandList)]

  syntax Operands ::= Ints

  syntax Operands ::= NonEmptyOperands

  syntax Ints ::= NonEmptyInts

  syntax NonEmptyOperands ::= NonEmptyInts

  syntax NonEmptyInts ::= NeList{Int, ","} [klabel(operandList)]







  syntax AssignInst ::= LValue "=" Operand







  syntax LoadInst ::= LValue "=" "load" /* cell */ Operand [hybrid, strict(2)]
  syntax LoadInst ::= LValue "=" "load" /* cell */ Operand "," /* offset in bytes */ Operand "," /* width in bytes */ Operand [hybrid, seqstrict(2,3,4)]
  syntax StoreInst ::= "store" /* value */ Operand "," /* cell */ Operand [hybrid, seqstrict(1,2)]
  syntax StoreInst ::= "store" /* value */ Operand "," /* cell */ Operand "," /* offset in bytes */ Operand "," /* width in bytes */ Operand [hybrid, seqstrict(1,2,3,4)]







  syntax SLoadInst ::= LValue "=" "sload" /* index */ Operand [hybrid, strict(2)]
  syntax SStoreInst ::= "sstore" /* value */ Operand "," /* index */ Operand [hybrid, seqstrict(1,2)]







  syntax IsZeroInst ::= LValue "=" "iszero" Operand [hybrid, strict(2)]
  syntax NotInst    ::= LValue "=" "not"    Operand [hybrid, strict(2)]

  syntax AddInst ::= LValue "=" "add" Operand "," Operand [hybrid, seqstrict(2,3)]
  syntax MulInst ::= LValue "=" "mul" Operand "," Operand [hybrid, seqstrict(2,3)]
  syntax SubInst ::= LValue "=" "sub" Operand "," Operand [hybrid, seqstrict(2,3)]
  syntax DivInst ::= LValue "=" "div" Operand "," Operand [hybrid, seqstrict(2,3)]
  syntax ExpInst ::= LValue "=" "exp" Operand "," Operand [hybrid, seqstrict(2,3)]
  syntax ModInst ::= LValue "=" "mod" Operand "," Operand [hybrid, seqstrict(2,3)]

  syntax AddModInst ::= LValue "=" "addmod" Operand "," Operand "," Operand [hybrid, seqstrict(2,3,4)]
  syntax MulModInst ::= LValue "=" "mulmod" Operand "," Operand "," Operand [hybrid, seqstrict(2,3,4)]
  syntax ExpModInst ::= LValue "=" "expmod" Operand "," Operand "," Operand [hybrid, seqstrict(2,3,4)]

  syntax LogInst   ::= LValue "=" "log2"                                    Operand             [hybrid, strict(2)]
  syntax ByteInst  ::= LValue "=" "byte"  /* byte index, 0 being the LSB */ Operand "," Operand [hybrid, seqstrict(2,3)]
  syntax SExtInst  ::= LValue "=" "sext"  /* width in bytes */              Operand "," Operand [hybrid, seqstrict(2,3)]
  syntax TwosInst  ::= LValue "=" "twos"  /* width in bytes */              Operand "," Operand [hybrid, seqstrict(2,3)]
  syntax BswapInst ::= LValue "=" "bswap" /* width in bytes */              Operand "," Operand [hybrid, seqstrict(2,3)]

  syntax AndInst   ::= LValue "=" "and"   Operand "," Operand [hybrid, seqstrict(2,3)]
  syntax OrInst    ::= LValue "=" "or"    Operand "," Operand [hybrid, seqstrict(2,3)]
  syntax XorInst   ::= LValue "=" "xor"   Operand "," Operand [hybrid, seqstrict(2,3)]
  syntax ShiftInst ::= LValue "=" "shift" Operand "," /* shift amount */ Operand [hybrid, seqstrict(2,3)]

  syntax Predicate ::= "lt" | "le" | "gt" | "ge" | "eq" | "ne"
  syntax CmpInst ::= LValue "=" "cmp" Predicate Operand "," Operand [hybrid, seqstrict(3,4)]

  syntax SHA3Inst ::= LValue "=" "sha3" Operand [hybrid, strict(2)]







  syntax JumpInst     ::= "br" IeleName
  syntax CondJumpInst ::= "br" Operand "," IeleName [hybrid, strict(1)]







  syntax LocalCallInst   ::= "call" Operand "(" Operands ")"
                           | LValues "=" "call" Operand "(" Operands ")" [hybrid, strict(2,3)]
  syntax AccountCallInst ::= LValues "=" "call" Operand "at" Operand "(" Operands ")" "send" Operand "," "gaslimit" Operand [hybrid, seqstrict(2,3,4,5,6)]
  syntax AccountCallInst ::= LValues "=" "staticcall" Operand "at" Operand "(" Operands ")" "gaslimit" Operand [hybrid, seqstrict(2,3,4,5)]

  syntax CallAddressInst ::= LValue "=" "calladdress" GlobalName "at" Operand [hybrid, strict(3)]

  syntax ReturnInst ::= "ret" NonEmptyOperands [hybrid, strict(1)]
                      | "ret" "void"
  syntax RevertInst ::= "revert" Operand [hybrid, strict(1)]







  syntax LogInst ::= "log" /* cell */ Operand [hybrid, strict(1)]
                   | "log" /* cell */ Operand "," NonEmptyOperands [hybrid, seqstrict(1,2)]







  syntax CreateInst ::= /* exit status */ LValue "," /* new account address */ LValue "=" "create"     /* contract name */    IeleName "(" Operands ")" "send" Operand [hybrid, seqstrict(4,5)]
  syntax CreateInst ::= /* exit status */ LValue "," /* new account address */ LValue "=" "copycreate" /* contract address */ Operand  "(" Operands ")" "send" Operand [hybrid, seqstrict(3,4,5)]

  syntax SelfdestructInst ::= "selfdestruct" /* account to send balance */ Operand [hybrid, strict(1)]







  syntax IeleName ::= "iele.invalid"     [token]
  // local state queries
  syntax IeleName ::= "iele.gas"         [token]
  syntax IeleName ::= "iele.gasprice"    [token]
  syntax IeleName ::= "iele.gaslimit"    [token]
  syntax IeleName ::= "iele.beneficiary" [token]
  syntax IeleName ::= "iele.timestamp"   [token]
  syntax IeleName ::= "iele.number"      [token]
  syntax IeleName ::= "iele.difficulty"  [token]
  syntax IeleName ::= "iele.address"     [token]
  syntax IeleName ::= "iele.origin"      [token]
  syntax IeleName ::= "iele.caller"      [token]
  syntax IeleName ::= "iele.callvalue"   [token]
  syntax IeleName ::= "iele.msize"       [token]
  syntax IeleName ::= "iele.codesize"    [token]
  syntax IeleName ::= "iele.blockhash"   [token]
  // account queries
  syntax IeleName ::= "iele.balance"     [token]
  syntax IeleName ::= "iele.extcodesize" [token]







  syntax IeleName ::= "iele.ecrec"     [token]
  syntax IeleName ::= "iele.sha256"    [token]
  syntax IeleName ::= "iele.rip160"    [token]
  syntax IeleName ::= "iele.id"        [token]
  syntax IeleName ::= "iele.ecadd"     [token]
  syntax IeleName ::= "iele.ecmul"     [token]
  syntax IeleName ::= "iele.ecpairing" [token]





  syntax Instruction ::=
    AssignInst
  | LoadInst
  | StoreInst
  | SLoadInst
  | SStoreInst
  | IsZeroInst
  | NotInst
  | AddInst
  | MulInst
  | SubInst
  | DivInst
  | ExpInst
  | ModInst
  | AddModInst
  | MulModInst
  | ExpModInst
  | ByteInst
  | SExtInst
  | TwosInst
  | BswapInst
  | AndInst
  | OrInst
  | XorInst
  | ShiftInst
  | CmpInst
  | SHA3Inst
  | JumpInst
  | CondJumpInst
  | LocalCallInst
  | AccountCallInst
  | CallAddressInst
  | ReturnInst
  | RevertInst
  | LogInst
  | CreateInst
  | SelfdestructInst

  syntax Instructions ::= List{Instruction, ""} [klabel(instructionList), format(%1%2%n%3)]

















  syntax TopLevelDefinition ::=
    FunctionDefinition
  | GlobalDefinition
  | ContractDeclaration

  syntax TopLevelDefinitions ::= List{TopLevelDefinition, ""} [klabel(topLevelDefinitionList), format(%1%2%n%3)]

  syntax ContractDefinition ::= "contract" IeleName "{" TopLevelDefinitions "}" [format(%1 %2 %3%n%i%4%d%n%5)]

  syntax Contract ::= List{ContractDefinition, ""} [klabel(contractDefinitionList)]







  syntax ContractDeclaration ::= "external" "contract" IeleName







  syntax GlobalDefinition ::= GlobalName "=" IntConstant  [klabel(globalDefinition)]







  syntax FunctionSignature ::= GlobalName "(" FunctionParameters ")"

  syntax FunctionParameters ::= LocalNames

  syntax FunctionDefinition ::= 
    "define" FunctionSignature "{" Blocks "}"             [format(%1 %2 %3%i%n%4%d%n%5%n)]
  | "define" "public" FunctionSignature "{" Blocks "}" [format(%1 %2 %3 %4%i%n%5%d%n%6%n)]







  syntax LabeledBlock ::= IeleName ":" Instructions

  syntax LabeledBlocks ::= List{LabeledBlock, ""} [klabel(labeledBlockList)]

  syntax UnlabeledBlock ::= Instructions

  syntax Blocks [flatPredicate]
  syntax Blocks ::= UnlabeledBlock LabeledBlocks [avoid] | LabeledBlocks








  syntax IeleName ::= "deposit" [token]

  syntax IeleName ::= "init" [token]







  rule call NAME ( ARGS ) => .LValues = call NAME ( ARGS ) [alias]
  rule ret void => ret .NonEmptyOperands [alias]
endmodule

