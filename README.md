# cli_with_go
CLI using Cobra and Go <br>
prerequisite: GOLANG must be installed on the machine. <br>
To run the CLI <br> 
1. Clone the project. <br>
2. Open the terminal from the project root directory and run  
```
go install
```

then,open the terminal from any directory and type 
```
ayvu-cli say hello
```
for printing any text 
```
ayvu-cli say hello -n [YOUR NAME or any text]
```
example:
```
ayvu-cli say hello -n shajir
```
or
```
ayvu-cli say hello --name shajir
```
for running a external python file run 
``` 
ayvu-cli run
```
