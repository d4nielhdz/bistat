grammar Bistat;

WS: [ \t\n\r]+ -> channel(HIDDEN);

program: 'Program' ID ';' varDeclaration* funcDef* main EOF;

varDeclaration: 'var' var_type ID ';';
var_type: TYPE_PRIMITIVE (CARDINALITY?);
CARDINALITY: ('[' (INT_CONS) ']') ('[' (INT_CONS) ']')?;
TYPE_PRIMITIVE: 'int' | 'float' | 'string' | 'bool' | 'void';

funcDef:
	'func' ID '(' paramDeclaration* ')' ':' TYPE_PRIMITIVE varDeclaration* funcBlockStart stmt+
		funcBlockEnd;
funcBlockStart: '{';
funcBlockEnd: '}';

paramDeclaration: 'var' TYPE_PRIMITIVE ID ';';

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

assignment:
	(indexing | ID) '=' (
		expression
		| listAssignment
		| matrixAssignment
	);
matrixAssignment: '[' listAssignment (',' listAssignment)* ']';
listAssignment: '[' (expression) (',' (expression))* ']';

comment: '#' (~'#')+ '#';

forLoop: forHeader '{' stmt+ '}';
forHeader: 'for' '(' ID 'in' expression ')';
forExprEnd: ')';

whileLoop: 'while' '(' expression whileExprEnd '{' stmt+ '}';
whileExprEnd: ')';

conditional: ifStmt elseIfStmt* elseStmt?;
ifStmt: 'if' '(' expression condExprEnd '{' stmt+ '}';
elseIfStmt:
	'else' 'if' '(' expression condExprEnd '{' stmt+ '}';
condExprEnd: ')';
elseStmt: 'else' '{' stmt+ '}';

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

expression: exp (logicOperator exp)?;
exp: term (opSec term)*;
term: factor ( opFirst factor)*;
factor:
	unaryMinus? (
		nestedExpression
		| indexing
		| specialFunction
		| functionCall
		| varCons
	);
unaryMinus: '-';
nestedExpression: '(' expression ')';
functionCall: ID '(' (expression (',' expression)*)? ')';
indexing: ID '[' expression ']' ('[' expression ']')?;
BOOL_CONS: 'true' | 'false';
ID: ('_' | Alpha)+ (Alpha | NUMBER | '_')*;
varCons: STRING_CONS | FLOAT_CONS | INT_CONS | BOOL_CONS | ID;
INT_CONS: NUMBER+;
NUMBER: '0' .. '9';
STRING_CONS: '"' (~'"')* '"';
FLOAT_CONS: NUMBER+ '.' NUMBER+;
opSec: '+' | '-';
opFirst: '/' | '*' | '%';
logicOperator:
	'<'
	| '>'
	| '<='
	| '>='
	| '=='
	| '!='
	| '&&'
	| '||';
fragment Alpha: 'A' .. 'Z' | 'a' .. 'z';
