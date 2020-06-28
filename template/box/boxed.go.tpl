{{ define "boxed" }}
/*Package box {{ .Version }} */
package box

func init(){
  initializeBox()
  box["{{ .Name }}"] = "{{ .Content }}"
}
{{ end }}
