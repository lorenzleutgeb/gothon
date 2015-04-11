#!/bin/bash
export CVSROOT=:pserver:anonymous@cvs.debian.org:/cvs/benchmarksgame
cvs login
cvs checkout benchmarksgame/bench
