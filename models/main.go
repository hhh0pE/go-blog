package models

type IPage interface {
    GetTemplates() []string
    Permalink() string
    GetByUrl(string) error
    GetByID(int) error
}