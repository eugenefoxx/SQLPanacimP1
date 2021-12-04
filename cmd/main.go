package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	//value uint16 = 3000
	value int = 4708
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	recipe := os.Getenv("recipe")

	//npm := readfileseeker("/home/eugenearch/Code/github.com/eugenefoxx/SQLPanacimP1/csvfolder/NPM_910-00473_A_recipte.csv")
	npm := readfileseeker(recipe)
	report, err := os.Create("/home/eugenearch/Code/github.com/eugenefoxx/SQLPanacimP1/csvfolder/report.csv")
	if err != nil {
		log.Println(err)
	}
	defer report.Close()

	split, err := os.OpenFile("/home/eugenearch/Code/github.com/eugenefoxx/SQLPanacimP1/csvfolder/report.csv", os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		log.Println(err)
		return
	}
	defer split.Close()
	//parts := readfileseeker("/home/eugenearch/Code/github.com/eugenefoxx/SQLPanacimP1/csvfolder/parts.csv")
	//var nmpparts = []string{}
	//var arr = [...]string{}
	//var arrofarr = [...]string{}
	for _, iter := range npm {
		//	for _, part := range par ts {
		//	fmt.Println(part[0])
		qtytotal, err := strconv.Atoi(iter[1])
		if err != nil {
			log.Println(err)
		}
		//	if iter[0] == part[0] {

		//	if iter[1] != "" {
		//	fmt.Println(iter[0], iter[1], uint16(qty)*value)
		qtytotal, err = strconv.Atoi(iter[1])
		if err != nil {
			log.Println(err)
		}
		//var arr = [3]string{iter[0], iter[1], int(uint16(qtytotal) * value)}
		//	fmt.Println(arr[2])

		//var result = []string{iter[0] + "," + iter[1] + "," + strconv.Itoa(int(uint16(qtytotal)*value))}
		var result = []string{iter[0] + "," + iter[1] + "," + strconv.Itoa(int(qtytotal)*value)}
		fmt.Println(result)
		for _, v := range result {
			_, err = fmt.Fprintln(split, v)
			if err != nil {
				split.Close()
				return
			}
		}
		//	for i := range arrofarr {
		//		arrofarr[i] = arr
		//	}
		//	}
		//	}
		//fmt.Println(arrofarr, "\n")

		/*	var arrayofslice [len(arrofarr)][]int
			for i := range arrofarr { // assign
				arrayofslice[i] = arrofarr[i][:]
			}
			//	fmt.Println(arrayofslice[0][2], "\n")
			var sliceofslices [][]int
			sliceofslices = arrayofslice[:]
			//sliceofslices[:][:] = [][]int{0, 0}
			fmt.Println(sliceofslices[4][2], "\n")*/
		/*
			parts := readfileseeker("/home/eugenearch/Code/github.com/eugenefoxx/SQLPanacimP1/csvfolder/parts.csv")
			for _, part := range parts {
				intSAP, err := strconv.Atoi(part[0])
				if err != nil {
					log.Println(err)
				}
				intPartSAP1, err := strconv.Atoi(part[1])
				if err != nil {
					log.Println(err)
				}
				intPartSAP2, err := strconv.Atoi(part[1])
				if err != nil {
					log.Println(err)
				}
				intPartSAP3, err := strconv.Atoi(part[1])
				if err != nil {
					log.Println(err)
				}
				//rrr = sliceofslices[0][2]
				if sliceofslices[0][0] == intSAP {
					sliceofslices = append(sliceofslices, part[1], intPartSAP2)
					//	var arr201 = [3]string{part[1], " ", " "}
					//	var arr202 = [3]string{part[1], " ", " "}
					//	var arr203 = [3]string{part[1], " ", " "}
					//arrofarr2
					for ii := range arrofarr {

						arrofarr[ii] = append() //arr201
						arrofarr[ii] = arr202
						arrofarr[ii] = arr203
					}
				}
			}
			fmt.Println(arrofarr)*/
		//	parts := readfileseeker("/home/eugenearch/Code/github.com/eugenefoxx/SQLPanacimP1/csvfolder/parts.csv")
		//	for _, part := range parts {
		//		if arrofarr[0][1]
		//	}
		//nmpparts = append(nmpparts, iter...)
		//	}
	}
	//fmt.Println(nmpparts[0], nmpparts[1])
	reportDGS := readfile("/home/eugenearch/Code/github.com/eugenefoxx/SQLPanacimP1/csvfolder/report.csv")
	reportParts := readfileseeker("/home/eugenearch/Code/github.com/eugenefoxx/SQLPanacimP1/csvfolder/parts.csv")
	panacimdata := readfileseeker("/home/eugenearch/Code/github.com/eugenefoxx/SQLPanacimP1/csvfolder/panacim.csv")

	for p := 0; p < len(reportDGS); p++ {
		parseParts(reportParts, reportDGS, panacimdata, reportDGS[p][0])
	}
	for p := 0; p < len(reportDGS); p++ {
		insertPanacimDataQty(panacimdata, reportDGS[p][0])
	}
	//  формируем файлы с подсчетом Итого установленных компонентов оригинал + замена
	for p := 0; p < len(reportDGS); p++ {
		insertPanacimDataQtyTotal(reportDGS[p][0])
	}
	reportSum, err := os.Create("/home/eugenearch/Code/github.com/eugenefoxx/SQLPanacimP1/csvfolder/reportSumComponent.csv")
	if err != nil {
		log.Println(err)
	}
	defer reportSum.Close()
	reportSumRead := readfile("/home/eugenearch/Code/github.com/eugenefoxx/SQLPanacimP1/csvfolder/reportSumComponent.csv")
	for r := 0; r < len(reportDGS); r++ {
		sumComponent(reportDGS, reportSumRead, reportDGS[r][0])
	}

	//var i int
	/*
			Стопосто, [01.12.2021 13:03]
		Загнать в мапу и проверить длинну мапа с массивом

		Viacheslav Poturaev, [01.12.2021 13:03]
		либо отсортировать массивы и пробежать соседей

		map[int]string использовать. В качестве ключа - индекс в строке
	*/

	reportSumRead2 := readfile("/home/eugenearch/Code/github.com/eugenefoxx/SQLPanacimP1/csvfolder/reportSumComponent.csv")
	//  строка 4, колонка 3
	fmt.Println(reportSumRead2[3][2])
	/* output each array element's value */
	//sum := 0

	//result := []string{}
	for i := 0; i < len(reportSumRead2); i++ {
		//for i, v := range reportSumRead2 {
		//	for j := 0; j < len(reportSumRead2); j++ {
		//	if reportSumRead2[i][0] == reportSumRead2[j][0] {
		//fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		//	fmt.Printf("a[%v] [%v] = %v\n", i, j, reportSumRead[i][j])

		fmt.Println("read", reportSumRead2[i])

		//	f, err := strconv.Atoi(reportSumRead2[i][2])
		//	if err != nil {
		//		log.Println(err)
		//	}
		//	sum += (f)

		//	}
		//	}
	}
	//fmt.Println("Total", sum)
	testFile := readfile("/home/eugenearch/Code/github.com/eugenefoxx/SQLPanacimP1/csvfolder/reportSumComponent.csv")

	for i := 0; i < len(testFile); i++ {
		total1, err := strconv.Atoi(testFile[i][2])
		if err != nil {
			log.Println(err)
		}
		total2, err := strconv.Atoi(testFile[i][3])
		if err != nil {
			log.Println(err)
		}
		if total2-total1 != 0 {
			fmt.Printf("read delta Testfile %s %d\n", testFile[i][0], total2-total1)
		}

	}

	//reportSumRead2 := readfile("/home/eugenearch/Code/github.com/eugenefoxx/SQLPanacimP1/csvfolder/reportSumComponent.csv")
	//fmt.Println("ddd", reportSumRead2)
	/*	for g := 0; g < len(reportSumRead2); g++ {
			for f := 0; f < len(reportSumRead2); f++ {
				if reportSumRead2[g][0] == reportSumRead2[f][0] {
					//fmt.Println("reportSumRead[yy][0]", reportSumRead[g][0])
					fmt.Println("Test duble")
					fmt.Println(reportSumRead2[g][1])
				}
			}
		}
	*/
	/*
		for i := 0; i < len(reportParts); i++ {
			for ii := 0; ii < len(reportDGS); ii++ {
				if reportParts[i][0] == "1013770" {
					fmt.Println("1013770", reportParts[i][1])
					break
				}
			}
		}
	*/
	/*
		reportCommon, err := os.Create("/home/eugenearch/Code/github.com/eugenefoxx/SQLPanacimP1/csvfolder/common.csv")
		if err != nil {
			log.Println(err)
		}
		defer reportCommon.Close()*/
	/*writer := csv.NewWriter(reportCommon)
	writer.Write([]string{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0"})
	writer.Comma = ','
	writer.Flush()*/
	/*
		splitCommon, err := os.OpenFile("/home/eugenearch/Code/github.com/eugenefoxx/SQLPanacimP1/csvfolder/common.csv", os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
			return
		}
		defer splitCommon.Close()
		// strings.Split(raw,",")
		for _, dgs := range reportDGS {
			ttt := readfile("/home/eugenearch/Code/github.com/eugenefoxx/SQLPanacimP1/csvfolder/" + dgs[0] + ".csv")

			//rr := fmt.Sprint(fmt.Printf("Read %s, %s, %s, %v \n", dgs[0], dgs[1], dgs[2], ttt))
			rr := fmt.Sprint(dgs[0]+","+dgs[1]+","+dgs[2]+",", ttt)
			//fmt.Println("rr", rr)
			str1 := strings.Replace(rr, "[", "", -1)
			str2 := strings.Replace(str1, "]", ",", -1)
			str3 := strings.Replace(str2, ",,", "", -1)
			str4 := strings.ReplaceAll(str3, " ", "")
			//str5 := strings.ReplaceAll(str4, " ", ",,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,")
			var result = []string{str4}
			//fmt.Println(result)
			for _, v := range result {
				_, err = fmt.Fprintln(splitCommon, v)
				if err != nil {
					split.Close()
					return
				}
			}
			//for _, t := range ttt {
			//	fmt.Println("T", dgs[0], t[0])
			//}
		}
	*/
}

func parseParts(reportParts, reportDGS, panacimdata [][]string, pars string) {

	report, err := os.Create("/home/eugenearch/Code/github.com/eugenefoxx/SQLPanacimP1/csvfolder/" + pars + ".csv")
	if err != nil {
		log.Println(err)
	}
	defer report.Close()

	split, err := os.OpenFile("/home/eugenearch/Code/github.com/eugenefoxx/SQLPanacimP1/csvfolder/"+pars+".csv", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
		return
	}
	defer split.Close()

	for p := 0; p < len(panacimdata); p++ {
		if panacimdata[p][0] == pars {
			//var resultp = []string{pars + "," + panacimdata[p][1]}
			var resultp = []string{pars}
			fmt.Println(resultp)
			for _, v := range resultp {
				_, err = fmt.Fprintln(split, v)
				if err != nil {
					split.Close()
					return
				}
			}

		}

	}

	//	for p := 0; p < len(panacimdata); p++ {
	for i := 0; i < len(reportParts); i++ {
		for ii := 0; ii < len(reportDGS); ii++ {
			if reportParts[i][0] == pars {
				//if panacimdata[p][0] == reportParts[i][0] {
				//	var result = []string{reportParts[i][1] + "," + panacimdata[p][1]}
				var result = []string{reportParts[i][1]}
				fmt.Println(result)
				for _, v := range result {
					_, err = fmt.Fprintln(split, v)
					if err != nil {
						split.Close()
						return
					}
				}

				//}
				break
				//			}

			}
		}
	}

	/*	for s := 0; s < len(pp); s++ {
		fmt.Println("Test", pp[s][0])
	}*/

}

func insertPanacimDataQty(panacimdata [][]string, pars string) {

	pp := readfile("/home/eugenearch/Code/github.com/eugenefoxx/SQLPanacimP1/csvfolder/" + pars + ".csv")

	report, err := os.Create("/home/eugenearch/Code/github.com/eugenefoxx/SQLPanacimP1/csvfolder/" + pars + "pana.csv")
	if err != nil {
		log.Println(err)
	}
	defer report.Close()
	split2, err := os.OpenFile("/home/eugenearch/Code/github.com/eugenefoxx/SQLPanacimP1/csvfolder/"+pars+"pana.csv", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
		return
	}
	defer split2.Close()
	for ppp := 0; ppp < len(panacimdata); ppp++ {
		for s := 0; s < len(pp); s++ {
			if panacimdata[ppp][0] == pp[s][0] {
				var resultp = []string{pp[s][0] + "," + panacimdata[ppp][1]}
				fmt.Println(resultp)
				for _, v := range resultp {
					_, err = fmt.Fprintln(split2, v)
					if err != nil {
						split2.Close()
						return
					}
				}
			}
		}
	}
}

func insertPanacimDataQtyTotal(pars string) {
	readFile := readfile("/home/eugenearch/Code/github.com/eugenefoxx/SQLPanacimP1/csvfolder/" + pars + "pana.csv")

	report, err := os.Create("/home/eugenearch/Code/github.com/eugenefoxx/SQLPanacimP1/csvfolder/" + pars + "panatotal.csv")
	if err != nil {
		log.Println(err)
	}
	defer report.Close()
	split, err := os.OpenFile("/home/eugenearch/Code/github.com/eugenefoxx/SQLPanacimP1/csvfolder/"+pars+"panatotal.csv", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
		return
	}
	defer split.Close()
	sumCol := 0
	for i := 0; i < len(readFile); i++ {
		convertsumCol, err := strconv.Atoi(readFile[i][1])
		if err != nil {
			log.Println(err)
		}
		sumCol += (convertsumCol)
	}

	var result = []string{"Total:" + "," + strconv.Itoa(sumCol)}
	for _, v := range result {
		_, err = fmt.Fprintln(split, v)
		if err != nil {
			split.Close()
			return
		}
	}

}

func sumComponent(reportDGS, reportSumRead [][]string, component string) {
	/*	report, err := os.Create("/home/eugenearch/Code/github.com/eugenefoxx/SQLPanacimP1/csvfolder/reportSumComponent.csv")
		if err != nil {
			log.Println(err)
		}
		defer report.Close()*/

	split, err := os.OpenFile("/home/eugenearch/Code/github.com/eugenefoxx/SQLPanacimP1/csvfolder/reportSumComponent.csv", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
		return
	}
	defer split.Close()

	parts := readfile("/home/eugenearch/Code/github.com/eugenefoxx/SQLPanacimP1/csvfolder/" + component + "panatotal.csv")
	fmt.Printf("readfile TEST %s %v\n", component, parts)

	for rp := 0; rp < len(reportDGS); rp++ {
		for p := 0; p < len(parts); p++ {
			//	for dub := 0; dub < len(reportSumRead); dub++ {
			//fmt.Println("testing")
			/*	if reportSumRead[dub][0] == component {
				sumc, err := strconv.Atoi(parts[p][1])
				if err != nil {
					log.Println(err)
				}
				sum2, err := strconv.Atoi(reportSumRead[dub][3])
				if err != nil {
					log.Println(err)
				}
				sum3 := sumc + sum2
				var result = []string{reportDGS[rp][0] + "," + reportDGS[rp][1] + "," + reportDGS[rp][2] + "," + strconv.Itoa(sum3)}
				for _, v := range result {
					_, err = fmt.Fprintln(split, v)
					if err != nil {
						split.Close()
						return
					}
				}
			}*/
			sumc, err := strconv.Atoi(parts[p][1])
			if err != nil {
				log.Println(err)
			}
			if reportDGS[rp][0] == component {
				fmt.Printf("reportDGS Test %v\n", reportDGS[rp][0])
				//sumc := 0

				// суммирую колонку с кол-вом установленных компонентов
				//sum += (sumc)
				fmt.Printf("reportDGS Test Sum %v %v\n", reportDGS[rp][0], sumc)
				var result = []string{reportDGS[rp][0] + "," + reportDGS[rp][1] + "," + reportDGS[rp][2] + "," + strconv.Itoa(sumc)}
				//var result = []string{reportDGS[rp][0] + "," + reportDGS[rp][1] + "," + reportDGS[rp][2] + "," + strconv.Itoa(sum)}
				fmt.Println("result TEST", result)
				for _, v := range result {
					_, err = fmt.Fprintln(split, v)
					if err != nil {
						split.Close()
						return
					}
				}
			}
			//	}
		}
	}
	if parts == nil {
		fmt.Println("nil", component)
		for rp := 0; rp < len(reportDGS); rp++ {
			if reportDGS[rp][0] == component {
				fmt.Printf("reportDGS Test %v\n", reportDGS[rp][0])

				fmt.Printf("reportDGS Test Sum %v %v\n", reportDGS[rp][0], "0")
				var result = []string{reportDGS[rp][0] + "," + reportDGS[rp][1] + "," + reportDGS[rp][2] + "," + "0"}
				fmt.Println("result TEST", result)
				for _, v := range result {
					_, err = fmt.Fprintln(split, v)
					if err != nil {
						split.Close()
						return
					}
				}
			}

		}
	}
}

func readfileseeker(name string) [][]string {
	f, err := os.Open(name)
	if err != nil {
		fmt.Println(err)

	}
	defer f.Close()

	cr, err := readseeker(f)
	if err != nil {
		log.Fatalf("error read %s", err)
	}

	return cr
}

func readseeker(rs io.ReadSeeker) ([][]string, error) {
	row1, err := bufio.NewReader(rs).ReadSlice('\n')
	if err != nil {
		return nil, err
	}

	_, err = rs.Seek(int64(len(row1)), io.SeekStart)
	if err != nil {
		return nil, err
	}

	lines := csv.NewReader(rs) //.ReadAll()
	//lines.Comma = '|'
	lines.Comma = ','
	lines.LazyQuotes = true

	CSVdata, err := lines.ReadAll()
	if err != nil {
		//	fmt.Println(err)
		//	os.Exit(1)
		log.Fatal(err)
	}

	return CSVdata, nil
}

func readfile(name string) [][]string {
	f, err := os.Open(name)
	if err != nil {
		fmt.Println(err)

	}
	defer f.Close()

	cr := csv.NewReader(f)

	cr.LazyQuotes = true
	cr.Comma = ','
	//	cr.FieldsPerRecord = 10

	CSVdata, err := cr.ReadAll()
	if err != nil {
		log.Fatalf("readfile %s", err)
	}
	return CSVdata
}

/*
func ragnearr1(args [][]string) [1][3]string {
	var arr = [3]string{}
	var arrofarr = [1][3]string{}
	for _, iter := range args {
		qty, err := strconv.Atoi(iter[1])
		if err != nil {
			log.Println(err)
		}
		arr = [3]string{iter[0], iter[1], strconv.Itoa(int(uint16(qty) * value))}

		//fmt.Println(arrofarr)
		//return arrofarr
	}
	for i := range arrofarr {
		arrofarr[i] = arr
	}
	return arrofarr
}
*/
