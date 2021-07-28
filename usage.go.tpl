{{$lb := "{" -}}
{{$rb := "}" -}}
{{if .Name -}}
NAME
  {{.Name}}
{{end -}}
{{if .Description}}
DESCRIPTION
  {{.Description}}
{{end -}}
{{if .List}}
OPTIONS
{{- end}}
{{range $k, $v := .List}}  -{{$v.Name -}}
    {{if ne "bool" $v.Value.Type}} {{$v.Value.Type -}}{{end -}}
    {{range $i, $a := $v.Aliases -}}
, -{{$a}}{{if ne "bool" $v.Value.Type}}={{$lb}}{{$v.Value.Type}}{{$rb}}{{end -}}
    {{end}}
        {{$v.Usage}}
{{end}}