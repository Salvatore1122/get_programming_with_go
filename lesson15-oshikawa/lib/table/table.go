package table

import (
	"fmt"
	"strings"
)


type Table struct {
	Header [2]string
	Body [][2]string
}
func NewTable(header [2]string, body [][2]string) *Table {
	return &Table{
		Header: header,
		Body: body,
	}
}

// @todo 以下Tableの責任ではない
func (t *Table) Display() {
	fmt.Println(
		strings.Join([]string{
			t.getSeparator(),
			t.getHeader(),
			t.getSeparator(),
			strings.Join(t.getBody(), "\n"),
			t.getSeparator(),
		}, "\n"),
	)
}

func (t *Table) getHeader() string {
	return t.getRow(t.Header)
}
func (t *Table) getBody() []string {
	var lines []string
	for _, r := range t.Body {
		lines = append(lines, t.getRow(r))
	}

	return lines
}
func (t *Table) getRow(r [2]string) string {
	return fmt.Sprintf("| %-8s | %-8s |", r[0], r[1])
}
func (t *Table) getSeparator() string {
	return "======================="
}