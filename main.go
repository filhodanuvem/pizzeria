package main

import (
	"github.com/cloudson/pizzeria/graph"
	"github.com/spf13/viper"

	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"crypto/sha256"
	"encoding/hex"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Error on read config file: %s \n", err))
	}
	address := "127.0.0.1"
	port := "8080"
	http.HandleFunc("/pie", printPieGraph)
	http.HandleFunc("/bar", printBarGraph)
	http.HandleFunc("/line", printLineGraph)

	log.Print("Pizzeria running on " + address + ":" + port)
	err = http.ListenAndServe(address+":"+port, nil)
	log.Print(err.Error())
}

func getHeightAndHeight(u *url.URL) (int, int, error) {
	query := u.Query()
	height, err := strconv.Atoi(query.Get("h"))
	if err != nil {
		return 0, 0, fmt.Errorf("height required as integer")
	}
	width, err := strconv.Atoi(query.Get("w"))
	if err != nil {
		return 0, 0, fmt.Errorf("Width required as integer")
	}

	return height, width, nil
}

func getDataAndLabels(u *url.URL) ([]float64, []string, error) {
	query := u.Query()
	data := query.Get("dt")
	valuesString := strings.Split(data, ",")
	numberOfData := len(valuesString)
	values := make([]float64, numberOfData)
	var err error
	for i := 0; i < len(valuesString); i++ {
		values[i], err = strconv.ParseFloat(valuesString[i], 64)
		if err != nil {
			return []float64{}, []string{}, fmt.Errorf("Data is required as integer values %s given", valuesString[i])
		}
	}

	labels := getLabels(u)
	numberOfLabels := len(labels)
	if numberOfLabels < numberOfData {
		return []float64{}, []string{}, fmt.Errorf("%d labels expected %d given", numberOfData, numberOfLabels)
	}

	return values, labels, nil
}

func getLabels(u *url.URL) []string {
	query := u.Query()
	labelsConcat := query.Get("lb")
	if labelsConcat == "" {
		return []string{}
	}
	labels := strings.Split(labelsConcat, ",")

	return labels
}

func getDataXY(u *url.URL) ([]float64, []float64, error) {
	query := u.Query()
	dataX := query.Get("dtx")
	valuesXString := strings.Split(dataX, ",")
	numberOfDataX := len(valuesXString)
	valuesX := make([]float64, numberOfDataX)
	var err error
	for i := 0; i < len(valuesXString); i++ {
		valuesX[i], err = strconv.ParseFloat(valuesXString[i], 64)
		if err != nil {
			return []float64{}, []float64{}, fmt.Errorf("Data X is required as integer values %s given", valuesXString[i])
		}
	}

	dataY := query.Get("dty")
	valuesYString := strings.Split(dataY, ",")
	numberOfDataY := len(valuesYString)
	valuesY := make([]float64, numberOfDataY)
	for i := 0; i < len(valuesYString); i++ {
		valuesY[i], err = strconv.ParseFloat(valuesYString[i], 64)
		if err != nil {
			return []float64{}, []float64{}, fmt.Errorf("Data Y is required as integer values %s given", valuesYString[i])
		}
	}

	return valuesX, valuesY, nil
}

func getColors(u *url.URL) ([]string, error) {
	query := u.Query()
	colorsString := query.Get("cl")
	if colorsString == "" {
		return []string{}, nil
	}

	colors := strings.Split(colorsString, ",")
	regex, _ := regexp.Compile("(^[a-fA-F0-9]{6}$)|(^[a-fA-F0-9]{3}$)")
	for i := 0; i < len(colors); i++ {
		if regex.MatchString(colors[i]) == false {
			return []string{}, fmt.Errorf("String %s is not a valid color, use the format, expected [A-F0-9]{6}", colors[i])
		}
	}

	return colors, nil
}

func checksum(u *url.URL, v *viper.Viper) bool {
	if !v.GetBool("enabled") {
		return true
	}
	secret := v.GetString("secret")
	queryString := u.String()
	regex, _ := regexp.Compile("&?hash=[a-f0-9]+")
	queryString = secret + string(regex.ReplaceAll([]byte(queryString), []byte("")))
	checksum := u.Query().Get("hash")

	hasher := sha256.New()
	hasher.Write([]byte(queryString))

	return checksum == hex.EncodeToString(hasher.Sum(nil))
}

/**
* @example http://localhost:8080/pie?h=200&w=200&dt=1,2,3&lb=cash,credit,debit
 */
func printPieGraph(w http.ResponseWriter, r *http.Request) {
	if !checksum(r.URL, viper.Sub("checksum")) {
		http.Error(w, "Hash is wrong", 404)
		return
	}
	width, height, err := getHeightAndHeight(r.URL)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	values, labels, err := getDataAndLabels(r.URL)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	colors, err := getColors(r.URL)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if len(colors) > 0 && len(colors) < len(values) {
		http.Error(w, fmt.Sprintf("Expected %d colors, found %d", len(values), len(colors)), 400)
		return
	}

	g := graph.NewPieGraph(height, width, values, labels)
	g.Colors = colors
	service := new(graph.PieService)
	w.Header().Set("Content-Type", "image/png")
	service.Build(g, w)
}

/**
* @example http://localhost:8080/bar?h=200&w=200&dt=1,2,3&lb=cash,credit,debit
 */
func printBarGraph(w http.ResponseWriter, r *http.Request) {
	if !checksum(r.URL, viper.Sub("checksum")) {
		http.Error(w, "Hash is wrong", 404)
		return
	}
	width, height, err := getHeightAndHeight(r.URL)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	values, labels, err := getDataAndLabels(r.URL)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	colors, err := getColors(r.URL)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if len(colors) > 0 && len(colors) < len(values) {
		http.Error(w, fmt.Sprintf("Expected %d colors, found %d", len(values), len(colors)), 400)
	}

	g := graph.NewBarGraph(width, height, values, labels)
	g.Colors = colors
	service := new(graph.BarService)
	w.Header().Set("Content-Type", "image/png")
	service.Build(g, w)
}

/**
* @example http://localhost:8080/line?h=200&w=200&dtx=1,2,3&dty=2,4,400
 */
func printLineGraph(w http.ResponseWriter, r *http.Request) {
	if !checksum(r.URL, viper.Sub("checksum")) {
		http.Error(w, "Hash is wrong", 404)
		return
	}
	width, height, err := getHeightAndHeight(r.URL)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	valuesX, valuesY, err := getDataXY(r.URL)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	labels := getLabels(r.URL)
	if len(labels) < len(valuesY) {
		http.Error(w, fmt.Sprintf("Expected %d labels, found %d", len(valuesY), len(labels)), 400)
		return
	}

	g := graph.NewLineGraph(width, height, valuesX, valuesY, labels)
	service := new(graph.LineService)
	w.Header().Set("Content-Type", "image/png")
	service.Build(g, w)

}

func printTimeSeriesGraph(w http.ResponseWriter, r *http.Request) {
	width, height, err := getHeightAndHeight(r.URL)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
	values, labels, err := getDataAndLabels(r.URL)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	g := graph.NewTimeSeriesGraph(width, height, values, labels)
	service := new(graph.TimeSeriesService)
	w.Header().Set("Content-Type", "image/png")
	service.Build(g, w)

}
