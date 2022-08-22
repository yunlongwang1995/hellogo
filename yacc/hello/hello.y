%{
package hello

import "hellogo/yacc/hello/sem/tree"
%}

%union {
    type_string string
}

%type <type_string> expression
%token <type_string> PRINT HELLO

%%
expression:        HELLO
        {
		st := tree.HelloStatement{
                		Name: "自定义参数",
                	}

                st.Exec()
        }
%%