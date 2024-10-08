# MANGO
This document contains the syntax rules used in a mango.build or MANGO.build file

# Warning
**The examples in this file are for educational purposes and are not all valid syntax**  
**They are here to highlight differences or give an indication on how the structure looks**

## Targets
Targets in MANGO can be specified with the "target" keyword followed by the name of the target followed by curly braces  
Structure: target <target name\>{}  
Examples:
```MANGO syntax
target install{
	# do installation things here
}
```

## Variables

### Intializing a variable
Initializing variables can be done anywhere outside a [target block]() by stating the name of the variable followed by the = sign followed by a value  
Structure: <Variable name\> = <Value\>  
Examples:
```MANGO syntax
CC = gcc
CXX = g++
GO = go
```

### Calling a variable
A variable may be called at any point in the build script with a \$ sign followed by the variable name in curly braces  
Structure: ${<Variable name\>}  
Examples:
```MANGO Syntax
${CC} file.c -o file.o
${CXX} file.cpp -o file.o
${GO} go run file.go
```

## Comments
In MANGO lines that start with the # symbol are not parsed, meaning that their content is not read or executed     
Structure: # <Content of the comment\>  
Examples: 
```MANGO syntax
# this line gets ignored
this line does not
```

## Commands
Currently in MANGO there is only one type of command user made commands  
User made commands may only exist inside a [target block]() and must be started by a - symbol  
Structure: - <command\>  
Examples:
```MANGO syntax
- gcc file.c -o file.o
- g++ file.c -o file.o
```
