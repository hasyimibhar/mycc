// Code generated from IR.g4 by ANTLR 4.9. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 41, 317,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 7,
	5, 123, 10, 5, 12, 5, 14, 5, 126, 11, 5, 3, 5, 3, 5, 3, 6, 6, 6, 131, 10,
	6, 13, 6, 14, 6, 132, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18,
	3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3,
	24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 29,
	3, 29, 3, 30, 3, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3,
	34, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36,
	7, 36, 261, 10, 36, 12, 36, 14, 36, 264, 11, 36, 3, 37, 6, 37, 267, 10,
	37, 13, 37, 14, 37, 268, 3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 3, 40, 5, 40,
	277, 10, 40, 3, 41, 3, 41, 5, 41, 281, 10, 41, 3, 42, 3, 42, 5, 42, 285,
	10, 42, 3, 43, 3, 43, 3, 44, 3, 44, 5, 44, 291, 10, 44, 3, 45, 3, 45, 7,
	45, 295, 10, 45, 12, 45, 14, 45, 298, 11, 45, 3, 45, 3, 45, 3, 46, 3, 46,
	3, 46, 3, 47, 3, 47, 3, 47, 7, 47, 308, 10, 47, 12, 47, 14, 47, 311, 11,
	47, 3, 47, 3, 47, 3, 48, 3, 48, 3, 48, 3, 309, 2, 49, 3, 3, 5, 4, 7, 5,
	9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27,
	15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45,
	24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63,
	33, 65, 34, 67, 35, 69, 36, 71, 37, 73, 38, 75, 2, 77, 2, 79, 2, 81, 2,
	83, 2, 85, 2, 87, 2, 89, 39, 91, 2, 93, 40, 95, 41, 3, 2, 6, 4, 2, 12,
	12, 15, 15, 5, 2, 11, 12, 15, 15, 34, 34, 5, 2, 38, 38, 47, 48, 97, 97,
	4, 2, 67, 72, 99, 104, 2, 320, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7,
	3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2,
	15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2,
	2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2,
	2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2,
	2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3,
	2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53,
	3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2,
	61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2,
	2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 89, 3, 2, 2,
	2, 2, 93, 3, 2, 2, 2, 2, 95, 3, 2, 2, 2, 3, 97, 3, 2, 2, 2, 5, 101, 3,
	2, 2, 2, 7, 106, 3, 2, 2, 2, 9, 120, 3, 2, 2, 2, 11, 130, 3, 2, 2, 2, 13,
	136, 3, 2, 2, 2, 15, 144, 3, 2, 2, 2, 17, 151, 3, 2, 2, 2, 19, 159, 3,
	2, 2, 2, 21, 168, 3, 2, 2, 2, 23, 178, 3, 2, 2, 2, 25, 191, 3, 2, 2, 2,
	27, 196, 3, 2, 2, 2, 29, 203, 3, 2, 2, 2, 31, 212, 3, 2, 2, 2, 33, 214,
	3, 2, 2, 2, 35, 216, 3, 2, 2, 2, 37, 218, 3, 2, 2, 2, 39, 220, 3, 2, 2,
	2, 41, 222, 3, 2, 2, 2, 43, 224, 3, 2, 2, 2, 45, 226, 3, 2, 2, 2, 47, 228,
	3, 2, 2, 2, 49, 230, 3, 2, 2, 2, 51, 232, 3, 2, 2, 2, 53, 234, 3, 2, 2,
	2, 55, 236, 3, 2, 2, 2, 57, 238, 3, 2, 2, 2, 59, 240, 3, 2, 2, 2, 61, 242,
	3, 2, 2, 2, 63, 244, 3, 2, 2, 2, 65, 246, 3, 2, 2, 2, 67, 249, 3, 2, 2,
	2, 69, 253, 3, 2, 2, 2, 71, 257, 3, 2, 2, 2, 73, 266, 3, 2, 2, 2, 75, 270,
	3, 2, 2, 2, 77, 272, 3, 2, 2, 2, 79, 276, 3, 2, 2, 2, 81, 280, 3, 2, 2,
	2, 83, 284, 3, 2, 2, 2, 85, 286, 3, 2, 2, 2, 87, 290, 3, 2, 2, 2, 89, 292,
	3, 2, 2, 2, 91, 301, 3, 2, 2, 2, 93, 304, 3, 2, 2, 2, 95, 314, 3, 2, 2,
	2, 97, 98, 7, 116, 2, 2, 98, 99, 7, 103, 2, 2, 99, 100, 7, 118, 2, 2, 100,
	4, 3, 2, 2, 2, 101, 102, 7, 101, 2, 2, 102, 103, 7, 99, 2, 2, 103, 104,
	7, 110, 2, 2, 104, 105, 7, 110, 2, 2, 105, 6, 3, 2, 2, 2, 106, 107, 7,
	105, 2, 2, 107, 108, 7, 103, 2, 2, 108, 109, 7, 118, 2, 2, 109, 110, 7,
	103, 2, 2, 110, 111, 7, 110, 2, 2, 111, 112, 7, 103, 2, 2, 112, 113, 7,
	111, 2, 2, 113, 114, 7, 103, 2, 2, 114, 115, 7, 112, 2, 2, 115, 116, 7,
	118, 2, 2, 116, 117, 7, 114, 2, 2, 117, 118, 7, 118, 2, 2, 118, 119, 7,
	116, 2, 2, 119, 8, 3, 2, 2, 2, 120, 124, 5, 53, 27, 2, 121, 123, 10, 2,
	2, 2, 122, 121, 3, 2, 2, 2, 123, 126, 3, 2, 2, 2, 124, 122, 3, 2, 2, 2,
	124, 125, 3, 2, 2, 2, 125, 127, 3, 2, 2, 2, 126, 124, 3, 2, 2, 2, 127,
	128, 8, 5, 2, 2, 128, 10, 3, 2, 2, 2, 129, 131, 9, 3, 2, 2, 130, 129, 3,
	2, 2, 2, 131, 132, 3, 2, 2, 2, 132, 130, 3, 2, 2, 2, 132, 133, 3, 2, 2,
	2, 133, 134, 3, 2, 2, 2, 134, 135, 8, 6, 2, 2, 135, 12, 3, 2, 2, 2, 136,
	137, 7, 102, 2, 2, 137, 138, 7, 103, 2, 2, 138, 139, 7, 101, 2, 2, 139,
	140, 7, 110, 2, 2, 140, 141, 7, 99, 2, 2, 141, 142, 7, 116, 2, 2, 142,
	143, 7, 103, 2, 2, 143, 14, 3, 2, 2, 2, 144, 145, 7, 102, 2, 2, 145, 146,
	7, 103, 2, 2, 146, 147, 7, 104, 2, 2, 147, 148, 7, 107, 2, 2, 148, 149,
	7, 112, 2, 2, 149, 150, 7, 103, 2, 2, 150, 16, 3, 2, 2, 2, 151, 152, 7,
	114, 2, 2, 152, 153, 7, 116, 2, 2, 153, 154, 7, 107, 2, 2, 154, 155, 7,
	120, 2, 2, 155, 156, 7, 99, 2, 2, 156, 157, 7, 118, 2, 2, 157, 158, 7,
	103, 2, 2, 158, 18, 3, 2, 2, 2, 159, 160, 7, 112, 2, 2, 160, 161, 7, 113,
	2, 2, 161, 162, 7, 119, 2, 2, 162, 163, 7, 112, 2, 2, 163, 164, 7, 121,
	2, 2, 164, 165, 7, 107, 2, 2, 165, 166, 7, 112, 2, 2, 166, 167, 7, 102,
	2, 2, 167, 20, 3, 2, 2, 2, 168, 169, 7, 112, 2, 2, 169, 170, 7, 113, 2,
	2, 170, 171, 7, 101, 2, 2, 171, 172, 7, 99, 2, 2, 172, 173, 7, 114, 2,
	2, 173, 174, 7, 118, 2, 2, 174, 175, 7, 119, 2, 2, 175, 176, 7, 116, 2,
	2, 176, 177, 7, 103, 2, 2, 177, 22, 3, 2, 2, 2, 178, 179, 7, 119, 2, 2,
	179, 180, 7, 112, 2, 2, 180, 181, 7, 112, 2, 2, 181, 182, 7, 99, 2, 2,
	182, 183, 7, 111, 2, 2, 183, 184, 7, 103, 2, 2, 184, 185, 7, 102, 2, 2,
	185, 186, 7, 97, 2, 2, 186, 187, 7, 99, 2, 2, 187, 188, 7, 102, 2, 2, 188,
	189, 7, 102, 2, 2, 189, 190, 7, 116, 2, 2, 190, 24, 3, 2, 2, 2, 191, 192,
	7, 112, 2, 2, 192, 193, 7, 119, 2, 2, 193, 194, 7, 110, 2, 2, 194, 195,
	7, 110, 2, 2, 195, 26, 3, 2, 2, 2, 196, 197, 7, 105, 2, 2, 197, 198, 7,
	110, 2, 2, 198, 199, 7, 113, 2, 2, 199, 200, 7, 100, 2, 2, 200, 201, 7,
	99, 2, 2, 201, 202, 7, 110, 2, 2, 202, 28, 3, 2, 2, 2, 203, 204, 7, 101,
	2, 2, 204, 205, 7, 113, 2, 2, 205, 206, 7, 112, 2, 2, 206, 207, 7, 117,
	2, 2, 207, 208, 7, 118, 2, 2, 208, 209, 7, 99, 2, 2, 209, 210, 7, 112,
	2, 2, 210, 211, 7, 118, 2, 2, 211, 30, 3, 2, 2, 2, 212, 213, 7, 42, 2,
	2, 213, 32, 3, 2, 2, 2, 214, 215, 7, 43, 2, 2, 215, 34, 3, 2, 2, 2, 216,
	217, 7, 93, 2, 2, 217, 36, 3, 2, 2, 2, 218, 219, 7, 95, 2, 2, 219, 38,
	3, 2, 2, 2, 220, 221, 7, 125, 2, 2, 221, 40, 3, 2, 2, 2, 222, 223, 7, 127,
	2, 2, 223, 42, 3, 2, 2, 2, 224, 225, 7, 63, 2, 2, 225, 44, 3, 2, 2, 2,
	226, 227, 7, 36, 2, 2, 227, 46, 3, 2, 2, 2, 228, 229, 7, 46, 2, 2, 229,
	48, 3, 2, 2, 2, 230, 231, 7, 44, 2, 2, 231, 50, 3, 2, 2, 2, 232, 233, 7,
	39, 2, 2, 233, 52, 3, 2, 2, 2, 234, 235, 7, 61, 2, 2, 235, 54, 3, 2, 2,
	2, 236, 237, 7, 35, 2, 2, 237, 56, 3, 2, 2, 2, 238, 239, 7, 66, 2, 2, 239,
	58, 3, 2, 2, 2, 240, 241, 7, 48, 2, 2, 241, 60, 3, 2, 2, 2, 242, 243, 7,
	122, 2, 2, 243, 62, 3, 2, 2, 2, 244, 245, 7, 47, 2, 2, 245, 64, 3, 2, 2,
	2, 246, 247, 7, 107, 2, 2, 247, 248, 7, 58, 2, 2, 248, 66, 3, 2, 2, 2,
	249, 250, 7, 107, 2, 2, 250, 251, 7, 53, 2, 2, 251, 252, 7, 52, 2, 2, 252,
	68, 3, 2, 2, 2, 253, 254, 7, 107, 2, 2, 254, 255, 7, 56, 2, 2, 255, 256,
	7, 54, 2, 2, 256, 70, 3, 2, 2, 2, 257, 262, 5, 81, 41, 2, 258, 261, 5,
	81, 41, 2, 259, 261, 5, 85, 43, 2, 260, 258, 3, 2, 2, 2, 260, 259, 3, 2,
	2, 2, 261, 264, 3, 2, 2, 2, 262, 260, 3, 2, 2, 2, 262, 263, 3, 2, 2, 2,
	263, 72, 3, 2, 2, 2, 264, 262, 3, 2, 2, 2, 265, 267, 5, 85, 43, 2, 266,
	265, 3, 2, 2, 2, 267, 268, 3, 2, 2, 2, 268, 266, 3, 2, 2, 2, 268, 269,
	3, 2, 2, 2, 269, 74, 3, 2, 2, 2, 270, 271, 4, 67, 92, 2, 271, 76, 3, 2,
	2, 2, 272, 273, 4, 99, 124, 2, 273, 78, 3, 2, 2, 2, 274, 277, 5, 75, 38,
	2, 275, 277, 5, 77, 39, 2, 276, 274, 3, 2, 2, 2, 276, 275, 3, 2, 2, 2,
	277, 80, 3, 2, 2, 2, 278, 281, 5, 79, 40, 2, 279, 281, 9, 4, 2, 2, 280,
	278, 3, 2, 2, 2, 280, 279, 3, 2, 2, 2, 281, 82, 3, 2, 2, 2, 282, 285, 5,
	81, 41, 2, 283, 285, 7, 94, 2, 2, 284, 282, 3, 2, 2, 2, 284, 283, 3, 2,
	2, 2, 285, 84, 3, 2, 2, 2, 286, 287, 4, 50, 59, 2, 287, 86, 3, 2, 2, 2,
	288, 291, 5, 85, 43, 2, 289, 291, 9, 5, 2, 2, 290, 288, 3, 2, 2, 2, 290,
	289, 3, 2, 2, 2, 291, 88, 3, 2, 2, 2, 292, 296, 5, 55, 28, 2, 293, 295,
	10, 2, 2, 2, 294, 293, 3, 2, 2, 2, 295, 298, 3, 2, 2, 2, 296, 294, 3, 2,
	2, 2, 296, 297, 3, 2, 2, 2, 297, 299, 3, 2, 2, 2, 298, 296, 3, 2, 2, 2,
	299, 300, 8, 45, 2, 2, 300, 90, 3, 2, 2, 2, 301, 302, 7, 94, 2, 2, 302,
	303, 7, 36, 2, 2, 303, 92, 3, 2, 2, 2, 304, 309, 7, 36, 2, 2, 305, 308,
	5, 91, 46, 2, 306, 308, 10, 2, 2, 2, 307, 305, 3, 2, 2, 2, 307, 306, 3,
	2, 2, 2, 308, 311, 3, 2, 2, 2, 309, 310, 3, 2, 2, 2, 309, 307, 3, 2, 2,
	2, 310, 312, 3, 2, 2, 2, 311, 309, 3, 2, 2, 2, 312, 313, 7, 36, 2, 2, 313,
	94, 3, 2, 2, 2, 314, 315, 7, 101, 2, 2, 315, 316, 5, 93, 47, 2, 316, 96,
	3, 2, 2, 2, 15, 2, 124, 132, 260, 262, 268, 276, 280, 284, 290, 296, 307,
	309, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'ret'", "'call'", "'getelementptr'", "", "", "'declare'", "'define'",
	"'private'", "'nounwind'", "'nocapture'", "'unnamed_addr'", "'null'", "'global'",
	"'constant'", "'('", "')'", "'['", "']'", "'{'", "'}'", "'='", "'\"'",
	"','", "'*'", "'%'", "';'", "'!'", "'@'", "'.'", "'x'", "'-'", "'i8'",
	"'i32'", "'i64'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "COMMENT", "WHITESPACE", "DECLARE", "DEFINE", "PRIVATE",
	"NOUNWIND", "NOCAPTURE", "UNNAMMED_ADDR", "NULL", "GLOBAL", "CONSTANT",
	"LPAREN", "RPAREN", "LBRACK", "RBRACK", "LBRACE", "RBRACE", "EQUAL", "DOUBLEQUOTE",
	"COMMA", "STAR", "PERCENT", "SEMICOLON", "EXCLAMATION", "AT", "DOT", "MULTIPLY",
	"MINUS", "TYPE_I8", "TYPE_I32", "TYPE_I64", "NAME", "DECIMALS", "METADATA",
	"QUOTED_STRING", "STRING_LIT",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "COMMENT", "WHITESPACE", "DECLARE", "DEFINE", "PRIVATE",
	"NOUNWIND", "NOCAPTURE", "UNNAMMED_ADDR", "NULL", "GLOBAL", "CONSTANT",
	"LPAREN", "RPAREN", "LBRACK", "RBRACK", "LBRACE", "RBRACE", "EQUAL", "DOUBLEQUOTE",
	"COMMA", "STAR", "PERCENT", "SEMICOLON", "EXCLAMATION", "AT", "DOT", "MULTIPLY",
	"MINUS", "TYPE_I8", "TYPE_I32", "TYPE_I64", "NAME", "DECIMALS", "ASCII_LETTER_UPPER",
	"ASCII_LETTER_LOWER", "ASCII_LETTER", "LETTER", "ESCAPE_LETTER", "DECIMAL_DIGIT",
	"HEX_DIGIT", "METADATA", "ESCAPED_QUOTE", "QUOTED_STRING", "STRING_LIT",
}

type IRLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewIRLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *IRLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewIRLexer(input antlr.CharStream) *IRLexer {
	l := new(IRLexer)
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
	l.GrammarFileName = "IR.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// IRLexer tokens.
const (
	IRLexerT__0          = 1
	IRLexerT__1          = 2
	IRLexerT__2          = 3
	IRLexerCOMMENT       = 4
	IRLexerWHITESPACE    = 5
	IRLexerDECLARE       = 6
	IRLexerDEFINE        = 7
	IRLexerPRIVATE       = 8
	IRLexerNOUNWIND      = 9
	IRLexerNOCAPTURE     = 10
	IRLexerUNNAMMED_ADDR = 11
	IRLexerNULL          = 12
	IRLexerGLOBAL        = 13
	IRLexerCONSTANT      = 14
	IRLexerLPAREN        = 15
	IRLexerRPAREN        = 16
	IRLexerLBRACK        = 17
	IRLexerRBRACK        = 18
	IRLexerLBRACE        = 19
	IRLexerRBRACE        = 20
	IRLexerEQUAL         = 21
	IRLexerDOUBLEQUOTE   = 22
	IRLexerCOMMA         = 23
	IRLexerSTAR          = 24
	IRLexerPERCENT       = 25
	IRLexerSEMICOLON     = 26
	IRLexerEXCLAMATION   = 27
	IRLexerAT            = 28
	IRLexerDOT           = 29
	IRLexerMULTIPLY      = 30
	IRLexerMINUS         = 31
	IRLexerTYPE_I8       = 32
	IRLexerTYPE_I32      = 33
	IRLexerTYPE_I64      = 34
	IRLexerNAME          = 35
	IRLexerDECIMALS      = 36
	IRLexerMETADATA      = 37
	IRLexerQUOTED_STRING = 38
	IRLexerSTRING_LIT    = 39
)