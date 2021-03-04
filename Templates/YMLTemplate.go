package plugintemplates

// GetYamlTemplate Outputs the string template for creating a plugin.yml
func GetYamlTemplate() string {
	return `main: com.{{.Author}}.net.Main
name: {{.Name}}
author: {{.Author}}
version: 1
api-version: 1.15
{{with .Cmds -}}
commands:{{range $val := .}}
  {{.SlashCommand}}:{{end}}
{{end}}
`
}
