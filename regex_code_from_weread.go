package main

import (
        "bufio"
        "fmt"
        "os"
        "regexp"
)

func main() {
        // 打开文件
        file, err := os.Open(os.Args[1])
        if err != nil {
                fmt.Println("Error opening file:", err)
                return
        }
        defer file.Close()

        // 正则表达式，用于匹配行首的数字、冒号和两个空格
        re := regexp.MustCompile(`^\d+:{0,1}\s{0,2}`)

        // 逐行读取文件
        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
                line := scanner.Text()
                // 使用正则表达式替换
                processedLine := re.ReplaceAllString(line, "")

                fmt.Println(processedLine)
        }

        // 检查扫描过程中是否有错误
        if err := scanner.Err(); err != nil {
                fmt.Println("Error reading file:", err)
        }
}
