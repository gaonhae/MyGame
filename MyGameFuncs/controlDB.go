package MyGameFuncs

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// ReadDB 한글 단어가 저장된 데이터베이스를 읽어 배열로 반환해줍니다(현재는 출력)
func ReadDB() []string{
	
	keys := make(map[string]struct{})
	res := make([]string, 0)  
	nouns := []string{}
	DB, err := os.Open("kr_korean.csv")
	HandleErr(err)

	r := csv.NewReader(DB)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if record[1] == "명사" && len(cleanDBString(record[0])) > 1 {
			nouns = append(nouns, cleanDBString(record[0]))
		}
	}

    for _, val := range nouns { 
        if _, ok := keys[val]; ok { 
            continue 
        } else { 
            keys[val] = struct{}{} 
            res = append(res, val) 
        } 
    }
	//res가 최종 배열

	fmt.Println(len(res))

	return res
}


// cleanDBString 데이터베이스에서 읽어온 문자들을 다듬습니다("-"등 제거)
func cleanDBString (s string) string{
	return strings.ReplaceAll(s, "-", "") + "\n"
}


// CleanDB 데이터베이스에서 읽어온 문자들을 정리합니다. (한글자, 중복 제거)
// func CleanDB() {
// 	nouns := []string{}

// 	file, err := os.Open("MyGameFuncs/KoreanNouns.txt")
// 	if err != nil {
// 		fmt.Print(err, "오픈 실패")
// 	}
//  	defer file.Close()
// 	reader := bufio.NewReader(file)
// 	for {
// 		line, isPrefix, err := reader.ReadLine()
// 		if isPrefix || err != nil {
// 		break
// 		}
// 		if len(string(line)) > 3 {
// 			nouns = append(nouns, string(line))
// 		}
// 	}
// 	nouns2,_ := os.Create("MyGameFuncs/KoreanNouns2.txt")
// 	keys := make(map[string]struct{}) 
//     res := make([]string, 0) 
//     for _, val := range nouns { 
//         if _, ok := keys[val]; ok { 
//             continue 
//         } else { 
//             keys[val] = struct{}{} 
//             res = append(res, val) 
//         } 
//     }
// 	for _, v := range res{
// 		nouns2.WriteString(v+"\n")
// 	}
// }

// buildNounTxts 정제된 단어들을 가지고 각 단어의 첫글자로 시작하는 파일을 생성합니다.
// func buildNounTxts() {
// 	nouns, err := os.Open("MyGameFuncs/KoreanNouns.txt")
// 	if err != nil {
// 		fmt.Print(err, "오픈 실패")
// 	}
//  	defer nouns.Close()

// 	reader := bufio.NewReader(nouns)

// 	for {
// 		line, isPrefix, err := reader.ReadLine()
// 		if isPrefix || err != nil {
// 		break
// 		}

// 		arr := []rune(string(line))


// 		filename := "MyGameFuncs/Nouns/" + string(arr[0]) + ".txt"
// 		fmt.Println(filename)

// 		file,err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE, 0755)
// 		if err != nil {
// 			fmt.Print(err)
// 		}
// 		defer file.Close()

// 		_,err = file.WriteString(string(line)+"\n")
// 		if err != nil {
// 			fmt.Print(err, "작성 실패")
// 		}
// 	}
// }