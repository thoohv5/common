package main

import (
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
)

var methodSets = make(map[string]int)

// generateFile generates a _http.pb.go file containing common errors definitions.
func generateFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	if len(file.Services) == 0 {
		return nil
	}
	filename := file.GeneratedFilenamePrefix + ".go"
	g := gen.NewGeneratedFile(strings.Replace(filename, "api", "internal/service", 1), file.GoImportPath)

	gfp := file.GeneratedFilenamePrefix
	idx := strings.LastIndex(gfp, "/")
	// suffix := gfp[idx:]
	ngf := gfp[:idx]

	g.P("package ", file.GoPackageName)
	g.P()
	g.P("import ( ")
	g.P("   \"context\"")
	g.P()
	g.P("   pb \"github.com/thoohv5/" + ngf + "\"")
	g.P(")")
	g.P()
	generateFileContent(gen, file, g)

	return g
}

// generateFileContent generates the common errors definitions, excluding the package statement.
func generateFileContent(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile) {
	if len(file.Services) == 0 {
		return
	}
	for _, service := range file.Services {
		genService(gen, file, g, service)
	}
}

func genService(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile, service *protogen.Service) {
	if service.Desc.Options().(*descriptorpb.ServiceOptions).GetDeprecated() {
		g.P("//")
	}

	// service.
	sd := &serviceDesc{
		ServiceType: service.GoName,
		ServiceName: string(service.Desc.FullName()),
		Metadata:    file.Desc.Path(),
	}
	for _, method := range service.Methods {
		if method.Desc.IsStreamingClient() || method.Desc.IsStreamingServer() {
			continue
		}
		sd.Methods = append(sd.Methods, buildMethodDesc(g, method, "POST", ""))
	}
	if len(sd.Methods) != 0 {
		g.P(sd.execute())
	}
}

func buildMethodDesc(g *protogen.GeneratedFile, m *protogen.Method, method, path string) *methodDesc {
	defer func() { methodSets[m.GoName]++ }()

	return &methodDesc{
		Name:         camelCase(m.GoName),
		OriginalName: string(m.Desc.Name()),
		Num:          methodSets[m.GoName],
		Request:      g.QualifiedGoIdent(m.Input.GoIdent),
		Reply:        g.QualifiedGoIdent(m.Output.GoIdent),
		Path:         path,
		Method:       method,
	}
}

// camelCase returns the CamelCased name.
// If there is an interior underscore followed by a lower case letter,
// drop the underscore and convert the letter to upper case.
// There is a remote possibility of this rewrite causing a name collision,
// but it's so remote we're prepared to pretend it's nonexistent - since the
// C++ generator lowercases names, it's extremely unlikely to have two fields
// with different capitalizations.
// In short, _my_field_name_2 becomes XMyFieldName_2.
func camelCase(s string) string {
	if s == "" {
		return ""
	}
	t := make([]byte, 0, 32)
	i := 0
	if s[0] == '_' {
		// Need a capital letter; drop the '_'.
		t = append(t, 'X')
		i++
	}
	// Invariant: if the next letter is lower case, it must be converted
	// to upper case.
	// That is, we process a word at a time, where words are marked by _ or
	// upper case letter. Digits are treated as words.
	for ; i < len(s); i++ {
		c := s[i]
		if c == '_' && i+1 < len(s) && isASCIILower(s[i+1]) {
			continue // Skip the underscore in s.
		}
		if isASCIIDigit(c) {
			t = append(t, c)
			continue
		}
		// Assume we have a letter now - if not, it's a bogus identifier.
		// The next word is a sequence of characters that must start upper case.
		if isASCIILower(c) {
			c ^= ' ' // Make it a capital letter.
		}
		t = append(t, c) // Guaranteed not lower case.
		// Accept lower case sequence that follows.
		for i+1 < len(s) && isASCIILower(s[i+1]) {
			i++
			t = append(t, s[i])
		}
	}
	return string(t)
}

// Is c an ASCII lower-case letter?
func isASCIILower(c byte) bool {
	return 'a' <= c && c <= 'z'
}

// Is c an ASCII digit?
func isASCIIDigit(c byte) bool {
	return '0' <= c && c <= '9'
}
