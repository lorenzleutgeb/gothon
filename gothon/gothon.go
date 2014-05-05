package main

import (
	"flag"
	"log"
	"os"
	"bufio"
	"fmt"
	"github.com/flowlo/gothon"
	"os/exec"
	"path"
//	"encoding/json"
)

var (
	verbose = flag.Bool("v", false, "verbose. if set, gothon will tell you what is going on.")
	debug = flag.Bool("d", true, "debug. if set, gothon will tell precicesly what it is doing. implies verbose")	
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [parameters] <filename>\n\nAvailable parameters:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprint(os.Stderr, "\nThis is an experimental interpreter for Python written by Lorenz Leutgeb.\n")
	}

	flag.Parse()

	if *debug { *verbose = true }

	if len(flag.Args()) != 1 {
		flag.Usage()
		os.Exit(1)
	}
	
	var target string = flag.Args()[0]

	if !path.IsAbs(target) {
		pwd, _ := os.Getwd()
		target = path.Join(pwd, target)
	}

	fmt.Println(target)

	if path.Ext(target) == ".py" {
		_, err := exec.LookPath("python3.3")
		if err != nil { log.Fatalf("If you want gothon to compile \"%s\", please install Python 3.3 (Could not find \"python3.3\")", path.Base(target)) }
		if *debug { log.Print("Found Python 3.3") }

		cmd := exec.Command("python3.3", "-m", "compileall", target)
		output, err := cmd.CombinedOutput()
		if err != nil { log.Fatal("compileall: ", err) }
		if len(output) > 0 && *debug { log.Print(string(output)) }
		if *verbose { log.Print("Compilation finished") }
		base := path.Base(target)
		target = path.Join(path.Dir(target), "__pycache__", base[:len(base) - 2] + "cpython-33.pyc")
	}

	var file *os.File

	file, err := os.Open(target)

	if err != nil { log.Fatal(err) }
	if *verbose { log.Print("Using file \"", target, "\"") }

	reader := gothon.Reader{*bufio.NewReader(file)}

	module := &gothon.Module{}
	
	module.Read(&reader)
	
/*	dump, err := json.Marshal(module)
	if err != nil { log.Fatal(err) }
	log.Printf("%s", dump) */

	file.Close()
	
	machine := &gothon.Machine{}
	machine.Execute(*module)
}
