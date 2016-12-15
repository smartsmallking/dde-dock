/**
 * Copyright (C) 2016 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

package audio

import (
	"encoding/json"
	dutils "pkg.deepin.io/lib/utils"
	"sync"
)

var (
	fileLocker    sync.Mutex
	configCache   *configInfo
	configHandler *dutils.Config
)

func init() {
	configHandler = new(dutils.Config)
	configHandler.SetConfigName("dde-daemon/audio")
}

type configInfo struct {
	Profiles   map[string]string // Profiles[cardName] = activeProfile
	Sink       string
	Source     string
	SinkPort   string
	SourcePort string

	SinkVolume   float64
	SourceVolume float64
}

func (info *configInfo) string() string {
	data, _ := json.Marshal(info)
	return string(data)
}

func readConfigInfo() (*configInfo, error) {
	fileLocker.Lock()
	defer fileLocker.Unlock()

	if configCache != nil {
		return configCache, nil
	}

	var info configInfo
	err := configHandler.Load(&info)
	if err != nil {
		return nil, err
	}

	configCache = &info
	return configCache, nil
}

func saveConfigInfo(info *configInfo) error {
	fileLocker.Lock()
	defer fileLocker.Unlock()

	logger.Debug("[saveConfigInfo] will save:", info.string())
	if configCache.string() == info.string() {
		logger.Debug("[saveConfigInfo] config info not changed")
		return nil
	}

	err := configHandler.Save(info)
	if err != nil {
		return err
	}

	configCache = info
	return nil
}
