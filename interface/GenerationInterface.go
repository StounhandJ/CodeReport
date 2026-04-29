package _interface

type GenerationInterface interface {
	CreateTable() TableGenerationInterface
	AddHeadingText(string)
	AddText(string)
	AddFileText(string) error
	Close() error
}
