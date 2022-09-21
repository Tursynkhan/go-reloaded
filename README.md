# go-reloaded

Foobar is a simple text completion/editing/auto-correction tool.

## Installation

Clone in your folder from git.

```bash
git clone git@git.01.alem.school:tursynkhan/go-reloaded.git
```

## Usage

Program receives as arguments the name of a file containing a text that needs some modifications and the name of the file the modified text should be placed in.   

Example: 
```bash
~/$ cd yourfolder/go-reloaded
~/yourfolder/go-reloaded$ go run . sample.txt result.txt
```
or 


```bash
~/$ cd yourfolder/go-reloaded
~/yourfolder/go-reloaded$ go run main.go mysample.txt myresult.txt
```
The list of modifications this program can do:
- Use ***(hex)*** before the hexadecimal number to replace it with decimal number
- Use ***(bin)*** before the binary number to replace it with decimal number
- Use ***(up)*** before the word you want to Uppercase version
- Use ***(low)*** before the word you want to Lowercase version
- Use ***(cap)*** before the word you want to capitalized version
    - For ***(cap)***, ***(low)***, ***(up)*** if a number appears next to it, like so: ***(low, <number>)*** it turns the previously specified number of words in lowercase, uppercase or capitalized accordingly.
Use it only with one space between **comma** and **number**. 

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## Authors and acknowledgment
@tursynkhan

## License
[MIT](https://choosealicense.com/licenses/mit/)