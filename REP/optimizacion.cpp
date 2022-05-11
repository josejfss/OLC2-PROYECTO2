#include <stdio.h>
double HEAP[100000];
double STACK[100000];
double SP;
double HP;
double T1,T2,T3,T4,T5,T7,T9,T10,T11,T12,T14,T15,T16,T17,T18,T19;
int main() {
SP=0;
HP=0;
L0:
T1=6/2;
T2=3*T1;
T3=4-2;
T4=T2*T3;
T5=2+5;
T17=T4-6;
L1:
T7=3+1;
T9=4*2;
T10=T7+2;
T18=HEAP[(int)T7];
STACK[(int)T19]=T5;
if(T7<=10) goto L2;
goto L3;
L2:
T11=6/2;
T12=3*T11;
T14=T12*T3;
T15=T5;
T16=T14;
L3:
HP=HP+1;
printf("%d", (int)T7);
printf("%d", (int)6);
printf("%d", (int)T2);
return 0;
}

