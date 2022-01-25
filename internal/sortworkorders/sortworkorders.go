package sortworkorders

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sort"

	"github.com/eugenefoxx/SQLPanacimP1/pkg/filereader"
	"github.com/eugenefoxx/SQLPanacimP1/pkg/logging"
)

type WorkOrders struct {
	JobID string `json: "jobid"`
}

func Getclosedworkorders() {
	logger := logging.GetLogger()
	dirWOpath := os.Getenv("dirWO")
	closedWORemovepath := os.Getenv("closedWORemove")
	processedWOpath := os.Getenv("processedWO")

	listWO := [][]string{{"5696"}, {"5697"}, {"5699"}}

	dirWO := dirWOpath
	//mode := 0755
	if _, err := os.Stat(dirWO); os.IsNotExist(err) {
		os.Mkdir(dirWO, 0755)
	}
	closedWORemove := closedWORemovepath

	if fileExists(closedWORemove) {
		os.Remove(closedWORemove)
	}

	closedWO := closedWORemovepath
	if _, err := os.Stat(closedWO); os.IsNotExist(err) {
		clwo, err := os.Create(closedWO)
		if err != nil {
			log.Println(err)
		}
		defer clwo.Close()

		writer := csv.NewWriter(clwo)
		writer.Write([]string{"0"})
		writer.Comma = ','
		writer.Flush()
	}

	//clwo := filereader.Readfile(closedWO)
	splitWO, err := os.OpenFile(closedWO, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		logger.Errorf(err.Error())
		return
	}
	defer splitWO.Close()

	for _, i := range listWO {

		var result = []string{i[0]}
		for _, v := range result {
			_, err = fmt.Fprintln(splitWO, v)
			if err != nil {
				splitWO.Close()
				return
			}
		}
	}
	// из файла забираем построчно номера job_id
	readclwo := filereader.Readfile(closedWO)
	fmt.Println(readclwo[1][0])
	fmt.Println(readclwo[2][0])
	str1 := readclwo[1][0]
	str2 := readclwo[2][0]
	str3 := readclwo[3][0]

	processedWO := processedWOpath

	if _, err := os.Stat(processedWO); os.IsNotExist(err) {
		clwo, err := os.Create(processedWO)
		if err != nil {
			log.Println(err)
		}
		defer clwo.Close()

		writer := csv.NewWriter(clwo)
		writer.Write([]string{"WO"})
		writer.Comma = ','
		writer.Flush()
	}

	//readPrWO := filereader.Readfileseeker(processedWO)

	//readPrWO := filereader.Readfile(processedWO)
	split, err := os.OpenFile(processedWO, os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		logger.Errorf(err.Error())
		return
	}
	defer split.Close()

	csvFile, err := os.Open(processedWO)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.LazyQuotes = true
	reader.Comma = ';'
	if err != nil {
		//	fmt.Println("Ошибка", err)
		fmt.Println(err)
		os.Exit(1)
	}
	defer csvFile.Close()

	var arrWO []WorkOrders

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		arrWO = append(arrWO, WorkOrders{
			JobID: line[0],
		})
	}
	arrwoJSON, _ := json.Marshal(arrWO)

	err = json.Unmarshal([]byte(arrwoJSON), &arrWO)
	crarr := arrWO
	sort.Slice(crarr, func(i, j int) bool {
		return crarr[i].JobID <= crarr[j].JobID
	})

	needlearr := str1
	idx := sort.Search(len(crarr), func(i int) bool {
		return string(crarr[i].JobID) >= needlearr
	})

	if crarr[idx].JobID == needlearr {
		//	fmt.Println("Found:", idx, crarr[idx])
	} else {
		//fmt.Println("Found noting: ", idx)
		var result = []string{str1}
		for _, v := range result {
			_, err = fmt.Fprintln(split, v)
			if err != nil {
				split.Close()
				return
			}
		}
	}
	needlearr = str2
	idx = sort.Search(len(crarr), func(i int) bool {
		return string(crarr[i].JobID) >= needlearr
	})

	if crarr[idx].JobID == needlearr {
		//	fmt.Println("Found:", idx, crarr[idx])
	} else {
		//fmt.Println("Found noting: ", idx)
		var result = []string{str2}
		for _, v := range result {
			_, err = fmt.Fprintln(split, v)
			if err != nil {
				split.Close()
				return
			}
		}
	}

	needlearr = str3
	idx = sort.Search(len(crarr), func(i int) bool {
		return string(crarr[i].JobID) >= needlearr
	})

	if crarr[idx].JobID == needlearr {
		//	fmt.Println("Found:", idx, crarr[idx])
	} else {
		//fmt.Println("Found noting: ", idx)
		var result = []string{str3}
		for _, v := range result {
			_, err = fmt.Fprintln(split, v)
			if err != nil {
				split.Close()
				return
			}
		}
	}

}

func GetLastJobIdValue1() (string, error) {
	logger := logging.GetLogger()
	lastWOpath := os.Getenv("lastJobId")
	processedWOpath := os.Getenv("processedWO")

	lastclosedWO := lastWOpath
	if _, err := os.Stat(lastclosedWO); os.IsNotExist(err) {
		lastwo, err := os.Create(lastclosedWO)
		if err != nil {
			logger.Errorf(err.Error())
		}
		defer lastwo.Close()

		writer := csv.NewWriter(lastwo)
		writer.Write([]string{"LastWO"})
		writer.Comma = ','
		writer.Flush()
	}

	splitWO, err := os.OpenFile(lastclosedWO, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		logger.Errorf(err.Error())
		//return
	}
	defer splitWO.Close()

	processedWO := processedWOpath
	csvFile, err := os.Open(processedWO)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	//reader.FieldPos()
	reader.LazyQuotes = true
	reader.Comma = ';'
	if err != nil {
		//	fmt.Println("Ошибка", err)
		fmt.Println(err)
		os.Exit(1)
	}
	defer csvFile.Close()

	var arrWO []WorkOrders

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		arrWO = append(arrWO, WorkOrders{
			JobID: line[0],
		})
	}
	arrwoJSON, _ := json.Marshal(arrWO)

	err = json.Unmarshal([]byte(arrwoJSON), &arrWO)
	crarr := arrWO
	sort.Slice(crarr, func(i, j int) bool {
		return crarr[i].JobID <= crarr[j].JobID
	})

	getresult := crarr[len(crarr)-1]
	// убираем заголовк {WO}
	getresult2 := crarr[:len(crarr)-1]
	// 1 последнее значение
	getresult3 := getresult2[len(getresult2)-1]
	// обрезаем еще на один элемент
	//test4 := test2[:len(test2)-1]
	//test5 := test4[len(test4)-1]
	//test := len(crarr) - 1
	fmt.Println("last 1", getresult)
	fmt.Println("last 2", getresult2)
	fmt.Println("last 3", getresult3)
	//fmt.Println("last 5", test5)
	//return string

	lastclosedWOread := lastclosedWO
	csvFilelastclosedWO, err := os.Open(lastclosedWOread)
	readerlastclosedWO := csv.NewReader(bufio.NewReader(csvFilelastclosedWO))
	//reader.FieldPos()
	readerlastclosedWO.LazyQuotes = true
	readerlastclosedWO.Comma = ';'
	if err != nil {
		//	fmt.Println("Ошибка", err)
		fmt.Println(err)
		os.Exit(1)
	}
	defer csvFilelastclosedWO.Close()

	var arrlastclosedWO []WorkOrders
	for {
		line, error := readerlastclosedWO.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		arrlastclosedWO = append(arrlastclosedWO, WorkOrders{
			JobID: line[0],
		})
	}
	arrlastclosedWOJSON, _ := json.Marshal(arrlastclosedWO)

	err = json.Unmarshal([]byte(arrlastclosedWOJSON), &arrlastclosedWO)

	crarrclosewo := arrlastclosedWO
	sort.Slice(crarrclosewo, func(i, j int) bool {
		return crarrclosewo[i].JobID <= crarrclosewo[j].JobID
	})

	//res := strings.Join(test3.JobID, "")
	needclosewo := getresult3.JobID
	idx2 := sort.Search(len(crarrclosewo), func(i int) bool {
		return string(crarrclosewo[i].JobID) >= needclosewo
	})
	var res string
	if crarrclosewo[idx2].JobID == needclosewo {
		//	fmt.Println("Found:", idx, crarr[idx])
		res = ""
		return res, nil
	} else {
		//fmt.Println("Found noting: ", idx)
		var result = []string{getresult3.JobID}
		for _, v := range result {
			_, err = fmt.Fprintln(splitWO, v)
			if err != nil {
				splitWO.Close()
			}
		}
		res = getresult3.JobID
		return res, nil
	}
	fmt.Println("test res1", res)
	return res, nil

}

func GetLastJobIdValue2() (string, error) {
	logger := logging.GetLogger()
	lastWOpath := os.Getenv("lastJobId")
	processedWOpath := os.Getenv("processedWO")

	lastclosedWO := lastWOpath
	if _, err := os.Stat(lastclosedWO); os.IsNotExist(err) {
		lastwo, err := os.Create(lastclosedWO)
		if err != nil {
			logger.Errorf(err.Error())
		}
		defer lastwo.Close()

		writer := csv.NewWriter(lastwo)
		writer.Write([]string{"LastWO"})
		writer.Comma = ','
		writer.Flush()
	}

	splitWO, err := os.OpenFile(lastclosedWO, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		logger.Errorf(err.Error())
		//return
	}
	defer splitWO.Close()

	processedWO := processedWOpath
	csvFile, err := os.Open(processedWO)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	//reader.FieldPos()
	reader.LazyQuotes = true
	reader.Comma = ';'
	if err != nil {
		//	fmt.Println("Ошибка", err)
		fmt.Println(err)
		os.Exit(1)
	}
	defer csvFile.Close()

	var arrWO []WorkOrders

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		arrWO = append(arrWO, WorkOrders{
			JobID: line[0],
		})
	}
	arrwoJSON, _ := json.Marshal(arrWO)

	err = json.Unmarshal([]byte(arrwoJSON), &arrWO)
	crarr := arrWO
	sort.Slice(crarr, func(i, j int) bool {
		return crarr[i].JobID <= crarr[j].JobID
	})

	getresult := crarr[len(crarr)-1]
	// убираем заголовк {WO}
	getresult2 := crarr[:len(crarr)-1]
	// 1 последнее значение
	getresult3 := getresult2[len(getresult2)-1]
	// обрезаем еще на один элемент
	getresult4 := getresult2[:len(getresult2)-1]
	getresult5 := getresult4[len(getresult4)-1]
	//test := len(crarr) - 1
	fmt.Println("last 1", getresult)
	fmt.Println("last 2", getresult2)
	fmt.Println("last 3", getresult3)
	//fmt.Println("last 5", test5)
	//return string

	lastclosedWOread := lastclosedWO
	csvFilelastclosedWO, err := os.Open(lastclosedWOread)
	readerlastclosedWO := csv.NewReader(bufio.NewReader(csvFilelastclosedWO))
	//reader.FieldPos()
	readerlastclosedWO.LazyQuotes = true
	readerlastclosedWO.Comma = ';'
	if err != nil {
		//	fmt.Println("Ошибка", err)
		fmt.Println(err)
		os.Exit(1)
	}
	defer csvFilelastclosedWO.Close()

	var arrlastclosedWO []WorkOrders
	for {
		line, error := readerlastclosedWO.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		arrlastclosedWO = append(arrlastclosedWO, WorkOrders{
			JobID: line[0],
		})
	}
	arrlastclosedWOJSON, _ := json.Marshal(arrlastclosedWO)

	err = json.Unmarshal([]byte(arrlastclosedWOJSON), &arrlastclosedWO)
	crarrclosewo := arrlastclosedWO
	sort.Slice(crarrclosewo, func(i, j int) bool {
		return crarrclosewo[i].JobID <= crarrclosewo[j].JobID
	})

	//res := strings.Join(test3.JobID, "")
	needclosewo := getresult5.JobID
	idx2 := sort.Search(len(crarrclosewo), func(i int) bool {
		return string(crarrclosewo[i].JobID) >= needclosewo
	})
	var res string
	if crarrclosewo[idx2].JobID == needclosewo {
		//	fmt.Println("Found:", idx, crarr[idx])
		res = ""
		return res, nil
	} else {
		//fmt.Println("Found noting: ", idx)
		var result = []string{getresult5.JobID}
		for _, v := range result {
			_, err = fmt.Fprintln(splitWO, v)
			if err != nil {
				splitWO.Close()
				//return test3.JobID, nil
			}
		}
		res = getresult5.JobID
		return res, nil
	}
	fmt.Println("test res2", res)
	return res, nil

}

func GetLastJobIdValue3() (string, error) {
	logger := logging.GetLogger()
	lastWOpath := os.Getenv("lastJobId")
	processedWOpath := os.Getenv("processedWO")

	lastclosedWO := lastWOpath
	if _, err := os.Stat(lastclosedWO); os.IsNotExist(err) {
		lastwo, err := os.Create(lastclosedWO)
		if err != nil {
			logger.Errorf(err.Error())
		}
		defer lastwo.Close()

		writer := csv.NewWriter(lastwo)
		writer.Write([]string{"LastWO"})
		writer.Comma = ','
		writer.Flush()
	}

	splitWO, err := os.OpenFile(lastclosedWO, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		logger.Errorf(err.Error())
		//return
	}
	defer splitWO.Close()

	processedWO := processedWOpath
	csvFile, err := os.Open(processedWO)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	//reader.FieldPos()
	reader.LazyQuotes = true
	reader.Comma = ';'
	if err != nil {
		//	fmt.Println("Ошибка", err)
		fmt.Println(err)
		os.Exit(1)
	}
	defer csvFile.Close()

	var arrWO []WorkOrders

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		arrWO = append(arrWO, WorkOrders{
			JobID: line[0],
		})
	}
	arrwoJSON, _ := json.Marshal(arrWO)

	err = json.Unmarshal([]byte(arrwoJSON), &arrWO)
	crarr := arrWO
	sort.Slice(crarr, func(i, j int) bool {
		return crarr[i].JobID <= crarr[j].JobID
	})

	getresult := crarr[len(crarr)-1]
	// убираем заголовк {WO}
	getresult2 := crarr[:len(crarr)-1]
	// 1 последнее значение
	getresult3 := getresult2[len(getresult2)-1]
	// обрезаем еще на один элемент
	getresult4 := getresult2[:len(getresult2)-1]
	//getresult5 := getresult4[len(getresult4)-1]
	// обрезаем 3-й элемент
	getresult6 := getresult4[:len(getresult4)-1]
	getresult7 := getresult6[len(getresult6)-1]
	//test := len(crarr) - 1
	fmt.Println("last 1", getresult)
	fmt.Println("last 2", getresult2)
	fmt.Println("last 3", getresult3)
	//fmt.Println("last 5", test5)
	//return string

	lastclosedWOread := lastclosedWO
	csvFilelastclosedWO, err := os.Open(lastclosedWOread)
	readerlastclosedWO := csv.NewReader(bufio.NewReader(csvFilelastclosedWO))
	//reader.FieldPos()
	readerlastclosedWO.LazyQuotes = true
	readerlastclosedWO.Comma = ';'
	if err != nil {
		//	fmt.Println("Ошибка", err)
		fmt.Println(err)
		os.Exit(1)
	}
	defer csvFilelastclosedWO.Close()

	var arrlastclosedWO []WorkOrders
	for {
		line, error := readerlastclosedWO.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		arrlastclosedWO = append(arrlastclosedWO, WorkOrders{
			JobID: line[0],
		})
	}
	arrlastclosedWOJSON, _ := json.Marshal(arrlastclosedWO)

	err = json.Unmarshal([]byte(arrlastclosedWOJSON), &arrlastclosedWO)
	crarrclosewo := arrlastclosedWO
	sort.Slice(crarrclosewo, func(i, j int) bool {
		return crarrclosewo[i].JobID <= crarrclosewo[j].JobID
	})

	//res := strings.Join(test3.JobID, "")
	needclosewo := getresult7.JobID
	idx2 := sort.Search(len(crarrclosewo), func(i int) bool {
		return string(crarrclosewo[i].JobID) >= needclosewo
	})
	var res string
	if crarrclosewo[idx2].JobID == needclosewo {
		//	fmt.Println("Found:", idx, crarr[idx])
		res = ""
		return res, nil
	} else {
		//fmt.Println("Found noting: ", idx)
		var result = []string{getresult7.JobID}
		for _, v := range result {
			_, err = fmt.Fprintln(splitWO, v)
			if err != nil {
				splitWO.Close()
				//return test3.JobID, nil

			}
		}
		res = getresult7.JobID
		return res, nil
	}
	fmt.Println("test res2", res)
	return res, nil

}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
