#include <stdbool.h>
#include <stdlib.h>

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct {
    // This value has to be the maximum length of the values in the file
    // third_party/cld3/src/task_context_params.cc
    // Note that initially that was a pointer, but the library is not returning
    // a pointer to statically allocated memory, as everytime it will create a
    // string out of the statically allocated const char*.
    char language[7];
    int len_language;
    float probability; // Language probability.
    bool is_reliable;  // Whether the prediction is reliable.

    // Proportion of bytes associated with the language. If FindLanguage is
    // called, this variable is set to 1.
    float proportion;
  } Result;

  typedef void* CLanguageIdentifier;

  CLanguageIdentifier new_language_identifier(int minNumBytes, int maxNumBytes);
  void free_language_identifier(CLanguageIdentifier);
  const Result find_language(CLanguageIdentifier li, char *data, int length);

#ifdef __cplusplus
}
#endif

