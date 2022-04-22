// yacc hello world

// part1: go code
%{
package main
import "fmt"
%}

// part2: type definition
%union {
    ystring string
}

%type <ystring> expression
%token <ystring> HELLO HI
%start expression

%%

// part3: grammer rule
expression : HELLO {
    fmt.Println("this is hello")
} | HI {
    fmt.Println("this is hi")
};