package hello

/*
#include <stdio.h>

// C言語でHello, World!を出力する関数
void printHello() {
    printf("Hello, World!\n");
}
*/
import "C"

// Hello 関数は C の printHello 関数を呼び出します
func Hello() {
	C.printHello()
}
