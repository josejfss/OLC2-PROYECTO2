// Code generated from optimizador_parser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // optimizador_parser

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import "OLC2-PROYECTO2/OPT_COMPILADOR/INTER"
import "OLC2-PROYECTO2/OPT_COMPILADOR/BLOQUE"

/* EXPRESIONES */
/* OPERACION ARITMETICA */
import "OLC2-PROYECTO2/OPT_COMPILADOR/AST/EXPRESION/OPERACION/ARITMETICA/SUMA"
import "OLC2-PROYECTO2/OPT_COMPILADOR/AST/EXPRESION/OPERACION/ARITMETICA/RESTA"
import "OLC2-PROYECTO2/OPT_COMPILADOR/AST/EXPRESION/OPERACION/ARITMETICA/MULTIPLICACION"
import "OLC2-PROYECTO2/OPT_COMPILADOR/AST/EXPRESION/OPERACION/ARITMETICA/DIVISION"
import "OLC2-PROYECTO2/OPT_COMPILADOR/AST/EXPRESION/OPERACION/ARITMETICA/MODULO"
import "OLC2-PROYECTO2/OPT_COMPILADOR/AST/EXPRESION/OPERACION/ARITMETICA/NEGATIVO"

/* OPERACION RELACIONAL */
import "OLC2-PROYECTO2/OPT_COMPILADOR/AST/EXPRESION/OPERACION/RELACIONAL/DESIGUALDAD"
import "OLC2-PROYECTO2/OPT_COMPILADOR/AST/EXPRESION/OPERACION/RELACIONAL/IGUALDAD"
import "OLC2-PROYECTO2/OPT_COMPILADOR/AST/EXPRESION/OPERACION/RELACIONAL/MAYOR"
import "OLC2-PROYECTO2/OPT_COMPILADOR/AST/EXPRESION/OPERACION/RELACIONAL/MAYORIGUAL"
import "OLC2-PROYECTO2/OPT_COMPILADOR/AST/EXPRESION/OPERACION/RELACIONAL/MENOR"
import "OLC2-PROYECTO2/OPT_COMPILADOR/AST/EXPRESION/OPERACION/RELACIONAL/MENORIGUAL"

/* PRIMITIVOS /home/iovana/go/src/  */
import "OLC2-PROYECTO2/OPT_COMPILADOR/AST/EXPRESION/PRIMITIIVOS"

/* ASIGNACIONES */
import "OLC2-PROYECTO2/OPT_COMPILADOR/AST/EXPRESION/ASIGNACIONES"

/* FUNCIONES */
/* ENCABEZADO */
import "OLC2-PROYECTO2/OPT_COMPILADOR/AST/INSTRUCCION/ENCABEZADO"
import "OLC2-PROYECTO2/OPT_COMPILADOR/AST/INSTRUCCION/FUNCION"
import "OLC2-PROYECTO2/OPT_COMPILADOR/AST/INSTRUCCION/DECLARAR"
import "OLC2-PROYECTO2/OPT_COMPILADOR/AST/INSTRUCCION/SIF"
import "OLC2-PROYECTO2/OPT_COMPILADOR/AST/INSTRUCCION/IMPRIMIR"
import "OLC2-PROYECTO2/OPT_COMPILADOR/AST/INSTRUCCION/ETIQUETAS"
import "OLC2-PROYECTO2/OPT_COMPILADOR/AST/INSTRUCCION/LLAMADA"
import arrayList "github.com/colegno/arraylist"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 54, 295,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 3, 2, 3, 2, 3, 2, 3, 3, 7,
	3, 47, 10, 3, 12, 3, 14, 3, 50, 11, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 76, 10, 6, 12, 6, 14, 6, 79,
	11, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 89, 10, 7,
	12, 7, 14, 7, 92, 11, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 5, 8, 116, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 142, 10, 11, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	5, 12, 166, 10, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5,
	13, 175, 10, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 7, 15, 191, 10, 15, 12, 15, 14,
	15, 194, 11, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 211, 10, 16, 3,
	17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 222,
	10, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 238, 10, 18, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 7, 19, 249, 10, 19, 12, 19,
	14, 19, 252, 11, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3,
	20, 5, 20, 262, 10, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 20, 7, 20, 274, 10, 20, 12, 20, 14, 20, 277, 11, 20, 3,
	21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21,
	3, 21, 3, 21, 3, 21, 5, 21, 293, 10, 21, 3, 21, 2, 7, 10, 12, 28, 36, 38,
	22, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36,
	38, 40, 2, 7, 3, 2, 11, 12, 3, 2, 18, 19, 3, 2, 36, 41, 3, 2, 33, 35, 3,
	2, 31, 32, 2, 301, 2, 42, 3, 2, 2, 2, 4, 48, 3, 2, 2, 2, 6, 53, 3, 2, 2,
	2, 8, 56, 3, 2, 2, 2, 10, 67, 3, 2, 2, 2, 12, 80, 3, 2, 2, 2, 14, 115,
	3, 2, 2, 2, 16, 117, 3, 2, 2, 2, 18, 126, 3, 2, 2, 2, 20, 141, 3, 2, 2,
	2, 22, 165, 3, 2, 2, 2, 24, 174, 3, 2, 2, 2, 26, 176, 3, 2, 2, 2, 28, 182,
	3, 2, 2, 2, 30, 210, 3, 2, 2, 2, 32, 221, 3, 2, 2, 2, 34, 237, 3, 2, 2,
	2, 36, 239, 3, 2, 2, 2, 38, 261, 3, 2, 2, 2, 40, 292, 3, 2, 2, 2, 42, 43,
	5, 4, 3, 2, 43, 44, 8, 2, 1, 2, 44, 3, 3, 2, 2, 2, 45, 47, 5, 6, 4, 2,
	46, 45, 3, 2, 2, 2, 47, 50, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 48, 49, 3,
	2, 2, 2, 49, 51, 3, 2, 2, 2, 50, 48, 3, 2, 2, 2, 51, 52, 8, 3, 1, 2, 52,
	5, 3, 2, 2, 2, 53, 54, 5, 8, 5, 2, 54, 55, 8, 4, 1, 2, 55, 7, 3, 2, 2,
	2, 56, 57, 7, 3, 2, 2, 57, 58, 7, 4, 2, 2, 58, 59, 7, 5, 2, 2, 59, 60,
	7, 6, 2, 2, 60, 61, 7, 7, 2, 2, 61, 62, 7, 21, 2, 2, 62, 63, 5, 10, 6,
	2, 63, 64, 7, 48, 2, 2, 64, 65, 5, 12, 7, 2, 65, 66, 8, 5, 1, 2, 66, 9,
	3, 2, 2, 2, 67, 68, 8, 6, 1, 2, 68, 69, 7, 28, 2, 2, 69, 70, 8, 6, 1, 2,
	70, 77, 3, 2, 2, 2, 71, 72, 12, 4, 2, 2, 72, 73, 7, 50, 2, 2, 73, 74, 7,
	28, 2, 2, 74, 76, 8, 6, 1, 2, 75, 71, 3, 2, 2, 2, 76, 79, 3, 2, 2, 2, 77,
	75, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 78, 11, 3, 2, 2, 2, 79, 77, 3, 2, 2,
	2, 80, 81, 8, 7, 1, 2, 81, 82, 5, 14, 8, 2, 82, 83, 8, 7, 1, 2, 83, 90,
	3, 2, 2, 2, 84, 85, 12, 4, 2, 2, 85, 86, 5, 14, 8, 2, 86, 87, 8, 7, 1,
	2, 87, 89, 3, 2, 2, 2, 88, 84, 3, 2, 2, 2, 89, 92, 3, 2, 2, 2, 90, 88,
	3, 2, 2, 2, 90, 91, 3, 2, 2, 2, 91, 13, 3, 2, 2, 2, 92, 90, 3, 2, 2, 2,
	93, 94, 7, 8, 2, 2, 94, 95, 7, 30, 2, 2, 95, 96, 7, 46, 2, 2, 96, 97, 7,
	47, 2, 2, 97, 98, 7, 42, 2, 2, 98, 99, 5, 28, 15, 2, 99, 100, 7, 10, 2,
	2, 100, 101, 7, 43, 2, 2, 101, 102, 8, 8, 1, 2, 102, 116, 3, 2, 2, 2, 103,
	104, 7, 22, 2, 2, 104, 105, 7, 16, 2, 2, 105, 106, 7, 46, 2, 2, 106, 107,
	7, 47, 2, 2, 107, 108, 7, 42, 2, 2, 108, 109, 7, 23, 2, 2, 109, 110, 7,
	24, 2, 2, 110, 111, 5, 28, 15, 2, 111, 112, 7, 17, 2, 2, 112, 113, 7, 43,
	2, 2, 113, 114, 8, 8, 1, 2, 114, 116, 3, 2, 2, 2, 115, 93, 3, 2, 2, 2,
	115, 103, 3, 2, 2, 2, 116, 15, 3, 2, 2, 2, 117, 118, 7, 20, 2, 2, 118,
	119, 7, 46, 2, 2, 119, 120, 5, 32, 17, 2, 120, 121, 7, 47, 2, 2, 121, 122,
	7, 9, 2, 2, 122, 123, 7, 29, 2, 2, 123, 124, 7, 48, 2, 2, 124, 125, 8,
	9, 1, 2, 125, 17, 3, 2, 2, 2, 126, 127, 7, 15, 2, 2, 127, 128, 7, 46, 2,
	2, 128, 129, 7, 27, 2, 2, 129, 130, 7, 50, 2, 2, 130, 131, 5, 20, 11, 2,
	131, 132, 5, 32, 17, 2, 132, 133, 7, 47, 2, 2, 133, 134, 7, 48, 2, 2, 134,
	135, 8, 10, 1, 2, 135, 19, 3, 2, 2, 2, 136, 137, 7, 14, 2, 2, 137, 142,
	8, 11, 1, 2, 138, 139, 7, 13, 2, 2, 139, 142, 8, 11, 1, 2, 140, 142, 8,
	11, 1, 2, 141, 136, 3, 2, 2, 2, 141, 138, 3, 2, 2, 2, 141, 140, 3, 2, 2,
	2, 142, 21, 3, 2, 2, 2, 143, 144, 7, 28, 2, 2, 144, 145, 7, 51, 2, 2, 145,
	146, 5, 32, 17, 2, 146, 147, 7, 48, 2, 2, 147, 148, 8, 12, 1, 2, 148, 166,
	3, 2, 2, 2, 149, 150, 9, 2, 2, 2, 150, 151, 7, 44, 2, 2, 151, 152, 7, 13,
	2, 2, 152, 153, 5, 32, 17, 2, 153, 154, 7, 45, 2, 2, 154, 155, 7, 51, 2,
	2, 155, 156, 5, 32, 17, 2, 156, 157, 7, 48, 2, 2, 157, 158, 8, 12, 1, 2,
	158, 166, 3, 2, 2, 2, 159, 160, 9, 3, 2, 2, 160, 161, 7, 51, 2, 2, 161,
	162, 5, 32, 17, 2, 162, 163, 7, 48, 2, 2, 163, 164, 8, 12, 1, 2, 164, 166,
	3, 2, 2, 2, 165, 143, 3, 2, 2, 2, 165, 149, 3, 2, 2, 2, 165, 159, 3, 2,
	2, 2, 166, 23, 3, 2, 2, 2, 167, 168, 7, 29, 2, 2, 168, 169, 7, 49, 2, 2,
	169, 175, 8, 13, 1, 2, 170, 171, 7, 9, 2, 2, 171, 172, 7, 29, 2, 2, 172,
	173, 7, 48, 2, 2, 173, 175, 8, 13, 1, 2, 174, 167, 3, 2, 2, 2, 174, 170,
	3, 2, 2, 2, 175, 25, 3, 2, 2, 2, 176, 177, 7, 30, 2, 2, 177, 178, 7, 46,
	2, 2, 178, 179, 7, 47, 2, 2, 179, 180, 7, 48, 2, 2, 180, 181, 8, 14, 1,
	2, 181, 27, 3, 2, 2, 2, 182, 183, 8, 15, 1, 2, 183, 184, 5, 30, 16, 2,
	184, 185, 8, 15, 1, 2, 185, 192, 3, 2, 2, 2, 186, 187, 12, 4, 2, 2, 187,
	188, 5, 30, 16, 2, 188, 189, 8, 15, 1, 2, 189, 191, 3, 2, 2, 2, 190, 186,
	3, 2, 2, 2, 191, 194, 3, 2, 2, 2, 192, 190, 3, 2, 2, 2, 192, 193, 3, 2,
	2, 2, 193, 29, 3, 2, 2, 2, 194, 192, 3, 2, 2, 2, 195, 196, 5, 18, 10, 2,
	196, 197, 8, 16, 1, 2, 197, 211, 3, 2, 2, 2, 198, 199, 5, 16, 9, 2, 199,
	200, 8, 16, 1, 2, 200, 211, 3, 2, 2, 2, 201, 202, 5, 22, 12, 2, 202, 203,
	8, 16, 1, 2, 203, 211, 3, 2, 2, 2, 204, 205, 5, 24, 13, 2, 205, 206, 8,
	16, 1, 2, 206, 211, 3, 2, 2, 2, 207, 208, 5, 26, 14, 2, 208, 209, 8, 16,
	1, 2, 209, 211, 3, 2, 2, 2, 210, 195, 3, 2, 2, 2, 210, 198, 3, 2, 2, 2,
	210, 201, 3, 2, 2, 2, 210, 204, 3, 2, 2, 2, 210, 207, 3, 2, 2, 2, 211,
	31, 3, 2, 2, 2, 212, 213, 5, 36, 19, 2, 213, 214, 8, 17, 1, 2, 214, 222,
	3, 2, 2, 2, 215, 216, 5, 38, 20, 2, 216, 217, 8, 17, 1, 2, 217, 222, 3,
	2, 2, 2, 218, 219, 5, 34, 18, 2, 219, 220, 8, 17, 1, 2, 220, 222, 3, 2,
	2, 2, 221, 212, 3, 2, 2, 2, 221, 215, 3, 2, 2, 2, 221, 218, 3, 2, 2, 2,
	222, 33, 3, 2, 2, 2, 223, 224, 7, 11, 2, 2, 224, 225, 7, 44, 2, 2, 225,
	226, 7, 13, 2, 2, 226, 227, 5, 32, 17, 2, 227, 228, 7, 45, 2, 2, 228, 229,
	8, 18, 1, 2, 229, 238, 3, 2, 2, 2, 230, 231, 7, 12, 2, 2, 231, 232, 7,
	44, 2, 2, 232, 233, 7, 13, 2, 2, 233, 234, 5, 32, 17, 2, 234, 235, 7, 45,
	2, 2, 235, 236, 8, 18, 1, 2, 236, 238, 3, 2, 2, 2, 237, 223, 3, 2, 2, 2,
	237, 230, 3, 2, 2, 2, 238, 35, 3, 2, 2, 2, 239, 240, 8, 19, 1, 2, 240,
	241, 5, 38, 20, 2, 241, 242, 8, 19, 1, 2, 242, 250, 3, 2, 2, 2, 243, 244,
	12, 4, 2, 2, 244, 245, 9, 4, 2, 2, 245, 246, 5, 36, 19, 5, 246, 247, 8,
	19, 1, 2, 247, 249, 3, 2, 2, 2, 248, 243, 3, 2, 2, 2, 249, 252, 3, 2, 2,
	2, 250, 248, 3, 2, 2, 2, 250, 251, 3, 2, 2, 2, 251, 37, 3, 2, 2, 2, 252,
	250, 3, 2, 2, 2, 253, 254, 8, 20, 1, 2, 254, 255, 7, 32, 2, 2, 255, 256,
	5, 38, 20, 6, 256, 257, 8, 20, 1, 2, 257, 262, 3, 2, 2, 2, 258, 259, 5,
	40, 21, 2, 259, 260, 8, 20, 1, 2, 260, 262, 3, 2, 2, 2, 261, 253, 3, 2,
	2, 2, 261, 258, 3, 2, 2, 2, 262, 275, 3, 2, 2, 2, 263, 264, 12, 5, 2, 2,
	264, 265, 9, 5, 2, 2, 265, 266, 5, 38, 20, 6, 266, 267, 8, 20, 1, 2, 267,
	274, 3, 2, 2, 2, 268, 269, 12, 4, 2, 2, 269, 270, 9, 6, 2, 2, 270, 271,
	5, 38, 20, 5, 271, 272, 8, 20, 1, 2, 272, 274, 3, 2, 2, 2, 273, 263, 3,
	2, 2, 2, 273, 268, 3, 2, 2, 2, 274, 277, 3, 2, 2, 2, 275, 273, 3, 2, 2,
	2, 275, 276, 3, 2, 2, 2, 276, 39, 3, 2, 2, 2, 277, 275, 3, 2, 2, 2, 278,
	279, 7, 26, 2, 2, 279, 293, 8, 21, 1, 2, 280, 281, 7, 25, 2, 2, 281, 293,
	8, 21, 1, 2, 282, 283, 7, 30, 2, 2, 283, 293, 8, 21, 1, 2, 284, 285, 7,
	27, 2, 2, 285, 293, 8, 21, 1, 2, 286, 287, 7, 28, 2, 2, 287, 293, 8, 21,
	1, 2, 288, 289, 7, 18, 2, 2, 289, 293, 8, 21, 1, 2, 290, 291, 7, 19, 2,
	2, 291, 293, 8, 21, 1, 2, 292, 278, 3, 2, 2, 2, 292, 280, 3, 2, 2, 2, 292,
	282, 3, 2, 2, 2, 292, 284, 3, 2, 2, 2, 292, 286, 3, 2, 2, 2, 292, 288,
	3, 2, 2, 2, 292, 290, 3, 2, 2, 2, 293, 41, 3, 2, 2, 2, 18, 48, 77, 90,
	115, 141, 165, 174, 192, 210, 221, 237, 250, 261, 273, 275, 292,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"'H'", "'#include <stdio.h>'", "'double HEAP[100000];'", "'double STACK[100000];'",
	"'double SP;'", "'double HP;'", "'void'", "'goto'", "'return;'", "'HEAP'",
	"'STACK'", "'(int)'", "'(char)'", "'printf'", "'main'", "'return 0;'",
	"'SP'", "'HP'", "'if'", "'double'", "'int'", "", "", "", "", "", "", "",
	"", "'+'", "'-'", "'*'", "'/'", "'%'", "'<='", "'>='", "'=='", "'!='",
	"'<'", "'>'", "'{'", "'}'", "'['", "']'", "'('", "')'", "';'", "':'", "','",
	"'='",
}
var symbolicNames = []string{
	"'H'", "TK_STDIOH", "TK_DHEAP", "TK_DSTACK", "TK_PSTAKC", "TK_PHEAP", "TK_VOID",
	"TK_GOTO", "TK_RETURN", "TK_HEAP", "TK_STACK", "TK_CASTINT", "TK_CASTCHAR",
	"TK_PRINTF", "TK_MAIN", "TK_RETMAIN", "TK_PUNSTACK", "TK_PUNHEAP", "TK_IF",
	"TK_DOUBLE", "TK_INT", "TK_ASIGPSTACK", "TK_ASIGHEAP", "TK_FLOAT", "TK_ENTERO",
	"TK_CADENA", "TK_TEMPORAL", "TK_ETIQUETA", "TK_IDENTIFICADOR", "TK_SUMA",
	"TK_RESTA", "TK_MULTI", "TK_DIVI", "TK_MODULO", "TK_MENORIGUAL", "TK_MAYORIGUAL",
	"TK_IGUALDAD", "TK_DESIGUALDAD", "TK_MENOR", "TK_MAYOR", "TK_LI", "TK_LD",
	"TK_CI", "TK_CD", "TK_PI", "TK_PD", "TK_PYC", "TK_DP", "TK_COMA", "TK_IGUAL",
	"COMENTARIO_MUL", "COMENTARIO_LIN", "WHITESPACE",
}

var ruleNames = []string{
	"start", "instrucss", "instrs", "encabezado", "l_temporales", "l_funcas",
	"funcas", "sent_if", "imprimir", "casteo", "declaracion", "etiquetas",
	"llamada", "l_bloque", "bloque", "expression", "asignaciones", "expre_relacional",
	"expre_aritmetica", "datos",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type optimizador_parser struct {
	*antlr.BaseParser
}

func Newoptimizador_parser(input antlr.TokenStream) *optimizador_parser {
	this := new(optimizador_parser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "optimizador_parser.g4"

	return this
}

// optimizador_parser tokens.
const (
	optimizador_parserEOF              = antlr.TokenEOF
	optimizador_parserTK_STDIOH        = 1
	optimizador_parserTK_DHEAP         = 2
	optimizador_parserTK_DSTACK        = 3
	optimizador_parserTK_PSTAKC        = 4
	optimizador_parserTK_PHEAP         = 5
	optimizador_parserTK_VOID          = 6
	optimizador_parserTK_GOTO          = 7
	optimizador_parserTK_RETURN        = 8
	optimizador_parserTK_HEAP          = 9
	optimizador_parserTK_STACK         = 10
	optimizador_parserTK_CASTINT       = 11
	optimizador_parserTK_CASTCHAR      = 12
	optimizador_parserTK_PRINTF        = 13
	optimizador_parserTK_MAIN          = 14
	optimizador_parserTK_RETMAIN       = 15
	optimizador_parserTK_PUNSTACK      = 16
	optimizador_parserTK_PUNHEAP       = 17
	optimizador_parserTK_IF            = 18
	optimizador_parserTK_DOUBLE        = 19
	optimizador_parserTK_INT           = 20
	optimizador_parserTK_ASIGPSTACK    = 21
	optimizador_parserTK_ASIGHEAP      = 22
	optimizador_parserTK_FLOAT         = 23
	optimizador_parserTK_ENTERO        = 24
	optimizador_parserTK_CADENA        = 25
	optimizador_parserTK_TEMPORAL      = 26
	optimizador_parserTK_ETIQUETA      = 27
	optimizador_parserTK_IDENTIFICADOR = 28
	optimizador_parserTK_SUMA          = 29
	optimizador_parserTK_RESTA         = 30
	optimizador_parserTK_MULTI         = 31
	optimizador_parserTK_DIVI          = 32
	optimizador_parserTK_MODULO        = 33
	optimizador_parserTK_MENORIGUAL    = 34
	optimizador_parserTK_MAYORIGUAL    = 35
	optimizador_parserTK_IGUALDAD      = 36
	optimizador_parserTK_DESIGUALDAD   = 37
	optimizador_parserTK_MENOR         = 38
	optimizador_parserTK_MAYOR         = 39
	optimizador_parserTK_LI            = 40
	optimizador_parserTK_LD            = 41
	optimizador_parserTK_CI            = 42
	optimizador_parserTK_CD            = 43
	optimizador_parserTK_PI            = 44
	optimizador_parserTK_PD            = 45
	optimizador_parserTK_PYC           = 46
	optimizador_parserTK_DP            = 47
	optimizador_parserTK_COMA          = 48
	optimizador_parserTK_IGUAL         = 49
	optimizador_parserCOMENTARIO_MUL   = 50
	optimizador_parserCOMENTARIO_LIN   = 51
	optimizador_parserWHITESPACE       = 52
)

// optimizador_parser rules.
const (
	optimizador_parserRULE_start            = 0
	optimizador_parserRULE_instrucss        = 1
	optimizador_parserRULE_instrs           = 2
	optimizador_parserRULE_encabezado       = 3
	optimizador_parserRULE_l_temporales     = 4
	optimizador_parserRULE_l_funcas         = 5
	optimizador_parserRULE_funcas           = 6
	optimizador_parserRULE_sent_if          = 7
	optimizador_parserRULE_imprimir         = 8
	optimizador_parserRULE_casteo           = 9
	optimizador_parserRULE_declaracion      = 10
	optimizador_parserRULE_etiquetas        = 11
	optimizador_parserRULE_llamada          = 12
	optimizador_parserRULE_l_bloque         = 13
	optimizador_parserRULE_bloque           = 14
	optimizador_parserRULE_expression       = 15
	optimizador_parserRULE_asignaciones     = 16
	optimizador_parserRULE_expre_relacional = 17
	optimizador_parserRULE_expre_aritmetica = 18
	optimizador_parserRULE_datos            = 19
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instrucss returns the _instrucss rule contexts.
	Get_instrucss() IInstrucssContext

	// Set_instrucss sets the _instrucss rule contexts.
	Set_instrucss(IInstrucssContext)

	// GetLista returns the lista attribute.
	GetLista() *arrayList.List

	// SetLista sets the lista attribute.
	SetLista(*arrayList.List)

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	lista      *arrayList.List
	_instrucss IInstrucssContext
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = optimizador_parserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = optimizador_parserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Get_instrucss() IInstrucssContext { return s._instrucss }

func (s *StartContext) Set_instrucss(v IInstrucssContext) { s._instrucss = v }

func (s *StartContext) GetLista() *arrayList.List { return s.lista }

func (s *StartContext) SetLista(v *arrayList.List) { s.lista = v }

func (s *StartContext) Instrucss() IInstrucssContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstrucssContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstrucssContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(optimizador_parserListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(optimizador_parserListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *optimizador_parser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, optimizador_parserRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(40)

		var _x = p.Instrucss()

		localctx.(*StartContext)._instrucss = _x
	}
	localctx.(*StartContext).lista = localctx.(*StartContext).Get_instrucss().GetLis()

	return localctx
}

// IInstrucssContext is an interface to support dynamic dispatch.
type IInstrucssContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instrs returns the _instrs rule contexts.
	Get_instrs() IInstrsContext

	// Set_instrs sets the _instrs rule contexts.
	Set_instrs(IInstrsContext)

	// GetE returns the e rule context list.
	GetE() []IInstrsContext

	// SetE sets the e rule context list.
	SetE([]IInstrsContext)

	// GetLis returns the lis attribute.
	GetLis() *arrayList.List

	// SetLis sets the lis attribute.
	SetLis(*arrayList.List)

	// IsInstrucssContext differentiates from other interfaces.
	IsInstrucssContext()
}

type InstrucssContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	lis     *arrayList.List
	_instrs IInstrsContext
	e       []IInstrsContext
}

func NewEmptyInstrucssContext() *InstrucssContext {
	var p = new(InstrucssContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = optimizador_parserRULE_instrucss
	return p
}

func (*InstrucssContext) IsInstrucssContext() {}

func NewInstrucssContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstrucssContext {
	var p = new(InstrucssContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = optimizador_parserRULE_instrucss

	return p
}

func (s *InstrucssContext) GetParser() antlr.Parser { return s.parser }

func (s *InstrucssContext) Get_instrs() IInstrsContext { return s._instrs }

func (s *InstrucssContext) Set_instrs(v IInstrsContext) { s._instrs = v }

func (s *InstrucssContext) GetE() []IInstrsContext { return s.e }

func (s *InstrucssContext) SetE(v []IInstrsContext) { s.e = v }

func (s *InstrucssContext) GetLis() *arrayList.List { return s.lis }

func (s *InstrucssContext) SetLis(v *arrayList.List) { s.lis = v }

func (s *InstrucssContext) AllInstrs() []IInstrsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstrsContext)(nil)).Elem())
	var tst = make([]IInstrsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstrsContext)
		}
	}

	return tst
}

func (s *InstrucssContext) Instrs(i int) IInstrsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstrsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstrsContext)
}

func (s *InstrucssContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstrucssContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstrucssContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(optimizador_parserListener); ok {
		listenerT.EnterInstrucss(s)
	}
}

func (s *InstrucssContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(optimizador_parserListener); ok {
		listenerT.ExitInstrucss(s)
	}
}

func (p *optimizador_parser) Instrucss() (localctx IInstrucssContext) {
	localctx = NewInstrucssContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, optimizador_parserRULE_instrucss)

	localctx.(*InstrucssContext).lis = arrayList.New()

	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == optimizador_parserTK_STDIOH {
		{
			p.SetState(43)

			var _x = p.Instrs()

			localctx.(*InstrucssContext)._instrs = _x
		}
		localctx.(*InstrucssContext).e = append(localctx.(*InstrucssContext).e, localctx.(*InstrucssContext)._instrs)

		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	listInt := localctx.(*InstrucssContext).GetE()
	for _, e := range listInt {
		localctx.(*InstrucssContext).lis.Add(e.GetInstr())
	}

	return localctx
}

// IInstrsContext is an interface to support dynamic dispatch.
type IInstrsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_encabezado returns the _encabezado rule contexts.
	Get_encabezado() IEncabezadoContext

	// Set_encabezado sets the _encabezado rule contexts.
	Set_encabezado(IEncabezadoContext)

	// GetInstr returns the instr attribute.
	GetInstr() inter.Instruccion

	// SetInstr sets the instr attribute.
	SetInstr(inter.Instruccion)

	// IsInstrsContext differentiates from other interfaces.
	IsInstrsContext()
}

type InstrsContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       inter.Instruccion
	_encabezado IEncabezadoContext
}

func NewEmptyInstrsContext() *InstrsContext {
	var p = new(InstrsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = optimizador_parserRULE_instrs
	return p
}

func (*InstrsContext) IsInstrsContext() {}

func NewInstrsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstrsContext {
	var p = new(InstrsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = optimizador_parserRULE_instrs

	return p
}

func (s *InstrsContext) GetParser() antlr.Parser { return s.parser }

func (s *InstrsContext) Get_encabezado() IEncabezadoContext { return s._encabezado }

func (s *InstrsContext) Set_encabezado(v IEncabezadoContext) { s._encabezado = v }

func (s *InstrsContext) GetInstr() inter.Instruccion { return s.instr }

func (s *InstrsContext) SetInstr(v inter.Instruccion) { s.instr = v }

func (s *InstrsContext) Encabezado() IEncabezadoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEncabezadoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEncabezadoContext)
}

func (s *InstrsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstrsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstrsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(optimizador_parserListener); ok {
		listenerT.EnterInstrs(s)
	}
}

func (s *InstrsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(optimizador_parserListener); ok {
		listenerT.ExitInstrs(s)
	}
}

func (p *optimizador_parser) Instrs() (localctx IInstrsContext) {
	localctx = NewInstrsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, optimizador_parserRULE_instrs)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(51)

		var _x = p.Encabezado()

		localctx.(*InstrsContext)._encabezado = _x
	}
	localctx.(*InstrsContext).instr = localctx.(*InstrsContext).Get_encabezado().GetInstr()

	return localctx
}

// IEncabezadoContext is an interface to support dynamic dispatch.
type IEncabezadoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_TK_STDIOH returns the _TK_STDIOH token.
	Get_TK_STDIOH() antlr.Token

	// Get_TK_DHEAP returns the _TK_DHEAP token.
	Get_TK_DHEAP() antlr.Token

	// Get_TK_DSTACK returns the _TK_DSTACK token.
	Get_TK_DSTACK() antlr.Token

	// Get_TK_PSTAKC returns the _TK_PSTAKC token.
	Get_TK_PSTAKC() antlr.Token

	// Get_TK_PHEAP returns the _TK_PHEAP token.
	Get_TK_PHEAP() antlr.Token

	// Set_TK_STDIOH sets the _TK_STDIOH token.
	Set_TK_STDIOH(antlr.Token)

	// Set_TK_DHEAP sets the _TK_DHEAP token.
	Set_TK_DHEAP(antlr.Token)

	// Set_TK_DSTACK sets the _TK_DSTACK token.
	Set_TK_DSTACK(antlr.Token)

	// Set_TK_PSTAKC sets the _TK_PSTAKC token.
	Set_TK_PSTAKC(antlr.Token)

	// Set_TK_PHEAP sets the _TK_PHEAP token.
	Set_TK_PHEAP(antlr.Token)

	// Get_l_temporales returns the _l_temporales rule contexts.
	Get_l_temporales() IL_temporalesContext

	// Get_l_funcas returns the _l_funcas rule contexts.
	Get_l_funcas() IL_funcasContext

	// Set_l_temporales sets the _l_temporales rule contexts.
	Set_l_temporales(IL_temporalesContext)

	// Set_l_funcas sets the _l_funcas rule contexts.
	Set_l_funcas(IL_funcasContext)

	// GetInstr returns the instr attribute.
	GetInstr() inter.Instruccion

	// SetInstr sets the instr attribute.
	SetInstr(inter.Instruccion)

	// IsEncabezadoContext differentiates from other interfaces.
	IsEncabezadoContext()
}

type EncabezadoContext struct {
	*antlr.BaseParserRuleContext
	parser        antlr.Parser
	instr         inter.Instruccion
	_TK_STDIOH    antlr.Token
	_TK_DHEAP     antlr.Token
	_TK_DSTACK    antlr.Token
	_TK_PSTAKC    antlr.Token
	_TK_PHEAP     antlr.Token
	_l_temporales IL_temporalesContext
	_l_funcas     IL_funcasContext
}

func NewEmptyEncabezadoContext() *EncabezadoContext {
	var p = new(EncabezadoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = optimizador_parserRULE_encabezado
	return p
}

func (*EncabezadoContext) IsEncabezadoContext() {}

func NewEncabezadoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EncabezadoContext {
	var p = new(EncabezadoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = optimizador_parserRULE_encabezado

	return p
}

func (s *EncabezadoContext) GetParser() antlr.Parser { return s.parser }

func (s *EncabezadoContext) Get_TK_STDIOH() antlr.Token { return s._TK_STDIOH }

func (s *EncabezadoContext) Get_TK_DHEAP() antlr.Token { return s._TK_DHEAP }

func (s *EncabezadoContext) Get_TK_DSTACK() antlr.Token { return s._TK_DSTACK }

func (s *EncabezadoContext) Get_TK_PSTAKC() antlr.Token { return s._TK_PSTAKC }

func (s *EncabezadoContext) Get_TK_PHEAP() antlr.Token { return s._TK_PHEAP }

func (s *EncabezadoContext) Set_TK_STDIOH(v antlr.Token) { s._TK_STDIOH = v }

func (s *EncabezadoContext) Set_TK_DHEAP(v antlr.Token) { s._TK_DHEAP = v }

func (s *EncabezadoContext) Set_TK_DSTACK(v antlr.Token) { s._TK_DSTACK = v }

func (s *EncabezadoContext) Set_TK_PSTAKC(v antlr.Token) { s._TK_PSTAKC = v }

func (s *EncabezadoContext) Set_TK_PHEAP(v antlr.Token) { s._TK_PHEAP = v }

func (s *EncabezadoContext) Get_l_temporales() IL_temporalesContext { return s._l_temporales }

func (s *EncabezadoContext) Get_l_funcas() IL_funcasContext { return s._l_funcas }

func (s *EncabezadoContext) Set_l_temporales(v IL_temporalesContext) { s._l_temporales = v }

func (s *EncabezadoContext) Set_l_funcas(v IL_funcasContext) { s._l_funcas = v }

func (s *EncabezadoContext) GetInstr() inter.Instruccion { return s.instr }

func (s *EncabezadoContext) SetInstr(v inter.Instruccion) { s.instr = v }

func (s *EncabezadoContext) TK_STDIOH() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_STDIOH, 0)
}

func (s *EncabezadoContext) TK_DHEAP() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_DHEAP, 0)
}

func (s *EncabezadoContext) TK_DSTACK() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_DSTACK, 0)
}

func (s *EncabezadoContext) TK_PSTAKC() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_PSTAKC, 0)
}

func (s *EncabezadoContext) TK_PHEAP() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_PHEAP, 0)
}

func (s *EncabezadoContext) TK_DOUBLE() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_DOUBLE, 0)
}

func (s *EncabezadoContext) L_temporales() IL_temporalesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IL_temporalesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IL_temporalesContext)
}

func (s *EncabezadoContext) TK_PYC() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_PYC, 0)
}

func (s *EncabezadoContext) L_funcas() IL_funcasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IL_funcasContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IL_funcasContext)
}

func (s *EncabezadoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EncabezadoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EncabezadoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(optimizador_parserListener); ok {
		listenerT.EnterEncabezado(s)
	}
}

func (s *EncabezadoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(optimizador_parserListener); ok {
		listenerT.ExitEncabezado(s)
	}
}

func (p *optimizador_parser) Encabezado() (localctx IEncabezadoContext) {
	localctx = NewEncabezadoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, optimizador_parserRULE_encabezado)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(54)

		var _m = p.Match(optimizador_parserTK_STDIOH)

		localctx.(*EncabezadoContext)._TK_STDIOH = _m
	}
	{
		p.SetState(55)

		var _m = p.Match(optimizador_parserTK_DHEAP)

		localctx.(*EncabezadoContext)._TK_DHEAP = _m
	}
	{
		p.SetState(56)

		var _m = p.Match(optimizador_parserTK_DSTACK)

		localctx.(*EncabezadoContext)._TK_DSTACK = _m
	}
	{
		p.SetState(57)

		var _m = p.Match(optimizador_parserTK_PSTAKC)

		localctx.(*EncabezadoContext)._TK_PSTAKC = _m
	}
	{
		p.SetState(58)

		var _m = p.Match(optimizador_parserTK_PHEAP)

		localctx.(*EncabezadoContext)._TK_PHEAP = _m
	}
	{
		p.SetState(59)
		p.Match(optimizador_parserTK_DOUBLE)
	}
	{
		p.SetState(60)

		var _x = p.l_temporales(0)

		localctx.(*EncabezadoContext)._l_temporales = _x
	}
	{
		p.SetState(61)
		p.Match(optimizador_parserTK_PYC)
	}
	{
		p.SetState(62)

		var _x = p.l_funcas(0)

		localctx.(*EncabezadoContext)._l_funcas = _x
	}

	localctx.(*EncabezadoContext).instr = encabezado.Nenca((func() string {
		if localctx.(*EncabezadoContext).Get_TK_STDIOH() == nil {
			return ""
		} else {
			return localctx.(*EncabezadoContext).Get_TK_STDIOH().GetText()
		}
	}()), (func() string {
		if localctx.(*EncabezadoContext).Get_TK_DHEAP() == nil {
			return ""
		} else {
			return localctx.(*EncabezadoContext).Get_TK_DHEAP().GetText()
		}
	}()), (func() string {
		if localctx.(*EncabezadoContext).Get_TK_DSTACK() == nil {
			return ""
		} else {
			return localctx.(*EncabezadoContext).Get_TK_DSTACK().GetText()
		}
	}()), (func() string {
		if localctx.(*EncabezadoContext).Get_TK_PSTAKC() == nil {
			return ""
		} else {
			return localctx.(*EncabezadoContext).Get_TK_PSTAKC().GetText()
		}
	}()), (func() string {
		if localctx.(*EncabezadoContext).Get_TK_PHEAP() == nil {
			return ""
		} else {
			return localctx.(*EncabezadoContext).Get_TK_PHEAP().GetText()
		}
	}()), localctx.(*EncabezadoContext).Get_l_temporales().GetLtemps(), localctx.(*EncabezadoContext).Get_l_funcas().GetLfuncas())

	return localctx
}

// IL_temporalesContext is an interface to support dynamic dispatch.
type IL_temporalesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_TK_TEMPORAL returns the _TK_TEMPORAL token.
	Get_TK_TEMPORAL() antlr.Token

	// Set_TK_TEMPORAL sets the _TK_TEMPORAL token.
	Set_TK_TEMPORAL(antlr.Token)

	// GetTemps returns the temps rule contexts.
	GetTemps() IL_temporalesContext

	// SetTemps sets the temps rule contexts.
	SetTemps(IL_temporalesContext)

	// GetLtemps returns the ltemps attribute.
	GetLtemps() *arrayList.List

	// SetLtemps sets the ltemps attribute.
	SetLtemps(*arrayList.List)

	// IsL_temporalesContext differentiates from other interfaces.
	IsL_temporalesContext()
}

type L_temporalesContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	ltemps       *arrayList.List
	temps        IL_temporalesContext
	_TK_TEMPORAL antlr.Token
}

func NewEmptyL_temporalesContext() *L_temporalesContext {
	var p = new(L_temporalesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = optimizador_parserRULE_l_temporales
	return p
}

func (*L_temporalesContext) IsL_temporalesContext() {}

func NewL_temporalesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *L_temporalesContext {
	var p = new(L_temporalesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = optimizador_parserRULE_l_temporales

	return p
}

func (s *L_temporalesContext) GetParser() antlr.Parser { return s.parser }

func (s *L_temporalesContext) Get_TK_TEMPORAL() antlr.Token { return s._TK_TEMPORAL }

func (s *L_temporalesContext) Set_TK_TEMPORAL(v antlr.Token) { s._TK_TEMPORAL = v }

func (s *L_temporalesContext) GetTemps() IL_temporalesContext { return s.temps }

func (s *L_temporalesContext) SetTemps(v IL_temporalesContext) { s.temps = v }

func (s *L_temporalesContext) GetLtemps() *arrayList.List { return s.ltemps }

func (s *L_temporalesContext) SetLtemps(v *arrayList.List) { s.ltemps = v }

func (s *L_temporalesContext) TK_TEMPORAL() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_TEMPORAL, 0)
}

func (s *L_temporalesContext) TK_COMA() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_COMA, 0)
}

func (s *L_temporalesContext) L_temporales() IL_temporalesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IL_temporalesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IL_temporalesContext)
}

func (s *L_temporalesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *L_temporalesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *L_temporalesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(optimizador_parserListener); ok {
		listenerT.EnterL_temporales(s)
	}
}

func (s *L_temporalesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(optimizador_parserListener); ok {
		listenerT.ExitL_temporales(s)
	}
}

func (p *optimizador_parser) L_temporales() (localctx IL_temporalesContext) {
	return p.l_temporales(0)
}

func (p *optimizador_parser) l_temporales(_p int) (localctx IL_temporalesContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewL_temporalesContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IL_temporalesContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 8
	p.EnterRecursionRule(localctx, 8, optimizador_parserRULE_l_temporales, _p)

	localctx.(*L_temporalesContext).ltemps = arrayList.New()

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(66)

		var _m = p.Match(optimizador_parserTK_TEMPORAL)

		localctx.(*L_temporalesContext)._TK_TEMPORAL = _m
	}
	localctx.(*L_temporalesContext).ltemps.Add((func() string {
		if localctx.(*L_temporalesContext).Get_TK_TEMPORAL() == nil {
			return ""
		} else {
			return localctx.(*L_temporalesContext).Get_TK_TEMPORAL().GetText()
		}
	}()))

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewL_temporalesContext(p, _parentctx, _parentState)
			localctx.(*L_temporalesContext).temps = _prevctx
			p.PushNewRecursionContext(localctx, _startState, optimizador_parserRULE_l_temporales)
			p.SetState(69)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(70)
				p.Match(optimizador_parserTK_COMA)
			}
			{
				p.SetState(71)

				var _m = p.Match(optimizador_parserTK_TEMPORAL)

				localctx.(*L_temporalesContext)._TK_TEMPORAL = _m
			}

			localctx.(*L_temporalesContext).GetTemps().GetLtemps().Add((func() string {
				if localctx.(*L_temporalesContext).Get_TK_TEMPORAL() == nil {
					return ""
				} else {
					return localctx.(*L_temporalesContext).Get_TK_TEMPORAL().GetText()
				}
			}()))
			localctx.(*L_temporalesContext).ltemps = localctx.(*L_temporalesContext).GetTemps().GetLtemps()

		}
		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}

	return localctx
}

// IL_funcasContext is an interface to support dynamic dispatch.
type IL_funcasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFuncs returns the funcs rule contexts.
	GetFuncs() IL_funcasContext

	// Get_funcas returns the _funcas rule contexts.
	Get_funcas() IFuncasContext

	// SetFuncs sets the funcs rule contexts.
	SetFuncs(IL_funcasContext)

	// Set_funcas sets the _funcas rule contexts.
	Set_funcas(IFuncasContext)

	// GetLfuncas returns the lfuncas attribute.
	GetLfuncas() *arrayList.List

	// SetLfuncas sets the lfuncas attribute.
	SetLfuncas(*arrayList.List)

	// IsL_funcasContext differentiates from other interfaces.
	IsL_funcasContext()
}

type L_funcasContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	lfuncas *arrayList.List
	funcs   IL_funcasContext
	_funcas IFuncasContext
}

func NewEmptyL_funcasContext() *L_funcasContext {
	var p = new(L_funcasContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = optimizador_parserRULE_l_funcas
	return p
}

func (*L_funcasContext) IsL_funcasContext() {}

func NewL_funcasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *L_funcasContext {
	var p = new(L_funcasContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = optimizador_parserRULE_l_funcas

	return p
}

func (s *L_funcasContext) GetParser() antlr.Parser { return s.parser }

func (s *L_funcasContext) GetFuncs() IL_funcasContext { return s.funcs }

func (s *L_funcasContext) Get_funcas() IFuncasContext { return s._funcas }

func (s *L_funcasContext) SetFuncs(v IL_funcasContext) { s.funcs = v }

func (s *L_funcasContext) Set_funcas(v IFuncasContext) { s._funcas = v }

func (s *L_funcasContext) GetLfuncas() *arrayList.List { return s.lfuncas }

func (s *L_funcasContext) SetLfuncas(v *arrayList.List) { s.lfuncas = v }

func (s *L_funcasContext) Funcas() IFuncasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncasContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncasContext)
}

func (s *L_funcasContext) L_funcas() IL_funcasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IL_funcasContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IL_funcasContext)
}

func (s *L_funcasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *L_funcasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *L_funcasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(optimizador_parserListener); ok {
		listenerT.EnterL_funcas(s)
	}
}

func (s *L_funcasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(optimizador_parserListener); ok {
		listenerT.ExitL_funcas(s)
	}
}

func (p *optimizador_parser) L_funcas() (localctx IL_funcasContext) {
	return p.l_funcas(0)
}

func (p *optimizador_parser) l_funcas(_p int) (localctx IL_funcasContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewL_funcasContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IL_funcasContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, optimizador_parserRULE_l_funcas, _p)

	localctx.(*L_funcasContext).lfuncas = arrayList.New()

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)

		var _x = p.Funcas()

		localctx.(*L_funcasContext)._funcas = _x
	}
	localctx.(*L_funcasContext).lfuncas.Add(localctx.(*L_funcasContext).Get_funcas().GetInstr())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewL_funcasContext(p, _parentctx, _parentState)
			localctx.(*L_funcasContext).funcs = _prevctx
			p.PushNewRecursionContext(localctx, _startState, optimizador_parserRULE_l_funcas)
			p.SetState(82)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(83)

				var _x = p.Funcas()

				localctx.(*L_funcasContext)._funcas = _x
			}

			localctx.(*L_funcasContext).GetFuncs().GetLfuncas().Add(localctx.(*L_funcasContext).Get_funcas().GetInstr())
			localctx.(*L_funcasContext).lfuncas = localctx.(*L_funcasContext).GetFuncs().GetLfuncas()

		}
		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// IFuncasContext is an interface to support dynamic dispatch.
type IFuncasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_TK_IDENTIFICADOR returns the _TK_IDENTIFICADOR token.
	Get_TK_IDENTIFICADOR() antlr.Token

	// Get_TK_MAIN returns the _TK_MAIN token.
	Get_TK_MAIN() antlr.Token

	// Set_TK_IDENTIFICADOR sets the _TK_IDENTIFICADOR token.
	Set_TK_IDENTIFICADOR(antlr.Token)

	// Set_TK_MAIN sets the _TK_MAIN token.
	Set_TK_MAIN(antlr.Token)

	// Get_l_bloque returns the _l_bloque rule contexts.
	Get_l_bloque() IL_bloqueContext

	// Set_l_bloque sets the _l_bloque rule contexts.
	Set_l_bloque(IL_bloqueContext)

	// GetInstr returns the instr attribute.
	GetInstr() inter.Instruccion

	// SetInstr sets the instr attribute.
	SetInstr(inter.Instruccion)

	// IsFuncasContext differentiates from other interfaces.
	IsFuncasContext()
}

type FuncasContext struct {
	*antlr.BaseParserRuleContext
	parser            antlr.Parser
	instr             inter.Instruccion
	_TK_IDENTIFICADOR antlr.Token
	_l_bloque         IL_bloqueContext
	_TK_MAIN          antlr.Token
}

func NewEmptyFuncasContext() *FuncasContext {
	var p = new(FuncasContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = optimizador_parserRULE_funcas
	return p
}

func (*FuncasContext) IsFuncasContext() {}

func NewFuncasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncasContext {
	var p = new(FuncasContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = optimizador_parserRULE_funcas

	return p
}

func (s *FuncasContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncasContext) Get_TK_IDENTIFICADOR() antlr.Token { return s._TK_IDENTIFICADOR }

func (s *FuncasContext) Get_TK_MAIN() antlr.Token { return s._TK_MAIN }

func (s *FuncasContext) Set_TK_IDENTIFICADOR(v antlr.Token) { s._TK_IDENTIFICADOR = v }

func (s *FuncasContext) Set_TK_MAIN(v antlr.Token) { s._TK_MAIN = v }

func (s *FuncasContext) Get_l_bloque() IL_bloqueContext { return s._l_bloque }

func (s *FuncasContext) Set_l_bloque(v IL_bloqueContext) { s._l_bloque = v }

func (s *FuncasContext) GetInstr() inter.Instruccion { return s.instr }

func (s *FuncasContext) SetInstr(v inter.Instruccion) { s.instr = v }

func (s *FuncasContext) TK_VOID() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_VOID, 0)
}

func (s *FuncasContext) TK_IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_IDENTIFICADOR, 0)
}

func (s *FuncasContext) TK_PI() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_PI, 0)
}

func (s *FuncasContext) TK_PD() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_PD, 0)
}

func (s *FuncasContext) TK_LI() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_LI, 0)
}

func (s *FuncasContext) L_bloque() IL_bloqueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IL_bloqueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IL_bloqueContext)
}

func (s *FuncasContext) TK_RETURN() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_RETURN, 0)
}

func (s *FuncasContext) TK_LD() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_LD, 0)
}

func (s *FuncasContext) TK_INT() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_INT, 0)
}

func (s *FuncasContext) TK_MAIN() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_MAIN, 0)
}

func (s *FuncasContext) TK_ASIGPSTACK() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_ASIGPSTACK, 0)
}

func (s *FuncasContext) TK_ASIGHEAP() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_ASIGHEAP, 0)
}

func (s *FuncasContext) TK_RETMAIN() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_RETMAIN, 0)
}

func (s *FuncasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(optimizador_parserListener); ok {
		listenerT.EnterFuncas(s)
	}
}

func (s *FuncasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(optimizador_parserListener); ok {
		listenerT.ExitFuncas(s)
	}
}

func (p *optimizador_parser) Funcas() (localctx IFuncasContext) {
	localctx = NewFuncasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, optimizador_parserRULE_funcas)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(113)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case optimizador_parserTK_VOID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(91)
			p.Match(optimizador_parserTK_VOID)
		}
		{
			p.SetState(92)

			var _m = p.Match(optimizador_parserTK_IDENTIFICADOR)

			localctx.(*FuncasContext)._TK_IDENTIFICADOR = _m
		}
		{
			p.SetState(93)
			p.Match(optimizador_parserTK_PI)
		}
		{
			p.SetState(94)
			p.Match(optimizador_parserTK_PD)
		}
		{
			p.SetState(95)
			p.Match(optimizador_parserTK_LI)
		}
		{
			p.SetState(96)

			var _x = p.l_bloque(0)

			localctx.(*FuncasContext)._l_bloque = _x
		}
		{
			p.SetState(97)
			p.Match(optimizador_parserTK_RETURN)
		}
		{
			p.SetState(98)
			p.Match(optimizador_parserTK_LD)
		}

		localctx.(*FuncasContext).instr = funcion.Nfunca((func() string {
			if localctx.(*FuncasContext).Get_TK_IDENTIFICADOR() == nil {
				return ""
			} else {
				return localctx.(*FuncasContext).Get_TK_IDENTIFICADOR().GetText()
			}
		}()), localctx.(*FuncasContext).Get_l_bloque().GetLbloque())

	case optimizador_parserTK_INT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(101)
			p.Match(optimizador_parserTK_INT)
		}
		{
			p.SetState(102)

			var _m = p.Match(optimizador_parserTK_MAIN)

			localctx.(*FuncasContext)._TK_MAIN = _m
		}
		{
			p.SetState(103)
			p.Match(optimizador_parserTK_PI)
		}
		{
			p.SetState(104)
			p.Match(optimizador_parserTK_PD)
		}
		{
			p.SetState(105)
			p.Match(optimizador_parserTK_LI)
		}
		{
			p.SetState(106)
			p.Match(optimizador_parserTK_ASIGPSTACK)
		}
		{
			p.SetState(107)
			p.Match(optimizador_parserTK_ASIGHEAP)
		}
		{
			p.SetState(108)

			var _x = p.l_bloque(0)

			localctx.(*FuncasContext)._l_bloque = _x
		}
		{
			p.SetState(109)
			p.Match(optimizador_parserTK_RETMAIN)
		}
		{
			p.SetState(110)
			p.Match(optimizador_parserTK_LD)
		}

		localctx.(*FuncasContext).instr = funcion.Nfuncamain((func() string {
			if localctx.(*FuncasContext).Get_TK_MAIN() == nil {
				return ""
			} else {
				return localctx.(*FuncasContext).Get_TK_MAIN().GetText()
			}
		}()), localctx.(*FuncasContext).Get_l_bloque().GetLbloque())

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISent_ifContext is an interface to support dynamic dispatch.
type ISent_ifContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_TK_IF returns the _TK_IF token.
	Get_TK_IF() antlr.Token

	// Get_TK_ETIQUETA returns the _TK_ETIQUETA token.
	Get_TK_ETIQUETA() antlr.Token

	// Set_TK_IF sets the _TK_IF token.
	Set_TK_IF(antlr.Token)

	// Set_TK_ETIQUETA sets the _TK_ETIQUETA token.
	Set_TK_ETIQUETA(antlr.Token)

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetInstr returns the instr attribute.
	GetInstr() inter.Instruccion

	// SetInstr sets the instr attribute.
	SetInstr(inter.Instruccion)

	// IsSent_ifContext differentiates from other interfaces.
	IsSent_ifContext()
}

type Sent_ifContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	instr        inter.Instruccion
	_TK_IF       antlr.Token
	_expression  IExpressionContext
	_TK_ETIQUETA antlr.Token
}

func NewEmptySent_ifContext() *Sent_ifContext {
	var p = new(Sent_ifContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = optimizador_parserRULE_sent_if
	return p
}

func (*Sent_ifContext) IsSent_ifContext() {}

func NewSent_ifContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sent_ifContext {
	var p = new(Sent_ifContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = optimizador_parserRULE_sent_if

	return p
}

func (s *Sent_ifContext) GetParser() antlr.Parser { return s.parser }

func (s *Sent_ifContext) Get_TK_IF() antlr.Token { return s._TK_IF }

func (s *Sent_ifContext) Get_TK_ETIQUETA() antlr.Token { return s._TK_ETIQUETA }

func (s *Sent_ifContext) Set_TK_IF(v antlr.Token) { s._TK_IF = v }

func (s *Sent_ifContext) Set_TK_ETIQUETA(v antlr.Token) { s._TK_ETIQUETA = v }

func (s *Sent_ifContext) Get_expression() IExpressionContext { return s._expression }

func (s *Sent_ifContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Sent_ifContext) GetInstr() inter.Instruccion { return s.instr }

func (s *Sent_ifContext) SetInstr(v inter.Instruccion) { s.instr = v }

func (s *Sent_ifContext) TK_IF() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_IF, 0)
}

func (s *Sent_ifContext) TK_PI() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_PI, 0)
}

func (s *Sent_ifContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Sent_ifContext) TK_PD() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_PD, 0)
}

func (s *Sent_ifContext) TK_GOTO() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_GOTO, 0)
}

func (s *Sent_ifContext) TK_ETIQUETA() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_ETIQUETA, 0)
}

func (s *Sent_ifContext) TK_PYC() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_PYC, 0)
}

func (s *Sent_ifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sent_ifContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sent_ifContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(optimizador_parserListener); ok {
		listenerT.EnterSent_if(s)
	}
}

func (s *Sent_ifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(optimizador_parserListener); ok {
		listenerT.ExitSent_if(s)
	}
}

func (p *optimizador_parser) Sent_if() (localctx ISent_ifContext) {
	localctx = NewSent_ifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, optimizador_parserRULE_sent_if)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)

		var _m = p.Match(optimizador_parserTK_IF)

		localctx.(*Sent_ifContext)._TK_IF = _m
	}
	{
		p.SetState(116)
		p.Match(optimizador_parserTK_PI)
	}
	{
		p.SetState(117)

		var _x = p.Expression()

		localctx.(*Sent_ifContext)._expression = _x
	}
	{
		p.SetState(118)
		p.Match(optimizador_parserTK_PD)
	}
	{
		p.SetState(119)
		p.Match(optimizador_parserTK_GOTO)
	}
	{
		p.SetState(120)

		var _m = p.Match(optimizador_parserTK_ETIQUETA)

		localctx.(*Sent_ifContext)._TK_ETIQUETA = _m
	}
	{
		p.SetState(121)
		p.Match(optimizador_parserTK_PYC)
	}

	localctx.(*Sent_ifContext).instr = sif.Nsenteif(localctx.(*Sent_ifContext).Get_expression().GetP(), (func() string {
		if localctx.(*Sent_ifContext).Get_TK_ETIQUETA() == nil {
			return ""
		} else {
			return localctx.(*Sent_ifContext).Get_TK_ETIQUETA().GetText()
		}
	}()), (func() int {
		if localctx.(*Sent_ifContext).Get_TK_IF() == nil {
			return 0
		} else {
			return localctx.(*Sent_ifContext).Get_TK_IF().GetLine()
		}
	}()))

	return localctx
}

// IImprimirContext is an interface to support dynamic dispatch.
type IImprimirContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_TK_PRINTF returns the _TK_PRINTF token.
	Get_TK_PRINTF() antlr.Token

	// Get_TK_CADENA returns the _TK_CADENA token.
	Get_TK_CADENA() antlr.Token

	// Set_TK_PRINTF sets the _TK_PRINTF token.
	Set_TK_PRINTF(antlr.Token)

	// Set_TK_CADENA sets the _TK_CADENA token.
	Set_TK_CADENA(antlr.Token)

	// Get_casteo returns the _casteo rule contexts.
	Get_casteo() ICasteoContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_casteo sets the _casteo rule contexts.
	Set_casteo(ICasteoContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetInstr returns the instr attribute.
	GetInstr() inter.Instruccion

	// SetInstr sets the instr attribute.
	SetInstr(inter.Instruccion)

	// IsImprimirContext differentiates from other interfaces.
	IsImprimirContext()
}

type ImprimirContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       inter.Instruccion
	_TK_PRINTF  antlr.Token
	_TK_CADENA  antlr.Token
	_casteo     ICasteoContext
	_expression IExpressionContext
}

func NewEmptyImprimirContext() *ImprimirContext {
	var p = new(ImprimirContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = optimizador_parserRULE_imprimir
	return p
}

func (*ImprimirContext) IsImprimirContext() {}

func NewImprimirContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImprimirContext {
	var p = new(ImprimirContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = optimizador_parserRULE_imprimir

	return p
}

func (s *ImprimirContext) GetParser() antlr.Parser { return s.parser }

func (s *ImprimirContext) Get_TK_PRINTF() antlr.Token { return s._TK_PRINTF }

func (s *ImprimirContext) Get_TK_CADENA() antlr.Token { return s._TK_CADENA }

func (s *ImprimirContext) Set_TK_PRINTF(v antlr.Token) { s._TK_PRINTF = v }

func (s *ImprimirContext) Set_TK_CADENA(v antlr.Token) { s._TK_CADENA = v }

func (s *ImprimirContext) Get_casteo() ICasteoContext { return s._casteo }

func (s *ImprimirContext) Get_expression() IExpressionContext { return s._expression }

func (s *ImprimirContext) Set_casteo(v ICasteoContext) { s._casteo = v }

func (s *ImprimirContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *ImprimirContext) GetInstr() inter.Instruccion { return s.instr }

func (s *ImprimirContext) SetInstr(v inter.Instruccion) { s.instr = v }

func (s *ImprimirContext) TK_PRINTF() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_PRINTF, 0)
}

func (s *ImprimirContext) TK_PI() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_PI, 0)
}

func (s *ImprimirContext) TK_CADENA() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_CADENA, 0)
}

func (s *ImprimirContext) TK_COMA() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_COMA, 0)
}

func (s *ImprimirContext) Casteo() ICasteoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICasteoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICasteoContext)
}

func (s *ImprimirContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ImprimirContext) TK_PD() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_PD, 0)
}

func (s *ImprimirContext) TK_PYC() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_PYC, 0)
}

func (s *ImprimirContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImprimirContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImprimirContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(optimizador_parserListener); ok {
		listenerT.EnterImprimir(s)
	}
}

func (s *ImprimirContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(optimizador_parserListener); ok {
		listenerT.ExitImprimir(s)
	}
}

func (p *optimizador_parser) Imprimir() (localctx IImprimirContext) {
	localctx = NewImprimirContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, optimizador_parserRULE_imprimir)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(124)

		var _m = p.Match(optimizador_parserTK_PRINTF)

		localctx.(*ImprimirContext)._TK_PRINTF = _m
	}
	{
		p.SetState(125)
		p.Match(optimizador_parserTK_PI)
	}
	{
		p.SetState(126)

		var _m = p.Match(optimizador_parserTK_CADENA)

		localctx.(*ImprimirContext)._TK_CADENA = _m
	}
	{
		p.SetState(127)
		p.Match(optimizador_parserTK_COMA)
	}
	{
		p.SetState(128)

		var _x = p.Casteo()

		localctx.(*ImprimirContext)._casteo = _x
	}
	{
		p.SetState(129)

		var _x = p.Expression()

		localctx.(*ImprimirContext)._expression = _x
	}
	{
		p.SetState(130)
		p.Match(optimizador_parserTK_PD)
	}
	{
		p.SetState(131)
		p.Match(optimizador_parserTK_PYC)
	}

	localctx.(*ImprimirContext).instr = imprimir.Nimprimir((func() string {
		if localctx.(*ImprimirContext).Get_TK_CADENA() == nil {
			return ""
		} else {
			return localctx.(*ImprimirContext).Get_TK_CADENA().GetText()
		}
	}()), localctx.(*ImprimirContext).Get_casteo().GetCast(), localctx.(*ImprimirContext).Get_expression().GetP(), (func() int {
		if localctx.(*ImprimirContext).Get_TK_PRINTF() == nil {
			return 0
		} else {
			return localctx.(*ImprimirContext).Get_TK_PRINTF().GetLine()
		}
	}()))

	return localctx
}

// ICasteoContext is an interface to support dynamic dispatch.
type ICasteoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_TK_CASTCHAR returns the _TK_CASTCHAR token.
	Get_TK_CASTCHAR() antlr.Token

	// Get_TK_CASTINT returns the _TK_CASTINT token.
	Get_TK_CASTINT() antlr.Token

	// Set_TK_CASTCHAR sets the _TK_CASTCHAR token.
	Set_TK_CASTCHAR(antlr.Token)

	// Set_TK_CASTINT sets the _TK_CASTINT token.
	Set_TK_CASTINT(antlr.Token)

	// GetCast returns the cast attribute.
	GetCast() string

	// SetCast sets the cast attribute.
	SetCast(string)

	// IsCasteoContext differentiates from other interfaces.
	IsCasteoContext()
}

type CasteoContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	cast         string
	_TK_CASTCHAR antlr.Token
	_TK_CASTINT  antlr.Token
}

func NewEmptyCasteoContext() *CasteoContext {
	var p = new(CasteoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = optimizador_parserRULE_casteo
	return p
}

func (*CasteoContext) IsCasteoContext() {}

func NewCasteoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CasteoContext {
	var p = new(CasteoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = optimizador_parserRULE_casteo

	return p
}

func (s *CasteoContext) GetParser() antlr.Parser { return s.parser }

func (s *CasteoContext) Get_TK_CASTCHAR() antlr.Token { return s._TK_CASTCHAR }

func (s *CasteoContext) Get_TK_CASTINT() antlr.Token { return s._TK_CASTINT }

func (s *CasteoContext) Set_TK_CASTCHAR(v antlr.Token) { s._TK_CASTCHAR = v }

func (s *CasteoContext) Set_TK_CASTINT(v antlr.Token) { s._TK_CASTINT = v }

func (s *CasteoContext) GetCast() string { return s.cast }

func (s *CasteoContext) SetCast(v string) { s.cast = v }

func (s *CasteoContext) TK_CASTCHAR() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_CASTCHAR, 0)
}

func (s *CasteoContext) TK_CASTINT() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_CASTINT, 0)
}

func (s *CasteoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CasteoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CasteoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(optimizador_parserListener); ok {
		listenerT.EnterCasteo(s)
	}
}

func (s *CasteoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(optimizador_parserListener); ok {
		listenerT.ExitCasteo(s)
	}
}

func (p *optimizador_parser) Casteo() (localctx ICasteoContext) {
	localctx = NewCasteoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, optimizador_parserRULE_casteo)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(139)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case optimizador_parserTK_CASTCHAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(134)

			var _m = p.Match(optimizador_parserTK_CASTCHAR)

			localctx.(*CasteoContext)._TK_CASTCHAR = _m
		}
		localctx.(*CasteoContext).cast = (func() string {
			if localctx.(*CasteoContext).Get_TK_CASTCHAR() == nil {
				return ""
			} else {
				return localctx.(*CasteoContext).Get_TK_CASTCHAR().GetText()
			}
		}())

	case optimizador_parserTK_CASTINT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(136)

			var _m = p.Match(optimizador_parserTK_CASTINT)

			localctx.(*CasteoContext)._TK_CASTINT = _m
		}
		localctx.(*CasteoContext).cast = (func() string {
			if localctx.(*CasteoContext).Get_TK_CASTINT() == nil {
				return ""
			} else {
				return localctx.(*CasteoContext).Get_TK_CASTINT().GetText()
			}
		}())

	case optimizador_parserTK_HEAP, optimizador_parserTK_STACK, optimizador_parserTK_PUNSTACK, optimizador_parserTK_PUNHEAP, optimizador_parserTK_FLOAT, optimizador_parserTK_ENTERO, optimizador_parserTK_CADENA, optimizador_parserTK_TEMPORAL, optimizador_parserTK_IDENTIFICADOR, optimizador_parserTK_RESTA:
		p.EnterOuterAlt(localctx, 3)
		localctx.(*CasteoContext).cast = ""

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDeclaracionContext is an interface to support dynamic dispatch.
type IDeclaracionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_TK_TEMPORAL returns the _TK_TEMPORAL token.
	Get_TK_TEMPORAL() antlr.Token

	// GetNom returns the nom token.
	GetNom() antlr.Token

	// Set_TK_TEMPORAL sets the _TK_TEMPORAL token.
	Set_TK_TEMPORAL(antlr.Token)

	// SetNom sets the nom token.
	SetNom(antlr.Token)

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// GetPos returns the pos rule contexts.
	GetPos() IExpressionContext

	// GetExp returns the exp rule contexts.
	GetExp() IExpressionContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// SetPos sets the pos rule contexts.
	SetPos(IExpressionContext)

	// SetExp sets the exp rule contexts.
	SetExp(IExpressionContext)

	// GetInstr returns the instr attribute.
	GetInstr() inter.Instruccion

	// SetInstr sets the instr attribute.
	SetInstr(inter.Instruccion)

	// IsDeclaracionContext differentiates from other interfaces.
	IsDeclaracionContext()
}

type DeclaracionContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	instr        inter.Instruccion
	_TK_TEMPORAL antlr.Token
	_expression  IExpressionContext
	nom          antlr.Token
	pos          IExpressionContext
	exp          IExpressionContext
}

func NewEmptyDeclaracionContext() *DeclaracionContext {
	var p = new(DeclaracionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = optimizador_parserRULE_declaracion
	return p
}

func (*DeclaracionContext) IsDeclaracionContext() {}

func NewDeclaracionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclaracionContext {
	var p = new(DeclaracionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = optimizador_parserRULE_declaracion

	return p
}

func (s *DeclaracionContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclaracionContext) Get_TK_TEMPORAL() antlr.Token { return s._TK_TEMPORAL }

func (s *DeclaracionContext) GetNom() antlr.Token { return s.nom }

func (s *DeclaracionContext) Set_TK_TEMPORAL(v antlr.Token) { s._TK_TEMPORAL = v }

func (s *DeclaracionContext) SetNom(v antlr.Token) { s.nom = v }

func (s *DeclaracionContext) Get_expression() IExpressionContext { return s._expression }

func (s *DeclaracionContext) GetPos() IExpressionContext { return s.pos }

func (s *DeclaracionContext) GetExp() IExpressionContext { return s.exp }

func (s *DeclaracionContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *DeclaracionContext) SetPos(v IExpressionContext) { s.pos = v }

func (s *DeclaracionContext) SetExp(v IExpressionContext) { s.exp = v }

func (s *DeclaracionContext) GetInstr() inter.Instruccion { return s.instr }

func (s *DeclaracionContext) SetInstr(v inter.Instruccion) { s.instr = v }

func (s *DeclaracionContext) TK_TEMPORAL() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_TEMPORAL, 0)
}

func (s *DeclaracionContext) TK_IGUAL() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_IGUAL, 0)
}

func (s *DeclaracionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *DeclaracionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DeclaracionContext) TK_PYC() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_PYC, 0)
}

func (s *DeclaracionContext) TK_CI() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_CI, 0)
}

func (s *DeclaracionContext) TK_CASTINT() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_CASTINT, 0)
}

func (s *DeclaracionContext) TK_CD() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_CD, 0)
}

func (s *DeclaracionContext) TK_STACK() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_STACK, 0)
}

func (s *DeclaracionContext) TK_HEAP() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_HEAP, 0)
}

func (s *DeclaracionContext) TK_PUNSTACK() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_PUNSTACK, 0)
}

func (s *DeclaracionContext) TK_PUNHEAP() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_PUNHEAP, 0)
}

func (s *DeclaracionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclaracionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclaracionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(optimizador_parserListener); ok {
		listenerT.EnterDeclaracion(s)
	}
}

func (s *DeclaracionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(optimizador_parserListener); ok {
		listenerT.ExitDeclaracion(s)
	}
}

func (p *optimizador_parser) Declaracion() (localctx IDeclaracionContext) {
	localctx = NewDeclaracionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, optimizador_parserRULE_declaracion)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(163)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case optimizador_parserTK_TEMPORAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(141)

			var _m = p.Match(optimizador_parserTK_TEMPORAL)

			localctx.(*DeclaracionContext)._TK_TEMPORAL = _m
		}
		{
			p.SetState(142)
			p.Match(optimizador_parserTK_IGUAL)
		}
		{
			p.SetState(143)

			var _x = p.Expression()

			localctx.(*DeclaracionContext)._expression = _x
		}
		{
			p.SetState(144)
			p.Match(optimizador_parserTK_PYC)
		}

		localctx.(*DeclaracionContext).instr = declarar.Ndeclaracion((func() string {
			if localctx.(*DeclaracionContext).Get_TK_TEMPORAL() == nil {
				return ""
			} else {
				return localctx.(*DeclaracionContext).Get_TK_TEMPORAL().GetText()
			}
		}()), localctx.(*DeclaracionContext).Get_expression().GetP(), (func() int {
			if localctx.(*DeclaracionContext).Get_TK_TEMPORAL() == nil {
				return 0
			} else {
				return localctx.(*DeclaracionContext).Get_TK_TEMPORAL().GetLine()
			}
		}()))

	case optimizador_parserTK_HEAP, optimizador_parserTK_STACK:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(147)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*DeclaracionContext).nom = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == optimizador_parserTK_HEAP || _la == optimizador_parserTK_STACK) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*DeclaracionContext).nom = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(148)
			p.Match(optimizador_parserTK_CI)
		}
		{
			p.SetState(149)
			p.Match(optimizador_parserTK_CASTINT)
		}
		{
			p.SetState(150)

			var _x = p.Expression()

			localctx.(*DeclaracionContext).pos = _x
		}
		{
			p.SetState(151)
			p.Match(optimizador_parserTK_CD)
		}
		{
			p.SetState(152)
			p.Match(optimizador_parserTK_IGUAL)
		}
		{
			p.SetState(153)

			var _x = p.Expression()

			localctx.(*DeclaracionContext).exp = _x
		}
		{
			p.SetState(154)
			p.Match(optimizador_parserTK_PYC)
		}

		localctx.(*DeclaracionContext).instr = declarar.Ndeclapunteros((func() string {
			if localctx.(*DeclaracionContext).GetNom() == nil {
				return ""
			} else {
				return localctx.(*DeclaracionContext).GetNom().GetText()
			}
		}()), localctx.(*DeclaracionContext).GetPos().GetP(), localctx.(*DeclaracionContext).GetExp().GetP(), (func() int {
			if localctx.(*DeclaracionContext).GetNom() == nil {
				return 0
			} else {
				return localctx.(*DeclaracionContext).GetNom().GetLine()
			}
		}()))

	case optimizador_parserTK_PUNSTACK, optimizador_parserTK_PUNHEAP:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(157)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*DeclaracionContext).nom = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == optimizador_parserTK_PUNSTACK || _la == optimizador_parserTK_PUNHEAP) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*DeclaracionContext).nom = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(158)
			p.Match(optimizador_parserTK_IGUAL)
		}
		{
			p.SetState(159)

			var _x = p.Expression()

			localctx.(*DeclaracionContext)._expression = _x
		}
		{
			p.SetState(160)
			p.Match(optimizador_parserTK_PYC)
		}

		localctx.(*DeclaracionContext).instr = declarar.Ndeclapstack((func() string {
			if localctx.(*DeclaracionContext).GetNom() == nil {
				return ""
			} else {
				return localctx.(*DeclaracionContext).GetNom().GetText()
			}
		}()), localctx.(*DeclaracionContext).Get_expression().GetP(), (func() int {
			if localctx.(*DeclaracionContext).GetNom() == nil {
				return 0
			} else {
				return localctx.(*DeclaracionContext).GetNom().GetLine()
			}
		}()))

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IEtiquetasContext is an interface to support dynamic dispatch.
type IEtiquetasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_TK_ETIQUETA returns the _TK_ETIQUETA token.
	Get_TK_ETIQUETA() antlr.Token

	// Set_TK_ETIQUETA sets the _TK_ETIQUETA token.
	Set_TK_ETIQUETA(antlr.Token)

	// GetInstr returns the instr attribute.
	GetInstr() inter.Instruccion

	// SetInstr sets the instr attribute.
	SetInstr(inter.Instruccion)

	// IsEtiquetasContext differentiates from other interfaces.
	IsEtiquetasContext()
}

type EtiquetasContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	instr        inter.Instruccion
	_TK_ETIQUETA antlr.Token
}

func NewEmptyEtiquetasContext() *EtiquetasContext {
	var p = new(EtiquetasContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = optimizador_parserRULE_etiquetas
	return p
}

func (*EtiquetasContext) IsEtiquetasContext() {}

func NewEtiquetasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EtiquetasContext {
	var p = new(EtiquetasContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = optimizador_parserRULE_etiquetas

	return p
}

func (s *EtiquetasContext) GetParser() antlr.Parser { return s.parser }

func (s *EtiquetasContext) Get_TK_ETIQUETA() antlr.Token { return s._TK_ETIQUETA }

func (s *EtiquetasContext) Set_TK_ETIQUETA(v antlr.Token) { s._TK_ETIQUETA = v }

func (s *EtiquetasContext) GetInstr() inter.Instruccion { return s.instr }

func (s *EtiquetasContext) SetInstr(v inter.Instruccion) { s.instr = v }

func (s *EtiquetasContext) TK_ETIQUETA() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_ETIQUETA, 0)
}

func (s *EtiquetasContext) TK_DP() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_DP, 0)
}

func (s *EtiquetasContext) TK_GOTO() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_GOTO, 0)
}

func (s *EtiquetasContext) TK_PYC() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_PYC, 0)
}

func (s *EtiquetasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EtiquetasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EtiquetasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(optimizador_parserListener); ok {
		listenerT.EnterEtiquetas(s)
	}
}

func (s *EtiquetasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(optimizador_parserListener); ok {
		listenerT.ExitEtiquetas(s)
	}
}

func (p *optimizador_parser) Etiquetas() (localctx IEtiquetasContext) {
	localctx = NewEtiquetasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, optimizador_parserRULE_etiquetas)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(172)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case optimizador_parserTK_ETIQUETA:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(165)

			var _m = p.Match(optimizador_parserTK_ETIQUETA)

			localctx.(*EtiquetasContext)._TK_ETIQUETA = _m
		}
		{
			p.SetState(166)
			p.Match(optimizador_parserTK_DP)
		}

		localctx.(*EtiquetasContext).instr = etiquetas.Netiqueta((func() string {
			if localctx.(*EtiquetasContext).Get_TK_ETIQUETA() == nil {
				return ""
			} else {
				return localctx.(*EtiquetasContext).Get_TK_ETIQUETA().GetText()
			}
		}()))

	case optimizador_parserTK_GOTO:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(168)
			p.Match(optimizador_parserTK_GOTO)
		}
		{
			p.SetState(169)

			var _m = p.Match(optimizador_parserTK_ETIQUETA)

			localctx.(*EtiquetasContext)._TK_ETIQUETA = _m
		}
		{
			p.SetState(170)
			p.Match(optimizador_parserTK_PYC)
		}

		localctx.(*EtiquetasContext).instr = etiquetas.Ngoto((func() string {
			if localctx.(*EtiquetasContext).Get_TK_ETIQUETA() == nil {
				return ""
			} else {
				return localctx.(*EtiquetasContext).Get_TK_ETIQUETA().GetText()
			}
		}()))

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILlamadaContext is an interface to support dynamic dispatch.
type ILlamadaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_TK_IDENTIFICADOR returns the _TK_IDENTIFICADOR token.
	Get_TK_IDENTIFICADOR() antlr.Token

	// Set_TK_IDENTIFICADOR sets the _TK_IDENTIFICADOR token.
	Set_TK_IDENTIFICADOR(antlr.Token)

	// GetInstr returns the instr attribute.
	GetInstr() inter.Instruccion

	// SetInstr sets the instr attribute.
	SetInstr(inter.Instruccion)

	// IsLlamadaContext differentiates from other interfaces.
	IsLlamadaContext()
}

type LlamadaContext struct {
	*antlr.BaseParserRuleContext
	parser            antlr.Parser
	instr             inter.Instruccion
	_TK_IDENTIFICADOR antlr.Token
}

func NewEmptyLlamadaContext() *LlamadaContext {
	var p = new(LlamadaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = optimizador_parserRULE_llamada
	return p
}

func (*LlamadaContext) IsLlamadaContext() {}

func NewLlamadaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LlamadaContext {
	var p = new(LlamadaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = optimizador_parserRULE_llamada

	return p
}

func (s *LlamadaContext) GetParser() antlr.Parser { return s.parser }

func (s *LlamadaContext) Get_TK_IDENTIFICADOR() antlr.Token { return s._TK_IDENTIFICADOR }

func (s *LlamadaContext) Set_TK_IDENTIFICADOR(v antlr.Token) { s._TK_IDENTIFICADOR = v }

func (s *LlamadaContext) GetInstr() inter.Instruccion { return s.instr }

func (s *LlamadaContext) SetInstr(v inter.Instruccion) { s.instr = v }

func (s *LlamadaContext) TK_IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_IDENTIFICADOR, 0)
}

func (s *LlamadaContext) TK_PI() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_PI, 0)
}

func (s *LlamadaContext) TK_PD() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_PD, 0)
}

func (s *LlamadaContext) TK_PYC() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_PYC, 0)
}

func (s *LlamadaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LlamadaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LlamadaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(optimizador_parserListener); ok {
		listenerT.EnterLlamada(s)
	}
}

func (s *LlamadaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(optimizador_parserListener); ok {
		listenerT.ExitLlamada(s)
	}
}

func (p *optimizador_parser) Llamada() (localctx ILlamadaContext) {
	localctx = NewLlamadaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, optimizador_parserRULE_llamada)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(174)

		var _m = p.Match(optimizador_parserTK_IDENTIFICADOR)

		localctx.(*LlamadaContext)._TK_IDENTIFICADOR = _m
	}
	{
		p.SetState(175)
		p.Match(optimizador_parserTK_PI)
	}
	{
		p.SetState(176)
		p.Match(optimizador_parserTK_PD)
	}
	{
		p.SetState(177)
		p.Match(optimizador_parserTK_PYC)
	}

	localctx.(*LlamadaContext).instr = llamada.Nllama((func() string {
		if localctx.(*LlamadaContext).Get_TK_IDENTIFICADOR() == nil {
			return ""
		} else {
			return localctx.(*LlamadaContext).Get_TK_IDENTIFICADOR().GetText()
		}
	}()))

	return localctx
}

// IL_bloqueContext is an interface to support dynamic dispatch.
type IL_bloqueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetBlock returns the block rule contexts.
	GetBlock() IL_bloqueContext

	// Get_bloque returns the _bloque rule contexts.
	Get_bloque() IBloqueContext

	// SetBlock sets the block rule contexts.
	SetBlock(IL_bloqueContext)

	// Set_bloque sets the _bloque rule contexts.
	Set_bloque(IBloqueContext)

	// GetLbloque returns the lbloque attribute.
	GetLbloque() *arrayList.List

	// SetLbloque sets the lbloque attribute.
	SetLbloque(*arrayList.List)

	// IsL_bloqueContext differentiates from other interfaces.
	IsL_bloqueContext()
}

type L_bloqueContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	lbloque *arrayList.List
	block   IL_bloqueContext
	_bloque IBloqueContext
}

func NewEmptyL_bloqueContext() *L_bloqueContext {
	var p = new(L_bloqueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = optimizador_parserRULE_l_bloque
	return p
}

func (*L_bloqueContext) IsL_bloqueContext() {}

func NewL_bloqueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *L_bloqueContext {
	var p = new(L_bloqueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = optimizador_parserRULE_l_bloque

	return p
}

func (s *L_bloqueContext) GetParser() antlr.Parser { return s.parser }

func (s *L_bloqueContext) GetBlock() IL_bloqueContext { return s.block }

func (s *L_bloqueContext) Get_bloque() IBloqueContext { return s._bloque }

func (s *L_bloqueContext) SetBlock(v IL_bloqueContext) { s.block = v }

func (s *L_bloqueContext) Set_bloque(v IBloqueContext) { s._bloque = v }

func (s *L_bloqueContext) GetLbloque() *arrayList.List { return s.lbloque }

func (s *L_bloqueContext) SetLbloque(v *arrayList.List) { s.lbloque = v }

func (s *L_bloqueContext) Bloque() IBloqueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *L_bloqueContext) L_bloque() IL_bloqueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IL_bloqueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IL_bloqueContext)
}

func (s *L_bloqueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *L_bloqueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *L_bloqueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(optimizador_parserListener); ok {
		listenerT.EnterL_bloque(s)
	}
}

func (s *L_bloqueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(optimizador_parserListener); ok {
		listenerT.ExitL_bloque(s)
	}
}

func (p *optimizador_parser) L_bloque() (localctx IL_bloqueContext) {
	return p.l_bloque(0)
}

func (p *optimizador_parser) l_bloque(_p int) (localctx IL_bloqueContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewL_bloqueContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IL_bloqueContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 26
	p.EnterRecursionRule(localctx, 26, optimizador_parserRULE_l_bloque, _p)

	localctx.(*L_bloqueContext).lbloque = arrayList.New()

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(181)

		var _x = p.Bloque()

		localctx.(*L_bloqueContext)._bloque = _x
	}
	localctx.(*L_bloqueContext).lbloque.Add(localctx.(*L_bloqueContext).Get_bloque().GetInstr())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewL_bloqueContext(p, _parentctx, _parentState)
			localctx.(*L_bloqueContext).block = _prevctx
			p.PushNewRecursionContext(localctx, _startState, optimizador_parserRULE_l_bloque)
			p.SetState(184)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(185)

				var _x = p.Bloque()

				localctx.(*L_bloqueContext)._bloque = _x
			}

			localctx.(*L_bloqueContext).GetBlock().GetLbloque().Add(localctx.(*L_bloqueContext).Get_bloque().GetInstr())
			localctx.(*L_bloqueContext).lbloque = localctx.(*L_bloqueContext).GetBlock().GetLbloque()

		}
		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}

	return localctx
}

// IBloqueContext is an interface to support dynamic dispatch.
type IBloqueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_imprimir returns the _imprimir rule contexts.
	Get_imprimir() IImprimirContext

	// Get_sent_if returns the _sent_if rule contexts.
	Get_sent_if() ISent_ifContext

	// Get_declaracion returns the _declaracion rule contexts.
	Get_declaracion() IDeclaracionContext

	// Get_etiquetas returns the _etiquetas rule contexts.
	Get_etiquetas() IEtiquetasContext

	// Get_llamada returns the _llamada rule contexts.
	Get_llamada() ILlamadaContext

	// Set_imprimir sets the _imprimir rule contexts.
	Set_imprimir(IImprimirContext)

	// Set_sent_if sets the _sent_if rule contexts.
	Set_sent_if(ISent_ifContext)

	// Set_declaracion sets the _declaracion rule contexts.
	Set_declaracion(IDeclaracionContext)

	// Set_etiquetas sets the _etiquetas rule contexts.
	Set_etiquetas(IEtiquetasContext)

	// Set_llamada sets the _llamada rule contexts.
	Set_llamada(ILlamadaContext)

	// GetInstr returns the instr attribute.
	GetInstr() inter.Instruccion

	// SetInstr sets the instr attribute.
	SetInstr(inter.Instruccion)

	// IsBloqueContext differentiates from other interfaces.
	IsBloqueContext()
}

type BloqueContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	instr        inter.Instruccion
	_imprimir    IImprimirContext
	_sent_if     ISent_ifContext
	_declaracion IDeclaracionContext
	_etiquetas   IEtiquetasContext
	_llamada     ILlamadaContext
}

func NewEmptyBloqueContext() *BloqueContext {
	var p = new(BloqueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = optimizador_parserRULE_bloque
	return p
}

func (*BloqueContext) IsBloqueContext() {}

func NewBloqueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BloqueContext {
	var p = new(BloqueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = optimizador_parserRULE_bloque

	return p
}

func (s *BloqueContext) GetParser() antlr.Parser { return s.parser }

func (s *BloqueContext) Get_imprimir() IImprimirContext { return s._imprimir }

func (s *BloqueContext) Get_sent_if() ISent_ifContext { return s._sent_if }

func (s *BloqueContext) Get_declaracion() IDeclaracionContext { return s._declaracion }

func (s *BloqueContext) Get_etiquetas() IEtiquetasContext { return s._etiquetas }

func (s *BloqueContext) Get_llamada() ILlamadaContext { return s._llamada }

func (s *BloqueContext) Set_imprimir(v IImprimirContext) { s._imprimir = v }

func (s *BloqueContext) Set_sent_if(v ISent_ifContext) { s._sent_if = v }

func (s *BloqueContext) Set_declaracion(v IDeclaracionContext) { s._declaracion = v }

func (s *BloqueContext) Set_etiquetas(v IEtiquetasContext) { s._etiquetas = v }

func (s *BloqueContext) Set_llamada(v ILlamadaContext) { s._llamada = v }

func (s *BloqueContext) GetInstr() inter.Instruccion { return s.instr }

func (s *BloqueContext) SetInstr(v inter.Instruccion) { s.instr = v }

func (s *BloqueContext) Imprimir() IImprimirContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImprimirContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImprimirContext)
}

func (s *BloqueContext) Sent_if() ISent_ifContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISent_ifContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISent_ifContext)
}

func (s *BloqueContext) Declaracion() IDeclaracionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclaracionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclaracionContext)
}

func (s *BloqueContext) Etiquetas() IEtiquetasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEtiquetasContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEtiquetasContext)
}

func (s *BloqueContext) Llamada() ILlamadaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILlamadaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILlamadaContext)
}

func (s *BloqueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BloqueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BloqueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(optimizador_parserListener); ok {
		listenerT.EnterBloque(s)
	}
}

func (s *BloqueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(optimizador_parserListener); ok {
		listenerT.ExitBloque(s)
	}
}

func (p *optimizador_parser) Bloque() (localctx IBloqueContext) {
	localctx = NewBloqueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, optimizador_parserRULE_bloque)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(208)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case optimizador_parserTK_PRINTF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(193)

			var _x = p.Imprimir()

			localctx.(*BloqueContext)._imprimir = _x
		}
		localctx.(*BloqueContext).instr = localctx.(*BloqueContext).Get_imprimir().GetInstr()

	case optimizador_parserTK_IF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(196)

			var _x = p.Sent_if()

			localctx.(*BloqueContext)._sent_if = _x
		}
		localctx.(*BloqueContext).instr = localctx.(*BloqueContext).Get_sent_if().GetInstr()

	case optimizador_parserTK_HEAP, optimizador_parserTK_STACK, optimizador_parserTK_PUNSTACK, optimizador_parserTK_PUNHEAP, optimizador_parserTK_TEMPORAL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(199)

			var _x = p.Declaracion()

			localctx.(*BloqueContext)._declaracion = _x
		}
		localctx.(*BloqueContext).instr = localctx.(*BloqueContext).Get_declaracion().GetInstr()

	case optimizador_parserTK_GOTO, optimizador_parserTK_ETIQUETA:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(202)

			var _x = p.Etiquetas()

			localctx.(*BloqueContext)._etiquetas = _x
		}
		localctx.(*BloqueContext).instr = localctx.(*BloqueContext).Get_etiquetas().GetInstr()

	case optimizador_parserTK_IDENTIFICADOR:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(205)

			var _x = p.Llamada()

			localctx.(*BloqueContext)._llamada = _x
		}
		localctx.(*BloqueContext).instr = localctx.(*BloqueContext).Get_llamada().GetInstr()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expre_relacional returns the _expre_relacional rule contexts.
	Get_expre_relacional() IExpre_relacionalContext

	// Get_expre_aritmetica returns the _expre_aritmetica rule contexts.
	Get_expre_aritmetica() IExpre_aritmeticaContext

	// Get_asignaciones returns the _asignaciones rule contexts.
	Get_asignaciones() IAsignacionesContext

	// Set_expre_relacional sets the _expre_relacional rule contexts.
	Set_expre_relacional(IExpre_relacionalContext)

	// Set_expre_aritmetica sets the _expre_aritmetica rule contexts.
	Set_expre_aritmetica(IExpre_aritmeticaContext)

	// Set_asignaciones sets the _asignaciones rule contexts.
	Set_asignaciones(IAsignacionesContext)

	// GetP returns the p attribute.
	GetP() inter.Expresion

	// SetP sets the p attribute.
	SetP(inter.Expresion)

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser            antlr.Parser
	p                 inter.Expresion
	_expre_relacional IExpre_relacionalContext
	_expre_aritmetica IExpre_aritmeticaContext
	_asignaciones     IAsignacionesContext
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = optimizador_parserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = optimizador_parserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Get_expre_relacional() IExpre_relacionalContext {
	return s._expre_relacional
}

func (s *ExpressionContext) Get_expre_aritmetica() IExpre_aritmeticaContext {
	return s._expre_aritmetica
}

func (s *ExpressionContext) Get_asignaciones() IAsignacionesContext { return s._asignaciones }

func (s *ExpressionContext) Set_expre_relacional(v IExpre_relacionalContext) { s._expre_relacional = v }

func (s *ExpressionContext) Set_expre_aritmetica(v IExpre_aritmeticaContext) { s._expre_aritmetica = v }

func (s *ExpressionContext) Set_asignaciones(v IAsignacionesContext) { s._asignaciones = v }

func (s *ExpressionContext) GetP() inter.Expresion { return s.p }

func (s *ExpressionContext) SetP(v inter.Expresion) { s.p = v }

func (s *ExpressionContext) Expre_relacional() IExpre_relacionalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpre_relacionalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpre_relacionalContext)
}

func (s *ExpressionContext) Expre_aritmetica() IExpre_aritmeticaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpre_aritmeticaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpre_aritmeticaContext)
}

func (s *ExpressionContext) Asignaciones() IAsignacionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAsignacionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAsignacionesContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(optimizador_parserListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(optimizador_parserListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *optimizador_parser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, optimizador_parserRULE_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(210)

			var _x = p.expre_relacional(0)

			localctx.(*ExpressionContext)._expre_relacional = _x
		}
		localctx.(*ExpressionContext).p = localctx.(*ExpressionContext).Get_expre_relacional().GetP()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(213)

			var _x = p.expre_aritmetica(0)

			localctx.(*ExpressionContext)._expre_aritmetica = _x
		}
		localctx.(*ExpressionContext).p = localctx.(*ExpressionContext).Get_expre_aritmetica().GetP()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(216)

			var _x = p.Asignaciones()

			localctx.(*ExpressionContext)._asignaciones = _x
		}
		localctx.(*ExpressionContext).p = localctx.(*ExpressionContext).Get_asignaciones().GetP()

	}

	return localctx
}

// IAsignacionesContext is an interface to support dynamic dispatch.
type IAsignacionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetP returns the p attribute.
	GetP() inter.Expresion

	// SetP sets the p attribute.
	SetP(inter.Expresion)

	// IsAsignacionesContext differentiates from other interfaces.
	IsAsignacionesContext()
}

type AsignacionesContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	p           inter.Expresion
	_expression IExpressionContext
}

func NewEmptyAsignacionesContext() *AsignacionesContext {
	var p = new(AsignacionesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = optimizador_parserRULE_asignaciones
	return p
}

func (*AsignacionesContext) IsAsignacionesContext() {}

func NewAsignacionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsignacionesContext {
	var p = new(AsignacionesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = optimizador_parserRULE_asignaciones

	return p
}

func (s *AsignacionesContext) GetParser() antlr.Parser { return s.parser }

func (s *AsignacionesContext) Get_expression() IExpressionContext { return s._expression }

func (s *AsignacionesContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *AsignacionesContext) GetP() inter.Expresion { return s.p }

func (s *AsignacionesContext) SetP(v inter.Expresion) { s.p = v }

func (s *AsignacionesContext) TK_HEAP() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_HEAP, 0)
}

func (s *AsignacionesContext) TK_CI() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_CI, 0)
}

func (s *AsignacionesContext) TK_CASTINT() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_CASTINT, 0)
}

func (s *AsignacionesContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AsignacionesContext) TK_CD() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_CD, 0)
}

func (s *AsignacionesContext) TK_STACK() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_STACK, 0)
}

func (s *AsignacionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsignacionesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(optimizador_parserListener); ok {
		listenerT.EnterAsignaciones(s)
	}
}

func (s *AsignacionesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(optimizador_parserListener); ok {
		listenerT.ExitAsignaciones(s)
	}
}

func (p *optimizador_parser) Asignaciones() (localctx IAsignacionesContext) {
	localctx = NewAsignacionesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, optimizador_parserRULE_asignaciones)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(235)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case optimizador_parserTK_HEAP:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(221)
			p.Match(optimizador_parserTK_HEAP)
		}
		{
			p.SetState(222)
			p.Match(optimizador_parserTK_CI)
		}
		{
			p.SetState(223)
			p.Match(optimizador_parserTK_CASTINT)
		}
		{
			p.SetState(224)

			var _x = p.Expression()

			localctx.(*AsignacionesContext)._expression = _x
		}
		{
			p.SetState(225)
			p.Match(optimizador_parserTK_CD)
		}

		localctx.(*AsignacionesContext).p = asignaciones.Nasiheap(localctx.(*AsignacionesContext).Get_expression().GetP())

	case optimizador_parserTK_STACK:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(228)
			p.Match(optimizador_parserTK_STACK)
		}
		{
			p.SetState(229)
			p.Match(optimizador_parserTK_CI)
		}
		{
			p.SetState(230)
			p.Match(optimizador_parserTK_CASTINT)
		}
		{
			p.SetState(231)

			var _x = p.Expression()

			localctx.(*AsignacionesContext)._expression = _x
		}
		{
			p.SetState(232)
			p.Match(optimizador_parserTK_CD)
		}

		localctx.(*AsignacionesContext).p = asignaciones.Nasistack(localctx.(*AsignacionesContext).Get_expression().GetP())

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExpre_relacionalContext is an interface to support dynamic dispatch.
type IExpre_relacionalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetOpIz returns the opIz rule contexts.
	GetOpIz() IExpre_relacionalContext

	// Get_expre_aritmetica returns the _expre_aritmetica rule contexts.
	Get_expre_aritmetica() IExpre_aritmeticaContext

	// GetOpDe returns the opDe rule contexts.
	GetOpDe() IExpre_relacionalContext

	// SetOpIz sets the opIz rule contexts.
	SetOpIz(IExpre_relacionalContext)

	// Set_expre_aritmetica sets the _expre_aritmetica rule contexts.
	Set_expre_aritmetica(IExpre_aritmeticaContext)

	// SetOpDe sets the opDe rule contexts.
	SetOpDe(IExpre_relacionalContext)

	// GetP returns the p attribute.
	GetP() inter.Expresion

	// SetP sets the p attribute.
	SetP(inter.Expresion)

	// IsExpre_relacionalContext differentiates from other interfaces.
	IsExpre_relacionalContext()
}

type Expre_relacionalContext struct {
	*antlr.BaseParserRuleContext
	parser            antlr.Parser
	p                 inter.Expresion
	opIz              IExpre_relacionalContext
	_expre_aritmetica IExpre_aritmeticaContext
	op                antlr.Token
	opDe              IExpre_relacionalContext
}

func NewEmptyExpre_relacionalContext() *Expre_relacionalContext {
	var p = new(Expre_relacionalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = optimizador_parserRULE_expre_relacional
	return p
}

func (*Expre_relacionalContext) IsExpre_relacionalContext() {}

func NewExpre_relacionalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expre_relacionalContext {
	var p = new(Expre_relacionalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = optimizador_parserRULE_expre_relacional

	return p
}

func (s *Expre_relacionalContext) GetParser() antlr.Parser { return s.parser }

func (s *Expre_relacionalContext) GetOp() antlr.Token { return s.op }

func (s *Expre_relacionalContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expre_relacionalContext) GetOpIz() IExpre_relacionalContext { return s.opIz }

func (s *Expre_relacionalContext) Get_expre_aritmetica() IExpre_aritmeticaContext {
	return s._expre_aritmetica
}

func (s *Expre_relacionalContext) GetOpDe() IExpre_relacionalContext { return s.opDe }

func (s *Expre_relacionalContext) SetOpIz(v IExpre_relacionalContext) { s.opIz = v }

func (s *Expre_relacionalContext) Set_expre_aritmetica(v IExpre_aritmeticaContext) {
	s._expre_aritmetica = v
}

func (s *Expre_relacionalContext) SetOpDe(v IExpre_relacionalContext) { s.opDe = v }

func (s *Expre_relacionalContext) GetP() inter.Expresion { return s.p }

func (s *Expre_relacionalContext) SetP(v inter.Expresion) { s.p = v }

func (s *Expre_relacionalContext) Expre_aritmetica() IExpre_aritmeticaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpre_aritmeticaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpre_aritmeticaContext)
}

func (s *Expre_relacionalContext) AllExpre_relacional() []IExpre_relacionalContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpre_relacionalContext)(nil)).Elem())
	var tst = make([]IExpre_relacionalContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpre_relacionalContext)
		}
	}

	return tst
}

func (s *Expre_relacionalContext) Expre_relacional(i int) IExpre_relacionalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpre_relacionalContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpre_relacionalContext)
}

func (s *Expre_relacionalContext) TK_MAYORIGUAL() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_MAYORIGUAL, 0)
}

func (s *Expre_relacionalContext) TK_MENORIGUAL() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_MENORIGUAL, 0)
}

func (s *Expre_relacionalContext) TK_IGUALDAD() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_IGUALDAD, 0)
}

func (s *Expre_relacionalContext) TK_DESIGUALDAD() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_DESIGUALDAD, 0)
}

func (s *Expre_relacionalContext) TK_MAYOR() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_MAYOR, 0)
}

func (s *Expre_relacionalContext) TK_MENOR() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_MENOR, 0)
}

func (s *Expre_relacionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expre_relacionalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expre_relacionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(optimizador_parserListener); ok {
		listenerT.EnterExpre_relacional(s)
	}
}

func (s *Expre_relacionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(optimizador_parserListener); ok {
		listenerT.ExitExpre_relacional(s)
	}
}

func (p *optimizador_parser) Expre_relacional() (localctx IExpre_relacionalContext) {
	return p.expre_relacional(0)
}

func (p *optimizador_parser) expre_relacional(_p int) (localctx IExpre_relacionalContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpre_relacionalContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpre_relacionalContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 34
	p.EnterRecursionRule(localctx, 34, optimizador_parserRULE_expre_relacional, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(238)

		var _x = p.expre_aritmetica(0)

		localctx.(*Expre_relacionalContext)._expre_aritmetica = _x
	}
	localctx.(*Expre_relacionalContext).p = localctx.(*Expre_relacionalContext).Get_expre_aritmetica().GetP()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpre_relacionalContext(p, _parentctx, _parentState)
			localctx.(*Expre_relacionalContext).opIz = _prevctx
			p.PushNewRecursionContext(localctx, _startState, optimizador_parserRULE_expre_relacional)
			p.SetState(241)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(242)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*Expre_relacionalContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(optimizador_parserTK_MENORIGUAL-34))|(1<<(optimizador_parserTK_MAYORIGUAL-34))|(1<<(optimizador_parserTK_IGUALDAD-34))|(1<<(optimizador_parserTK_DESIGUALDAD-34))|(1<<(optimizador_parserTK_MENOR-34))|(1<<(optimizador_parserTK_MAYOR-34)))) != 0) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*Expre_relacionalContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(243)

				var _x = p.expre_relacional(3)

				localctx.(*Expre_relacionalContext).opDe = _x
			}

			if (func() string {
				if localctx.(*Expre_relacionalContext).GetOp() == nil {
					return ""
				} else {
					return localctx.(*Expre_relacionalContext).GetOp().GetText()
				}
			}()) == "<" {
				localctx.(*Expre_relacionalContext).p = menor.Noperacionmenor(localctx.(*Expre_relacionalContext).GetOpIz().GetP(), localctx.(*Expre_relacionalContext).GetOpDe().GetP())
			} else if (func() string {
				if localctx.(*Expre_relacionalContext).GetOp() == nil {
					return ""
				} else {
					return localctx.(*Expre_relacionalContext).GetOp().GetText()
				}
			}()) == "<=" {
				localctx.(*Expre_relacionalContext).p = menorigual.Noperacionmenorigual(localctx.(*Expre_relacionalContext).GetOpIz().GetP(), localctx.(*Expre_relacionalContext).GetOpDe().GetP())
			} else if (func() string {
				if localctx.(*Expre_relacionalContext).GetOp() == nil {
					return ""
				} else {
					return localctx.(*Expre_relacionalContext).GetOp().GetText()
				}
			}()) == ">" {
				localctx.(*Expre_relacionalContext).p = mayor.Noperacionmayor(localctx.(*Expre_relacionalContext).GetOpIz().GetP(), localctx.(*Expre_relacionalContext).GetOpDe().GetP())
			} else if (func() string {
				if localctx.(*Expre_relacionalContext).GetOp() == nil {
					return ""
				} else {
					return localctx.(*Expre_relacionalContext).GetOp().GetText()
				}
			}()) == ">=" {
				localctx.(*Expre_relacionalContext).p = mayorigual.Noperacionmayorigual(localctx.(*Expre_relacionalContext).GetOpIz().GetP(), localctx.(*Expre_relacionalContext).GetOpDe().GetP())
			} else if (func() string {
				if localctx.(*Expre_relacionalContext).GetOp() == nil {
					return ""
				} else {
					return localctx.(*Expre_relacionalContext).GetOp().GetText()
				}
			}()) == "==" {
				localctx.(*Expre_relacionalContext).p = igualdad.Noperacionigualdad(localctx.(*Expre_relacionalContext).GetOpIz().GetP(), localctx.(*Expre_relacionalContext).GetOpDe().GetP())
			} else if (func() string {
				if localctx.(*Expre_relacionalContext).GetOp() == nil {
					return ""
				} else {
					return localctx.(*Expre_relacionalContext).GetOp().GetText()
				}
			}()) == "!=" {
				localctx.(*Expre_relacionalContext).p = desigualdad.Noperaciondesigualdad(localctx.(*Expre_relacionalContext).GetOpIz().GetP(), localctx.(*Expre_relacionalContext).GetOpDe().GetP())
			}

		}
		p.SetState(250)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
	}

	return localctx
}

// IExpre_aritmeticaContext is an interface to support dynamic dispatch.
type IExpre_aritmeticaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOpera returns the opera token.
	GetOpera() antlr.Token

	// SetOpera sets the opera token.
	SetOpera(antlr.Token)

	// GetOpIz returns the opIz rule contexts.
	GetOpIz() IExpre_aritmeticaContext

	// GetOpUn returns the opUn rule contexts.
	GetOpUn() IExpre_aritmeticaContext

	// Get_datos returns the _datos rule contexts.
	Get_datos() IDatosContext

	// GetOpDe returns the opDe rule contexts.
	GetOpDe() IExpre_aritmeticaContext

	// SetOpIz sets the opIz rule contexts.
	SetOpIz(IExpre_aritmeticaContext)

	// SetOpUn sets the opUn rule contexts.
	SetOpUn(IExpre_aritmeticaContext)

	// Set_datos sets the _datos rule contexts.
	Set_datos(IDatosContext)

	// SetOpDe sets the opDe rule contexts.
	SetOpDe(IExpre_aritmeticaContext)

	// GetP returns the p attribute.
	GetP() inter.Expresion

	// SetP sets the p attribute.
	SetP(inter.Expresion)

	// IsExpre_aritmeticaContext differentiates from other interfaces.
	IsExpre_aritmeticaContext()
}

type Expre_aritmeticaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	p      inter.Expresion
	opIz   IExpre_aritmeticaContext
	opera  antlr.Token
	opUn   IExpre_aritmeticaContext
	_datos IDatosContext
	opDe   IExpre_aritmeticaContext
}

func NewEmptyExpre_aritmeticaContext() *Expre_aritmeticaContext {
	var p = new(Expre_aritmeticaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = optimizador_parserRULE_expre_aritmetica
	return p
}

func (*Expre_aritmeticaContext) IsExpre_aritmeticaContext() {}

func NewExpre_aritmeticaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expre_aritmeticaContext {
	var p = new(Expre_aritmeticaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = optimizador_parserRULE_expre_aritmetica

	return p
}

func (s *Expre_aritmeticaContext) GetParser() antlr.Parser { return s.parser }

func (s *Expre_aritmeticaContext) GetOpera() antlr.Token { return s.opera }

func (s *Expre_aritmeticaContext) SetOpera(v antlr.Token) { s.opera = v }

func (s *Expre_aritmeticaContext) GetOpIz() IExpre_aritmeticaContext { return s.opIz }

func (s *Expre_aritmeticaContext) GetOpUn() IExpre_aritmeticaContext { return s.opUn }

func (s *Expre_aritmeticaContext) Get_datos() IDatosContext { return s._datos }

func (s *Expre_aritmeticaContext) GetOpDe() IExpre_aritmeticaContext { return s.opDe }

func (s *Expre_aritmeticaContext) SetOpIz(v IExpre_aritmeticaContext) { s.opIz = v }

func (s *Expre_aritmeticaContext) SetOpUn(v IExpre_aritmeticaContext) { s.opUn = v }

func (s *Expre_aritmeticaContext) Set_datos(v IDatosContext) { s._datos = v }

func (s *Expre_aritmeticaContext) SetOpDe(v IExpre_aritmeticaContext) { s.opDe = v }

func (s *Expre_aritmeticaContext) GetP() inter.Expresion { return s.p }

func (s *Expre_aritmeticaContext) SetP(v inter.Expresion) { s.p = v }

func (s *Expre_aritmeticaContext) TK_RESTA() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_RESTA, 0)
}

func (s *Expre_aritmeticaContext) AllExpre_aritmetica() []IExpre_aritmeticaContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpre_aritmeticaContext)(nil)).Elem())
	var tst = make([]IExpre_aritmeticaContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpre_aritmeticaContext)
		}
	}

	return tst
}

func (s *Expre_aritmeticaContext) Expre_aritmetica(i int) IExpre_aritmeticaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpre_aritmeticaContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpre_aritmeticaContext)
}

func (s *Expre_aritmeticaContext) Datos() IDatosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDatosContext)
}

func (s *Expre_aritmeticaContext) TK_MULTI() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_MULTI, 0)
}

func (s *Expre_aritmeticaContext) TK_DIVI() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_DIVI, 0)
}

func (s *Expre_aritmeticaContext) TK_MODULO() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_MODULO, 0)
}

func (s *Expre_aritmeticaContext) TK_SUMA() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_SUMA, 0)
}

func (s *Expre_aritmeticaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expre_aritmeticaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expre_aritmeticaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(optimizador_parserListener); ok {
		listenerT.EnterExpre_aritmetica(s)
	}
}

func (s *Expre_aritmeticaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(optimizador_parserListener); ok {
		listenerT.ExitExpre_aritmetica(s)
	}
}

func (p *optimizador_parser) Expre_aritmetica() (localctx IExpre_aritmeticaContext) {
	return p.expre_aritmetica(0)
}

func (p *optimizador_parser) expre_aritmetica(_p int) (localctx IExpre_aritmeticaContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpre_aritmeticaContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpre_aritmeticaContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 36
	p.EnterRecursionRule(localctx, 36, optimizador_parserRULE_expre_aritmetica, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(259)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case optimizador_parserTK_RESTA:
		{
			p.SetState(252)

			var _m = p.Match(optimizador_parserTK_RESTA)

			localctx.(*Expre_aritmeticaContext).opera = _m
		}
		{
			p.SetState(253)

			var _x = p.expre_aritmetica(4)

			localctx.(*Expre_aritmeticaContext).opUn = _x
		}

		localctx.(*Expre_aritmeticaContext).p = negativo.Nnegativo(localctx.(*Expre_aritmeticaContext).GetOpUn().GetP())

	case optimizador_parserTK_PUNSTACK, optimizador_parserTK_PUNHEAP, optimizador_parserTK_FLOAT, optimizador_parserTK_ENTERO, optimizador_parserTK_CADENA, optimizador_parserTK_TEMPORAL, optimizador_parserTK_IDENTIFICADOR:
		{
			p.SetState(256)

			var _x = p.Datos()

			localctx.(*Expre_aritmeticaContext)._datos = _x
		}
		localctx.(*Expre_aritmeticaContext).p = localctx.(*Expre_aritmeticaContext).Get_datos().GetP()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(271)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpre_aritmeticaContext(p, _parentctx, _parentState)
				localctx.(*Expre_aritmeticaContext).opIz = _prevctx
				p.PushNewRecursionContext(localctx, _startState, optimizador_parserRULE_expre_aritmetica)
				p.SetState(261)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(262)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expre_aritmeticaContext).opera = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-31)&-(0x1f+1)) == 0 && ((1<<uint((_la-31)))&((1<<(optimizador_parserTK_MULTI-31))|(1<<(optimizador_parserTK_DIVI-31))|(1<<(optimizador_parserTK_MODULO-31)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expre_aritmeticaContext).opera = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(263)

					var _x = p.expre_aritmetica(4)

					localctx.(*Expre_aritmeticaContext).opDe = _x
				}

				if (func() string {
					if localctx.(*Expre_aritmeticaContext).GetOpera() == nil {
						return ""
					} else {
						return localctx.(*Expre_aritmeticaContext).GetOpera().GetText()
					}
				}()) == "*" {
					localctx.(*Expre_aritmeticaContext).p = multiplicacion.Noperacionmultiplicacion(localctx.(*Expre_aritmeticaContext).GetOpIz().GetP(), localctx.(*Expre_aritmeticaContext).GetOpDe().GetP())
				} else if (func() string {
					if localctx.(*Expre_aritmeticaContext).GetOpera() == nil {
						return ""
					} else {
						return localctx.(*Expre_aritmeticaContext).GetOpera().GetText()
					}
				}()) == "/" {
					localctx.(*Expre_aritmeticaContext).p = division.Noperaciondivision(localctx.(*Expre_aritmeticaContext).GetOpIz().GetP(), localctx.(*Expre_aritmeticaContext).GetOpDe().GetP())
				} else if (func() string {
					if localctx.(*Expre_aritmeticaContext).GetOpera() == nil {
						return ""
					} else {
						return localctx.(*Expre_aritmeticaContext).GetOpera().GetText()
					}
				}()) == "%" {
					localctx.(*Expre_aritmeticaContext).p = modulo.Noperacionmodulo(localctx.(*Expre_aritmeticaContext).GetOpIz().GetP(), localctx.(*Expre_aritmeticaContext).GetOpDe().GetP())
				}

			case 2:
				localctx = NewExpre_aritmeticaContext(p, _parentctx, _parentState)
				localctx.(*Expre_aritmeticaContext).opIz = _prevctx
				p.PushNewRecursionContext(localctx, _startState, optimizador_parserRULE_expre_aritmetica)
				p.SetState(266)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(267)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expre_aritmeticaContext).opera = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == optimizador_parserTK_SUMA || _la == optimizador_parserTK_RESTA) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expre_aritmeticaContext).opera = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(268)

					var _x = p.expre_aritmetica(3)

					localctx.(*Expre_aritmeticaContext).opDe = _x
				}

				if (func() string {
					if localctx.(*Expre_aritmeticaContext).GetOpera() == nil {
						return ""
					} else {
						return localctx.(*Expre_aritmeticaContext).GetOpera().GetText()
					}
				}()) == "+" {
					localctx.(*Expre_aritmeticaContext).p = suma.Noperacionsuma(localctx.(*Expre_aritmeticaContext).GetOpIz().GetP(), localctx.(*Expre_aritmeticaContext).GetOpDe().GetP())
				} else {
					localctx.(*Expre_aritmeticaContext).p = resta.Noperacionresta(localctx.(*Expre_aritmeticaContext).GetOpIz().GetP(), localctx.(*Expre_aritmeticaContext).GetOpDe().GetP())
				}

			}

		}
		p.SetState(275)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())
	}

	return localctx
}

// IDatosContext is an interface to support dynamic dispatch.
type IDatosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_TK_ENTERO returns the _TK_ENTERO token.
	Get_TK_ENTERO() antlr.Token

	// Get_TK_FLOAT returns the _TK_FLOAT token.
	Get_TK_FLOAT() antlr.Token

	// Get_TK_IDENTIFICADOR returns the _TK_IDENTIFICADOR token.
	Get_TK_IDENTIFICADOR() antlr.Token

	// Get_TK_CADENA returns the _TK_CADENA token.
	Get_TK_CADENA() antlr.Token

	// Get_TK_TEMPORAL returns the _TK_TEMPORAL token.
	Get_TK_TEMPORAL() antlr.Token

	// Get_TK_PUNSTACK returns the _TK_PUNSTACK token.
	Get_TK_PUNSTACK() antlr.Token

	// Get_TK_PUNHEAP returns the _TK_PUNHEAP token.
	Get_TK_PUNHEAP() antlr.Token

	// Set_TK_ENTERO sets the _TK_ENTERO token.
	Set_TK_ENTERO(antlr.Token)

	// Set_TK_FLOAT sets the _TK_FLOAT token.
	Set_TK_FLOAT(antlr.Token)

	// Set_TK_IDENTIFICADOR sets the _TK_IDENTIFICADOR token.
	Set_TK_IDENTIFICADOR(antlr.Token)

	// Set_TK_CADENA sets the _TK_CADENA token.
	Set_TK_CADENA(antlr.Token)

	// Set_TK_TEMPORAL sets the _TK_TEMPORAL token.
	Set_TK_TEMPORAL(antlr.Token)

	// Set_TK_PUNSTACK sets the _TK_PUNSTACK token.
	Set_TK_PUNSTACK(antlr.Token)

	// Set_TK_PUNHEAP sets the _TK_PUNHEAP token.
	Set_TK_PUNHEAP(antlr.Token)

	// GetP returns the p attribute.
	GetP() inter.Expresion

	// SetP sets the p attribute.
	SetP(inter.Expresion)

	// IsDatosContext differentiates from other interfaces.
	IsDatosContext()
}

type DatosContext struct {
	*antlr.BaseParserRuleContext
	parser            antlr.Parser
	p                 inter.Expresion
	_TK_ENTERO        antlr.Token
	_TK_FLOAT         antlr.Token
	_TK_IDENTIFICADOR antlr.Token
	_TK_CADENA        antlr.Token
	_TK_TEMPORAL      antlr.Token
	_TK_PUNSTACK      antlr.Token
	_TK_PUNHEAP       antlr.Token
}

func NewEmptyDatosContext() *DatosContext {
	var p = new(DatosContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = optimizador_parserRULE_datos
	return p
}

func (*DatosContext) IsDatosContext() {}

func NewDatosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DatosContext {
	var p = new(DatosContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = optimizador_parserRULE_datos

	return p
}

func (s *DatosContext) GetParser() antlr.Parser { return s.parser }

func (s *DatosContext) Get_TK_ENTERO() antlr.Token { return s._TK_ENTERO }

func (s *DatosContext) Get_TK_FLOAT() antlr.Token { return s._TK_FLOAT }

func (s *DatosContext) Get_TK_IDENTIFICADOR() antlr.Token { return s._TK_IDENTIFICADOR }

func (s *DatosContext) Get_TK_CADENA() antlr.Token { return s._TK_CADENA }

func (s *DatosContext) Get_TK_TEMPORAL() antlr.Token { return s._TK_TEMPORAL }

func (s *DatosContext) Get_TK_PUNSTACK() antlr.Token { return s._TK_PUNSTACK }

func (s *DatosContext) Get_TK_PUNHEAP() antlr.Token { return s._TK_PUNHEAP }

func (s *DatosContext) Set_TK_ENTERO(v antlr.Token) { s._TK_ENTERO = v }

func (s *DatosContext) Set_TK_FLOAT(v antlr.Token) { s._TK_FLOAT = v }

func (s *DatosContext) Set_TK_IDENTIFICADOR(v antlr.Token) { s._TK_IDENTIFICADOR = v }

func (s *DatosContext) Set_TK_CADENA(v antlr.Token) { s._TK_CADENA = v }

func (s *DatosContext) Set_TK_TEMPORAL(v antlr.Token) { s._TK_TEMPORAL = v }

func (s *DatosContext) Set_TK_PUNSTACK(v antlr.Token) { s._TK_PUNSTACK = v }

func (s *DatosContext) Set_TK_PUNHEAP(v antlr.Token) { s._TK_PUNHEAP = v }

func (s *DatosContext) GetP() inter.Expresion { return s.p }

func (s *DatosContext) SetP(v inter.Expresion) { s.p = v }

func (s *DatosContext) TK_ENTERO() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_ENTERO, 0)
}

func (s *DatosContext) TK_FLOAT() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_FLOAT, 0)
}

func (s *DatosContext) TK_IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_IDENTIFICADOR, 0)
}

func (s *DatosContext) TK_CADENA() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_CADENA, 0)
}

func (s *DatosContext) TK_TEMPORAL() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_TEMPORAL, 0)
}

func (s *DatosContext) TK_PUNSTACK() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_PUNSTACK, 0)
}

func (s *DatosContext) TK_PUNHEAP() antlr.TerminalNode {
	return s.GetToken(optimizador_parserTK_PUNHEAP, 0)
}

func (s *DatosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DatosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(optimizador_parserListener); ok {
		listenerT.EnterDatos(s)
	}
}

func (s *DatosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(optimizador_parserListener); ok {
		listenerT.ExitDatos(s)
	}
}

func (p *optimizador_parser) Datos() (localctx IDatosContext) {
	localctx = NewDatosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, optimizador_parserRULE_datos)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(290)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case optimizador_parserTK_ENTERO:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(276)

			var _m = p.Match(optimizador_parserTK_ENTERO)

			localctx.(*DatosContext)._TK_ENTERO = _m
		}

		localctx.(*DatosContext).p = primitiivos.NuevoPrimitivo((func() string {
			if localctx.(*DatosContext).Get_TK_ENTERO() == nil {
				return ""
			} else {
				return localctx.(*DatosContext).Get_TK_ENTERO().GetText()
			}
		}()), bloque.INTEGER)

	case optimizador_parserTK_FLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(278)

			var _m = p.Match(optimizador_parserTK_FLOAT)

			localctx.(*DatosContext)._TK_FLOAT = _m
		}

		localctx.(*DatosContext).p = primitiivos.NuevoPrimitivo((func() string {
			if localctx.(*DatosContext).Get_TK_FLOAT() == nil {
				return ""
			} else {
				return localctx.(*DatosContext).Get_TK_FLOAT().GetText()
			}
		}()), bloque.FLOAT)

	case optimizador_parserTK_IDENTIFICADOR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(280)

			var _m = p.Match(optimizador_parserTK_IDENTIFICADOR)

			localctx.(*DatosContext)._TK_IDENTIFICADOR = _m
		}

		localctx.(*DatosContext).p = primitiivos.NuevoPrimitivo((func() string {
			if localctx.(*DatosContext).Get_TK_IDENTIFICADOR() == nil {
				return ""
			} else {
				return localctx.(*DatosContext).Get_TK_IDENTIFICADOR().GetText()
			}
		}()), bloque.NULL)

	case optimizador_parserTK_CADENA:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(282)

			var _m = p.Match(optimizador_parserTK_CADENA)

			localctx.(*DatosContext)._TK_CADENA = _m
		}

		localctx.(*DatosContext).p = primitiivos.NuevoPrimitivo((func() string {
			if localctx.(*DatosContext).Get_TK_CADENA() == nil {
				return ""
			} else {
				return localctx.(*DatosContext).Get_TK_CADENA().GetText()
			}
		}()), bloque.NULL)

	case optimizador_parserTK_TEMPORAL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(284)

			var _m = p.Match(optimizador_parserTK_TEMPORAL)

			localctx.(*DatosContext)._TK_TEMPORAL = _m
		}

		localctx.(*DatosContext).p = primitiivos.NuevoPrimitivo((func() string {
			if localctx.(*DatosContext).Get_TK_TEMPORAL() == nil {
				return ""
			} else {
				return localctx.(*DatosContext).Get_TK_TEMPORAL().GetText()
			}
		}()), bloque.TEMPORAL)

	case optimizador_parserTK_PUNSTACK:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(286)

			var _m = p.Match(optimizador_parserTK_PUNSTACK)

			localctx.(*DatosContext)._TK_PUNSTACK = _m
		}

		localctx.(*DatosContext).p = primitiivos.NuevoPrimitivo((func() string {
			if localctx.(*DatosContext).Get_TK_PUNSTACK() == nil {
				return ""
			} else {
				return localctx.(*DatosContext).Get_TK_PUNSTACK().GetText()
			}
		}()), bloque.NULL)

	case optimizador_parserTK_PUNHEAP:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(288)

			var _m = p.Match(optimizador_parserTK_PUNHEAP)

			localctx.(*DatosContext)._TK_PUNHEAP = _m
		}

		localctx.(*DatosContext).p = primitiivos.NuevoPrimitivo((func() string {
			if localctx.(*DatosContext).Get_TK_PUNHEAP() == nil {
				return ""
			} else {
				return localctx.(*DatosContext).Get_TK_PUNHEAP().GetText()
			}
		}()), bloque.NULL)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *optimizador_parser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 4:
		var t *L_temporalesContext = nil
		if localctx != nil {
			t = localctx.(*L_temporalesContext)
		}
		return p.L_temporales_Sempred(t, predIndex)

	case 5:
		var t *L_funcasContext = nil
		if localctx != nil {
			t = localctx.(*L_funcasContext)
		}
		return p.L_funcas_Sempred(t, predIndex)

	case 13:
		var t *L_bloqueContext = nil
		if localctx != nil {
			t = localctx.(*L_bloqueContext)
		}
		return p.L_bloque_Sempred(t, predIndex)

	case 17:
		var t *Expre_relacionalContext = nil
		if localctx != nil {
			t = localctx.(*Expre_relacionalContext)
		}
		return p.Expre_relacional_Sempred(t, predIndex)

	case 18:
		var t *Expre_aritmeticaContext = nil
		if localctx != nil {
			t = localctx.(*Expre_aritmeticaContext)
		}
		return p.Expre_aritmetica_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *optimizador_parser) L_temporales_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *optimizador_parser) L_funcas_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *optimizador_parser) L_bloque_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *optimizador_parser) Expre_relacional_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *optimizador_parser) Expre_aritmetica_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
