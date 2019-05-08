package ieletestingmodel 

// Sort ... a K sort identifier
type Sort int

// SortCallerCellOpt ... CallerCellOpt
const SortCallerCellOpt Sort = 0

// SortSubstateCellFragment ... SubstateCellFragment
const SortSubstateCellFragment Sort = 1

// SortTxNonceCellOpt ... TxNonceCellOpt
const SortTxNonceCellOpt Sort = 2

// SortModeCellOpt ... ModeCellOpt
const SortModeCellOpt Sort = 3

// SortGeneratedTopCell ... GeneratedTopCell
const SortGeneratedTopCell Sort = 4

// SortK ... K
const SortK Sort = 5

// SortInts ... Ints
const SortInts Sort = 6

// SortQuadOp ... QuadOp
const SortQuadOp Sort = 7

// SortStorageCellOpt ... StorageCellOpt
const SortStorageCellOpt Sort = 8

// SortScheduleCell ... ScheduleCell
const SortScheduleCell Sort = 9

// SortCurrentFunctionCellOpt ... CurrentFunctionCellOpt
const SortCurrentFunctionCellOpt Sort = 10

// SortIELECommand ... IELECommand
const SortIELECommand Sort = 11

// SortNotInst ... NotInst
const SortNotInst Sort = 12

// SortGasUsedCell ... GasUsedCell
const SortGasUsedCell Sort = 13

// SortArgsCell ... ArgsCell
const SortArgsCell Sort = 14

// SortFunctionsCellOpt ... FunctionsCellOpt
const SortFunctionsCellOpt Sort = 15

// SortContractDefinition ... ContractDefinition
const SortContractDefinition Sort = 16

// SortRevertInst ... RevertInst
const SortRevertInst Sort = 17

// SortInstructionsCell ... InstructionsCell
const SortInstructionsCell Sort = 18

// SortCurrentFunctionCellFragment ... CurrentFunctionCellFragment
const SortCurrentFunctionCellFragment Sort = 19

// SortFuncIDsCell ... FuncIdsCell
const SortFuncIDsCell Sort = 20

// SortInstructionsCellOpt ... InstructionsCellOpt
const SortInstructionsCellOpt Sort = 21

// SortActiveAccountsCell ... ActiveAccountsCell
const SortActiveAccountsCell Sort = 22

// SortOrInst ... OrInst
const SortOrInst Sort = 23

// SortCodeCellOpt ... CodeCellOpt
const SortCodeCellOpt Sort = 24

// SortSendtoCell ... SendtoCell
const SortSendtoCell Sort = 25

// SortExportedCell ... ExportedCell
const SortExportedCell Sort = 26

// SortTypeCheckingCell ... TypeCheckingCell
const SortTypeCheckingCell Sort = 27

// SortFunctionCell ... FunctionCell
const SortFunctionCell Sort = 28

// SortBalanceCellOpt ... BalanceCellOpt
const SortBalanceCellOpt Sort = 29

// SortAndInst ... AndInst
const SortAndInst Sort = 30

// SortSelfdestructInst ... SelfdestructInst
const SortSelfdestructInst Sort = 31

// SortBinOp ... BinOp
const SortBinOp Sort = 32

// SortFunctionSignature ... FunctionSignature
const SortFunctionSignature Sort = 33

// SortTxNonceCell ... TxNonceCell
const SortTxNonceCell Sort = 34

// SortBlockhashCellOpt ... BlockhashCellOpt
const SortBlockhashCellOpt Sort = 35

// SortMsgIDCell ... MsgIDCell
const SortMsgIDCell Sort = 36

// SortRegsCell ... RegsCell
const SortRegsCell Sort = 37

// SortFuncIDCellOpt ... FuncIdCellOpt
const SortFuncIDCellOpt Sort = 38

// SortReturnOp ... ReturnOp
const SortReturnOp Sort = 39

// SortWellFormednessCellFragment ... WellFormednessCellFragment
const SortWellFormednessCellFragment Sort = 40

// SortValueCellOpt ... ValueCellOpt
const SortValueCellOpt Sort = 41

// SortProgramSizeCellOpt ... ProgramSizeCellOpt
const SortProgramSizeCellOpt Sort = 42

// SortBool ... Bool
const SortBool Sort = 43

// SortOpCode ... OpCode
const SortOpCode Sort = 44

// SortBeneficiaryCell ... BeneficiaryCell
const SortBeneficiaryCell Sort = 45

// SortNonceCellOpt ... NonceCellOpt
const SortNonceCellOpt Sort = 46

// SortKCell ... KCell
const SortKCell Sort = 47

// SortGasLimitCellOpt ... GasLimitCellOpt
const SortGasLimitCellOpt Sort = 48

// SortFuncIDCell ... FuncIdCell
const SortFuncIDCell Sort = 49

// SortLocalCallsCell ... LocalCallsCell
const SortLocalCallsCell Sort = 50

// SortCallFrameCellOpt ... CallFrameCellOpt
const SortCallFrameCellOpt Sort = 51

// SortPseudoInstruction ... PseudoInstruction
const SortPseudoInstruction Sort = 52

// SortLocalNames ... LocalNames
const SortLocalNames Sort = 53

// SortRefundCell ... RefundCell
const SortRefundCell Sort = 54

// SortOperand ... Operand
const SortOperand Sort = 55

// SortDataCellOpt ... DataCellOpt
const SortDataCellOpt Sort = 56

// SortNonEmptyInts ... NonEmptyInts
const SortNonEmptyInts Sort = 57

// SortStringIeleName ... StringIeleName
const SortStringIeleName Sort = 58

// SortBalanceCell ... BalanceCell
const SortBalanceCell Sort = 59

// SortStrategy ... Strategy
const SortStrategy Sort = 60

// SortGeneratedTopCellFragment ... GeneratedTopCellFragment
const SortGeneratedTopCellFragment Sort = 61

// SortKCellOpt ... KCellOpt
const SortKCellOpt Sort = 62

// SortNumberCellOpt ... NumberCellOpt
const SortNumberCellOpt Sort = 63

// SortTypesCell ... TypesCell
const SortTypesCell Sort = 64

// SortOriginCellOpt ... OriginCellOpt
const SortOriginCellOpt Sort = 65

// SortCurrentInstructionsCellOpt ... CurrentInstructionsCellOpt
const SortCurrentInstructionsCellOpt Sort = 66

// SortXhashUpperID ... #UpperId
const SortXhashUpperID Sort = 67

// SortExitCodeCell ... ExitCodeCell
const SortExitCodeCell Sort = 68

// SortMessageCell ... MessageCell
const SortMessageCell Sort = 69

// SortGasCellOpt ... GasCellOpt
const SortGasCellOpt Sort = 70

// SortProgramCell ... ProgramCell
const SortProgramCell Sort = 71

// SortUnlabeledBlock ... UnlabeledBlock
const SortUnlabeledBlock Sort = 72

// SortCallFrameCellFragment ... CallFrameCellFragment
const SortCallFrameCellFragment Sort = 73

// SortCreateOp ... CreateOp
const SortCreateOp Sort = 74

// SortContractDeclaration ... ContractDeclaration
const SortContractDeclaration Sort = 75

// SortNparamsCellOpt ... NparamsCellOpt
const SortNparamsCellOpt Sort = 76

// SortMode ... Mode
const SortMode Sort = 77

// SortConstant ... Constant
const SortConstant Sort = 78

// SortBeneficiaryCellOpt ... BeneficiaryCellOpt
const SortBeneficiaryCellOpt Sort = 79

// SortContractNameCell ... ContractNameCell
const SortContractNameCell Sort = 80

// SortLocalCallsCellOpt ... LocalCallsCellOpt
const SortLocalCallsCellOpt Sort = 81

// SortProgramSizeCell ... ProgramSizeCell
const SortProgramSizeCell Sort = 82

// SortLocalCall ... LocalCall
const SortLocalCall Sort = 83

// SortLogInst ... LogInst
const SortLogInst Sort = 84

// SortContractCodeCellOpt ... ContractCodeCellOpt
const SortContractCodeCellOpt Sort = 85

// SortLengthPrefix ... LengthPrefix
const SortLengthPrefix Sort = 86

// SortTypeCheckingCellOpt ... TypeCheckingCellOpt
const SortTypeCheckingCellOpt Sort = 87

// SortFunctionCellFragment ... FunctionCellFragment
const SortFunctionCellFragment Sort = 88

// SortIeleCellOpt ... IeleCellOpt
const SortIeleCellOpt Sort = 89

// SortJSONKey ... JSONKey
const SortJSONKey Sort = 90

// SortCallDepthCellOpt ... CallDepthCellOpt
const SortCallDepthCellOpt Sort = 91

// SortDataCell ... DataCell
const SortDataCell Sort = 92

// SortLogDataCell ... LogDataCell
const SortLogDataCell Sort = 93

// SortMulInst ... MulInst
const SortMulInst Sort = 94

// SortTopLevelDefinitions ... TopLevelDefinitions
const SortTopLevelDefinitions Sort = 95

// SortOutputCell ... OutputCell
const SortOutputCell Sort = 96

// SortMessageCellFragment ... MessageCellFragment
const SortMessageCellFragment Sort = 97

// SortSubstateStackCellOpt ... SubstateStackCellOpt
const SortSubstateStackCellOpt Sort = 98

// SortLocalCallInst ... LocalCallInst
const SortLocalCallInst Sort = 99

// SortSStoreInst ... SStoreInst
const SortSStoreInst Sort = 100

// SortCurrentMemoryCellOpt ... CurrentMemoryCellOpt
const SortCurrentMemoryCellOpt Sort = 101

// SortGlobalDefinition ... GlobalDefinition
const SortGlobalDefinition Sort = 102

// SortSubstateCell ... SubstateCell
const SortSubstateCell Sort = 103

// SortInterimStatesCell ... InterimStatesCell
const SortInterimStatesCell Sort = 104

// SortLabelsCellOpt ... LabelsCellOpt
const SortLabelsCellOpt Sort = 105

// SortAccountsCell ... AccountsCell
const SortAccountsCell Sort = 106

// SortAccountCell ... AccountCell
const SortAccountCell Sort = 107

// SortBlocks ... Blocks
const SortBlocks Sort = 108

// SortSubstateStackCell ... SubstateStackCell
const SortSubstateStackCell Sort = 109

// SortDifficultyCellOpt ... DifficultyCellOpt
const SortDifficultyCellOpt Sort = 110

// SortTwosInst ... TwosInst
const SortTwosInst Sort = 111

// SortScheduleConst ... ScheduleConst
const SortScheduleConst Sort = 112

// SortLogDataCellOpt ... LogDataCellOpt
const SortLogDataCellOpt Sort = 113

// SortWellFormednessCell ... WellFormednessCell
const SortWellFormednessCell Sort = 114

// SortString ... String
const SortString Sort = 115

// SortModeCell ... ModeCell
const SortModeCell Sort = 116

// SortStorageCell ... StorageCell
const SortStorageCell Sort = 117

// SortCurrentContractCell ... CurrentContractCell
const SortCurrentContractCell Sort = 118

// SortTxGasPriceCellOpt ... TxGasPriceCellOpt
const SortTxGasPriceCellOpt Sort = 119

// SortCurrentInstructionsCell ... CurrentInstructionsCell
const SortCurrentInstructionsCell Sort = 120

// SortExitCodeCellOpt ... ExitCodeCellOpt
const SortExitCodeCellOpt Sort = 121

// SortTxOrderCell ... TxOrderCell
const SortTxOrderCell Sort = 122

// SortStaticCell ... StaticCell
const SortStaticCell Sort = 123

// SortPeakMemoryCell ... PeakMemoryCell
const SortPeakMemoryCell Sort = 124

// SortAcctIDCell ... AcctIDCell
const SortAcctIDCell Sort = 125

// SortFunctionsCellFragment ... FunctionsCellFragment
const SortFunctionsCellFragment Sort = 126

// SortLValues ... LValues
const SortLValues Sort = 127

// SortAccount ... Account
const SortAccount Sort = 128

// SortXhashRuleTag ... #RuleTag
const SortXhashRuleTag Sort = 129

// SortTernOp ... TernOp
const SortTernOp Sort = 130

// SortCallAddressInst ... CallAddressInst
const SortCallAddressInst Sort = 131

// SortRegsCellOpt ... RegsCellOpt
const SortRegsCellOpt Sort = 132

// SortCallFrameCell ... CallFrameCell
const SortCallFrameCell Sort = 133

// SortBytes ... Bytes
const SortBytes Sort = 134

// SortGasUsedCellOpt ... GasUsedCellOpt
const SortGasUsedCellOpt Sort = 135

// SortFunctionParameters ... FunctionParameters
const SortFunctionParameters Sort = 136

// SortTypes ... Types
const SortTypes Sort = 137

// SortActiveAccountsCellOpt ... ActiveAccountsCellOpt
const SortActiveAccountsCellOpt Sort = 138

// SortStringBuffer ... StringBuffer
const SortStringBuffer Sort = 139

// SortExportedCellOpt ... ExportedCellOpt
const SortExportedCellOpt Sort = 140

// SortModInst ... ModInst
const SortModInst Sort = 141

// SortFuncLabelsCellOpt ... FuncLabelsCellOpt
const SortFuncLabelsCellOpt Sort = 142

// SortScheduleFlag ... ScheduleFlag
const SortScheduleFlag Sort = 143

// SortCheckGasCellOpt ... CheckGasCellOpt
const SortCheckGasCellOpt Sort = 144

// SortSHA3Inst ... SHA3Inst
const SortSHA3Inst Sort = 145

// SortStoreInst ... StoreInst
const SortStoreInst Sort = 146

// SortOperands ... Operands
const SortOperands Sort = 147

// SortSubstateLogEntry ... SubstateLogEntry
const SortSubstateLogEntry Sort = 148

// SortFuncCellOpt ... FuncCellOpt
const SortFuncCellOpt Sort = 149

// SortJumpTableCellOpt ... JumpTableCellOpt
const SortJumpTableCellOpt Sort = 150

// SortLabeledBlock ... LabeledBlock
const SortLabeledBlock Sort = 151

// SortFuncLabelsCell ... FuncLabelsCell
const SortFuncLabelsCell Sort = 152

// SortAssignInst ... AssignInst
const SortAssignInst Sort = 153

// SortLocalName ... LocalName
const SortLocalName Sort = 154

// SortExpInst ... ExpInst
const SortExpInst Sort = 155

// SortIsZeroInst ... IsZeroInst
const SortIsZeroInst Sort = 156

// SortInterimStatesCellOpt ... InterimStatesCellOpt
const SortInterimStatesCellOpt Sort = 157

// SortLengthPrefixType ... LengthPrefixType
const SortLengthPrefixType Sort = 158

// SortReturnType ... ReturnType
const SortReturnType Sort = 159

// SortAccounts ... Accounts
const SortAccounts Sort = 160

// SortJumpInst ... JumpInst
const SortJumpInst Sort = 161

// SortGasLimitCell ... GasLimitCell
const SortGasLimitCell Sort = 162

// SortBlockhashCell ... BlockhashCell
const SortBlockhashCell Sort = 163

// SortReturnInst ... ReturnInst
const SortReturnInst Sort = 164

// SortDifficultyCell ... DifficultyCell
const SortDifficultyCell Sort = 165

// SortCmpInst ... CmpInst
const SortCmpInst Sort = 166

// SortIDCellOpt ... IdCellOpt
const SortIDCellOpt Sort = 167

// SortException ... Exception
const SortException Sort = 168

// SortCheckGasCell ... CheckGasCell
const SortCheckGasCell Sort = 169

// SortCopyCreateOp ... CopyCreateOp
const SortCopyCreateOp Sort = 170

// SortPreviousGasCell ... PreviousGasCell
const SortPreviousGasCell Sort = 171

// SortXorInst ... XorInst
const SortXorInst Sort = 172

// SortPrecompiledOp ... PrecompiledOp
const SortPrecompiledOp Sort = 173

// SortTxGasLimitCell ... TxGasLimitCell
const SortTxGasLimitCell Sort = 174

// SortTimestampCellOpt ... TimestampCellOpt
const SortTimestampCellOpt Sort = 175

// SortKItem ... KItem
const SortKItem Sort = 176

// SortJSON ... JSON
const SortJSON Sort = 177

// SortCallDepthCell ... CallDepthCell
const SortCallDepthCell Sort = 178

// SortDivInst ... DivInst
const SortDivInst Sort = 179

// SortGasCell ... GasCell
const SortGasCell Sort = 180

// SortIDCell ... IdCell
const SortIDCell Sort = 181

// SortFunctionBodiesCellOpt ... FunctionBodiesCellOpt
const SortFunctionBodiesCellOpt Sort = 182

// SortByteInst ... ByteInst
const SortByteInst Sort = 183

// SortCallDataCellOpt ... CallDataCellOpt
const SortCallDataCellOpt Sort = 184

// SortLocalMemCellOpt ... LocalMemCellOpt
const SortLocalMemCellOpt Sort = 185

// SortMessageCellMap ... MessageCellMap
const SortMessageCellMap Sort = 186

// SortEndianness ... Endianness
const SortEndianness Sort = 187

// SortNparamsCell ... NparamsCell
const SortNparamsCell Sort = 188

// SortCallerCell ... CallerCell
const SortCallerCell Sort = 189

// SortSelfDestructCell ... SelfDestructCell
const SortSelfDestructCell Sort = 190

// SortSendtoCellOpt ... SendtoCellOpt
const SortSendtoCellOpt Sort = 191

// SortMsgIDCellOpt ... MsgIDCellOpt
const SortMsgIDCellOpt Sort = 192

// SortGasPriceCell ... GasPriceCell
const SortGasPriceCell Sort = 193

// SortSet ... Set
const SortSet Sort = 194

// SortOutputCellOpt ... OutputCellOpt
const SortOutputCellOpt Sort = 195

// SortSubInst ... SubInst
const SortSubInst Sort = 196

// SortScheduleCellOpt ... ScheduleCellOpt
const SortScheduleCellOpt Sort = 197

// SortNonceCell ... NonceCell
const SortNonceCell Sort = 198

// SortSCell ... SCell
const SortSCell Sort = 199

// SortMulModInst ... MulModInst
const SortMulModInst Sort = 200

// SortMInt ... MInt
const SortMInt Sort = 201

// SortCell ... Cell
const SortCell Sort = 202

// SortSubstateCellOpt ... SubstateCellOpt
const SortSubstateCellOpt Sort = 203

// SortAccountCellMap ... AccountCellMap
const SortAccountCellMap Sort = 204

// SortOriginCell ... OriginCell
const SortOriginCell Sort = 205

// SortTxPendingCellOpt ... TxPendingCellOpt
const SortTxPendingCellOpt Sort = 206

// SortBswapInst ... BswapInst
const SortBswapInst Sort = 207

// SortTxGasPriceCell ... TxGasPriceCell
const SortTxGasPriceCell Sort = 208

// SortSLoadInst ... SLoadInst
const SortSLoadInst Sort = 209

// SortInstruction ... Instruction
const SortInstruction Sort = 210

// SortContractsCellOpt ... ContractsCellOpt
const SortContractsCellOpt Sort = 211

// SortCurrentContractCellFragment ... CurrentContractCellFragment
const SortCurrentContractCellFragment Sort = 212

// SortTxPendingCell ... TxPendingCell
const SortTxPendingCell Sort = 213

// SortCallStackCell ... CallStackCell
const SortCallStackCell Sort = 214

// SortMessagesCellOpt ... MessagesCellOpt
const SortMessagesCellOpt Sort = 215

// SortKResult ... KResult
const SortKResult Sort = 216

// SortCallDataCell ... CallDataCell
const SortCallDataCell Sort = 217

// SortWordStack ... WordStack
const SortWordStack Sort = 218

// SortRefundCellOpt ... RefundCellOpt
const SortRefundCellOpt Sort = 219

// SortFuncIDsCellOpt ... FuncIdsCellOpt
const SortFuncIDsCellOpt Sort = 220

// SortIELESimulation ... IELESimulation
const SortIELESimulation Sort = 221

// SortG2Point ... G2Point
const SortG2Point Sort = 222

// SortContractsCell ... ContractsCell
const SortContractsCell Sort = 223

// SortFunctionNameCell ... FunctionNameCell
const SortFunctionNameCell Sort = 224

// SortMessagesCellFragment ... MessagesCellFragment
const SortMessagesCellFragment Sort = 225

// SortLabeledBlocks ... LabeledBlocks
const SortLabeledBlocks Sort = 226

// SortIeleCell ... IeleCell
const SortIeleCell Sort = 227

// SortType ... Type
const SortType Sort = 228

// SortGlobalName ... GlobalName
const SortGlobalName Sort = 229

// SortInt ... Int
const SortInt Sort = 230

// SortCurrentMemoryCell ... CurrentMemoryCell
const SortCurrentMemoryCell Sort = 231

// SortInternalOp ... InternalOp
const SortInternalOp Sort = 232

// SortMessagesCell ... MessagesCell
const SortMessagesCell Sort = 233

// SortFuncCell ... FuncCell
const SortFuncCell Sort = 234

// SortWellFormednessScheduleCellOpt ... WellFormednessScheduleCellOpt
const SortWellFormednessScheduleCellOpt Sort = 235

// SortAccountsCellFragment ... AccountsCellFragment
const SortAccountsCellFragment Sort = 236

// SortLValue ... LValue
const SortLValue Sort = 237

// SortGasPriceCellOpt ... GasPriceCellOpt
const SortGasPriceCellOpt Sort = 238

// SortNregsCell ... NregsCell
const SortNregsCell Sort = 239

// SortLocalCallOp ... LocalCallOp
const SortLocalCallOp Sort = 240

// SortContract ... Contract
const SortContract Sort = 241

// SortCallOp ... CallOp
const SortCallOp Sort = 242

// SortAccountCellFragment ... AccountCellFragment
const SortAccountCellFragment Sort = 243

// SortNullOp ... NullOp
const SortNullOp Sort = 244

// SortLabelsCell ... LabelsCell
const SortLabelsCell Sort = 245

// SortAccountsCellOpt ... AccountsCellOpt
const SortAccountsCellOpt Sort = 246

// SortFidCell ... FidCell
const SortFidCell Sort = 247

// SortCallValueCell ... CallValueCell
const SortCallValueCell Sort = 248

// SortPredicate ... Predicate
const SortPredicate Sort = 249

// SortCurrentFunctionCell ... CurrentFunctionCell
const SortCurrentFunctionCell Sort = 250

// SortSignedness ... Signedness
const SortSignedness Sort = 251

// SortSExtInst ... SExtInst
const SortSExtInst Sort = 252

// SortCurrentContractCellOpt ... CurrentContractCellOpt
const SortCurrentContractCellOpt Sort = 253

// SortAcctIDCellOpt ... AcctIDCellOpt
const SortAcctIDCellOpt Sort = 254

// SortStaticCellOpt ... StaticCellOpt
const SortStaticCellOpt Sort = 255

// SortBExp ... BExp
const SortBExp Sort = 256

// SortCallValueCellOpt ... CallValueCellOpt
const SortCallValueCellOpt Sort = 257

// SortAddModInst ... AddModInst
const SortAddModInst Sort = 258

// SortMap ... Map
const SortMap Sort = 259

// SortProgramCellFragment ... ProgramCellFragment
const SortProgramCellFragment Sort = 260

// SortDeclaredContractsCell ... DeclaredContractsCell
const SortDeclaredContractsCell Sort = 261

// SortPreviousGasCellOpt ... PreviousGasCellOpt
const SortPreviousGasCellOpt Sort = 262

// SortFiveOp ... FiveOp
const SortFiveOp Sort = 263

// SortHexConstant ... HexConstant
const SortHexConstant Sort = 264

// SortCreateInst ... CreateInst
const SortCreateInst Sort = 265

// SortIeleCellFragment ... IeleCellFragment
const SortIeleCellFragment Sort = 266

// SortLocalMemCell ... LocalMemCell
const SortLocalMemCell Sort = 267

// SortProgramCellOpt ... ProgramCellOpt
const SortProgramCellOpt Sort = 268

// SortTypesCellOpt ... TypesCellOpt
const SortTypesCellOpt Sort = 269

// SortCondJumpInst ... CondJumpInst
const SortCondJumpInst Sort = 270

// SortPeakMemoryCellOpt ... PeakMemoryCellOpt
const SortPeakMemoryCellOpt Sort = 271

// SortFunctionsCell ... FunctionsCell
const SortFunctionsCell Sort = 272

// SortFromCell ... FromCell
const SortFromCell Sort = 273

// SortContractCodeCell ... ContractCodeCell
const SortContractCodeCell Sort = 274

// SortStream ... Stream
const SortStream Sort = 275

// SortTxGasLimitCellOpt ... TxGasLimitCellOpt
const SortTxGasLimitCellOpt Sort = 276

// SortNetworkCellOpt ... NetworkCellOpt
const SortNetworkCellOpt Sort = 277

// SortIeleName ... IeleName
const SortIeleName Sort = 278

// SortLoadInst ... LoadInst
const SortLoadInst Sort = 279

// SortCallStackCellOpt ... CallStackCellOpt
const SortCallStackCellOpt Sort = 280

// SortSCellOpt ... SCellOpt
const SortSCellOpt Sort = 281

// SortNumberCell ... NumberCell
const SortNumberCell Sort = 282

// SortFunctionBodiesCell ... FunctionBodiesCell
const SortFunctionBodiesCell Sort = 283

// SortFloat ... Float
const SortFloat Sort = 284

// SortWellFormednessScheduleCell ... WellFormednessScheduleCell
const SortWellFormednessScheduleCell Sort = 285

// SortUnOp ... UnOp
const SortUnOp Sort = 286

// SortNonEmptyOperands ... NonEmptyOperands
const SortNonEmptyOperands Sort = 287

// SortTopLevelDefinition ... TopLevelDefinition
const SortTopLevelDefinition Sort = 288

// SortNregsCellOpt ... NregsCellOpt
const SortNregsCellOpt Sort = 289

// SortTxOrderCellOpt ... TxOrderCellOpt
const SortTxOrderCellOpt Sort = 290

// SortXhashLowerID ... #LowerId
const SortXhashLowerID Sort = 291

// SortArray ... Array
const SortArray Sort = 292

// SortIntConstant ... IntConstant
const SortIntConstant Sort = 293

// SortJumpTableCell ... JumpTableCell
const SortJumpTableCell Sort = 294

// SortFidCellOpt ... FidCellOpt
const SortFidCellOpt Sort = 295

// SortAccountCallInst ... AccountCallInst
const SortAccountCallInst Sort = 296

// SortExpModInst ... ExpModInst
const SortExpModInst Sort = 297

// SortFromCellOpt ... FromCellOpt
const SortFromCellOpt Sort = 298

// SortArgsCellOpt ... ArgsCellOpt
const SortArgsCellOpt Sort = 299

// SortTimestampCell ... TimestampCell
const SortTimestampCell Sort = 300

// SortNetworkCellFragment ... NetworkCellFragment
const SortNetworkCellFragment Sort = 301

// SortShiftInst ... ShiftInst
const SortShiftInst Sort = 302

// SortIOError ... IOError
const SortIOError Sort = 303

// SortSelfDestructCellOpt ... SelfDestructCellOpt
const SortSelfDestructCellOpt Sort = 304

// SortFunctionCellMap ... FunctionCellMap
const SortFunctionCellMap Sort = 305

// SortKConfigVar ... KConfigVar
const SortKConfigVar Sort = 306

// SortJSONList ... JSONList
const SortJSONList Sort = 307

// SortSchedule ... Schedule
const SortSchedule Sort = 308

// SortG1Point ... G1Point
const SortG1Point Sort = 309

// SortInstructions ... Instructions
const SortInstructions Sort = 310

// SortAddInst ... AddInst
const SortAddInst Sort = 311

// SortNumericIeleName ... NumericIeleName
const SortNumericIeleName Sort = 312

// SortContractNameCellOpt ... ContractNameCellOpt
const SortContractNameCellOpt Sort = 313

// SortFunctionDefinition ... FunctionDefinition
const SortFunctionDefinition Sort = 314

// SortWellFormednessCellOpt ... WellFormednessCellOpt
const SortWellFormednessCellOpt Sort = 315

// SortCodeCell ... CodeCell
const SortCodeCell Sort = 316

// SortValueCell ... ValueCell
const SortValueCell Sort = 317

// SortFunctionNameCellOpt ... FunctionNameCellOpt
const SortFunctionNameCellOpt Sort = 318

// SortID ... Id
const SortID Sort = 319

// SortNetworkCell ... NetworkCell
const SortNetworkCell Sort = 320

// SortDeclaredContractsCellOpt ... DeclaredContractsCellOpt
const SortDeclaredContractsCellOpt Sort = 321

// SortList ... List
const SortList Sort = 322

// Name ... Sort name
func (s Sort) Name () string {
	switch s {
		case SortCallerCellOpt:
			return "CallerCellOpt"
		case SortSubstateCellFragment:
			return "SubstateCellFragment"
		case SortTxNonceCellOpt:
			return "TxNonceCellOpt"
		case SortModeCellOpt:
			return "ModeCellOpt"
		case SortGeneratedTopCell:
			return "GeneratedTopCell"
		case SortK:
			return "K"
		case SortInts:
			return "Ints"
		case SortQuadOp:
			return "QuadOp"
		case SortStorageCellOpt:
			return "StorageCellOpt"
		case SortScheduleCell:
			return "ScheduleCell"
		case SortCurrentFunctionCellOpt:
			return "CurrentFunctionCellOpt"
		case SortIELECommand:
			return "IELECommand"
		case SortNotInst:
			return "NotInst"
		case SortGasUsedCell:
			return "GasUsedCell"
		case SortArgsCell:
			return "ArgsCell"
		case SortFunctionsCellOpt:
			return "FunctionsCellOpt"
		case SortContractDefinition:
			return "ContractDefinition"
		case SortRevertInst:
			return "RevertInst"
		case SortInstructionsCell:
			return "InstructionsCell"
		case SortCurrentFunctionCellFragment:
			return "CurrentFunctionCellFragment"
		case SortFuncIDsCell:
			return "FuncIdsCell"
		case SortInstructionsCellOpt:
			return "InstructionsCellOpt"
		case SortActiveAccountsCell:
			return "ActiveAccountsCell"
		case SortOrInst:
			return "OrInst"
		case SortCodeCellOpt:
			return "CodeCellOpt"
		case SortSendtoCell:
			return "SendtoCell"
		case SortExportedCell:
			return "ExportedCell"
		case SortTypeCheckingCell:
			return "TypeCheckingCell"
		case SortFunctionCell:
			return "FunctionCell"
		case SortBalanceCellOpt:
			return "BalanceCellOpt"
		case SortAndInst:
			return "AndInst"
		case SortSelfdestructInst:
			return "SelfdestructInst"
		case SortBinOp:
			return "BinOp"
		case SortFunctionSignature:
			return "FunctionSignature"
		case SortTxNonceCell:
			return "TxNonceCell"
		case SortBlockhashCellOpt:
			return "BlockhashCellOpt"
		case SortMsgIDCell:
			return "MsgIDCell"
		case SortRegsCell:
			return "RegsCell"
		case SortFuncIDCellOpt:
			return "FuncIdCellOpt"
		case SortReturnOp:
			return "ReturnOp"
		case SortWellFormednessCellFragment:
			return "WellFormednessCellFragment"
		case SortValueCellOpt:
			return "ValueCellOpt"
		case SortProgramSizeCellOpt:
			return "ProgramSizeCellOpt"
		case SortBool:
			return "Bool"
		case SortOpCode:
			return "OpCode"
		case SortBeneficiaryCell:
			return "BeneficiaryCell"
		case SortNonceCellOpt:
			return "NonceCellOpt"
		case SortKCell:
			return "KCell"
		case SortGasLimitCellOpt:
			return "GasLimitCellOpt"
		case SortFuncIDCell:
			return "FuncIdCell"
		case SortLocalCallsCell:
			return "LocalCallsCell"
		case SortCallFrameCellOpt:
			return "CallFrameCellOpt"
		case SortPseudoInstruction:
			return "PseudoInstruction"
		case SortLocalNames:
			return "LocalNames"
		case SortRefundCell:
			return "RefundCell"
		case SortOperand:
			return "Operand"
		case SortDataCellOpt:
			return "DataCellOpt"
		case SortNonEmptyInts:
			return "NonEmptyInts"
		case SortStringIeleName:
			return "StringIeleName"
		case SortBalanceCell:
			return "BalanceCell"
		case SortStrategy:
			return "Strategy"
		case SortGeneratedTopCellFragment:
			return "GeneratedTopCellFragment"
		case SortKCellOpt:
			return "KCellOpt"
		case SortNumberCellOpt:
			return "NumberCellOpt"
		case SortTypesCell:
			return "TypesCell"
		case SortOriginCellOpt:
			return "OriginCellOpt"
		case SortCurrentInstructionsCellOpt:
			return "CurrentInstructionsCellOpt"
		case SortXhashUpperID:
			return "#UpperId"
		case SortExitCodeCell:
			return "ExitCodeCell"
		case SortMessageCell:
			return "MessageCell"
		case SortGasCellOpt:
			return "GasCellOpt"
		case SortProgramCell:
			return "ProgramCell"
		case SortUnlabeledBlock:
			return "UnlabeledBlock"
		case SortCallFrameCellFragment:
			return "CallFrameCellFragment"
		case SortCreateOp:
			return "CreateOp"
		case SortContractDeclaration:
			return "ContractDeclaration"
		case SortNparamsCellOpt:
			return "NparamsCellOpt"
		case SortMode:
			return "Mode"
		case SortConstant:
			return "Constant"
		case SortBeneficiaryCellOpt:
			return "BeneficiaryCellOpt"
		case SortContractNameCell:
			return "ContractNameCell"
		case SortLocalCallsCellOpt:
			return "LocalCallsCellOpt"
		case SortProgramSizeCell:
			return "ProgramSizeCell"
		case SortLocalCall:
			return "LocalCall"
		case SortLogInst:
			return "LogInst"
		case SortContractCodeCellOpt:
			return "ContractCodeCellOpt"
		case SortLengthPrefix:
			return "LengthPrefix"
		case SortTypeCheckingCellOpt:
			return "TypeCheckingCellOpt"
		case SortFunctionCellFragment:
			return "FunctionCellFragment"
		case SortIeleCellOpt:
			return "IeleCellOpt"
		case SortJSONKey:
			return "JSONKey"
		case SortCallDepthCellOpt:
			return "CallDepthCellOpt"
		case SortDataCell:
			return "DataCell"
		case SortLogDataCell:
			return "LogDataCell"
		case SortMulInst:
			return "MulInst"
		case SortTopLevelDefinitions:
			return "TopLevelDefinitions"
		case SortOutputCell:
			return "OutputCell"
		case SortMessageCellFragment:
			return "MessageCellFragment"
		case SortSubstateStackCellOpt:
			return "SubstateStackCellOpt"
		case SortLocalCallInst:
			return "LocalCallInst"
		case SortSStoreInst:
			return "SStoreInst"
		case SortCurrentMemoryCellOpt:
			return "CurrentMemoryCellOpt"
		case SortGlobalDefinition:
			return "GlobalDefinition"
		case SortSubstateCell:
			return "SubstateCell"
		case SortInterimStatesCell:
			return "InterimStatesCell"
		case SortLabelsCellOpt:
			return "LabelsCellOpt"
		case SortAccountsCell:
			return "AccountsCell"
		case SortAccountCell:
			return "AccountCell"
		case SortBlocks:
			return "Blocks"
		case SortSubstateStackCell:
			return "SubstateStackCell"
		case SortDifficultyCellOpt:
			return "DifficultyCellOpt"
		case SortTwosInst:
			return "TwosInst"
		case SortScheduleConst:
			return "ScheduleConst"
		case SortLogDataCellOpt:
			return "LogDataCellOpt"
		case SortWellFormednessCell:
			return "WellFormednessCell"
		case SortString:
			return "String"
		case SortModeCell:
			return "ModeCell"
		case SortStorageCell:
			return "StorageCell"
		case SortCurrentContractCell:
			return "CurrentContractCell"
		case SortTxGasPriceCellOpt:
			return "TxGasPriceCellOpt"
		case SortCurrentInstructionsCell:
			return "CurrentInstructionsCell"
		case SortExitCodeCellOpt:
			return "ExitCodeCellOpt"
		case SortTxOrderCell:
			return "TxOrderCell"
		case SortStaticCell:
			return "StaticCell"
		case SortPeakMemoryCell:
			return "PeakMemoryCell"
		case SortAcctIDCell:
			return "AcctIDCell"
		case SortFunctionsCellFragment:
			return "FunctionsCellFragment"
		case SortLValues:
			return "LValues"
		case SortAccount:
			return "Account"
		case SortXhashRuleTag:
			return "#RuleTag"
		case SortTernOp:
			return "TernOp"
		case SortCallAddressInst:
			return "CallAddressInst"
		case SortRegsCellOpt:
			return "RegsCellOpt"
		case SortCallFrameCell:
			return "CallFrameCell"
		case SortBytes:
			return "Bytes"
		case SortGasUsedCellOpt:
			return "GasUsedCellOpt"
		case SortFunctionParameters:
			return "FunctionParameters"
		case SortTypes:
			return "Types"
		case SortActiveAccountsCellOpt:
			return "ActiveAccountsCellOpt"
		case SortStringBuffer:
			return "StringBuffer"
		case SortExportedCellOpt:
			return "ExportedCellOpt"
		case SortModInst:
			return "ModInst"
		case SortFuncLabelsCellOpt:
			return "FuncLabelsCellOpt"
		case SortScheduleFlag:
			return "ScheduleFlag"
		case SortCheckGasCellOpt:
			return "CheckGasCellOpt"
		case SortSHA3Inst:
			return "SHA3Inst"
		case SortStoreInst:
			return "StoreInst"
		case SortOperands:
			return "Operands"
		case SortSubstateLogEntry:
			return "SubstateLogEntry"
		case SortFuncCellOpt:
			return "FuncCellOpt"
		case SortJumpTableCellOpt:
			return "JumpTableCellOpt"
		case SortLabeledBlock:
			return "LabeledBlock"
		case SortFuncLabelsCell:
			return "FuncLabelsCell"
		case SortAssignInst:
			return "AssignInst"
		case SortLocalName:
			return "LocalName"
		case SortExpInst:
			return "ExpInst"
		case SortIsZeroInst:
			return "IsZeroInst"
		case SortInterimStatesCellOpt:
			return "InterimStatesCellOpt"
		case SortLengthPrefixType:
			return "LengthPrefixType"
		case SortReturnType:
			return "ReturnType"
		case SortAccounts:
			return "Accounts"
		case SortJumpInst:
			return "JumpInst"
		case SortGasLimitCell:
			return "GasLimitCell"
		case SortBlockhashCell:
			return "BlockhashCell"
		case SortReturnInst:
			return "ReturnInst"
		case SortDifficultyCell:
			return "DifficultyCell"
		case SortCmpInst:
			return "CmpInst"
		case SortIDCellOpt:
			return "IdCellOpt"
		case SortException:
			return "Exception"
		case SortCheckGasCell:
			return "CheckGasCell"
		case SortCopyCreateOp:
			return "CopyCreateOp"
		case SortPreviousGasCell:
			return "PreviousGasCell"
		case SortXorInst:
			return "XorInst"
		case SortPrecompiledOp:
			return "PrecompiledOp"
		case SortTxGasLimitCell:
			return "TxGasLimitCell"
		case SortTimestampCellOpt:
			return "TimestampCellOpt"
		case SortKItem:
			return "KItem"
		case SortJSON:
			return "JSON"
		case SortCallDepthCell:
			return "CallDepthCell"
		case SortDivInst:
			return "DivInst"
		case SortGasCell:
			return "GasCell"
		case SortIDCell:
			return "IdCell"
		case SortFunctionBodiesCellOpt:
			return "FunctionBodiesCellOpt"
		case SortByteInst:
			return "ByteInst"
		case SortCallDataCellOpt:
			return "CallDataCellOpt"
		case SortLocalMemCellOpt:
			return "LocalMemCellOpt"
		case SortMessageCellMap:
			return "MessageCellMap"
		case SortEndianness:
			return "Endianness"
		case SortNparamsCell:
			return "NparamsCell"
		case SortCallerCell:
			return "CallerCell"
		case SortSelfDestructCell:
			return "SelfDestructCell"
		case SortSendtoCellOpt:
			return "SendtoCellOpt"
		case SortMsgIDCellOpt:
			return "MsgIDCellOpt"
		case SortGasPriceCell:
			return "GasPriceCell"
		case SortSet:
			return "Set"
		case SortOutputCellOpt:
			return "OutputCellOpt"
		case SortSubInst:
			return "SubInst"
		case SortScheduleCellOpt:
			return "ScheduleCellOpt"
		case SortNonceCell:
			return "NonceCell"
		case SortSCell:
			return "SCell"
		case SortMulModInst:
			return "MulModInst"
		case SortMInt:
			return "MInt"
		case SortCell:
			return "Cell"
		case SortSubstateCellOpt:
			return "SubstateCellOpt"
		case SortAccountCellMap:
			return "AccountCellMap"
		case SortOriginCell:
			return "OriginCell"
		case SortTxPendingCellOpt:
			return "TxPendingCellOpt"
		case SortBswapInst:
			return "BswapInst"
		case SortTxGasPriceCell:
			return "TxGasPriceCell"
		case SortSLoadInst:
			return "SLoadInst"
		case SortInstruction:
			return "Instruction"
		case SortContractsCellOpt:
			return "ContractsCellOpt"
		case SortCurrentContractCellFragment:
			return "CurrentContractCellFragment"
		case SortTxPendingCell:
			return "TxPendingCell"
		case SortCallStackCell:
			return "CallStackCell"
		case SortMessagesCellOpt:
			return "MessagesCellOpt"
		case SortKResult:
			return "KResult"
		case SortCallDataCell:
			return "CallDataCell"
		case SortWordStack:
			return "WordStack"
		case SortRefundCellOpt:
			return "RefundCellOpt"
		case SortFuncIDsCellOpt:
			return "FuncIdsCellOpt"
		case SortIELESimulation:
			return "IELESimulation"
		case SortG2Point:
			return "G2Point"
		case SortContractsCell:
			return "ContractsCell"
		case SortFunctionNameCell:
			return "FunctionNameCell"
		case SortMessagesCellFragment:
			return "MessagesCellFragment"
		case SortLabeledBlocks:
			return "LabeledBlocks"
		case SortIeleCell:
			return "IeleCell"
		case SortType:
			return "Type"
		case SortGlobalName:
			return "GlobalName"
		case SortInt:
			return "Int"
		case SortCurrentMemoryCell:
			return "CurrentMemoryCell"
		case SortInternalOp:
			return "InternalOp"
		case SortMessagesCell:
			return "MessagesCell"
		case SortFuncCell:
			return "FuncCell"
		case SortWellFormednessScheduleCellOpt:
			return "WellFormednessScheduleCellOpt"
		case SortAccountsCellFragment:
			return "AccountsCellFragment"
		case SortLValue:
			return "LValue"
		case SortGasPriceCellOpt:
			return "GasPriceCellOpt"
		case SortNregsCell:
			return "NregsCell"
		case SortLocalCallOp:
			return "LocalCallOp"
		case SortContract:
			return "Contract"
		case SortCallOp:
			return "CallOp"
		case SortAccountCellFragment:
			return "AccountCellFragment"
		case SortNullOp:
			return "NullOp"
		case SortLabelsCell:
			return "LabelsCell"
		case SortAccountsCellOpt:
			return "AccountsCellOpt"
		case SortFidCell:
			return "FidCell"
		case SortCallValueCell:
			return "CallValueCell"
		case SortPredicate:
			return "Predicate"
		case SortCurrentFunctionCell:
			return "CurrentFunctionCell"
		case SortSignedness:
			return "Signedness"
		case SortSExtInst:
			return "SExtInst"
		case SortCurrentContractCellOpt:
			return "CurrentContractCellOpt"
		case SortAcctIDCellOpt:
			return "AcctIDCellOpt"
		case SortStaticCellOpt:
			return "StaticCellOpt"
		case SortBExp:
			return "BExp"
		case SortCallValueCellOpt:
			return "CallValueCellOpt"
		case SortAddModInst:
			return "AddModInst"
		case SortMap:
			return "Map"
		case SortProgramCellFragment:
			return "ProgramCellFragment"
		case SortDeclaredContractsCell:
			return "DeclaredContractsCell"
		case SortPreviousGasCellOpt:
			return "PreviousGasCellOpt"
		case SortFiveOp:
			return "FiveOp"
		case SortHexConstant:
			return "HexConstant"
		case SortCreateInst:
			return "CreateInst"
		case SortIeleCellFragment:
			return "IeleCellFragment"
		case SortLocalMemCell:
			return "LocalMemCell"
		case SortProgramCellOpt:
			return "ProgramCellOpt"
		case SortTypesCellOpt:
			return "TypesCellOpt"
		case SortCondJumpInst:
			return "CondJumpInst"
		case SortPeakMemoryCellOpt:
			return "PeakMemoryCellOpt"
		case SortFunctionsCell:
			return "FunctionsCell"
		case SortFromCell:
			return "FromCell"
		case SortContractCodeCell:
			return "ContractCodeCell"
		case SortStream:
			return "Stream"
		case SortTxGasLimitCellOpt:
			return "TxGasLimitCellOpt"
		case SortNetworkCellOpt:
			return "NetworkCellOpt"
		case SortIeleName:
			return "IeleName"
		case SortLoadInst:
			return "LoadInst"
		case SortCallStackCellOpt:
			return "CallStackCellOpt"
		case SortSCellOpt:
			return "SCellOpt"
		case SortNumberCell:
			return "NumberCell"
		case SortFunctionBodiesCell:
			return "FunctionBodiesCell"
		case SortFloat:
			return "Float"
		case SortWellFormednessScheduleCell:
			return "WellFormednessScheduleCell"
		case SortUnOp:
			return "UnOp"
		case SortNonEmptyOperands:
			return "NonEmptyOperands"
		case SortTopLevelDefinition:
			return "TopLevelDefinition"
		case SortNregsCellOpt:
			return "NregsCellOpt"
		case SortTxOrderCellOpt:
			return "TxOrderCellOpt"
		case SortXhashLowerID:
			return "#LowerId"
		case SortArray:
			return "Array"
		case SortIntConstant:
			return "IntConstant"
		case SortJumpTableCell:
			return "JumpTableCell"
		case SortFidCellOpt:
			return "FidCellOpt"
		case SortAccountCallInst:
			return "AccountCallInst"
		case SortExpModInst:
			return "ExpModInst"
		case SortFromCellOpt:
			return "FromCellOpt"
		case SortArgsCellOpt:
			return "ArgsCellOpt"
		case SortTimestampCell:
			return "TimestampCell"
		case SortNetworkCellFragment:
			return "NetworkCellFragment"
		case SortShiftInst:
			return "ShiftInst"
		case SortIOError:
			return "IOError"
		case SortSelfDestructCellOpt:
			return "SelfDestructCellOpt"
		case SortFunctionCellMap:
			return "FunctionCellMap"
		case SortKConfigVar:
			return "KConfigVar"
		case SortJSONList:
			return "JSONList"
		case SortSchedule:
			return "Schedule"
		case SortG1Point:
			return "G1Point"
		case SortInstructions:
			return "Instructions"
		case SortAddInst:
			return "AddInst"
		case SortNumericIeleName:
			return "NumericIeleName"
		case SortContractNameCellOpt:
			return "ContractNameCellOpt"
		case SortFunctionDefinition:
			return "FunctionDefinition"
		case SortWellFormednessCellOpt:
			return "WellFormednessCellOpt"
		case SortCodeCell:
			return "CodeCell"
		case SortValueCell:
			return "ValueCell"
		case SortFunctionNameCellOpt:
			return "FunctionNameCellOpt"
		case SortID:
			return "Id"
		case SortNetworkCell:
			return "NetworkCell"
		case SortDeclaredContractsCellOpt:
			return "DeclaredContractsCellOpt"
		case SortList:
			return "List"
		default:
			panic("Unexpected Sort.")
	}
}

// ParseSort ... Yields the sort with the given name
func ParseSort (name string) Sort {
	switch name {
		case "CallerCellOpt":
			return SortCallerCellOpt
		case "SubstateCellFragment":
			return SortSubstateCellFragment
		case "TxNonceCellOpt":
			return SortTxNonceCellOpt
		case "ModeCellOpt":
			return SortModeCellOpt
		case "GeneratedTopCell":
			return SortGeneratedTopCell
		case "K":
			return SortK
		case "Ints":
			return SortInts
		case "QuadOp":
			return SortQuadOp
		case "StorageCellOpt":
			return SortStorageCellOpt
		case "ScheduleCell":
			return SortScheduleCell
		case "CurrentFunctionCellOpt":
			return SortCurrentFunctionCellOpt
		case "IELECommand":
			return SortIELECommand
		case "NotInst":
			return SortNotInst
		case "GasUsedCell":
			return SortGasUsedCell
		case "ArgsCell":
			return SortArgsCell
		case "FunctionsCellOpt":
			return SortFunctionsCellOpt
		case "ContractDefinition":
			return SortContractDefinition
		case "RevertInst":
			return SortRevertInst
		case "InstructionsCell":
			return SortInstructionsCell
		case "CurrentFunctionCellFragment":
			return SortCurrentFunctionCellFragment
		case "FuncIdsCell":
			return SortFuncIDsCell
		case "InstructionsCellOpt":
			return SortInstructionsCellOpt
		case "ActiveAccountsCell":
			return SortActiveAccountsCell
		case "OrInst":
			return SortOrInst
		case "CodeCellOpt":
			return SortCodeCellOpt
		case "SendtoCell":
			return SortSendtoCell
		case "ExportedCell":
			return SortExportedCell
		case "TypeCheckingCell":
			return SortTypeCheckingCell
		case "FunctionCell":
			return SortFunctionCell
		case "BalanceCellOpt":
			return SortBalanceCellOpt
		case "AndInst":
			return SortAndInst
		case "SelfdestructInst":
			return SortSelfdestructInst
		case "BinOp":
			return SortBinOp
		case "FunctionSignature":
			return SortFunctionSignature
		case "TxNonceCell":
			return SortTxNonceCell
		case "BlockhashCellOpt":
			return SortBlockhashCellOpt
		case "MsgIDCell":
			return SortMsgIDCell
		case "RegsCell":
			return SortRegsCell
		case "FuncIdCellOpt":
			return SortFuncIDCellOpt
		case "ReturnOp":
			return SortReturnOp
		case "WellFormednessCellFragment":
			return SortWellFormednessCellFragment
		case "ValueCellOpt":
			return SortValueCellOpt
		case "ProgramSizeCellOpt":
			return SortProgramSizeCellOpt
		case "Bool":
			return SortBool
		case "OpCode":
			return SortOpCode
		case "BeneficiaryCell":
			return SortBeneficiaryCell
		case "NonceCellOpt":
			return SortNonceCellOpt
		case "KCell":
			return SortKCell
		case "GasLimitCellOpt":
			return SortGasLimitCellOpt
		case "FuncIdCell":
			return SortFuncIDCell
		case "LocalCallsCell":
			return SortLocalCallsCell
		case "CallFrameCellOpt":
			return SortCallFrameCellOpt
		case "PseudoInstruction":
			return SortPseudoInstruction
		case "LocalNames":
			return SortLocalNames
		case "RefundCell":
			return SortRefundCell
		case "Operand":
			return SortOperand
		case "DataCellOpt":
			return SortDataCellOpt
		case "NonEmptyInts":
			return SortNonEmptyInts
		case "StringIeleName":
			return SortStringIeleName
		case "BalanceCell":
			return SortBalanceCell
		case "Strategy":
			return SortStrategy
		case "GeneratedTopCellFragment":
			return SortGeneratedTopCellFragment
		case "KCellOpt":
			return SortKCellOpt
		case "NumberCellOpt":
			return SortNumberCellOpt
		case "TypesCell":
			return SortTypesCell
		case "OriginCellOpt":
			return SortOriginCellOpt
		case "CurrentInstructionsCellOpt":
			return SortCurrentInstructionsCellOpt
		case "#UpperId":
			return SortXhashUpperID
		case "ExitCodeCell":
			return SortExitCodeCell
		case "MessageCell":
			return SortMessageCell
		case "GasCellOpt":
			return SortGasCellOpt
		case "ProgramCell":
			return SortProgramCell
		case "UnlabeledBlock":
			return SortUnlabeledBlock
		case "CallFrameCellFragment":
			return SortCallFrameCellFragment
		case "CreateOp":
			return SortCreateOp
		case "ContractDeclaration":
			return SortContractDeclaration
		case "NparamsCellOpt":
			return SortNparamsCellOpt
		case "Mode":
			return SortMode
		case "Constant":
			return SortConstant
		case "BeneficiaryCellOpt":
			return SortBeneficiaryCellOpt
		case "ContractNameCell":
			return SortContractNameCell
		case "LocalCallsCellOpt":
			return SortLocalCallsCellOpt
		case "ProgramSizeCell":
			return SortProgramSizeCell
		case "LocalCall":
			return SortLocalCall
		case "LogInst":
			return SortLogInst
		case "ContractCodeCellOpt":
			return SortContractCodeCellOpt
		case "LengthPrefix":
			return SortLengthPrefix
		case "TypeCheckingCellOpt":
			return SortTypeCheckingCellOpt
		case "FunctionCellFragment":
			return SortFunctionCellFragment
		case "IeleCellOpt":
			return SortIeleCellOpt
		case "JSONKey":
			return SortJSONKey
		case "CallDepthCellOpt":
			return SortCallDepthCellOpt
		case "DataCell":
			return SortDataCell
		case "LogDataCell":
			return SortLogDataCell
		case "MulInst":
			return SortMulInst
		case "TopLevelDefinitions":
			return SortTopLevelDefinitions
		case "OutputCell":
			return SortOutputCell
		case "MessageCellFragment":
			return SortMessageCellFragment
		case "SubstateStackCellOpt":
			return SortSubstateStackCellOpt
		case "LocalCallInst":
			return SortLocalCallInst
		case "SStoreInst":
			return SortSStoreInst
		case "CurrentMemoryCellOpt":
			return SortCurrentMemoryCellOpt
		case "GlobalDefinition":
			return SortGlobalDefinition
		case "SubstateCell":
			return SortSubstateCell
		case "InterimStatesCell":
			return SortInterimStatesCell
		case "LabelsCellOpt":
			return SortLabelsCellOpt
		case "AccountsCell":
			return SortAccountsCell
		case "AccountCell":
			return SortAccountCell
		case "Blocks":
			return SortBlocks
		case "SubstateStackCell":
			return SortSubstateStackCell
		case "DifficultyCellOpt":
			return SortDifficultyCellOpt
		case "TwosInst":
			return SortTwosInst
		case "ScheduleConst":
			return SortScheduleConst
		case "LogDataCellOpt":
			return SortLogDataCellOpt
		case "WellFormednessCell":
			return SortWellFormednessCell
		case "String":
			return SortString
		case "ModeCell":
			return SortModeCell
		case "StorageCell":
			return SortStorageCell
		case "CurrentContractCell":
			return SortCurrentContractCell
		case "TxGasPriceCellOpt":
			return SortTxGasPriceCellOpt
		case "CurrentInstructionsCell":
			return SortCurrentInstructionsCell
		case "ExitCodeCellOpt":
			return SortExitCodeCellOpt
		case "TxOrderCell":
			return SortTxOrderCell
		case "StaticCell":
			return SortStaticCell
		case "PeakMemoryCell":
			return SortPeakMemoryCell
		case "AcctIDCell":
			return SortAcctIDCell
		case "FunctionsCellFragment":
			return SortFunctionsCellFragment
		case "LValues":
			return SortLValues
		case "Account":
			return SortAccount
		case "#RuleTag":
			return SortXhashRuleTag
		case "TernOp":
			return SortTernOp
		case "CallAddressInst":
			return SortCallAddressInst
		case "RegsCellOpt":
			return SortRegsCellOpt
		case "CallFrameCell":
			return SortCallFrameCell
		case "Bytes":
			return SortBytes
		case "GasUsedCellOpt":
			return SortGasUsedCellOpt
		case "FunctionParameters":
			return SortFunctionParameters
		case "Types":
			return SortTypes
		case "ActiveAccountsCellOpt":
			return SortActiveAccountsCellOpt
		case "StringBuffer":
			return SortStringBuffer
		case "ExportedCellOpt":
			return SortExportedCellOpt
		case "ModInst":
			return SortModInst
		case "FuncLabelsCellOpt":
			return SortFuncLabelsCellOpt
		case "ScheduleFlag":
			return SortScheduleFlag
		case "CheckGasCellOpt":
			return SortCheckGasCellOpt
		case "SHA3Inst":
			return SortSHA3Inst
		case "StoreInst":
			return SortStoreInst
		case "Operands":
			return SortOperands
		case "SubstateLogEntry":
			return SortSubstateLogEntry
		case "FuncCellOpt":
			return SortFuncCellOpt
		case "JumpTableCellOpt":
			return SortJumpTableCellOpt
		case "LabeledBlock":
			return SortLabeledBlock
		case "FuncLabelsCell":
			return SortFuncLabelsCell
		case "AssignInst":
			return SortAssignInst
		case "LocalName":
			return SortLocalName
		case "ExpInst":
			return SortExpInst
		case "IsZeroInst":
			return SortIsZeroInst
		case "InterimStatesCellOpt":
			return SortInterimStatesCellOpt
		case "LengthPrefixType":
			return SortLengthPrefixType
		case "ReturnType":
			return SortReturnType
		case "Accounts":
			return SortAccounts
		case "JumpInst":
			return SortJumpInst
		case "GasLimitCell":
			return SortGasLimitCell
		case "BlockhashCell":
			return SortBlockhashCell
		case "ReturnInst":
			return SortReturnInst
		case "DifficultyCell":
			return SortDifficultyCell
		case "CmpInst":
			return SortCmpInst
		case "IdCellOpt":
			return SortIDCellOpt
		case "Exception":
			return SortException
		case "CheckGasCell":
			return SortCheckGasCell
		case "CopyCreateOp":
			return SortCopyCreateOp
		case "PreviousGasCell":
			return SortPreviousGasCell
		case "XorInst":
			return SortXorInst
		case "PrecompiledOp":
			return SortPrecompiledOp
		case "TxGasLimitCell":
			return SortTxGasLimitCell
		case "TimestampCellOpt":
			return SortTimestampCellOpt
		case "KItem":
			return SortKItem
		case "JSON":
			return SortJSON
		case "CallDepthCell":
			return SortCallDepthCell
		case "DivInst":
			return SortDivInst
		case "GasCell":
			return SortGasCell
		case "IdCell":
			return SortIDCell
		case "FunctionBodiesCellOpt":
			return SortFunctionBodiesCellOpt
		case "ByteInst":
			return SortByteInst
		case "CallDataCellOpt":
			return SortCallDataCellOpt
		case "LocalMemCellOpt":
			return SortLocalMemCellOpt
		case "MessageCellMap":
			return SortMessageCellMap
		case "Endianness":
			return SortEndianness
		case "NparamsCell":
			return SortNparamsCell
		case "CallerCell":
			return SortCallerCell
		case "SelfDestructCell":
			return SortSelfDestructCell
		case "SendtoCellOpt":
			return SortSendtoCellOpt
		case "MsgIDCellOpt":
			return SortMsgIDCellOpt
		case "GasPriceCell":
			return SortGasPriceCell
		case "Set":
			return SortSet
		case "OutputCellOpt":
			return SortOutputCellOpt
		case "SubInst":
			return SortSubInst
		case "ScheduleCellOpt":
			return SortScheduleCellOpt
		case "NonceCell":
			return SortNonceCell
		case "SCell":
			return SortSCell
		case "MulModInst":
			return SortMulModInst
		case "MInt":
			return SortMInt
		case "Cell":
			return SortCell
		case "SubstateCellOpt":
			return SortSubstateCellOpt
		case "AccountCellMap":
			return SortAccountCellMap
		case "OriginCell":
			return SortOriginCell
		case "TxPendingCellOpt":
			return SortTxPendingCellOpt
		case "BswapInst":
			return SortBswapInst
		case "TxGasPriceCell":
			return SortTxGasPriceCell
		case "SLoadInst":
			return SortSLoadInst
		case "Instruction":
			return SortInstruction
		case "ContractsCellOpt":
			return SortContractsCellOpt
		case "CurrentContractCellFragment":
			return SortCurrentContractCellFragment
		case "TxPendingCell":
			return SortTxPendingCell
		case "CallStackCell":
			return SortCallStackCell
		case "MessagesCellOpt":
			return SortMessagesCellOpt
		case "KResult":
			return SortKResult
		case "CallDataCell":
			return SortCallDataCell
		case "WordStack":
			return SortWordStack
		case "RefundCellOpt":
			return SortRefundCellOpt
		case "FuncIdsCellOpt":
			return SortFuncIDsCellOpt
		case "IELESimulation":
			return SortIELESimulation
		case "G2Point":
			return SortG2Point
		case "ContractsCell":
			return SortContractsCell
		case "FunctionNameCell":
			return SortFunctionNameCell
		case "MessagesCellFragment":
			return SortMessagesCellFragment
		case "LabeledBlocks":
			return SortLabeledBlocks
		case "IeleCell":
			return SortIeleCell
		case "Type":
			return SortType
		case "GlobalName":
			return SortGlobalName
		case "Int":
			return SortInt
		case "CurrentMemoryCell":
			return SortCurrentMemoryCell
		case "InternalOp":
			return SortInternalOp
		case "MessagesCell":
			return SortMessagesCell
		case "FuncCell":
			return SortFuncCell
		case "WellFormednessScheduleCellOpt":
			return SortWellFormednessScheduleCellOpt
		case "AccountsCellFragment":
			return SortAccountsCellFragment
		case "LValue":
			return SortLValue
		case "GasPriceCellOpt":
			return SortGasPriceCellOpt
		case "NregsCell":
			return SortNregsCell
		case "LocalCallOp":
			return SortLocalCallOp
		case "Contract":
			return SortContract
		case "CallOp":
			return SortCallOp
		case "AccountCellFragment":
			return SortAccountCellFragment
		case "NullOp":
			return SortNullOp
		case "LabelsCell":
			return SortLabelsCell
		case "AccountsCellOpt":
			return SortAccountsCellOpt
		case "FidCell":
			return SortFidCell
		case "CallValueCell":
			return SortCallValueCell
		case "Predicate":
			return SortPredicate
		case "CurrentFunctionCell":
			return SortCurrentFunctionCell
		case "Signedness":
			return SortSignedness
		case "SExtInst":
			return SortSExtInst
		case "CurrentContractCellOpt":
			return SortCurrentContractCellOpt
		case "AcctIDCellOpt":
			return SortAcctIDCellOpt
		case "StaticCellOpt":
			return SortStaticCellOpt
		case "BExp":
			return SortBExp
		case "CallValueCellOpt":
			return SortCallValueCellOpt
		case "AddModInst":
			return SortAddModInst
		case "Map":
			return SortMap
		case "ProgramCellFragment":
			return SortProgramCellFragment
		case "DeclaredContractsCell":
			return SortDeclaredContractsCell
		case "PreviousGasCellOpt":
			return SortPreviousGasCellOpt
		case "FiveOp":
			return SortFiveOp
		case "HexConstant":
			return SortHexConstant
		case "CreateInst":
			return SortCreateInst
		case "IeleCellFragment":
			return SortIeleCellFragment
		case "LocalMemCell":
			return SortLocalMemCell
		case "ProgramCellOpt":
			return SortProgramCellOpt
		case "TypesCellOpt":
			return SortTypesCellOpt
		case "CondJumpInst":
			return SortCondJumpInst
		case "PeakMemoryCellOpt":
			return SortPeakMemoryCellOpt
		case "FunctionsCell":
			return SortFunctionsCell
		case "FromCell":
			return SortFromCell
		case "ContractCodeCell":
			return SortContractCodeCell
		case "Stream":
			return SortStream
		case "TxGasLimitCellOpt":
			return SortTxGasLimitCellOpt
		case "NetworkCellOpt":
			return SortNetworkCellOpt
		case "IeleName":
			return SortIeleName
		case "LoadInst":
			return SortLoadInst
		case "CallStackCellOpt":
			return SortCallStackCellOpt
		case "SCellOpt":
			return SortSCellOpt
		case "NumberCell":
			return SortNumberCell
		case "FunctionBodiesCell":
			return SortFunctionBodiesCell
		case "Float":
			return SortFloat
		case "WellFormednessScheduleCell":
			return SortWellFormednessScheduleCell
		case "UnOp":
			return SortUnOp
		case "NonEmptyOperands":
			return SortNonEmptyOperands
		case "TopLevelDefinition":
			return SortTopLevelDefinition
		case "NregsCellOpt":
			return SortNregsCellOpt
		case "TxOrderCellOpt":
			return SortTxOrderCellOpt
		case "#LowerId":
			return SortXhashLowerID
		case "Array":
			return SortArray
		case "IntConstant":
			return SortIntConstant
		case "JumpTableCell":
			return SortJumpTableCell
		case "FidCellOpt":
			return SortFidCellOpt
		case "AccountCallInst":
			return SortAccountCallInst
		case "ExpModInst":
			return SortExpModInst
		case "FromCellOpt":
			return SortFromCellOpt
		case "ArgsCellOpt":
			return SortArgsCellOpt
		case "TimestampCell":
			return SortTimestampCell
		case "NetworkCellFragment":
			return SortNetworkCellFragment
		case "ShiftInst":
			return SortShiftInst
		case "IOError":
			return SortIOError
		case "SelfDestructCellOpt":
			return SortSelfDestructCellOpt
		case "FunctionCellMap":
			return SortFunctionCellMap
		case "KConfigVar":
			return SortKConfigVar
		case "JSONList":
			return SortJSONList
		case "Schedule":
			return SortSchedule
		case "G1Point":
			return SortG1Point
		case "Instructions":
			return SortInstructions
		case "AddInst":
			return SortAddInst
		case "NumericIeleName":
			return SortNumericIeleName
		case "ContractNameCellOpt":
			return SortContractNameCellOpt
		case "FunctionDefinition":
			return SortFunctionDefinition
		case "WellFormednessCellOpt":
			return SortWellFormednessCellOpt
		case "CodeCell":
			return SortCodeCell
		case "ValueCell":
			return SortValueCell
		case "FunctionNameCellOpt":
			return SortFunctionNameCellOpt
		case "Id":
			return SortID
		case "NetworkCell":
			return SortNetworkCell
		case "DeclaredContractsCellOpt":
			return SortDeclaredContractsCellOpt
		case "List":
			return SortList
		default:
			panic("Parsing Sort failed. Unexpected Sort name:" + name)
	}
}

