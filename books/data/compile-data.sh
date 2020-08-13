#!/bin/bash

echo "package data" > data.go
echo "" >> data.go

echo "var Booklist = \`" >> data.go
cat _booklist.csv >> data.go
echo "\`" >> data.go

echo "" >> data.go
echo "var Books = initBookMap()" >> data.go
echo "" >> data.go

echo "func initBookMap() map[string]string {" >> data.go
echo "	b := make(map[string]string)" >> data.go

for f in *.txt
do
  echo "" >> data.go
  echo "	b["'"'"$f"'"'"] = \`" >> data.go
  cat "$f" >> data.go
  echo "\`" >> data.go
done

echo "return b" >> data.go
echo "}" >> data.go