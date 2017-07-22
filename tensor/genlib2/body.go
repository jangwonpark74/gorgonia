package main

// generic loop templates

type LoopBody struct {
	TypedOp
	Range string
	Left  string
	Right string

	Index0, Index1, Index2 string

	IterName0, IterName1, IterName2 string
}

const (
	genericLoopRaw = `for i := range {{.Range}} {
		{{template "check" . -}}
		{{template "loopbody" .}}
	}`

	genericUnaryIterLoopRaw = `var {{.Index0}} int
	var valid{{.Index0}} bool
	for {
		if {{.Index0}}, valid{{.Index0}}, err = {{.IterName0}}.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if valid{{.Index0}} {
			{{template "check" . -}}
			{{template "loopbody" . -}}
		}
	}`

	genericBinaryIterLoopRaw = `var {{.Index0}}, {{.Index1}} int
	var valid{{.Index0}}, valid{{.Index1}} bool
	for {
		if {{.Index0}}, valid{{.Index0}}, err = {{.IterName0}}.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if {{.Index1}}, valid{{.Index1}}, err = {{.IterName1}}.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if valid{{.Index0}} && valid{{.Index1}} {
			{{template "check" . -}}
			{{template "loopbody" . -}}
		}
	}`

	genericTernaryIterLoopRaw = `var {{.Index0}}, {{.Index1}}, {{.Index2}} int
	var valid{{.Index0}}, valid{{.Index1}}, valid{{.Index2}} bool
	for {
		if {{.Index0}}, valid{{.Index0}}, err = {{.IterName0}}.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if {{.Index1}}, valid{{.Index1}}, err = {{.IterName1}}.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if {{.Index2}}, valid{{.Index2}}, err = {{.IterName2}}.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if valid{{.Index0}} && valid{{.Index1}} && valid{{.Index2}} {
			{{template "check" . -}}
			{{template "loopbody" . -}}
		}
	}`

	// ALL THE SYNTACTIC ABSTRACTIONS!
	// did I mention how much I hate C-style macros? Now I'm doing them instead

	basicSet = `{{if .IsFunc -}}
			{{.Range}}[i] = {{ template "callFunc" . -}}
		{{else -}}
			{{.Range}}[i] = {{template "opDo" . -}}
		{{end -}}`

	basicIncr = `{{if .IsFunc -}}
			{{.Range}}[i] += {{template "callFunc" . -}}
		{{else -}}
			{{.Range}}[i] += {{template "opDo" . -}}
		{{end -}}`

	iterIncrLoopBody = `{{if .IsFunc -}}
			{{.Range}}[k] += {{template "callFunc" . -}}
		{{else -}}
			{{.Range}}[k] += {{template "opDo" . -}}
		{{end -}}`

	sameSet = `if {{template "opDo" . }} {
		{{.Range}}[i] = {{trueValue .Kind}}
	}else{
		{{.Range}}[i] = {{falseValue .Kind}}
	}
	`

	ternaryIterSet = `{{.Range}}[k] = {{template "opDo" . -}}`

	binOpCallFunc = `{{if eq "complex64" .Kind.String -}}
		complex64({{template "symbol" .Kind}}(complex128({{.Left}}), complex128({{.Right}})))
		{{else -}}
		{{template "symbol" .Kind}}({{.Left}}, {{.Right}})
		{{end -}}`

	binOpDo = `{{.Left}} {{template "symbol" .Kind}} {{.Right}}`

	unaryOpDo = `{{template "symbol" .Kind}}{{.Left}}`

	unaryOpCallFunc = `{{template "symbol" .Kind}}({{.Left}})`

	check0 = `if {{.Right}} == 0 {
		errs = append(errs, i)
		{{.Range}}[i] = 0
		continue
	}
	`
)

// renamed
const (
	vvLoopRaw         = genericLoopRaw
	vvIncrLoopRaw     = genericLoopRaw
	vvIterLoopRaw     = genericBinaryIterLoopRaw
	vvIterIncrLoopRaw = genericTernaryIterLoopRaw

	mixedLoopRaw         = genericLoopRaw
	mixedIncrLoopRaw     = genericLoopRaw
	mixedIterLoopRaw     = genericUnaryIterLoopRaw
	mixedIterIncrLoopRaw = genericBinaryIterLoopRaw
)
