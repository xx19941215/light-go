package main

import (
	"io"
	"io/ioutil"
	"os"
	"fmt"
	"bufio"
)

//创建一个新写文件，写入5句话
func main() {
	filePath := "/var/log/file1.log"
	file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	defer file.Close()
	str := "hello world\n"
	//写入时，使用带缓存的*Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}

	//因为writer是带缓存的，因此在调用writer的时候，其实内容先是
	//写到缓存中的，所以需要使用Flush方法，将缓存的数据真正的写入到文件中
	writer.Flush()
}

//覆盖写入、追加写入
func overWrite() {
	filePath := "/var/log/file.log"
	file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_TRUNC, 0666)
	// file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}

	defer file.Close()

	str := "你好\r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}

	writer.Flush()
}

//将一个文件内容写入另一个文件
func writeCopy() {
	file1Path := "/var/log/path1.log"
	file2Path := "/var/log/path2.log"

	data, err := ioutil.ReadFile(file1Path)
	if err != nil {
		fmt.Printf("read file err=%v", err)
		return
	}

	err = ioutil.WriteFile(file2Path, data, 0666)
	if err != nil {
		fmt.Printf("write file err = %v", err)
	}
}

//判断文件是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	//文件存在
	if err == nil {
		return true, nil
	}
	//文件或文件夹不存在
	if os.IsNotExist(err) {
		return false, nil
	}
	//如果错误类型是其他类型，不确定是否存在
	return false, err
}

func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("open file err=%v", err)
	}

	defer srcFile.Close()
	//通过 srcfile ,获取到 Reader

	reader := bufio.NewReader(srcFile)

	//打开 dstFileName
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v", err)
		return
	}

	//通过 dstFile, 获取到 Writer
	writter := bufio.NewWriter(dstFile)

	defer dstFile.Close()

	return io.Copy(writter, reader)
}

//统计英文、数字、空格和其他字符数量
type CharCount struct {
	ChCount int //英文个数
	NumCount int //记录数字个数
	SpaceCount int //空格个数
	OtherCount int //其他个数
}
func StatisticsCharactor() {
	fileName := "/var/log/file.log"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
	}

	defer file.Close()
	var count CharCount

	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		newStr := []rune(str)
		for _, v := range newStr {
			switch {
				case v >= 'a' && v <= 'z':
				fallthrough
				case v >= 'A' && v <= 'Z':
					count.ChCount++
				case v == ' ' || v == '\t':
					count.SpaceCount++
				case v >= '0' && v <= '9':
					count.NumCount++
				default:
					count.OtherCount++
			}
		}

		fmt.Printf("字符的个数为=%v 数字的个数为=%v 空格的个数为=%v 其它字符个数=%v",
		count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)
	}
	
}