grammar Bistat;

WS: [ \t\n\r]+ -> channel(HIDDEN);

program: 'Program' ID ';' varDeclaration* funcDef* main EOF;

varDeclaration: 'var' var_type ID ';';
var_type: TYPE_PRIMITIVE (CARDINALITY?);
CARDINALITY: ('[' (INT_CONS?) ']') ('[' (INT_CONS?) ']')?;
TYPE_PRIMITIVE: 'int' | 'float' | 'string' | 'bool' | 'void';

funcDef:
	'func' ID '(' paramDeclaration* ')' ':' var_type varDeclaration* '{' stmt+ funcEnd;
funcEnd: '}';

paramDeclaration: 'var' var_type ID ';';

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
	ID '=' (var_cons | listAssignment | matrixAssignment);
nested_stmt: (
		(
			assignment
			| print
			| specialFunction
			| functionCall
			| returnStmt
		) ';'
	)
	| conditional
	| whileLoop
	| comment;

matrixAssignment: '[' listAssignment (',' listAssignment)* ']';
listAssignment: '[' var_cons (',' var_cons)* ']';

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
	'-'? (
		( '(' expression ')')
		| indexing
		| specialFunction
		| functionCall
		| var_cons
	);

functionCall: ID '(' (expression (',' expression)*)? ')';
indexing: ID '[' expression ']' ('[' expression ']')?;
ID: ('_' | Alpha)+ (Alpha | NUMBER | '_')*;
var_cons: STRING_CONS | FLOAT_CONS | INT_CONS | BOOL_CONS | ID;
INT_CONS: NUMBER+;
NUMBER: '0' .. '9';
BOOL_CONS: 'true' | 'false';
STRING_CONS: '"' (~'"')* '"';
FLOAT_CONS: NUMBER+ '.' NUMBER+;
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
