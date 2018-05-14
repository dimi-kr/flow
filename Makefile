
ANTLR4 = java -Xmx500M -cp "bin/antlr-4.7.1-complete.jar:${CLASSPATH}" org.antlr.v4.Tool
GRUN = java -Xmx500M -cp "bin/antlr-4.7.1-complete.jar:./parser/:${CLASSPATH}" org.antlr.v4.gui.TestRig

default: build test

build:
	    $(ANTLR4) -Dlanguage=Go ./parser/Flow.g4
	    go build flow.go

build-test:
		$(ANTLR4) ./parser/Flow.g4
		CLASSPATH="./bin/antlr-4.7.1-complete.jar:./parser/" javac ./parser/Flow*.java
test: build-test
		$(GRUN) Flow flows -tree ./test.flow

test-gui: build-test
		$(GRUN) Flow flows -gui ./test.flow

run: build
		./flow test.flow

