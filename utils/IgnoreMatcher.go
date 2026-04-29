// Author: Simernin Matvei
// Created: 2026-04-26
// Description: Пакет для работы с правилами исключения файлов и папок при анализе проекта.
// Поддерживает синтаксис файла .codereportignore, аналогичный gitignore.
// Позволяет отфильтровать ненужные файлы, директории и типы расширений из отчета.
package utils

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

type ignoreRule struct {
	pattern  string
	negate   bool
	dirOnly  bool
	anchored bool
	hasSlash bool
}

type IgnoreMatcher struct {
	rules []ignoreRule
}

// NewIgnoreMatcher инициализирует новый IgnoreMatcher с правилами по умолчанию
// и загружает дополнительные правила из файла .codereportignore в корне проекта.
func NewIgnoreMatcher(root string) *IgnoreMatcher {
	matcher := &IgnoreMatcher{}
	matcher.addDefaults()
	matcher.loadFile(filepath.Join(root, ".codereportignore"))
	return matcher
}

// ShouldIgnore проверяет, должен ли файл или папка быть исключен из обработки.
// Применяет все загруженные правила и возвращает true, если файл нужно пропустить.
func (m *IgnoreMatcher) ShouldIgnore(relPath string, isDir bool) bool {
	relPath = filepath.ToSlash(filepath.Clean(relPath))
	if relPath == "." || relPath == "" {
		return false
	}

	ignored := false
	for _, rule := range m.rules {
		if rule.matches(relPath, isDir) {
			ignored = !rule.negate
		}
	}
	return ignored
}

// addDefaults добавляет встроенные правила исключения для распространённых
// папок кэша, файлов сборки, медиа и других вспомогательных файлов.
func (m *IgnoreMatcher) addDefaults() {
	defaultPatterns := []string{
		".git/",
		".idea/",
		".vscode/",
		"__pycache__/",
		".pytest_cache/",
		".mypy_cache/",
		".ruff_cache/",
		".tox/",
		".cache/",
		".venv/",
		"venv/",
		"env/",
		"node_modules/",
		"vendor/",
		"dist/",
		"build/",
		"target/",
		"out/",
		"bin/",
		"obj/",
		".next/",
		".nuxt/",
		"coverage/",
		"htmlcov/",
		"*.pyc",
		"*.pyo",
		"*.pyd",
		"*.exe",
		"*.dll",
		"*.so",
		"*.dylib",
		"*.class",
		"*.jar",
		"*.war",
		"*.ear",
		"*.zip",
		"*.tar",
		"*.gz",
		"*.rar",
		"*.7z",
		"*.doc",
		"*.docx",
		"*.xlsx",
		"*.pptx",
		"*.pdf",
		"*.png",
		"*.jpg",
		"*.jpeg",
		"*.gif",
		"*.svg",
		"*.webp",
		"*.ico",
		"*.mp3",
		"*.mp4",
		"*.webm",
		"*.ttf",
		"*.woff",
		"*.woff2",
		"*.eot",
		"*.otf",
		"*.map",
		"*.min.js",
		"*.min.css",
		"package-lock.json",
		"yarn.lock",
		"pnpm-lock.yaml",
		"composer.lock",
		"poetry.lock",
		"Pipfile.lock",
		"*.sqlite",
		"*.sqlite3",
		"*.db",
		"*.log",
		".env",
		".DS_Store",
	}

	for _, pattern := range defaultPatterns {
		m.addRule(pattern)
	}
}

// loadFile загружает и парсит правила исключения из файла .codereportignore.
// Если файл не существует, ошибка молча игнорируется.
func (m *IgnoreMatcher) loadFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		m.addRule(scanner.Text())
	}
}

// addRule парсит одну строку из файла .codereportignore и добавляет соответствующее правило.
// Поддерживает: комментарии (#), отрицание (!), якорь (/) и фильтр по папкам (/).
func (m *IgnoreMatcher) addRule(line string) {
	line = strings.TrimSpace(line)
	if line == "" || strings.HasPrefix(line, "#") {
		return
	}

	rule := ignoreRule{}
	if strings.HasPrefix(line, "!") {
		rule.negate = true
		line = strings.TrimSpace(strings.TrimPrefix(line, "!"))
	}
	if line == "" {
		return
	}

	line = filepath.ToSlash(line)
	if strings.HasPrefix(line, "/") {
		rule.anchored = true
		line = strings.TrimPrefix(line, "/")
	}
	if strings.HasSuffix(line, "/") {
		rule.dirOnly = true
		line = strings.TrimSuffix(line, "/")
	}

	rule.pattern = strings.Trim(line, "/")
	rule.hasSlash = strings.Contains(rule.pattern, "/")
	if rule.pattern != "" {
		m.rules = append(m.rules, rule)
	}
}

// matches проверяет, совпадает ли путь с правилом.
// Учитывает якорирование, наличие слэшей и тип элемента (файл/папка).
func (r ignoreRule) matches(relPath string, isDir bool) bool {
	if r.dirOnly && !isDir {
		return false
	}

	if r.anchored || r.hasSlash {
		return matchPattern(r.pattern, relPath)
	}

	parts := strings.Split(relPath, "/")
	for _, part := range parts {
		if matchPattern(r.pattern, part) {
			return true
		}
	}
	return false
}

// matchPattern проверяет соответствие паттерна (с поддержкой * и ?) имени файла или пути.
func matchPattern(pattern, name string) bool {
	matched, err := filepath.Match(pattern, name)
	if err == nil && matched {
		return true
	}
	return pattern == name || strings.HasPrefix(name, pattern+"/")
}
