package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {

	file := flag.String("file", "", "The file containing the animations you would like to render")
	scene := flag.String("scene", "", "The scene within the file you would like to render")
	args := flag.String("args", "", "The manim arguments you would like to render your scene with")

	required := []string{"file", "args", "scene"}
	flag.Parse()

	seen := make(map[string]bool)
	flag.VisitAll(func(f *flag.Flag) {
		if f.Value.String() != "" {
			seen[f.Name] = true
		}
	})
	for _, req := range required {
		if !seen[req] {
			fmt.Fprintf(os.Stderr, "Missing required flag: -%s\n", req)
			os.Exit(2)
		}
	}

	fmt.Println("Manimating the", *scene, "scene in the file", *file, "with the", *args, "options...")

	runCmd := `cmd /C python3 C:\Users\safin\Documents\manim\manim.py C:\Users\safin\Documents\manim\my_animations\` + *file + ` ` + *scene + ` -` + *args

	fmt.Println(runCmd)

	cmd, err := exec.Command(runCmd).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", cmd)
}
