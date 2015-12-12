package main

import (
	"go/ast"
	"go/token"
	"strconv"
	"strings"
	"text/template"
)

type instance struct {
	Name string
	Addr string
}

type multiCtx struct {
	Pkg       string
	Periph    string
	Instances []instance
	Regs      []reg
	Len       int
}

func instances(f string, lines []string) ([]instance, []string) {
	var insts []instance
	for len(lines) > 0 {
		line := strings.TrimSpace(lines[0])
		lines = lines[1:]
		if line == "Registers:" {
			break
		}
		name, addr := nameval(line, ':')
		_, err := strconv.ParseUint(addr, 0, 64)
		if err != nil {
			fdie(f, "bad %s address '%s': %v", name, addr, err)
		}
		insts = append(insts, instance{Name: name, Addr: addr})
	}
	return insts, lines
}

func multi(pkg, f, txt string, decls []ast.Decl) {
	lines := strings.Split(txt, "\n")
	p := lines[0]
	lines = lines[1:]
	p = strings.TrimSpace(p[len("Peripheral:"):])
	var insts []instance
	for len(lines) > 0 {
		line := strings.TrimSpace(lines[0])
		lines = lines[1:]
		if line == "Instances:" {
			insts, lines = instances(f, lines)
			break
		} else if line == "Registers:" {
			break
		}
	}
	ctx := &multiCtx{
		Pkg:       pkg,
		Periph:    p,
		Instances: insts,
	}
	ctx.Regs, ctx.Len = regs(f, lines)
	regmap := make(map[string]*reg)
	for i := range ctx.Regs {
		regmap[ctx.Regs[i].Reg] = &ctx.Regs[i]
	}
	for _, d := range decls {
		g, ok := d.(*ast.GenDecl)
		if !ok || g.Tok != token.CONST {
			continue
		}
		for _, s := range g.Specs {
			v := s.(*ast.ValueSpec)
			t, ok := v.Type.(*ast.Ident)
			if !ok {
				continue
			}
			i := strings.LastIndexByte(t.Name, '_')
			if i < 0 {
				continue
			}
			reg := t.Name[:i]
			typ := t.Name[i+1:]
			var und string
			switch typ {
			case "Bits":
				und = "uint32"
			default:
				continue
			}
			for _, i := range v.Names {
				if r := regmap[reg]; r != nil {
					r.Decls = append(r.Decls, decl{i.Name, typ, und})
				}
			}
		}
	}
	save(f, multiTmpl, ctx)
}

const multiText = `package {{.Pkg}}

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"
)
{{$p := .Periph}}
type {{$p}} struct{}

func (p *{{$p}}) r(n uint) *mmio.U32 {
	return &(*[{{.Len}}]mmio.U32)(unsafe.Pointer(p))[n]
}

{{range .Instances}}
var {{.Name}} = (*{{$p}})(unsafe.Pointer(uintptr({{.Addr}}))){{end}}
{{range .Regs}}
{{$prn := print "p.r(" .N ")"}}
type {{.Bits}} uint32
{{range .Decls}}
func (p *{{$p}}) {{.Name}}() mmio.{{.Typ}}32 {return mmio.{{.Typ}}32{{"{"}}{{$prn}}, {{.Und}}({{.Name}})} }{{end}}

func (p *{{$p}}) {{.Reg}}_Load() {{.Bits}}   { return {{.Bits}}({{$prn}}.Load()) }
func (p *{{$p}}) {{.Reg}}_Store(b {{.Bits}}) { {{$prn}}.Store(uint32(b)) }

func (b {{.Bits}}) Field(mask {{.Bits}}) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func Make_{{.Reg}}(v int, mask {{.Bits}}) {{.Bits}} {
	return {{.Bits}}(bits.Make32(v, uint32(mask)))
}
{{end}}
`

var multiTmpl = template.Must(template.New("multi").Parse(multiText))
