#!/bin/bash

suffix='.python3-4.python3'

export CVSROOT=:pserver:anonymous@cvs.debian.org:/cvs/benchmarksgame
cvs checkout benchmarksgame/bench

for file in $(find benchmarksgame -name "*${suffix}")
do
	mv -v $file ./test/$(basename $file $suffix).py
done
