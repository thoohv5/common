package goshared

const inTpl = `{{ $f := .Field }}{{ $r := .Rules }}
	{{ if $r.In }}
		if _, ok := {{ lookup $f "InLookup" }}[{{ accessor . }}]; !ok {
		{{ if $r.GetMessage -}}
			err := {{ err . ($r.GetMessage) }}
		{{ else -}}
			err := {{ err . "value must be in list " $r.In }}
		{{ end -}}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ else if $r.NotIn }}
		if _, ok := {{ lookup $f "NotInLookup" }}[{{ accessor . }}]; ok {
		{{ if $r.GetMessage -}}
			err := {{ err . ($r.GetMessage) }}
		{{ else -}}
			err := {{ err . "value must not be in list " $r.NotIn }}
		{{ end -}}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}
`
