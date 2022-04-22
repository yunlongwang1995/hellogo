%{
package parser

import (
    "fmt"
)
%}

%union {
   var string
   num int
}

%token <num> NUM
%token <var> ADD SUB '+' '-'

%type <val> expr

%start expr

%%

expr: NUM {
  $$ =$1
}
| expr '+' NUM {
  $$ = $1 + $3
}
| expr '-' NUM {
  $$ = $1 + $3
}
;