// +build cgo 

#include <cstddef>
#include <string.h>
#include <stdio.h>
#include <string>

#include "compact_lang_det.h"
#include "cld2.h"

const char* DetectLang(char *data, int length) {

    bool is_plain_text = true;
    CLD2::CLDHints cldhints = {NULL, NULL, 0, CLD2::UNKNOWN_LANGUAGE};
    bool allow_extended_lang = true;
    int flags = 0;
    CLD2::Language language3[3];
    int percent3[3];
    double normalized_score3[3];
    CLD2::ResultChunkVector resultchunkvector;
    int text_bytes;
    bool is_reliable;

    if (length <= 0) {
        length = strlen(data);
    }

    CLD2::Language summary_lang = CLD2::UNKNOWN_LANGUAGE;

    summary_lang = CLD2::ExtDetectLanguageSummary(data, 
            length,
            is_plain_text,
            &cldhints,
            flags,
            language3,
            percent3,
            normalized_score3,
            &resultchunkvector,
            &text_bytes,
            &is_reliable);

    return CLD2::LanguageCode(summary_lang);
}

int DetectLangCode(char *data, int length) {
    bool is_plain_text = true;
    CLD2::CLDHints cldhints = {NULL, NULL, 0, CLD2::UNKNOWN_LANGUAGE};
    bool allow_extended_lang = true;
    int flags = 0;
    CLD2::Language language3[3];
    int percent3[3];
    double normalized_score3[3];
    CLD2::ResultChunkVector resultchunkvector;
    int text_bytes;
    bool is_reliable;

    if (length <= 0) {
        length = strlen(data);
    }

    CLD2::Language summary_lang = CLD2::UNKNOWN_LANGUAGE;

    summary_lang = CLD2::ExtDetectLanguageSummary(data, 
            length,
            is_plain_text,
            &cldhints,
            flags,
            language3,
            percent3,
            normalized_score3,
            &resultchunkvector,
            &text_bytes,
            &is_reliable);

    return int(summary_lang);
}

void DetectThree(result *dst, char *data, int length) {
    bool is_plain_text = true;
    CLD2::CLDHints cldhints = {NULL, NULL, 0, CLD2::UNKNOWN_LANGUAGE};
    bool allow_extended_lang = true;
    int flags = 0;
    CLD2::Language language3[3];
    int percent3[3];
    double normalized_score3[3];
    CLD2::ResultChunkVector resultchunkvector;
    int text_bytes;
    bool is_reliable;

    if (length <= 0) {
        length = strlen(data);
    }

    CLD2::Language summary_lang = CLD2::UNKNOWN_LANGUAGE;

    summary_lang = CLD2::ExtDetectLanguageSummary(data, 
            length,
            is_plain_text,
            &cldhints,
            flags,
            language3,
            percent3,
            normalized_score3,
            &resultchunkvector,
            &text_bytes,
            &is_reliable);

    memcpy(&dst->language[0], &language3[0], sizeof(language3));
    memcpy(&dst->percent[0], &percent3[0], sizeof(percent3));
    memcpy(&dst->normalized_score[0], &normalized_score3[0], sizeof(normalized_score3));
    dst->reliable = char(is_reliable);
    dst->text_bytes = text_bytes;
    return;
}
