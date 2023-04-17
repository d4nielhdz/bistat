grammar Bistat;

program: 'Program' ID ';' varDeclaration* funcDef* main EOF;

varDeclaration: 'var' NON_VOID_TYPE ID ';';
funcDef:
	'func' ID '(' paramDeclaration* ')' ':' RETURN_TYPE varDeclaration* '{' stmt+ '}';
main: 'main' '(' ')' '{' stmt+ '}';

stmt: ((assignment | print | specialFunction | functionCall) ';')
	| conditional
	| whileLoop
	| forLoop
	| comment;

paramDeclaration: 'var' PARAM_TYPE ID ';';
assignment:
	ID '=' (VAR_CONS | listAssignment | matrixAssignment);
nested_stmt: (
		(assignment | print | specialFunction | functionCall) ';'
	)
	| conditional
	| whileLoop
	| comment;

matrixAssignment: '[' listAssignment (',' listAssignment)* ']';
listAssignment: '[' VAR_CONS (',' VAR_CONS)* ']';

comment: '#' (~'#')+ '#';
forLoop: 'for' '(' ID 'in' expression ')' '{' nested_stmt+ '}';
whileLoop: 'while' '(' expression ')' '{' stmt+ '}';
conditional:
	'if' '(' expression ')' '{' stmt+ '}' (
		'else' 'if' '(' expression ')' '{' stmt+ '}'
	)* ('else' '{' stmt+ '}')?;

specialFunction:
	inputRead
	| print
	| listAdd
	| listPop
	| length
	| range
	| plot
	| sum
	| min
	| prod
	| avg
	| sMode
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
inputRead: 'read' '(' ID (',' ID)* ')';
print: 'print' '(' expression ( ',' expression)* ')';
listAdd: 'listAdd' '(' expression ',' expression ')';
listPop: 'listPop' '(' expression ')';
length: 'length' '(' expression ')';
range: 'range' '(' expression (',' expression)? ')';
plot: 'plot' '(' expression ')';
sum: 'sum' '(' expression ')';
min: 'min' '(' expression ')';
prod: 'prod' '(' expression ')';
avg: 'avg' '(' expression ')';
sMode: 'sMode' '(' expression ')';
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
		| specialFunction
		| functionCall
		| VAR_CONS
	);
functionCall: ID '(' (expression (',' expression)*)? ')';
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
