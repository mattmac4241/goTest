package main

import (
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/Jeffail/gabs"
	_ "github.com/lib/pq" //Needed for postgres

	"github.com/jinzhu/gorm"
)

//Asset struct
type Asset struct {
	gorm.Model
	ModelNumber     string `json:"model_number"`
	ProductTypeName string `json:"product_type_name"`
	Quantity        int    `json:"quantity"`
	ProfilePicture  string `json:"profile_picture"`
	RetailerName    string `json:"retailer_name"`
}

func main() {

	body := `{"after_claim_rooms":[],"auto_claim_damage_picture_notes":[],"auto_claim_damages":[],"auto_claim_exterior_picture_notes":[],"auto_claim_exteriors":[],"auto_claim_interior_picture_notes":[],"auto_claim_interiors":[],"auto_claim_item_picture_notes":[],"auto_claim_pdf_report_requests":[],"auto_claim_pdf_reports":[],"auto_inspection_exterior_picture_notes":[],"auto_inspection_exteriors":[],"auto_inspection_feature_picture_notes":[],"auto_inspection_features":[],"auto_inspection_interior_picture_notes":[],"auto_inspection_interiors":[],"auto_inspection_item_picture_notes":[],"claim_pdf_report_requests":[],"claim_pdf_reports":[],"claim_picture_documents":[],"claim_picture_notes":[],"claim_room_picture_notes":[],"claim_sketch_picture_nodes":[],"claim_sketches":[],"compressed_key":"19d225ac-0bcb-4921-a544-7664ab051b74/compressed","created":1.469634096147e+12,"creator":46141,"creator_email":"lindsay.gaskins@gmail.com","full_uri":"https://tor01.objectstorage.softlayer.net/v1/AUTH_291c4033-af47-446c-aeba-5dbbb1dad45a/Pictures/0d114595-d0e3-4024-ac4b-c5343b60f47a?temp_url_expires=1475798400&temp_url_sig=a2acbf875ed45285233b5789cc9dc477a6d4de28","height":1080,"id":2.956488e+06,"key":"0d114595-d0e3-4024-ac4b-c5343b60f47a","picture_notes":[],"profile_assets":[{"id":568095}],"profile_auto_claim_items":[],"profile_auto_inspection_items":[],"profile_claim_assets":[],"property_inspection_pdf_report_requests":[],"property_inspection_pdf_reports":[],"receipt_assets":[],"receipt_auto_claim_items":[],"receipt_auto_inspection_items":[],"receipt_claim_assets":[],"rooms":[],"server_permissions":{"delete":true,"read":{"asset":true,"compressed_key":true,"created":true,"creator":true,"height":true,"id":true,"key":true,"tap_x":true,"tap_y":true,"thumbnails":true,"width":true},"update":{"asset":true,"thumbnails":true}},"tag_assets":[],"tag_auto_claim_items":[],"tag_auto_inspection_items":[],"tag_claim_assets":[],"tap_x":null,"tap_y":null,"uri":"https://tor01.objectstorage.softlayer.net/v1/AUTH_291c4033-af47-446c-aeba-5dbbb1dad45a/Pictures/19d225ac-0bcb-4921-a544-7664ab051b74/compressed?temp_url_sig=0b9fcf45326822ce35d2ca008e94c86a89e8c92a\u0026temp_url_expires=1475798400","user":46141,"videos":[],"width":1920}`
	jsonParsed, _ := gabs.ParseJSON([]byte(body))
	uri := jsonParsed.Path("full_uri").String()
	uri, _ = strconv.Unquote(uri)
	out, _ := os.Create("test.jpg")
	defer out.Close()
	resp, _ := http.Get(uri)
	defer resp.Body.Close()
	io.Copy(out, resp.Body)
}
