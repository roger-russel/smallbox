{{ define "boxed" }}
/*Package box {{ .Version }} genereted this file on {{ .Date }}*/
package box

func init(){
  initializeBox()
  box["{{ .Name }}"] = "{{ .Content }}"
}
{{ end }}
