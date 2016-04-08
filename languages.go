package cld2

import "strings"

// Single Language estimate
type Estimate struct {
	Language Language
	Percent  int // text percentage 0..100 of the top 3 languages.

	// NormScore is internal language scores as a ratio to normal score for real text in that language.
	// Scores close to 1.0 indicate normal text, while scores far away
	// from 1.0 indicate badly-skewed text or gibberish.
	NormScore float64
}

// Language is a single language.
// Note that the zero value is "ENGLISH".
type Language uint16

// Languages are probable languages of the supplied text
type Languages struct {
	Estimates []Estimate // Possible languages returned in order of confidence
	TextBytes int        // the amount of non-tag/letters-only text found
	Reliable  bool       // Does CLD2 see the result as reliable?
}

func (l Language) String() string {
	return strings.Replace(strings.Title(strings.ToLower(languageToCName[int(l)])), "_", " ", -1)
}

func (l Language) Code() string {
	return languageToCode[int(l)]
}

// NewLanguage supplies a safe way of returning a uint16
// to a Language.
// If an invalid id is supplied, UNKNOWN_LANGUAGE is returned.
func NewLanguage(id uint16) Language {
	if id >= uint16(NUM_LANGUAGES) {
		return UNKNOWN_LANGUAGE
	}
	return Language(id)
}

// LanguageFromCode returns the language associated with the
// code. Returns UNKNOWN_LANGUAGE if the code isn't known.
func LanguageFromCode(code string) Language {
	if l, ok := codeToLanguage[code]; ok {
		return l
	}
	return UNKNOWN_LANGUAGE
}

// Copied from "generated_language.h"
const (
	ENGLISH                  Language = 0   // en
	DANISH                   Language = 1   // da
	DUTCH                    Language = 2   // nl
	FINNISH                  Language = 3   // fi
	FRENCH                   Language = 4   // fr
	GERMAN                   Language = 5   // de
	HEBREW                   Language = 6   // iw
	ITALIAN                  Language = 7   // it
	JAPANESE                 Language = 8   // ja
	KOREAN                   Language = 9   // ko
	NORWEGIAN                Language = 10  // no
	POLISH                   Language = 11  // pl
	PORTUGUESE               Language = 12  // pt
	RUSSIAN                  Language = 13  // ru
	SPANISH                  Language = 14  // es
	SWEDISH                  Language = 15  // sv
	CHINESE                  Language = 16  // zh
	CZECH                    Language = 17  // cs
	GREEK                    Language = 18  // el
	ICELANDIC                Language = 19  // is
	LATVIAN                  Language = 20  // lv
	LITHUANIAN               Language = 21  // lt
	ROMANIAN                 Language = 22  // ro
	HUNGARIAN                Language = 23  // hu
	ESTONIAN                 Language = 24  // et
	TG_UNKNOWN_LANGUAGE      Language = 25  // xxx
	UNKNOWN_LANGUAGE         Language = 26  // un
	BULGARIAN                Language = 27  // bg
	CROATIAN                 Language = 28  // hr
	SERBIAN                  Language = 29  // sr
	IRISH                    Language = 30  // ga
	GALICIAN                 Language = 31  // gl
	TAGALOG                  Language = 32  // tl
	TURKISH                  Language = 33  // tr
	UKRAINIAN                Language = 34  // uk
	HINDI                    Language = 35  // hi
	MACEDONIAN               Language = 36  // mk
	BENGALI                  Language = 37  // bn
	INDONESIAN               Language = 38  // id
	LATIN                    Language = 39  // la
	MALAY                    Language = 40  // ms
	MALAYALAM                Language = 41  // ml
	WELSH                    Language = 42  // cy
	NEPALI                   Language = 43  // ne
	TELUGU                   Language = 44  // te
	ALBANIAN                 Language = 45  // sq
	TAMIL                    Language = 46  // ta
	BELARUSIAN               Language = 47  // be
	JAVANESE                 Language = 48  // jw
	OCCITAN                  Language = 49  // oc
	URDU                     Language = 50  // ur
	BIHARI                   Language = 51  // bh
	GUJARATI                 Language = 52  // gu
	THAI                     Language = 53  // th
	ARABIC                   Language = 54  // ar
	CATALAN                  Language = 55  // ca
	ESPERANTO                Language = 56  // eo
	BASQUE                   Language = 57  // eu
	INTERLINGUA              Language = 58  // ia
	KANNADA                  Language = 59  // kn
	PUNJABI                  Language = 60  // pa
	SCOTS_GAELIC             Language = 61  // gd
	SWAHILI                  Language = 62  // sw
	SLOVENIAN                Language = 63  // sl
	MARATHI                  Language = 64  // mr
	MALTESE                  Language = 65  // mt
	VIETNAMESE               Language = 66  // vi
	FRISIAN                  Language = 67  // fy
	SLOVAK                   Language = 68  // sk
	CHINESE_T                Language = 69  // zh-Hant
	FAROESE                  Language = 70  // fo
	SUNDANESE                Language = 71  // su
	UZBEK                    Language = 72  // uz
	AMHARIC                  Language = 73  // am
	AZERBAIJANI              Language = 74  // az
	GEORGIAN                 Language = 75  // ka
	TIGRINYA                 Language = 76  // ti
	PERSIAN                  Language = 77  // fa
	BOSNIAN                  Language = 78  // bs
	SINHALESE                Language = 79  // si
	NORWEGIAN_N              Language = 80  // nn
	X_81                     Language = 81  //
	X_82                     Language = 82  //
	XHOSA                    Language = 83  // xh
	ZULU                     Language = 84  // zu
	GUARANI                  Language = 85  // gn
	SESOTHO                  Language = 86  // st
	TURKMEN                  Language = 87  // tk
	KYRGYZ                   Language = 88  // ky
	BRETON                   Language = 89  // br
	TWI                      Language = 90  // tw
	YIDDISH                  Language = 91  // yi
	X_92                     Language = 92  //
	SOMALI                   Language = 93  // so
	UIGHUR                   Language = 94  // ug
	KURDISH                  Language = 95  // ku
	MONGOLIAN                Language = 96  // mn
	ARMENIAN                 Language = 97  // hy
	LAOTHIAN                 Language = 98  // lo
	SINDHI                   Language = 99  // sd
	RHAETO_ROMANCE           Language = 100 // rm
	AFRIKAANS                Language = 101 // af
	LUXEMBOURGISH            Language = 102 // lb
	BURMESE                  Language = 103 // my
	KHMER                    Language = 104 // km
	TIBETAN                  Language = 105 // bo
	DHIVEHI                  Language = 106 // dv
	CHEROKEE                 Language = 107 // chr
	SYRIAC                   Language = 108 // syr
	LIMBU                    Language = 109 // lif
	ORIYA                    Language = 110 // or
	ASSAMESE                 Language = 111 // as
	CORSICAN                 Language = 112 // co
	INTERLINGUE              Language = 113 // ie
	KAZAKH                   Language = 114 // kk
	LINGALA                  Language = 115 // ln
	X_116                    Language = 116 //
	PASHTO                   Language = 117 // ps
	QUECHUA                  Language = 118 // qu
	SHONA                    Language = 119 // sn
	TAJIK                    Language = 120 // tg
	TATAR                    Language = 121 // tt
	TONGA                    Language = 122 // to
	YORUBA                   Language = 123 // yo
	X_124                    Language = 124 //
	X_125                    Language = 125 //
	X_126                    Language = 126 //
	X_127                    Language = 127 //
	MAORI                    Language = 128 // mi
	WOLOF                    Language = 129 // wo
	ABKHAZIAN                Language = 130 // ab
	AFAR                     Language = 131 // aa
	AYMARA                   Language = 132 // ay
	BASHKIR                  Language = 133 // ba
	BISLAMA                  Language = 134 // bi
	DZONGKHA                 Language = 135 // dz
	FIJIAN                   Language = 136 // fj
	GREENLANDIC              Language = 137 // kl
	HAUSA                    Language = 138 // ha
	HAITIAN_CREOLE           Language = 139 // ht
	INUPIAK                  Language = 140 // ik
	INUKTITUT                Language = 141 // iu
	KASHMIRI                 Language = 142 // ks
	KINYARWANDA              Language = 143 // rw
	MALAGASY                 Language = 144 // mg
	NAURU                    Language = 145 // na
	OROMO                    Language = 146 // om
	RUNDI                    Language = 147 // rn
	SAMOAN                   Language = 148 // sm
	SANGO                    Language = 149 // sg
	SANSKRIT                 Language = 150 // sa
	SISWANT                  Language = 151 // ss
	TSONGA                   Language = 152 // ts
	TSWANA                   Language = 153 // tn
	VOLAPUK                  Language = 154 // vo
	ZHUANG                   Language = 155 // za
	KHASI                    Language = 156 // kha
	SCOTS                    Language = 157 // sco
	GANDA                    Language = 158 // lg
	MANX                     Language = 159 // gv
	MONTENEGRIN              Language = 160 // sr-ME
	AKAN                     Language = 161 // ak
	IGBO                     Language = 162 // ig
	MAURITIAN_CREOLE         Language = 163 // mfe
	HAWAIIAN                 Language = 164 // haw
	CEBUANO                  Language = 165 // ceb
	EWE                      Language = 166 // ee
	GA                       Language = 167 // gaa
	HMONG                    Language = 168 // blu
	KRIO                     Language = 169 // kri
	LOZI                     Language = 170 // loz
	LUBA_LULUA               Language = 171 // lua
	LUO_KENYA_AND_TANZANIA   Language = 172 // luo
	NEWARI                   Language = 173 // new
	NYANJA                   Language = 174 // ny
	OSSETIAN                 Language = 175 // os
	PAMPANGA                 Language = 176 // pam
	PEDI                     Language = 177 // nso
	RAJASTHANI               Language = 178 // raj
	SESELWA                  Language = 179 // crs
	TUMBUKA                  Language = 180 // tum
	VENDA                    Language = 181 // ve
	WARAY_PHILIPPINES        Language = 182 // war
	X_183                    Language = 183 //
	X_184                    Language = 184 //
	X_185                    Language = 185 //
	X_186                    Language = 186 //
	X_187                    Language = 187 //
	X_188                    Language = 188 //
	X_189                    Language = 189 //
	X_190                    Language = 190 //
	X_191                    Language = 191 //
	X_192                    Language = 192 //
	X_193                    Language = 193 //
	X_194                    Language = 194 //
	X_195                    Language = 195 //
	X_196                    Language = 196 //
	X_197                    Language = 197 //
	X_198                    Language = 198 //
	X_199                    Language = 199 //
	X_200                    Language = 200 //
	X_201                    Language = 201 //
	X_202                    Language = 202 //
	X_203                    Language = 203 //
	X_204                    Language = 204 //
	X_205                    Language = 205 //
	X_206                    Language = 206 //
	X_207                    Language = 207 //
	X_208                    Language = 208 //
	X_209                    Language = 209 //
	X_210                    Language = 210 //
	X_211                    Language = 211 //
	X_212                    Language = 212 //
	X_213                    Language = 213 //
	X_214                    Language = 214 //
	X_215                    Language = 215 //
	X_216                    Language = 216 //
	X_217                    Language = 217 //
	X_218                    Language = 218 //
	X_219                    Language = 219 //
	X_220                    Language = 220 //
	X_221                    Language = 221 //
	X_222                    Language = 222 //
	X_223                    Language = 223 //
	X_224                    Language = 224 //
	X_225                    Language = 225 //
	X_226                    Language = 226 //
	X_227                    Language = 227 //
	X_228                    Language = 228 //
	X_229                    Language = 229 //
	X_230                    Language = 230 //
	X_231                    Language = 231 //
	X_232                    Language = 232 //
	X_233                    Language = 233 //
	X_234                    Language = 234 //
	X_235                    Language = 235 //
	X_236                    Language = 236 //
	X_237                    Language = 237 //
	X_238                    Language = 238 //
	X_239                    Language = 239 //
	X_240                    Language = 240 //
	X_241                    Language = 241 //
	X_242                    Language = 242 //
	X_243                    Language = 243 //
	X_244                    Language = 244 //
	X_245                    Language = 245 //
	X_246                    Language = 246 //
	X_247                    Language = 247 //
	X_248                    Language = 248 //
	X_249                    Language = 249 //
	X_250                    Language = 250 //
	X_251                    Language = 251 //
	X_252                    Language = 252 //
	X_253                    Language = 253 //
	X_254                    Language = 254 //
	X_255                    Language = 255 //
	X_256                    Language = 256 //
	X_257                    Language = 257 //
	X_258                    Language = 258 //
	X_259                    Language = 259 //
	X_260                    Language = 260 //
	X_261                    Language = 261 //
	X_262                    Language = 262 //
	X_263                    Language = 263 //
	X_264                    Language = 264 //
	X_265                    Language = 265 //
	X_266                    Language = 266 //
	X_267                    Language = 267 //
	X_268                    Language = 268 //
	X_269                    Language = 269 //
	X_270                    Language = 270 //
	X_271                    Language = 271 //
	X_272                    Language = 272 //
	X_273                    Language = 273 //
	X_274                    Language = 274 //
	X_275                    Language = 275 //
	X_276                    Language = 276 //
	X_277                    Language = 277 //
	X_278                    Language = 278 //
	X_279                    Language = 279 //
	X_280                    Language = 280 //
	X_281                    Language = 281 //
	X_282                    Language = 282 //
	X_283                    Language = 283 //
	X_284                    Language = 284 //
	X_285                    Language = 285 //
	X_286                    Language = 286 //
	X_287                    Language = 287 //
	X_288                    Language = 288 //
	X_289                    Language = 289 //
	X_290                    Language = 290 //
	X_291                    Language = 291 //
	X_292                    Language = 292 //
	X_293                    Language = 293 //
	X_294                    Language = 294 //
	X_295                    Language = 295 //
	X_296                    Language = 296 //
	X_297                    Language = 297 //
	X_298                    Language = 298 //
	X_299                    Language = 299 //
	X_300                    Language = 300 //
	X_301                    Language = 301 //
	X_302                    Language = 302 //
	X_303                    Language = 303 //
	X_304                    Language = 304 //
	X_305                    Language = 305 //
	X_306                    Language = 306 //
	X_307                    Language = 307 //
	X_308                    Language = 308 //
	X_309                    Language = 309 //
	X_310                    Language = 310 //
	X_311                    Language = 311 //
	X_312                    Language = 312 //
	X_313                    Language = 313 //
	X_314                    Language = 314 //
	X_315                    Language = 315 //
	X_316                    Language = 316 //
	X_317                    Language = 317 //
	X_318                    Language = 318 //
	X_319                    Language = 319 //
	X_320                    Language = 320 //
	X_321                    Language = 321 //
	X_322                    Language = 322 //
	X_323                    Language = 323 //
	X_324                    Language = 324 //
	X_325                    Language = 325 //
	X_326                    Language = 326 //
	X_327                    Language = 327 //
	X_328                    Language = 328 //
	X_329                    Language = 329 //
	X_330                    Language = 330 //
	X_331                    Language = 331 //
	X_332                    Language = 332 //
	X_333                    Language = 333 //
	X_334                    Language = 334 //
	X_335                    Language = 335 //
	X_336                    Language = 336 //
	X_337                    Language = 337 //
	X_338                    Language = 338 //
	X_339                    Language = 339 //
	X_340                    Language = 340 //
	X_341                    Language = 341 //
	X_342                    Language = 342 //
	X_343                    Language = 343 //
	X_344                    Language = 344 //
	X_345                    Language = 345 //
	X_346                    Language = 346 //
	X_347                    Language = 347 //
	X_348                    Language = 348 //
	X_349                    Language = 349 //
	X_350                    Language = 350 //
	X_351                    Language = 351 //
	X_352                    Language = 352 //
	X_353                    Language = 353 //
	X_354                    Language = 354 //
	X_355                    Language = 355 //
	X_356                    Language = 356 //
	X_357                    Language = 357 //
	X_358                    Language = 358 //
	X_359                    Language = 359 //
	X_360                    Language = 360 //
	X_361                    Language = 361 //
	X_362                    Language = 362 //
	X_363                    Language = 363 //
	X_364                    Language = 364 //
	X_365                    Language = 365 //
	X_366                    Language = 366 //
	X_367                    Language = 367 //
	X_368                    Language = 368 //
	X_369                    Language = 369 //
	X_370                    Language = 370 //
	X_371                    Language = 371 //
	X_372                    Language = 372 //
	X_373                    Language = 373 //
	X_374                    Language = 374 //
	X_375                    Language = 375 //
	X_376                    Language = 376 //
	X_377                    Language = 377 //
	X_378                    Language = 378 //
	X_379                    Language = 379 //
	X_380                    Language = 380 //
	X_381                    Language = 381 //
	X_382                    Language = 382 //
	X_383                    Language = 383 //
	X_384                    Language = 384 //
	X_385                    Language = 385 //
	X_386                    Language = 386 //
	X_387                    Language = 387 //
	X_388                    Language = 388 //
	X_389                    Language = 389 //
	X_390                    Language = 390 //
	X_391                    Language = 391 //
	X_392                    Language = 392 //
	X_393                    Language = 393 //
	X_394                    Language = 394 //
	X_395                    Language = 395 //
	X_396                    Language = 396 //
	X_397                    Language = 397 //
	X_398                    Language = 398 //
	X_399                    Language = 399 //
	X_400                    Language = 400 //
	X_401                    Language = 401 //
	X_402                    Language = 402 //
	X_403                    Language = 403 //
	X_404                    Language = 404 //
	X_405                    Language = 405 //
	X_406                    Language = 406 //
	X_407                    Language = 407 //
	X_408                    Language = 408 //
	X_409                    Language = 409 //
	X_410                    Language = 410 //
	X_411                    Language = 411 //
	X_412                    Language = 412 //
	X_413                    Language = 413 //
	X_414                    Language = 414 //
	X_415                    Language = 415 //
	X_416                    Language = 416 //
	X_417                    Language = 417 //
	X_418                    Language = 418 //
	X_419                    Language = 419 //
	X_420                    Language = 420 //
	X_421                    Language = 421 //
	X_422                    Language = 422 //
	X_423                    Language = 423 //
	X_424                    Language = 424 //
	X_425                    Language = 425 //
	X_426                    Language = 426 //
	X_427                    Language = 427 //
	X_428                    Language = 428 //
	X_429                    Language = 429 //
	X_430                    Language = 430 //
	X_431                    Language = 431 //
	X_432                    Language = 432 //
	X_433                    Language = 433 //
	X_434                    Language = 434 //
	X_435                    Language = 435 //
	X_436                    Language = 436 //
	X_437                    Language = 437 //
	X_438                    Language = 438 //
	X_439                    Language = 439 //
	X_440                    Language = 440 //
	X_441                    Language = 441 //
	X_442                    Language = 442 //
	X_443                    Language = 443 //
	X_444                    Language = 444 //
	X_445                    Language = 445 //
	X_446                    Language = 446 //
	X_447                    Language = 447 //
	X_448                    Language = 448 //
	X_449                    Language = 449 //
	X_450                    Language = 450 //
	X_451                    Language = 451 //
	X_452                    Language = 452 //
	X_453                    Language = 453 //
	X_454                    Language = 454 //
	X_455                    Language = 455 //
	X_456                    Language = 456 //
	X_457                    Language = 457 //
	X_458                    Language = 458 //
	X_459                    Language = 459 //
	X_460                    Language = 460 //
	X_461                    Language = 461 //
	X_462                    Language = 462 //
	X_463                    Language = 463 //
	X_464                    Language = 464 //
	X_465                    Language = 465 //
	X_466                    Language = 466 //
	X_467                    Language = 467 //
	X_468                    Language = 468 //
	X_469                    Language = 469 //
	X_470                    Language = 470 //
	X_471                    Language = 471 //
	X_472                    Language = 472 //
	X_473                    Language = 473 //
	X_474                    Language = 474 //
	X_475                    Language = 475 //
	X_476                    Language = 476 //
	X_477                    Language = 477 //
	X_478                    Language = 478 //
	X_479                    Language = 479 //
	X_480                    Language = 480 //
	X_481                    Language = 481 //
	X_482                    Language = 482 //
	X_483                    Language = 483 //
	X_484                    Language = 484 //
	X_485                    Language = 485 //
	X_486                    Language = 486 //
	X_487                    Language = 487 //
	X_488                    Language = 488 //
	X_489                    Language = 489 //
	X_490                    Language = 490 //
	X_491                    Language = 491 //
	X_492                    Language = 492 //
	X_493                    Language = 493 //
	X_494                    Language = 494 //
	X_495                    Language = 495 //
	X_496                    Language = 496 //
	X_497                    Language = 497 //
	X_498                    Language = 498 //
	X_499                    Language = 499 //
	X_500                    Language = 500 //
	X_501                    Language = 501 //
	X_502                    Language = 502 //
	X_503                    Language = 503 //
	X_504                    Language = 504 //
	X_505                    Language = 505 //
	NDEBELE                  Language = 506 // nr
	X_BORK_BORK_BORK         Language = 507 // zzb
	X_PIG_LATIN              Language = 508 // zzp
	X_HACKER                 Language = 509 // zzh
	X_KLINGON                Language = 510 // tlh
	X_ELMER_FUDD             Language = 511 // zze
	X_Common                 Language = 512 // xx-Zyyy
	X_Latin                  Language = 513 // xx-Latn
	X_Greek                  Language = 514 // xx-Grek
	X_Cyrillic               Language = 515 // xx-Cyrl
	X_Armenian               Language = 516 // xx-Armn
	X_Hebrew                 Language = 517 // xx-Hebr
	X_Arabic                 Language = 518 // xx-Arab
	X_Syriac                 Language = 519 // xx-Syrc
	X_Thaana                 Language = 520 // xx-Thaa
	X_Devanagari             Language = 521 // xx-Deva
	X_Bengali                Language = 522 // xx-Beng
	X_Gurmukhi               Language = 523 // xx-Guru
	X_Gujarati               Language = 524 // xx-Gujr
	X_Oriya                  Language = 525 // xx-Orya
	X_Tamil                  Language = 526 // xx-Taml
	X_Telugu                 Language = 527 // xx-Telu
	X_Kannada                Language = 528 // xx-Knda
	X_Malayalam              Language = 529 // xx-Mlym
	X_Sinhala                Language = 530 // xx-Sinh
	X_Thai                   Language = 531 // xx-Thai
	X_Lao                    Language = 532 // xx-Laoo
	X_Tibetan                Language = 533 // xx-Tibt
	X_Myanmar                Language = 534 // xx-Mymr
	X_Georgian               Language = 535 // xx-Geor
	X_Hangul                 Language = 536 // xx-Hang
	X_Ethiopic               Language = 537 // xx-Ethi
	X_Cherokee               Language = 538 // xx-Cher
	X_Canadian_Aboriginal    Language = 539 // xx-Cans
	X_Ogham                  Language = 540 // xx-Ogam
	X_Runic                  Language = 541 // xx-Runr
	X_Khmer                  Language = 542 // xx-Khmr
	X_Mongolian              Language = 543 // xx-Mong
	X_Hiragana               Language = 544 // xx-Hira
	X_Katakana               Language = 545 // xx-Kana
	X_Bopomofo               Language = 546 // xx-Bopo
	X_Han                    Language = 547 // xx-Hani
	X_Yi                     Language = 548 // xx-Yiii
	X_Old_Italic             Language = 549 // xx-Ital
	X_Gothic                 Language = 550 // xx-Goth
	X_Deseret                Language = 551 // xx-Dsrt
	X_Inherited              Language = 552 // xx-Qaai
	X_Tagalog                Language = 553 // xx-Tglg
	X_Hanunoo                Language = 554 // xx-Hano
	X_Buhid                  Language = 555 // xx-Buhd
	X_Tagbanwa               Language = 556 // xx-Tagb
	X_Limbu                  Language = 557 // xx-Limb
	X_Tai_Le                 Language = 558 // xx-Tale
	X_Linear_B               Language = 559 // xx-Linb
	X_Ugaritic               Language = 560 // xx-Ugar
	X_Shavian                Language = 561 // xx-Shaw
	X_Osmanya                Language = 562 // xx-Osma
	X_Cypriot                Language = 563 // xx-Cprt
	X_Braille                Language = 564 // xx-Brai
	X_Buginese               Language = 565 // xx-Bugi
	X_Coptic                 Language = 566 // xx-Copt
	X_New_Tai_Lue            Language = 567 // xx-Talu
	X_Glagolitic             Language = 568 // xx-Glag
	X_Tifinagh               Language = 569 // xx-Tfng
	X_Syloti_Nagri           Language = 570 // xx-Sylo
	X_Old_Persian            Language = 571 // xx-Xpeo
	X_Kharoshthi             Language = 572 // xx-Khar
	X_Balinese               Language = 573 // xx-Bali
	X_Cuneiform              Language = 574 // xx-Xsux
	X_Phoenician             Language = 575 // xx-Phnx
	X_Phags_Pa               Language = 576 // xx-Phag
	X_Nko                    Language = 577 // xx-Nkoo
	X_Sundanese              Language = 578 // xx-Sund
	X_Lepcha                 Language = 579 // xx-Lepc
	X_Ol_Chiki               Language = 580 // xx-Olck
	X_Vai                    Language = 581 // xx-Vaii
	X_Saurashtra             Language = 582 // xx-Saur
	X_Kayah_Li               Language = 583 // xx-Kali
	X_Rejang                 Language = 584 // xx-Rjng
	X_Lycian                 Language = 585 // xx-Lyci
	X_Carian                 Language = 586 // xx-Cari
	X_Lydian                 Language = 587 // xx-Lydi
	X_Cham                   Language = 588 // xx-Cham
	X_Tai_Tham               Language = 589 // xx-Lana
	X_Tai_Viet               Language = 590 // xx-Tavt
	X_Avestan                Language = 591 // xx-Avst
	X_Egyptian_Hieroglyphs   Language = 592 // xx-Egyp
	X_Samaritan              Language = 593 // xx-Samr
	X_Lisu                   Language = 594 // xx-Lisu
	X_Bamum                  Language = 595 // xx-Bamu
	X_Javanese               Language = 596 // xx-Java
	X_Meetei_Mayek           Language = 597 // xx-Mtei
	X_Imperial_Aramaic       Language = 598 // xx-Armi
	X_Old_South_Arabian      Language = 599 // xx-Sarb
	X_Inscriptional_Parthian Language = 600 // xx-Prti
	X_Inscriptional_Pahlavi  Language = 601 // xx-Phli
	X_Old_Turkic             Language = 602 // xx-Orkh
	X_Kaithi                 Language = 603 // xx-Kthi
	X_Batak                  Language = 604 // xx-Batk
	X_Brahmi                 Language = 605 // xx-Brah
	X_Mandaic                Language = 606 // xx-Mand
	X_Chakma                 Language = 607 // xx-Cakm
	X_Meroitic_Cursive       Language = 608 // xx-Merc
	X_Meroitic_Hieroglyphs   Language = 609 // xx-Mero
	X_Miao                   Language = 610 // xx-Plrd
	X_Sharada                Language = 611 // xx-Shrd
	X_Sora_Sompeng           Language = 612 // xx-Sora
	X_Takri                  Language = 613 // xx-Takr
	NUM_LANGUAGES            Language = 614
)

var codeToLanguage = make(map[string]Language)

func init() {
	for i, code := range languageToCode {
		if code != "" {
			codeToLanguage[code] = Language(i)
		}
	}
}

var languageToCode = []string{
	"en",      // 0 ENGLISH
	"da",      // 1 DANISH
	"nl",      // 2 DUTCH
	"fi",      // 3 FINNISH
	"fr",      // 4 FRENCH
	"de",      // 5 GERMAN
	"iw",      // 6 HEBREW
	"it",      // 7 ITALIAN
	"ja",      // 8 Japanese
	"ko",      // 9 Korean
	"no",      // 10 NORWEGIAN
	"pl",      // 11 POLISH
	"pt",      // 12 PORTUGUESE
	"ru",      // 13 RUSSIAN
	"es",      // 14 SPANISH
	"sv",      // 15 SWEDISH
	"zh",      // 16 Chinese
	"cs",      // 17 CZECH
	"el",      // 18 GREEK
	"is",      // 19 ICELANDIC
	"lv",      // 20 LATVIAN
	"lt",      // 21 LITHUANIAN
	"ro",      // 22 ROMANIAN
	"hu",      // 23 HUNGARIAN
	"et",      // 24 ESTONIAN
	"xxx",     // 25 Ignore
	"un",      // 26 Unknown
	"bg",      // 27 BULGARIAN
	"hr",      // 28 CROATIAN
	"sr",      // 29 SERBIAN
	"ga",      // 30 IRISH
	"gl",      // 31 GALICIAN
	"tl",      // 32 TAGALOG
	"tr",      // 33 TURKISH
	"uk",      // 34 UKRAINIAN
	"hi",      // 35 HINDI
	"mk",      // 36 MACEDONIAN
	"bn",      // 37 BENGALI
	"id",      // 38 INDONESIAN
	"la",      // 39 LATIN
	"ms",      // 40 MALAY
	"ml",      // 41 MALAYALAM
	"cy",      // 42 WELSH
	"ne",      // 43 NEPALI
	"te",      // 44 TELUGU
	"sq",      // 45 ALBANIAN
	"ta",      // 46 TAMIL
	"be",      // 47 BELARUSIAN
	"jw",      // 48 JAVANESE
	"oc",      // 49 OCCITAN
	"ur",      // 50 URDU
	"bh",      // 51 BIHARI
	"gu",      // 52 GUJARATI
	"th",      // 53 THAI
	"ar",      // 54 ARABIC
	"ca",      // 55 CATALAN
	"eo",      // 56 ESPERANTO
	"eu",      // 57 BASQUE
	"ia",      // 58 INTERLINGUA
	"kn",      // 59 KANNADA
	"pa",      // 60 PUNJABI
	"gd",      // 61 SCOTS_GAELIC
	"sw",      // 62 SWAHILI
	"sl",      // 63 SLOVENIAN
	"mr",      // 64 MARATHI
	"mt",      // 65 MALTESE
	"vi",      // 66 VIETNAMESE
	"fy",      // 67 FRISIAN
	"sk",      // 68 SLOVAK
	"zh-Hant", // 69 ChineseT
	"fo",      // 70 FAROESE
	"su",      // 71 SUNDANESE
	"uz",      // 72 UZBEK
	"am",      // 73 AMHARIC
	"az",      // 74 AZERBAIJANI
	"ka",      // 75 GEORGIAN
	"ti",      // 76 TIGRINYA
	"fa",      // 77 PERSIAN
	"bs",      // 78 BOSNIAN
	"si",      // 79 SINHALESE
	"nn",      // 80 NORWEGIAN_N
	"",        // 81 81
	"",        // 82 82
	"xh",      // 83 XHOSA
	"zu",      // 84 ZULU
	"gn",      // 85 GUARANI
	"st",      // 86 SESOTHO
	"tk",      // 87 TURKMEN
	"ky",      // 88 KYRGYZ
	"br",      // 89 BRETON
	"tw",      // 90 TWI
	"yi",      // 91 YIDDISH
	"",        // 92 92
	"so",      // 93 SOMALI
	"ug",      // 94 UIGHUR
	"ku",      // 95 KURDISH
	"mn",      // 96 MONGOLIAN
	"hy",      // 97 ARMENIAN
	"lo",      // 98 LAOTHIAN
	"sd",      // 99 SINDHI
	"rm",      // 100 RHAETO_ROMANCE
	"af",      // 101 AFRIKAANS
	"lb",      // 102 LUXEMBOURGISH
	"my",      // 103 BURMESE
	"km",      // 104 KHMER
	"bo",      // 105 TIBETAN
	"dv",      // 106 DHIVEHI
	"chr",     // 107 CHEROKEE
	"syr",     // 108 SYRIAC
	"lif",     // 109 LIMBU
	"or",      // 110 ORIYA
	"as",      // 111 ASSAMESE
	"co",      // 112 CORSICAN
	"ie",      // 113 INTERLINGUE
	"kk",      // 114 KAZAKH
	"ln",      // 115 LINGALA
	"",        // 116 116
	"ps",      // 117 PASHTO
	"qu",      // 118 QUECHUA
	"sn",      // 119 SHONA
	"tg",      // 120 TAJIK
	"tt",      // 121 TATAR
	"to",      // 122 TONGA
	"yo",      // 123 YORUBA
	"",        // 124 124
	"",        // 125 125
	"",        // 126 126
	"",        // 127 127
	"mi",      // 128 MAORI
	"wo",      // 129 WOLOF
	"ab",      // 130 ABKHAZIAN
	"aa",      // 131 AFAR
	"ay",      // 132 AYMARA
	"ba",      // 133 BASHKIR
	"bi",      // 134 BISLAMA
	"dz",      // 135 DZONGKHA
	"fj",      // 136 FIJIAN
	"kl",      // 137 GREENLANDIC
	"ha",      // 138 HAUSA
	"ht",      // 139 HAITIAN_CREOLE
	"ik",      // 140 INUPIAK
	"iu",      // 141 INUKTITUT
	"ks",      // 142 KASHMIRI
	"rw",      // 143 KINYARWANDA
	"mg",      // 144 MALAGASY
	"na",      // 145 NAURU
	"om",      // 146 OROMO
	"rn",      // 147 RUNDI
	"sm",      // 148 SAMOAN
	"sg",      // 149 SANGO
	"sa",      // 150 SANSKRIT
	"ss",      // 151 SISWANT
	"ts",      // 152 TSONGA
	"tn",      // 153 TSWANA
	"vo",      // 154 VOLAPUK
	"za",      // 155 ZHUANG
	"kha",     // 156 KHASI
	"sco",     // 157 SCOTS
	"lg",      // 158 GANDA
	"gv",      // 159 MANX
	"sr-ME",   // 160 MONTENEGRIN
	"ak",      // 161 AKAN
	"ig",      // 162 IGBO
	"mfe",     // 163 MAURITIAN_CREOLE
	"haw",     // 164 HAWAIIAN
	"ceb",     // 165 CEBUANO
	"ee",      // 166 EWE
	"gaa",     // 167 GA
	"hmn",     // 168 HMONG
	"kri",     // 169 KRIO
	"loz",     // 170 LOZI
	"lua",     // 171 LUBA_LULUA
	"luo",     // 172 LUO_KENYA_AND_TANZANIA
	"new",     // 173 NEWARI
	"ny",      // 174 NYANJA
	"os",      // 175 OSSETIAN
	"pam",     // 176 PAMPANGA
	"nso",     // 177 PEDI
	"raj",     // 178 RAJASTHANI
	"crs",     // 179 SESELWA
	"tum",     // 180 TUMBUKA
	"ve",      // 181 VENDA
	"war",     // 182 WARAY_PHILIPPINES
	"",        // 183 183
	"",        // 184 184
	"",        // 185 185
	"",        // 186 186
	"",        // 187 187
	"",        // 188 188
	"",        // 189 189
	"",        // 190 190
	"",        // 191 191
	"",        // 192 192
	"",        // 193 193
	"",        // 194 194
	"",        // 195 195
	"",        // 196 196
	"",        // 197 197
	"",        // 198 198
	"",        // 199 199
	"",        // 200 200
	"",        // 201 201
	"",        // 202 202
	"",        // 203 203
	"",        // 204 204
	"",        // 205 205
	"",        // 206 206
	"",        // 207 207
	"",        // 208 208
	"",        // 209 209
	"",        // 210 210
	"",        // 211 211
	"",        // 212 212
	"",        // 213 213
	"",        // 214 214
	"",        // 215 215
	"",        // 216 216
	"",        // 217 217
	"",        // 218 218
	"",        // 219 219
	"",        // 220 220
	"",        // 221 221
	"",        // 222 222
	"",        // 223 223
	"",        // 224 224
	"",        // 225 225
	"",        // 226 226
	"",        // 227 227
	"",        // 228 228
	"",        // 229 229
	"",        // 230 230
	"",        // 231 231
	"",        // 232 232
	"",        // 233 233
	"",        // 234 234
	"",        // 235 235
	"",        // 236 236
	"",        // 237 237
	"",        // 238 238
	"",        // 239 239
	"",        // 240 240
	"",        // 241 241
	"",        // 242 242
	"",        // 243 243
	"",        // 244 244
	"",        // 245 245
	"",        // 246 246
	"",        // 247 247
	"",        // 248 248
	"",        // 249 249
	"",        // 250 250
	"",        // 251 251
	"",        // 252 252
	"",        // 253 253
	"",        // 254 254
	"",        // 255 255
	"",        // 256 256
	"",        // 257 257
	"",        // 258 258
	"",        // 259 259
	"",        // 260 260
	"",        // 261 261
	"",        // 262 262
	"",        // 263 263
	"",        // 264 264
	"",        // 265 265
	"",        // 266 266
	"",        // 267 267
	"",        // 268 268
	"",        // 269 269
	"",        // 270 270
	"",        // 271 271
	"",        // 272 272
	"",        // 273 273
	"",        // 274 274
	"",        // 275 275
	"",        // 276 276
	"",        // 277 277
	"",        // 278 278
	"",        // 279 279
	"",        // 280 280
	"",        // 281 281
	"",        // 282 282
	"",        // 283 283
	"",        // 284 284
	"",        // 285 285
	"",        // 286 286
	"",        // 287 287
	"",        // 288 288
	"",        // 289 289
	"",        // 290 290
	"",        // 291 291
	"",        // 292 292
	"",        // 293 293
	"",        // 294 294
	"",        // 295 295
	"",        // 296 296
	"",        // 297 297
	"",        // 298 298
	"",        // 299 299
	"",        // 300 300
	"",        // 301 301
	"",        // 302 302
	"",        // 303 303
	"",        // 304 304
	"",        // 305 305
	"",        // 306 306
	"",        // 307 307
	"",        // 308 308
	"",        // 309 309
	"",        // 310 310
	"",        // 311 311
	"",        // 312 312
	"",        // 313 313
	"",        // 314 314
	"",        // 315 315
	"",        // 316 316
	"",        // 317 317
	"",        // 318 318
	"",        // 319 319
	"",        // 320 320
	"",        // 321 321
	"",        // 322 322
	"",        // 323 323
	"",        // 324 324
	"",        // 325 325
	"",        // 326 326
	"",        // 327 327
	"",        // 328 328
	"",        // 329 329
	"",        // 330 330
	"",        // 331 331
	"",        // 332 332
	"",        // 333 333
	"",        // 334 334
	"",        // 335 335
	"",        // 336 336
	"",        // 337 337
	"",        // 338 338
	"",        // 339 339
	"",        // 340 340
	"",        // 341 341
	"",        // 342 342
	"",        // 343 343
	"",        // 344 344
	"",        // 345 345
	"",        // 346 346
	"",        // 347 347
	"",        // 348 348
	"",        // 349 349
	"",        // 350 350
	"",        // 351 351
	"",        // 352 352
	"",        // 353 353
	"",        // 354 354
	"",        // 355 355
	"",        // 356 356
	"",        // 357 357
	"",        // 358 358
	"",        // 359 359
	"",        // 360 360
	"",        // 361 361
	"",        // 362 362
	"",        // 363 363
	"",        // 364 364
	"",        // 365 365
	"",        // 366 366
	"",        // 367 367
	"",        // 368 368
	"",        // 369 369
	"",        // 370 370
	"",        // 371 371
	"",        // 372 372
	"",        // 373 373
	"",        // 374 374
	"",        // 375 375
	"",        // 376 376
	"",        // 377 377
	"",        // 378 378
	"",        // 379 379
	"",        // 380 380
	"",        // 381 381
	"",        // 382 382
	"",        // 383 383
	"",        // 384 384
	"",        // 385 385
	"",        // 386 386
	"",        // 387 387
	"",        // 388 388
	"",        // 389 389
	"",        // 390 390
	"",        // 391 391
	"",        // 392 392
	"",        // 393 393
	"",        // 394 394
	"",        // 395 395
	"",        // 396 396
	"",        // 397 397
	"",        // 398 398
	"",        // 399 399
	"",        // 400 400
	"",        // 401 401
	"",        // 402 402
	"",        // 403 403
	"",        // 404 404
	"",        // 405 405
	"",        // 406 406
	"",        // 407 407
	"",        // 408 408
	"",        // 409 409
	"",        // 410 410
	"",        // 411 411
	"",        // 412 412
	"",        // 413 413
	"",        // 414 414
	"",        // 415 415
	"",        // 416 416
	"",        // 417 417
	"",        // 418 418
	"",        // 419 419
	"",        // 420 420
	"",        // 421 421
	"",        // 422 422
	"",        // 423 423
	"",        // 424 424
	"",        // 425 425
	"",        // 426 426
	"",        // 427 427
	"",        // 428 428
	"",        // 429 429
	"",        // 430 430
	"",        // 431 431
	"",        // 432 432
	"",        // 433 433
	"",        // 434 434
	"",        // 435 435
	"",        // 436 436
	"",        // 437 437
	"",        // 438 438
	"",        // 439 439
	"",        // 440 440
	"",        // 441 441
	"",        // 442 442
	"",        // 443 443
	"",        // 444 444
	"",        // 445 445
	"",        // 446 446
	"",        // 447 447
	"",        // 448 448
	"",        // 449 449
	"",        // 450 450
	"",        // 451 451
	"",        // 452 452
	"",        // 453 453
	"",        // 454 454
	"",        // 455 455
	"",        // 456 456
	"",        // 457 457
	"",        // 458 458
	"",        // 459 459
	"",        // 460 460
	"",        // 461 461
	"",        // 462 462
	"",        // 463 463
	"",        // 464 464
	"",        // 465 465
	"",        // 466 466
	"",        // 467 467
	"",        // 468 468
	"",        // 469 469
	"",        // 470 470
	"",        // 471 471
	"",        // 472 472
	"",        // 473 473
	"",        // 474 474
	"",        // 475 475
	"",        // 476 476
	"",        // 477 477
	"",        // 478 478
	"",        // 479 479
	"",        // 480 480
	"",        // 481 481
	"",        // 482 482
	"",        // 483 483
	"",        // 484 484
	"",        // 485 485
	"",        // 486 486
	"",        // 487 487
	"",        // 488 488
	"",        // 489 489
	"",        // 490 490
	"",        // 491 491
	"",        // 492 492
	"",        // 493 493
	"",        // 494 494
	"",        // 495 495
	"",        // 496 496
	"",        // 497 497
	"",        // 498 498
	"",        // 499 499
	"",        // 500 500
	"",        // 501 501
	"",        // 502 502
	"",        // 503 503
	"",        // 504 504
	"",        // 505 505
	"nr",      // 506 NDEBELE
	"zzb",     // 507 X_BORK_BORK_BORK
	"zzp",     // 508 X_PIG_LATIN
	"zzh",     // 509 X_HACKER
	"tlh",     // 510 X_KLINGON
	"zze",     // 511 X_ELMER_FUDD
	"xx-Zyyy", // 512 X_Common
	"xx-Latn", // 513 X_Latin
	"xx-Grek", // 514 X_Greek
	"xx-Cyrl", // 515 X_Cyrillic
	"xx-Armn", // 516 X_Armenian
	"xx-Hebr", // 517 X_Hebrew
	"xx-Arab", // 518 X_Arabic
	"xx-Syrc", // 519 X_Syriac
	"xx-Thaa", // 520 X_Thaana
	"xx-Deva", // 521 X_Devanagari
	"xx-Beng", // 522 X_Bengali
	"xx-Guru", // 523 X_Gurmukhi
	"xx-Gujr", // 524 X_Gujarati
	"xx-Orya", // 525 X_Oriya
	"xx-Taml", // 526 X_Tamil
	"xx-Telu", // 527 X_Telugu
	"xx-Knda", // 528 X_Kannada
	"xx-Mlym", // 529 X_Malayalam
	"xx-Sinh", // 530 X_Sinhala
	"xx-Thai", // 531 X_Thai
	"xx-Laoo", // 532 X_Lao
	"xx-Tibt", // 533 X_Tibetan
	"xx-Mymr", // 534 X_Myanmar
	"xx-Geor", // 535 X_Georgian
	"xx-Hang", // 536 X_Hangul
	"xx-Ethi", // 537 X_Ethiopic
	"xx-Cher", // 538 X_Cherokee
	"xx-Cans", // 539 X_Canadian_Aboriginal
	"xx-Ogam", // 540 X_Ogham
	"xx-Runr", // 541 X_Runic
	"xx-Khmr", // 542 X_Khmer
	"xx-Mong", // 543 X_Mongolian
	"xx-Hira", // 544 X_Hiragana
	"xx-Kana", // 545 X_Katakana
	"xx-Bopo", // 546 X_Bopomofo
	"xx-Hani", // 547 X_Han
	"xx-Yiii", // 548 X_Yi
	"xx-Ital", // 549 X_Old_Italic
	"xx-Goth", // 550 X_Gothic
	"xx-Dsrt", // 551 X_Deseret
	"xx-Qaai", // 552 X_Inherited
	"xx-Tglg", // 553 X_Tagalog
	"xx-Hano", // 554 X_Hanunoo
	"xx-Buhd", // 555 X_Buhid
	"xx-Tagb", // 556 X_Tagbanwa
	"xx-Limb", // 557 X_Limbu
	"xx-Tale", // 558 X_Tai_Le
	"xx-Linb", // 559 X_Linear_B
	"xx-Ugar", // 560 X_Ugaritic
	"xx-Shaw", // 561 X_Shavian
	"xx-Osma", // 562 X_Osmanya
	"xx-Cprt", // 563 X_Cypriot
	"xx-Brai", // 564 X_Braille
	"xx-Bugi", // 565 X_Buginese
	"xx-Copt", // 566 X_Coptic
	"xx-Talu", // 567 X_New_Tai_Lue
	"xx-Glag", // 568 X_Glagolitic
	"xx-Tfng", // 569 X_Tifinagh
	"xx-Sylo", // 570 X_Syloti_Nagri
	"xx-Xpeo", // 571 X_Old_Persian
	"xx-Khar", // 572 X_Kharoshthi
	"xx-Bali", // 573 X_Balinese
	"xx-Xsux", // 574 X_Cuneiform
	"xx-Phnx", // 575 X_Phoenician
	"xx-Phag", // 576 X_Phags_Pa
	"xx-Nkoo", // 577 X_Nko
	"xx-Sund", // 578 X_Sundanese
	"xx-Lepc", // 579 X_Lepcha
	"xx-Olck", // 580 X_Ol_Chiki
	"xx-Vaii", // 581 X_Vai
	"xx-Saur", // 582 X_Saurashtra
	"xx-Kali", // 583 X_Kayah_Li
	"xx-Rjng", // 584 X_Rejang
	"xx-Lyci", // 585 X_Lycian
	"xx-Cari", // 586 X_Carian
	"xx-Lydi", // 587 X_Lydian
	"xx-Cham", // 588 X_Cham
	"xx-Lana", // 589 X_Tai_Tham
	"xx-Tavt", // 590 X_Tai_Viet
	"xx-Avst", // 591 X_Avestan
	"xx-Egyp", // 592 X_Egyptian_Hieroglyphs
	"xx-Samr", // 593 X_Samaritan
	"xx-Lisu", // 594 X_Lisu
	"xx-Bamu", // 595 X_Bamum
	"xx-Java", // 596 X_Javanese
	"xx-Mtei", // 597 X_Meetei_Mayek
	"xx-Armi", // 598 X_Imperial_Aramaic
	"xx-Sarb", // 599 X_Old_South_Arabian
	"xx-Prti", // 600 X_Inscriptional_Parthian
	"xx-Phli", // 601 X_Inscriptional_Pahlavi
	"xx-Orkh", // 602 X_Old_Turkic
	"xx-Kthi", // 603 X_Kaithi
	"xx-Batk", // 604 X_Batak
	"xx-Brah", // 605 X_Brahmi
	"xx-Mand", // 606 X_Mandaic
	"xx-Cakm", // 607 X_Chakma
	"xx-Merc", // 608 X_Meroitic_Cursive
	"xx-Mero", // 609 X_Meroitic_Hieroglyphs
	"xx-Plrd", // 610 X_Miao
	"xx-Shrd", // 611 X_Sharada
	"xx-Sora", // 612 X_Sora_Sompeng
	"xx-Takr", // 613 X_Takri
}

var languageToCName = []string{
	"ENGLISH",                  // 0 en
	"DANISH",                   // 1 da
	"DUTCH",                    // 2 nl
	"FINNISH",                  // 3 fi
	"FRENCH",                   // 4 fr
	"GERMAN",                   // 5 de
	"HEBREW",                   // 6 iw
	"ITALIAN",                  // 7 it
	"JAPANESE",                 // 8 ja
	"KOREAN",                   // 9 ko
	"NORWEGIAN",                // 10 no
	"POLISH",                   // 11 pl
	"PORTUGUESE",               // 12 pt
	"RUSSIAN",                  // 13 ru
	"SPANISH",                  // 14 es
	"SWEDISH",                  // 15 sv
	"CHINESE",                  // 16 zh
	"CZECH",                    // 17 cs
	"GREEK",                    // 18 el
	"ICELANDIC",                // 19 is
	"LATVIAN",                  // 20 lv
	"LITHUANIAN",               // 21 lt
	"ROMANIAN",                 // 22 ro
	"HUNGARIAN",                // 23 hu
	"ESTONIAN",                 // 24 et
	"TG_UNKNOWN_LANGUAGE",      // 25 xxx
	"UNKNOWN_LANGUAGE",         // 26 un
	"BULGARIAN",                // 27 bg
	"CROATIAN",                 // 28 hr
	"SERBIAN",                  // 29 sr
	"IRISH",                    // 30 ga
	"GALICIAN",                 // 31 gl
	"TAGALOG",                  // 32 tl
	"TURKISH",                  // 33 tr
	"UKRAINIAN",                // 34 uk
	"HINDI",                    // 35 hi
	"MACEDONIAN",               // 36 mk
	"BENGALI",                  // 37 bn
	"INDONESIAN",               // 38 id
	"LATIN",                    // 39 la
	"MALAY",                    // 40 ms
	"MALAYALAM",                // 41 ml
	"WELSH",                    // 42 cy
	"NEPALI",                   // 43 ne
	"TELUGU",                   // 44 te
	"ALBANIAN",                 // 45 sq
	"TAMIL",                    // 46 ta
	"BELARUSIAN",               // 47 be
	"JAVANESE",                 // 48 jw
	"OCCITAN",                  // 49 oc
	"URDU",                     // 50 ur
	"BIHARI",                   // 51 bh
	"GUJARATI",                 // 52 gu
	"THAI",                     // 53 th
	"ARABIC",                   // 54 ar
	"CATALAN",                  // 55 ca
	"ESPERANTO",                // 56 eo
	"BASQUE",                   // 57 eu
	"INTERLINGUA",              // 58 ia
	"KANNADA",                  // 59 kn
	"PUNJABI",                  // 60 pa
	"SCOTS_GAELIC",             // 61 gd
	"SWAHILI",                  // 62 sw
	"SLOVENIAN",                // 63 sl
	"MARATHI",                  // 64 mr
	"MALTESE",                  // 65 mt
	"VIETNAMESE",               // 66 vi
	"FRISIAN",                  // 67 fy
	"SLOVAK",                   // 68 sk
	"CHINESE_T",                // 69 zh-Hant
	"FAROESE",                  // 70 fo
	"SUNDANESE",                // 71 su
	"UZBEK",                    // 72 uz
	"AMHARIC",                  // 73 am
	"AZERBAIJANI",              // 74 az
	"GEORGIAN",                 // 75 ka
	"TIGRINYA",                 // 76 ti
	"PERSIAN",                  // 77 fa
	"BOSNIAN",                  // 78 bs
	"SINHALESE",                // 79 si
	"NORWEGIAN_N",              // 80 nn
	"X_81",                     // 81
	"X_82",                     // 82
	"XHOSA",                    // 83 xh
	"ZULU",                     // 84 zu
	"GUARANI",                  // 85 gn
	"SESOTHO",                  // 86 st
	"TURKMEN",                  // 87 tk
	"KYRGYZ",                   // 88 ky
	"BRETON",                   // 89 br
	"TWI",                      // 90 tw
	"YIDDISH",                  // 91 yi
	"X_92",                     // 92
	"SOMALI",                   // 93 so
	"UIGHUR",                   // 94 ug
	"KURDISH",                  // 95 ku
	"MONGOLIAN",                // 96 mn
	"ARMENIAN",                 // 97 hy
	"LAOTHIAN",                 // 98 lo
	"SINDHI",                   // 99 sd
	"RHAETO_ROMANCE",           // 100 rm
	"AFRIKAANS",                // 101 af
	"LUXEMBOURGISH",            // 102 lb
	"BURMESE",                  // 103 my
	"KHMER",                    // 104 km
	"TIBETAN",                  // 105 bo
	"DHIVEHI",                  // 106 dv
	"CHEROKEE",                 // 107 chr
	"SYRIAC",                   // 108 syr
	"LIMBU",                    // 109 lif
	"ORIYA",                    // 110 or
	"ASSAMESE",                 // 111 as
	"CORSICAN",                 // 112 co
	"INTERLINGUE",              // 113 ie
	"KAZAKH",                   // 114 kk
	"LINGALA",                  // 115 ln
	"X_116",                    // 116
	"PASHTO",                   // 117 ps
	"QUECHUA",                  // 118 qu
	"SHONA",                    // 119 sn
	"TAJIK",                    // 120 tg
	"TATAR",                    // 121 tt
	"TONGA",                    // 122 to
	"YORUBA",                   // 123 yo
	"X_124",                    // 124
	"X_125",                    // 125
	"X_126",                    // 126
	"X_127",                    // 127
	"MAORI",                    // 128 mi
	"WOLOF",                    // 129 wo
	"ABKHAZIAN",                // 130 ab
	"AFAR",                     // 131 aa
	"AYMARA",                   // 132 ay
	"BASHKIR",                  // 133 ba
	"BISLAMA",                  // 134 bi
	"DZONGKHA",                 // 135 dz
	"FIJIAN",                   // 136 fj
	"GREENLANDIC",              // 137 kl
	"HAUSA",                    // 138 ha
	"HAITIAN_CREOLE",           // 139 ht
	"INUPIAK",                  // 140 ik
	"INUKTITUT",                // 141 iu
	"KASHMIRI",                 // 142 ks
	"KINYARWANDA",              // 143 rw
	"MALAGASY",                 // 144 mg
	"NAURU",                    // 145 na
	"OROMO",                    // 146 om
	"RUNDI",                    // 147 rn
	"SAMOAN",                   // 148 sm
	"SANGO",                    // 149 sg
	"SANSKRIT",                 // 150 sa
	"SISWANT",                  // 151 ss
	"TSONGA",                   // 152 ts
	"TSWANA",                   // 153 tn
	"VOLAPUK",                  // 154 vo
	"ZHUANG",                   // 155 za
	"KHASI",                    // 156 kha
	"SCOTS",                    // 157 sco
	"GANDA",                    // 158 lg
	"MANX",                     // 159 gv
	"MONTENEGRIN",              // 160 sr-ME
	"AKAN",                     // 161 ak
	"IGBO",                     // 162 ig
	"MAURITIAN_CREOLE",         // 163 mfe
	"HAWAIIAN",                 // 164 haw
	"CEBUANO",                  // 165 ceb
	"EWE",                      // 166 ee
	"GA",                       // 167 gaa
	"HMONG",                    // 168 hmn
	"KRIO",                     // 169 kri
	"LOZI",                     // 170 loz
	"LUBA_LULUA",               // 171 lua
	"LUO_KENYA_AND_TANZANIA",   // 172 luo
	"NEWARI",                   // 173 new
	"NYANJA",                   // 174 ny
	"OSSETIAN",                 // 175 os
	"PAMPANGA",                 // 176 pam
	"PEDI",                     // 177 nso
	"RAJASTHANI",               // 178 raj
	"SESELWA",                  // 179 crs
	"TUMBUKA",                  // 180 tum
	"VENDA",                    // 181 ve
	"WARAY_PHILIPPINES",        // 182 war
	"X_183",                    // 183
	"X_184",                    // 184
	"X_185",                    // 185
	"X_186",                    // 186
	"X_187",                    // 187
	"X_188",                    // 188
	"X_189",                    // 189
	"X_190",                    // 190
	"X_191",                    // 191
	"X_192",                    // 192
	"X_193",                    // 193
	"X_194",                    // 194
	"X_195",                    // 195
	"X_196",                    // 196
	"X_197",                    // 197
	"X_198",                    // 198
	"X_199",                    // 199
	"X_200",                    // 200
	"X_201",                    // 201
	"X_202",                    // 202
	"X_203",                    // 203
	"X_204",                    // 204
	"X_205",                    // 205
	"X_206",                    // 206
	"X_207",                    // 207
	"X_208",                    // 208
	"X_209",                    // 209
	"X_210",                    // 210
	"X_211",                    // 211
	"X_212",                    // 212
	"X_213",                    // 213
	"X_214",                    // 214
	"X_215",                    // 215
	"X_216",                    // 216
	"X_217",                    // 217
	"X_218",                    // 218
	"X_219",                    // 219
	"X_220",                    // 220
	"X_221",                    // 221
	"X_222",                    // 222
	"X_223",                    // 223
	"X_224",                    // 224
	"X_225",                    // 225
	"X_226",                    // 226
	"X_227",                    // 227
	"X_228",                    // 228
	"X_229",                    // 229
	"X_230",                    // 230
	"X_231",                    // 231
	"X_232",                    // 232
	"X_233",                    // 233
	"X_234",                    // 234
	"X_235",                    // 235
	"X_236",                    // 236
	"X_237",                    // 237
	"X_238",                    // 238
	"X_239",                    // 239
	"X_240",                    // 240
	"X_241",                    // 241
	"X_242",                    // 242
	"X_243",                    // 243
	"X_244",                    // 244
	"X_245",                    // 245
	"X_246",                    // 246
	"X_247",                    // 247
	"X_248",                    // 248
	"X_249",                    // 249
	"X_250",                    // 250
	"X_251",                    // 251
	"X_252",                    // 252
	"X_253",                    // 253
	"X_254",                    // 254
	"X_255",                    // 255
	"X_256",                    // 256
	"X_257",                    // 257
	"X_258",                    // 258
	"X_259",                    // 259
	"X_260",                    // 260
	"X_261",                    // 261
	"X_262",                    // 262
	"X_263",                    // 263
	"X_264",                    // 264
	"X_265",                    // 265
	"X_266",                    // 266
	"X_267",                    // 267
	"X_268",                    // 268
	"X_269",                    // 269
	"X_270",                    // 270
	"X_271",                    // 271
	"X_272",                    // 272
	"X_273",                    // 273
	"X_274",                    // 274
	"X_275",                    // 275
	"X_276",                    // 276
	"X_277",                    // 277
	"X_278",                    // 278
	"X_279",                    // 279
	"X_280",                    // 280
	"X_281",                    // 281
	"X_282",                    // 282
	"X_283",                    // 283
	"X_284",                    // 284
	"X_285",                    // 285
	"X_286",                    // 286
	"X_287",                    // 287
	"X_288",                    // 288
	"X_289",                    // 289
	"X_290",                    // 290
	"X_291",                    // 291
	"X_292",                    // 292
	"X_293",                    // 293
	"X_294",                    // 294
	"X_295",                    // 295
	"X_296",                    // 296
	"X_297",                    // 297
	"X_298",                    // 298
	"X_299",                    // 299
	"X_300",                    // 300
	"X_301",                    // 301
	"X_302",                    // 302
	"X_303",                    // 303
	"X_304",                    // 304
	"X_305",                    // 305
	"X_306",                    // 306
	"X_307",                    // 307
	"X_308",                    // 308
	"X_309",                    // 309
	"X_310",                    // 310
	"X_311",                    // 311
	"X_312",                    // 312
	"X_313",                    // 313
	"X_314",                    // 314
	"X_315",                    // 315
	"X_316",                    // 316
	"X_317",                    // 317
	"X_318",                    // 318
	"X_319",                    // 319
	"X_320",                    // 320
	"X_321",                    // 321
	"X_322",                    // 322
	"X_323",                    // 323
	"X_324",                    // 324
	"X_325",                    // 325
	"X_326",                    // 326
	"X_327",                    // 327
	"X_328",                    // 328
	"X_329",                    // 329
	"X_330",                    // 330
	"X_331",                    // 331
	"X_332",                    // 332
	"X_333",                    // 333
	"X_334",                    // 334
	"X_335",                    // 335
	"X_336",                    // 336
	"X_337",                    // 337
	"X_338",                    // 338
	"X_339",                    // 339
	"X_340",                    // 340
	"X_341",                    // 341
	"X_342",                    // 342
	"X_343",                    // 343
	"X_344",                    // 344
	"X_345",                    // 345
	"X_346",                    // 346
	"X_347",                    // 347
	"X_348",                    // 348
	"X_349",                    // 349
	"X_350",                    // 350
	"X_351",                    // 351
	"X_352",                    // 352
	"X_353",                    // 353
	"X_354",                    // 354
	"X_355",                    // 355
	"X_356",                    // 356
	"X_357",                    // 357
	"X_358",                    // 358
	"X_359",                    // 359
	"X_360",                    // 360
	"X_361",                    // 361
	"X_362",                    // 362
	"X_363",                    // 363
	"X_364",                    // 364
	"X_365",                    // 365
	"X_366",                    // 366
	"X_367",                    // 367
	"X_368",                    // 368
	"X_369",                    // 369
	"X_370",                    // 370
	"X_371",                    // 371
	"X_372",                    // 372
	"X_373",                    // 373
	"X_374",                    // 374
	"X_375",                    // 375
	"X_376",                    // 376
	"X_377",                    // 377
	"X_378",                    // 378
	"X_379",                    // 379
	"X_380",                    // 380
	"X_381",                    // 381
	"X_382",                    // 382
	"X_383",                    // 383
	"X_384",                    // 384
	"X_385",                    // 385
	"X_386",                    // 386
	"X_387",                    // 387
	"X_388",                    // 388
	"X_389",                    // 389
	"X_390",                    // 390
	"X_391",                    // 391
	"X_392",                    // 392
	"X_393",                    // 393
	"X_394",                    // 394
	"X_395",                    // 395
	"X_396",                    // 396
	"X_397",                    // 397
	"X_398",                    // 398
	"X_399",                    // 399
	"X_400",                    // 400
	"X_401",                    // 401
	"X_402",                    // 402
	"X_403",                    // 403
	"X_404",                    // 404
	"X_405",                    // 405
	"X_406",                    // 406
	"X_407",                    // 407
	"X_408",                    // 408
	"X_409",                    // 409
	"X_410",                    // 410
	"X_411",                    // 411
	"X_412",                    // 412
	"X_413",                    // 413
	"X_414",                    // 414
	"X_415",                    // 415
	"X_416",                    // 416
	"X_417",                    // 417
	"X_418",                    // 418
	"X_419",                    // 419
	"X_420",                    // 420
	"X_421",                    // 421
	"X_422",                    // 422
	"X_423",                    // 423
	"X_424",                    // 424
	"X_425",                    // 425
	"X_426",                    // 426
	"X_427",                    // 427
	"X_428",                    // 428
	"X_429",                    // 429
	"X_430",                    // 430
	"X_431",                    // 431
	"X_432",                    // 432
	"X_433",                    // 433
	"X_434",                    // 434
	"X_435",                    // 435
	"X_436",                    // 436
	"X_437",                    // 437
	"X_438",                    // 438
	"X_439",                    // 439
	"X_440",                    // 440
	"X_441",                    // 441
	"X_442",                    // 442
	"X_443",                    // 443
	"X_444",                    // 444
	"X_445",                    // 445
	"X_446",                    // 446
	"X_447",                    // 447
	"X_448",                    // 448
	"X_449",                    // 449
	"X_450",                    // 450
	"X_451",                    // 451
	"X_452",                    // 452
	"X_453",                    // 453
	"X_454",                    // 454
	"X_455",                    // 455
	"X_456",                    // 456
	"X_457",                    // 457
	"X_458",                    // 458
	"X_459",                    // 459
	"X_460",                    // 460
	"X_461",                    // 461
	"X_462",                    // 462
	"X_463",                    // 463
	"X_464",                    // 464
	"X_465",                    // 465
	"X_466",                    // 466
	"X_467",                    // 467
	"X_468",                    // 468
	"X_469",                    // 469
	"X_470",                    // 470
	"X_471",                    // 471
	"X_472",                    // 472
	"X_473",                    // 473
	"X_474",                    // 474
	"X_475",                    // 475
	"X_476",                    // 476
	"X_477",                    // 477
	"X_478",                    // 478
	"X_479",                    // 479
	"X_480",                    // 480
	"X_481",                    // 481
	"X_482",                    // 482
	"X_483",                    // 483
	"X_484",                    // 484
	"X_485",                    // 485
	"X_486",                    // 486
	"X_487",                    // 487
	"X_488",                    // 488
	"X_489",                    // 489
	"X_490",                    // 490
	"X_491",                    // 491
	"X_492",                    // 492
	"X_493",                    // 493
	"X_494",                    // 494
	"X_495",                    // 495
	"X_496",                    // 496
	"X_497",                    // 497
	"X_498",                    // 498
	"X_499",                    // 499
	"X_500",                    // 500
	"X_501",                    // 501
	"X_502",                    // 502
	"X_503",                    // 503
	"X_504",                    // 504
	"X_505",                    // 505
	"NDEBELE",                  // 506 nr
	"X_BORK_BORK_BORK",         // 507 zzb
	"X_PIG_LATIN",              // 508 zzp
	"X_HACKER",                 // 509 zzh
	"X_KLINGON",                // 510 tlh
	"X_ELMER_FUDD",             // 511 zze
	"X_Common",                 // 512 xx-Zyyy
	"X_Latin",                  // 513 xx-Latn
	"X_Greek",                  // 514 xx-Grek
	"X_Cyrillic",               // 515 xx-Cyrl
	"X_Armenian",               // 516 xx-Armn
	"X_Hebrew",                 // 517 xx-Hebr
	"X_Arabic",                 // 518 xx-Arab
	"X_Syriac",                 // 519 xx-Syrc
	"X_Thaana",                 // 520 xx-Thaa
	"X_Devanagari",             // 521 xx-Deva
	"X_Bengali",                // 522 xx-Beng
	"X_Gurmukhi",               // 523 xx-Guru
	"X_Gujarati",               // 524 xx-Gujr
	"X_Oriya",                  // 525 xx-Orya
	"X_Tamil",                  // 526 xx-Taml
	"X_Telugu",                 // 527 xx-Telu
	"X_Kannada",                // 528 xx-Knda
	"X_Malayalam",              // 529 xx-Mlym
	"X_Sinhala",                // 530 xx-Sinh
	"X_Thai",                   // 531 xx-Thai
	"X_Lao",                    // 532 xx-Laoo
	"X_Tibetan",                // 533 xx-Tibt
	"X_Myanmar",                // 534 xx-Mymr
	"X_Georgian",               // 535 xx-Geor
	"X_Hangul",                 // 536 xx-Hang
	"X_Ethiopic",               // 537 xx-Ethi
	"X_Cherokee",               // 538 xx-Cher
	"X_Canadian_Aboriginal",    // 539 xx-Cans
	"X_Ogham",                  // 540 xx-Ogam
	"X_Runic",                  // 541 xx-Runr
	"X_Khmer",                  // 542 xx-Khmr
	"X_Mongolian",              // 543 xx-Mong
	"X_Hiragana",               // 544 xx-Hira
	"X_Katakana",               // 545 xx-Kana
	"X_Bopomofo",               // 546 xx-Bopo
	"X_Han",                    // 547 xx-Hani
	"X_Yi",                     // 548 xx-Yiii
	"X_Old_Italic",             // 549 xx-Ital
	"X_Gothic",                 // 550 xx-Goth
	"X_Deseret",                // 551 xx-Dsrt
	"X_Inherited",              // 552 xx-Qaai
	"X_Tagalog",                // 553 xx-Tglg
	"X_Hanunoo",                // 554 xx-Hano
	"X_Buhid",                  // 555 xx-Buhd
	"X_Tagbanwa",               // 556 xx-Tagb
	"X_Limbu",                  // 557 xx-Limb
	"X_Tai_Le",                 // 558 xx-Tale
	"X_Linear_B",               // 559 xx-Linb
	"X_Ugaritic",               // 560 xx-Ugar
	"X_Shavian",                // 561 xx-Shaw
	"X_Osmanya",                // 562 xx-Osma
	"X_Cypriot",                // 563 xx-Cprt
	"X_Braille",                // 564 xx-Brai
	"X_Buginese",               // 565 xx-Bugi
	"X_Coptic",                 // 566 xx-Copt
	"X_New_Tai_Lue",            // 567 xx-Talu
	"X_Glagolitic",             // 568 xx-Glag
	"X_Tifinagh",               // 569 xx-Tfng
	"X_Syloti_Nagri",           // 570 xx-Sylo
	"X_Old_Persian",            // 571 xx-Xpeo
	"X_Kharoshthi",             // 572 xx-Khar
	"X_Balinese",               // 573 xx-Bali
	"X_Cuneiform",              // 574 xx-Xsux
	"X_Phoenician",             // 575 xx-Phnx
	"X_Phags_Pa",               // 576 xx-Phag
	"X_Nko",                    // 577 xx-Nkoo
	"X_Sundanese",              // 578 xx-Sund
	"X_Lepcha",                 // 579 xx-Lepc
	"X_Ol_Chiki",               // 580 xx-Olck
	"X_Vai",                    // 581 xx-Vaii
	"X_Saurashtra",             // 582 xx-Saur
	"X_Kayah_Li",               // 583 xx-Kali
	"X_Rejang",                 // 584 xx-Rjng
	"X_Lycian",                 // 585 xx-Lyci
	"X_Carian",                 // 586 xx-Cari
	"X_Lydian",                 // 587 xx-Lydi
	"X_Cham",                   // 588 xx-Cham
	"X_Tai_Tham",               // 589 xx-Lana
	"X_Tai_Viet",               // 590 xx-Tavt
	"X_Avestan",                // 591 xx-Avst
	"X_Egyptian_Hieroglyphs",   // 592 xx-Egyp
	"X_Samaritan",              // 593 xx-Samr
	"X_Lisu",                   // 594 xx-Lisu
	"X_Bamum",                  // 595 xx-Bamu
	"X_Javanese",               // 596 xx-Java
	"X_Meetei_Mayek",           // 597 xx-Mtei
	"X_Imperial_Aramaic",       // 598 xx-Armi
	"X_Old_South_Arabian",      // 599 xx-Sarb
	"X_Inscriptional_Parthian", // 600 xx-Prti
	"X_Inscriptional_Pahlavi",  // 601 xx-Phli
	"X_Old_Turkic",             // 602 xx-Orkh
	"X_Kaithi",                 // 603 xx-Kthi
	"X_Batak",                  // 604 xx-Batk
	"X_Brahmi",                 // 605 xx-Brah
	"X_Mandaic",                // 606 xx-Mand
	"X_Chakma",                 // 607 xx-Cakm
	"X_Meroitic_Cursive",       // 608 xx-Merc
	"X_Meroitic_Hieroglyphs",   // 609 xx-Mero
	"X_Miao",                   // 610 xx-Plrd
	"X_Sharada",                // 611 xx-Shrd
	"X_Sora_Sompeng",           // 612 xx-Sora
	"X_Takri",                  // 613 xx-Takr
}
