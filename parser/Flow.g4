grammar Flow;

flows: flow+ EOF;
flow: stmt+ END;

stmt: node | edge;

edge: node edge_rhs;
edge_rhs: ( op node )+ | (op node AND node)+;
node: LINK_PREV | LINK_NEXT | ID | question | block | TEXT;
question: '<' STRING as_id? '>';
as_id: as ID;

block: '[' STRING as_id? ']' caption?;
TEXT: '"' STRING '"';
as: ' as '| ':';
with: ' with ';
caption: with TEXT;
AND: ',';
STRING: '"' ( '\\"' | . )*? '"';
op: '->' | '--';
LINK_PREV: '$prev';
LINK_NEXT: '$next';
END: ';';
ID: [a-zA-Z][a-zA-Z0-9]*;

COMMENT : '//' .+? ('\n'|EOF) -> skip;
WS: [ \t\n\r] + -> skip;
