// Code generated from Apple documentation. DO NOT EDIT.

package naturallanguage

import (
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
)

var ()

var ()

var ()

var ()

var ()

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLContextualEmbeddingKeyLanguages"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLContextualEmbeddingKeys.Languages = NLContextualEmbeddingKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLContextualEmbeddingKeyRevision"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLContextualEmbeddingKeys.Revision = NLContextualEmbeddingKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLContextualEmbeddingKeyScripts"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLContextualEmbeddingKeys.Scripts = NLContextualEmbeddingKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageAmharic"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Amharic = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageArabic"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Arabic = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageArmenian"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Armenian = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageBengali"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Bengali = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageBulgarian"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Bulgarian = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageBurmese"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Burmese = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageCatalan"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Catalan = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageCherokee"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Cherokee = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageCroatian"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Croatian = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageCzech"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Czech = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageDanish"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Danish = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageDutch"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Dutch = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageEnglish"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.English = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageFinnish"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Finnish = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageFrench"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.French = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageGeorgian"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Georgian = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageGerman"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.German = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageGreek"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Greek = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageGujarati"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Gujarati = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageHebrew"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Hebrew = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageHindi"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Hindi = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageHungarian"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Hungarian = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageIcelandic"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Icelandic = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageIndonesian"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Indonesian = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageItalian"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Italian = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageJapanese"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Japanese = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageKannada"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Kannada = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageKazakh"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Kazakh = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageKhmer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Khmer = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageKorean"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Korean = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageLao"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Lao = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageMalay"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Malay = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageMalayalam"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Malayalam = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageMarathi"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Marathi = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageMongolian"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Mongolian = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageNorwegian"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Norwegian = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageOriya"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Oriya = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguagePersian"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Persian = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguagePolish"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Polish = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguagePortuguese"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Portuguese = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguagePunjabi"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Punjabi = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageRomanian"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Romanian = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageRussian"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Russian = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageSimplifiedChinese"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.SimplifiedChinese = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageSinhalese"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Sinhalese = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageSlovak"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Slovak = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageSpanish"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Spanish = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageSwedish"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Swedish = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageTamil"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Tamil = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageTelugu"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Telugu = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageThai"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Thai = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageTibetan"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Tibetan = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageTraditionalChinese"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.TraditionalChinese = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageTurkish"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Turkish = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageUkrainian"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Ukrainian = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageUndetermined"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Undetermined = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageUrdu"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Urdu = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLLanguageVietnamese"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLLanguages.Vietnamese = NLLanguage(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLScriptArabic"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLScripts.Arabic = NLScript(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLScriptArmenian"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLScripts.Armenian = NLScript(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLScriptBengali"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLScripts.Bengali = NLScript(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLScriptCanadianAboriginalSyllabics"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLScripts.CanadianAboriginalSyllabics = NLScript(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLScriptCherokee"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLScripts.Cherokee = NLScript(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLScriptCyrillic"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLScripts.Cyrillic = NLScript(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLScriptDevanagari"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLScripts.Devanagari = NLScript(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLScriptEthiopic"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLScripts.Ethiopic = NLScript(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLScriptGeorgian"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLScripts.Georgian = NLScript(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLScriptGreek"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLScripts.Greek = NLScript(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLScriptGujarati"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLScripts.Gujarati = NLScript(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLScriptGurmukhi"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLScripts.Gurmukhi = NLScript(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLScriptHebrew"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLScripts.Hebrew = NLScript(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLScriptJapanese"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLScripts.Japanese = NLScript(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLScriptKannada"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLScripts.Kannada = NLScript(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLScriptKhmer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLScripts.Khmer = NLScript(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLScriptKorean"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLScripts.Korean = NLScript(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLScriptLao"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLScripts.Lao = NLScript(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLScriptLatin"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLScripts.Latin = NLScript(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLScriptMalayalam"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLScripts.Malayalam = NLScript(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLScriptMongolian"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLScripts.Mongolian = NLScript(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLScriptMyanmar"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLScripts.Myanmar = NLScript(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLScriptOriya"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLScripts.Oriya = NLScript(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLScriptSimplifiedChinese"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLScripts.SimplifiedChinese = NLScript(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLScriptSinhala"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLScripts.Sinhala = NLScript(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLScriptTamil"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLScripts.Tamil = NLScript(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLScriptTelugu"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLScripts.Telugu = NLScript(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLScriptThai"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLScripts.Thai = NLScript(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLScriptTibetan"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLScripts.Tibetan = NLScript(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLScriptTraditionalChinese"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLScripts.TraditionalChinese = NLScript(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLScriptUndetermined"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLScripts.Undetermined = NLScript(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLTagAdjective"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLTags.Adjective = NLTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLTagAdverb"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLTags.Adverb = NLTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLTagClassifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLTags.Classifier = NLTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLTagCloseParenthesis"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLTags.CloseParenthesis = NLTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLTagCloseQuote"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLTags.CloseQuote = NLTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLTagConjunction"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLTags.Conjunction = NLTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLTagDash"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLTags.Dash = NLTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLTagDeterminer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLTags.Determiner = NLTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLTagIdiom"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLTags.Idiom = NLTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLTagInterjection"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLTags.Interjection = NLTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLTagNoun"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLTags.Noun = NLTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLTagNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLTags.Number = NLTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLTagOpenParenthesis"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLTags.OpenParenthesis = NLTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLTagOpenQuote"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLTags.OpenQuote = NLTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLTagOrganizationName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLTags.OrganizationName = NLTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLTagOther"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLTags.Other = NLTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLTagOtherPunctuation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLTags.OtherPunctuation = NLTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLTagOtherWhitespace"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLTags.OtherWhitespace = NLTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLTagOtherWord"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLTags.OtherWord = NLTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLTagParagraphBreak"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLTags.ParagraphBreak = NLTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLTagParticle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLTags.Particle = NLTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLTagPersonalName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLTags.PersonalName = NLTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLTagPlaceName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLTags.PlaceName = NLTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLTagPreposition"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLTags.Preposition = NLTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLTagPronoun"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLTags.Pronoun = NLTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLTagPunctuation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLTags.Punctuation = NLTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLTagSchemeLanguage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLTagSchemes.Language = NLTagScheme(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLTagSchemeLemma"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLTagSchemes.Lemma = NLTagScheme(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLTagSchemeLexicalClass"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLTagSchemes.LexicalClass = NLTagScheme(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLTagSchemeNameType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLTagSchemes.NameType = NLTagScheme(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLTagSchemeNameTypeOrLexicalClass"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLTagSchemes.NameTypeOrLexicalClass = NLTagScheme(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLTagSchemeScript"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLTagSchemes.Script = NLTagScheme(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLTagSchemeSentimentScore"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLTagSchemes.SentimentScore = NLTagScheme(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLTagSchemeTokenType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLTagSchemes.TokenType = NLTagScheme(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLTagSentenceTerminator"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLTags.SentenceTerminator = NLTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLTagVerb"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLTags.Verb = NLTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLTagWhitespace"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLTags.Whitespace = NLTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLTagWord"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLTags.Word = NLTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NLTagWordJoiner"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NLTags.WordJoiner = NLTag(objc.GoString(cstr))
			}
		}
	}

}

// NLContextualEmbeddingKeys provides typed accessors for [NLContextualEmbeddingKey] constants.
var NLContextualEmbeddingKeys struct {
	// Languages: A key that identifies the languages in a contextual embedding.
	Languages NLContextualEmbeddingKey
	// Revision: A key that identifies the revision for a contextual embedding.
	Revision NLContextualEmbeddingKey
	// Scripts: A key that identifies the scripts in a contextual embedding.
	Scripts NLContextualEmbeddingKey
}

// NLLanguages provides typed accessors for [NLLanguage] constants.
var NLLanguages struct {
	// Amharic: The unique identifier string for the Amharic language.
	Amharic NLLanguage
	// Arabic: The unique identifier string for the Arabic language.
	Arabic NLLanguage
	// Armenian: The unique identifier string for the Armenian language.
	Armenian NLLanguage
	// Bengali: The unique identifier string for the Bengali language.
	Bengali NLLanguage
	// Bulgarian: The unique identifier string for the Bulgarian language.
	Bulgarian NLLanguage
	// Burmese: The unique identifier string for the Burmese language.
	Burmese NLLanguage
	// Catalan: The unique identifier string for the Catalan language.
	Catalan NLLanguage
	// Cherokee: The unique identifier string for the Cherokee language.
	Cherokee NLLanguage
	// Croatian: The unique identifier string for the Croatian language.
	Croatian NLLanguage
	// Czech: The unique identifier string for the Czech language.
	Czech NLLanguage
	// Danish: The unique identifier string for the Danish language.
	Danish NLLanguage
	// Dutch: The unique identifier string for the Dutch language.
	Dutch NLLanguage
	// English: The unique identifier string for the English language.
	English NLLanguage
	// Finnish: The unique identifier string for the Finnish language.
	Finnish NLLanguage
	// French: The unique identifier string for the French language.
	French NLLanguage
	// Georgian: The unique identifier string for the Georgian language.
	Georgian NLLanguage
	// German: The unique identifier string for the German language.
	German NLLanguage
	// Greek: The unique identifier string for the Greek language.
	Greek NLLanguage
	// Gujarati: The unique identifier string for the Gujarati language.
	Gujarati NLLanguage
	// Hebrew: The unique identifier string for the Hebrew language.
	Hebrew NLLanguage
	// Hindi: The unique identifier string for the Hindi language.
	Hindi NLLanguage
	// Hungarian: The unique identifier string for the Hungarian language.
	Hungarian NLLanguage
	// Icelandic: The unique identifier string for the Icelandic language.
	Icelandic NLLanguage
	// Indonesian: The unique identifier string for the Indonesian language.
	Indonesian NLLanguage
	// Italian: The unique identifier string for the Italian language.
	Italian NLLanguage
	// Japanese: The unique identifier string for the Japanese language.
	Japanese NLLanguage
	// Kannada: The unique identifier string for the Kannada language.
	Kannada NLLanguage
	// Kazakh: The unique identifier string for the Kazakh language.
	Kazakh NLLanguage
	// Khmer: The unique identifier string for the Khmer language.
	Khmer NLLanguage
	// Korean: The unique identifier string for the Korean language.
	Korean NLLanguage
	// Lao: The unique identifier string for the Lao language.
	Lao NLLanguage
	// Malay: The unique identifier string for the Malay language.
	Malay NLLanguage
	// Malayalam: The unique identifier string for the Malayalam language.
	Malayalam NLLanguage
	// Marathi: The unique identifier string for the Marathi language.
	Marathi NLLanguage
	// Mongolian: The unique identifier string for the Mongolian language.
	Mongolian NLLanguage
	// Norwegian: The unique identifier string for the Norwegian language.
	Norwegian NLLanguage
	// Oriya: The unique identifier string for the Oriya language.
	Oriya NLLanguage
	// Persian: The unique identifier string for the Persian language.
	Persian NLLanguage
	// Polish: The unique identifier string for the Polish language.
	Polish NLLanguage
	// Portuguese: The unique identifier string for the Portuguese language.
	Portuguese NLLanguage
	// Punjabi: The unique identifier string for the Punjabi language.
	Punjabi NLLanguage
	// Romanian: The unique identifier string for the Romanian language.
	Romanian NLLanguage
	// Russian: The unique identifier string for the Russian language.
	Russian NLLanguage
	// SimplifiedChinese: The unique identifier string for the Simplified Chinese language.
	SimplifiedChinese NLLanguage
	// Sinhalese: The unique identifier string for the Sinhalese language.
	Sinhalese NLLanguage
	// Slovak: The unique identifier string for the Slovak language.
	Slovak NLLanguage
	// Spanish: The unique identifier string for the Spanish language.
	Spanish NLLanguage
	// Swedish: The unique identifier string for the Swedish language.
	Swedish NLLanguage
	// Tamil: The unique identifier string for the Tamil language.
	Tamil NLLanguage
	// Telugu: The unique identifier string for the Telugu language.
	Telugu NLLanguage
	// Thai: The unique identifier string for the Thai language.
	Thai NLLanguage
	// Tibetan: The unique identifier string for the Tibetan language.
	Tibetan NLLanguage
	// TraditionalChinese: The unique identifier string for the Traditional Chinese language.
	TraditionalChinese NLLanguage
	// Turkish: The unique identifier string for the Turkish language.
	Turkish NLLanguage
	// Ukrainian: The unique identifier string for the Ukrainian language.
	Ukrainian NLLanguage
	// Undetermined: The unique identifier string for a language the Natural Language framework doesn’t recognize.
	Undetermined NLLanguage
	// Urdu: The unique identifier string for the Urdu language.
	Urdu NLLanguage
	// Vietnamese: The unique identifier string for the Vietnamese language.
	Vietnamese NLLanguage
}

// NLScripts provides typed accessors for [NLScript] constants.
var NLScripts struct {
	// Arabic: The unique identifier string for the Arabic script.
	Arabic NLScript
	// Armenian: The unique identifier string for the Armenian script.
	Armenian NLScript
	// Bengali: The unique identifier string for the Bengali script.
	Bengali NLScript
	// CanadianAboriginalSyllabics: The unique identifier string for the Canadian Aboriginal Syllabics.
	CanadianAboriginalSyllabics NLScript
	// Cherokee: The unique identifier string for the Cherokee script.
	Cherokee NLScript
	// Cyrillic: The unique identifier string for the Cyrillic script.
	Cyrillic NLScript
	// Devanagari: The unique identifier string for the Devanagari script.
	Devanagari NLScript
	// Ethiopic: The unique identifier string for the Ethiopic script.
	Ethiopic NLScript
	// Georgian: The unique identifier string for the Georgian script.
	Georgian NLScript
	// Greek: The unique identifier string for the Greek script.
	Greek NLScript
	// Gujarati: The unique identifier string for the Gujarati script.
	Gujarati NLScript
	// Gurmukhi: The unique identifier string for the Gurmukhi script.
	Gurmukhi NLScript
	// Hebrew: The unique identifier string for the Hebrew script.
	Hebrew NLScript
	// Japanese: The unique identifier string for the Japanese script.
	Japanese NLScript
	// Kannada: The unique identifier string for the Kannada script.
	Kannada NLScript
	// Khmer: The unique identifier string for the Khmer script.
	Khmer NLScript
	// Korean: The unique identifier string for the Korean script.
	Korean NLScript
	// Lao: The unique identifier string for the Lao script.
	Lao NLScript
	// Latin: The unique identifier string for the Latin script.
	Latin NLScript
	// Malayalam: The unique identifier string for the Malayalam script.
	Malayalam NLScript
	// Mongolian: The unique identifier string for the Mongolian script.
	Mongolian NLScript
	// Myanmar: The unique identifier string for the Myanmar script.
	Myanmar NLScript
	// Oriya: The unique identifier string for the Oriya script.
	Oriya NLScript
	// SimplifiedChinese: The unique identifier string for the simplified Chinese script.
	SimplifiedChinese NLScript
	// Sinhala: The unique identifier string for the Sinhala script.
	Sinhala NLScript
	// Tamil: The unique identifier string for the Tamil script.
	Tamil NLScript
	// Telugu: The unique identifier string for the Telugu script.
	Telugu NLScript
	// Thai: The unique identifier string for the Thai script.
	Thai NLScript
	// Tibetan: The unique identifier string for the Tibetan script.
	Tibetan NLScript
	// TraditionalChinese: The unique identifier string for the traditional Chinese script.
	TraditionalChinese NLScript
	// Undetermined: The unique identifier string for a script the Natural Language framework doesn’t recognize.
	Undetermined NLScript
}

// NLTagSchemes provides typed accessors for [NLTagScheme] constants.
var NLTagSchemes struct {
	// Language: A scheme that supplies the language for a token, if it can determine one.
	Language NLTagScheme
	// Lemma: A scheme that supplies a stem form of a word token, if known.
	Lemma NLTagScheme
	// LexicalClass: A scheme that classifies tokens according to class: part of speech, type of punctuation, or whitespace.
	LexicalClass NLTagScheme
	// NameType: A scheme that classifies tokens according to whether they are part of a named entity.
	NameType NLTagScheme
	// NameTypeOrLexicalClass: A scheme that classifies tokens corresponding to names according to [nameType](<doc://com.apple.naturallanguage/documentation/NaturalLanguage/NLTagScheme/nameType>), and classifies all other tokens according to [lexicalClass](<doc://com.apple.naturallanguage/documentation/NaturalLanguage/NLTagScheme/lexicalClass>).
	NameTypeOrLexicalClass NLTagScheme
	// Script: A scheme that supplies the script for a token, if it can determine one.
	Script NLTagScheme
	// SentimentScore: A scheme that scores text as positive, negative, or neutral based on its sentiment polarity.
	SentimentScore NLTagScheme
	// TokenType: A scheme that classifies tokens according to their broad type: word, punctuation, or whitespace.
	TokenType NLTagScheme
}

// NLTags provides typed accessors for [NLTag] constants.
var NLTags struct {
	// Adjective: A tag indicating that the token is an adjective
	Adjective NLTag
	// Adverb: A tag indicating that the token is an adverb.
	Adverb NLTag
	// Classifier: A tag indicating that the token is a classifier.
	Classifier NLTag
	// CloseParenthesis: A tag indicating that the token is a close parenthesis.
	CloseParenthesis NLTag
	// CloseQuote: A tag indicating that the token is a close quote.
	CloseQuote NLTag
	// Conjunction: A tag indicating that the token is a conjunction.
	Conjunction NLTag
	// Dash: A tag indicating that the token is a dash.
	Dash NLTag
	// Determiner: A tag indicating that the token is a determiner.
	Determiner NLTag
	// Idiom: A tag indicating that the token is an idiom.
	Idiom NLTag
	// Interjection: A tag indicating that the token is an interjection.
	Interjection NLTag
	// Noun: A tag indicating that the token is a noun.
	Noun NLTag
	// Number: A tag indicating that the token is a number.
	Number NLTag
	// OpenParenthesis: A tag indicating that the token is an open parenthesis.
	OpenParenthesis NLTag
	// OpenQuote: A tag indicating that the token is an open quote.
	OpenQuote NLTag
	// OrganizationName: A tag indicating that the token is an organization name.
	OrganizationName NLTag
	// Other: A tag indicating that the token is a non-linguistic item, such as a symbol.
	Other NLTag
	// OtherPunctuation: A tag indicating that the token is punctuation other than a kind described by other lexical classes (sentence terminator, open or close quote, open or close parenthesis, word joiner, and dash).
	OtherPunctuation NLTag
	// OtherWhitespace: A tag indicating that the token is whitespace other than a kind described by other lexical classes (paragraph break).
	OtherWhitespace NLTag
	// OtherWord: A tag indicating that the token is a word other than a kind described by other lexical classes (noun, verb, adjective, adverb, pronoun, determiner, particle, preposition, number, conjunction, interjection, classifier, and idiom).
	OtherWord NLTag
	// ParagraphBreak: A tag indicating that the token is a paragraph break.
	ParagraphBreak NLTag
	// Particle: A tag indicating that the token is a particle.
	Particle NLTag
	// PersonalName: A tag indicating that the token is a personal name.
	PersonalName NLTag
	// PlaceName: A tag indicating that the token is a place name.
	PlaceName NLTag
	// Preposition: A tag indicating that the token is a preposition.
	Preposition NLTag
	// Pronoun: A tag indicating that the token is a pronoun.
	Pronoun NLTag
	// Punctuation: A tag indicating that the token is punctuation.
	Punctuation NLTag
	// SentenceTerminator: A tag indicating that the token is punctuation at the end of a sentence.
	SentenceTerminator NLTag
	// Verb: A tag indicating that the token is a verb.
	Verb NLTag
	// Whitespace: A tag indicating that the token is white space of any sort.
	Whitespace NLTag
	// Word: A tag indicating that the token is a word.
	Word NLTag
	// WordJoiner: A tag indicating that the token is a word joiner, signifying that two tokens on each side should not be broken up.
	WordJoiner NLTag
}
