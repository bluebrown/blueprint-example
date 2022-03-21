package main

import (
    "fmt"
    {{ if .Values.service.enabled }}
    "{{ .Values.git.provider }}/{{ .Values.git.user }}/{{ .ProjectName }}/svc/{{.Values.service.name }}"
    {{- end }}
)

func main() {
	fmt.Println("{{ .Values.message }}")
    {{- if .Values.service.enabled }}
    {{.Values.service.name }}.parse()
    {{- end }}
}
