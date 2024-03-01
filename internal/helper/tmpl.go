package helper

import (
	"bytes"
	"errors"
	"html/template"
	"os"

	"github.com/XxThunderBlast/thunder-api/domain"
)

func BuildEmailTmpl(m domain.Message, tmplPath string) (buff bytes.Buffer, err error) {
	htmlBuff, htmlErr := os.ReadFile(tmplPath)
	if htmlErr != nil {
		return buff, errors.New("Error while reading template: " + htmlErr.Error())
	}

	tmpl, tmplErr := template.New("email").Parse(string(htmlBuff))
	if tmplErr != nil {
		return buff, errors.New("Error while parsing template: " + tmplErr.Error())
	}
	if err := tmpl.Execute(&buff, m); err != nil {
		return buff, errors.New("Error while executing template: " + err.Error())
	}

	return buff, nil
}
