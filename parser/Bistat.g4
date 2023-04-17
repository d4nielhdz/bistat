grammar Bistat;

WS: [ \t\n\r]+ -> channel(HIDDEN);

program: 'Program' ID ';' varDeclaration* funcDef* main EOF;

varDeclaration: 'var' NON_VOID_TYPE ID ';';
funcDef:
	'func' ID '(' paramDeclaration* ')' ':' RETURN_TYPE varDeclaration* '{' stmt+ '}';
main: 'main' '(' ')' '{' stmt+ '}';

stmt: (
		(
			assignment
			| specialFunction
			| functionCall
			| returnStmt
		) ';'
	)
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
returnStmt: 'return' expression;

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

NON_VOID_TYPE: ('int' | 'float' | 'string' | 'bool') (
		CARDINALITY?
	);
PARAM_TYPE: ('int' | 'float' | 'string' | 'bool') (
		PARAM_CARDINALITY?
	);
RETURN_TYPE:
	(('int' | 'float' | 'string' | 'bool') (PARAM_CARDINALITY?))
	| 'void';

CARDINALITY: ('[' INT_CONS ']')
	| ('[' INT_CONS ']' '[' INT_CONS ']');
PARAM_CARDINALITY: ('[' ']') | ('[' ']' '[' ']');

ID: ('_' | Alpha)+ (Alpha | NUMBER | '_')*;
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
