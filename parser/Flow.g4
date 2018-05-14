grammar Flow;

flows: flow+ EOF;
flow: stmt+ END;

stmt: node | edge;

edge: node edge_rhs;
edge_rhs: ( op node )+;
node: LINK_PREV | LINK_NEXT | ID | question | block | TEXT;
question: '<' STRING '>';
block: '[' STRING ']';
TEXT: '"' STRING '"';

AND: ',';
STRING: '"' ( '\\"' | . )*? '"';
op: '->' | '--';
LINK_PREV: '$prev';
LINK_NEXT: '$next';
END: ';';
ID: [a-zA-Z][a-zA-Z0-9]*;

COMMENT : '//' .+? ('\n'|EOF) -> skip;
WS: [ \t\n\r] + -> skip;
