package postgres

import (
	"fmt"
	dbModels "github.com/mustafakocatepe/Tamircity/pkg/models/db"
	"github.com/mustafakocatepe/Tamircity/pkg/utils"
	"gorm.io/gorm"
	"log"
)

func Seeder(db *gorm.DB) error {
	log.Printf("Seeding started")

	// Brand Seeding
	var brands []*dbModels.Brand
	brands = []*dbModels.Brand{
		{Name: "Samsung", IsActive: true},
		{Name: "Xiomi", IsActive: true},
	}
	for _, brand := range brands {
		if err := db.Create(&brand).Error; err != nil {
			return err
		}
	}

	// DeviceType Seeding
	deviceTypeResults, _ := utils.ReadToJson("pkg/store/shared/db/seeds/device_types.json")
	//println(deviceTypeResults["Device_Types"])
	res := deviceTypeResults["Device_Types"].map[string]interface{}
	println(res)
	//var deviceTypes []*dbModels.DeviceType
	//deviceTypes = []*dbModels.DeviceType{
	//	{Name: deviceTypeResults, IsActive: true},
	//	{Name: "Xiomi", IsActive: true},
	//}

	for _, data := range res {
		for _, v := range data.(map[string]interface{}) {
			switch t := v.(type) {
			case string, []int:
				fmt.Println(t)
			default:
				fmt.Println("wrong type")
			}
		}
	}

	//birds := deviceTypeResults["Device_Types"].(map[string]interface{})
	//println(birds)
	//for key, value := range res {
	//	// Each value is an interface{} type, that is type asserted as a string
	//	fmt.Println(key, value.(string))
	//}

	for _, deviceType := range deviceTypeResults {
		println(deviceType)
		if err := db.Create(&deviceType).Error; err != nil {
			return err
		}
	}

	log.Printf("Seeding done")
	return error(nil)
}
