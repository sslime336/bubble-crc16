package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	list     list.Model
	choice   string
	txtinput textinput.Model
	opDepth  int
}

func initializeModel() model {
	l := list.New(items, itemDelegate{}, 20, 20)
	l.Title = "请选择一个CRC16模型："
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = lipgloss.NewStyle().MarginLeft(2)
	l.Styles.PaginationStyle = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	l.Styles.HelpStyle = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)

	ti := textinput.New()
	ti.Width = 20
	ti.CharLimit = 256

	m := model{
		list:     l,
		choice:   "",
		txtinput: ti,
		opDepth:  0,
	}
	return m
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			if m.opDepth == 0 {
				m.opDepth++
				i, ok := m.list.SelectedItem().(item)
				if ok {
					m.choice = string(i)
				}
				return m, m.txtinput.Focus()
			}
			value := m.txtinput.Value()
			if len(value)%2 != 0 {
				value = "0" + value
			}
			byts, _ := hex.DecodeString(value)
			res := Checksum(parseAlgorithm(m.choice), byts)
			checksumRecord.WriteString(fmt.Sprintf("[%d] %s 使用 %s 算法的结果是：%#x\n", recordCount+1, m.txtinput.Value(), m.choice, res))
			m.txtinput.SetValue("")
			recordCount++
		}
	}

	if m.opDepth > 0 {
		var cmd tea.Cmd
		m.txtinput, cmd = m.txtinput.Update(msg)
		return m, cmd
	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

var checksumRecord = strings.Builder{}
var recordCount = 0

func (m model) View() string {
	if m.opDepth > 0 {
		res := checksumRecord.String()
		if res != "" {
			return "历史记录：\n" + res + "\n" + "请输入十六进制数(忽略前导'0x')：\n" + m.txtinput.View()
		}
		return "请输入十六进制数(忽略前导'0x')：\n" + m.txtinput.View()
	}
	return "\n" + m.list.View()
}

var items = func() []list.Item {
	res := make([]list.Item, SupportedCRC16ModelLen)
	var i int
	for modelKey := range CrcModelMap {
		res[i] = item(modelKey.String())
		i++
	}
	return res
}()

func parseAlgorithm(target string) Algorithm {
	for modelKey, alg := range CrcModelMap {
		if modelKey.String() == target {
			return alg
		}
	}
	panic("unreachable when parsing algorithm")
}

type item string

func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                         { return 1 }
func (d itemDelegate) Spacing() int                        { return 0 }
func (d itemDelegate) Update(tea.Msg, *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, _ := listItem.(item)

	str := fmt.Sprintf("%d. %s", index+1, i)

	itemStyle := lipgloss.NewStyle().PaddingLeft(4)
	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return lipgloss.
				NewStyle().
				PaddingLeft(2).
				Foreground(lipgloss.Color("170")).
				Render("> " + strings.Join(s, " "))
		}
	}

	_, _ = fmt.Fprint(w, fn(str))
}
