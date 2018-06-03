package main

import (
  "flag"
  "fmt"
  "io"
  "log"
  "os"
  "github.com/justinbarrick/go-terraform-plan/plan"
)

func main() {
  planFile := flag.String("plan", "", "Path to plan file to read.")
  flag.Parse()

  var err error
  var planReader io.Reader
  if *planFile != "" {
    planReader, err = os.Open(*planFile)
    if err != nil {
      log.Fatal(err)
    }
  } else {
    planReader = os.Stdin
  }

  plan, err := plan.ReadPlanFile(planReader)
  if err != nil {
    log.Fatal(err)
  }

  output, err := plan.PlanJson()
  if err != nil {
    log.Fatal("Error encoding json:", err)
  }

  fmt.Println(output)
}
