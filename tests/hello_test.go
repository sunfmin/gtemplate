package test

import (
	"bytes"
	"github.com/sunfmin/gtemplate/tests/templates"
	"testing"
)

func TestGreeting(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	d := &templates.HelloData{}
	d.Name = "Felix"
	d.Email = "sunfmin@gmail.com"
	d.Entries = append(d.Entries, &templates.Entry{"post", "Hello gtemplate"})
	d.Entries = append(d.Entries, &templates.Entry{"comment", "Compiled? really?"})
	templates.Hello(buf, d)
	expected := `<h1>Hi, Felix <em>sunfmin@gmail.com</em></h1>
<ul>
<li><div class="post">Hello gtemplate</div>
Unknown Language
</li>
<li><div class="comment">Compiled? really?</div>
Unknown Language
</li>
</ul>
`
	if buf.String() != expected {
		t.Errorf("expected:\n%s, but was:\n%s", expected, buf.String())
	}
}
