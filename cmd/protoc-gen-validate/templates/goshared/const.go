package goshared

const constTpl = `{{ $r := .Rules }}
	{{ if $r.Const }}
		if {{ accessor . }} != {{ lit $r.GetConst }} {
		{{ if $r.GetMessage -}}
			err := {{ err . ($r.GetMessage) }}
		{{ else -}}
			err := {{ err . "value must equal " $r.GetConst }}
		{{ end -}}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}
`
