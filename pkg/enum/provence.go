package enum

type Provence string

const (
	AnHui Provence = "安徽"
	// Aomen       = "澳门"
	BeiJing      Provence = "北京"
	ChongQing    Provence = "重庆"
	FuJian       Provence = "福建"
	Gansu        Provence = "甘肃"
	GuangDong    Provence = "广东"
	GuangXi      Provence = "广西"
	GuiZhou      Provence = "贵州"
	HaiNan       Provence = "海南"
	HeiLongJiang Provence = "黑龙江"
	HeBei        Provence = "河北"
	HeNan        Provence = "河南"
	HuBei        Provence = "湖北"
	HuNan        Provence = "湖南"
	JiangSu      Provence = "江苏"
	JiangXi      Provence = "江西"
	JiLin        Provence = "吉林"
	LiaoNing     Provence = "辽宁"
	NeiMenggu    Provence = "内蒙古"
	NingXia      Provence = "宁夏"
	QingHai      Provence = "青海"
	ShanDong     Provence = "山东"
	ShanXi       Provence = "山西"
	ShaanXi      Provence = "陕西"
	ShangHai     Provence = "上海"
	SiChuan      Provence = "四川"
	//TaiWan      = "台湾"
	TianJin Provence = "天津"
	//XiangGang   = "香港"
	XinJiang Provence = "新疆"
	YunNan   Provence = "云南"
	ZheJiang Provence = "浙江"
)

func (obj Provence) Values() []string {
	return []string{
		string(AnHui),
		string(BeiJing),
		string(ChongQing),
		string(FuJian),
		string(Gansu),
		string(GuangDong),
		string(GuangXi),
		string(GuiZhou),
		string(HaiNan),
		string(HeiLongJiang),
		string(HeBei),
		string(HeNan),
		string(HuBei),
		string(HuNan),
		string(JiangSu),
		string(JiangXi),
		string(JiLin),
		string(LiaoNing),
		string(NeiMenggu),
		string(NingXia),
		string(QingHai),
		string(ShanDong),
		string(ShanXi),
		string(ShaanXi),
		string(ShangHai),
		string(SiChuan),
		string(TianJin),
		string(XinJiang),
		string(YunNan),
		string(ZheJiang),
	}
}
func (obj Provence) Ptr() *Provence {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
