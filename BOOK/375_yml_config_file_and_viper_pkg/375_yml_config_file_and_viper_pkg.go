package main

// YAML — это еще один популярный текстовый формат, который используется
// для конфигурационных файлов. Здесь показано, как читать
// конфигурационные файлы формата YAML с помощью пакета viper. Имя
// конфигурационного YAML-файла будет передаваться утилите в
// виде аргумента командной строки. Кроме того, в утилите будет
// использоваться функция viper.AddConfigPath() для добавления трех
// путей поиска, по которым viper будет автоматически искать
// конфигурационные файлы.

// Для того чтобы viper мог работать с аргументами командной строки,
//нам нужно импортировать пакет pflag
import (
	"fmt"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
)

func main() {
	var configFile *string = flag.String("c", "myConfig",
		"Setting the configuration file")
	// После определения аргументов командной строки, которые
	// поддерживает программа, всегда нужно вызывать flag.Parse().
	flag.Parse()

	_, err := os.Stat(*configFile)

	// код проверяет, существует ли файл, указанный в значении флага
	// конфигурации (--c), используя вызов os.Stat()
	if err == nil {
		fmt.Println("Using User Specified Configuration file!")
		viper.SetConfigFile(*configFile)
	} else {
		viper.SetConfigName(*configFile)
		viper.AddConfigPath("/tmp")
		viper.AddConfigPath("$HOME")
		viper.AddConfigPath(".")
	}
	// мы не указываем точно, что хотим использовать конфигурационный
	// файл в формате YAML, — программа будет искать все поддерживаемые
	// форматы файлов при условии, что имя файла без расширения —
	// myConfig. Именно так работает viper.

	// Чтение и синтаксический анализ YAML-файла
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("Using config: %s\n", viper.ConfigFileUsed())

	if viper.IsSet("item.k1") {
		fmt.Println("item1.val1:", viper.Get("item1.k1"))
	} else {
		fmt.Println("item1.k1 not set!")
	}

	if viper.IsSet("item1.k2") {
		fmt.Println("item1.val2:", viper.Get("item1.k2"))
	} else {
		fmt.Println("item1.k2 not set!")
	}

	if !viper.IsSet("item3.k1") {
		fmt.Println("item3.k1 is not set!")
	}
}
