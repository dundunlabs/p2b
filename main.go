package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/dundunlabs/p2b/generator"
	"github.com/dundunlabs/p2b/jsonrpc"
	"github.com/dundunlabs/p2b/prisma-generator"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	maxCapacity := 10 * 1024 * 1024
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, maxCapacity)

	for scanner.Scan() {
		req, err := jsonrpc.Read(scanner.Bytes())
		if err != nil {
			log.Fatal(err)
		}

		switch req.Method {
		case "getManifest":
			err = jsonrpc.Write(os.Stderr, req.ID, prisma.ManifestResult{
				Manifest: prisma.Manifest{
					PrettyName:    "Prisma to Bun",
					DefaultOutput: "../model",
				},
			})
		case "generate":
			var p prisma.GenerateParams
			if err := json.Unmarshal(req.Params, &p); err != nil {
				log.Fatal(err)
			}

			g := generator.New(p)
			if err := g.Generate(); err != nil {
				log.Fatal(err)
			}

			err = jsonrpc.Write(os.Stderr, req.ID, nil)
		default:
			err = fmt.Errorf("unsupported method: %s", req.Method)
		}

		if err != nil {
			log.Fatal(err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
