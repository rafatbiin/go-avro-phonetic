package avro_phonetic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasic(t *testing.T) {
	assert.Equal(t, "ভ্ল", ParseOrDie("bhl"))
	assert.Equal(t, "ব্জ", ParseOrDie("bj"))
	assert.Equal(t, "ব্দ", ParseOrDie("bd"))
	assert.Equal(t, "ব্ব", ParseOrDie("bb"))
	assert.Equal(t, "ব্ল", ParseOrDie("bl"))
	assert.Equal(t, "ভ", ParseOrDie("bh"))
	assert.Equal(t, "ভ্ল", ParseOrDie("vl"))
	assert.Equal(t, "ব", ParseOrDie("b"))
	assert.Equal(t, "ভ", ParseOrDie("v"))
	assert.Equal(t, "চ্ঞ", ParseOrDie("cNG"))
	assert.Equal(t, "চ্ছ", ParseOrDie("cch"))
	assert.Equal(t, "চ্চ", ParseOrDie("cc"))
	assert.Equal(t, "ছ", ParseOrDie("ch"))
	assert.Equal(t, "চ", ParseOrDie("c"))
	assert.Equal(t, "ধ্ন", ParseOrDie("dhn"))
	assert.Equal(t, "ধ্ম", ParseOrDie("dhm"))
	assert.Equal(t, "দ্ঘ", ParseOrDie("dgh"))
	assert.Equal(t, "দ্ধ", ParseOrDie("ddh"))
	assert.Equal(t, "দ্ভ", ParseOrDie("dbh"))
	assert.Equal(t, "দ্ভ", ParseOrDie("dv"))
	assert.Equal(t, "দ্ম", ParseOrDie("dm"))
	assert.Equal(t, "ড্ড", ParseOrDie("DD"))
	assert.Equal(t, "ঢ", ParseOrDie("Dh"))
	assert.Equal(t, "ধ", ParseOrDie("dh"))
	assert.Equal(t, "দ্গ", ParseOrDie("dg"))
	assert.Equal(t, "দ্দ", ParseOrDie("dd"))
	assert.Equal(t, "ড", ParseOrDie("D"))
	assert.Equal(t, "দ", ParseOrDie("d"))
	assert.Equal(t, "...", ParseOrDie("..."))
	assert.Equal(t, ".", ParseOrDie(".`"))
	assert.Equal(t, "।।", ParseOrDie(".."))
	assert.Equal(t, "।", ParseOrDie("."))
	assert.Equal(t, "ঘ্ন", ParseOrDie("ghn"))
	assert.Equal(t, "ঘ্ন", ParseOrDie("Ghn"))
	assert.Equal(t, "গ্ধ", ParseOrDie("gdh"))
	assert.Equal(t, "গ্ণ", ParseOrDie("gN"))
	assert.Equal(t, "গ্ণ", ParseOrDie("GN"))
	assert.Equal(t, "গ্ন", ParseOrDie("gn"))
	assert.Equal(t, "গ্ম", ParseOrDie("gm"))
	assert.Equal(t, "গ্ম", ParseOrDie("Gm"))
	assert.Equal(t, "গ্ল", ParseOrDie("gl"))
	assert.Equal(t, "গ্ল", ParseOrDie("Gl"))
	assert.Equal(t, "জ্ঞ", ParseOrDie("gg"))
	assert.Equal(t, "জ্ঞ", ParseOrDie("GG"))
	assert.Equal(t, "জ্ঞ", ParseOrDie("Gg"))
	assert.Equal(t, "জ্ঞ", ParseOrDie("gG"))
	assert.Equal(t, "ঘ", ParseOrDie("gh"))
	assert.Equal(t, "ঘ", ParseOrDie("Gh"))
	assert.Equal(t, "গ", ParseOrDie("g"))
	assert.Equal(t, "হ্ণ", ParseOrDie("hN"))
	assert.Equal(t, "হ্ন", ParseOrDie("hn"))
	assert.Equal(t, "হ্ম", ParseOrDie("hm"))
	assert.Equal(t, "হ্ল", ParseOrDie("hl"))
	assert.Equal(t, "হ", ParseOrDie("h"))
	assert.Equal(t, "জ্ঝ", ParseOrDie("jjh"))
	assert.Equal(t, "জ্ঞ", ParseOrDie("jNG"))
	assert.Equal(t, "ঝ", ParseOrDie("jh"))
	assert.Equal(t, "জ্জ", ParseOrDie("jj"))
	assert.Equal(t, "জ", ParseOrDie("j"))
	assert.Equal(t, "জ", ParseOrDie("J"))
	assert.Equal(t, "ক্ষ্ণ", ParseOrDie("kkhN"))
	assert.Equal(t, "ক্ষ্ণ", ParseOrDie("kShN"))
	assert.Equal(t, "ক্ষ্ম", ParseOrDie("kkhm"))
	assert.Equal(t, "ক্ষ্ম", ParseOrDie("kShm"))
	assert.Equal(t, "ক্ষ্ণ", ParseOrDie("kxN"))
	assert.Equal(t, "ক্ষ্ম", ParseOrDie("kxm"))
	assert.Equal(t, "ক্ষ", ParseOrDie("kkh"))
	assert.Equal(t, "ক্ষ", ParseOrDie("kSh"))
	assert.Equal(t, "কশ", ParseOrDie("ksh"))
	assert.Equal(t, "ক্ষ", ParseOrDie("kx"))
	assert.Equal(t, "ক্ক", ParseOrDie("kk"))
	assert.Equal(t, "ক্ট", ParseOrDie("kT"))
	assert.Equal(t, "ক্ত", ParseOrDie("kt"))
	assert.Equal(t, "ক্ল", ParseOrDie("kl"))
	assert.Equal(t, "ক্স", ParseOrDie("ks"))
	assert.Equal(t, "খ", ParseOrDie("kh"))
	assert.Equal(t, "ক", ParseOrDie("k"))
	assert.Equal(t, "ল্ভ", ParseOrDie("lbh"))
	assert.Equal(t, "ল্ধ", ParseOrDie("ldh"))
	assert.Equal(t, "লখ", ParseOrDie("lkh"))
	assert.Equal(t, "লঘ", ParseOrDie("lgh"))
	assert.Equal(t, "লফ", ParseOrDie("lph"))
	assert.Equal(t, "ল্ক", ParseOrDie("lk"))
	assert.Equal(t, "ল্গ", ParseOrDie("lg"))
	assert.Equal(t, "ল্ট", ParseOrDie("lT"))
	assert.Equal(t, "ল্ড", ParseOrDie("lD"))
	assert.Equal(t, "ল্প", ParseOrDie("lp"))
	assert.Equal(t, "ল্ভ", ParseOrDie("lv"))
	assert.Equal(t, "ল্ম", ParseOrDie("lm"))
	assert.Equal(t, "ল্ল", ParseOrDie("ll"))
	assert.Equal(t, "ল্ব", ParseOrDie("lb"))
	assert.Equal(t, "ল", ParseOrDie("l"))
	assert.Equal(t, "ম্থ", ParseOrDie("mth"))
	assert.Equal(t, "ম্ফ", ParseOrDie("mph"))
	assert.Equal(t, "ম্ভ", ParseOrDie("mbh"))
	assert.Equal(t, "মপ্ল", ParseOrDie("mpl"))
	assert.Equal(t, "ম্ন", ParseOrDie("mn"))
	assert.Equal(t, "ম্প", ParseOrDie("mp"))
	assert.Equal(t, "ম্ভ", ParseOrDie("mv"))
	assert.Equal(t, "ম্ম", ParseOrDie("mm"))
	assert.Equal(t, "ম্ল", ParseOrDie("ml"))
	assert.Equal(t, "ম্ব", ParseOrDie("mb"))
	assert.Equal(t, "ম্ফ", ParseOrDie("mf"))
	assert.Equal(t, "ম", ParseOrDie("m"))
	assert.Equal(t, "০", ParseOrDie("0"))
	assert.Equal(t, "১", ParseOrDie("1"))
	assert.Equal(t, "২", ParseOrDie("2"))
	assert.Equal(t, "৩", ParseOrDie("3"))
	assert.Equal(t, "৪", ParseOrDie("4"))
	assert.Equal(t, "৫", ParseOrDie("5"))
	assert.Equal(t, "৬", ParseOrDie("6"))
	assert.Equal(t, "৭", ParseOrDie("7"))
	assert.Equal(t, "৮", ParseOrDie("8"))
	assert.Equal(t, "৯", ParseOrDie("9"))
	assert.Equal(t, "ঙ্ক্ষ", ParseOrDie("NgkSh"))
	assert.Equal(t, "ঙ্ক্ষ", ParseOrDie("Ngkkh"))
	assert.Equal(t, "ঞ্ছ", ParseOrDie("NGch"))
	assert.Equal(t, "ঙ্ঘ", ParseOrDie("Nggh"))
	assert.Equal(t, "ঙ্খ", ParseOrDie("Ngkh"))
	assert.Equal(t, "ঞ্ঝ", ParseOrDie("NGjh"))
	assert.Equal(t, "ঙ্গৌ", ParseOrDie("ngOU"))
	assert.Equal(t, "ঙ্গৈ", ParseOrDie("ngOI"))
	assert.Equal(t, "ঙ্ক্ষ", ParseOrDie("Ngkx"))
	assert.Equal(t, "ঞ্চ", ParseOrDie("NGc"))
	assert.Equal(t, "ঞ্ছ", ParseOrDie("nch"))
	assert.Equal(t, "ঞ্ঝ", ParseOrDie("njh"))
	assert.Equal(t, "ঙ্ঘ", ParseOrDie("ngh"))
	assert.Equal(t, "ঙ্ক", ParseOrDie("Ngk"))
	assert.Equal(t, "ঙ্ষ", ParseOrDie("Ngx"))
	assert.Equal(t, "ঙ্গ", ParseOrDie("Ngg"))
	assert.Equal(t, "ঙ্ম", ParseOrDie("Ngm"))
	assert.Equal(t, "ঞ্জ", ParseOrDie("NGj"))
	assert.Equal(t, "ন্ধ", ParseOrDie("ndh"))
	assert.Equal(t, "ন্ঠ", ParseOrDie("nTh"))
	assert.Equal(t, "ণ্ঠ", ParseOrDie("NTh"))
	assert.Equal(t, "ন্থ", ParseOrDie("nth"))
	assert.Equal(t, "ঙ্খ", ParseOrDie("nkh"))
	assert.Equal(t, "ঙ্গ", ParseOrDie("ngo"))
	assert.Equal(t, "ঙ্গা", ParseOrDie("nga"))
	assert.Equal(t, "ঙ্গি", ParseOrDie("ngi"))
	assert.Equal(t, "ঙ্গী", ParseOrDie("ngI"))
	assert.Equal(t, "ঙ্গু", ParseOrDie("ngu"))
	assert.Equal(t, "ঙ্গূ", ParseOrDie("ngU"))
	assert.Equal(t, "ঙ্গে", ParseOrDie("nge"))
	assert.Equal(t, "ঙ্গো", ParseOrDie("ngO"))
	assert.Equal(t, "ণ্ঢ", ParseOrDie("NDh"))
	assert.Equal(t, "নশ", ParseOrDie("nsh"))
	assert.Equal(t, "ঙর", ParseOrDie("Ngr"))
	assert.Equal(t, "ঞর", ParseOrDie("NGr"))
	assert.Equal(t, "ংর", ParseOrDie("ngr"))
	assert.Equal(t, "ঞ্জ", ParseOrDie("nj"))
	assert.Equal(t, "ঙ", ParseOrDie("Ng"))
	assert.Equal(t, "ঞ", ParseOrDie("NG"))
	assert.Equal(t, "ঙ্ক", ParseOrDie("nk"))
	assert.Equal(t, "ং", ParseOrDie("ng"))
	assert.Equal(t, "ন্ন", ParseOrDie("nn"))
	assert.Equal(t, "ণ্ণ", ParseOrDie("NN"))
	assert.Equal(t, "ণ্ন", ParseOrDie("Nn"))
	assert.Equal(t, "ন্ম", ParseOrDie("nm"))
	assert.Equal(t, "ণ্ম", ParseOrDie("Nm"))
	assert.Equal(t, "ন্দ", ParseOrDie("nd"))
	assert.Equal(t, "ন্ট", ParseOrDie("nT"))
	assert.Equal(t, "ণ্ট", ParseOrDie("NT"))
	assert.Equal(t, "ন্ড", ParseOrDie("nD"))
	assert.Equal(t, "ণ্ড", ParseOrDie("ND"))
	assert.Equal(t, "ন্ত", ParseOrDie("nt"))
	assert.Equal(t, "ন্স", ParseOrDie("ns"))
	assert.Equal(t, "ঞ্চ", ParseOrDie("nc"))
	assert.Equal(t, "ন", ParseOrDie("n"))
	assert.Equal(t, "ণ", ParseOrDie("N"))
	assert.Equal(t, "ৈ", ParseOrDie("OI`"))
	assert.Equal(t, "ৌ", ParseOrDie("OU`"))
	assert.Equal(t, "ো", ParseOrDie("O`"))
	assert.Equal(t, "ঐ", ParseOrDie("OI"))
	assert.Equal(t, "কৈ", ParseOrDie("kOI"))
	assert.Equal(t, " ঐ", ParseOrDie(" OI"))
	assert.Equal(t, "(ঐ", ParseOrDie("(OI"))
	assert.Equal(t, "।ঐ", ParseOrDie(".OI"))
	assert.Equal(t, "ঔ", ParseOrDie("OU"))
	assert.Equal(t, "কৌ", ParseOrDie("kOU"))
	assert.Equal(t, " ঔ", ParseOrDie(" OU"))
	assert.Equal(t, "-ঔ", ParseOrDie("-OU"))
	assert.Equal(t, "্‌ঔ", ParseOrDie(",,OU"))
	assert.Equal(t, "ও", ParseOrDie("O"))
	assert.Equal(t, "পো", ParseOrDie("pO"))
	assert.Equal(t, " ও", ParseOrDie(" O"))
	assert.Equal(t, "ইও", ParseOrDie("iO"))
	assert.Equal(t, "ও", ParseOrDie("`O"))
	assert.Equal(t, "ফ্ল", ParseOrDie("phl"))
	assert.Equal(t, "প্ট", ParseOrDie("pT"))
	assert.Equal(t, "প্ত", ParseOrDie("pt"))
	assert.Equal(t, "প্ন", ParseOrDie("pn"))
	assert.Equal(t, "প্প", ParseOrDie("pp"))
	assert.Equal(t, "প্ল", ParseOrDie("pl"))
	assert.Equal(t, "প্স", ParseOrDie("ps"))
	assert.Equal(t, "ফ", ParseOrDie("ph"))
	assert.Equal(t, "ফ্ল", ParseOrDie("fl"))
	assert.Equal(t, "ফ", ParseOrDie("f"))
	assert.Equal(t, "প", ParseOrDie("p"))
	assert.Equal(t, "ৃ", ParseOrDie("rri`"))
	assert.Equal(t, "ঋ", ParseOrDie("rri"))
	assert.Equal(t, "কৃ", ParseOrDie("krri"))
	assert.Equal(t, "ঈঋ", ParseOrDie("Irri"))
	assert.Equal(t, "ঁঋ", ParseOrDie("^rri"))
	assert.Equal(t, "ঃঋ", ParseOrDie(":rri"))
	assert.Equal(t, "র‍্য", ParseOrDie("rZ"))
	assert.Equal(t, "ক্র্য", ParseOrDie("krZ"))
	assert.Equal(t, "রর‍্য", ParseOrDie("rrZ"))
	assert.Equal(t, "ইয়র‍্য", ParseOrDie("yrZ"))
	assert.Equal(t, "ওর‍্য", ParseOrDie("wrZ"))
	assert.Equal(t, "এক্সর‍্য", ParseOrDie("xrZ"))
	assert.Equal(t, "ইর‍্য", ParseOrDie("irZ"))
	assert.Equal(t, "-র‍্য", ParseOrDie("-rZ"))
	assert.Equal(t, "ররর‍্য", ParseOrDie("rrrZ"))
	assert.Equal(t, "র‍্য", ParseOrDie("ry"))
	assert.Equal(t, "ক্র্য", ParseOrDie("qry"))
	assert.Equal(t, "রর‍্য", ParseOrDie("rry"))
	assert.Equal(t, "ইয়র‍্য", ParseOrDie("yry"))
	assert.Equal(t, "ওর‍্য", ParseOrDie("wry"))
	assert.Equal(t, "এক্সর‍্য", ParseOrDie("xry"))
	assert.Equal(t, "০র‍্য", ParseOrDie("0ry"))
	assert.Equal(t, "রররর‍্য", ParseOrDie("rrrry"))
	assert.Equal(t, "ড়্র্য", ParseOrDie("Rry"))
	assert.Equal(t, "রর", ParseOrDie("rr"))
	assert.Equal(t, "আরর", ParseOrDie("arr"))
	assert.Equal(t, "আর্ক", ParseOrDie("arrk"))
	assert.Equal(t, "আররা", ParseOrDie("arra"))
	assert.Equal(t, "আরর", ParseOrDie("arr"))
	assert.Equal(t, "আরর!", ParseOrDie("arr!"))
	assert.Equal(t, "ক্রর", ParseOrDie("krr"))
	assert.Equal(t, "ক্ররা", ParseOrDie("krra"))
	assert.Equal(t, "ড়্গ", ParseOrDie("Rg"))
	assert.Equal(t, "ঢ়", ParseOrDie("Rh"))
	assert.Equal(t, "ড়", ParseOrDie("R"))
	assert.Equal(t, "র", ParseOrDie("r"))
	assert.Equal(t, "অর", ParseOrDie("or"))
	assert.Equal(t, "ম্র", ParseOrDie("mr"))
	assert.Equal(t, "১র", ParseOrDie("1r"))
	assert.Equal(t, "+র", ParseOrDie("+r"))
	assert.Equal(t, "রর", ParseOrDie("rr"))
	assert.Equal(t, "ইয়র", ParseOrDie("yr"))
	assert.Equal(t, "ওর", ParseOrDie("wr"))
	assert.Equal(t, "এক্সর", ParseOrDie("xr"))
	assert.Equal(t, "য্র", ParseOrDie("zr"))
	assert.Equal(t, "ম্রি", ParseOrDie("mri"))
	assert.Equal(t, "শ্ছ", ParseOrDie("shch"))
	assert.Equal(t, "ষ্ঠ", ParseOrDie("ShTh"))
	assert.Equal(t, "ষ্ফ", ParseOrDie("Shph"))
	assert.Equal(t, "শ্ছ", ParseOrDie("Sch"))
	assert.Equal(t, "স্ক্ল", ParseOrDie("skl"))
	assert.Equal(t, "স্খ", ParseOrDie("skh"))
	assert.Equal(t, "স্থ", ParseOrDie("sth"))
	assert.Equal(t, "স্ফ", ParseOrDie("sph"))
	assert.Equal(t, "শ্চ", ParseOrDie("shc"))
	assert.Equal(t, "শ্ত", ParseOrDie("sht"))
	assert.Equal(t, "শ্ন", ParseOrDie("shn"))
	assert.Equal(t, "শ্ম", ParseOrDie("shm"))
	assert.Equal(t, "শ্ল", ParseOrDie("shl"))
	assert.Equal(t, "ষ্ক", ParseOrDie("Shk"))
	assert.Equal(t, "ষ্ট", ParseOrDie("ShT"))
	assert.Equal(t, "ষ্ণ", ParseOrDie("ShN"))
	assert.Equal(t, "ষ্প", ParseOrDie("Shp"))
	assert.Equal(t, "ষ্ফ", ParseOrDie("Shf"))
	assert.Equal(t, "ষ্ম", ParseOrDie("Shm"))
	assert.Equal(t, "স্প্ল", ParseOrDie("spl"))
	assert.Equal(t, "স্ক", ParseOrDie("sk"))
	assert.Equal(t, "শ্চ", ParseOrDie("Sc"))
	assert.Equal(t, "স্ট", ParseOrDie("sT"))
	assert.Equal(t, "স্ত", ParseOrDie("st"))
	assert.Equal(t, "স্ন", ParseOrDie("sn"))
	assert.Equal(t, "স্প", ParseOrDie("sp"))
	assert.Equal(t, "স্ফ", ParseOrDie("sf"))
	assert.Equal(t, "স্ম", ParseOrDie("sm"))
	assert.Equal(t, "স্ল", ParseOrDie("sl"))
	assert.Equal(t, "শ", ParseOrDie("sh"))
	assert.Equal(t, "শ্চ", ParseOrDie("Sc"))
	assert.Equal(t, "শ্ত", ParseOrDie("St"))
	assert.Equal(t, "শ্ন", ParseOrDie("Sn"))
	assert.Equal(t, "শ্ম", ParseOrDie("Sm"))
	assert.Equal(t, "শ্ল", ParseOrDie("Sl"))
	assert.Equal(t, "ষ", ParseOrDie("Sh"))
	assert.Equal(t, "স", ParseOrDie("s"))
	assert.Equal(t, "শ", ParseOrDie("S"))
	assert.Equal(t, "উ", ParseOrDie("oo"))
	assert.Equal(t, "ওও", ParseOrDie("OO"))
	assert.Equal(t, "ু", ParseOrDie("oo`"))
	assert.Equal(t, "কু", ParseOrDie("koo"))
	assert.Equal(t, "উঅ", ParseOrDie("ooo"))
	assert.Equal(t, "!উ", ParseOrDie("!oo"))
	assert.Equal(t, "!উঅ", ParseOrDie("!ooo"))
	assert.Equal(t, "আউ", ParseOrDie("aoo"))
	assert.Equal(t, "উপ", ParseOrDie("oop"))
	assert.Equal(t, "উ", ParseOrDie("ooo`"))
	assert.Equal(t, "", ParseOrDie("o`"))
	assert.Equal(t, "অ্য", ParseOrDie("oZ"))
	assert.Equal(t, "অয়", ParseOrDie("oY"))
	assert.Equal(t, "অ", ParseOrDie("o"))
	assert.Equal(t, "!অ", ParseOrDie("!o"))
	assert.Equal(t, "ঁঅ", ParseOrDie("^o"))
	assert.Equal(t, "*অ", ParseOrDie("*o"))
	assert.Equal(t, "ইও", ParseOrDie("io"))
	assert.Equal(t, "ইয়", ParseOrDie("yo"))
	assert.Equal(t, "ন", ParseOrDie("no"))
	assert.Equal(t, "ত্থ", ParseOrDie("tth"))
	assert.Equal(t, "ৎ", ParseOrDie("t``"))
	assert.Equal(t, "ৎ", ParseOrDie("`t``"))
	assert.Equal(t, "ৎৎ", ParseOrDie("t``t``"))
	assert.Equal(t, "ৎ", ParseOrDie("t```"))
	assert.Equal(t, "ট্ট", ParseOrDie("TT"))
	assert.Equal(t, "ট্ম", ParseOrDie("Tm"))
	assert.Equal(t, "ঠ", ParseOrDie("Th"))
	assert.Equal(t, "ত্ন", ParseOrDie("tn"))
	assert.Equal(t, "ত্ম", ParseOrDie("tm"))
	assert.Equal(t, "থ", ParseOrDie("th"))
	assert.Equal(t, "ত্ত", ParseOrDie("tt"))
	assert.Equal(t, "ট", ParseOrDie("T"))
	assert.Equal(t, "ত", ParseOrDie("t"))
	assert.Equal(t, "অ্যা", ParseOrDie("aZ"))
	assert.Equal(t, "আঅ্যা", ParseOrDie("aaZ"))
	assert.Equal(t, "অ্যা", ParseOrDie("AZ"))
	assert.Equal(t, "া", ParseOrDie("a`"))
	assert.Equal(t, "া", ParseOrDie("a``"))
	assert.Equal(t, "কা", ParseOrDie("ka`"))
	assert.Equal(t, "া", ParseOrDie("A`"))
	assert.Equal(t, "আ", ParseOrDie("a"))
	assert.Equal(t, "আ", ParseOrDie("`a"))
	assert.Equal(t, "কআ", ParseOrDie("k`a"))
	assert.Equal(t, "ইয়া", ParseOrDie("ia"))
	assert.Equal(t, "আআআা", ParseOrDie("aaaa`"))
	assert.Equal(t, "ি", ParseOrDie("i`"))
	assert.Equal(t, "ই", ParseOrDie("i"))
	assert.Equal(t, "ই", ParseOrDie("`i"))
	assert.Equal(t, "হি", ParseOrDie("hi"))
	assert.Equal(t, "ইহ", ParseOrDie("ih"))
	assert.Equal(t, "িহ", ParseOrDie("i`h"))
	assert.Equal(t, "ী", ParseOrDie("I`"))
	assert.Equal(t, "ঈ", ParseOrDie("I"))
	assert.Equal(t, "চী", ParseOrDie("cI"))
	assert.Equal(t, "ঈক্স", ParseOrDie("Ix"))
	assert.Equal(t, "ঈঈ", ParseOrDie("II"))
	assert.Equal(t, "০ঈ", ParseOrDie("0I"))
	assert.Equal(t, "অঈ", ParseOrDie("oI"))
	assert.Equal(t, "ু", ParseOrDie("u`"))
	assert.Equal(t, "উ", ParseOrDie("u"))
	assert.Equal(t, "কু", ParseOrDie("ku"))
	assert.Equal(t, "উক", ParseOrDie("uk"))
	assert.Equal(t, "উউ", ParseOrDie("uu"))
	assert.Equal(t, "ইউ", ParseOrDie("iu"))
	assert.Equal(t, "&উ", ParseOrDie("&u"))
	assert.Equal(t, "উ&", ParseOrDie("u&"))
	assert.Equal(t, "ূ", ParseOrDie("U`"))
	assert.Equal(t, "ঊ", ParseOrDie("U"))
	assert.Equal(t, "ইয়ূ", ParseOrDie("yU"))
	assert.Equal(t, "ঊয়", ParseOrDie("Uy"))
	assert.Equal(t, "ঁঊ", ParseOrDie("^U"))
	assert.Equal(t, "ঊঁ", ParseOrDie("U^"))
	assert.Equal(t, "ঈ", ParseOrDie("EE"))
	assert.Equal(t, "ঈ", ParseOrDie("ee"))
	assert.Equal(t, "ঈ", ParseOrDie("Ee"))
	assert.Equal(t, "ঈ", ParseOrDie("eE"))
	assert.Equal(t, "ী", ParseOrDie("ee`"))
	assert.Equal(t, "কী", ParseOrDie("kee"))
	assert.Equal(t, "ঈক", ParseOrDie("eek"))
	assert.Equal(t, "০ঈ", ParseOrDie("0ee"))
	assert.Equal(t, "ঈ৮", ParseOrDie("ee8"))
	assert.Equal(t, "(ঈ)", ParseOrDie("(ee)"))
	assert.Equal(t, "ে", ParseOrDie("e`"))
	assert.Equal(t, "এ", ParseOrDie("e"))
	assert.Equal(t, "কে", ParseOrDie("ke"))
	assert.Equal(t, "ওয়ে", ParseOrDie("we"))
	assert.Equal(t, "#এ#", ParseOrDie("#e#"))
	assert.Equal(t, "ে", ParseOrDie("`e`"))
	assert.Equal(t, "য", ParseOrDie("z"))
	assert.Equal(t, "্য", ParseOrDie("Z"))
	assert.Equal(t, "র‍্য", ParseOrDie("rZ"))
	assert.Equal(t, "ক্যশ", ParseOrDie("kZS"))
	assert.Equal(t, "ইয়", ParseOrDie("y"))
	assert.Equal(t, "অয়", ParseOrDie("oy"))
	assert.Equal(t, "ক্য", ParseOrDie("ky"))
	assert.Equal(t, "ইয়া", ParseOrDie("ya"))
	assert.Equal(t, "ইয়াআ", ParseOrDie("yaa"))
	assert.Equal(t, "য়", ParseOrDie("Y"))
	assert.Equal(t, "য়য়", ParseOrDie("YY"))
	assert.Equal(t, "ইয়", ParseOrDie("iY"))
	assert.Equal(t, "কয়", ParseOrDie("kY"))
	assert.Equal(t, "ক", ParseOrDie("q"))
	assert.Equal(t, "ক", ParseOrDie("Q"))
	assert.Equal(t, "ও", ParseOrDie("w"))
	assert.Equal(t, "ওয়া", ParseOrDie("wa"))
	assert.Equal(t, "-ওয়া-", ParseOrDie("-wa-"))
	assert.Equal(t, "ওয়ু", ParseOrDie("woo"))
	assert.Equal(t, "ওরে", ParseOrDie("wre"))
	assert.Equal(t, "ক্ব", ParseOrDie("kw"))
	assert.Equal(t, "এক্স", ParseOrDie("x"))
	assert.Equal(t, "এক্স", ParseOrDie("ex"))
	assert.Equal(t, "বক্স", ParseOrDie("bx"))
	assert.Equal(t, ":", ParseOrDie(":`"))
	assert.Equal(t, "ঃ", ParseOrDie(":"))
	assert.Equal(t, "^", ParseOrDie("^`"))
	assert.Equal(t, "ঁ", ParseOrDie("^"))
	assert.Equal(t, "কঁ", ParseOrDie("k^"))
	assert.Equal(t, "কঁই", ParseOrDie("k^i"))
	assert.Equal(t, "কিঁ", ParseOrDie("ki^"))
	assert.Equal(t, "্‌", ParseOrDie(",,"))
	assert.Equal(t, "্‌,", ParseOrDie(",,,"))
	assert.Equal(t, "্‌,", ParseOrDie(",,`,"))
	assert.Equal(t, "্‌", ParseOrDie("`,,"))
	assert.Equal(t, ",,", ParseOrDie(",`,"))
	assert.Equal(t, "৳", ParseOrDie("$"))
	assert.Equal(t, "", ParseOrDie("`"))
	assert.Equal(t, "ব্ধ", ParseOrDie("bdh"))
}

func TestRandomChars(t *testing.T) {
	assert.Equal(t, ParseOrDie("!"), "!")
}

func TestSentence(t *testing.T) {
	assert.Equal(t, "আমি বাংলায় গান গাই", ParseOrDie("ami banglay gan gai"))
	assert.Equal(t, "আমাদের ভালোবাসা হয়ে গেল ঘাস, খেয়ে গেল গরু আর দিয়ে গেল বাঁশ", ParseOrDie("amader valObasa hoye gel ghas, kheye gel goru ar diye gelo ba^sh"))
}
