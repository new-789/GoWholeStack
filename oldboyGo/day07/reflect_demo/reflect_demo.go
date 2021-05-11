package reflect_demo

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

func LocadIni(fileName string, data interface{}) (err error) {
	// 0. 参数的校验
	// 0.1. 传进来的 data 参数必须是指针类型(因为需要在函数中对其赋值)
	t := reflect.TypeOf(data)
	if t.Kind() != reflect.Ptr {
		err = errors.New("data param should be a pointer")
		return
	}
	// 0.2.传进来的 data 参数，必须是结构体类型指针(因为配置文件中由各种键值对需要赋值给结构体中的字段)
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data param should be a struct pointer")
		return
	}
	// 1. 读取文件得到字节类型数据
	// var fileInfo []byte
	fileInfo, rerr := ioutil.ReadFile(fileName)
	if rerr != nil {
		err = rerr
		return
	}
	// string(byteInfo) // 将字节类型的文件内容转换成字符串
	lineSlice := strings.Split(string(fileInfo), "\n")
	// 2. 一行行的读文件中的数据
	var structName string
	for i, line := range lineSlice {
		// 去掉字符串首尾的空格
		line = strings.TrimSpace(line)
		// 如果是空行则直接跳过
		if len(line) == 0 {
			continue
		}
		// 2.1. 如果是注释，就跳过
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		// 2.2. 如果是方括号开头或者方括号结尾的就表示是节点（section）
		if strings.HasPrefix(line, "[") {
			// 去除配置文件中只有单边括号的语法错误
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d Bracket stntax error ", i+1)
				return
			}
			// 把这一行首尾的 [] 去掉，取到中间的内容去掉首尾空格后如果长度为 0 则格式为 [   ]，则说明不正确
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d syntax error Bracket is empty", i+1)
				return
			}
			// 根据字符 sectionName 去 data 里面根据反射找到对应的结构体
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					// 说明找到了对应的结构体，把字段名记下来
					structName = field.Name
					fmt.Printf("找到%s对应的结构体%s\n", sectionName, structName)
				}
			}
		} else {
			// 2.3. 如果不是【 开头就是 = 分隔的键值对
			// 1. 以等号分隔这一行，等号左边为 key，等号右边为 value
			if strings.HasPrefix(line, "]") {
				err = fmt.Errorf("line:%d Bracket stntax error ", i+1)
				return
			}
			// 如过 line 中不包含 = 或者 line 以 = 号开头，或者  line 以 = 结尾，则配置文件内容都不符合条件
			if !strings.Contains(line, "=") || strings.HasPrefix(line, "=") || strings.HasSuffix(line, "=") {
				err = fmt.Errorf("line:%d sytax error", i+1)
				return
			}
			index := strings.Index(line, "=")
			key := strings.TrimSpace(line[:index])
			value := strings.TrimSpace(line[index+1:])
			// 2. 根据 structName 去 data 中将对应的结构体字段取出来
			v := reflect.ValueOf(data)
			sValue := v.Elem().FieldByName(structName) // 拿到嵌套结构体的值信息
			sType := sValue.Type()                     // 拿到嵌套结构体的类型信息
			if sType.Kind() != reflect.Struct {
				err = fmt.Errorf("data 中的 %s 字段应该是一个结构体", structName)
				return
			}
			var fieldName string
			// var fieldType reflect.StructField
			// 3. 遍历嵌套结构体的每一个字段，判断 tag 是不是等于 key
			for i := 0; i < sValue.NumField(); i++ {
				field := sType.Field(i) // tag 信息是存储在类型信息中的
				// fieldType = field
				if field.Tag.Get("ini") == key {
					// 找到了对应的字段
					fieldName = field.Name
					// break
				}
			}
			// 4. 如果key = tag, 给这个字符安赋值
			// 4.1: 根据 fieldName 去取出这个字段
			if len(fieldName) == 0 {
				// 在结构体中找不到对应的字段
				continue
			}
			fieldObj := sValue.FieldByName(fieldName)
			// 4.2: 对其赋值
			// fmt.Println(fieldName, fieldObj.Type().Kind())
			switch fieldObj.Type().Kind() {
			case reflect.String:
				fieldObj.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				// 将字符串按照十进制数转换为64位的 int 类型
				var valueInt int64
				valueInt, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					err = fmt.Errorf("config file is line:%d value type error", i+1)
					return
				}
				fieldObj.SetInt(valueInt)
			case reflect.Bool:
				var valueBool bool
				valueBool, err = strconv.ParseBool(value)
				if err != nil {
					err = fmt.Errorf("config file is line:%d value type error", i+1)
					return
				}
				fieldObj.SetBool(valueBool)
			case reflect.Float32, reflect.Float64:
				var valueFloat float64
				valueFloat, err = strconv.ParseFloat(value, 64)
				fieldObj.SetFloat(valueFloat)
			}
		}
	}
	return
}
