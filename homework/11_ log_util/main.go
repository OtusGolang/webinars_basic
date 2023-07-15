package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const (
	defaultLogLevel = "info"
)

var (
	logFileFlag    = flag.String("file", "", "Path to the log file")
	logLevelFlag   = flag.String("level", defaultLogLevel, "Log level for analysis")
	outputFileFlag = flag.String("output", "", "Path to the output file")
)

func main() {
	flag.Parse()

	logFilePath := getLogFilePath()
	logLevel := getLogLevel()
	outputFilePath := getOutputFilePath()

	// Чтение лог-файла
	logData, err := os.ReadFile(logFilePath)
	if err != nil {
		log.Fatalf("Failed to read log file: %v", err)
	}

	// Анализ логов и сбор статистики
	statistics := analyzeLogs(string(logData), logLevel)

	// Вывод статистики
	if outputFilePath != "" {
		err = writeToFile(outputFilePath, statistics)
		if err != nil {
			log.Fatalf("Failed to write statistics to file: %v", err)
		}
	} else {
		fmt.Println(statistics)
	}
}

func getLogFilePath() string {
	logFilePath := *logFileFlag
	if logFilePath == "" {
		logFilePath = os.Getenv("LOG_ANALYZER_FILE")
	}

	if logFilePath == "" {
		log.Fatal("Log file path is required")
	}

	return logFilePath
}

func getLogLevel() string {
	logLevel := *logLevelFlag
	if logLevel == defaultLogLevel {
		logLevel = os.Getenv("LOG_ANALYZER_LEVEL")
	}

	return logLevel
}

func getOutputFilePath() string {
	outputFilePath := *outputFileFlag
	if outputFilePath == "" {
		outputFilePath = os.Getenv("LOG_ANALYZER_OUTPUT")
	}

	return outputFilePath
}

func analyzeLogs(logData, logLevel string) string {
	// Здесь происходит анализ логов и сбор статистики в соответствии с заданным уровнем логов.
	// Реализуйте эту логику в соответствии с вашими требованиями.
	// Возвращаемая строка должна содержать статистику.

	// Пример временной реализации, возвращающей просто информацию о лог-файле:
	statistics := fmt.Sprintf("Log File: %s\nLog Level: %s\nLog Data:\n%s", getLogFilePath(), logLevel, logData)
	return statistics
}

func writeToFile(filePath, data string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		return err
	}

	return nil
}
