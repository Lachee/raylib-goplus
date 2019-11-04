package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var (
	format      = flag.Bool("format", true, "run gofmt to the resulting output")
	input       = flag.String("in", "headers.txt", "raylib-convert compatible header file")
	manualDir   = flag.String("manual", "manual/", "directory that stores the manual files")
	output      = flag.String("out", "out/headers.go", "the output go file")
	errorOutput = flag.String("err", "out/headers.failed.txt", "the output file for hte failed headers")
)

func main() {

	//Parse the flags
	flag.Parse()

	//Read the file
	file, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}

	//Defer the close until the end
	defer file.Close()

	//PRepare a bunch of variables to hold our succes, failed and prototypes
	prototypes := make([]*prototype, 0)
	failed := make([]string, 0)
	success := make([]string, 0)

	//Read each line of the input file and generate the prototypes
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		//Read the line
		line := scanner.Text()
		p, err := parseLine(line)

		if err != nil {
			//Failed to parse the file
			fmt.Println("Failed: ", line, err)
			failed = append(failed, "\n//"+err.Error()+"\n"+line)
		} else {
			//We parsed it, but comments return no errors and nil prototypes.
			if p != nil {

				//Add the prototype to our list of processed.
				prototypes = append(prototypes, p)

				//Translate it. If we are successful then add it to our success list,
				// otherwise add it to our fail list
				trans, terr := translatePrototype(p)
				if terr == nil {
					success = append(success, trans)
				} else {
					fmt.Println("Failed: ", line, terr)
					failed = append(failed, "\n//"+terr.Error()+"\n"+line)
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//WRite the headers and the failures
	failedResults := strings.Join(failed, "\n")
	sucessResults := "package raylib\n" + strings.Join(success, "\n")

	//Write the failures
	ioutil.WriteFile(*errorOutput, []byte(failedResults), 0644)

	if *format {
		fmt.Println("Formatting...")
		cmd := exec.Command("gofmt")
		cmd.Stdin = strings.NewReader(sucessResults)
		var out bytes.Buffer
		cmd.Stdout = &out
		if err := cmd.Run(); err != nil {
			fmt.Println("Failed to format!", err)
			ioutil.WriteFile(*output, []byte(sucessResults), 0644)
		} else {
			ioutil.WriteFile(*output, out.Bytes(), 0644)
		}
	}

	fmt.Println("Completed ", len(success), " / ", len(prototypes), " functions (", (float64(len(success)) / float64(len(prototypes)) * 100), "% Yield)")
}

func translatePrototype(prototype *prototype) (string, error) {

	//We have a manual definition, so use that instead
	if _, err := os.Stat(*manualDir + prototype.name + ".go"); err == nil {
		bt, fe := ioutil.ReadFile(*manualDir + prototype.name + ".go")
		return "\n" + string(bt), fe
	}

	//We do not support return types really yet, but when we do we have a special case for pointers
	if prototype.returnArg.GetPraticalPointerDepth() >= 1 {
		return "", errors.New("cannot process pointer return types")
	}

	//Prepare some variables
	argHeaders := make([]string, len(prototype.args))
	bodyArgs := make([]string, len(prototype.args))
	bodyArgsTally := 0
	returnHeaders := make([]string, 0)
	returnExpre := make([]string, 0)

	body := "C." + prototype.name + "("

	//Add the first item to the return headers
	if prototype.returnArg.valueType != "void" {
		returnHeaders = append(returnHeaders, convertType(prototype.returnArg.valueType))
		returnExpre = append(returnExpre, castToGo("&res", prototype.returnArg.valueType))

		body = "res := " + body
	}

	//Convert the arguments into their headers
	for i, arg := range prototype.args {
		if arg == nil {
			continue
		}

		//Make sure it's a valid type
		if arg.GetPraticalPointerDepth() > 1 {
			return "", errors.New("cannot process pointer of pointer arg types")
		}

		//Append to the header
		argHeaders[i] = arg.name + " " + convertType(arg.valueType)
		bodyArgPart, bodyPrefixPart := castToC(*arg)

		if arg.GetPraticalPointerDepth() == 1 {
			returnHeaders = append(returnHeaders, convertType(arg.valueType))
			returnExpre = append(returnExpre, castToGo(bodyArgPart, arg.valueType))
			bodyArgPart = "&" + bodyArgPart
		}

		//Append to C function header
		bodyArgs[bodyArgsTally] = bodyArgPart
		bodyArgsTally++

		//If we have a definition, then prepend it to the body
		if len(bodyPrefixPart) > 0 {
			body = bodyPrefixPart + "\n" + body
		}
	}

	//Prepare the return headers

	//Finish the body and add everythign back
	body = body + strings.Join(bodyArgs, ", ") + ")"
	returnFooter := ""
	if len(returnExpre) > 0 {
		returnFooter = "return " + strings.Join(returnExpre, ", ")
	}

	//Prepare the function
	text := "func " + prototype.name + "(" + strings.Join(argHeaders, ", ") + ") (" + strings.Join(returnHeaders, ", ") + ") {\n" + body + "\n" + returnFooter + "\n}"
	return "//" + prototype.name + " : " + prototype.comment + "\n" + text, nil
}

//castType creates a cast for a type, returning first the name of the variable and then the definition of the variable.
// There are some cases where there is no definition.
func castToC(a argument) (string, string) {
	csname := "c" + a.name

	switch a.valueType {
	default:
		deref := "*"
		if a.HasPointer() {
			deref = ""
		}
		return csname, csname + " := " + deref + a.name + ".cptr()"
	case "float":
		fallthrough
	case "int":
		fallthrough
	case "uint8":
		fallthrough
	case "bool":
		if a.GetPraticalPointerDepth() == 1 {
			return csname, csname + " := C." + a.valueType + "(" + a.name + ")"
		}

		return "C." + a.valueType + "(" + a.name + ")", ""
	case "void":
		return a.name, ""
	case "char":
		return csname, csname + " := C.CString(" + a.name + ")\ndefer C.free(unsafe.Pointer(" + csname + "))"
	}
}

//casts a c type to a go type
func castToGo(variable, t string) string {
	switch t {
	case "float":
		fallthrough
	case "double":
		fallthrough
	case "int":
		fallthrough
	case "uint8":
		fallthrough
	case "bool":
		return convertType(t) + "(" + variable + ")"
	case "char":
		return "C.GoString(" + variable + ")"
	case "void":
		return "unsafe.Pointer(" + variable + ")"
	default:
		return "new" + convertType(t) + "FromPointer(unsafe.Pointer(" + variable + "))"
	}
}

//convertType converts a c type to a go type
func convertType(t string) string {
	switch t {
	default:
		return t
	case "float":
		return "float32"
	case "char":
		return "string"
	case "void":
		return "unsafe.Pointer"
	case "double":
		return "float64"
	}
}

//Parses a line and generates a prototype
func parseLine(line string) (*prototype, error) {
	//Trim the line and validate it
	line = strings.Trim(line, " ")
	if len(line) < 4 || strings.HasPrefix(line, "//") {
		return nil, nil
	}

	rePrototype := regexp.MustCompile(`(RAYGUIDEF|RLAPI) (const |unsigned )?([a-zA-Z0-9]+) (\**)([a-zA-Z0-9]+)\s?\(([^!@#$+%^]+?)\);\s*(\/\/(.*))?`)
	reArgument := regexp.MustCompile(`(const |unsigned )?([a-zA-Z0-9]+) (\**)([a-zA-Z0-9]+)`)

	matches := rePrototype.FindAllStringSubmatch(line, -1)
	if len(matches) != 1 {
		return nil, errors.New("invalid amount of matches for header")
	}

	//Prepare the prototype
	p := &prototype{
		entire:    matches[0][0],
		returnArg: argument{name: "return", valueType: matches[0][3], pointerDepth: len(strings.Trim(matches[0][4], " "))},
		name:      matches[0][5],
		comment:   strings.Trim(matches[0][7], " /"),
	}

	//Prepare the arguments
	parts := strings.Split(matches[0][6], ",")
	arguments := make([]*argument, len(parts))
	i := 0
	for _, p := range parts {
		matches := reArgument.FindAllStringSubmatch(p, -1)

		if len(matches) != 1 {
			if p == "void" {
				break
			} else {
				return nil, errors.New("invalid amount of matches for arguments")
			}
		}

		name := matches[0][4]
		if name == "type" || name == "interface" || name == "return" {
			name = "g" + name
		}

		arguments[i] = &argument{
			entire:       matches[0][0],
			valueType:    matches[0][2],
			pointerDepth: len(strings.Trim(matches[0][3], " ")),
			name:         name,
		}

		i++
	}

	p.args = arguments
	return p, nil
}

type prototype struct {
	entire    string
	name      string
	args      []*argument
	returnArg argument
	comment   string
}

type argument struct {
	entire       string
	valueType    string
	name         string
	pointerDepth int
}

func (p *argument) HasPointer() bool { return p.pointerDepth > 0 }
func (p *argument) GetPraticalPointerDepth() int {
	if p.valueType == "char" || p.valueType == "void" {
		return p.pointerDepth - 1
	}
	return p.pointerDepth
}
