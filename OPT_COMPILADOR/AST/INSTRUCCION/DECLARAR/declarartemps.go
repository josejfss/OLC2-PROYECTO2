package declarar

import (
	reportes "OLC2-PROYECTO2/COMPILADOR/REPORTES"
	bloque "OLC2-PROYECTO2/OPT_COMPILADOR/BLOQUE"
	generando "OLC2-PROYECTO2/OPT_COMPILADOR/GENERANDO"
	inter "OLC2-PROYECTO2/OPT_COMPILADOR/INTER"
	"strconv"
)

type Declaracion struct {
	Temporal string
	Valor    inter.Expresion
	Linea    int
}

func Ndeclaracion(tem string, val inter.Expresion, lin int) Declaracion {
	decla := Declaracion{Temporal: tem, Valor: val, Linea: lin}
	return decla
}

func (decla Declaracion) Opti_Instruccion(block *bloque.BloquesOpt, genopt *generando.GenerandoOpti) interface{} {
	valopt := decla.Valor.Opti_Expresion(block)

	if len(block.TablaObjBloques) != 0 {
		verdad, valor := block.Regla1(valopt.Valor)
		op_iz := block.Obtener_Temporal(valopt.OpeIz)
		op_de := block.Obtener_Temporal(valopt.OpeDe)
		if verdad {
			reportes.ReporteOpti("Bloques", "Regla 1", decla.Temporal+"="+valopt.Valor, decla.Temporal+"="+valor, strconv.Itoa(decla.Linea))
			simb_decla := bloque.ObjetoBloque{Tipo: bloque.TEMPORAL,
				Asignacion: decla.Temporal,
				OpeIz:      valor,
				Ope:        valopt.Ope,
				OpeDe:      valor,
				Valor:      valor,
				Linea:      decla.Linea}
			block.Guardar(decla.Temporal, simb_decla)
			guardar := bloque.Nobjetolista(decla.Temporal, decla.Temporal+"="+valor+";")
			genopt.AgegarCodigoOpt(guardar)
		} else if op_iz.Tipo == bloque.TEMPORAL {
			reportes.ReporteOpti("Bloques", "Regla 2", decla.Temporal+"="+valopt.Valor, decla.Temporal+"="+op_iz.Valor+valopt.Ope+valopt.OpeDe, strconv.Itoa(decla.Linea))
			//reportes.ReporteOpti("Bloques", "Regla 3", op_iz.Asignacion+"="+op_iz.Valor, "--", strconv.Itoa(decla.Linea))
			genopt.EliminarTempo(op_iz.Asignacion, op_iz.Valor, strconv.Itoa(decla.Linea))
			block.EliminarTemporalLista(op_iz.Asignacion)

			if valopt.Ope == "HEAP" {
				simb_decla := bloque.ObjetoBloque{Tipo: bloque.NULL,
					Asignacion: decla.Temporal,
					OpeIz:      op_iz.Valor,
					Ope:        valopt.Ope,
					OpeDe:      valopt.OpeDe,
					Valor:      op_iz.Valor,
					Linea:      decla.Linea}
				block.Guardar(decla.Temporal, simb_decla)
				guardar := bloque.Nobjetolista(decla.Temporal, decla.Temporal+"="+"HEAP[(int)"+op_iz.Valor+"];")
				genopt.AgegarCodigoOpt(guardar)
			} else if valopt.Ope == "STACK" {
				simb_decla := bloque.ObjetoBloque{Tipo: bloque.NULL,
					Asignacion: decla.Temporal,
					OpeIz:      op_iz.Valor,
					Ope:        valopt.Ope,
					OpeDe:      valopt.OpeDe,
					Valor:      op_iz.Valor,
					Linea:      decla.Linea}
				block.Guardar(decla.Temporal, simb_decla)
				guardar := bloque.Nobjetolista(decla.Temporal, decla.Temporal+"="+"STACK[(int)"+op_iz.Valor+"];")
				genopt.AgegarCodigoOpt(guardar)
			} else {
				simb_decla := bloque.ObjetoBloque{Tipo: valopt.Tipo,
					Asignacion: decla.Temporal,
					OpeIz:      op_iz.Valor,
					Ope:        valopt.Ope,
					OpeDe:      valopt.OpeDe,
					Valor:      op_iz.Valor + valopt.Ope + valopt.OpeDe,
					Linea:      decla.Linea}
				block.Guardar(decla.Temporal, simb_decla)
				guardar := bloque.Nobjetolista(decla.Temporal, decla.Temporal+"="+op_iz.Valor+valopt.Ope+valopt.OpeDe+";")
				genopt.AgegarCodigoOpt(guardar)
			}

		} else if op_de.Tipo == bloque.TEMPORAL {
			reportes.ReporteOpti("Bloques", "Regla 2", decla.Temporal+"="+valopt.Valor, decla.Temporal+"="+valopt.OpeIz+valopt.Ope+op_de.Valor, strconv.Itoa(decla.Linea))
			//reportes.ReporteOpti("Bloques", "Regla 3", op_de.Asignacion+"="+op_de.Valor, "--", strconv.Itoa(decla.Linea))
			genopt.EliminarTempo(op_de.Asignacion, op_de.Valor, strconv.Itoa(decla.Linea))
			block.EliminarTemporalLista(op_de.Asignacion)
			simb_decla := bloque.ObjetoBloque{Tipo: valopt.Tipo,
				Asignacion: decla.Temporal,
				OpeIz:      valopt.OpeIz,
				Ope:        valopt.Ope,
				OpeDe:      op_de.Valor,
				Valor:      valopt.OpeIz + valopt.Ope + op_de.Valor,
				Linea:      decla.Linea}
			block.Guardar(decla.Temporal, simb_decla)
			guardar := bloque.Nobjetolista(decla.Temporal, decla.Temporal+"="+valopt.OpeIz+valopt.Ope+op_de.Valor+";")
			genopt.AgegarCodigoOpt(guardar)
		} else if op_iz.Tipo == bloque.CONSTANTE {
			reportes.ReporteOpti("Bloques", "Regla 4", decla.Temporal+"="+valopt.Valor, decla.Temporal+"="+op_iz.Valor+valopt.Ope+valopt.OpeDe, strconv.Itoa(decla.Linea))
			//reportes.ReporteOpti("Bloques", "Regla 3", op_iz.Asignacion+"="+op_iz.Valor, "--", strconv.Itoa(decla.Linea))
			genopt.EliminarTempo(op_iz.Asignacion, op_iz.Valor, strconv.Itoa(decla.Linea))
			block.EliminarTemporalLista(op_iz.Asignacion)
			simb_decla := bloque.ObjetoBloque{Tipo: valopt.Tipo,
				Asignacion: decla.Temporal,
				OpeIz:      op_iz.Valor,
				Ope:        valopt.Ope,
				OpeDe:      valopt.OpeDe,
				Valor:      valor,
				Linea:      decla.Linea}
			block.Guardar(decla.Temporal, simb_decla)
			guardar := bloque.Nobjetolista(decla.Temporal, decla.Temporal+"="+op_iz.Valor+valopt.Ope+valopt.OpeDe+";")
			genopt.AgegarCodigoOpt(guardar)
		} else if op_de.Tipo == bloque.CONSTANTE {
			reportes.ReporteOpti("Bloques", "Regla 4", decla.Temporal+"="+valopt.Valor, decla.Temporal+"="+valopt.OpeIz+valopt.Ope+op_de.Valor, strconv.Itoa(decla.Linea))
			//reportes.ReporteOpti("Bloques", "Regla 3", op_de.Asignacion+"="+op_de.Valor, "--", strconv.Itoa(decla.Linea))
			genopt.EliminarTempo(op_de.Asignacion, op_de.Valor, strconv.Itoa(decla.Linea))
			block.EliminarTemporalLista(op_de.Asignacion)
			simb_decla := bloque.ObjetoBloque{Tipo: valopt.Tipo,
				Asignacion: decla.Temporal,
				OpeIz:      valopt.OpeIz,
				Ope:        valopt.Ope,
				OpeDe:      op_de.Valor,
				Valor:      valor,
				Linea:      decla.Linea}
			block.Guardar(decla.Temporal, simb_decla)
			guardar := bloque.Nobjetolista(decla.Temporal, decla.Temporal+"="+valopt.OpeIz+valopt.Ope+op_de.Valor+";")
			genopt.AgegarCodigoOpt(guardar)
		} else {
			if valopt.Tipo == bloque.INTEGER || valopt.Tipo == bloque.FLOAT {
				simb_decla := bloque.ObjetoBloque{Tipo: bloque.CONSTANTE,
					Asignacion: decla.Temporal,
					OpeIz:      valopt.OpeIz,
					Ope:        valopt.Ope,
					OpeDe:      valopt.OpeDe,
					Valor:      valopt.Valor,
					Linea:      decla.Linea}
				block.Guardar(decla.Temporal, simb_decla)
				guardar := bloque.Nobjetolista(decla.Temporal, decla.Temporal+"="+valopt.Valor+";")
				genopt.AgegarCodigoOpt(guardar)
			} else if valopt.Ope == "HEAP" {
				simb_decla := bloque.ObjetoBloque{Tipo: bloque.NULL,
					Asignacion: decla.Temporal,
					OpeIz:      op_iz.Valor,
					Ope:        valopt.Ope,
					OpeDe:      valopt.OpeDe,
					Valor:      op_iz.Valor,
					Linea:      decla.Linea}
				block.Guardar(decla.Temporal, simb_decla)
				guardar := bloque.Nobjetolista(decla.Temporal, decla.Temporal+"="+"HEAP[(int)"+op_iz.Valor+"];")
				genopt.AgegarCodigoOpt(guardar)
			} else if valopt.Ope == "STACK" {
				simb_decla := bloque.ObjetoBloque{Tipo: bloque.NULL,
					Asignacion: decla.Temporal,
					OpeIz:      op_iz.Valor,
					Ope:        valopt.Ope,
					OpeDe:      valopt.OpeDe,
					Valor:      op_iz.Valor,
					Linea:      decla.Linea}
				block.Guardar(decla.Temporal, simb_decla)
				guardar := bloque.Nobjetolista(decla.Temporal, decla.Temporal+"="+"STACK[(int)"+op_iz.Valor+"];")
				genopt.AgegarCodigoOpt(guardar)
			} else {
				simb_decla := bloque.ObjetoBloque{Tipo: valopt.Tipo,
					Asignacion: decla.Temporal,
					OpeIz:      valopt.OpeIz,
					Ope:        valopt.Ope,
					OpeDe:      valopt.OpeDe,
					Valor:      valopt.Valor,
					Linea:      decla.Linea}
				block.Guardar(decla.Temporal, simb_decla)
				guardar := bloque.Nobjetolista(decla.Temporal, decla.Temporal+"="+valopt.Valor+";")
				genopt.AgegarCodigoOpt(guardar)
			}
		}

	} else {
		if valopt.Tipo == bloque.INTEGER || valopt.Tipo == bloque.FLOAT {
			simb_decla := bloque.ObjetoBloque{Tipo: bloque.CONSTANTE,
				Asignacion: decla.Temporal,
				OpeIz:      valopt.OpeIz,
				Ope:        valopt.Ope,
				OpeDe:      valopt.OpeDe,
				Valor:      valopt.Valor,
				Linea:      decla.Linea}
			block.Guardar(decla.Temporal, simb_decla)
			guardar := bloque.Nobjetolista(decla.Temporal, decla.Temporal+"="+valopt.Valor+";")
			genopt.AgegarCodigoOpt(guardar)
		} else {
			simb_decla := bloque.ObjetoBloque{Tipo: valopt.Tipo,
				Asignacion: decla.Temporal,
				OpeIz:      valopt.OpeIz,
				Ope:        valopt.Ope,
				OpeDe:      valopt.OpeDe,
				Valor:      valopt.Valor,
				Linea:      decla.Linea}
			block.Guardar(decla.Temporal, simb_decla)
			guardar := bloque.Nobjetolista(decla.Temporal, decla.Temporal+"="+valopt.Valor+";")
			genopt.AgegarCodigoOpt(guardar)
		}
	}

	return decla.Temporal
}
