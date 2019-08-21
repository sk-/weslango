// Package cld3 implements language detection using the Compact Language Detector v3.
package cld3

// #include "cld3.h"
import "C"
import (
	"errors"
	"strings"
	"unsafe"
)

const (
	latinSuffix = "-Latn"
)

type LanguageDetector struct {
	li C.CLanguageIdentifier
}

var (
	ErrMaxLessThanOrEqToZero  = errors.New("cld3: maxNumBytes passed to NewLanguageIdentifier must be greater than 0")
	ErrMinLessThanZero        = errors.New("cld3: minNumBytes passed to NewLanguageIdentifier must be greater than or equal to 0")
	ErrMaxSmallerOrEqualToMin = errors.New("cld3: maxNumBytes passed to NewLanguageIdentifier must be larger than minNumBytes")

	langCodes = map[string]string{
		"af":  "afr",
		"am":  "amh",
		"ar":  "ara",
		"az":  "aze",
		"be":  "bel",
		"bg":  "bul",
		"bn":  "ben",
		"bs":  "bos",
		"ca":  "cat",
		"ceb": "ceb",
		"co":  "cos",
		"cs":  "ces",
		"cy":  "cym",
		"da":  "dan",
		"de":  "deu",
		"el":  "ell",
		"en":  "eng",
		"eo":  "epo",
		"es":  "spa",
		"et":  "est",
		"eu":  "eus",
		"fa":  "fas",
		"fi":  "fin",
		"fil": "fil",
		"fr":  "fra",
		"fy":  "fry",
		"ga":  "gle",
		"gd":  "gla",
		"gl":  "glg",
		"gu":  "guj",
		"ha":  "hau",
		"haw": "haw",
		"hi":  "hin",
		"hmn": "hmn",
		"hr":  "hrv",
		"ht":  "hat",
		"hu":  "hun",
		"hy":  "hye",
		"id":  "ind",
		"ig":  "ibo",
		"is":  "isl",
		"it":  "ita",
		"iw":  "heb",
		"ja":  "jpn",
		"jv":  "jav",
		"ka":  "kat",
		"kk":  "kaz",
		"km":  "khm",
		"kn":  "kan",
		"ko":  "kor",
		"ku":  "kur",
		"ky":  "kir",
		"la":  "lat",
		"lb":  "ltz",
		"lo":  "lao",
		"lt":  "lit",
		"lv":  "lav",
		"mg":  "mlg",
		"mi":  "mri",
		"mk":  "mkd",
		"ml":  "mal",
		"mn":  "mon",
		"mr":  "mar",
		"ms":  "msa",
		"mt":  "mlt",
		"my":  "mya",
		"ne":  "nep",
		"nl":  "nld",
		"no":  "nor",
		"ny":  "nya",
		"pa":  "pan",
		"pl":  "pol",
		"ps":  "pus",
		"pt":  "por",
		"ro":  "ron",
		"ru":  "rus",
		"sd":  "snd",
		"si":  "sin",
		"sk":  "slk",
		"sl":  "slv",
		"sm":  "smo",
		"sn":  "sna",
		"so":  "som",
		"sq":  "sqi",
		"sr":  "srp",
		"st":  "sot",
		"su":  "sun",
		"sv":  "swe",
		"sw":  "swa",
		"ta":  "tam",
		"te":  "tel",
		"tg":  "tgk",
		"th":  "tha",
		"tr":  "tur",
		"uk":  "ukr",
		"ur":  "urd",
		"uz":  "uzb",
		"vi":  "vie",
		"xh":  "xho",
		"yi":  "yid",
		"yo":  "yor",
		"zh":  "zho",
		"zu":  "zul",
	}
)

// New returns a LanguageDetector. minNumBytes is the
// minimum numbers of bytes to consider in the text before making a decision and
// maxNumBytes is the maximum of the same. Chromium uses 0 and 512, respectively
// for its i18n work. LanguageIdentifier must be deallocated explicitly with
// Free.
// TODO: use uint instead of int
func NewDetector(minNumBytes, maxNumBytes int) (*LanguageDetector, error) {
	// We do these checks even though they exist in NNetLanguageIdentifier's
	// constructor because the CLD3_CHECK calls cause inscrutable "illegal
	// instruction" crashes if they are violated.
	if maxNumBytes <= 0 {
		return nil, ErrMaxLessThanOrEqToZero
	}
	if minNumBytes < 0 {
		return nil, ErrMinLessThanZero
	}
	if maxNumBytes <= minNumBytes {
		return nil, ErrMaxSmallerOrEqualToMin
	}
	// TODO: check if we can safely convert to int regardless the size
	return &LanguageDetector{C.new_language_identifier(C.int(minNumBytes), C.int(maxNumBytes))}, nil
	// TODO maybe user runtime setfinalizer to automatically clean up
}

// TODO: maje this to set the pointer to nil
func FreeDetector(li *LanguageDetector) {
	C.free_language_identifier(li.li)
}

// FindLanguage detects the language in a given text.
func (li *LanguageDetector) FindLanguage(text string) Result {
	// TODO: maybe do the splitting of the text here.
	cs := C.CString(text)
	defer C.free(unsafe.Pointer(cs))
	res := C.find_language(li.li, cs, C.int(len(text)))
	r := Result{}
	lang := C.GoStringN(&res.language[0], res.len_language)
	r.Language = langCodes[strings.TrimSuffix(lang, latinSuffix)]
	r.Probability = float32(res.probability)
	r.IsReliable = bool(res.is_reliable)
	r.Latin = strings.HasSuffix(lang, latinSuffix)
	return r
}

type Result struct {
	// ISO 639-3 code of the detected language.
	// We use this instead of 639-1 codes, as the underlying library supports some
	// languages which don't have a two letter representation.
	Language string

	// Probability is the probability from 0 to 1 of the text being in the
	// returned Language.
	Probability float32

	// This is just the result returned by the underlying library. It is set to
	// true if the probability is >= 70%, except in case when the detected
	// language is Croatian or Serbian, in which case the required probability is
	// just 50%.
	IsReliable bool

	// Whether the text was written using the Latin script. This can be
	// true for those languages that are in a non-Latin script. Currently the only
	// ones that can have this fiels set to true are: Bulgarian, Chinese, Greek,
	// Hindi, Japanese and Russian.
	Latin bool
}
