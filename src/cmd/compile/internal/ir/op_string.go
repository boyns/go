// Code generated by "stringer -type=Op -trimprefix=O node.go"; DO NOT EDIT.

package ir

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[OXXX-0]
	_ = x[ONAME-1]
	_ = x[ONONAME-2]
	_ = x[OTYPE-3]
	_ = x[OPACK-4]
	_ = x[OLITERAL-5]
	_ = x[ONIL-6]
	_ = x[OADD-7]
	_ = x[OSUB-8]
	_ = x[OOR-9]
	_ = x[OXOR-10]
	_ = x[OADDSTR-11]
	_ = x[OADDR-12]
	_ = x[OANDAND-13]
	_ = x[OAPPEND-14]
	_ = x[OBYTES2STR-15]
	_ = x[OBYTES2STRTMP-16]
	_ = x[ORUNES2STR-17]
	_ = x[OSTR2BYTES-18]
	_ = x[OSTR2BYTESTMP-19]
	_ = x[OSTR2RUNES-20]
	_ = x[OAS-21]
	_ = x[OAS2-22]
	_ = x[OAS2DOTTYPE-23]
	_ = x[OAS2FUNC-24]
	_ = x[OAS2MAPR-25]
	_ = x[OAS2RECV-26]
	_ = x[OASOP-27]
	_ = x[OCALL-28]
	_ = x[OCALLFUNC-29]
	_ = x[OCALLMETH-30]
	_ = x[OCALLINTER-31]
	_ = x[OCALLPART-32]
	_ = x[OCAP-33]
	_ = x[OCLOSE-34]
	_ = x[OCLOSURE-35]
	_ = x[OCOMPLIT-36]
	_ = x[OMAPLIT-37]
	_ = x[OSTRUCTLIT-38]
	_ = x[OARRAYLIT-39]
	_ = x[OSLICELIT-40]
	_ = x[OPTRLIT-41]
	_ = x[OCONV-42]
	_ = x[OCONVIFACE-43]
	_ = x[OCONVNOP-44]
	_ = x[OCOPY-45]
	_ = x[ODCL-46]
	_ = x[ODCLFUNC-47]
	_ = x[ODCLCONST-48]
	_ = x[ODCLTYPE-49]
	_ = x[ODELETE-50]
	_ = x[ODOT-51]
	_ = x[ODOTPTR-52]
	_ = x[ODOTMETH-53]
	_ = x[ODOTINTER-54]
	_ = x[OXDOT-55]
	_ = x[ODOTTYPE-56]
	_ = x[ODOTTYPE2-57]
	_ = x[OEQ-58]
	_ = x[ONE-59]
	_ = x[OLT-60]
	_ = x[OLE-61]
	_ = x[OGE-62]
	_ = x[OGT-63]
	_ = x[ODEREF-64]
	_ = x[OINDEX-65]
	_ = x[OINDEXMAP-66]
	_ = x[OKEY-67]
	_ = x[OSTRUCTKEY-68]
	_ = x[OLEN-69]
	_ = x[OMAKE-70]
	_ = x[OMAKECHAN-71]
	_ = x[OMAKEMAP-72]
	_ = x[OMAKESLICE-73]
	_ = x[OMAKESLICECOPY-74]
	_ = x[OMUL-75]
	_ = x[ODIV-76]
	_ = x[OMOD-77]
	_ = x[OLSH-78]
	_ = x[ORSH-79]
	_ = x[OAND-80]
	_ = x[OANDNOT-81]
	_ = x[ONEW-82]
	_ = x[ONOT-83]
	_ = x[OBITNOT-84]
	_ = x[OPLUS-85]
	_ = x[ONEG-86]
	_ = x[OOROR-87]
	_ = x[OPANIC-88]
	_ = x[OPRINT-89]
	_ = x[OPRINTN-90]
	_ = x[OPAREN-91]
	_ = x[OSEND-92]
	_ = x[OSLICE-93]
	_ = x[OSLICEARR-94]
	_ = x[OSLICESTR-95]
	_ = x[OSLICE3-96]
	_ = x[OSLICE3ARR-97]
	_ = x[OSLICEHEADER-98]
	_ = x[ORECOVER-99]
	_ = x[ORECV-100]
	_ = x[ORUNESTR-101]
	_ = x[OSELRECV2-102]
	_ = x[OIOTA-103]
	_ = x[OREAL-104]
	_ = x[OIMAG-105]
	_ = x[OCOMPLEX-106]
	_ = x[OALIGNOF-107]
	_ = x[OOFFSETOF-108]
	_ = x[OSIZEOF-109]
	_ = x[OMETHEXPR-110]
	_ = x[OSTMTEXPR-111]
	_ = x[OBLOCK-112]
	_ = x[OBREAK-113]
	_ = x[OCASE-114]
	_ = x[OCONTINUE-115]
	_ = x[ODEFER-116]
	_ = x[OFALL-117]
	_ = x[OFOR-118]
	_ = x[OFORUNTIL-119]
	_ = x[OGOTO-120]
	_ = x[OIF-121]
	_ = x[OLABEL-122]
	_ = x[OGO-123]
	_ = x[ORANGE-124]
	_ = x[ORETURN-125]
	_ = x[OSELECT-126]
	_ = x[OSWITCH-127]
	_ = x[OTYPESW-128]
	_ = x[OTCHAN-129]
	_ = x[OTMAP-130]
	_ = x[OTSTRUCT-131]
	_ = x[OTINTER-132]
	_ = x[OTFUNC-133]
	_ = x[OTARRAY-134]
	_ = x[OTSLICE-135]
	_ = x[OINLCALL-136]
	_ = x[OEFACE-137]
	_ = x[OITAB-138]
	_ = x[OIDATA-139]
	_ = x[OSPTR-140]
	_ = x[OCFUNC-141]
	_ = x[OCHECKNIL-142]
	_ = x[OVARDEF-143]
	_ = x[OVARKILL-144]
	_ = x[OVARLIVE-145]
	_ = x[ORESULT-146]
	_ = x[OINLMARK-147]
	_ = x[OLINKSYMOFFSET-148]
	_ = x[OTAILCALL-149]
	_ = x[OGETG-150]
	_ = x[OEND-151]
}

const _Op_name = "XXXNAMENONAMETYPEPACKLITERALNILADDSUBORXORADDSTRADDRANDANDAPPENDBYTES2STRBYTES2STRTMPRUNES2STRSTR2BYTESSTR2BYTESTMPSTR2RUNESASAS2AS2DOTTYPEAS2FUNCAS2MAPRAS2RECVASOPCALLCALLFUNCCALLMETHCALLINTERCALLPARTCAPCLOSECLOSURECOMPLITMAPLITSTRUCTLITARRAYLITSLICELITPTRLITCONVCONVIFACECONVNOPCOPYDCLDCLFUNCDCLCONSTDCLTYPEDELETEDOTDOTPTRDOTMETHDOTINTERXDOTDOTTYPEDOTTYPE2EQNELTLEGEGTDEREFINDEXINDEXMAPKEYSTRUCTKEYLENMAKEMAKECHANMAKEMAPMAKESLICEMAKESLICECOPYMULDIVMODLSHRSHANDANDNOTNEWNOTBITNOTPLUSNEGORORPANICPRINTPRINTNPARENSENDSLICESLICEARRSLICESTRSLICE3SLICE3ARRSLICEHEADERRECOVERRECVRUNESTRSELRECV2IOTAREALIMAGCOMPLEXALIGNOFOFFSETOFSIZEOFMETHEXPRSTMTEXPRBLOCKBREAKCASECONTINUEDEFERFALLFORFORUNTILGOTOIFLABELGORANGERETURNSELECTSWITCHTYPESWTCHANTMAPTSTRUCTTINTERTFUNCTARRAYTSLICEINLCALLEFACEITABIDATASPTRCFUNCCHECKNILVARDEFVARKILLVARLIVERESULTINLMARKLINKSYMOFFSETTAILCALLGETGEND"

var _Op_index = [...]uint16{0, 3, 7, 13, 17, 21, 28, 31, 34, 37, 39, 42, 48, 52, 58, 64, 73, 85, 94, 103, 115, 124, 126, 129, 139, 146, 153, 160, 164, 168, 176, 184, 193, 201, 204, 209, 216, 223, 229, 238, 246, 254, 260, 264, 273, 280, 284, 287, 294, 302, 309, 315, 318, 324, 331, 339, 343, 350, 358, 360, 362, 364, 366, 368, 370, 375, 380, 388, 391, 400, 403, 407, 415, 422, 431, 444, 447, 450, 453, 456, 459, 462, 468, 471, 474, 480, 484, 487, 491, 496, 501, 507, 512, 516, 521, 529, 537, 543, 552, 563, 570, 574, 581, 589, 593, 597, 601, 608, 615, 623, 629, 637, 645, 650, 655, 659, 667, 672, 676, 679, 687, 691, 693, 698, 700, 705, 711, 717, 723, 729, 734, 738, 745, 751, 756, 762, 768, 775, 780, 784, 789, 793, 798, 806, 812, 819, 826, 832, 839, 852, 860, 864, 867}

func (i Op) String() string {
	if i >= Op(len(_Op_index)-1) {
		return "Op(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Op_name[_Op_index[i]:_Op_index[i+1]]
}
