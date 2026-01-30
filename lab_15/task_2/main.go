package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"golang.org/x/term"
)

const pageSize = 10

type Viewer struct {
	filename    string
	lines       []string
	pos         int
	showNumbers bool
}

func clearScreen() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	_ = cmd.Run()

	// fallback (если терминал не очистился)
	fmt.Print("\033[H")
}

func main() {
	showNumbers := flag.Bool("N", false, "show line numbers")
	flag.Parse()

	args := flag.Args()

	filename := ""
	for _, a := range args {
		if !strings.HasPrefix(a, "-") {
			filename = a
			break
		}
	}

	if filename == "" {
		fmt.Println("Missing filename.")
		return
	}

	lines, err := readAllLines(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	v := &Viewer{
		filename:    filename,
		lines:       lines,
		pos:         0,
		showNumbers: *showNumbers,
	}

	// включаем raw режим, чтобы ловить стрелки
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	reader := bufio.NewReader(os.Stdin)

	v.render()

	for {
		b, err := reader.ReadByte()
		if err != nil {
			return
		}

		// q = выйти
		if b == 'q' {
			clearScreen()
			return
		}

		// стрелки: ESC [ A/B
		if b == 27 {
			b2, _ := reader.ReadByte()
			b3, _ := reader.ReadByte()

			if b2 == '[' {
				switch b3 {
				case 'A': // up
					v.scrollUp()
				case 'B': // down
					v.scrollDown()
				}
				v.render()
				continue
			}
		}

		switch b {
		case ' ': // page down
			v.pageDown()
		case 'b': // page up
			v.pageUp()
		case 'g': // top
			v.goTop()
		case 'G': // bottom
			v.goBottom()
		}

		v.render()
	}
}

// ===================== логика навигации =====================

func (v *Viewer) scrollUp() {
	if v.pos > 0 {
		v.pos--
	}
}

func (v *Viewer) scrollDown() {
	if v.pos < len(v.lines)-1 {
		v.pos++
	}
}

func (v *Viewer) pageDown() {
	v.pos += pageSize
	if v.pos >= len(v.lines) {
		if len(v.lines) > 0 {
			v.pos = len(v.lines) - 1
		} else {
			v.pos = 0
		}
	}
}

func (v *Viewer) pageUp() {
	v.pos -= pageSize
	if v.pos < 0 {
		v.pos = 0
	}
}

func (v *Viewer) goTop() {
	v.pos = 0
}

func (v *Viewer) goBottom() {
	if len(v.lines) == 0 {
		v.pos = 0
		return
	}

	v.pos = len(v.lines) - pageSize
	if v.pos < 0 {
		v.pos = 0
	}
}

// ===================== вывод =====================

func (v *Viewer) render() {
	clearScreen()

	start := v.pos
	end := v.pos + pageSize
	if end > len(v.lines) {
		end = len(v.lines)
	}

	for i := start; i < end; i++ {
		line := v.lines[i]

		if v.showNumbers {
			fmt.Printf("%4d %s\n", i+1, line)
		} else {
			fmt.Println(line)
		}
	}

	fmt.Printf("\n%s\n", v.filename)
	fmt.Println("Controls: ↑ ↓  Space  b  g  G  q")
}

func clearScreen() {
	fmt.Print("\033[H")
}

// ===================== чтение файла =====================

func readAllLines(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines, sc.Err()
}
