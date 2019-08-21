#include <algorithm>
#include <string>
#include "third_party/cld3/src/nnet_language_identifier.h"
#include "cld3.h"

using chrome_lang_id::NNetLanguageIdentifier;

CLanguageIdentifier new_language_identifier(int minNumBytes, int maxNumBytes) {
  NNetLanguageIdentifier* lang_id = new NNetLanguageIdentifier(minNumBytes, maxNumBytes);
  return (void *)lang_id;
}

void free_language_identifier(CLanguageIdentifier li) {
  NNetLanguageIdentifier* lang_id = (NNetLanguageIdentifier*)li;
  delete lang_id;
}

const Result find_language(CLanguageIdentifier li, char *data, int length) {
  NNetLanguageIdentifier* lang_id = (NNetLanguageIdentifier*)li;
  std::string text(data, length);
  const NNetLanguageIdentifier::Result res = lang_id->FindLanguage(text);
  Result out;

  out.len_language = res.language.length();
  res.language.copy(out.language, std::min(out.len_language, (int)sizeof(res.language)));
  out.probability = res.probability;
  out.is_reliable = res.is_reliable;
  out.proportion = res.proportion;
  return out;
}
