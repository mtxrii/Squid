package main

func GetCharGlyphs() map[rune]string {
	var charGlyphs = map[rune]string{
		'a': "║╝",
		'b': "╚╗",
		'c': "╔╝",
		'd': "╔╗",
		'e': "║╗",
		'f': "╔═",
		'g': "═╗",
		'h': "╚╝",
		'i': "║║",
		'j': "═╝",
		'k': "══",
		'l': "╚═",
		'm': "╗╝",
		'n': "╝╝",
		'o': "╔║",
		'p': "╝╔",
		'q': "╝╗",
		'r': "╚╚",
		's': "═╚",
		't': "╔╔",
		'u': "╚║",
		'v': "═╔",
		'w': "╗╗",
		'x': "╗═",
		'y': "═║",
		'z': "╗╚",
	}
	return charGlyphs
}

func GetNumberGlyphs() map[rune]string {
	var numberGlyphs = map[rune]string{
		'1': "(ƞ",
		'2': "@#",
		'3': "Ʃ^",
		'4': "*¢",
		'5': "(Ʈ",
		'6': "+Ĳ",
		'7': "ƴƴ",
		'8': "_Ɨ",
		'9': "Ɩ)",
		'0': "Ɖ)",
	}
	return numberGlyphs
}

func GetSpecialGlyphs() map[rune]string {
	var specialGlyphs = map[rune]string{
		'.': "Ͳ╗",
		',': "═Ȝ",
		'"': ">(",
		'-': "Ϟ╝",
		';': "╚$",
		':': "ī╔",
	}
	return specialGlyphs
}

/*
Rules:

lowercase - normal charGlyph replacement
numbers - normal numberGlyph replacement
uppercase - '{' + charGlyph
spaces - stay same
special characters - if in 'specialGlyphs' replace with , else stay as defined

*/
