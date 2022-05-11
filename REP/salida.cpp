/*------ENACABEZADO------*/
#include <stdio.h>
/*------DECLARACION HEAP------*/
double HEAP[100000];
/*------DECLARACION STACK------*/
double STACK[100000];
/*------DECLARACION PUNTERO STACK------*/
double SP;
/*------DECLARACION PUNTERO HEAP------*/
double HP;
/*------DECLARACION TEMPORALES------*/
double T0, T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16, T17, T18;

/*INICIO DECLARACION FUNCION NATIVA PRINT*/
void print() {
T0 = SP;
T1 = STACK[(int)T0];
L0:
T2 = HEAP[(int)T1];
if ( T2 == -1) goto L1;
printf("%c", (char)T2);
T1 = T1 + 1;
goto L0;
L1:
return;
}

/*FIN DECLARACION FUNCION NATIVA PRINT*/
/*------------------ENTRANDO FUNCION PRINCIPAL-------------------*/
int main() {
SP=0;		//INICIANDO PUNTERO STACK
HP=0;		//INICIANDO PUNTERO HEAP


/*INICIO DECLARACION ARREGLO -- xy*/
T3 = SP + 0;		// POSICION: xy
STACK[(int)T3] = HP;
T4 = HP;
HP = HP +50;
HEAP[(int)HP] = -2;
HP = HP + 1;
T5 = T4 + 0;
HEAP[(int)T5] = 1;
T5 = T4 + 1;
HEAP[(int)T5] = 2;
T5 = T4 + 2;
HEAP[(int)T5] = 3;
T5 = T4 + 3;
HEAP[(int)T5] = 4;
T5 = T4 + 4;
HEAP[(int)T5] = 5;
T5 = T4 + 5;
HEAP[(int)T5] = 6;
T5 = T4 + 6;
HEAP[(int)T5] = 7;
T5 = T4 + 7;
HEAP[(int)T5] = 8;
T5 = T4 + 8;
HEAP[(int)T5] = 9;
T5 = T4 + 9;
HEAP[(int)T5] = 10;
T5 = T4 + 10;
HEAP[(int)T5] = 11;
T5 = T4 + 11;
HEAP[(int)T5] = 12;
T5 = T4 + 12;
HEAP[(int)T5] = 13;
T5 = T4 + 13;
HEAP[(int)T5] = 14;
T5 = T4 + 14;
HEAP[(int)T5] = 15;
T5 = T4 + 15;
HEAP[(int)T5] = 16;
T5 = T4 + 16;
HEAP[(int)T5] = 17;
T5 = T4 + 17;
HEAP[(int)T5] = 18;
T5 = T4 + 18;
HEAP[(int)T5] = 19;
T5 = T4 + 19;
HEAP[(int)T5] = 20;
T5 = T4 + 20;
HEAP[(int)T5] = 21;
T5 = T4 + 21;
HEAP[(int)T5] = 22;
T5 = T4 + 22;
HEAP[(int)T5] = 31;
T5 = T4 + 23;
HEAP[(int)T5] = 32;
T5 = T4 + 24;
HEAP[(int)T5] = 33;
T5 = T4 + 25;
HEAP[(int)T5] = 34;
T5 = T4 + 26;
HEAP[(int)T5] = 35;
T5 = T4 + 27;
HEAP[(int)T5] = 36;
T5 = T4 + 28;
HEAP[(int)T5] = 37;
T5 = T4 + 29;
HEAP[(int)T5] = 38;
T5 = T4 + 30;
HEAP[(int)T5] = 39;
T5 = T4 + 31;
HEAP[(int)T5] = 30;
T5 = T4 + 32;
HEAP[(int)T5] = 41;
T5 = T4 + 33;
HEAP[(int)T5] = 42;
T5 = T4 + 34;
HEAP[(int)T5] = 43;
T5 = T4 + 35;
HEAP[(int)T5] = 53;
T5 = T4 + 36;
HEAP[(int)T5] = 54;
T5 = T4 + 37;
HEAP[(int)T5] = 55;
T5 = T4 + 38;
HEAP[(int)T5] = 56;
T5 = T4 + 39;
HEAP[(int)T5] = 57;
T5 = T4 + 40;
HEAP[(int)T5] = 58;
T5 = T4 + 41;
HEAP[(int)T5] = 59;
T5 = T4 + 42;
HEAP[(int)T5] = 60;
T5 = T4 + 43;
HEAP[(int)T5] = 61;
T5 = T4 + 44;
HEAP[(int)T5] = 71;
T5 = T4 + 45;
HEAP[(int)T5] = 72;
T5 = T4 + 46;
HEAP[(int)T5] = 73;
T5 = T4 + 47;
HEAP[(int)T5] = 74;
T5 = T4 + 48;
HEAP[(int)T5] = 75;
T5 = T4 + 49;
HEAP[(int)T5] = 76;
T5 = T4 + 50;
HEAP[(int)T5] = 77;
T5 = T4 + 51;
HEAP[(int)T5] = 78;
T5 = T4 + 52;
HEAP[(int)T5] = 79;
T5 = T4 + 53;
HEAP[(int)T5] = 110;
T5 = T4 + 54;
HEAP[(int)T5] = 111;
T5 = T4 + 55;
HEAP[(int)T5] = 81;
T5 = T4 + 56;
HEAP[(int)T5] = 82;
T5 = T4 + 57;
HEAP[(int)T5] = 83;
T5 = T4 + 58;
HEAP[(int)T5] = 84;
T5 = T4 + 59;
HEAP[(int)T5] = 85;
T5 = T4 + 60;
HEAP[(int)T5] = 86;
T5 = T4 + 61;
HEAP[(int)T5] = 87;
T5 = T4 + 62;
HEAP[(int)T5] = 88;
T5 = T4 + 63;
HEAP[(int)T5] = 89;
T5 = T4 + 64;
HEAP[(int)T5] = 100;
T5 = T4 + 65;
HEAP[(int)T5] = 101;
/*FIN DECLARACION ARREGLO -- xy*/
T8 = SP + 0;
T9 = STACK[(int)T8];
/*-IMPRIMIR TEXTO-*/
T6 = HP;
HEAP[(int)HP] = 84; //LETRA-> T
HP = HP + 1;
HEAP[(int)HP] = 111; //LETRA-> o
HP = HP + 1;
HEAP[(int)HP] = 100; //LETRA-> d
HP = HP + 1;
HEAP[(int)HP] = 111; //LETRA-> o
HP = HP + 1;
HEAP[(int)HP] = 32; //LETRA->  
HP = HP + 1;
HEAP[(int)HP] = 97; //LETRA-> a
HP = HP + 1;
HEAP[(int)HP] = 114; //LETRA-> r
HP = HP + 1;
HEAP[(int)HP] = 114; //LETRA-> r
HP = HP + 1;
HEAP[(int)HP] = 101; //LETRA-> e
HP = HP + 1;
HEAP[(int)HP] = 103; //LETRA-> g
HP = HP + 1;
HEAP[(int)HP] = 108; //LETRA-> l
HP = HP + 1;
HEAP[(int)HP] = 111; //LETRA-> o
HP = HP + 1;
HEAP[(int)HP] = 32; //LETRA->  
HP = HP + 1;
HEAP[(int)HP] = 50; //LETRA-> 2
HP = HP + 1;
HEAP[(int)HP] = 100; //LETRA-> d
HP = HP + 1;
HEAP[(int)HP] = 105; //LETRA-> i
HP = HP + 1;
HEAP[(int)HP] = 109; //LETRA-> m
HP = HP + 1;
HEAP[(int)HP] = 58; //LETRA-> :
HP = HP + 1;
HEAP[(int)HP] = 32; //LETRA->  
HP = HP + 1;
HEAP[(int)HP] = -1;
HP = HP + 1;
T7 = SP + 1; //POSICION CADENA EN STACK
STACK[(int)T7] = T6;
SP = SP + 1;
print();
SP = SP - 1;
L2:
T10 = HEAP[(int)T9];
if ( T10 == -2) goto L3;
printf("%d",(int)T10);
printf("%c", (char) 32);
T9 = T9 + 1;
goto L2;
L3:
printf("%c", (char)10);

/*INCIO ACCESO A ARREGLO -- xy*/
T13 = 2 * 10;
T14 = T13 + 1;
if (T14 > 50) goto L4;
T15 = SP + 0;
T16 = STACK[(int)T15];
T17 = T16 + T14;
T18 = HEAP[(int)T17];
goto L5;
L4:
printf("%c", 73); //LETRA-> I
printf("%c", 78); //LETRA-> N
printf("%c", 68); //LETRA-> D
printf("%c", 73); //LETRA-> I
printf("%c", 67); //LETRA-> C
printf("%c", 69); //LETRA-> E
printf("%c", 32); //LETRA->  
printf("%c", 70); //LETRA-> F
printf("%c", 85); //LETRA-> U
printf("%c", 69); //LETRA-> E
printf("%c", 82); //LETRA-> R
printf("%c", 65); //LETRA-> A
printf("%c", 32); //LETRA->  
printf("%c", 68); //LETRA-> D
printf("%c", 69); //LETRA-> E
printf("%c", 32); //LETRA->  
printf("%c", 76); //LETRA-> L
printf("%c", 73); //LETRA-> I
printf("%c", 77); //LETRA-> M
printf("%c", 73); //LETRA-> I
printf("%c", 84); //LETRA-> T
printf("%c", 69); //LETRA-> E
printf("%c", 46);	//PUNTO
printf("%c", 32); //LETRA->  
printf("%c", 69); //LETRA-> E
printf("%c", 110); //LETRA-> n
printf("%c", 32); //LETRA->  
printf("%c", 108); //LETRA-> l
printf("%c", 97); //LETRA-> a
printf("%c", 32); //LETRA->  
printf("%c", 76); //LETRA-> L
printf("%c", 105); //LETRA-> i
printf("%c", 110); //LETRA-> n
printf("%c", 101); //LETRA-> e
printf("%c", 97); //LETRA-> a
printf("%c", 58); //LETRA-> :
printf("%c", 32); //LETRA->  
printf("%o", (int)10);	//NUM
printf("%c", 46);	//PUNTO
printf("%c", 32);	//ESPACIO
printf("%c", 32); //LETRA->  
printf("%c", 69); //LETRA-> E
printf("%c", 110); //LETRA-> n
printf("%c", 32); //LETRA->  
printf("%c", 108); //LETRA-> l
printf("%c", 97); //LETRA-> a
printf("%c", 32); //LETRA->  
printf("%c", 99); //LETRA-> c
printf("%c", 111); //LETRA-> o
printf("%c", 108); //LETRA-> l
printf("%c", 117); //LETRA-> u
printf("%c", 109); //LETRA-> m
printf("%c", 110); //LETRA-> n
printf("%c", 97); //LETRA-> a
printf("%c", 58); //LETRA-> :
printf("%c", 32); //LETRA->  
printf("%d", (int)48);	//NUM
printf("%c", 46);	//PUNTO
printf("%c", 10);	//SALTO LINEA

L5:
/*-IMPRIMIR TEXTO-*/
T11 = HP;
HEAP[(int)HP] = 80; //LETRA-> P
HP = HP + 1;
HEAP[(int)HP] = 111; //LETRA-> o
HP = HP + 1;
HEAP[(int)HP] = 115; //LETRA-> s
HP = HP + 1;
HEAP[(int)HP] = 105; //LETRA-> i
HP = HP + 1;
HEAP[(int)HP] = 99; //LETRA-> c
HP = HP + 1;
HEAP[(int)HP] = 105; //LETRA-> i
HP = HP + 1;
HEAP[(int)HP] = 111; //LETRA-> o
HP = HP + 1;
HEAP[(int)HP] = 110; //LETRA-> n
HP = HP + 1;
HEAP[(int)HP] = 32; //LETRA->  
HP = HP + 1;
HEAP[(int)HP] = 91; //LETRA-> [
HP = HP + 1;
HEAP[(int)HP] = 50; //LETRA-> 2
HP = HP + 1;
HEAP[(int)HP] = 93; //LETRA-> ]
HP = HP + 1;
HEAP[(int)HP] = 91; //LETRA-> [
HP = HP + 1;
HEAP[(int)HP] = 48; //LETRA-> 0
HP = HP + 1;
HEAP[(int)HP] = 93; //LETRA-> ]
HP = HP + 1;
HEAP[(int)HP] = 32; //LETRA->  
HP = HP + 1;
HEAP[(int)HP] = 97; //LETRA-> a
HP = HP + 1;
HEAP[(int)HP] = 114; //LETRA-> r
HP = HP + 1;
HEAP[(int)HP] = 114; //LETRA-> r
HP = HP + 1;
HEAP[(int)HP] = 101; //LETRA-> e
HP = HP + 1;
HEAP[(int)HP] = 103; //LETRA-> g
HP = HP + 1;
HEAP[(int)HP] = 108; //LETRA-> l
HP = HP + 1;
HEAP[(int)HP] = 111; //LETRA-> o
HP = HP + 1;
HEAP[(int)HP] = 32; //LETRA->  
HP = HP + 1;
HEAP[(int)HP] = 50; //LETRA-> 2
HP = HP + 1;
HEAP[(int)HP] = 100; //LETRA-> d
HP = HP + 1;
HEAP[(int)HP] = 105; //LETRA-> i
HP = HP + 1;
HEAP[(int)HP] = 109; //LETRA-> m
HP = HP + 1;
HEAP[(int)HP] = 58; //LETRA-> :
HP = HP + 1;
HEAP[(int)HP] = 32; //LETRA->  
HP = HP + 1;
HEAP[(int)HP] = -1;
HP = HP + 1;
T12 = SP + 1; //POSICION CADENA EN STACK
STACK[(int)T12] = T11;
SP = SP + 1;
print();
SP = SP - 1;
printf("%d",(int)T18);
printf("%c", (char)10);

return 0;
}
/*------------------SALIENDO FUNCION PRINCIPAL-------------------*/
