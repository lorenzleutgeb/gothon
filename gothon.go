package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"
)

var (
	debug = flag.Bool("d", true, "debug. if set, gothon will tell precicesly what it is doing.")
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [parameters] <filename>\n\nAvailable parameters:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprint(os.Stderr, "\nThis is an experimental interpreter for Python written by Lorenz Leutgeb.\n")
	}

	flag.Parse()

	if len(flag.Args()) != 1 {
		flag.Usage()
		os.Exit(1)
	}

	file, err := resolve(flag.Args()[0])
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	module := &Module{}
	reader := Reader{*bufio.NewReader(file), *module}

	module.Read(&reader, 0)

	frame := NewFrame(module.Code)
	frame.Execute()
}

func resolve(target string) (file *os.File, err error) {
	if path.Ext(target) == ".pyc" {
		file, err = os.Open(target)
		return
	}

	if _, err := exec.LookPath("python3.4"); err != nil {
		err = fmt.Errorf("gothon: python3.4 needed for compilation")
		return nil, err
	}

	cmd := exec.Command("python3.4", "-m", "compileall", "-l", target)
	output, err := cmd.CombinedOutput()

	//if len(output) > 0 {
	fmt.Print(string(output))
	//}

	if err != nil {
		return
	}

	base := path.Base(target)
	target = path.Join(path.Dir(target), "__pycache__", base[:len(base)-2]+"cpython-34.pyc")
	file, err = os.Open(target)
	return
}
