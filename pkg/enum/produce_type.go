package enum

type ProduceType string

const (
	Tea                ProduceType = "tea"
	EdibleFungi        ProduceType = "edible-fungi"
	Vegetable          ProduceType = "vegetable"
	Seedlings          ProduceType = "seedlings"
	MedicinalMaterials ProduceType = "medicinal-materials"
	GrainAndOilCrops   ProduceType = "grain-and-oil-crops"
	Fisheries          ProduceType = "fisheries"
	AnimalHusbandry    ProduceType = "animal-husbandry"
)

func (obj ProduceType) Values() []string {
	return []string{
		string(Tea),
		string(EdibleFungi),
		string(Vegetable),
		string(Seedlings),
		string(MedicinalMaterials),
		string(GrainAndOilCrops),
		string(Fisheries),
		string(AnimalHusbandry),
	}
}
func (obj ProduceType) Ptr() *ProduceType {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
