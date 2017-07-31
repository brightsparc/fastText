package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/brightsparc/fasttextgo"
	flags "github.com/jessevdk/go-flags"
)

func main() {
	var opts = struct {
		Positional struct {
			Command   string `required:"yes"`
			ModelFile string `required:"yes"`
			TestFile  string `required:"yes"`
			K         int    `optional:"yes"` // Defaults to 0
		} `positional-args:"yes"`
	}{}

	var parser = flags.NewParser(&opts, flags.IgnoreUnknown)
	_, err := parser.Parse()
	if err != nil {
		log.Fatal(err)
	}

	// Load model
	t0 := time.Now()
	fasttextgo.LoadModel(opts.Positional.ModelFile)
	log.Printf("Model loaded in %s\n", time.Since(t0))

	switch opts.Positional.Command {
	case "predict", "predict-prob":
		f, err := os.Open(opts.Positional.TestFile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			if opts.Positional.K > 1 {
				log.Fatal("K>1 not supported yet")
			}
			prob, label, err := fasttextgo.Predict(scanner.Text())
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Print(label)
				if opts.Positional.Command == "predict-prob" {
					fmt.Printf(" %f\n", prob)
				} else {
					fmt.Println()
				}
			}
		}
		if err := scanner.Err(); err != nil {
			log.Printf("Error reading input: %s\n", err)
		}
	default:
		log.Fatalf("Command %s not supported", opts.Positional.Command)
	}
}
