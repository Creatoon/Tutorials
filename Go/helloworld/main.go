// package == project == workspace
// whenever you see the keyword package think of it like a project, so a package is a collection of common source code files and  these source code files are present to create one package.
// suppose I and You are working on same application then we will be writing many go modules for one single package.
// every file inside pur package must end with [ .go ] extension.
// every files must have their package name included at the top if they corresponds to a same package.
// why package main only ?
// there are 2 types of packages a) executable b) reusable
// so executable is just for compile and execute and must have a func called 'main'
// and reusable package is that can be used as a dependency(helper code)
// so we can understand "main" is standard or sacred and if use other name after keyword package then it will not get excecuted and will be a reusable package
// package === repository, in easy terms
package main

// import statements are just used to import code from other module to the current module in which import statement is used.
// fmt is just a part of standard library in the go, full form FORMAT, which is just used to print out something into the shell in different formats just to give a better sense of debugging.
// other standard libraries of go are :- debug, fmt, io, crypto, encoding, math
// also we can import reusable packages
// there should be a double inverted commas after the import statement
import "fmt"

// so func keyword is used to define a function followed by a function name and then set of arguements and a function body
// There should be a double inverted commas around the string during calling Println method
// Using single quotes will throw an error of "illegal rune literal"
func main() {
	fmt.Println("Hi there!")
}
