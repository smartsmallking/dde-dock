package main

import (
	"path"
	"strings"
)

const (
	_THEME_DIR       = "/boot/grub/themes"
	_THEME_MAIN_FILE = "theme.txt"
	_THEME_TPL_FILE  = "theme.tpl"
)

type ThemeTpl struct {
	Background, ItemColor, SelectedItemColor string
}

type Theme struct {
	name    string
	tplfile string
	tpldata ThemeTpl
	// zipfile  string
	// mainfile string
}

type ThemeManager struct {
	enabledThemeMainFile string
}

// TODO
func NewThemeManager() *ThemeManager {
	tm := &ThemeManager{}
	tm.enabledThemeMainFile = ""
	return tm
}

func (tm *ThemeManager) setEnabledThemeMainFile(file string) {
	// if the theme.txt file is not under theme dir(/boot/grub/themes), ignore it
	if strings.HasPrefix(file, _THEME_DIR) {
		tm.enabledThemeMainFile = file
	}
}

func (tm *ThemeManager) getEnabledThemeMainFile() string {
	return tm.enabledThemeMainFile
}

func (tm *ThemeManager) isThemeValid(themeName string) bool {
	_, okPath := tm.getThemePath(themeName)
	_, okMainFile := tm.getThemeMainFile(themeName)
	if okPath && okMainFile {
		return true
	} else {
		return false
	}
}

func (tm *ThemeManager) isThemeArchiveValid(archive string) bool {
	_, err := findFileInTarGz(archive, _THEME_MAIN_FILE)
	if err != nil {
		return false
	}
	return true
}

func (tm *ThemeManager) getThemeName(themeMainFile string) string {
	if len(themeMainFile) == 0 {
		return ""
	}
	return path.Base(path.Dir(themeMainFile))
}

func (tm *ThemeManager) getThemePath(themeName string) (string, bool) {
	themePath := path.Join(_THEME_DIR, themeName)
	if isFileExist(themePath) {
		return themePath, true
	}
	return "", false
}

func (tm *ThemeManager) getThemeMainFile(themeName string) (string, bool) {
	mainFile := path.Join(_THEME_DIR, themeName, _THEME_MAIN_FILE)
	if isFileExist(mainFile) {
		return mainFile, true
	}
	return "", false
}

func (tm *ThemeManager) getThemeTplFile(themeName string) (string, bool) {
	tplFile := path.Join(_THEME_DIR, themeName, _THEME_TPL_FILE)
	if isFileExist(tplFile) {
		return tplFile, true
	}
	return "", false
}
