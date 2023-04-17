grammar Bistat;

program: 'Program' ID ';' var_declaration* func_def* main EOF;

var_declaration: 'var' NON_VOID_TYPE ID ';';
func_def:
	'func' ID '(' param_declaration* ')' ':' RETURN_TYPE var_declaration* '{' stmt+ '}';
main: 'main' '(' ')' '{' stmt+ '}';

stmt: (
		(assignment | print | special_function | function_call) ';'
	)
	| conditional
	| while_loop
	| for_loop
	| comment;

param_declaration: 'var' PARAM_TYPE ID ';';
assignment:
	ID '=' (VAR_CONS | list_assignment | matrix_assignment);
nested_stmt: (
		(assignment | print | special_function | function_call) ';'
	)
	| conditional
	| while_loop
	| comment;

matrix_assignment:
	'[' list_assignment (',' list_assignment)* ']';
list_assignment: '[' VAR_CONS (',' VAR_CONS)* ']';

comment: '#' (~'#')+ '#';
for_loop: 'for' '(' ID 'in' expression ')' '{' nested_stmt+ '}';
while_loop: 'while' '(' expression ')' '{' stmt+ '}';
conditional:
	'if' '(' expression ')' '{' stmt+ '}' (
		'else' 'if' '(' expression ')' '{' stmt+ '}'
	)* ('else' '{' stmt+ '}')?;

special_function:
	input_read
	| print
	| list_add
	| list_pop
	| length
	| range
	| plot
	| sum
	| min
	| prod
	| avg
	| s_mode
	| median
	| sin
	| cos
	| tan
	| sort
	| sqrt
	| floor
	| ceil
	| abs
	| not;
input_read: 'read' '(' ID (',' ID)* ')';
print: 'print' '(' expression ( ',' expression)* ')';
list_add: 'list_add' '(' expression ',' expression ')';
list_pop: 'list_pop' '(' expression ')';
length: 'length' '(' expression ')';
range: 'range' '(' expression (',' expression)? ')';
plot: 'plot' '(' expression ')';
sum: 'sum' '(' expression ')';
min: 'min' '(' expression ')';
prod: 'prod' '(' expression ')';
avg: 'avg' '(' expression ')';
s_mode: 's_mode' '(' expression ')';
median: 'median' '(' expression ')';
sin: 'sin' '(' expression ')';
tan: 'tan' '(' expression ')';
cos: 'cos' '(' expression ')';
sort: 'sort' '(' expression ')';
sqrt: 'sqrt' '(' expression ')';
floor: 'floor' '(' expression ')';
ceil: 'ceil' '(' expression ')';
abs: 'abs' '(' expression ')';
not: 'not' '(' expression ')';

expression: exp (LOGIC_OPERATOR exp)?;
exp: term (OP_SEC exp)?;
term: factor ( OP_FIRST term)?;
factor:
	OP_SEC? (
		( '(' expression ')')
		| indexing
		| special_function
		| function_call
		| VAR_CONS
	);
function_call: ID '(' (expression (',' expression)*)? ')';
indexing: ID '[' expression ']' ('[' expression ']')?;

WS: [ \t\n\r]+ -> channel(HIDDEN);
NON_VOID_TYPE: Type_Primitive CARDINALITY?;
RETURN_TYPE: NON_VOID_TYPE | 'void';
PARAM_TYPE: Type_Primitive PARAM_CARDINALITY?;
fragment Type_Primitive: 'int' | 'float' | 'string' | 'bool';
ID: ('_' | Alpha)+ (Alpha | NUMBER | '_')*;
PARAM_CARDINALITY: ('[' ']') | ('[' ']' '[' ']');
CARDINALITY: ('[' INT_CONS ']')
	| ('[' INT_CONS ']' '[' INT_CONS ']');
VAR_CONS: STRING_CONS | FLOAT_CONS | INT_CONS | BOOL_CONS | ID;
BOOL_CONS: 'true' | 'false';
STRING_CONS: '"' (~'"')* '"';
FLOAT_CONS: NUMBER+ '.' NUMBER+;
INT_CONS: NUMBER+;
OP_SEC: '+' | '-';
OP_FIRST: '/' | '*' | '%';
LOGIC_OPERATOR:
	'<'
	| '>'
	| '<='
	| '>='
	| '=='
	| '!='
	| '&&'
	| '||';
fragment Alpha: 'A' .. 'Z' | 'a' .. 'z';
NUMBER: '0' .. '9';
