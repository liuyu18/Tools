package unmarshal

// 转换成的对象
type Flower struct {
	PlantId    string `json:"plantId"`
	Name    string `json:"name"`
	Description    string `json:"description"`
	GrowZoneNumber    int `json:"growZoneNumber"`
	WateringInterval    int `json:"wateringInterval"`
	ImageUrl    string `json:"imageUrl"`

 }