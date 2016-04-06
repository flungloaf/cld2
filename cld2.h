#ifdef __cplusplus
extern "C" {
#endif

typedef struct _result {
   int language[3];
   int percent[3];
   double normalized_score[3];
   int text_bytes;
   char reliable;
} result;


const char* DetectLang(char *data, int length);
int DetectLangCode(char *data, int length);
void DetectThree(result *dst, char *data, int length);

#ifdef __cplusplus
}
#endif
