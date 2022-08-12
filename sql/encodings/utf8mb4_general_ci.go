// Copyright 2022 Dolthub, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package encodings

// Utf8mb4_general_ci_RuneWeight returns the weight of a given rune based on its relational sort order from
// the `utf8mb4_general_ci` collation.
func Utf8mb4_general_ci_RuneWeight(r rune) int32 {
	weight, ok := utf8mb4_general_ci_Weights[r]
	if ok {
		return weight
	} else if r >= 659 && r <= 836 {
		return r - 342
	} else if r >= 1415 && r <= 7679 {
		return r - 586
	} else if r >= 8189 && r <= 8544 {
		return r - 1040
	} else if r >= 8575 && r <= 9398 {
		return r - 1056
	} else if r >= 9449 && r <= 55295 {
		return r - 1082
	} else if r >= 57345 && r <= 65313 {
		return r - 3130
	} else if r >= 65371 && r <= 65533 {
		return r - 3156
	} else if r >= 65536 && r <= 1114111 {
		return 62377
	} else {
		return 2147483647
	}
}

// utf8mb4_general_ci_Weights contain a map from rune to weight for the `utf8mb4_general_ci` collation. The
// map primarily contains mappings that have a random order. Mappings that fit into a sequential range (and are long
// enough) are defined in the calling function to save space.
var utf8mb4_general_ci_Weights = map[rune]int32{
	0:     0,
	1:     1,
	2:     2,
	3:     3,
	4:     4,
	5:     5,
	6:     6,
	7:     7,
	8:     8,
	9:     9,
	10:    10,
	11:    11,
	12:    12,
	13:    13,
	14:    14,
	15:    15,
	16:    16,
	17:    17,
	18:    18,
	19:    19,
	20:    20,
	21:    21,
	22:    22,
	23:    23,
	24:    24,
	25:    25,
	26:    26,
	27:    27,
	28:    28,
	29:    29,
	30:    30,
	31:    31,
	32:    32,
	33:    33,
	34:    34,
	35:    35,
	36:    36,
	37:    37,
	38:    38,
	39:    39,
	40:    40,
	41:    41,
	42:    42,
	43:    43,
	44:    44,
	45:    45,
	46:    46,
	47:    47,
	48:    48,
	49:    49,
	50:    50,
	51:    51,
	52:    52,
	53:    53,
	54:    54,
	55:    55,
	56:    56,
	57:    57,
	58:    58,
	59:    59,
	60:    60,
	61:    61,
	62:    62,
	63:    63,
	64:    64,
	65:    65,
	97:    65,
	192:   65,
	193:   65,
	194:   65,
	195:   65,
	196:   65,
	197:   65,
	224:   65,
	225:   65,
	226:   65,
	227:   65,
	228:   65,
	229:   65,
	256:   65,
	257:   65,
	258:   65,
	259:   65,
	260:   65,
	261:   65,
	461:   65,
	462:   65,
	478:   65,
	479:   65,
	480:   65,
	481:   65,
	506:   65,
	507:   65,
	512:   65,
	513:   65,
	514:   65,
	515:   65,
	550:   65,
	551:   65,
	7680:  65,
	7681:  65,
	7840:  65,
	7841:  65,
	7842:  65,
	7843:  65,
	7844:  65,
	7845:  65,
	7846:  65,
	7847:  65,
	7848:  65,
	7849:  65,
	7850:  65,
	7851:  65,
	7852:  65,
	7853:  65,
	7854:  65,
	7855:  65,
	7856:  65,
	7857:  65,
	7858:  65,
	7859:  65,
	7860:  65,
	7861:  65,
	7862:  65,
	7863:  65,
	66:    66,
	98:    66,
	7682:  66,
	7683:  66,
	7684:  66,
	7685:  66,
	7686:  66,
	7687:  66,
	67:    67,
	99:    67,
	199:   67,
	231:   67,
	262:   67,
	263:   67,
	264:   67,
	265:   67,
	266:   67,
	267:   67,
	268:   67,
	269:   67,
	7688:  67,
	7689:  67,
	68:    68,
	100:   68,
	270:   68,
	271:   68,
	7690:  68,
	7691:  68,
	7692:  68,
	7693:  68,
	7694:  68,
	7695:  68,
	7696:  68,
	7697:  68,
	7698:  68,
	7699:  68,
	69:    69,
	101:   69,
	200:   69,
	201:   69,
	202:   69,
	203:   69,
	232:   69,
	233:   69,
	234:   69,
	235:   69,
	274:   69,
	275:   69,
	276:   69,
	277:   69,
	278:   69,
	279:   69,
	280:   69,
	281:   69,
	282:   69,
	283:   69,
	516:   69,
	517:   69,
	518:   69,
	519:   69,
	552:   69,
	553:   69,
	7700:  69,
	7701:  69,
	7702:  69,
	7703:  69,
	7704:  69,
	7705:  69,
	7706:  69,
	7707:  69,
	7708:  69,
	7709:  69,
	7864:  69,
	7865:  69,
	7866:  69,
	7867:  69,
	7868:  69,
	7869:  69,
	7870:  69,
	7871:  69,
	7872:  69,
	7873:  69,
	7874:  69,
	7875:  69,
	7876:  69,
	7877:  69,
	7878:  69,
	7879:  69,
	70:    70,
	102:   70,
	7710:  70,
	7711:  70,
	71:    71,
	103:   71,
	284:   71,
	285:   71,
	286:   71,
	287:   71,
	288:   71,
	289:   71,
	290:   71,
	291:   71,
	486:   71,
	487:   71,
	500:   71,
	501:   71,
	7712:  71,
	7713:  71,
	72:    72,
	104:   72,
	292:   72,
	293:   72,
	542:   72,
	543:   72,
	7714:  72,
	7715:  72,
	7716:  72,
	7717:  72,
	7718:  72,
	7719:  72,
	7720:  72,
	7721:  72,
	7722:  72,
	7723:  72,
	7830:  72,
	73:    73,
	105:   73,
	204:   73,
	205:   73,
	206:   73,
	207:   73,
	236:   73,
	237:   73,
	238:   73,
	239:   73,
	296:   73,
	297:   73,
	298:   73,
	299:   73,
	300:   73,
	301:   73,
	302:   73,
	303:   73,
	304:   73,
	305:   73,
	463:   73,
	464:   73,
	520:   73,
	521:   73,
	522:   73,
	523:   73,
	7724:  73,
	7725:  73,
	7726:  73,
	7727:  73,
	7880:  73,
	7881:  73,
	7882:  73,
	7883:  73,
	74:    74,
	106:   74,
	308:   74,
	309:   74,
	496:   74,
	75:    75,
	107:   75,
	310:   75,
	311:   75,
	488:   75,
	489:   75,
	7728:  75,
	7729:  75,
	7730:  75,
	7731:  75,
	7732:  75,
	7733:  75,
	76:    76,
	108:   76,
	313:   76,
	314:   76,
	315:   76,
	316:   76,
	317:   76,
	318:   76,
	7734:  76,
	7735:  76,
	7736:  76,
	7737:  76,
	7738:  76,
	7739:  76,
	7740:  76,
	7741:  76,
	77:    77,
	109:   77,
	7742:  77,
	7743:  77,
	7744:  77,
	7745:  77,
	7746:  77,
	7747:  77,
	78:    78,
	110:   78,
	209:   78,
	241:   78,
	323:   78,
	324:   78,
	325:   78,
	326:   78,
	327:   78,
	328:   78,
	504:   78,
	505:   78,
	7748:  78,
	7749:  78,
	7750:  78,
	7751:  78,
	7752:  78,
	7753:  78,
	7754:  78,
	7755:  78,
	79:    79,
	111:   79,
	210:   79,
	211:   79,
	212:   79,
	213:   79,
	214:   79,
	242:   79,
	243:   79,
	244:   79,
	245:   79,
	246:   79,
	332:   79,
	333:   79,
	334:   79,
	335:   79,
	336:   79,
	337:   79,
	416:   79,
	417:   79,
	465:   79,
	466:   79,
	490:   79,
	491:   79,
	492:   79,
	493:   79,
	524:   79,
	525:   79,
	526:   79,
	527:   79,
	554:   79,
	555:   79,
	556:   79,
	557:   79,
	558:   79,
	559:   79,
	560:   79,
	561:   79,
	7756:  79,
	7757:  79,
	7758:  79,
	7759:  79,
	7760:  79,
	7761:  79,
	7762:  79,
	7763:  79,
	7884:  79,
	7885:  79,
	7886:  79,
	7887:  79,
	7888:  79,
	7889:  79,
	7890:  79,
	7891:  79,
	7892:  79,
	7893:  79,
	7894:  79,
	7895:  79,
	7896:  79,
	7897:  79,
	7898:  79,
	7899:  79,
	7900:  79,
	7901:  79,
	7902:  79,
	7903:  79,
	7904:  79,
	7905:  79,
	7906:  79,
	7907:  79,
	80:    80,
	112:   80,
	7764:  80,
	7765:  80,
	7766:  80,
	7767:  80,
	81:    81,
	113:   81,
	82:    82,
	114:   82,
	340:   82,
	341:   82,
	342:   82,
	343:   82,
	344:   82,
	345:   82,
	528:   82,
	529:   82,
	530:   82,
	531:   82,
	7768:  82,
	7769:  82,
	7770:  82,
	7771:  82,
	7772:  82,
	7773:  82,
	7774:  82,
	7775:  82,
	83:    83,
	115:   83,
	223:   83,
	346:   83,
	347:   83,
	348:   83,
	349:   83,
	350:   83,
	351:   83,
	352:   83,
	353:   83,
	383:   83,
	536:   83,
	537:   83,
	7776:  83,
	7777:  83,
	7778:  83,
	7779:  83,
	7780:  83,
	7781:  83,
	7782:  83,
	7783:  83,
	7784:  83,
	7785:  83,
	7835:  83,
	84:    84,
	116:   84,
	354:   84,
	355:   84,
	356:   84,
	357:   84,
	538:   84,
	539:   84,
	7786:  84,
	7787:  84,
	7788:  84,
	7789:  84,
	7790:  84,
	7791:  84,
	7792:  84,
	7793:  84,
	7831:  84,
	85:    85,
	117:   85,
	217:   85,
	218:   85,
	219:   85,
	220:   85,
	249:   85,
	250:   85,
	251:   85,
	252:   85,
	360:   85,
	361:   85,
	362:   85,
	363:   85,
	364:   85,
	365:   85,
	366:   85,
	367:   85,
	368:   85,
	369:   85,
	370:   85,
	371:   85,
	431:   85,
	432:   85,
	467:   85,
	468:   85,
	469:   85,
	470:   85,
	471:   85,
	472:   85,
	473:   85,
	474:   85,
	475:   85,
	476:   85,
	532:   85,
	533:   85,
	534:   85,
	535:   85,
	7794:  85,
	7795:  85,
	7796:  85,
	7797:  85,
	7798:  85,
	7799:  85,
	7800:  85,
	7801:  85,
	7802:  85,
	7803:  85,
	7908:  85,
	7909:  85,
	7910:  85,
	7911:  85,
	7912:  85,
	7913:  85,
	7914:  85,
	7915:  85,
	7916:  85,
	7917:  85,
	7918:  85,
	7919:  85,
	7920:  85,
	7921:  85,
	86:    86,
	118:   86,
	7804:  86,
	7805:  86,
	7806:  86,
	7807:  86,
	87:    87,
	119:   87,
	372:   87,
	373:   87,
	7808:  87,
	7809:  87,
	7810:  87,
	7811:  87,
	7812:  87,
	7813:  87,
	7814:  87,
	7815:  87,
	7816:  87,
	7817:  87,
	7832:  87,
	88:    88,
	120:   88,
	7818:  88,
	7819:  88,
	7820:  88,
	7821:  88,
	89:    89,
	121:   89,
	221:   89,
	253:   89,
	255:   89,
	374:   89,
	375:   89,
	376:   89,
	562:   89,
	563:   89,
	7822:  89,
	7823:  89,
	7833:  89,
	7922:  89,
	7923:  89,
	7924:  89,
	7925:  89,
	7926:  89,
	7927:  89,
	7928:  89,
	7929:  89,
	90:    90,
	122:   90,
	377:   90,
	378:   90,
	379:   90,
	380:   90,
	381:   90,
	382:   90,
	7824:  90,
	7825:  90,
	7826:  90,
	7827:  90,
	7828:  90,
	7829:  90,
	91:    91,
	92:    92,
	93:    93,
	94:    94,
	95:    95,
	96:    96,
	123:   97,
	124:   98,
	125:   99,
	126:   100,
	127:   101,
	128:   102,
	129:   103,
	130:   104,
	131:   105,
	132:   106,
	133:   107,
	134:   108,
	135:   109,
	136:   110,
	137:   111,
	138:   112,
	139:   113,
	140:   114,
	141:   115,
	142:   116,
	143:   117,
	144:   118,
	145:   119,
	146:   120,
	147:   121,
	148:   122,
	149:   123,
	150:   124,
	151:   125,
	152:   126,
	153:   127,
	154:   128,
	155:   129,
	156:   130,
	157:   131,
	158:   132,
	159:   133,
	160:   134,
	161:   135,
	162:   136,
	163:   137,
	164:   138,
	165:   139,
	166:   140,
	167:   141,
	168:   142,
	169:   143,
	170:   144,
	171:   145,
	172:   146,
	173:   147,
	174:   148,
	175:   149,
	176:   150,
	177:   151,
	178:   152,
	179:   153,
	180:   154,
	182:   155,
	183:   156,
	184:   157,
	185:   158,
	186:   159,
	187:   160,
	188:   161,
	189:   162,
	190:   163,
	191:   164,
	198:   165,
	230:   165,
	482:   165,
	483:   165,
	508:   165,
	509:   165,
	208:   166,
	240:   166,
	215:   167,
	216:   168,
	248:   168,
	510:   168,
	511:   168,
	222:   169,
	254:   169,
	247:   170,
	272:   171,
	273:   171,
	294:   172,
	295:   172,
	306:   173,
	307:   173,
	312:   174,
	319:   175,
	320:   175,
	321:   176,
	322:   176,
	329:   177,
	330:   178,
	331:   178,
	338:   179,
	339:   179,
	358:   180,
	359:   180,
	384:   181,
	385:   182,
	595:   182,
	386:   183,
	387:   183,
	388:   184,
	389:   184,
	390:   185,
	596:   185,
	391:   186,
	392:   186,
	393:   187,
	598:   187,
	394:   188,
	599:   188,
	395:   189,
	396:   189,
	397:   190,
	398:   191,
	477:   191,
	399:   192,
	601:   192,
	400:   193,
	603:   193,
	401:   194,
	402:   194,
	403:   195,
	608:   195,
	404:   196,
	611:   196,
	406:   197,
	617:   197,
	407:   198,
	616:   198,
	408:   199,
	409:   199,
	410:   200,
	411:   201,
	412:   202,
	623:   202,
	413:   203,
	626:   203,
	414:   204,
	415:   205,
	629:   205,
	418:   206,
	419:   206,
	420:   207,
	421:   207,
	422:   208,
	640:   208,
	423:   209,
	424:   209,
	425:   210,
	643:   210,
	426:   211,
	427:   212,
	428:   213,
	429:   213,
	430:   214,
	648:   214,
	433:   215,
	650:   215,
	434:   216,
	651:   216,
	435:   217,
	436:   217,
	437:   218,
	438:   218,
	439:   219,
	494:   219,
	495:   219,
	658:   219,
	440:   220,
	441:   220,
	442:   221,
	443:   222,
	444:   223,
	445:   223,
	446:   224,
	448:   225,
	449:   226,
	450:   227,
	451:   228,
	452:   229,
	453:   229,
	454:   229,
	455:   230,
	456:   230,
	457:   230,
	458:   231,
	459:   231,
	460:   231,
	484:   232,
	485:   232,
	497:   233,
	498:   233,
	499:   233,
	405:   234,
	502:   234,
	447:   235,
	503:   235,
	540:   236,
	541:   236,
	544:   237,
	545:   238,
	546:   239,
	547:   239,
	548:   240,
	549:   240,
	564:   241,
	565:   242,
	566:   243,
	567:   244,
	568:   245,
	569:   246,
	570:   247,
	571:   248,
	572:   249,
	573:   250,
	574:   251,
	575:   252,
	576:   253,
	577:   254,
	578:   255,
	579:   256,
	580:   257,
	581:   258,
	582:   259,
	583:   260,
	584:   261,
	585:   262,
	586:   263,
	587:   264,
	588:   265,
	589:   266,
	590:   267,
	591:   268,
	592:   269,
	593:   270,
	594:   271,
	597:   272,
	600:   273,
	602:   274,
	604:   275,
	605:   276,
	606:   277,
	607:   278,
	609:   279,
	610:   280,
	612:   281,
	613:   282,
	614:   283,
	615:   284,
	618:   285,
	619:   286,
	620:   287,
	621:   288,
	622:   289,
	624:   290,
	625:   291,
	627:   292,
	628:   293,
	630:   294,
	631:   295,
	632:   296,
	633:   297,
	634:   298,
	635:   299,
	636:   300,
	637:   301,
	638:   302,
	639:   303,
	641:   304,
	642:   305,
	644:   306,
	645:   307,
	646:   308,
	647:   309,
	649:   310,
	652:   311,
	653:   312,
	654:   313,
	655:   314,
	656:   315,
	657:   316,
	838:   495,
	839:   496,
	840:   497,
	841:   498,
	842:   499,
	843:   500,
	844:   501,
	845:   502,
	846:   503,
	847:   504,
	848:   505,
	849:   506,
	850:   507,
	851:   508,
	852:   509,
	853:   510,
	854:   511,
	855:   512,
	856:   513,
	857:   514,
	858:   515,
	859:   516,
	860:   517,
	861:   518,
	862:   519,
	863:   520,
	864:   521,
	865:   522,
	866:   523,
	867:   524,
	868:   525,
	869:   526,
	870:   527,
	871:   528,
	872:   529,
	873:   530,
	874:   531,
	875:   532,
	876:   533,
	877:   534,
	878:   535,
	879:   536,
	880:   537,
	881:   538,
	882:   539,
	883:   540,
	884:   541,
	885:   542,
	886:   543,
	887:   544,
	888:   545,
	889:   546,
	890:   547,
	891:   548,
	892:   549,
	893:   550,
	894:   551,
	895:   552,
	896:   553,
	897:   554,
	898:   555,
	899:   556,
	900:   557,
	901:   558,
	903:   559,
	907:   560,
	909:   561,
	902:   562,
	913:   562,
	940:   562,
	945:   562,
	7936:  562,
	7937:  562,
	7938:  562,
	7939:  562,
	7940:  562,
	7941:  562,
	7942:  562,
	7943:  562,
	7944:  562,
	7945:  562,
	7946:  562,
	7947:  562,
	7948:  562,
	7949:  562,
	7950:  562,
	7951:  562,
	8048:  562,
	8064:  562,
	8065:  562,
	8066:  562,
	8067:  562,
	8068:  562,
	8069:  562,
	8070:  562,
	8071:  562,
	8072:  562,
	8073:  562,
	8074:  562,
	8075:  562,
	8076:  562,
	8077:  562,
	8078:  562,
	8079:  562,
	8112:  562,
	8113:  562,
	8114:  562,
	8115:  562,
	8116:  562,
	8118:  562,
	8119:  562,
	8120:  562,
	8121:  562,
	8122:  562,
	8124:  562,
	914:   563,
	946:   563,
	976:   563,
	915:   564,
	947:   564,
	916:   565,
	948:   565,
	904:   566,
	917:   566,
	941:   566,
	949:   566,
	7952:  566,
	7953:  566,
	7954:  566,
	7955:  566,
	7956:  566,
	7957:  566,
	7960:  566,
	7961:  566,
	7962:  566,
	7963:  566,
	7964:  566,
	7965:  566,
	8050:  566,
	8136:  566,
	918:   567,
	950:   567,
	905:   568,
	919:   568,
	942:   568,
	951:   568,
	7968:  568,
	7969:  568,
	7970:  568,
	7971:  568,
	7972:  568,
	7973:  568,
	7974:  568,
	7975:  568,
	7976:  568,
	7977:  568,
	7978:  568,
	7979:  568,
	7980:  568,
	7981:  568,
	7982:  568,
	7983:  568,
	8052:  568,
	8080:  568,
	8081:  568,
	8082:  568,
	8083:  568,
	8084:  568,
	8085:  568,
	8086:  568,
	8087:  568,
	8088:  568,
	8089:  568,
	8090:  568,
	8091:  568,
	8092:  568,
	8093:  568,
	8094:  568,
	8095:  568,
	8130:  568,
	8131:  568,
	8132:  568,
	8134:  568,
	8135:  568,
	8138:  568,
	8140:  568,
	920:   569,
	952:   569,
	977:   569,
	837:   570,
	906:   570,
	912:   570,
	921:   570,
	938:   570,
	943:   570,
	953:   570,
	970:   570,
	7984:  570,
	7985:  570,
	7986:  570,
	7987:  570,
	7988:  570,
	7989:  570,
	7990:  570,
	7991:  570,
	7992:  570,
	7993:  570,
	7994:  570,
	7995:  570,
	7996:  570,
	7997:  570,
	7998:  570,
	7999:  570,
	8054:  570,
	8126:  570,
	8144:  570,
	8145:  570,
	8146:  570,
	8150:  570,
	8151:  570,
	8152:  570,
	8153:  570,
	8154:  570,
	922:   571,
	954:   571,
	1008:  571,
	923:   572,
	955:   572,
	181:   573,
	924:   573,
	956:   573,
	925:   574,
	957:   574,
	926:   575,
	958:   575,
	908:   576,
	927:   576,
	959:   576,
	972:   576,
	8000:  576,
	8001:  576,
	8002:  576,
	8003:  576,
	8004:  576,
	8005:  576,
	8008:  576,
	8009:  576,
	8010:  576,
	8011:  576,
	8012:  576,
	8013:  576,
	8056:  576,
	8184:  576,
	928:   577,
	960:   577,
	982:   577,
	929:   578,
	961:   578,
	1009:  578,
	8164:  578,
	8165:  578,
	8172:  578,
	930:   579,
	931:   580,
	962:   580,
	963:   580,
	1010:  580,
	932:   581,
	964:   581,
	910:   582,
	933:   582,
	939:   582,
	944:   582,
	965:   582,
	971:   582,
	973:   582,
	8016:  582,
	8017:  582,
	8018:  582,
	8019:  582,
	8020:  582,
	8021:  582,
	8022:  582,
	8023:  582,
	8025:  582,
	8027:  582,
	8029:  582,
	8031:  582,
	8058:  582,
	8160:  582,
	8161:  582,
	8162:  582,
	8166:  582,
	8167:  582,
	8168:  582,
	8169:  582,
	8170:  582,
	934:   583,
	966:   583,
	981:   583,
	935:   584,
	967:   584,
	936:   585,
	968:   585,
	911:   586,
	937:   586,
	969:   586,
	974:   586,
	8032:  586,
	8033:  586,
	8034:  586,
	8035:  586,
	8036:  586,
	8037:  586,
	8038:  586,
	8039:  586,
	8040:  586,
	8041:  586,
	8042:  586,
	8043:  586,
	8044:  586,
	8045:  586,
	8046:  586,
	8047:  586,
	8060:  586,
	8096:  586,
	8097:  586,
	8098:  586,
	8099:  586,
	8100:  586,
	8101:  586,
	8102:  586,
	8103:  586,
	8104:  586,
	8105:  586,
	8106:  586,
	8107:  586,
	8108:  586,
	8109:  586,
	8110:  586,
	8111:  586,
	8178:  586,
	8179:  586,
	8180:  586,
	8182:  586,
	8183:  586,
	8186:  586,
	8188:  586,
	975:   587,
	978:   588,
	979:   588,
	980:   588,
	983:   589,
	984:   590,
	985:   591,
	986:   592,
	987:   592,
	988:   593,
	989:   593,
	990:   594,
	991:   594,
	992:   595,
	993:   595,
	994:   596,
	995:   596,
	996:   597,
	997:   597,
	998:   598,
	999:   598,
	1000:  599,
	1001:  599,
	1002:  600,
	1003:  600,
	1004:  601,
	1005:  601,
	1006:  602,
	1007:  602,
	1011:  603,
	1012:  604,
	1013:  605,
	1014:  606,
	1015:  607,
	1016:  608,
	1017:  609,
	1018:  610,
	1019:  611,
	1020:  612,
	1021:  613,
	1022:  614,
	1023:  615,
	1026:  616,
	1106:  616,
	1028:  617,
	1108:  617,
	1029:  618,
	1109:  618,
	1030:  619,
	1031:  619,
	1110:  619,
	1111:  619,
	1032:  620,
	1112:  620,
	1033:  621,
	1113:  621,
	1034:  622,
	1114:  622,
	1035:  623,
	1115:  623,
	1039:  624,
	1119:  624,
	1040:  625,
	1072:  625,
	1232:  625,
	1233:  625,
	1234:  625,
	1235:  625,
	1041:  626,
	1073:  626,
	1042:  627,
	1074:  627,
	1027:  628,
	1043:  628,
	1075:  628,
	1107:  628,
	1044:  629,
	1076:  629,
	1024:  630,
	1025:  630,
	1045:  630,
	1077:  630,
	1104:  630,
	1105:  630,
	1238:  630,
	1239:  630,
	1046:  631,
	1078:  631,
	1217:  631,
	1218:  631,
	1244:  631,
	1245:  631,
	1047:  632,
	1079:  632,
	1246:  632,
	1247:  632,
	1037:  633,
	1048:  633,
	1080:  633,
	1117:  633,
	1250:  633,
	1251:  633,
	1252:  633,
	1253:  633,
	1049:  634,
	1081:  634,
	1036:  635,
	1050:  635,
	1082:  635,
	1116:  635,
	1051:  636,
	1083:  636,
	1052:  637,
	1084:  637,
	1053:  638,
	1085:  638,
	1054:  639,
	1086:  639,
	1254:  639,
	1255:  639,
	1055:  640,
	1087:  640,
	1056:  641,
	1088:  641,
	1057:  642,
	1089:  642,
	1058:  643,
	1090:  643,
	1038:  644,
	1059:  644,
	1091:  644,
	1118:  644,
	1262:  644,
	1263:  644,
	1264:  644,
	1265:  644,
	1266:  644,
	1267:  644,
	1060:  645,
	1092:  645,
	1061:  646,
	1093:  646,
	1062:  647,
	1094:  647,
	1063:  648,
	1095:  648,
	1268:  648,
	1269:  648,
	1064:  649,
	1096:  649,
	1065:  650,
	1097:  650,
	1066:  651,
	1098:  651,
	1067:  652,
	1099:  652,
	1272:  652,
	1273:  652,
	1068:  653,
	1100:  653,
	1069:  654,
	1101:  654,
	1260:  654,
	1261:  654,
	1070:  655,
	1102:  655,
	1071:  656,
	1103:  656,
	1120:  657,
	1121:  657,
	1122:  658,
	1123:  658,
	1124:  659,
	1125:  659,
	1126:  660,
	1127:  660,
	1128:  661,
	1129:  661,
	1130:  662,
	1131:  662,
	1132:  663,
	1133:  663,
	1134:  664,
	1135:  664,
	1136:  665,
	1137:  665,
	1138:  666,
	1139:  666,
	1140:  667,
	1141:  667,
	1142:  667,
	1143:  667,
	1144:  668,
	1145:  668,
	1146:  669,
	1147:  669,
	1148:  670,
	1149:  670,
	1150:  671,
	1151:  671,
	1152:  672,
	1153:  672,
	1154:  673,
	1155:  674,
	1156:  675,
	1157:  676,
	1158:  677,
	1159:  678,
	1160:  679,
	1161:  680,
	1162:  681,
	1163:  682,
	1164:  683,
	1165:  683,
	1166:  684,
	1167:  684,
	1168:  685,
	1169:  685,
	1170:  686,
	1171:  686,
	1172:  687,
	1173:  687,
	1174:  688,
	1175:  688,
	1176:  689,
	1177:  689,
	1178:  690,
	1179:  690,
	1180:  691,
	1181:  691,
	1182:  692,
	1183:  692,
	1184:  693,
	1185:  693,
	1186:  694,
	1187:  694,
	1188:  695,
	1189:  695,
	1190:  696,
	1191:  696,
	1192:  697,
	1193:  697,
	1194:  698,
	1195:  698,
	1196:  699,
	1197:  699,
	1198:  700,
	1199:  700,
	1200:  701,
	1201:  701,
	1202:  702,
	1203:  702,
	1204:  703,
	1205:  703,
	1206:  704,
	1207:  704,
	1208:  705,
	1209:  705,
	1210:  706,
	1211:  706,
	1212:  707,
	1213:  707,
	1214:  708,
	1215:  708,
	1216:  709,
	1219:  710,
	1220:  710,
	1221:  711,
	1222:  712,
	1223:  713,
	1224:  713,
	1225:  714,
	1226:  715,
	1227:  716,
	1228:  716,
	1229:  717,
	1230:  718,
	1231:  719,
	1236:  720,
	1237:  720,
	1240:  721,
	1241:  721,
	1242:  721,
	1243:  721,
	1248:  722,
	1249:  722,
	1256:  723,
	1257:  723,
	1258:  723,
	1259:  723,
	1270:  724,
	1271:  725,
	1274:  726,
	1275:  727,
	1276:  728,
	1277:  729,
	1278:  730,
	1279:  731,
	1280:  732,
	1281:  733,
	1282:  734,
	1283:  735,
	1284:  736,
	1285:  737,
	1286:  738,
	1287:  739,
	1288:  740,
	1289:  741,
	1290:  742,
	1291:  743,
	1292:  744,
	1293:  745,
	1294:  746,
	1295:  747,
	1296:  748,
	1297:  749,
	1298:  750,
	1299:  751,
	1300:  752,
	1301:  753,
	1302:  754,
	1303:  755,
	1304:  756,
	1305:  757,
	1306:  758,
	1307:  759,
	1308:  760,
	1309:  761,
	1310:  762,
	1311:  763,
	1312:  764,
	1313:  765,
	1314:  766,
	1315:  767,
	1316:  768,
	1317:  769,
	1318:  770,
	1319:  771,
	1320:  772,
	1321:  773,
	1322:  774,
	1323:  775,
	1324:  776,
	1325:  777,
	1326:  778,
	1327:  779,
	1328:  780,
	1329:  781,
	1377:  781,
	1330:  782,
	1378:  782,
	1331:  783,
	1379:  783,
	1332:  784,
	1380:  784,
	1333:  785,
	1381:  785,
	1334:  786,
	1382:  786,
	1335:  787,
	1383:  787,
	1336:  788,
	1384:  788,
	1337:  789,
	1385:  789,
	1338:  790,
	1386:  790,
	1339:  791,
	1387:  791,
	1340:  792,
	1388:  792,
	1341:  793,
	1389:  793,
	1342:  794,
	1390:  794,
	1343:  795,
	1391:  795,
	1344:  796,
	1392:  796,
	1345:  797,
	1393:  797,
	1346:  798,
	1394:  798,
	1347:  799,
	1395:  799,
	1348:  800,
	1396:  800,
	1349:  801,
	1397:  801,
	1350:  802,
	1398:  802,
	1351:  803,
	1399:  803,
	1352:  804,
	1400:  804,
	1353:  805,
	1401:  805,
	1354:  806,
	1402:  806,
	1355:  807,
	1403:  807,
	1356:  808,
	1404:  808,
	1357:  809,
	1405:  809,
	1358:  810,
	1406:  810,
	1359:  811,
	1407:  811,
	1360:  812,
	1408:  812,
	1361:  813,
	1409:  813,
	1362:  814,
	1410:  814,
	1363:  815,
	1411:  815,
	1364:  816,
	1412:  816,
	1365:  817,
	1413:  817,
	1366:  818,
	1414:  818,
	1367:  819,
	1368:  820,
	1369:  821,
	1370:  822,
	1371:  823,
	1372:  824,
	1373:  825,
	1374:  826,
	1375:  827,
	1376:  828,
	7834:  7094,
	7836:  7095,
	7837:  7096,
	7838:  7097,
	7839:  7098,
	7930:  7099,
	7931:  7100,
	7932:  7101,
	7933:  7102,
	7934:  7103,
	7935:  7104,
	7958:  7105,
	7959:  7106,
	7966:  7107,
	7967:  7108,
	8006:  7109,
	8007:  7110,
	8014:  7111,
	8015:  7112,
	8024:  7113,
	8026:  7114,
	8028:  7115,
	8030:  7116,
	8062:  7117,
	8063:  7118,
	8117:  7119,
	8049:  7120,
	8123:  7120,
	8125:  7121,
	8127:  7122,
	8128:  7123,
	8129:  7124,
	8133:  7125,
	8051:  7126,
	8137:  7126,
	8053:  7127,
	8139:  7127,
	8141:  7128,
	8142:  7129,
	8143:  7130,
	8147:  7131,
	8148:  7132,
	8149:  7133,
	8055:  7134,
	8155:  7134,
	8156:  7135,
	8157:  7136,
	8158:  7137,
	8159:  7138,
	8163:  7139,
	8059:  7140,
	8171:  7140,
	8173:  7141,
	8174:  7142,
	8175:  7143,
	8176:  7144,
	8177:  7145,
	8181:  7146,
	8057:  7147,
	8185:  7147,
	8061:  7148,
	8187:  7148,
	8560:  7504,
	8545:  7505,
	8561:  7505,
	8546:  7506,
	8562:  7506,
	8547:  7507,
	8563:  7507,
	8548:  7508,
	8564:  7508,
	8549:  7509,
	8565:  7509,
	8550:  7510,
	8566:  7510,
	8551:  7511,
	8567:  7511,
	8552:  7512,
	8568:  7512,
	8553:  7513,
	8569:  7513,
	8554:  7514,
	8570:  7514,
	8555:  7515,
	8571:  7515,
	8556:  7516,
	8572:  7516,
	8557:  7517,
	8573:  7517,
	8558:  7518,
	8574:  7518,
	8559:  7519,
	9424:  8342,
	9399:  8343,
	9425:  8343,
	9400:  8344,
	9426:  8344,
	9401:  8345,
	9427:  8345,
	9402:  8346,
	9428:  8346,
	9403:  8347,
	9429:  8347,
	9404:  8348,
	9430:  8348,
	9405:  8349,
	9431:  8349,
	9406:  8350,
	9432:  8350,
	9407:  8351,
	9433:  8351,
	9408:  8352,
	9434:  8352,
	9409:  8353,
	9435:  8353,
	9410:  8354,
	9436:  8354,
	9411:  8355,
	9437:  8355,
	9412:  8356,
	9438:  8356,
	9413:  8357,
	9439:  8357,
	9414:  8358,
	9440:  8358,
	9415:  8359,
	9441:  8359,
	9416:  8360,
	9442:  8360,
	9417:  8361,
	9443:  8361,
	9418:  8362,
	9444:  8362,
	9419:  8363,
	9445:  8363,
	9420:  8364,
	9446:  8364,
	9421:  8365,
	9447:  8365,
	9422:  8366,
	9448:  8366,
	9423:  8367,
	57344: 54214,
	65345: 62183,
	65314: 62184,
	65346: 62184,
	65315: 62185,
	65347: 62185,
	65316: 62186,
	65348: 62186,
	65317: 62187,
	65349: 62187,
	65318: 62188,
	65350: 62188,
	65319: 62189,
	65351: 62189,
	65320: 62190,
	65352: 62190,
	65321: 62191,
	65353: 62191,
	65322: 62192,
	65354: 62192,
	65323: 62193,
	65355: 62193,
	65324: 62194,
	65356: 62194,
	65325: 62195,
	65357: 62195,
	65326: 62196,
	65358: 62196,
	65327: 62197,
	65359: 62197,
	65328: 62198,
	65360: 62198,
	65329: 62199,
	65361: 62199,
	65330: 62200,
	65362: 62200,
	65331: 62201,
	65363: 62201,
	65332: 62202,
	65364: 62202,
	65333: 62203,
	65365: 62203,
	65334: 62204,
	65366: 62204,
	65335: 62205,
	65367: 62205,
	65336: 62206,
	65368: 62206,
	65337: 62207,
	65369: 62207,
	65338: 62208,
	65370: 62208,
	65339: 62209,
	65340: 62210,
	65341: 62211,
	65342: 62212,
	65343: 62213,
	65344: 62214,
	65534: 62378,
	65535: 62379,
}
