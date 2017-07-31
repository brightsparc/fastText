package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"

	"github.com/brightsparc/fasttextgo"
	flags "github.com/jessevdk/go-flags"
)

// make fasttext-go && ./fasttext-go predict-prob result/dbpedia.bin data/dbpedia.test 1 -p main.pprof > nul
// go tool pprof ./fasttext-go main.pprof

func main() {
	var opts = struct {
		CPUProfile string `short:"p" long:"cpuprofile" description:"Enable cpu profile"`
		Positional struct {
			Command   string
			ModelFile string
			TestFile  string
			K         int
		} `positional-args:"yes" required:"yes"`
	}{}

	var parser = flags.NewParser(&opts, flags.IgnoreUnknown)
	_, err := parser.Parse()
	if err != nil {
		log.Fatal(err)
	}

	// Enable CPU profile
	if opts.CPUProfile != "" {
		f, err := os.Create(opts.CPUProfile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
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
