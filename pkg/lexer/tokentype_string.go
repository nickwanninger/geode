// Code generated by "stringer -type=TokenType TokenType.go"; DO NOT EDIT.

package lexer

import "strconv"

const _TokenType_name = "TokErrorTokNoEmitTokWhitespaceTokCharTokStringTokNumberTokBoolTokDotTokElipsisTokOperTokNamespaceAccessTokOperatorStartTokStarTokPlusTokMinusTokDivTokExpTokLTTokLTETokGTTokGTETokOperatorEndTokSemiColonTokDefereferenceTokReferenceTokAssignmentTokEqualityTokRightParenTokLeftParenTokRightCurlyTokLeftCurlyTokRightBraceTokLeftBraceTokRightArrowTokLeftArrowTokInfoTokCompoundAssignmentTokQuestionMarkTokForTokWhileTokIfTokElseTokReturnTokFuncDefnTokClassDefnTokNamespaceTokLetTokAsTokNilTokDependencyTokTypeTokCommaTokIdentTokSymbolTokComment"

var _TokenType_index = [...]uint16{0, 8, 17, 30, 37, 46, 55, 62, 68, 78, 85, 103, 119, 126, 133, 141, 147, 153, 158, 164, 169, 175, 189, 201, 217, 229, 242, 253, 266, 278, 291, 303, 316, 328, 341, 353, 360, 381, 396, 402, 410, 415, 422, 431, 442, 454, 466, 472, 477, 483, 496, 503, 511, 519, 528, 538}

func (i TokenType) String() string {
	if i < 0 || i >= TokenType(len(_TokenType_index)-1) {
		return "TokenType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TokenType_name[_TokenType_index[i]:_TokenType_index[i+1]]
}
