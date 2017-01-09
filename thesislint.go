package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"path/filepath"
	"sync"
	"regexp"
)

func main() {
	if err := os.Mkdir("tmp", 0777); err != nil {
		if err := os.RemoveAll("thesislint_tmp"); err != nil {
			fmt.Println(err)
		}
	}

	files, err := ioutil.ReadDir("./")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		var filenames []string
		var wg sync.WaitGroup
		for _, f := range files {
			file_path := f.Name()
			ext := filepath.Ext(file_path)
			if ext == ".tex" {
				s := strings.Split(file_path, ".")
				filename := s[0]
				name := filename
				filenames = append(filenames, "tmp/" + filename + ".md")
				wg.Add(1)
				go func() {
					exec.Command("pandoc", "-s", name + ".tex", "-o", "tmp/" + name + ".md").Run()
					wg.Done()
				}()
			}
		}
		wg.Wait()
		var cmd = exec.Command("textlint", "--dry-run")
		cmd.Args = filenames
		out, _ := cmd.CombinedOutput()
		var o = string(out)
		var rep = regexp.MustCompile("✖.*\n")
		o = rep.ReplaceAllString(o, "\x1b[31m$0\x1b[0m")
		rep = regexp.MustCompile(".*.md\n")
		o = rep.ReplaceAllString(o, "\x1b[35m$0\x1b[0m")
		rep = regexp.MustCompile(" error ")
		o = rep.ReplaceAllString(o, "\x1b[31m error \x1b[0m")
		
		fmt.Println(o)

		if err := os.RemoveAll("thesislint_tmp"); err != nil {
			fmt.Println(err)
		}
	}
}
