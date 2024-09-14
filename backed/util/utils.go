package util

import "regexp"

const ()

var (
	Rule_name, _     = regexp.Compile("[a-zA-Z0-9]{4,20}")
	Rule_password, _ = regexp.Compile("[a-zA-Z0-9]{6,30}")
	Rule_email, _    = regexp.Compile("[a-zA-Z0-9][\\w\\\\.-]*[a-zA-Z0-9]@[a-zA-Z0-9][\\w\\\\.-]*[a-zA-Z0-9]\\.[a-zA-Z][a-zA-Z\\\\.]*[a-zA-Z]")
)
