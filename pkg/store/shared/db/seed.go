package postgres

import (
	"fmt"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
)

func Seeder(db *gorm.DB) error {
	log.Printf("Seeding started")

	//// Brands Seeding
	//var brands []*dbModels.Brand
	//brands = []*dbModels.Brand{
	//	{Name: "Samsung", IsActive: true},
	//	{Name: "Xiomi", IsActive: true},
	//}
	//for _, brand := range brands {
	//	if err := db.Create(&brand).Error; err != nil {
	//		return err
	//	}
	//}

	// Technical Services Seeding
	//response, err := http.Get("http://localhost:8080/api/v1/technical-services/ping")
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
	fmt.Println("hellllllo")

	//var todo Todo
	//json.Unmarshal(bodyBytes, &todo) //bu method ile json objesini  vartodo objesine aktardık
	////json to &tdo object data convertion
	//fmt.Println(todo)

	// DeviceTypes Seeding
	//deviceTypeResults, _ := utils.ReadToJson("pkg/store/shared/db/seeds/device_types.json")
	////println(deviceTypeResults["Device_Types"])
	//res := deviceTypeResults["Device_Types"].map[string]interface{}
	//println(res)
	//var deviceTypes []*dbModels.DeviceType
	//deviceTypes = []*dbModels.DeviceType{
	//	{Name: deviceTypeResults, IsActive: true},
	//	{Name: "Xiomi", IsActive: true},
	//}

	//for _, data := range res {
	//	for _, v := range data.(map[string]interface{}) {
	//		switch t := v.(type) {
	//		case string, []int:
	//			fmt.Println(t)
	//		default:
	//			fmt.Println("wrong type")
	//		}
	//	}
	//}

	//birds := deviceTypeResults["Device_Types"].(map[string]interface{})
	//println(birds)
	//for key, value := range res {
	//	// Each value is an interface{} type, that is type asserted as a string
	//	fmt.Println(key, value.(string))
	//}

	//for _, deviceType := range deviceTypeResults {
	//	println(deviceType)
	//	if err := db.Create(&deviceType).Error; err != nil {
	//		return err
	//	}
	//}

	log.Printf("Seeding done")
	return error(nil)
}
