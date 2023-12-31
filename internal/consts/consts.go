package consts

type Course struct {
	Number   string
	Code     string
	Name     string
	Teacher  string
	Time     string
	Location string
}

var Courses = [...]Course{
	{"0885", "01UG003", "西洋古典音樂", "劉思涵", "四8-9", "公館\nＢ102"},
	{"0886", "01UG004", "音樂鑑賞 ", "李娓娓", "四3-4", "本部\n誠101"},
	{"0888", "01UG004", "音樂鑑賞 ", "劉錦紋", "二8-9", "公館\nＢ102"},
	{"0887", "01UG004", "音樂鑑賞 ", "陳惠齡", "二3-4", "本部\n綜210展覽廳"},
	{"0889", "01UG005", "歌劇大師經典作品賞析", "王望舒", "五6-7", "本部\n誠104"},
	{"0890", "01UG006", "臺灣小說選讀", "林淑慧", "三3-4", "本部\n教202會議廳"},
	{"0891", "01UG008", "旅行文學", "陳室如", "二3-4", "本部\n美術館601"},
	{"0892", "01UG011", "亞裔美國文學與電影的藝術形式", "左乙萱", "五8-9", "本部\n誠402"},
	{"0893", "01UG011", "亞裔美國文學與電影的藝術形式", "左乙萱", "◎面授/同步", "◎面授/同步"},
	{"0894", "01UG012", "舞蹈鑑賞 ", "王宏豪", "二6-8", "本部\n誠201"},
	{"0895", "01UG014", "藝術鑑賞與視覺文化 ", "李立鈞", "一9-10", "本部\n美101"},
	{"0896", "01UG016", "戲劇鑑賞 ", "未定", "三3-4", "本部\n正105"},
	{"0897", "01UG018", "民間文學與說唱藝術", "蔡孟珍", "三3-4", "本部\n誠202"},
	{"0898", "01UG021", "文學與電影", "蘇榕", "◎面授/同步", "◎面授/同步"},
	{"0899", "01UG027", "日本美術鑑賞", "蔡家丘", "四8-9", "本部\n美術館601"},
	{"0900", "01UG029", "古典詩選讀", "何維剛", "一3-4", "本部\n樸202"},
	{"0901", "01UG031", "戲曲旋轉門", "陳芳", "一8-9", "本部\n誠208"},
	{"0902", "01UG033", "音樂與戲劇", "王望舒", "三3-4", "本部\n誠102"},
	{"0903", "01UG037", "高行健與表演藝術", "黃于真", "三8-9", "本部\n美術館601"},
	{"0904", "02UG001", "哲學入門", "傅皓政", "一3-4", "本部\n誠301"},
	{"0905", "02UG001", "哲學入門", "王鍾山", "二6-7", "公館\n理圖003"},
	{"0906", "02UG002", "思考方法", "紀金慶", "四6-7", "公館\n理圖001"},
	{"0907", "02UG005", "道德推理", "洪仁進", "一8-9", "本部\n教201演講廳"},
	{"0908", "02UG007", "中國哲人的生命與心靈", "黃瑩暖", "三8-9", "本部\n誠203"},
	{"0909", "02UG008", "中國古代思想的現代意義", "謝聰輝", "三8-9", "本部\n誠302"},
	{"0910", "02UG010", "人生哲學與批判性思考", "鄭鈞瑋", "三3-4", "本部\n誠108"},
	{"0911", "02UG010", "人生哲學與批判性思考", "鄭鈞瑋", "三6-7", "本部\n誠203"},
	{"0912", "02UG010", "人生哲學與批判性思考", "紀金慶", "四3-4", "公館\n理圖001"},
	{"0913", "02UG011", "從精神分析看生命潛則", "李秀娟", "三3-4", "本部\n誠203"},
	{"0914", "02UG015", "生物哲學導論", "林陳涌", "二8-9", "公館\n理圖003"},
	{"0915", "02UG019", "佛學中的人生智慧", "釋見見", "三8-9", "本部\n教201演講廳"},
	{"0916", "03UG006", "憲法與人權 ", "呂啟民", "四6-7", "本部\n誠301"},
	{"0917", "03UG012", "臺灣流行文化", "莊佳穎", "三6-7", "本部\n誠101"},
	{"0918", "03UG024", "環境與傳播", "楊樺", "一3-4", "公館\nＥ102"},
	{"0919", "03UG027", "女性文學、性別平等理論與婦運", "左乙萱", "五6-7", "本部\n誠402"},
	{"0920", "03UG027", "女性文學、性別平等理論與婦運", "左乙萱", "◎面授/同步", "◎面授/同步"},
	{"0921", "03UG031", "鑑識科學概論", "姜雲生", "四3-4", "本部\n綜210展覽廳"},
	{"0922", "03UG041", "管理學入門", "廖經維", "一1-2", "本部\n誠203"},
	{"0923", "03UG041", "管理學入門", "廖經維", "一3-4", "本部\n誠203"},
	{"0924", "03UG041", "管理學入門", "徐揚", "五6-7", "公館\n理圖001"},
	{"0926", "03UG043", "國際企業策略", "徐揚", "三6-7", "公館\nＳ501"},
	{"0927", "03UG047", "生活中的SDGs", "蘇淑娟", "三3-4", "本部\n誠104"},
	{"0928", "04UG003", "文明的探索", "陳健文", "一8-9", "本部\n誠307"},
	{"0929", "04UG003", "文明的探索", "陳南旭", "三3-4", "本部\n誠306"},
	{"0930", "04UG007", "西方文明史", "陳建元", "五3-4", "本部\n誠201"},
	{"0932", "04UG013", "臺灣歷史與電影", "陳佳宏", "五3-4", "本部\n誠101"},
	{"0934", "04UG024", "臺灣文史與城市發展", "黃玫瑄", "二3-4", "本部\n誠301"},
	{"0935", "04UG024", "臺灣文史與城市發展", "黃玫瑄", "二6-7", "本部\n誠301"},
	{"0936", "04UG029", "《楚辭‧九歌》與神話、巫風", "郭乃禎", "三3-4", "本部\n樸301"},
	{"0937", "04UG030", "客家社會探究", "黃玫瑄", "四3-4", "本部\n誠301"},
	{"0938", "04UG033", "東南亞文化與歷史概略", "廖弘民", "五6-7", "本部\n誠105"},
	{"0939", "04UG036", "臺灣原住民族文化展演 ", "王宏豪", "三3-4", "本部\n誠101"},
	{"0940", "05UG001", "邏輯思考與應用", "王銀國", "二3-4", "本部\n誠101"},
	{"0941", "05UG001", "邏輯思考與應用", "王銀國", "二8-9", "本部\n誠101"},
	{"0942", "05UG001", "邏輯思考與應用", "蔡承志", "二6-7", "本部\n誠109"},
	{"0943", "05UG001", "邏輯思考與應用", "王鍾山", "二3-4", "公館\n理圖001"},
	{"0944", "05UG002", "科學思維 ", "鍾兆晉", "一10-A", "本部\n誠109"},
	{"0945", "05UG002", "科學思維 ", "鍾兆晉", "一B-C", "本部\n誠109"},
	{"0946", "05UG007", "科學本質", "邱國力", "三3-4", "本部\n誠401"},
	{"0947", "05UG009", "科技與社會 ", "林陳涌", "二3-4", "公館\n理圖003"},
	{"0948", "05UG010", "科學的語言", "蔡承志", "二8-9", "本部\n誠109"},
	{"0949", "05UG011", "科技倫理", "劉啟民", "一6-7", "本部\n誠109"},
	{"0950", "05UG011", "科技倫理", "劉啟民", "一8-9", "本部\n誠109"},
	{"0951", "05UG012", "科技史哲思維下的世界觀", "陳珍源", "五1-2", "本部\n誠201"},
	{"0952", "05UG015", "生活中的統計", "林昌平", "三6-7", "本部\n誠401"},
	{"0953", "05UG016", "運算思維與程式設計 ", "張凌倩", "一9-10", "本部\n教401"},
	{"0954", "05UG016", "運算思維與程式設計 ", "簡培修", "二8-9", "本部\n教401"},
	{"0955", "05UG016", "運算思維與程式設計 ", "張凌倩", "二9-10", "公館\n理圖801"},
	{"0956", "05UG016", "運算思維與程式設計 ", "簡培修", "三8-9", "本部\n教401"},
	{"0957", "05UG016", "運算思維與程式設計 ", "蔡明男", "四6-7", "本部\n教401"},
	{"0958", "05UG016", "運算思維與程式設計 ", "邵森蘭", "四6-7", "公館\n理圖801"},
	{"0959", "05UG016", "運算思維與程式設計 ", "蔡明男", "四8-9", "本部\n教401"},
	{"0960", "05UG016", "運算思維與程式設計 ", "江政杰", "五1-2", "公館\n理圖801"},
	{"0961", "05UG016", "運算思維與程式設計 ", "江政杰", "五3-4", "公館\n理圖801"},
	{"0962", "05UG016", "運算思維與程式設計 ", "黃芳蘭", "五9-10", "公館\n理圖801"},
	{"0963", "05UG017", "資料科學與程式設計 ", "林耘逸", "三9-10", "本部\n教403"},
	{"0964", "05UG017", "資料科學與程式設計 ", "黃浩霆", "五3-4", "本部\n誠401"},
	{"0965", "05UG019", "行動App程式設計 ", "劉建成", "四6-7", "本部\n教402"},
	{"0966", "05UG020", "統計學", "何英奇", "二6-7", "本部\n教401"},
	{"0967", "05UG021", "試算表進階應用與程式設計", "吳清輝", "三8-9", "本部\n教402"},
	{"0968", "05UG021", "試算表進階應用與程式設計", "吳清輝", "五8-9", "本部\n教401"},
	{"0969", "05UG022", "學程式玩音樂 ", "林宜徵", "二3-4", "本部\n教401"},
	{"0970", "05UG023", "文本分析與程式設計", "王祥安", "五3-4", "本部\n教403"},
	{"0971", "05UG023", "文本分析與程式設計", "王祥安", "五6-7", "本部\n教402"},
	{"0972", "05UG025", "遊戲程式設計 ", "劉建成", "二6-7", "公館\n理圖801"},
	{"0973", "05UG025", "遊戲程式設計 ", "劉建成", "五3-4", "本部\n教402"},
	{"0974", "05UG026", "運動數據分析與程式設計", "蔣宗哲", "三3-4", "公館\n理圖801"},
	{"0975", "05UG027", "數位音樂與聲音合成之基礎程式設計", "林宜徵", "二6-7", "本部\n教404"},
	{"0976", "05UG028", "試算表與商務資料分析", "劉建成", "二3-4", "公館\n理圖801"},
	{"0977", "05UG028", "試算表與商務資料分析", "劉建成", "四3-4", "本部\n教401"},
	{"0978", "05UG029", "生醫與健康數據分析", "林耘逸", "◎面授/同步", "◎面授/同步"},
	{"0979", "05UG030", "大數據程式設計 ", "林政宏", "一2-4", "本部\n教401"},
	{"0980", "05UG031", "資料科學在教育上的應用 ", "鍾明倫", "四2-4", "本部\n教402"},
	{"0981", "05UG032", "學習分析工具實務應用 ", "廖執善", "二6-8", "本部\n教402"},
	{"0982", "05UG033", "藝術創作與程式設計", "謝諭", "三3-4", "本部\n誠109"},
	{"0983", "05UG034", "網頁程式設計入門", "周遵儒", "三9-10", "本部\n綜1001"},
	{"0984", "05UG035", "數位敘事與學習科技", "陳志洪", "三6-7", "本部\n教401"},
	{"0985", "06UG004", "星星月亮太陽－天文漫談 ", "橋本康弘", "三3-4", "公館\nＥ202"},
	{"0986", "06UG005", "奈米科技", "鄭淳護", "二6-7", "本部\n誠203"},
	{"0987", "06UG005", "奈米科技", "郭金國", "五3-4", "本部\n教201演講廳"},
	{"0988", "06UG006", "環境倫理與永續發展 ", "李清安", "一A-B", "本部\n誠201"},
	{"3449", "06UG006", "環境倫理與永續發展 ", "拉帕契", "二3-4", "公館\nＳ307"},
	{"0989", "06UG007", "生活中的化學", "姜雲生", "四1-2", "本部\n綜210展覽廳"},
	{"5894", "06UG007", "生活中的化學", "宋蕙伶", "二6-7", "本部\n綜10樓教室"},
	{"0990", "06UG012", "資訊科技與生活 ", "劉建成", "一3-4", "本部\n誠104"},
	{"0991", "06UG012", "資訊科技與生活 ", "劉建成", "一6-7", "本部\n誠104"},
	{"0992", "06UG014", "科技與人文的對話", "王銀國", "三3-4", "本部\n誠201"},
	{"0993", "06UG018", "學習與生活中的腦科學 ", "顏妙璇", "一3-4", "公館\n理圖001"},
	{"0994", "06UG018", "學習與生活中的腦科學 ", "顏妙璇", "二3-4", "公館\nＥ101"},
	{"0995", "06UG021", "宇宙中的生命與太空環境 ", "橋本康弘", "", ""},
	{"0996", "06UG030", "聯合國永續發展目標", "張子超", "四6-7", "公館\nＳ303"},
	{"0997", "0AUG203", "音樂概論", "吳佩蓉", "三3-4", "本部\n音308"},
	{"1000", "0AUG402", "心理學 ", "蘇宜芬", "三8-9", "本部\n教103"},
	{"1001", "0AUG402", "心理學 ", "張祐瑄", "一3-4", "本部\n正201"},
	{"1002", "0AUG402", "心理學 ", "楊諮燕", "四3-4", "本部\n正201"},
	{"1003", "0AUG402", "心理學 ", "黃美齡", "二3-4", "本部\n誠401"},
	{"1004", "0AUG402", "心理學 ", "朱采翎", "四6-7", "本部\n誠202"},
	{"1005", "0AUG402", "心理學 ", "朱采翎", "四8-9", "本部\n誠202"},
	{"1006", "0AUG426", "物理與生活", "盧志權", "三3-4", "公館\nＳ304"},
	{"1007", "0AUG435", "歷史與人物", "陳健文", "二6-7", "本部\n誠108"},
	{"1008", "0AUG442", "聖經與人生", "范俊銘", "三8-9", "公館\nＳ501"},
	{"1009", "0AUG462", "個人投資理財", "朱文增", "三3-4", "本部\n綜210展覽廳"},
	{"1010", "0AUG463", "休閒與生活", "李晶", "三3-4", "本部\n美術館601"},
	{"1011", "0HUG011", "學術寫作", "謝秀卉", "二3-4", "本部\n正201"},
	{"1012", "0HUG210", "閩南語概論", "林香薇", "五3-4", "本部\n教103"},
	{"1013", "0HUG223", "禪與人生", "黃敬家", "五6-7", "本部\n樸401"},
	{"1014", "0HUG249", "樂府詩與人生", "林淑雲", "一6-7", "本部\n誠101"},
	{"1015", "0HUG432", "莎士比亞的故事", "狄亞倫", "二6-7", "本部\n誠202"},
	{"1016", "0HUG501", "日語（一）", "張怡倩", "四3-4", "公館\n理圖003"},
	{"1017", "0HUG501", "日語（一）", "張怡倩", "四6-7", "公館\n理圖003"},
	{"1018", "0HUG501", "日語（一）", "許菁娟", "一3-4", "本部\n誠109"},
	{"1019", "0HUG501", "日語（一）", "許菁娟", "一6-7", "本部\n正201"},
	{"1020", "0HUG502", "法語（一） ", "楊慧娟", "三6-7", "本部\n正301"},
	{"1021", "0HUG502", "法語（一） ", "楊慧娟", "三8-9", "本部\n正301"},
	{"1022", "0HUG503", "西班牙語（一） ", "張乃翠", "三3-4", "本部\n正203"},
	{"1023", "0HUG503", "西班牙語（一） ", "張乃翠", "三6-7", "本部\n正203"},
	{"1024", "0HUG506", "韓語（一）", "賀中慧", "一6-7", "公館\n理圖002"},
	{"1025", "0HUG506", "韓語（一）", "賀中慧", "一8-9", "公館\n理圖002"},
	{"1026", "0HUG509", "泰語（一） ", "譚華德", "四6-7", "本部\n樸306"},
	{"1027", "0HUG509", "泰語（一） ", "譚華德", "四8-9", "本部\n樸306"},
	{"1028", "0HUG649", "越南語（一） ", "劉陳石草", "三A-B", "本部\n誠109"},
	{"1029", "0HUG649", "越南語（一） ", "劉陳石草", "四A-B", "本部\n誠206"},
	{"1030", "0HUG652", "日語（二）", "張怡倩", "四8-9", "公館\n理圖003"},
	{"1031", "0HUG652", "日語（二）", "許菁娟", "一8-9", "本部\n正201"},
	{"1032", "0HUG655", "馬來語（一）", "陳際宇", "一A-B", "本部\n誠205"},
	{"1033", "0HUG655", "馬來語（一）", "陳際宇", "二A-B", "本部\n誠205"},
	{"1034", "0NUG238", "動物與人生", "李明忠", "三3-4", "公館\nＢ101"},
	{"1035", "0NUG245", "人類與自然資源", "葉孟宛", "三3-4", "公館\nＳ201"},
	{"1036", "0NUG246", "生理心理學 ", "呂國棟", "三3-4", "公館\nＥ102"},
	{"1037", "0NUG253", "３Ｄ列印與創新應用", "沈林琥", "三3-4", "本部\n教103"},
	{"1038", "0SUG514", "博物館探索", "黃玫瑄", "三6-7", "本部\n誠301"},
	{"1039", "0SUG523", "通識教育講座", "鄭怡庭", "三3-4", "本部\n誠301"},
	{"1040", "0SUG526", "大學入門", "陳伊琳", "三3-4", "本部\n教9樓綜合教室"},
	{"1041", "0SUG526", "大學入門", "陳柏邑", "三3-4", "本部\n教513"},
	{"1042", "0SUG526", "大學入門", "徐敏雄", "二8-9", "本部\n誠203"},
	{"1043", "0SUG526", "大學入門", "張家臻", "三3-4", "本部\n誠207"},
	{"1044", "0SUG526", "大學入門", "吳志文", "三3-4", "本部\n樸401"},
	{"1045", "0SUG526", "大學入門", "周有熙", "三3-4", "本部\n綜202演講廳"},
	{"1046", "0SUG526", "大學入門", "廖靜宜", "三3-4", "本部\n特112"},
	{"1047", "0SUG526", "大學入門", "林宗進", "三3-4", "本部\n誠204"},
	{"1048", "0SUG526", "大學入門", "胡翠君", "四3-4", "本部\n誠106"},
	{"1049", "0SUG531", "大學入門", "徐國能", "三3-4", "本部\n教201演講廳"},
	{"1050", "0SUG531", "大學入門", "江璧羽", "三3-4", "本部\n誠307"},
	{"1051", "0SUG531", "大學入門", "陳昭揚", "四3-4", "本部\n誠208"},
	{"1052", "0SUG531", "大學入門", "李宗祐", "三3-4", "本部\n誠107"},
	{"1053", "0SUG531", "大學入門", "洪立三", "三3-4", "本部\n誠302"},
	{"1054", "0SUG531", "大學入門", "劉承賢", "三3-4", "本部\n誠206"},
	{"1055", "0SUG532", "大學入門", "朱亮儒\n許志農", "三3-4", "公館\nＢ103"},
	{"1056", "0SUG532", "大學入門", "傅祖怡", "三3-4", "公館\nＳ303"},
	{"1057", "0SUG532", "大學入門", "呂家榮\n吳學亮\n陳頌方", "三3-4", "公館\nＥ302"},
	{"1058", "0SUG532", "大學入門", "賴韻如", "三3-4", "公館\nＥ201"},
	{"1059", "0SUG532", "大學入門", "曾瑋玲", "三3-4", "公館\nＳ404"},
	{"1060", "0SUG532", "大學入門", "蔣宗哲", "二6-7", "公館\nＢ102"},
	{"1061", "0SUG533", "大學入門", "程金保", "五3-4", "本部\n工501(TA501)"},
	{"1062", "0SUG534", "大學入門", "潘正宸", "二1-2", "本部\n正104"},
	{"1063", "0SUG534", "大學入門", "林靜萍", "二1-2", "本部\n正105"},
	{"1064", "0SUG534", "大學入門", "周台英", "四1-2", "公館\n體育館二3樓會議室"},
	{"1065", "0SUG535", "大學入門", "孫翼華", "四6-7", "本部\n美101"},
	{"1066", "0SUG535", "大學入門", "黃進龍", "四6-7", "本部\n美402"},
	{"1067", "0SUG535", "大學入門", "薛佑廷", "五8-9", "本部\n設計系2A教室"},
	{"1068", "0SUG536", "大學入門", "吳佩蓉", "二3-4", "本部\n綜202演講廳"},
	{"1069", "0SUG536", "大學入門", "林娟儀", "二6-7", "本部\n樸306"},
	{"1070", "0SUG537", "大學入門", "邱皓政", "三3-4", "本部\n誠402"},
	{"1071", "0SUG538", "大學入門", "林昌平", "三9-10", "本部\n誠202"},
	{"1072", "0SUG538", "大學入門", "王慧娟", "一8-9", "本部\n正202"},
	{"1073", "0SUG538", "大學入門", "吳龍雲", "一6-7", "本部\n誠208"},
	{"3299", "0SUG539", "自主學習：專題探究", "謝宜蓉", "二A-B", "本部\n誠108"},
}
