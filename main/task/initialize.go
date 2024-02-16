package task

import (
	"encoding/json"
	"fmt"
	"os"
)

type CountryConfig struct {
	Name []string `json:"Name"`
	URL  []string `json:"URL"`
}

type Config struct {
	Countries map[string]CountryConfig `json:"countries"`
}

func ReadFile(path string, selectedCountry string) ([]string, []string, error) {
	// 打开json配置文件
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, fmt.Errorf("打开文件失败: %v", err)
	}
	defer file.Close()

	// 解码json配置文件到Config中
	var config Config

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, nil, fmt.Errorf("解码失败: %v", err)
	}

	// 根据选择的国家获取相应的镜像站点信息
	countryConfig, found := config.Countries[selectedCountry]
	if !found {
		return nil, nil, fmt.Errorf("未找到国家：%s", selectedCountry)
	}

	return countryConfig.Name, countryConfig.URL, nil
}
