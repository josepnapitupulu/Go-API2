package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	
	"github.com/gorilla/mux"
	"github.com/josepnapitupulu/API_Create_Nikah/models"
	"github.com/josepnapitupulu/API_Create_Nikah/utils"
)

var NewNikah models.Nikah
// var tmpl * template.Template
// func init(){
// 	tmpl = template.Must(template.ParseFiles("jemaat.html"))
// }

// func Index(w http.ResponseWriter, r *http.Request){
// 	temp, _ := template.ParseFiles("views/jemaat.html")
// 	temp.Execute(w, nil)
// }

// func Create(w http.ResponseWriter, r *http.Request){
// 	temp, _ := template.ParseFiles("views/createJemaat.html")
// 	temp.Execute(w, nil)
// }

func GetNikah(w http.ResponseWriter, r *http.Request) {
	newNikahs := models.GetAllNikahs()
	res, _ := json.Marshal(newNikahs)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetNikahById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nikahId := vars["nikahId"]
	ID, err := strconv.ParseInt(nikahId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	nikahDetails, _ := models.GetNikahbyId(ID)
	res, _ := json.Marshal(nikahDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateNikah(w http.ResponseWriter, r *http.Request) {
	CreateNikah := &models.Nikah{}
	utils.ParseBody(r, CreateNikah)
	b := CreateNikah.CreateNikah()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteNikah(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nikahId := vars["nikahId"]
	ID, err := strconv.ParseInt(nikahId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	nikah := models.DeleteNikah(ID)
	res, _ := json.Marshal(nikah)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateNikah(w http.ResponseWriter, r *http.Request) {
    var updateNikah = &models.Nikah{}
    utils.ParseBody(r, updateNikah)
    vars := mux.Vars(r)
    nikahId := vars["nikahId"]
    ID, err := strconv.ParseInt(nikahId, 0, 0)
    if err != nil {
        fmt.Println("error while parsing")
    }
    nikahDetails, db := models.GetNikahbyId(ID)
    if updateNikah.NamaMempelaiLaki != "" {
        nikahDetails.NamaMempelaiLaki = updateNikah.NamaMempelaiLaki
    }
    if updateNikah.AlamatMempelaiLaki != "" {
        nikahDetails.AlamatMempelaiLaki = updateNikah.AlamatMempelaiLaki
    }
    if updateNikah.TanggalPernikahan != "" {
        nikahDetails.TanggalPernikahan = updateNikah.TanggalPernikahan
    }
    if updateNikah.NamaAyahMempelaiLaki != "" {
        nikahDetails.NamaAyahMempelaiLaki = updateNikah.NamaAyahMempelaiLaki
    }
    if updateNikah.NamaIbuMempelaiLaki != "" {
        nikahDetails.NamaIbuMempelaiLaki = updateNikah.NamaIbuMempelaiLaki
    }
	if updateNikah.NamaLengkapMempelaiPerempuan != "" {
        nikahDetails.NamaLengkapMempelaiPerempuan = updateNikah.NamaLengkapMempelaiPerempuan
    }
	if updateNikah.AlamatMempelaiPerempuan != "" {
        nikahDetails.AlamatMempelaiPerempuan = updateNikah.AlamatMempelaiPerempuan
    }
	if updateNikah.NamaAyahMempelaiPerempuan != "" {
        nikahDetails.NamaAyahMempelaiPerempuan = updateNikah.NamaAyahMempelaiPerempuan
    }
	if updateNikah.NamaIbuMempelaiPerempuan != "" {
        nikahDetails.NamaIbuMempelaiPerempuan = updateNikah.NamaIbuMempelaiPerempuan
    }
    db.Save(&nikahDetails)
    res, _ := json.Marshal(nikahDetails)
    w.Header().Set("Content-Type", "pkglication/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}