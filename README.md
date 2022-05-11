# OLC2-PROYECTO2
Traducción de DB-RUST a C++ y su optimización.
En este proyecto se realiza la traducción del lenguaje DB-RUST a lenguaje de bajo nivel C++.
Al finalizar se realiza la optimización por bloques. 
Se realiza una primera pasada para guardar en la tabla de simbolos las variables, arreglos, vectores, struct y modulos.
Y luego se realiza la traduccion.
Para la optimización se realiza otra gramatica, donde se analiza el c3d generado en la compilación.

## COMPILADOR
En esta carpeta se realiza la traducción del lenguaje.

### AST 
En esta carpeta se encuentra las expresiones e instrucciones utilizadas en la traducción del lenguaje.

### ENTORNO
En esta carpeta se encuentra el manejo de entorno para la interpretacion.

### GENERADOR
En esta carpeta se encuentra el archivo que se utiliza para escribir C3D.

### INTERFACES
En esta carpeta se encuentra los archivos que se utilizan para las interfaces del patrón.

### REPORTES
En esta carpeta se encuentran los archivos para generar los reportes correspondientes.

### SIMBOLOS
En esta carpeta estan los archivos donde se definen los simbolos a utilizar.

### db_rustlexer.g4
Este archivo es el analizador léxico.

### db_rustparser.g4
Este archivo es el analizador sintactico.

## ENUNCIADO

## OPT_COMPILADOR
En esta carpeta se realiza la optimizacion del C3D.

### AST
En esta carpeta se encuentra las expresiones e instrucciones utilizadas en la optimización del lenguaje generado con C#D.

### BLOQUE
En esta carpeta se encuentra la separación de bloques del C3D, los simbolos y los tipos.

### ENVIAR
En esta carpeta se encuentra el archivo donde se encuentran los metodos para llamar a la gramatica desde el main

### GENERANDO
En esta carpeta se encuentra el archivo para escribir el C3D optimzado.

### INTER
En esta carpeta se encuentra el archivo donde estan las interfaces para la optimización.

### optimizador_lexer.g4
Este archivo es el analizador lexico.

### optimizador_parser.g4
Este archivo es el analizador sintactico.

## REP
Carpeta donde se encuentran los reportes generados.

## pruebas.rs
Este es el archivo de pruebas de la gramatica.