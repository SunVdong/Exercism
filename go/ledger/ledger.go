package ledger

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

// 支持的货币类型
const (
	CurrencyEUR = "EUR"
	CurrencyUSD = "USD"
)

// 支持的地区
const (
	LocaleDutch = "nl-NL"
	LocaleUS    = "en-US"
)

// 格式化常量
const (
	MaxDescriptionLength = 25
	TruncateLength       = 22
	DateFormat           = "2006-01-02"
	DutchDateFormat      = "02-01-2006"
	USDateFormat         = "01/02/2006"
)

// 错误信息
var (
	ErrInvalidCurrency = errors.New("invalid currency")
	ErrInvalidLocale   = errors.New("invalid locale")
	ErrInvalidDate     = errors.New("invalid date")
)

// Entry 表示一个账目条目
type Entry struct {
	Date        string // 格式: "YYYY-MM-DD"
	Description string
	Change      int // 以分为单位
}

// EntryList 实现 sort.Interface 用于排序
type EntryList []Entry

// 表头映射
var headers = map[string][]string{
	LocaleDutch: {"Datum", "Omschrijving", "Verandering"},
	LocaleUS:    {"Date", "Description", "Change"},
}

// 货币符号映射
var currencySymbols = map[string]string{
	CurrencyEUR: "€",
	CurrencyUSD: "$",
}

// FormatLedger 格式化账目列表
func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	if err := validateInput(currency, locale); err != nil {
		return "", err
	}

	// 创建副本避免修改原始数据
	entriesCopy := make([]Entry, len(entries))
	copy(entriesCopy, entries)
	sort.Sort(EntryList(entriesCopy))

	return buildLedgerOutput(currency, locale, entriesCopy)
}

// validateInput 验证输入参数
func validateInput(currency, locale string) error {
	if currency != CurrencyEUR && currency != CurrencyUSD {
		return ErrInvalidCurrency
	}
	if locale != LocaleDutch && locale != LocaleUS {
		return ErrInvalidLocale
	}
	return nil
}

// buildLedgerOutput 构建账目输出
func buildLedgerOutput(currency, locale string, entries []Entry) (string, error) {
	var result strings.Builder

	// 添加表头
	result.WriteString(formatHeader(locale))

	// 添加条目
	for _, entry := range entries {
		formattedEntry, err := formatEntry(entry, currency, locale)
		if err != nil {
			return "", err
		}
		result.WriteString(formattedEntry)
	}

	return result.String(), nil
}

// formatEntry 格式化单个条目
func formatEntry(entry Entry, currency, locale string) (string, error) {
	date, err := formatDate(entry.Date, locale)
	if err != nil {
		return "", err
	}

	description := formatDescription(entry.Description)
	change := formatChange(entry.Change, currency, locale)

	return fmt.Sprintf("%-10s | %-25s | %13s\n", date, description, change), nil
}

// formatDate 格式化日期
func formatDate(dateStr, locale string) (string, error) {
	date, err := time.Parse(DateFormat, dateStr)
	if err != nil {
		return "", ErrInvalidDate
	}

	switch locale {
	case LocaleDutch:
		return date.Format(DutchDateFormat), nil
	case LocaleUS:
		return date.Format(USDateFormat), nil
	default:
		return "", ErrInvalidLocale
	}
}

// formatDescription 格式化描述，截断过长的描述
func formatDescription(description string) string {
	if len(description) > MaxDescriptionLength {
		return description[:TruncateLength] + "..."
	}
	return description
}

// formatChange 格式化金额变化
func formatChange(change int, currency, locale string) string {
	symbol := currencySymbols[currency]
	isNegative := change < 0
	absChange := abs(change)

	switch locale {
	case LocaleDutch:
		return formatDutchChange(absChange, symbol, isNegative)
	case LocaleUS:
		return formatUSChange(absChange, symbol, isNegative)
	default:
		return ""
	}
}

// formatDutchChange 格式化荷兰格式的金额
func formatDutchChange(absChange int, symbol string, isNegative bool) string {
	dollars := absChange / 100
	cents := absChange % 100

	var result string
	if dollars >= 1000 {
		// 大数字使用点号分隔千位
		dollarsStr := formatNumberWithDots(dollars)
		result = fmt.Sprintf("%s %s,%02d", symbol, dollarsStr, cents)
	} else {
		result = fmt.Sprintf("%s %d,%02d", symbol, dollars, cents)
	}

	// 荷兰格式：负数在末尾加负号
	if isNegative {
		result += "-"
	} else {
		result += " "
	}

	return result
}

// formatUSChange 格式化美式格式的金额
func formatUSChange(absChange int, symbol string, isNegative bool) string {
	dollars := absChange / 100
	cents := absChange % 100
	dollarsStr := formatNumberWithCommas(dollars)

	if isNegative {
		// 美式格式：负数用括号包围
		return fmt.Sprintf("(%s%s.%02d)", symbol, dollarsStr, cents)
	} else {
		return fmt.Sprintf(" %s%s.%02d ", symbol, dollarsStr, cents)
	}
}

// formatNumberWithCommas 为数字添加千位逗号分隔符
func formatNumberWithCommas(n int) string {
	str := strconv.Itoa(n)
	if n < 0 {
		str = str[1:] // 移除负号
	}

	// 从右往左每三位添加逗号
	for i := len(str) - 3; i > 0; i -= 3 {
		str = str[:i] + "," + str[i:]
	}

	if n < 0 {
		str = "-" + str
	}
	return str
}

// formatNumberWithDots 为数字添加千位点号分隔符
func formatNumberWithDots(n int) string {
	str := strconv.Itoa(n)

	// 从右往左每三位添加点号
	for i := len(str) - 3; i > 0; i -= 3 {
		str = str[:i] + "." + str[i:]
	}

	return str
}

// abs 返回整数的绝对值
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// formatHeader 格式化表头
func formatHeader(locale string) string {
	titles, ok := headers[locale]
	if !ok {
		return "Unsupported locale\n"
	}
	return fmt.Sprintf("%-10s | %-25s | %s\n", titles[0], titles[1], titles[2])
}

// sort.Interface 实现
func (el EntryList) Len() int {
	return len(el)
}

// Less 定义排序规则：日期 > 描述 > 金额
func (el EntryList) Less(i, j int) bool {
	// 首先按日期排序
	timeI, _ := time.Parse(DateFormat, el[i].Date)
	timeJ, _ := time.Parse(DateFormat, el[j].Date)
	if !timeI.Equal(timeJ) {
		return timeI.Before(timeJ)
	}

	// 日期相同时按描述排序
	if el[i].Description != el[j].Description {
		return el[i].Description < el[j].Description
	}

	// 描述相同时按金额排序
	return el[i].Change < el[j].Change
}

func (el EntryList) Swap(i, j int) {
	el[i], el[j] = el[j], el[i]
}
