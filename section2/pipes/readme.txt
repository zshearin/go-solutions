uses pipes to count the number of occurrences of a string 

commands to run:

input:
echo "some string" | go run pipes.go

output:
string: 1
some: 1


input:
echo "test case test" | go run pipes.go

output:
test: 2
case: 1

