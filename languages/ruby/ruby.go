package ruby

import (
	"bytes"
	"fmt"
	. "github.com/deevatech/runner/types"
	"io/ioutil"
	"log"
	"os/exec"
)

type Context RunContext

func NewContext(p RunParams) *Context {
	return &Context{
		Path:   "/home/deeva/code/ruby",
		Params: p,
	}
}

func Run(p RunParams) RunResults {
	log.Printf("Running ruby code: %#v", p)

	ctx := NewContext(p)
	ctx.createSourceFile()
	ctx.createSpecFile()

	if err := ctx.runSpec(); err != nil {
		log.Fatal(err)
	}

	return ctx.Results
}

func (ctx Context) createSourceFile() error {
	data := []byte(ctx.Params.Source)
	return ioutil.WriteFile(ctx.sourceFilename(), data, 0644)
}

func (ctx Context) createSpecFile() error {
	data := []byte(ctx.Params.Spec)
	return ioutil.WriteFile(ctx.specFilename(), data, 0644)
}

func (ctx Context) sourceFilename() string {
	return fmt.Sprintf("%s/lib/solution.rb", ctx.Path)
}

func (ctx Context) specFilename() string {
	return fmt.Sprintf("%s/spec/solution_spec.rb", ctx.Path)
}

func (ctx *Context) runSpec() error {
	rspec, err := exec.LookPath("rspec")
	if err != nil {
		return fmt.Errorf("Unable to find rspec!")
	}

	var out bytes.Buffer
	cmd := exec.Command(rspec, "-fd")
	cmd.Stdout = &out
	cmd.Stderr = &out

	if err := cmd.Run(); err != nil {
		return err
	}

	ctx.Results.Output = out.String()

	return nil
}
