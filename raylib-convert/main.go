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
	"time"
)

var (
	fileSuffix        = flag.String("suffix", "_gen", "the suffic to append to each file")
	format            = flag.Bool("format", true, "run gofmt to the resulting output")
	input             = flag.String("in", "HeaderList.h", "raylib-convert compatible header file")
	ignoreOOPFile     = flag.String("ioop", "StructList.txt", "file that contains a list of types that cannot be OOP")
	manualDir         = flag.String("manual", "manual/", "directory that stores the manual files")
	output            = flag.String("out", "out/", "the output directory")
	functionalConvert = flag.Bool("use_func", true, "tells the converter to use newTypeFromPointer and cptr() functions")
	oopOnly           = flag.Bool("oop_only", false, "should only the OOP version of the function be generated?")
	trackUnloadables  = flag.Bool("track_unloadables", true, "should unloadables track when they are being loaded and unloaded to our list. Only applicable with OOP")
)

var ignoreOOPs []string
var patterns []matchPattern
var enums []matchEnum

type matchPattern struct {
	pattern *regexp.Regexp
	replace string
}

type matchEnum struct {
	pattern *regexp.Regexp
	enum    string
}

func (p *matchPattern) replaceAll(s string) string {
	return p.pattern.ReplaceAllString(s, p.replace)
}

func main() {

	if _, err := os.Stat(*output); os.IsNotExist(err) {
		os.Mkdir(*output, os.ModePerm)
	}

	//Parse the flags
	flag.Parse()

	//Read the ignore FileExists
	ifile, ierr := os.Open(*ignoreOOPFile)
	ignoreOOPs = make([]string, 0)
	if ierr == nil {
		defer ifile.Close()

		//Read each line of the input file and generate the prototypes
		iscanner := bufio.NewScanner(ifile)
		for iscanner.Scan() {
			ignoreOOPs = append(ignoreOOPs, iscanner.Text())
		}
	}

	//Read each line of the headers now
	file, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}

	//Defer the close until the end
	defer file.Close()

	successTally := 0
	failureTally := 0

	//PRepare a bunch of variables to hold our succes, failed and prototypes
	prototypes := make([]*prototype, 0)
	failed := make([]string, 0)
	success := make([]string, 0)
	patterns = make([]matchPattern, 0)
	enums = make([]matchEnum, 0)

	defaultHeader := "//Generated " + time.Now().Format(time.RFC3339) + "\n#include \"raylib.h\"\n#include <stdlib.h>\n#include \"go.h\"\n"

	fileHeader := defaultHeader
	filenameSuccess := "main" + *fileSuffix + ".go"
	filenameFailed := "main" + *fileSuffix + ".failed"

	ignoring := false
	asOOP := false

	//Read each line of the input file and generate the prototypes
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		//Read the line
		line := strings.Trim(scanner.Text(), " ")

		//We are special instructions
		if strings.HasPrefix(line, "#convert_") {
			fmt.Println("Processing Command")
			parts := strings.Split(line, "_")
			if len(parts) > 0 {
				switch parts[1] {
				default:

					//A file group
				case "group":

					//Write the old filename and recrate the values
					failedResults := strings.Join(failed, "\n")
					sucessResults := "package raylib\n/*\n" + fileHeader + "*/\nimport \"C\"\nimport \"unsafe\"\n" + strings.Join(success, "\n")
					saveProgress(filenameFailed, filenameSuccess, sucessResults, failedResults)

					//Clear previous arrays
					fileHeader = defaultHeader
					failed = make([]string, 0)
					success = make([]string, 0)
					patterns = make([]matchPattern, 0)
					enums = make([]matchEnum, 0)

					//Prepare the new filename
					filenameSuccess = parts[2] + *fileSuffix + ".go"
					filenameFailed = parts[2] + *fileSuffix + ".failed"

					//A new header line
				case "cgo":
					fileHeader = fileHeader + strings.Join(parts[2:], ":") + "\n"
					fmt.Println("Header: " + fileHeader)

				case "ignore":
					ignoring = parts[2] == "begin"

				case "oop":
					asOOP = parts[2] == "begin"

				case "replace":
					patterns = append(patterns, matchPattern{
						pattern: regexp.MustCompile(parts[2]),
						replace: parts[3],
					})

				case "enum":
					//conv:enum:/int gamepad/g:GamepadNumber
					enums = append(enums, matchEnum{
						pattern: regexp.MustCompile(parts[2]),
						enum:    parts[3],
					})
				}

			}

			//Skip because its a command line
			continue
		}

		//We are ignoring, so skip the lines
		if ignoring {
			continue
		}

		//Process the line
		p, err := parseLine(line)

		if err != nil {
			//Failed to parse the file
			fmt.Println("Failed: ", line)
			failed = append(failed, "\n//"+err.Error()+"\n"+line)
		} else {
			//We parsed it, but comments return no errors and nil prototypes.
			if p != nil {

				//Add the prototype to our list of processed.
				prototypes = append(prototypes, p)

				//Translate it. If we are successful then add it to our success list,
				// otherwise add it to our fail list
				trans, terr := translatePrototype(p, asOOP)
				if terr == nil {
					success = append(success, trans)
					successTally++
				} else {
					fmt.Println("Failed: ", line)
					failed = append(failed, "\n//"+terr.Error()+"\n"+line)
					failureTally++
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//Write the old filename and recrate the values
	failedResults := strings.Join(failed, "\n")
	sucessResults := "package raylib\n/*\n" + fileHeader + "*/\nimport \"C\"\nimport \"unsafe\"\n" + strings.Join(success, "\n")
	saveProgress(filenameFailed, filenameSuccess, sucessResults, failedResults)

	//Complete
	fmt.Println("Completed ", successTally, " / ", len(prototypes), " functions (", (float64(successTally) / float64(len(prototypes)) * 100), "% Yield)")
}

func saveProgress(filenameFailed string, filenameSuccess string, successResults string, failureResults string) {

	//Write the failures
	if len(failureResults) > 0 {
		ioutil.WriteFile(*output+"/"+filenameFailed, []byte(failureResults), 0644)
	}

	if *format {
		fmt.Println("Formatting...")
		cmd := exec.Command("goimports")
		cmd.Stdin = strings.NewReader(successResults)
		var out bytes.Buffer
		cmd.Stdout = &out
		if err := cmd.Run(); err != nil {
			fmt.Println("Failed to format!", err)
			ioutil.WriteFile(*output+"/"+filenameSuccess, []byte(successResults), 0644)
		} else {
			ioutil.WriteFile(*output+"/"+filenameSuccess, out.Bytes(), 0644)
		}
	}

}

func translatePrototype(prototype *prototype, objectOriented bool) (string, error) {

	//We have a manual definition, so use that instead
	if _, err := os.Stat(*manualDir + prototype.name + ".go"); err == nil {
		bt, fe := ioutil.ReadFile(*manualDir + prototype.name + ".go")
		return "\n" + string(bt), fe
	}

	//We do not support return types really yet, but when we do we have a special case for pointers
	if prototype.returnArg.GetPraticalPointerDepth() >= 1 {
		return "", errors.New("cannot process pointer return types")
	}

	isOOP := false
	if len(prototype.args) >= 1 && prototype.args[0] != nil &&
		isObject(prototype.args[0].valueType) && objectOriented &&
		!strings.Contains(prototype.name, "Load") &&
		!contains(ignoreOOPs, prototype.args[0].valueType) {
		fmt.Println("OOP: ", prototype.name)
		isOOP = true
	}

	//Prepare some variables
	argHeaders := make([]string, len(prototype.args))
	bodyArgs := make([]string, len(prototype.args))
	argNames := make([]string, len(prototype.args))
	bodyArgsTally := 0
	returnHeaders := make([]string, 0)
	returnExpre := make([]string, 0)
	returnPointer := false

	body := "C." + prototype.name + "("

	//Add the first item to the return headers
	if prototype.returnArg.valueType != "void" {
		spacing := " "

		if isReferencedObject(prototype.returnArg.valueType) {
			spacing = " *"
			returnPointer = true
		}

		returnHeaders = append(returnHeaders, spacing+convertType(prototype.returnArg.valueType, prototype.returnArg.unsigned))
		returnExpre = append(returnExpre, castToGo("res", prototype.returnArg.valueType, prototype.returnArg.HasPointer(), prototype.returnArg.unsigned))
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

		spacing := " "
		if isReferencedObject(arg.valueType) && objectOriented {
			spacing = " *"
		}

		//Update our name
		//Append to the header
		argNames[i] = arg.name
		argHeaders[i] = arg.name + spacing + convertType(arg.valueType, arg.unsigned)
		if arg.enumType != "" {
			//We have an enum type, so we will update our argHeader to use that instead
			argHeaders[i] = arg.name + spacing + arg.enumType
		}

		bodyArgPart, bodyPrefixPart, pointerless := castToC(*arg)

		//We dont want to do this if we are OOP and its the first item
		if (!isOOP || i > 0) && arg.GetPraticalPointerDepth() == 1 {

			if !objectOriented {
				spacing = ""
			}

			returnHeaders = append(returnHeaders, spacing+convertType(arg.valueType, arg.unsigned))
			returnExpre = append(returnExpre, castToGo(bodyArgPart, arg.valueType, arg.HasPointer(), arg.unsigned))

			if !pointerless {
				bodyArgPart = "&" + bodyArgPart
			}
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

		//If we are enabling unloadable tracking, then we need to add the addUnloadable to the end
		// before returning
		if objectOriented && *trackUnloadables && (returnPointer || strings.HasPrefix(prototype.name, "Load")) {
			returnFooter = "\nretval := " + returnExpre[0]
			returnExpre[0] = "retval"

			returnFooter += "\nRegisterUnloadable(retval)"
			returnFooter += "\nreturn " + strings.Join(returnExpre, ", ")

		} else {
			returnFooter = "\nreturn " + strings.Join(returnExpre, ", ")
		}
	}

	//If we are enabling unloadable tracking and this is an unload, then lets unload
	isUnload := *trackUnloadables && strings.HasPrefix(prototype.name, "Unload")
	if isUnload && isOOP {
		body += "\nUnregisterUnloadable(" + argNames[0] + ")"
	}

	//Finally combine the body
	body = body + returnFooter
	oopName := prototype.name

	//The definition is what we will write, including comment at the top
	definition := ""
	if isOOP {

		//Prepare the function name
		argName := prototype.args[0].valueType
		oopName = strings.Replace(oopName, "From"+argName, "", 1)
		oopName = strings.Replace(oopName, argName, "", 1)
		if isUnload {
			oopName = "Unload"
		}

		spacing := " *"
		if !isReferencedObject(prototype.args[0].valueType) {
			spacing = " "
		}

		retName := prototype.args[0].name + spacing + prototype.args[0].valueType

		//add the comment and the line
		definition += "//" + oopName + " : " + prototype.comment + "\n"
		definition += fmt.Sprintf("func (%s) %s(%s) (%s) {\n %s \n}\n", retName, oopName, strings.Join(argHeaders[1:], ", "), strings.Join(returnHeaders, ", "), body)
	}

	if !*oopOnly || !isOOP {

		//Add the top comment
		definition += "//" + prototype.name + " : " + prototype.comment + "\n"

		//We are just going to call the OOP function
		if isOOP && !*oopOnly {

			//Prepare the new body
			body = fmt.Sprintf("%s.%s(%s)", prototype.args[0].name, oopName, strings.Join(argNames[1:], ", "))
			if len(returnFooter) > 0 {
				body = "return " + body
			}

			//Add an recommendation comments
			definition += fmt.Sprintf("//Recommended to use %s.%s(%s) instead\n", prototype.args[0].name, oopName, strings.Join(argNames[1:], ", "))
		}

		//Add the definition
		definition += fmt.Sprintf("func %s(%s) (%s) {\n %s \n}\n", prototype.name, strings.Join(argHeaders, ", "), strings.Join(returnHeaders, ", "), body)
	}

	return definition, nil
}

//castType creates a cast for a type, returning first the name of the variable and then the definition of the variable.
// There are some cases where there is no definition.
func castToC(a argument) (string, string, bool) {
	csname := "c" + a.name

	switch a.valueType {
	default:
		deref := "*"
		if a.HasPointer() {
			deref = ""
		}

		if *functionalConvert {
			return csname, csname + " := " + deref + a.name + ".cptr()", true
		}

		//func (v *Vector3) cptr() *C.Vector3 { return (*C.Vector3)(unsafe.Pointer(v)) 	}
		return csname, csname + " := " + deref + "( (*C." + a.valueType + ")(unsafe.Pointer(&" + a.name + ")) )", true
	case "int":
		//WE are unsigned, so do a converntional convert
		if a.unsigned {
			if a.GetPraticalPointerDepth() == 1 {
				return csname, csname + " := C.u" + a.valueType + "(" + a.name + ")", false
			}
			return "C.u" + a.valueType + "(" + a.name + ")", "", false
		}

		//We are not unsigned, so proceed as normal
		if a.GetPraticalPointerDepth() == 1 {
			return csname, csname + " := C." + a.valueType + "(int32(" + a.name + "))", false
		}
		return "C." + a.valueType + "(int32(" + a.name + "))", "", false
	case "float":
		fallthrough
	case "uint8":
		fallthrough
	case "bool":
		if a.GetPraticalPointerDepth() == 1 {
			return csname, csname + " := C." + a.valueType + "(" + a.name + ")", false
		}
		return "C." + a.valueType + "(" + a.name + ")", "", false
	case "void":
		return a.name, "", true
	case "char":
		return csname, csname + " := C.CString(" + a.name + ")\ndefer C.free(unsafe.Pointer(" + csname + "))", true
	}
}

//casts a c type to a go type
func castToGo(variable, t string, isPointer bool, unsigned bool) string {
	addr := ""
	if !isPointer {
		addr = "&"
	}

	switch t {
	case "int":
		if unsigned {
			return convertType(t, unsigned) + "(" + variable + ")"
		}
		return convertType(t, false) + "(int32(" + variable + "))"
	case "float":
		fallthrough
	case "double":
		fallthrough
	case "uint8":
		fallthrough
	case "bool":
		return convertType(t, unsigned) + "(" + variable + ")"
	case "char":
		return "C.GoString(" + variable + ")"
	case "void":
		return "unsafe.Pointer(" + addr + variable + ")"
	default:
		//func newRectangleFromPointer(ptr unsafe.Pointer) Rectangle { return *(*Rectangle)(ptr) }
		if *functionalConvert {
			return "new" + convertType(t, unsigned) + "FromPointer(unsafe.Pointer(" + addr + variable + "))"
		}

		return "*(*" + convertType(t, unsigned) + ")(unsafe.Pointer(" + addr + variable + "))"
	}
}

//convertType converts a c type to a go type
func convertType(t string, unsigned bool) string {
	switch t {
	default:
		if unsigned {
			return "u" + t
		}
		return t
	case "float":
		return "float32"
	case "char":
		return "string"
	case "void":
		return "unsafe.Pointer"
	case "double":
		return "float64"
	case "int":
		if unsigned {
			return "uint32"
		}
		return "int"
	}
}

func isReferencedObject(t string) bool {
	if !isObject(t) || t == "Texture2D" || t == "RenderTexture2D" || t == "Shader" {
		return false
	}

	return true
}

func isObject(t string) bool {
	switch t {
	default:
		return !contains(ignoreOOPs, t)
	case "float":
		fallthrough
	case "int":
		fallthrough
	case "char":
		fallthrough
	case "void":
		fallthrough
	case "double":
		fallthrough
	case "bool":
		return false
	}
}

//Parses a line and generates a prototype
func parseLine(line string) (*prototype, error) {
	//Trim the line and validate it
	line = strings.Trim(line, " ")
	if len(line) < 4 || strings.HasPrefix(line, "//") || strings.HasPrefix(line, "#") {
		return nil, nil
	}

	rePrototype := regexp.MustCompile(`(RAYGUIDEF|RLAPI) (const |unsigned )?([a-zA-Z0-9]+) (\**)([a-zA-Z0-9]+)\s?\(([^!@#$+%^]+?)\);\s*(\/\/(.*))?`)
	reArgument := regexp.MustCompile(`(const |unsigned )?([a-zA-Z0-9]+) (\**)([a-zA-Z0-9]+)`)

	matches := rePrototype.FindAllStringSubmatch(line, -1)
	if len(matches) != 1 {
		return nil, errors.New("invalid amount of matches for header")
	}

	//Prepare the prototype
	proto := &prototype{
		entire: matches[0][0],
		returnArg: argument{
			name:         "return",
			constant:     strings.Trim(matches[0][2], " ") == "const",
			unsigned:     strings.Trim(matches[0][2], " ") == "unsigned",
			valueType:    matches[0][3],
			enumType:     "",
			pointerDepth: len(strings.Trim(matches[0][4], " ")),
		},
		name:    matches[0][5],
		comment: strings.Trim(matches[0][7], " /"),
	}

	//Prepare the arguments
	argline := matches[0][6]
	for _, p := range patterns {
		argline = p.replaceAll(argline)
	}

	//Split it up
	parts := strings.Split(argline, ",")
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

		//Prepare the enum type
		enumLine := proto.name + "(" + p + ")"
		enumType := ""
		for _, et := range enums {
			if et.pattern.MatchString(enumLine) {
				enumType = et.enum
				break
			}
		}

		arguments[i] = &argument{
			entire:       matches[0][0],
			constant:     strings.Trim(matches[0][1], " ") == "const",
			unsigned:     strings.Trim(matches[0][1], " ") == "unsigned",
			valueType:    matches[0][2],
			enumType:     enumType,
			pointerDepth: len(strings.Trim(matches[0][3], " ")),
			name:         name,
		}

		i++
	}

	proto.args = arguments
	return proto, nil
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
	enumType     string
	pointerDepth int
	constant     bool
	unsigned     bool
}

func (p *argument) HasPointer() bool { return p.pointerDepth > 0 }
func (p *argument) GetPraticalPointerDepth() int {
	if (p.constant && p.valueType == "char") || p.valueType == "void" {
		return p.pointerDepth - 1
	}
	return p.pointerDepth
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
