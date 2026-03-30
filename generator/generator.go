package generator

type Policy struct {
	Length           int
	IncludeUpper     bool
	IncludeLower     bool
	IncludeDigits    bool
	IncludeSymbols   bool
	RequireEachClass bool
}
