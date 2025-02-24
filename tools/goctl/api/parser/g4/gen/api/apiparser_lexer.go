package api

import (
	"fmt"
	"unicode"

	"github.com/xiaoshouchen/antlr"
)

// Suppress unused import error
var (
	_ = fmt.Printf
	_ = unicode.IsLetter
)

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 26, 276,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 3, 2, 3, 2,
	3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3,
	13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3,
	17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 6, 19, 142, 10,
	19, 13, 19, 14, 19, 143, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 7, 20,
	152, 10, 20, 12, 20, 14, 20, 155, 11, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3,
	20, 3, 21, 3, 21, 3, 21, 3, 21, 7, 21, 166, 10, 21, 12, 21, 14, 21, 169,
	11, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 7, 22, 176, 10, 22, 12, 22,
	14, 22, 179, 11, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 6, 23, 186, 10,
	23, 13, 23, 14, 23, 187, 3, 23, 3, 23, 3, 24, 3, 24, 7, 24, 194, 10, 24,
	12, 24, 14, 24, 197, 11, 24, 3, 24, 3, 24, 7, 24, 201, 10, 24, 12, 24,
	14, 24, 204, 11, 24, 5, 24, 206, 10, 24, 3, 25, 3, 25, 7, 25, 210, 10,
	25, 12, 25, 14, 25, 213, 11, 25, 3, 26, 3, 26, 5, 26, 217, 10, 26, 3, 26,
	3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 5, 27, 225, 10, 27, 3, 27, 5, 27, 228,
	10, 27, 3, 27, 3, 27, 3, 27, 6, 27, 233, 10, 27, 13, 27, 14, 27, 234, 3,
	27, 3, 27, 3, 27, 3, 27, 3, 27, 5, 27, 242, 10, 27, 3, 28, 3, 28, 3, 28,
	7, 28, 247, 10, 28, 12, 28, 14, 28, 250, 11, 28, 3, 28, 5, 28, 253, 10,
	28, 3, 29, 3, 29, 3, 30, 3, 30, 7, 30, 259, 10, 30, 12, 30, 14, 30, 262,
	11, 30, 3, 30, 5, 30, 265, 10, 30, 3, 31, 3, 31, 5, 31, 269, 10, 31, 3,
	32, 3, 32, 3, 32, 3, 32, 5, 32, 275, 10, 32, 3, 153, 2, 33, 3, 3, 5, 4,
	7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14,
	27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23,
	45, 24, 47, 25, 49, 26, 51, 2, 53, 2, 55, 2, 57, 2, 59, 2, 61, 2, 63, 2,
	3, 2, 20, 5, 2, 11, 12, 14, 15, 34, 34, 4, 2, 12, 12, 15, 15, 4, 2, 36,
	36, 94, 94, 6, 2, 12, 12, 15, 15, 94, 94, 98, 98, 4, 2, 11, 11, 34, 34,
	6, 2, 12, 12, 15, 15, 36, 36, 98, 98, 4, 2, 71, 71, 103, 103, 4, 2, 45,
	45, 47, 47, 10, 2, 36, 36, 41, 41, 94, 94, 100, 100, 104, 104, 112, 112,
	116, 116, 118, 118, 3, 2, 50, 53, 3, 2, 50, 57, 5, 2, 50, 59, 67, 72, 99,
	104, 3, 2, 50, 59, 4, 2, 50, 59, 97, 97, 6, 2, 38, 38, 67, 92, 97, 97,
	99, 124, 4, 2, 2, 129, 55298, 56321, 3, 2, 55298, 56321, 3, 2, 56322, 57345,
	2, 293, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3,
	2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17,
	3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2,
	25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2,
	2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2,
	2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2,
	2, 2, 2, 49, 3, 2, 2, 2, 3, 65, 3, 2, 2, 2, 5, 67, 3, 2, 2, 2, 7, 69, 3,
	2, 2, 2, 9, 71, 3, 2, 2, 2, 11, 73, 3, 2, 2, 2, 13, 75, 3, 2, 2, 2, 15,
	77, 3, 2, 2, 2, 17, 87, 3, 2, 2, 2, 19, 89, 3, 2, 2, 2, 21, 91, 3, 2, 2,
	2, 23, 99, 3, 2, 2, 2, 25, 101, 3, 2, 2, 2, 27, 103, 3, 2, 2, 2, 29, 106,
	3, 2, 2, 2, 31, 111, 3, 2, 2, 2, 33, 120, 3, 2, 2, 2, 35, 132, 3, 2, 2,
	2, 37, 141, 3, 2, 2, 2, 39, 147, 3, 2, 2, 2, 41, 161, 3, 2, 2, 2, 43, 172,
	3, 2, 2, 2, 45, 182, 3, 2, 2, 2, 47, 191, 3, 2, 2, 2, 49, 207, 3, 2, 2,
	2, 51, 214, 3, 2, 2, 2, 53, 241, 3, 2, 2, 2, 55, 243, 3, 2, 2, 2, 57, 254,
	3, 2, 2, 2, 59, 256, 3, 2, 2, 2, 61, 268, 3, 2, 2, 2, 63, 274, 3, 2, 2,
	2, 65, 66, 7, 63, 2, 2, 66, 4, 3, 2, 2, 2, 67, 68, 7, 42, 2, 2, 68, 6,
	3, 2, 2, 2, 69, 70, 7, 43, 2, 2, 70, 8, 3, 2, 2, 2, 71, 72, 7, 125, 2,
	2, 72, 10, 3, 2, 2, 2, 73, 74, 7, 127, 2, 2, 74, 12, 3, 2, 2, 2, 75, 76,
	7, 44, 2, 2, 76, 14, 3, 2, 2, 2, 77, 78, 7, 118, 2, 2, 78, 79, 7, 107,
	2, 2, 79, 80, 7, 111, 2, 2, 80, 81, 7, 103, 2, 2, 81, 82, 7, 48, 2, 2,
	82, 83, 7, 86, 2, 2, 83, 84, 7, 107, 2, 2, 84, 85, 7, 111, 2, 2, 85, 86,
	7, 103, 2, 2, 86, 16, 3, 2, 2, 2, 87, 88, 7, 93, 2, 2, 88, 18, 3, 2, 2,
	2, 89, 90, 7, 95, 2, 2, 90, 20, 3, 2, 2, 2, 91, 92, 7, 116, 2, 2, 92, 93,
	7, 103, 2, 2, 93, 94, 7, 118, 2, 2, 94, 95, 7, 119, 2, 2, 95, 96, 7, 116,
	2, 2, 96, 97, 7, 112, 2, 2, 97, 98, 7, 117, 2, 2, 98, 22, 3, 2, 2, 2, 99,
	100, 7, 47, 2, 2, 100, 24, 3, 2, 2, 2, 101, 102, 7, 49, 2, 2, 102, 26,
	3, 2, 2, 2, 103, 104, 7, 49, 2, 2, 104, 105, 7, 60, 2, 2, 105, 28, 3, 2,
	2, 2, 106, 107, 7, 66, 2, 2, 107, 108, 7, 102, 2, 2, 108, 109, 7, 113,
	2, 2, 109, 110, 7, 101, 2, 2, 110, 30, 3, 2, 2, 2, 111, 112, 7, 66, 2,
	2, 112, 113, 7, 106, 2, 2, 113, 114, 7, 99, 2, 2, 114, 115, 7, 112, 2,
	2, 115, 116, 7, 102, 2, 2, 116, 117, 7, 110, 2, 2, 117, 118, 7, 103, 2,
	2, 118, 119, 7, 116, 2, 2, 119, 32, 3, 2, 2, 2, 120, 121, 7, 107, 2, 2,
	121, 122, 7, 112, 2, 2, 122, 123, 7, 118, 2, 2, 123, 124, 7, 103, 2, 2,
	124, 125, 7, 116, 2, 2, 125, 126, 7, 104, 2, 2, 126, 127, 7, 99, 2, 2,
	127, 128, 7, 101, 2, 2, 128, 129, 7, 103, 2, 2, 129, 130, 7, 125, 2, 2,
	130, 131, 7, 127, 2, 2, 131, 34, 3, 2, 2, 2, 132, 133, 7, 66, 2, 2, 133,
	134, 7, 117, 2, 2, 134, 135, 7, 103, 2, 2, 135, 136, 7, 116, 2, 2, 136,
	137, 7, 120, 2, 2, 137, 138, 7, 103, 2, 2, 138, 139, 7, 116, 2, 2, 139,
	36, 3, 2, 2, 2, 140, 142, 9, 2, 2, 2, 141, 140, 3, 2, 2, 2, 142, 143, 3,
	2, 2, 2, 143, 141, 3, 2, 2, 2, 143, 144, 3, 2, 2, 2, 144, 145, 3, 2, 2,
	2, 145, 146, 8, 19, 2, 2, 146, 38, 3, 2, 2, 2, 147, 148, 7, 49, 2, 2, 148,
	149, 7, 44, 2, 2, 149, 153, 3, 2, 2, 2, 150, 152, 11, 2, 2, 2, 151, 150,
	3, 2, 2, 2, 152, 155, 3, 2, 2, 2, 153, 154, 3, 2, 2, 2, 153, 151, 3, 2,
	2, 2, 154, 156, 3, 2, 2, 2, 155, 153, 3, 2, 2, 2, 156, 157, 7, 44, 2, 2,
	157, 158, 7, 49, 2, 2, 158, 159, 3, 2, 2, 2, 159, 160, 8, 20, 3, 2, 160,
	40, 3, 2, 2, 2, 161, 162, 7, 49, 2, 2, 162, 163, 7, 49, 2, 2, 163, 167,
	3, 2, 2, 2, 164, 166, 10, 3, 2, 2, 165, 164, 3, 2, 2, 2, 166, 169, 3, 2,
	2, 2, 167, 165, 3, 2, 2, 2, 167, 168, 3, 2, 2, 2, 168, 170, 3, 2, 2, 2,
	169, 167, 3, 2, 2, 2, 170, 171, 8, 21, 3, 2, 171, 42, 3, 2, 2, 2, 172,
	177, 7, 36, 2, 2, 173, 176, 10, 4, 2, 2, 174, 176, 5, 53, 27, 2, 175, 173,
	3, 2, 2, 2, 175, 174, 3, 2, 2, 2, 176, 179, 3, 2, 2, 2, 177, 175, 3, 2,
	2, 2, 177, 178, 3, 2, 2, 2, 178, 180, 3, 2, 2, 2, 179, 177, 3, 2, 2, 2,
	180, 181, 7, 36, 2, 2, 181, 44, 3, 2, 2, 2, 182, 185, 7, 98, 2, 2, 183,
	186, 10, 5, 2, 2, 184, 186, 5, 53, 27, 2, 185, 183, 3, 2, 2, 2, 185, 184,
	3, 2, 2, 2, 186, 187, 3, 2, 2, 2, 187, 185, 3, 2, 2, 2, 187, 188, 3, 2,
	2, 2, 188, 189, 3, 2, 2, 2, 189, 190, 7, 98, 2, 2, 190, 46, 3, 2, 2, 2,
	191, 195, 7, 60, 2, 2, 192, 194, 9, 6, 2, 2, 193, 192, 3, 2, 2, 2, 194,
	197, 3, 2, 2, 2, 195, 193, 3, 2, 2, 2, 195, 196, 3, 2, 2, 2, 196, 205,
	3, 2, 2, 2, 197, 195, 3, 2, 2, 2, 198, 206, 5, 43, 22, 2, 199, 201, 10,
	7, 2, 2, 200, 199, 3, 2, 2, 2, 201, 204, 3, 2, 2, 2, 202, 200, 3, 2, 2,
	2, 202, 203, 3, 2, 2, 2, 203, 206, 3, 2, 2, 2, 204, 202, 3, 2, 2, 2, 205,
	198, 3, 2, 2, 2, 205, 202, 3, 2, 2, 2, 206, 48, 3, 2, 2, 2, 207, 211, 5,
	63, 32, 2, 208, 210, 5, 61, 31, 2, 209, 208, 3, 2, 2, 2, 210, 213, 3, 2,
	2, 2, 211, 209, 3, 2, 2, 2, 211, 212, 3, 2, 2, 2, 212, 50, 3, 2, 2, 2,
	213, 211, 3, 2, 2, 2, 214, 216, 9, 8, 2, 2, 215, 217, 9, 9, 2, 2, 216,
	215, 3, 2, 2, 2, 216, 217, 3, 2, 2, 2, 217, 218, 3, 2, 2, 2, 218, 219,
	5, 59, 30, 2, 219, 52, 3, 2, 2, 2, 220, 221, 7, 94, 2, 2, 221, 242, 9,
	10, 2, 2, 222, 227, 7, 94, 2, 2, 223, 225, 9, 11, 2, 2, 224, 223, 3, 2,
	2, 2, 224, 225, 3, 2, 2, 2, 225, 226, 3, 2, 2, 2, 226, 228, 9, 12, 2, 2,
	227, 224, 3, 2, 2, 2, 227, 228, 3, 2, 2, 2, 228, 229, 3, 2, 2, 2, 229,
	242, 9, 12, 2, 2, 230, 232, 7, 94, 2, 2, 231, 233, 7, 119, 2, 2, 232, 231,
	3, 2, 2, 2, 233, 234, 3, 2, 2, 2, 234, 232, 3, 2, 2, 2, 234, 235, 3, 2,
	2, 2, 235, 236, 3, 2, 2, 2, 236, 237, 5, 57, 29, 2, 237, 238, 5, 57, 29,
	2, 238, 239, 5, 57, 29, 2, 239, 240, 5, 57, 29, 2, 240, 242, 3, 2, 2, 2,
	241, 220, 3, 2, 2, 2, 241, 222, 3, 2, 2, 2, 241, 230, 3, 2, 2, 2, 242,
	54, 3, 2, 2, 2, 243, 252, 5, 57, 29, 2, 244, 247, 5, 57, 29, 2, 245, 247,
	7, 97, 2, 2, 246, 244, 3, 2, 2, 2, 246, 245, 3, 2, 2, 2, 247, 250, 3, 2,
	2, 2, 248, 246, 3, 2, 2, 2, 248, 249, 3, 2, 2, 2, 249, 251, 3, 2, 2, 2,
	250, 248, 3, 2, 2, 2, 251, 253, 5, 57, 29, 2, 252, 248, 3, 2, 2, 2, 252,
	253, 3, 2, 2, 2, 253, 56, 3, 2, 2, 2, 254, 255, 9, 13, 2, 2, 255, 58, 3,
	2, 2, 2, 256, 264, 9, 14, 2, 2, 257, 259, 9, 15, 2, 2, 258, 257, 3, 2,
	2, 2, 259, 262, 3, 2, 2, 2, 260, 258, 3, 2, 2, 2, 260, 261, 3, 2, 2, 2,
	261, 263, 3, 2, 2, 2, 262, 260, 3, 2, 2, 2, 263, 265, 9, 14, 2, 2, 264,
	260, 3, 2, 2, 2, 264, 265, 3, 2, 2, 2, 265, 60, 3, 2, 2, 2, 266, 269, 5,
	63, 32, 2, 267, 269, 9, 14, 2, 2, 268, 266, 3, 2, 2, 2, 268, 267, 3, 2,
	2, 2, 269, 62, 3, 2, 2, 2, 270, 275, 9, 16, 2, 2, 271, 275, 10, 17, 2,
	2, 272, 273, 9, 18, 2, 2, 273, 275, 9, 19, 2, 2, 274, 270, 3, 2, 2, 2,
	274, 271, 3, 2, 2, 2, 274, 272, 3, 2, 2, 2, 275, 64, 3, 2, 2, 2, 26, 2,
	143, 153, 167, 175, 177, 185, 187, 195, 202, 205, 211, 216, 224, 227, 234,
	241, 246, 248, 252, 260, 264, 268, 274, 4, 2, 3, 2, 2, 90, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'='", "'('", "')'", "'{'", "'}'", "'*'", "'time.Time'", "'['", "']'",
	"'returns'", "'-'", "'/'", "'/:'", "'@doc'", "'@handler'", "'interface{}'",
	"'@server'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "ATDOC", "ATHANDLER",
	"INTERFACE", "ATSERVER", "WS", "COMMENT", "LINE_COMMENT", "STRING", "RAW_STRING",
	"LINE_VALUE", "ID",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "ATDOC", "ATHANDLER", "INTERFACE", "ATSERVER",
	"WS", "COMMENT", "LINE_COMMENT", "STRING", "RAW_STRING", "LINE_VALUE",
	"ID", "ExponentPart", "EscapeSequence", "HexDigits", "HexDigit", "Digits",
	"LetterOrDigit", "Letter",
}

type ApiParserLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewApiParserLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *ApiParserLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewApiParserLexer(input antlr.CharStream) *ApiParserLexer {
	l := new(ApiParserLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "ApiParser.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ApiParserLexer tokens.
const (
	ApiParserLexerT__0         = 1
	ApiParserLexerT__1         = 2
	ApiParserLexerT__2         = 3
	ApiParserLexerT__3         = 4
	ApiParserLexerT__4         = 5
	ApiParserLexerT__5         = 6
	ApiParserLexerT__6         = 7
	ApiParserLexerT__7         = 8
	ApiParserLexerT__8         = 9
	ApiParserLexerT__9         = 10
	ApiParserLexerT__10        = 11
	ApiParserLexerT__11        = 12
	ApiParserLexerT__12        = 13
	ApiParserLexerATDOC        = 14
	ApiParserLexerATHANDLER    = 15
	ApiParserLexerINTERFACE    = 16
	ApiParserLexerATSERVER     = 17
	ApiParserLexerWS           = 18
	ApiParserLexerCOMMENT      = 19
	ApiParserLexerLINE_COMMENT = 20
	ApiParserLexerSTRING       = 21
	ApiParserLexerRAW_STRING   = 22
	ApiParserLexerLINE_VALUE   = 23
	ApiParserLexerID           = 24
)

const COMEMNTS = 88
