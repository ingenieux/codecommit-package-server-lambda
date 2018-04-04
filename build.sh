#!/bin/bash

gb build -ldflags='-s -w'

for i in bin/* ; do
	strip $i
	upx $i
done
