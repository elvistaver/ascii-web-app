package ascii

import("os"
"errors"
"bufio")


func Banner(file string)(map[rune][]string, error){
	filename, err:=os.Open("banners/"+file+".txt")

	if err !=nil{
		return nil, errors.New("error opening file")
	}
	defer filename.Close()

	scanner:=bufio.NewScanner(filename)
	slice:=[]string{}

	for scanner.Scan(){
		lines:=scanner.Text()
		slice=append(slice, lines)
	}
	if err:=scanner.Err(); err!=nil{
		return nil, errors.New("faild reading file")
	}
	font:=map[rune][]string{}

	for i:=1; i<len(slice); i+=9{
		char:=rune((i/9)+32)
		block:=slice[i:i+8]
		font[char]=block
	}
	return font, nil
}