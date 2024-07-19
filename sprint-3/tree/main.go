package main

/*
Course `Web services on Go`, week 1, homework, `tree` program.
See: week_01\materials.zip\week_1\99_hw\tree

mkdir -p week01_homework/tree
pushd week01_homework/tree
go mod init tree
go mod tidy
pushd ..
go work init
go work use ./tree/
go vet tree
gofmt -w tree
go test -v tree
go run tree . -f
go run tree ./tree/testdata
cd tree && docker build -t mailgo_hw1 .

https://en.wikipedia.org/wiki/Tree_(command)
https://mama.indstate.edu/users/ice/tree/
https://stackoverflow.com/questions/32151776/visualize-tree-in-bash-like-the-output-of-unix-tree

*/

import (
	"fmt"
	"io"
	"os"
	"strings"
	"slices"
)

/*
	Example output:

	├───project
	│	└───gopher.png (70372b)
	├───static
	│	├───a_lorem
	│	│	├───dolor.txt (empty)
	│	├───css
	│	│	└───body.css (28b)
	...
	│			└───gopher.png (70372b)

	- path should point to a directory,
	- output all dir items in sorted order, w/o distinction file/dir
	- last element prefix is `└───`
	- other elements prefix is `├───`
	- nested elements aligned with one tab `	` for each level
*/

const (
	EOL             = "\n"
	BRANCHING_TRUNK = "├───"
	LAST_BRANCH     = "└───"
	TRUNC_TAB       = "│\t"
	LAST_TAB        = "\t"
	EMPTY_FILE      = "empty"
	ROOT_PREFIX     = ""

	USE_RECURSION_ENV_KEY = "RECURSIVE_TREE"
	USE_RECURSION_ENV_VAL = "YES"
)

func main() {
	// This code is given
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage: go run main.go . [-f]")
	}

	out := os.Stdout
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"

	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

// dirTree: `tree` program implementation, top-level function, signature is fixed.
// Write `path` dir listing to `out`. If `prinFiles` is set, files is listed along with directories.

func dirTreeSup(out io.Writer, path string, mask []byte, printFiles bool) error{
	curDir, err := os.ReadDir(path)
	fmt.Println(path)
	fmt.Println(curDir)
	fmt.Println(err)

	if err != nil {
		return err
	}
	if len(curDir) == 0 {
		return fmt.Errorf("Empty dir: %s", path)
	}

	var lastEl int

	if !printFiles {
		for i:=len(curDir)-1; i !=0 ; i-- {
			if curDir[i].IsDir() {
				lastEl = i
				break
			}
		}
	} else {
		lastEl = len(curDir) -1
	}
	for _, i := range curDir[:lastEl] {
		if i.IsDir() {
			out.Write(slices.Concat(mask, []byte(BRANCHING_TRUNK), []byte(i.Name()),
				[]byte(EOL)))
			dirTreeSup(out, path+i.Name()+"/", append(mask,[]byte(TRUNC_TAB)...), printFiles)

		} else if printFiles{

			tbArr := slices.Concat( mask, []byte(BRANCHING_TRUNK),[]byte(i.Name()))
			tmp, _ := i.Info()
			size := tmp.Size()

			if size != 0 {
				tbArr = slices.Concat(tbArr, []byte(fmt.Sprintf(" (%db)", size)),[]byte(EOL))
			} else {
				tbArr = slices.Concat(tbArr, []byte(fmt.Sprint(" (empty)")),[]byte(EOL))
			}

			out.Write(tbArr)
			tbArr = nil
		}
	}


	if curDir[lastEl].IsDir() {

		out.Write(slices.Concat(mask, []byte(LAST_BRANCH), []byte(curDir[lastEl].Name()), []byte(EOL)))
		dirTreeSup(out, path+curDir[lastEl].Name() + "/", append(mask, []byte(LAST_TAB)...), printFiles)

	} else if printFiles{

		tbArr := slices.Concat( mask, []byte(LAST_BRANCH),[]byte(curDir[lastEl].Name()))
		tmp, _ := curDir[lastEl].Info()
		size := tmp.Size()
		if (size != 0) {
			tbArr = slices.Concat(tbArr, []byte(fmt.Sprintf(" (%db)", size)), []byte(EOL))
		} else {
			tbArr = slices.Concat(tbArr, []byte(fmt.Sprint(" (empty)")), []byte(EOL))
		}
		out.Write(tbArr)
		tbArr = nil
	}
	return nil
}
func dirTree(out io.Writer, path string, printFiles bool) error {
	path+="/"
	if _, ok := os.ReadDir(path); ok != nil {
		return nil
	}

	var sb  strings.Builder
	sb.WriteString(path)

	dirTreeSup(out, path, []byte{}, printFiles)
	return nil
}

