{{ define "boxed" }}
/*Package box {{ .Version }} genereted this file on {{ .Date }}*/
package box

func init(){
  box["{{ .Name }}"] = "{{ .Content }}"
}
{{ end }}
