package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

type Symbol struct {
	Name string
	File string
	Line string
}

const repo = "http://github.com/qcoh/dmgb"

func main() {
	cmd := exec.Command("/bin/sh", "cloneandtag.sh", repo)
	out, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	db := make(map[string][]Symbol)

	scanner := bufio.NewScanner(out)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		sym := Symbol{
			Name: s[0],
			File: s[1][20:],
			Line: s[2],
		}
		db[sym.Name] = append(db[sym.Name], sym)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, v := range db[os.Args[1]] {
		fmt.Printf("%s/blob/master/%s#L%s\n", repo, v.File, v.Line)
	}
}
