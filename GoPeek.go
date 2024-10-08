package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"bufio"
	"io/fs"
	"go/parser"
    "go/token"
    "go/ast"
    "unicode"
    "unicode/utf8"
    "strings"
)

/*
	Reads a specified Go source file
	and prints out details such as:
		Total Line Count,
		Methods (exported/unexported)
*/

var cliVersion = "1.0.3"
var githubURL = "https://github.com/SarryGeezOwO/GoPeek"

func main() {

	// Get cmd Arguments
	if(len(os.Args) < 2) {
		fmt.Println("Usage: GoPeek <targetFile.go>")
		fmt.Println("**Ensure the file exists within the current working directory.")
		os.Exit(1)
	}
	var targetFile string = os.Args[1]


	// Get CD
	var dir string
	var err error
	dir, err = os.Getwd()
	if(err != nil) {
		log.Fatalf("Error in getting current working directory: %v", err);
	}


	// get all info related to the Go file
	var file string    		= findGoFile(dir, targetFile);
	var lineCount int  		= checkLineCount(file);
	var commentCount int  	= getCommentCount(file);
	var methodList []string = getMethodList(targetFile);

	// Print Out File Details
	fmt.Println("--GoPeek--")
	fmt.Println("Version:", cliVersion)
	fmt.Println("Github:", githubURL)

	fmt.Println("\nFile:", file)
	fmt.Println("   --> Line Count:", lineCount)
	fmt.Println("   --> Comments:", commentCount)
	fmt.Println("   --> Methods Found:")
	for _, val := range methodList {
		var modifier string
		
		if isFuncPrivate(val) {
			modifier = "-- Private | Unexported"
		}else {
			modifier = "-- Public  | Exported"
		}

		fmt.Printf("        > %-25s %s\n", val, modifier)
	}
}

// Returns the absolute path of a given source Go File
func findGoFile(dir string, targetFile string) (file string) {

	// Find The first instance of GO file
	var root fs.FS

	// files <- is a string[]
	// fs.Glob() <- in java terms is like GetFileList(dir, pattern)
	root = os.DirFS(dir)
	files, err := fs.Glob(root, "*.go")
	if(err != nil) {
		log.Fatal(err)
	}

	// Find one matching file to the specified target and return
	for _, value := range files {
		if(value == targetFile) {
			file = path.Join(dir, value);
			return;	
		}
	}
	file = " <404 not Found> "
	return;
}

func getCommentCount(filePath string) (count int) {
	file, err := os.Open(filePath)
	if(err != nil) {
		log.Fatal(err)
	}
	defer file.Close() // <- Will be called after the function is done

	// Read file line by line
	scanner := bufio.NewScanner(file)
	multiCommentFlag := false;
	for scanner.Scan() {
		if(strings.Contains(scanner.Text(), "/*")) {
			multiCommentFlag = true;
		}

		if(strings.Contains(scanner.Text(), "*/")) {
			if(multiCommentFlag) {
				count++;
				multiCommentFlag = false;
			}
		}

		if(strings.Contains(scanner.Text(), "//")) {
			count++;
		}
	}


	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func checkLineCount(filePath string) (lineCount int) {
	file, err := os.Open(filePath)
	if(err != nil) {
		log.Fatal(err)
	}
	defer file.Close() // <- Will be called after the function is done

	// Read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func getMethodList(filePath string) (methods []string) {
	file, err := os.Open(filePath)
	if(err != nil) {
		log.Fatal(err)
		return
	}
	defer file.Close()

	// Create a New Scanner and then parse the specified Go file
	fileSet := token.NewFileSet()
	node, err := parser.ParseFile(fileSet, "", file, parser.AllErrors)
	if(err != nil) {
		fmt.Println("Error file Parsing")
		return
	}

	// Traverse the AST using Recurssion and find all methods
	ast.Inspect(node, func(n ast.Node) bool {
		if funcDeclaration, ok := n.(*ast.FuncDecl); ok {
			if(funcDeclaration != nil) {
				methods = append(methods, funcDeclaration.Name.Name)
			}
		}
		return true
	})
	return
}

// Returns true if the function is private false otherwise
func isFuncPrivate(method string) (bool) {
	firstRune, _ := utf8.DecodeRuneInString(method) // First Character in a String
	if unicode.IsLower(firstRune) {
		return true
	}else {
		return false
	}	
}