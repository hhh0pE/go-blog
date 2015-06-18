package db

import (
    "strings"
)

const TMP_DIR = "templates/"

type Template struct {
	ID, ParentID         int
	Parent      *Template
	Name, File string
}

func (t Template) ToStrings() []string {
    stemps := []string{}

    stemps = append(stemps, TMP_DIR+t.File)

    temp := t
    for temp.Parent!=nil {
        stemps = append(stemps, TMP_DIR+temp.Parent.File)
        temp = *temp.Parent
    }

    return stemps
}

func (t Template) ToString() string {
    return strings.Join(t.ToStrings(), ", ")
}