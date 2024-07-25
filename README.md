
# ASCII ART OUTPUT
The ASCII Art Output program written in go language, is designed to take a string as an argument and output in a graphic representation, is written out to a file. This graphical representation is created by rendering the input string with ASCII characters to form a visual representation.

File Structure: main folder(asciiart), contains(3 banner files; shadow.txt,standard.txt and thinkertoy.txt, a ReadMe.md file containing the documantaion of the program, 5 function folders with their test files and a go.mod file that contains modules that help in running the program )


## Demo

Usage:
+ To run the program, follow these steps:

1. Clone the repository:

```bash
    git clone https://learn.zone01kisumu.ke/git/bshisia/ascii-art-output
```

2. Navigate to the program folder by copying the command below to the terminal:

```bash
    cd ascii-art-output
```

3. Run the program using the following command/ copy the command below to the terminal:
```bash
    Usage: go run . [OPTION] [STRING] [BANNER]

    EX: go run . --output=<fileName.txt> something standard
```

4. To display the contents of the file on the terminal, run:
```bash
    cat banner.txt
```
                
Example:
For example, running the program with the commands below:

command 1.

```bash
go run . --output=banner.txt "Hello Alice!" shadow
```

The ouput written to the file will be an ASCII art representation using the shadow banner style as shown below:
```bash
                                                                           
_|    _|          _| _|                  _|_|   _| _|                   _| 
_|    _|   _|_|   _| _|   _|_|         _|    _| _|      _|_|_|   _|_|   _| 
_|_|_|_| _|_|_|_| _| _| _|    _|       _|_|_|_| _| _| _|       _|_|_|_| _| 
_|    _| _|       _| _| _|    _|       _|    _| _| _| _|       _|          
_|    _|   _|_|_| _| _|   _|_|         _|    _| _| _|   _|_|_|   _|_|_| _| 
                                                                           
```

command 2.

```bash
go run . --output=banner.txt "Hello Brian!" thinkertoy
```

The ouput written to the file will be an ASCII art representation using the shadow banner style as shown below:
```bash
                                               
o  o     o o           o--o                  o 
|  |     | |           |   |     o           | 
O--O o-o | | o-o       O--o  o-o    oo  o-o  o 
|  | |-' | | | |       |   | |   | | |  |  |   
o  o o-o o o o-o       o--o  o   | o-o- o  o O 

``` 
command 3.

```bash
go run . --output=banner.txt "Hello Joe!" standard
```

The ouput written to the file will be an ASCII art representation using the shadow banner style as shown below:
```bash

 _    _          _   _                      _                  _  
| |  | |        | | | |                    | |                | | 
| |__| |   ___  | | | |   ___              | |   ___     ___  | | 
|  __  |  / _ \ | | | |  / _ \         _   | |  / _ \   / _ \ | | 
| |  | | |  __/ | | | | | (_) |       | |__| | | (_) | |  __/ |_| 
|_|  |_|  \___| |_| |_|  \___/         \____/   \___/   \___| (_) 
  
```
## Running Tests

To run tests, run the following command:
```bash
 go test
```
## Authors

- [@Aliceokingo](https://learn.zone01kisumu.ke/git/aokingo)
- [@joelamos](https://learn.zone01kisumu.ke/git/jamos)
- [@brianshisia](https://learn.zone01kisumu.ke/git/bshisia)
## License
This project is licensed under [MIT](https://learn.zone01kisumu.ke/git/bshisia/ascii-art-output/LICENSE)

