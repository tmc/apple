// Code generated from internal/generator/templates/runtime/objc.txtar by applegen. DO NOT EDIT.

// Package objcinspect provides Objective-C type encoding and signature string parsing.
package objcinspect

import (
	"fmt"
	"strconv"
	"strings"
)

// Kind represents the high-level category of an Objective-C type.
type Kind int

const (
	KindInvalid Kind = iota
	KindVoid
	KindId
	KindClass
	KindSEL
	KindBool
	KindChar
	KindUChar
	KindShort
	KindUShort
	KindInt
	KindUInt
	KindLong
	KindULong
	KindLongLong
	KindULongLong
	KindFloat
	KindDouble
	KindBitfield
	KindPointer
	KindArray
	KindStruct
	KindUnion
	KindBlock
	KindProtocol
	KindUnknown
)

// String returns a human-readable representation of Kind.
func (k Kind) String() string {
	switch k {
	case KindVoid:
		return "void"
	case KindId:
		return "id"
	case KindClass:
		return "Class"
	case KindSEL:
		return "SEL"
	case KindBool:
		return "BOOL"
	case KindChar:
		return "char"
	case KindUChar:
		return "unsigned char"
	case KindShort:
		return "short"
	case KindUShort:
		return "unsigned short"
	case KindInt:
		return "int"
	case KindUInt:
		return "unsigned int"
	case KindLong:
		return "long"
	case KindULong:
		return "unsigned long"
	case KindLongLong:
		return "long long"
	case KindULongLong:
		return "unsigned long long"
	case KindFloat:
		return "float"
	case KindDouble:
		return "double"
	case KindBitfield:
		return "bitfield"
	case KindPointer:
		return "pointer"
	case KindArray:
		return "array"
	case KindStruct:
		return "struct"
	case KindUnion:
		return "union"
	case KindBlock:
		return "block"
	case KindProtocol:
		return "Protocol"
	default:
		return "unknown"
	}
}

// Type represents a parsed Objective-C type.
//
// The parser does not record the qualifiers r, n, N, o, O, R and V: they are
// consumed and discarded, so a const-qualified type is indistinguishable from
// its unqualified form.
type Type struct {
	// Encoding is the type's own encoding, without the stack offset that
	// follows it in a method signature: parsing "@\"NSString\"16" yields
	// `@"NSString"`, not `@"NSString"16`. Callers match it against literals
	// and re-parse it, so a trailing offset would defeat both.
	Encoding  string
	Name      string
	FieldName string
	// HasQuote reports whether the field was written with a quoted name.
	// It is not redundant with FieldName != "": 522 encodings in the macOS
	// corpus carry an empty quoted field name (""), for example
	// `"__begin_"^@"__end_"^@""{?=...}`, so only HasQuote distinguishes an
	// unnamed field from a field named "".
	HasQuote bool
	IsAtomic bool
	Kind     Kind
	Elem     *Type  // Element type for pointer/array
	Fields   []Type // Struct/union fields
}

// Align returns the alignment requirement in bytes for the type.
func (t Type) Align() int {
	switch t.Kind {
	case KindChar, KindUChar, KindBool, KindBitfield:
		return 1
	case KindShort, KindUShort:
		return 2
	case KindInt, KindUInt, KindFloat:
		return 4
	case KindLong, KindULong, KindLongLong, KindULongLong, KindDouble, KindId, KindClass, KindSEL, KindPointer, KindBlock:
		return 8
	case KindArray:
		if t.Elem != nil {
			return t.Elem.Align()
		}
		return 8
	case KindStruct, KindUnion:
		maxAlign := 1
		for _, f := range t.Fields {
			a := f.Align()
			if a > maxAlign {
				maxAlign = a
			}
		}
		if maxAlign < 1 {
			maxAlign = 8
		}
		return maxAlign
	default:
		return 8
	}
}

// NaturalSize returns the size in bytes of the type assuming natural alignment.
// It returns -1 if the size cannot be determined (e.g. for opaque struct references,
// structs containing bitfields, or unparseable fields). Returns 0 for empty structs.
func (t Type) NaturalSize() int {
	switch t.Kind {
	case KindChar, KindUChar, KindBool:
		return 1
	case KindShort, KindUShort:
		return 2
	case KindInt, KindUInt, KindFloat:
		return 4
	case KindLong, KindULong, KindLongLong, KindULongLong, KindDouble, KindId, KindClass, KindSEL, KindPointer, KindBlock:
		return 8
	case KindBitfield:
		return -1
	case KindArray:
		if t.Elem != nil {
			elemSize := t.Elem.NaturalSize()
			if elemSize < 0 {
				return -1
			}
			count := 0
			enc := t.Encoding
			if strings.HasPrefix(enc, "[") && strings.HasSuffix(enc, "]") {
				inner := enc[1 : len(enc)-1]
				idx := 0
				for idx < len(inner) && inner[idx] >= '0' && inner[idx] <= '9' {
					idx++
				}
				if idx > 0 {
					count, _ = strconv.Atoi(inner[:idx])
				}
			}
			return count * elemSize
		}
		return -1
	case KindStruct:
		if !strings.Contains(t.Encoding, "=") {
			return -1
		}
		inner := t.Encoding[1 : len(t.Encoding)-1]
		eq := strings.IndexByte(inner, '=')
		if eq >= 0 && eq+1 == len(inner) {
			return -1
		}
		if len(t.Fields) > 0 {
			return calcNaturalSize(t.Fields)
		}
		return 0
	case KindUnion:
		if !strings.Contains(t.Encoding, "=") {
			return -1
		}
		inner := t.Encoding[1 : len(t.Encoding)-1]
		eq := strings.IndexByte(inner, '=')
		if eq >= 0 && eq+1 == len(inner) {
			return -1
		}
		if len(t.Fields) > 0 {
			maxSize := 0
			maxAlign := 1
			for _, f := range t.Fields {
				s := f.NaturalSize()
				if s < 0 {
					return -1
				}
				if s > maxSize {
					maxSize = s
				}
				a := f.Align()
				if a > maxAlign {
					maxAlign = a
				}
			}
			if remainder := maxSize % maxAlign; remainder != 0 {
				maxSize += maxAlign - remainder
			}
			return maxSize
		}
		return 0
	default:
		return 8
	}
}

// Signature represents a parsed Objective-C method or block signature.
type Signature struct {
	Return Type
	Args   []Type
}

// ParseSignature parses an Objective-C method signature string.
// Standard method signatures include return type followed by argument types
// and stack offset numbers (e.g. "v24@0:8@16").
func ParseSignature(encoding string) (Signature, error) {
	if len(encoding) == 0 {
		return Signature{}, fmt.Errorf("empty encoding")
	}

	p := &parser{s: encoding}

	ret, err := p.parseType()
	if err != nil {
		return Signature{}, fmt.Errorf("parse return type: %w", err)
	}

	var args []Type
	for p.i < len(p.s) {
		arg, err := p.parseType()
		if err != nil {
			return Signature{}, fmt.Errorf("parse argument %d: %w", len(args), err)
		}
		args = append(args, arg)
	}

	return Signature{Return: ret, Args: args}, nil
}

// ParseStruct parses an Objective-C struct type encoding (e.g. "{CGPoint=dd}").
//
// NOTE: Objective-C type encodings do NOT describe C struct packing, alignment,
// or padding bytes inserted by the C compiler. The naturalSize returned is calculated
// assuming natural alignment of each field. True memory layout and stride must be
// determined through C ABI rules or unsafe.Sizeof.
func ParseStruct(encoding string) (name string, fields []Type, naturalSize int, err error) {
	if !strings.HasPrefix(encoding, "{") || !strings.HasSuffix(encoding, "}") {
		return "", nil, -1, fmt.Errorf("invalid struct encoding %q", encoding)
	}

	inner := encoding[1 : len(encoding)-1]
	eq := strings.IndexByte(inner, '=')

	if eq == -1 {
		// Struct with name only, no field details (e.g. "{CGPoint}")
		sName := inner
		if sName == "?" {
			sName = ""
		}
		return sName, nil, -1, nil
	}

	sName := inner[:eq]
	if sName == "?" {
		sName = ""
	}

	if eq+1 == len(inner) {
		// Opaque struct with name and trailing '=', no fields (e.g. "{CGImage=}")
		return sName, nil, -1, nil
	}

	p := &parser{s: inner[eq+1:], inFields: true}
	for p.i < len(p.s) {
		var fieldName string
		hasQuote := false
		if p.i < len(p.s) && p.s[p.i] == '"' {
			hasQuote = true
			p.i++
			nStart := p.i
			for p.i < len(p.s) && p.s[p.i] != '"' {
				p.i++
			}
			fieldName = p.s[nStart:p.i]
			if p.i < len(p.s) && p.s[p.i] == '"' {
				p.i++
			}
		}
		var f Type
		if p.i >= len(p.s) || p.s[p.i] == '"' || p.s[p.i] == '}' || p.s[p.i] == ')' {
			f = Type{Encoding: "?", Kind: KindUnknown}
		} else {
			f, err = p.parseType()
			if err != nil {
				return sName, fields, -1, fmt.Errorf("parse field %d: %w", len(fields), err)
			}
		}
		f.FieldName = fieldName
		f.HasQuote = hasQuote
		fields = append(fields, f)
	}

	naturalSize = calcNaturalSize(fields)
	return sName, fields, naturalSize, nil
}

// ParseBlockEncoding parses an extended or bare block type encoding.
// Bare block arguments appear as "@?", while extended block signatures
// appear as "@?<v@?I>" carrying inner return and argument types.
func ParseBlockEncoding(encoding string) (Signature, error) {
	if encoding == "@?" {
		block := Type{Encoding: "@?", Kind: KindBlock}
		return Signature{
			Return: Type{Encoding: "v", Kind: KindVoid},
			Args:   []Type{block},
		}, nil
	}

	if strings.HasPrefix(encoding, "@?<") && strings.HasSuffix(encoding, ">") {
		inner := encoding[3 : len(encoding)-1]
		return ParseSignature(inner)
	}

	if strings.HasPrefix(encoding, "@?") {
		return ParseSignature(encoding[2:])
	}

	return ParseSignature(encoding)
}

// looksLikeObjectAddr reports whether id plausibly points to an Objective-C object.
func looksLikeObjectAddr(id uintptr) bool {
	if id == 0 {
		return false
	}
	u := uint64(id)
	// Check for arm64 tagged pointer (top bit set)
	if (u & (1 << 63)) != 0 {
		return true
	}
	// Reject addresses below first mapped page
	if u < 0x1000 {
		return false
	}
	// Reject misaligned pointers on 64-bit platforms
	if (u & 7) != 0 {
		return false
	}
	return true
}

type parser struct {
	s string
	i int
	// inFields reports whether s is the body of a struct or union field list,
	// where a quoted string may be a field name rather than a class name.
	// It propagates through pointer elements, which are parsed from the same
	// input, but not into a nested field list or array, each of which gets its
	// own parser and its own delimiters.
	inFields bool
}

// startsType reports whether a type encoding begins at index i, skipping any
// qualifiers and stack offsets that precede it.
func (p *parser) startsType(i int) bool {
	for i < len(p.s) {
		ch := p.s[i]
		if ch >= '0' && ch <= '9' || isQualifier(ch) {
			i++
			continue
		}
		break
	}
	if i >= len(p.s) {
		return false
	}
	switch p.s[i] {
	case 'c', 'C', 's', 'S', 'i', 'I', 'l', 'L', 'q', 'Q', 'f', 'd', 'B', 'v',
		'@', '#', ':', '*', '^', '{', '(', '[', 'b', '?':
		return true
	}
	return false
}

// isQualifier reports whether ch is a type qualifier: r (const), n (in),
// N (inout), o (out), O (bycopy), R (byref) or V (oneway).
func isQualifier(ch byte) bool {
	return ch == 'r' || ch == 'n' || ch == 'N' || ch == 'o' || ch == 'O' || ch == 'R' || ch == 'V'
}

func (p *parser) skipQualifiersAndOffsets() {
	for p.i < len(p.s) {
		ch := p.s[p.i]
		if isQualifier(ch) {
			p.i++
			continue
		}
		// Skip stack offset digits
		if ch >= '0' && ch <= '9' {
			p.i++
			continue
		}
		break
	}
}

// skipQualifiers advances past type qualifiers, leaving any stack offset in
// place. Use it where a following digit belongs to the enclosing signature
// rather than to the type being parsed.
func (p *parser) skipQualifiers() {
	for p.i < len(p.s) && isQualifier(p.s[p.i]) {
		p.i++
	}
}

func (p *parser) parseType() (Type, error) {
	p.skipQualifiersAndOffsets()
	if p.i >= len(p.s) {
		return Type{}, fmt.Errorf("unexpected end of encoding")
	}

	start := p.i
	ch := p.s[p.i]
	p.i++

	switch ch {
	case 'v':
		p.skipQualifiersAndOffsets()
		return Type{Encoding: "v", Kind: KindVoid}, nil
	case 'A':
		elem, err := p.parseType()
		if err != nil {
			return Type{}, fmt.Errorf("atomic elem: %w", err)
		}
		elem.Encoding = "A" + elem.Encoding
		elem.IsAtomic = true
		return elem, nil
	case '@':
		if p.i < len(p.s) && p.s[p.i] == '?' {
			p.i++
			if p.i < len(p.s) && p.s[p.i] == '<' {
				// Extended block signature: @?<...>
				depth := 1
				p.i++
				for p.i < len(p.s) && depth > 0 {
					if p.s[p.i] == '<' {
						depth++
					} else if p.s[p.i] == '>' {
						depth--
					}
					p.i++
				}
				enc := p.s[start:p.i]
				p.skipQualifiersAndOffsets()
				return Type{Encoding: enc, Kind: KindBlock}, nil
			}
			p.skipQualifiersAndOffsets()
			return Type{Encoding: "@?", Kind: KindBlock}, nil
		}
		// A quoted string after '@' is the class name, except inside a struct
		// or union field list, where it may instead be the next field's name.
		// The two are locally indistinguishable:
		//
		//	{__cfobservers_t="slot"@"next"^{__cfobservers_t}}
		//	@?<v@?@"GTMioUSCTraceData"^{GTMioUSCKickMetadata=...}Q>
		//
		// have the same shape, but "next" names a field and
		// "GTMioUSCTraceData" names a class. Only the context separates them.
		// A signature has no field names, so there the quote is always a class
		// name; inside a field list, a type character after the closing quote
		// means the quote opened the next field instead.
		if p.i < len(p.s) && p.s[p.i] == '"' {
			qEnd := p.i + 1
			for qEnd < len(p.s) && p.s[qEnd] != '"' {
				qEnd++
			}
			if qEnd < len(p.s) && p.s[qEnd] == '"' && !(p.inFields && p.startsType(qEnd+1)) {
				name := p.s[p.i+1 : qEnd]
				p.i = qEnd + 1
				enc := p.s[start:p.i]
				p.skipQualifiersAndOffsets()
				return Type{Encoding: enc, Kind: KindId, Name: name}, nil
			}
		}
		p.skipQualifiersAndOffsets()
		return Type{Encoding: "@", Kind: KindId}, nil

	case '#':
		p.skipQualifiersAndOffsets()
		return Type{Encoding: "#", Kind: KindClass}, nil
	case ':':
		p.skipQualifiersAndOffsets()
		return Type{Encoding: ":", Kind: KindSEL}, nil
	case 'c':
		p.skipQualifiersAndOffsets()
		return Type{Encoding: "c", Kind: KindChar}, nil
	case 'C':
		p.skipQualifiersAndOffsets()
		return Type{Encoding: "C", Kind: KindUChar}, nil
	case 's':
		p.skipQualifiersAndOffsets()
		return Type{Encoding: "s", Kind: KindShort}, nil
	case 'S':
		p.skipQualifiersAndOffsets()
		return Type{Encoding: "S", Kind: KindUShort}, nil
	case 'i':
		p.skipQualifiersAndOffsets()
		return Type{Encoding: "i", Kind: KindInt}, nil
	case 'I':
		p.skipQualifiersAndOffsets()
		return Type{Encoding: "I", Kind: KindUInt}, nil
	case 'l':
		p.skipQualifiersAndOffsets()
		return Type{Encoding: "l", Kind: KindLong}, nil
	case 'L':
		p.skipQualifiersAndOffsets()
		return Type{Encoding: "L", Kind: KindULong}, nil
	case 'q':
		p.skipQualifiersAndOffsets()
		return Type{Encoding: "q", Kind: KindLongLong}, nil
	case 'Q':
		p.skipQualifiersAndOffsets()
		return Type{Encoding: "Q", Kind: KindULongLong}, nil
	case 'f':
		p.skipQualifiersAndOffsets()
		return Type{Encoding: "f", Kind: KindFloat}, nil
	case 'd':
		p.skipQualifiersAndOffsets()
		return Type{Encoding: "d", Kind: KindDouble}, nil
	case 'B':
		p.skipQualifiersAndOffsets()
		return Type{Encoding: "B", Kind: KindBool}, nil
	case '*':
		p.skipQualifiersAndOffsets()
		return Type{Encoding: "*", Kind: KindPointer, Elem: &Type{Encoding: "c", Kind: KindChar}}, nil
	case '^':
		// Skip qualifiers but not offsets. A digit after '^' belongs to the
		// enclosing signature: it is the stack offset that follows a pointer
		// written with no pointee. Consuming it here would let the next
		// argument stand in as the pointee, so "^16@0:8q16" would report a
		// return type of ^@, lose self, and read q where _cmd belongs,
		// shifting every argument left by one.
		p.skipQualifiers()
		if p.i >= len(p.s) || (p.s[p.i] >= '0' && p.s[p.i] <= '9') || p.s[p.i] == '"' || p.s[p.i] == '}' || p.s[p.i] == ')' {
			p.skipQualifiersAndOffsets()
			return Type{Encoding: "^", Kind: KindPointer}, nil
		}
		elem, err := p.parseType()
		if err != nil {
			return Type{}, fmt.Errorf("pointer elem: %w", err)
		}
		return Type{Encoding: "^" + elem.Encoding, Kind: KindPointer, Elem: &elem}, nil
	case '{':
		depth := 1
		sStart := start
		var name string
		eq := -1
		for p.i < len(p.s) && depth > 0 {
			if p.s[p.i] == '{' {
				depth++
			} else if p.s[p.i] == '}' {
				depth--
			} else if p.s[p.i] == '=' && depth == 1 && eq == -1 {
				eq = p.i
			}
			p.i++
		}
		if depth > 0 {
			return Type{}, fmt.Errorf("unclosed struct encoding")
		}
		fullEnc := p.s[sStart:p.i]
		if eq != -1 {
			name = p.s[sStart+1 : eq]
			if name == "?" {
				name = ""
			}
			fieldsStr := p.s[eq+1 : p.i-1]
			fp := &parser{s: fieldsStr, inFields: true}
			var fields []Type
			for fp.i < len(fp.s) {
				var fieldName string
				hasQuote := false
				if fp.i < len(fp.s) && fp.s[fp.i] == '"' {
					hasQuote = true
					fp.i++
					nStart := fp.i
					for fp.i < len(fp.s) && fp.s[fp.i] != '"' {
						fp.i++
					}
					fieldName = fp.s[nStart:fp.i]
					if fp.i < len(fp.s) && fp.s[fp.i] == '"' {
						fp.i++
					}
				}
				var f Type
				if fp.i >= len(fp.s) || fp.s[fp.i] == '"' || fp.s[fp.i] == '}' || fp.s[fp.i] == ')' {
					f = Type{Encoding: "?", Kind: KindUnknown}
				} else {
					var err error
					f, err = fp.parseType()
					if err != nil {
						return Type{}, fmt.Errorf("parse struct field %d: %w", len(fields), err)
					}
				}
				f.FieldName = fieldName
				f.HasQuote = hasQuote
				fields = append(fields, f)
			}
			p.skipQualifiersAndOffsets()
			return Type{Encoding: fullEnc, Name: name, Kind: KindStruct, Fields: fields}, nil
		}
		name = p.s[sStart+1 : p.i-1]
		if name == "?" {
			name = ""
		}
		p.skipQualifiersAndOffsets()
		return Type{Encoding: fullEnc, Name: name, Kind: KindStruct}, nil
	case '(':
		depth := 1
		uStart := start
		var name string
		eq := -1
		for p.i < len(p.s) && depth > 0 {
			if p.s[p.i] == '(' {
				depth++
			} else if p.s[p.i] == ')' {
				depth--
			} else if p.s[p.i] == '=' && depth == 1 && eq == -1 {
				eq = p.i
			}
			p.i++
		}
		if depth > 0 {
			return Type{}, fmt.Errorf("unclosed union encoding")
		}
		fullEnc := p.s[uStart:p.i]
		if eq != -1 {
			name = p.s[uStart+1 : eq]
			if name == "?" {
				name = ""
			}
			fieldsStr := p.s[eq+1 : p.i-1]
			fp := &parser{s: fieldsStr, inFields: true}
			var fields []Type
			for fp.i < len(fp.s) {
				var fieldName string
				hasQuote := false
				if fp.i < len(fp.s) && fp.s[fp.i] == '"' {
					hasQuote = true
					fp.i++
					nStart := fp.i
					for fp.i < len(fp.s) && fp.s[fp.i] != '"' {
						fp.i++
					}
					fieldName = fp.s[nStart:fp.i]
					if fp.i < len(fp.s) && fp.s[fp.i] == '"' {
						fp.i++
					}
				}
				var f Type
				if fp.i >= len(fp.s) || fp.s[fp.i] == '"' || fp.s[fp.i] == '}' || fp.s[fp.i] == ')' {
					f = Type{Encoding: "?", Kind: KindUnknown}
				} else {
					var err error
					f, err = fp.parseType()
					if err != nil {
						return Type{}, fmt.Errorf("parse union field %d: %w", len(fields), err)
					}
				}
				f.FieldName = fieldName
				f.HasQuote = hasQuote
				fields = append(fields, f)
			}
			p.skipQualifiersAndOffsets()
			return Type{Encoding: fullEnc, Name: name, Kind: KindUnion, Fields: fields}, nil
		}
		name = p.s[uStart+1 : p.i-1]
		if name == "?" {
			name = ""
		}
		p.skipQualifiersAndOffsets()
		return Type{Encoding: fullEnc, Name: name, Kind: KindUnion}, nil
	case '[':
		depth := 1
		aStart := start
		for p.i < len(p.s) && depth > 0 {
			if p.s[p.i] == '[' {
				depth++
			} else if p.s[p.i] == ']' {
				depth--
			}
			p.i++
		}
		if depth > 0 {
			return Type{}, fmt.Errorf("unclosed array encoding")
		}
		fullEnc := p.s[aStart:p.i]
		inner := fullEnc[1 : len(fullEnc)-1]
		idx := 0
		for idx < len(inner) && inner[idx] >= '0' && inner[idx] <= '9' {
			idx++
		}
		var elem *Type
		if idx < len(inner) {
			ep := &parser{s: inner[idx:]}
			eType, err := ep.parseType()
			if err != nil {
				return Type{}, fmt.Errorf("parse array elem: %w", err)
			}
			elem = &eType
		}
		p.skipQualifiersAndOffsets()
		return Type{Encoding: fullEnc, Kind: KindArray, Elem: elem}, nil
	case 'b':
		for p.i < len(p.s) && p.s[p.i] >= '0' && p.s[p.i] <= '9' {
			p.i++
		}
		fullEnc := p.s[start:p.i]
		p.skipQualifiersAndOffsets()
		return Type{Encoding: fullEnc, Kind: KindBitfield}, nil
	case '?':
		p.skipQualifiersAndOffsets()
		return Type{Encoding: "?", Kind: KindUnknown}, nil
	default:
		p.skipQualifiersAndOffsets()
		return Type{Encoding: string(ch), Kind: KindUnknown}, nil
	}
}

func calcNaturalSize(fields []Type) int {
	offset := 0
	maxAlign := 1
	for _, f := range fields {
		fSize := f.NaturalSize()
		if fSize < 0 {
			return -1
		}
		align := f.Align()
		if align > maxAlign {
			maxAlign = align
		}
		if align > 0 {
			if remainder := offset % align; remainder != 0 {
				offset += align - remainder
			}
		}
		offset += fSize
	}
	if maxAlign > 1 {
		if remainder := offset % maxAlign; remainder != 0 {
			offset += maxAlign - remainder
		}
	}
	return offset
}

func formatType(t Type) string {
	prefix := ""
	if t.IsAtomic {
		prefix = "A"
		t.IsAtomic = false
	}
	return prefix + formatTypeBase(t)
}

func formatTypeBase(t Type) string {
	switch t.Kind {
	case KindVoid:
		return "v"
	case KindId:
		if t.Name != "" {
			return "@\"" + t.Name + "\""
		}
		return "@"
	case KindClass:
		return "#"
	case KindSEL:
		return ":"
	case KindBool:
		return "B"
	case KindChar:
		return "c"
	case KindUChar:
		return "C"
	case KindShort:
		return "s"
	case KindUShort:
		return "S"
	case KindInt:
		return "i"
	case KindUInt:
		return "I"
	case KindLong:
		return "l"
	case KindULong:
		return "L"
	case KindLongLong:
		return "q"
	case KindULongLong:
		return "Q"
	case KindFloat:
		return "f"
	case KindDouble:
		return "d"
	case KindBitfield:
		return t.Encoding
	case KindPointer:
		if t.Encoding == "*" {
			return "*"
		}
		if t.Elem == nil {
			return "^"
		}
		return "^" + formatType(*t.Elem)
	case KindBlock:
		return t.Encoding
	case KindStruct:
		name := t.Name
		if name == "" {
			name = "?"
		}
		if !strings.Contains(t.Encoding, "=") {
			return "{" + name + "}"
		}
		var sb strings.Builder
		sb.WriteString("{")
		sb.WriteString(name)
		sb.WriteString("=")
		for _, f := range t.Fields {
			if f.HasQuote {
				sb.WriteString("\"")
				sb.WriteString(f.FieldName)
				sb.WriteString("\"")
			}
			sb.WriteString(formatType(f))
		}
		sb.WriteString("}")
		return sb.String()
	case KindUnion:
		name := t.Name
		if name == "" {
			name = "?"
		}
		if !strings.Contains(t.Encoding, "=") {
			return "(" + name + ")"
		}
		var sb strings.Builder
		sb.WriteString("(")
		sb.WriteString(name)
		sb.WriteString("=")
		for _, f := range t.Fields {
			if f.HasQuote {
				sb.WriteString("\"")
				sb.WriteString(f.FieldName)
				sb.WriteString("\"")
			}
			sb.WriteString(formatType(f))
		}
		sb.WriteString(")")
		return sb.String()
	case KindArray:
		if t.Elem == nil {
			return t.Encoding
		}
		countStr := ""
		if strings.HasPrefix(t.Encoding, "[") {
			inner := t.Encoding[1:]
			idx := 0
			for idx < len(inner) && inner[idx] >= '0' && inner[idx] <= '9' {
				idx++
			}
			countStr = inner[:idx]
		}
		return "[" + countStr + formatType(*t.Elem) + "]"
	default:
		if t.Kind == KindUnknown && t.Encoding == "?" && t.HasQuote {
			return ""
		}
		return t.Encoding
	}
}
