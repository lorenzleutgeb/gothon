[![Build Status](https://travis-ci.org/flowlo/gothon.svg?branch=master)](https://travis-ci.org/flowlo/gothon) [![Go Report Card](http://goreportcard.com/badge/flowlo/gothon)](http://goreportcard.com/report/flowlo/gothon) [![GoDoc](https://godoc.org/github.com/flowlo/gothon?status.svg)](https://godoc.org/github.com/flowlo/gothon)

# Gothon <img src="gothon.png" alt="Gothon Logo" align="right" />

A Python interpreter written in Go written as part of the course [*Abstract Machines* (german)](http://www.complang.tuwien.ac.at/andi/185966.html) at Vienna University of Technology.

:rotating_light: **ATTENTION:** This codebase is _not maintained_. It only covers a small subset of all Python features. It should not be used in production. It exists for _educational purposes_.

## Scope

Gothon is only an interpreter, not a compiler. It relies on [CPython](https://wiki.python.org/moin/CPython) (the de facto standard Python runtime) to compile human readable source code into [machine readable bytecode](https://docs.python.org/3/library/dis.html).

Gothon is intended to interpret bytecode (which is generated using CPython ahead of time) and actually run the program. It is limited in this aspect on purpose. The goal of the project (at least for now) is not to build a compiler, but a virtual machine that leverages the compiler built into CPython.

## Current Status

Gothon is able to interpret the most basic instructions generated by CPython. This means arithmetic and function calls.

## Future

It would be cool if Gothon would at some point be able to

 * resolve `import` statements.
 * understand classes.
 * deal with exception handling.
 * ...

like a production-grade runtime. Contributions to the project are welcome.

<sub>The Go Gopher was created by [Renée French](http://reneefrench.blogspot.co.at/). The Python Logo is a [registered trademark of the Python Software Foundation](https://www.python.org/psf/trademarks/)</sub>
