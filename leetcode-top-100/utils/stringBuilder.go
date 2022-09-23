package utils

import (
	"bytes"
	"fmt"
	"strings"
)

type StringBuilder struct {
	buffer bytes.Buffer
}

func NewStringBuilderNoParam() *StringBuilder {
	return &StringBuilder{}
}

func NewStringBuilder(s ...string) *StringBuilder {
	builder := StringBuilder{}
	builder.buffer.WriteString(strings.Join(s, ""))
	return &builder
}

func (builder *StringBuilder) Append(s string) *StringBuilder {
	builder.buffer.WriteString(s)
	return builder
}

func (builder *StringBuilder) AppendFormat(format string, s ...interface{}) *StringBuilder {
	builder.buffer.WriteString(fmt.Sprintf(format, s...))
	return builder
}

func (builder *StringBuilder) String() string {
	return builder.buffer.String()
}
