{{ define "boxed" }}
/*Package box {{ .Version }} generated at: {{ .Date }}*/
package box

func init(){
  initializeBox()
  box["{{ .Name }}"] = "{{ .Content }}"
}
{{ end }}
