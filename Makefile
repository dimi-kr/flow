
ANTLR4 = java -Xmx500M -cp "bin/antlr-4.7.1-complete.jar:${CLASSPATH}" org.antlr.v4.Tool
GRUN = java -Xmx500M -cp "bin/antlr-4.7.1-complete.jar:${CLASSPATH}" org.antlr.v4.gui.TestRig

build:
	    $(ANTLR4) -Dlanguage=Go ./parser/Flow.g4
	    go build flow.go

build-test:
		$(ANTLR4) ./parser/Flow.g4
		javac ./parser/Flow*.java
test: build-test
		$(GRUN) Flow flows -gui ./parses/Flow.g4

run:
		./flow test.flow

