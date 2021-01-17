package main

import(
	"fmt"
	"os"
	"bufio"
	"strings"
	"errors"
)

type Table struct{
	names []string
	columns [][]string
}

func main() {
	argsWithoutProgram := os.Args[1:]
	if (len(argsWithoutProgram) != 5){
		fmt.Println("Usage:: join ACTION FILE KEY FILE KEY")
		return
	}
	action := argsWithoutProgram[0]
	file1 := argsWithoutProgram[1]
	key1 := argsWithoutProgram[2]
	file2 := argsWithoutProgram[3]
	key2 := argsWithoutProgram[4]
	table1 , err := readCsv(file1)
	if err != nil{
		fmt.Printf("%s was not propely formatted missing some , ?" , file1)
		fmt.Println(err)
		return
	}
	table2 , err := readCsv(file2)
	if err != nil{
		fmt.Printf("%s was not propely formatted missing some , ?" , file2)
		fmt.Println(err)
		return
	}
	result , err := hashInnerJoin(table1,table2,key1,key2)
	if err != nil {
		fmt.Println("Join failed, are keys inside of the csv headers?")
		return
	}
	printTable(* result)
	fmt.Printf("joining %s on %s with %s on %s using %s",file1,key1,file2,key2,action)
}

func readCsv(fileName string) (*Table, error){
	result := new(Table)
	file, err := os.Open(fileName)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if err != nil{
		return nil, err
	}
	if (scanner.Scan()){
		result.names = strings.Split(scanner.Text(),",")
	}
	for scanner.Scan(){
		result.columns = append(result.columns,strings.Split(scanner.Text(),","))
	}
	return result , nil
}
func printTable(t Table){
	for i:= 0 ;i<len(t.names); i ++{
		fmt.Print(t.names[i])
		if i < len(t.names) - 1{
			fmt.Print(",")
		}
	}
	fmt.Println()
	for i:= 0;i<len(t.columns);i++{
		for j:= 0;j<len(t.columns[i]);j++{
			fmt.Print(t.columns[i][j])
			if j < (len(t.columns[i]) - 1){
				fmt.Print(",")
			}
		}
		fmt.Println()
	}
}

func hashInnerJoin(firstTable *Table,secondTable *Table,firstKey string, secondKey string) (*Table, error){
	hash := make(map[string][]string)
	sizeFirst := len(firstTable.columns)
	sizeSecond := len(secondTable.columns)
	var smallest *Table
	var bigger *Table
	join := new(Table)
	
	if sizeFirst < sizeSecond {
		smallest = firstTable
		bigger = secondTable
	}else{
		smallest = secondTable
		bigger = firstTable
	}
	join.names = append(smallest.names,bigger.names...)
	indexSmall := -1
	indexBig := -1
	for i := 0; i < len(smallest.names); i ++ {
		if (smallest.names[i] == firstKey){
			indexSmall = i
			break
		}
	}
	for i := 0; i < len(bigger.names); i ++ {
		if (bigger.names[i] == secondKey){
			indexBig = i
			break
		}
	}
	if (indexBig == -1 || indexSmall == -1 ){
		return nil, errors.New("key  not found")
	}
	for i:= 0; i < len(smallest.columns);i++{
		hash[smallest.columns[i][indexSmall]] = smallest.columns[i]
	}

	for i:= 0; i < len(bigger.columns); i++{
		if val, ok := hash[bigger.columns[i][indexBig]]; ok  {
			joinItem := append(val,bigger.columns[i]...)
			join.columns = append(join.columns,joinItem)
		}
	}

	return join , nil
}

func sortedMergeInnerJoin(leftTable Table,rightTable Table,leftKey string, rightKey string) *Table{
	hash := make(map[string][]string)
	join := new(Table)
	indexLeft := -1
	indexRight := -1
	for i := 0; i < len(leftTable.names); i ++ {
		if leftTable.names[i] == leftKey {
			indexLeft = i
		}
	}

	for i := 0; i < len(rightTable.names); i ++ {
		if rightTable.names[i] == rightKey {
			indexRight = i
		}
	}

	for i := 0; i < len(rightTable.columns); i++ {
		hash[rightTable.columns[i][indexRight]] = rightTable.columns[i]
	}
	for i := 0; i < len(leftTable.columns); i++ {
		if val, ok := hash[leftTable.columns[i][indexLeft]]; ok{
			join.columns = append(join.columns,append(leftTable.columns[i],val...))
		}else{
			nulls := makeNulls(len(rightTable.names))
			join.columns = append(join.columns,append(leftTable.columns[i],nulls...))
		}
	}
	return new(Table)
}

func makeNulls(total int)[]string{
	result := []string{}
	for i:= 0; i<total;i++{
		result = append(result,",")
	}
	return result
}

func hashLeftJoin(firstTable Table,secondTable Table,firstKey string, secondKey string) *Table{

	return new(Table)
}

func sortedMergeLeftJoin(firstTable Table,secondTable Table,firstKey string, secondKey string) *Table{

	return new(Table)
}
func hashOuterJoin(firstTable Table,secondTable Table,firstKey string, secondKey string) *Table{

	return new(Table)
}

func sortedMergeOuterJoin(firstTable Table,secondTable Table,firstKey string, secondKey string) *Table{

	return new(Table)
}