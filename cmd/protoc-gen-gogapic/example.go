// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import "github.com/golang/protobuf/protoc-gen-go/descriptor"

func (g *generator) genExampleFile(serv *descriptor.ServiceDescriptorProto, pkgName string) {
	servName := reduceServName(*serv.Name, pkgName)
	p := g.printf

	p("func ExampleNew%sClient() {", servName)
	g.exampleInitClient(pkgName, servName)
	p("  // TODO: Use client.")
	p("  _ = c")
	p("}")
	p("")
	g.imports[importSpec{path: "golang.org/x/net/context"}] = true

	for _, m := range serv.Method {
		g.exampleMethod(pkgName, servName, m)
	}
}

func (g *generator) exampleInitClient(pkgName, servName string) {
	p := g.printf

	p("ctx := context.Background()")
	p("c, err := %s.New%sClient(ctx)", pkgName, servName)
	p("if err != nil {")
	p("  // TODO: Handle error.")
	p("}")
}

func (g *generator) exampleMethod(pkgName, servName string, m *descriptor.MethodDescriptorProto) {
	p := g.printf

	inType := g.types[*m.InputType]
	inSpec := g.importSpec(inType)
	g.imports[inSpec] = true

	p("func Example%sClient_%s() {", servName, *m.Name)
	g.exampleInitClient(pkgName, servName)
	p("")
	p("  req := &%s.%s{", inSpec.name, *inType.Name)
	p("    // TODO: Fill request struct fields.")
	p("  }")

	if isLRO(m) {
		g.exampleLROCall(m)
	} else if *m.OutputType == emptyType {
		g.exampleEmptyCall(m)
	} else {
		g.exampleUnaryCall(m)
	}

	p("}")
	p("")
}

func (g *generator) exampleLROCall(m *descriptor.MethodDescriptorProto) {
	p := g.printf

	p("op, err := c.%s(ctx, req)", *m.Name)
	p("if err != nil {")
	p("  // TODO: Handle error.")
	p("}")
	p("")

	p("resp, err := op.Wait(ctx)")
	p("if err != nil {")
	p("  // TODO: Handle error.")
	p("}")
	p("// TODO: Use resp.")
	p("_ = resp")
}

func (g *generator) exampleUnaryCall(m *descriptor.MethodDescriptorProto) {
	p := g.printf

	p("resp, err := c.%s(ctx, req)", *m.Name)
	p("if err != nil {")
	p("  // TODO: Handle error.")
	p("}")
	p("// TODO: Use resp.")
	p("_ = resp")
}

func (g *generator) exampleEmptyCall(m *descriptor.MethodDescriptorProto) {
	p := g.printf

	p("err = c.%s(ctx, req)", *m.Name)
	p("if err != nil {")
	p("  // TODO: Handle error.")
	p("}")
}
