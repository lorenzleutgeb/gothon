package main

import ("flag" ; "log" ; "os" ; "encoding/binary" ; "bufio" ; "time" ; "fmt" ; "github.com/flowlo/gothon" ; "os/exec" ; "path" )

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
		if err != nil { log.Fatalf("If you want gothon to compile \"%s\", please install Python 3.3", path.Base(target)) }
		if *debug { log.Print("Found Python 3.3") }

		cmd := exec.Command("python3.3", "/usr/lib/python3.3/compileall.py", target)
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

	reader := bufio.NewReader(file)

	var version, check, size uint16

	binary.Read(reader, binary.LittleEndian, &version)
	binary.Read(reader, binary.LittleEndian, &check)

	if check != 0xa0d {
		log.Fatalf("Second two bytes of magic number are incorrect (0x0a0d expected, 0x%x found)", check)
	}

	var stamp int32

	binary.Read(reader, binary.LittleEndian, &stamp)

	if *verbose { log.Print("Last modification: ", time.Unix(int64(stamp), 0)) }

	binary.Read(reader, binary.LittleEndian, &size)

	if *debug { log.Print("Claimed file size: ", size) }

	var c byte

	for {
		c, err = reader.ReadByte()
		if err != nil { log.Fatal(err) }
		log.Print(gothon.GetMnemonic(c))
		if gothon.HasArg(c) {
			c, err = reader.ReadByte()
			c, err = reader.ReadByte()
		}
	}

	file.Close()
}
