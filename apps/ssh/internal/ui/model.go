package ui

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"ssh-portfolio/internal/data"
	"ssh-portfolio/internal/styles"
)

// ── Pages ──────────────────────────────────────────────────────────────

type Page int

const (
	PageHome Page = iota
	PageAbout
	PageExperience
	PageProjects
	PageSkills
	PageAchievements
	PageContact
)

var pageNames = []string{
	"Home",
	"About",
	"Experience",
	"Projects",
	"Skills",
	"Achievements",
	"Contact",
}

var pageIcons = []string{
	"◆",
	"●",
	"▶",
	"◈",
	"⬟",
	"★",
	"◉",
}

// ── Phases ─────────────────────────────────────────────────────────────

type Phase int

const (
	PhaseGlitchBoot Phase = iota // Fake kernel panic / BSOD
	PhaseGlitchTear              // Screen tears + corruption
	PhaseTermRepair              // "Repairing" terminal typing
	PhaseReady                   // Portfolio is visible
)

// ── Tick Messages ──────────────────────────────────────────────────────

type tickMsg time.Time
type glitchTickMsg time.Time
type typeTickMsg time.Time

// ── Model ──────────────────────────────────────────────────────────────

type Model struct {
	cv       *data.CV
	page     Page
	width    int
	height   int
	scrollY  int
	maxScroll int
	username string
	ready    bool

	// Animation state
	phase         Phase
	animFrame     int
	glitchLines   []string
	glitchIdx     int
	repairLines   []string
	repairIdx     int
	showCursor    bool
	cursorTick    int
	pageTransition int // frames remaining for page transition
	prevPage      Page

	// Home page animation
	homeRevealed  int // lines of ASCII art revealed
	homeTyped     int // characters of subtitle typed
	statsRevealed int // stats items revealed
}

func NewModel(cv *data.CV, username string) Model {
	return Model{
		cv:       cv,
		page:     PageHome,
		width:    80,
		height:   24,
		username: username,
		phase:    PhaseGlitchBoot,
	}
}

// ── Glitch Content ────────────────────────────────────────────────────

var glitchBootLines = []string{
	"[    0.000000] Linux version 6.1.0-portfolio (farhan@cloud) (gcc 12.2.0)",
	"[    0.000000] Command line: BOOT_IMAGE=/vmlinuz root=/dev/sda1",
	"[    0.000012] BIOS-provided physical RAM map:",
	"[    0.000015]  BIOS-e820: [mem 0x0000000000000000-0x000000000009fbff] usable",
	"[    0.000031] ACPI: RSDP 0x00000000000F6A10 000024 (v02 PTLTD )",
	"[    0.001204] CPU: Intel Core i9-13900K @ 5.80GHz",
	"[    0.001330] x86/fpu: x87 FPU on chip",
	"[    0.002100] smpboot: CPU0: booting...",
	"[    0.003442] Mounting root filesystem...",
	"[    0.004100] systemd[1]: Starting Portfolio Service...",
	"[    0.004500] portfolio.service: Loading cv.json...",
	"[    0.005100] portfolio.service: Initializing SSH server...",
	"",
	"[    0.005842] ████████████████████████████████████████████",
	"[    0.005843] █ KERNEL PANIC - not syncing: portfolio.sys █",
	"[    0.005844] ████████████████████████████████████████████",
	"",
	"[    0.005900] CPU: 0 PID: 1 Comm: portfolio Not tainted 6.1.0",
	"[    0.005901] Call Trace:",
	"[    0.005902]  dump_stack_lvl+0x37/0x4d",
	"[    0.005903]  panic+0x107/0x2d8",
	"[    0.005904]  portfolio_init+0x42/0x50",
	"[    0.005905]  mount_block_root+0x161/0x21b",
	"",
	"[    0.006000] ---[ end Kernel panic - not syncing ]---",
}

var repairTermLines = []string{
	"$ sudo systemctl restart portfolio.service",
	"  Restarting portfolio.service...",
	"",
	"$ portfolio --diagnostics",
	"  [████████████████████] 100%",
	"  System check: OK",
	"  SSH module:   OK",
	"  TUI engine:   OK",
	"  CV data:      LOADED",
	"",
	"$ portfolio --serve --port 2222",
	"  Initializing Bubble Tea...",
	"  Loading theme: zinc-monochrome",
	"  Mounting components...",
	"",
	"  ✓ Portfolio server is live!",
	"  ✓ Welcome, visitor.",
}

// ── Init ──────────────────────────────────────────────────────────────

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		tickGlitch(),
		tickCursor(),
	)
}

func tickGlitch() tea.Cmd {
	return tea.Tick(80*time.Millisecond, func(t time.Time) tea.Msg {
		return glitchTickMsg(t)
	})
}

func tickType() tea.Cmd {
	return tea.Tick(50*time.Millisecond, func(t time.Time) tea.Msg {
		return typeTickMsg(t)
	})
}

func tickCursor() tea.Cmd {
	return tea.Tick(530*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func tickAnim() tea.Cmd {
	return tea.Tick(60*time.Millisecond, func(t time.Time) tea.Msg {
		return typeTickMsg(t)
	})
}

// ── Update ─────────────────────────────────────────────────────────────

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.ready = true
		m.scrollY = 0

	case glitchTickMsg:
		if m.phase == PhaseGlitchBoot {
			if m.glitchIdx < len(glitchBootLines) {
				m.glitchLines = append(m.glitchLines, glitchBootLines[m.glitchIdx])
				m.glitchIdx++
				return m, tickGlitch()
			}
			// Move to tear phase
			m.phase = PhaseGlitchTear
			m.animFrame = 0
			return m, tickGlitch()
		}
		if m.phase == PhaseGlitchTear {
			m.animFrame++
			if m.animFrame > 15 {
				// Move to repair phase
				m.phase = PhaseTermRepair
				m.repairIdx = 0
				return m, tickType()
			}
			return m, tickGlitch()
		}

	case typeTickMsg:
		if m.phase == PhaseTermRepair {
			if m.repairIdx < len(repairTermLines) {
				m.repairLines = append(m.repairLines, repairTermLines[m.repairIdx])
				m.repairIdx++
				return m, tickType()
			}
			// Done repairing, transition to portfolio
			m.phase = PhaseReady
			m.homeRevealed = 0
			m.homeTyped = 0
			m.statsRevealed = 0
			return m, tickAnim()
		}
		if m.phase == PhaseReady {
			// Animate home page reveal
			changed := false
			if m.homeRevealed < 8 {
				m.homeRevealed++
				changed = true
			} else if m.homeTyped < len(m.cv.Basics.Label) {
				m.homeTyped += 2
				if m.homeTyped > len(m.cv.Basics.Label) {
					m.homeTyped = len(m.cv.Basics.Label)
				}
				changed = true
			} else if m.statsRevealed < 5 {
				m.statsRevealed++
				changed = true
			}
			if changed {
				return m, tickAnim()
			}
			// Animation complete
		}
		if m.pageTransition > 0 {
			m.pageTransition--
			if m.pageTransition > 0 {
				return m, tickAnim()
			}
		}

	case tickMsg:
		m.cursorTick++
		m.showCursor = m.cursorTick%2 == 0
		return m, tickCursor()

	case tea.KeyMsg:
		// Skip intro on any key
		if m.phase != PhaseReady {
			m.phase = PhaseReady
			m.homeRevealed = 8
			m.homeTyped = len(m.cv.Basics.Label)
			m.statsRevealed = 5
			return m, nil
		}

		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit

		// Tab navigation
		case "tab", "l", "right":
			m.prevPage = m.page
			m.page = (m.page + 1) % Page(len(pageNames))
			m.scrollY = 0
			m.pageTransition = 4
			return m, tickAnim()
		case "shift+tab", "h", "left":
			m.prevPage = m.page
			m.page = (m.page - 1 + Page(len(pageNames))) % Page(len(pageNames))
			m.scrollY = 0
			m.pageTransition = 4
			return m, tickAnim()

		// Direct page jumps
		case "1":
			m.prevPage = m.page
			m.page = PageHome
			m.scrollY = 0
		case "2":
			m.prevPage = m.page
			m.page = PageAbout
			m.scrollY = 0
		case "3":
			m.prevPage = m.page
			m.page = PageExperience
			m.scrollY = 0
		case "4":
			m.prevPage = m.page
			m.page = PageProjects
			m.scrollY = 0
		case "5":
			m.prevPage = m.page
			m.page = PageSkills
			m.scrollY = 0
		case "6":
			m.prevPage = m.page
			m.page = PageAchievements
			m.scrollY = 0
		case "7":
			m.prevPage = m.page
			m.page = PageContact
			m.scrollY = 0

		// Scrolling
		case "j", "down":
			m.scrollY += 2
		case "k", "up":
			m.scrollY -= 2
			if m.scrollY < 0 {
				m.scrollY = 0
			}
		case "g":
			m.scrollY = 0
		case "G":
			m.scrollY = 9999
		}
	}

	return m, nil
}

// ── View ───────────────────────────────────────────────────────────────

func (m Model) View() tea.View {
	if !m.ready {
		v := tea.NewView("\n  Connecting...")
		v.AltScreen = true
		return v
	}

	// During intro animation
	if m.phase != PhaseReady {
		content := m.renderIntro()
		v := tea.NewView(content)
		v.AltScreen = true
		return v
	}

	w := m.width
	if w > 100 {
		w = 100
	}

	contentWidth := w - 4

	// Build layout
	var b strings.Builder

	// Top bar
	b.WriteString(m.renderTopBar(w))
	b.WriteString("\n")

	// Tab bar
	b.WriteString(m.renderTabBar(w))
	b.WriteString("\n")
	b.WriteString(styles.Divider(w))
	b.WriteString("\n")

	// Page content
	content := m.renderPage(contentWidth)

	// Apply scrolling
	lines := strings.Split(content, "\n")
	availableHeight := m.height - 7

	if m.scrollY > len(lines)-availableHeight {
		m.scrollY = maxInt(0, len(lines)-availableHeight)
	}

	endLine := m.scrollY + availableHeight
	if endLine > len(lines) {
		endLine = len(lines)
	}

	if m.scrollY < len(lines) {
		visibleLines := lines[m.scrollY:endLine]
		b.WriteString(strings.Join(visibleLines, "\n"))
	}

	b.WriteString("\n")

	// Status bar
	b.WriteString(m.renderStatusBar(w, len(lines), availableHeight))

	// Center the whole thing
	fullContent := b.String()
	if m.width > 100 {
		fullContent = lipgloss.NewStyle().
			Width(m.width).
			Align(lipgloss.Center).
			Render(fullContent)
	}

	v := tea.NewView(fullContent)
	v.AltScreen = true
	return v
}

// ── Intro Screens ──────────────────────────────────────────────────────

func (m Model) renderIntro() string {
	var b strings.Builder

	switch m.phase {
	case PhaseGlitchBoot:
		// Kernel panic style boot
		for _, line := range m.glitchLines {
			if strings.Contains(line, "KERNEL PANIC") || strings.Contains(line, "████") {
				b.WriteString(styles.GlitchStyle.Render(line))
			} else if strings.Contains(line, "end Kernel panic") {
				b.WriteString(styles.GlitchStyle.Render(line))
			} else {
				b.WriteString(styles.DimTextStyle.Render(line))
			}
			b.WriteString("\n")
		}
		if m.showCursor {
			b.WriteString(styles.DimTextStyle.Render("_"))
		}

	case PhaseGlitchTear:
		// Corrupted screen — glitch art
		for i := 0; i < m.height; i++ {
			line := m.generateGlitchLine(m.width, i)
			b.WriteString(line)
			b.WriteString("\n")
		}

	case PhaseTermRepair:
		// Terminal repair sequence
		termWidth := m.width - 4
		if termWidth > 70 {
			termWidth = 70
		}

		// Terminal window chrome
		topBar := lipgloss.NewStyle().
			Foreground(styles.Zinc600).
			Render("  ● ● ●  ") +
			lipgloss.NewStyle().
				Foreground(styles.Zinc500).
				Render("recovery — ~/portfolio") +
			lipgloss.NewStyle().
				Foreground(styles.Emerald500).
				Render("  ● LIVE")

		b.WriteString("\n\n")
		b.WriteString(topBar)
		b.WriteString("\n")
		b.WriteString(styles.Divider(termWidth + 4))
		b.WriteString("\n")

		for _, line := range m.repairLines {
			if strings.HasPrefix(line, "$") {
				b.WriteString(lipgloss.NewStyle().
					Foreground(styles.Zinc200).
					Bold(true).
					Render("  " + line))
			} else if strings.Contains(line, "✓") {
				b.WriteString(lipgloss.NewStyle().
					Foreground(styles.Emerald500).
					Bold(true).
					Render("  " + line))
			} else if strings.Contains(line, "████") {
				b.WriteString(lipgloss.NewStyle().
					Foreground(styles.Emerald400).
					Render("  " + line))
			} else {
				b.WriteString(styles.DimTextStyle.Render("  " + line))
			}
			b.WriteString("\n")
		}

		if m.showCursor {
			b.WriteString(styles.DimTextStyle.Render("  _"))
		}
	}

	// Add skip hint at bottom
	skipHint := lipgloss.NewStyle().
		Foreground(styles.Zinc600).
		Italic(true).
		Render("\n\n  Press any key to skip...")

	// Pad to fill screen
	rendered := b.String()
	lines := strings.Count(rendered, "\n")
	for i := lines; i < m.height-3; i++ {
		b.WriteString("\n")
	}
	b.WriteString(skipHint)

	return b.String()
}

func (m Model) generateGlitchLine(width, row int) string {
	glitchChars := []rune("█▓▒░╔╗╚╝║═╬╣╠╩╦▄▀■□▪▫●◆◇◈★☆")
	normalChars := []rune(" .:;+=xX#@")

	var line strings.Builder
	intensity := float64(m.animFrame) / 15.0

	for col := 0; col < width; col++ {
		// Random corruption with varying intensity
		r := rand.Float64()
		if r < intensity*0.3 {
			// Glitch characters
			ch := glitchChars[rand.Intn(len(glitchChars))]
			if rand.Float64() < 0.3 {
				line.WriteString(lipgloss.NewStyle().Foreground(styles.Red500).Render(string(ch)))
			} else if rand.Float64() < 0.5 {
				line.WriteString(lipgloss.NewStyle().Foreground(styles.Emerald500).Render(string(ch)))
			} else {
				line.WriteString(lipgloss.NewStyle().Foreground(styles.Zinc700).Render(string(ch)))
			}
		} else if r < intensity*0.6 {
			ch := normalChars[rand.Intn(len(normalChars))]
			line.WriteString(lipgloss.NewStyle().Foreground(styles.Zinc800).Render(string(ch)))
		} else {
			line.WriteString(" ")
		}
	}

	return line.String()
}

// ── Top Bar ────────────────────────────────────────────────────────────

func (m Model) renderTopBar(width int) string {
	left := styles.LogoStyle.Render("  " + m.cv.Basics.Name)

	visitorLabel := "visitor"
	if m.username != "anonymous" && m.username != "" {
		visitorLabel = m.username
	}
	right := styles.DimTextStyle.Render(fmt.Sprintf("ssh · %s ", visitorLabel))

	gap := width - lipgloss.Width(left) - lipgloss.Width(right)
	if gap < 0 {
		gap = 0
	}

	return left + strings.Repeat(" ", gap) + right
}

// ── Tab Bar ────────────────────────────────────────────────────────────

func (m Model) renderTabBar(width int) string {
	var tabs []string
	for i, name := range pageNames {
		var label string
		if width < 80 {
			label = fmt.Sprintf(" %s ", pageIcons[i])
		} else {
			label = fmt.Sprintf(" %s %s ", pageIcons[i], name)
		}

		isActive := Page(i) == m.page
		if isActive {
			tabs = append(tabs, styles.ActiveTabStyle.Render(label))
		} else {
			tabs = append(tabs, styles.InactiveTabStyle.Render(label))
		}
	}
	return lipgloss.JoinHorizontal(lipgloss.Top, tabs...)
}

// ── Status Bar ─────────────────────────────────────────────────────────

func (m Model) renderStatusBar(width, totalLines, visibleLines int) string {
	// Minimal, clean status bar
	keys := []struct {
		key  string
		desc string
	}{
		{"←→", "nav"},
		{"↑↓", "scroll"},
		{"1-7", "jump"},
		{"q", "quit"},
	}

	var parts []string
	for _, k := range keys {
		parts = append(parts,
			styles.KeyActiveStyle.Render(k.key)+
				styles.KeyHintStyle.Render(" "+k.desc),
		)
	}

	nav := "  " + strings.Join(parts, styles.KeyHintStyle.Render(" · "))

	scrollInfo := ""
	if totalLines > visibleLines {
		pct := 0
		if totalLines-visibleLines > 0 {
			pct = m.scrollY * 100 / (totalLines - visibleLines)
		}
		if pct > 100 {
			pct = 100
		}
		scrollInfo = styles.ScrollStyle.Render(fmt.Sprintf(" %d%% ", pct))
	}

	gap := width - lipgloss.Width(nav) - lipgloss.Width(scrollInfo)
	if gap < 0 {
		gap = 0
	}

	return nav + strings.Repeat(" ", gap) + scrollInfo
}

// ── Page Router ────────────────────────────────────────────────────────

func (m Model) renderPage(width int) string {
	switch m.page {
	case PageHome:
		return m.renderHome(width)
	case PageAbout:
		return m.renderAbout(width)
	case PageExperience:
		return m.renderExperience(width)
	case PageProjects:
		return m.renderProjects(width)
	case PageSkills:
		return m.renderSkills(width)
	case PageAchievements:
		return m.renderAchievements(width)
	case PageContact:
		return m.renderContact(width)
	default:
		return ""
	}
}

// ── Page: Home ─────────────────────────────────────────────────────────

func (m Model) renderHome(width int) string {
	var b strings.Builder

	// ASCII Art Name — with reveal animation
	asciiLines := []string{
		"   ███████╗ █████╗ ██████╗ ██╗  ██╗ █████╗ ███╗   ██╗",
		"   ██╔════╝██╔══██╗██╔══██╗██║  ██║██╔══██╗████╗  ██║",
		"   █████╗  ███████║██████╔╝███████║███████║██╔██╗ ██║",
		"   ██╔══╝  ██╔══██║██╔══██╗██╔══██║██╔══██║██║╚██╗██║",
		"   ██║     ██║  ██║██║  ██║██║  ██║██║  ██║██║ ╚████║",
		"   ╚═╝     ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝",
	}

	b.WriteString("\n")

	linesToShow := m.homeRevealed
	if linesToShow > len(asciiLines) {
		linesToShow = len(asciiLines)
	}

	for i := 0; i < linesToShow; i++ {
		styled := lipgloss.NewStyle().
			Foreground(styles.Zinc100).
			Bold(true).
			Render(asciiLines[i])
		b.WriteString(styled)
		b.WriteString("\n")
	}

	// Fill remaining ASCII lines with empty space during animation
	for i := linesToShow; i < len(asciiLines); i++ {
		b.WriteString("\n")
	}

	b.WriteString("\n")

	// Subtitle with typing effect
	subtitle := m.cv.Basics.Label
	charsToShow := m.homeTyped
	if charsToShow > len(subtitle) {
		charsToShow = len(subtitle)
	}

	typedText := subtitle[:charsToShow]
	cursor := ""
	if charsToShow < len(subtitle) && m.showCursor {
		cursor = "█"
	}

	subtitleStyled := lipgloss.NewStyle().
		Foreground(styles.Zinc400).
		Render("   " + typedText + cursor)
	b.WriteString(subtitleStyled)
	b.WriteString("\n\n")

	// Thin divider
	b.WriteString("   ")
	b.WriteString(styles.ThinDivider(width - 6))
	b.WriteString("\n\n")

	// Summary card — clean, minimal
	summaryCard := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(styles.Zinc800).
		Padding(1, 2).
		MarginLeft(2).
		Width(width - 2).
		Render(
			styles.DimTextStyle.Render("   \"") +
				styles.TextStyle.Render(m.cv.Basics.Summary) +
				styles.DimTextStyle.Render("\""),
		)
	b.WriteString(summaryCard)
	b.WriteString("\n\n")

	// Quick stats with reveal animation
	b.WriteString("  ")
	b.WriteString(styles.SectionIconStyle.Render("◆ "))
	b.WriteString(styles.SectionTitleStyle.Render("Quick Overview"))
	b.WriteString("\n")

	stats := []struct {
		label string
		value string
	}{
		{"Projects", fmt.Sprintf("%d shipped products", len(m.cv.Projects))},
		{"Experience", fmt.Sprintf("%d roles", len(m.cv.Work))},
		{"Skills", fmt.Sprintf("%d+ technologies", len(m.cv.Skills))},
		{"Achievements", fmt.Sprintf("%d awards", len(m.cv.Achievements))},
		{"Community", fmt.Sprintf("%d orgs", len(m.cv.ExperiencesInOrganization))},
	}

	statsToShow := m.statsRevealed
	if statsToShow > len(stats) {
		statsToShow = len(stats)
	}

	for i := 0; i < statsToShow; i++ {
		s := stats[i]
		line := fmt.Sprintf("   %s %s  %s",
			styles.BulletStyle.Render("▸"),
			styles.BoldTextStyle.Render(s.label),
			styles.DimTextStyle.Render(s.value),
		)
		b.WriteString(line)
		b.WriteString("\n")
	}

	b.WriteString("\n")

	// Navigation hint
	hint := lipgloss.NewStyle().
		Foreground(styles.Zinc600).
		Italic(true).
		Render("   Use ← → or Tab to explore sections")
	b.WriteString(hint)
	b.WriteString("\n")

	return b.String()
}

// ── Page: About ────────────────────────────────────────────────────────

func (m Model) renderAbout(width int) string {
	var b strings.Builder

	b.WriteString("\n")
	b.WriteString("  ")
	b.WriteString(styles.SectionIconStyle.Render("● "))
	b.WriteString(styles.SectionTitleStyle.Render("About Me"))
	b.WriteString("\n\n")

	// Bio card
	bioCard := styles.CardStyle.Width(width - 2).Render(
		styles.TextStyle.Render(m.cv.Basics.Summary),
	)
	b.WriteString(bioCard)
	b.WriteString("\n\n")

	// Education
	b.WriteString("  ")
	b.WriteString(styles.SectionIconStyle.Render("▸ "))
	b.WriteString(styles.BoldTextStyle.Render("Education"))
	b.WriteString("\n\n")

	for _, edu := range m.cv.Education {
		b.WriteString(fmt.Sprintf("    %s  %s\n",
			styles.BulletStyle.Render("▸"),
			styles.BoldTextStyle.Render(edu.Institution),
		))
		b.WriteString(fmt.Sprintf("       %s  %s\n",
			styles.PositionStyle.Render(edu.Area),
			styles.DateStyle.Render(edu.StartDate+" — "+edu.EndDate),
		))
		b.WriteString(fmt.Sprintf("       %s\n\n",
			styles.DimTextStyle.Render(edu.StudyType),
		))
	}

	// Certifications
	b.WriteString("  ")
	b.WriteString(styles.SectionIconStyle.Render("▸ "))
	b.WriteString(styles.BoldTextStyle.Render("Certifications"))
	b.WriteString("\n\n")

	for _, cert := range m.cv.Certifications {
		b.WriteString(fmt.Sprintf("    %s  %s  %s  %s\n",
			styles.BulletStyle.Render("▸"),
			styles.BoldTextStyle.Render(cert.Name),
			styles.DateStyle.Render("("+cert.Date+")"),
			styles.StackStyle.Render("Score: "+cert.Score),
		))
	}
	b.WriteString("\n")

	// Talks
	if len(m.cv.Talks) > 0 {
		b.WriteString("  ")
		b.WriteString(styles.SectionIconStyle.Render("▸ "))
		b.WriteString(styles.BoldTextStyle.Render("Speaking"))
		b.WriteString("\n\n")

		for _, talk := range m.cv.Talks {
			talkCard := styles.HighlightCardStyle.Width(width - 2).Render(
				styles.BoldTextStyle.Render(talk.Title) + "\n" +
					styles.PositionStyle.Render(talk.Event) + " · " +
					styles.DateStyle.Render(talk.Date) + "\n\n" +
					styles.DimTextStyle.Render(talk.Summary),
			)
			b.WriteString(talkCard)
			b.WriteString("\n")
		}
	}

	// Organizations
	b.WriteString("\n")
	b.WriteString("  ")
	b.WriteString(styles.SectionIconStyle.Render("▸ "))
	b.WriteString(styles.BoldTextStyle.Render("Community & Organizations"))
	b.WriteString("\n\n")

	for _, org := range m.cv.ExperiencesInOrganization {
		b.WriteString(fmt.Sprintf("    %s  %s  %s\n",
			styles.BulletStyle.Render("▸"),
			styles.CompanyStyle.Render(org.Organization),
			styles.DateStyle.Render(org.StartDate+" — "+org.EndDate),
		))
		b.WriteString(fmt.Sprintf("       %s\n",
			styles.PositionStyle.Render(org.Position),
		))
		if org.Summary != "" {
			b.WriteString(fmt.Sprintf("       %s\n",
				styles.DimTextStyle.Render(org.Summary),
			))
		}
		b.WriteString("\n")
	}

	return b.String()
}

// ── Page: Experience ───────────────────────────────────────────────────

func (m Model) renderExperience(width int) string {
	var b strings.Builder

	b.WriteString("\n")
	b.WriteString("  ")
	b.WriteString(styles.SectionIconStyle.Render("▶ "))
	b.WriteString(styles.SectionTitleStyle.Render("Work Experience"))
	b.WriteString("\n\n")

	for i, work := range m.cv.Work {
		// Build card content with timeline inside
		var cardContent strings.Builder

		// Company + Date header
		cardContent.WriteString(styles.CompanyStyle.Render(work.Company))
		cardContent.WriteString("  ")
		cardContent.WriteString(styles.DateStyle.Render(work.StartDate + " — " + work.EndDate))
		cardContent.WriteString("\n")

		// Position
		cardContent.WriteString(styles.PositionStyle.Render(work.Position))
		cardContent.WriteString("\n")

		// Highlights
		for _, hl := range work.Highlights {
			wrapped := wordWrap(hl, width-14)
			lines := strings.Split(wrapped, "\n")
			cardContent.WriteString(fmt.Sprintf("\n%s %s",
				styles.BulletStyle.Render("▸"),
				styles.TextStyle.Render(lines[0]),
			))
			for _, line := range lines[1:] {
				cardContent.WriteString(fmt.Sprintf("\n  %s",
					styles.TextStyle.Render(line),
				))
			}
		}

		// Use highlighted card for first (current) role, normal for rest
		var card string
		if i == 0 {
			card = styles.HighlightCardStyle.Width(width - 2).Render(cardContent.String())
		} else {
			card = styles.CardStyle.Width(width - 2).Render(cardContent.String())
		}
		b.WriteString(card)
		b.WriteString("\n\n")
	}

	return b.String()
}

// ── Page: Projects ─────────────────────────────────────────────────────

func (m Model) renderProjects(width int) string {
	var b strings.Builder

	b.WriteString("\n")
	b.WriteString("  ")
	b.WriteString(styles.SectionIconStyle.Render("◈ "))
	b.WriteString(styles.SectionTitleStyle.Render("Projects"))
	b.WriteString("\n\n")

	for _, proj := range m.cv.Projects {
		var cardContent strings.Builder

		// Name with emerald accent
		nameAndDate := styles.ProjectNameStyle.Render(proj.Name)
		if proj.EndDate != "" {
			nameAndDate += "  " + styles.DateStyle.Render(proj.Date+" — "+proj.EndDate)
		} else {
			nameAndDate += "  " + styles.DateStyle.Render(proj.Date)
		}
		cardContent.WriteString(nameAndDate)
		cardContent.WriteString("\n")

		// URL
		if proj.URL != "" {
			cardContent.WriteString(styles.LinkStyle.Render(proj.URL))
			cardContent.WriteString("\n")
		}
		cardContent.WriteString("\n")

		// Summary
		if proj.Summary != "" {
			wrapped := wordWrap(proj.Summary, width-12)
			cardContent.WriteString(styles.TextStyle.Render(wrapped))
			cardContent.WriteString("\n")
		}

		// Highlights
		if len(proj.Highlights) > 0 {
			cardContent.WriteString("\n")
			for _, hl := range proj.Highlights {
				wrapped := wordWrap(hl, width-14)
				lines := strings.Split(wrapped, "\n")
				cardContent.WriteString(fmt.Sprintf("%s %s\n",
					styles.BulletStyle.Render("▸"),
					styles.DimTextStyle.Render(lines[0]),
				))
				for _, line := range lines[1:] {
					cardContent.WriteString(fmt.Sprintf("  %s\n",
						styles.DimTextStyle.Render(line),
					))
				}
			}
		}

		// Stack
		if proj.Stack != "" {
			cardContent.WriteString("\n")
			cardContent.WriteString(styles.StackStyle.Render("⚙ " + proj.Stack))
		}

		card := styles.CardStyle.Width(width - 2).Render(cardContent.String())
		b.WriteString(card)
		b.WriteString("\n\n")
	}

	return b.String()
}

// ── Page: Skills ───────────────────────────────────────────────────────

func (m Model) renderSkills(width int) string {
	var b strings.Builder

	b.WriteString("\n")
	b.WriteString("  ")
	b.WriteString(styles.SectionIconStyle.Render("⬟ "))
	b.WriteString(styles.SectionTitleStyle.Render("Skills & Technologies"))
	b.WriteString("\n\n")

	// Categorize skills
	categories := map[string][]string{
		"Languages":      {},
		"Frontend":       {},
		"Backend":        {},
		"Databases":      {},
		"Cloud & DevOps": {},
		"AI/ML":          {},
	}

	categoryOrder := []string{"Languages", "Frontend", "Backend", "Databases", "Cloud & DevOps", "AI/ML"}

	langKeywords := []string{"JavaScript", "TypeScript", "Python", "HTML", "CSS", "GraphQL"}
	frontendKeywords := []string{"Svelte", "SvelteKit", "Next.js", "React"}
	backendKeywords := []string{"FastAPI", "Node.js"}
	dbKeywords := []string{"MongoDB", "PostgreSQL", "Supabase", "Redis", "ClickHouse", "Kafka"}
	cloudKeywords := []string{"Docker", "AWS", "GCP", "Cloudflare", "DigitalOcean", "Terraform", "CI/CD", "Git", "Linux"}
	aiKeywords := []string{"LiteLLM", "LangChain", "Strands", "AI/ML"}

	for _, skill := range m.cv.Skills {
		placed := false
		for _, kw := range langKeywords {
			if strings.Contains(skill, kw) {
				categories["Languages"] = append(categories["Languages"], skill)
				placed = true
				break
			}
		}
		if placed {
			continue
		}
		for _, kw := range frontendKeywords {
			if strings.Contains(skill, kw) {
				categories["Frontend"] = append(categories["Frontend"], skill)
				placed = true
				break
			}
		}
		if placed {
			continue
		}
		for _, kw := range backendKeywords {
			if strings.Contains(skill, kw) {
				categories["Backend"] = append(categories["Backend"], skill)
				placed = true
				break
			}
		}
		if placed {
			continue
		}
		for _, kw := range dbKeywords {
			if strings.Contains(skill, kw) {
				categories["Databases"] = append(categories["Databases"], skill)
				placed = true
				break
			}
		}
		if placed {
			continue
		}
		for _, kw := range cloudKeywords {
			if strings.Contains(skill, kw) {
				categories["Cloud & DevOps"] = append(categories["Cloud & DevOps"], skill)
				placed = true
				break
			}
		}
		if placed {
			continue
		}
		for _, kw := range aiKeywords {
			if strings.Contains(skill, kw) {
				categories["AI/ML"] = append(categories["AI/ML"], skill)
				placed = true
				break
			}
		}
		if !placed {
			categories["Languages"] = append(categories["Languages"], skill)
		}
	}

	catIcons := map[string]string{
		"Languages":      "◆",
		"Frontend":       "◇",
		"Backend":        "▣",
		"Databases":      "◈",
		"Cloud & DevOps": "☁",
		"AI/ML":          "◎",
	}

	for _, cat := range categoryOrder {
		skills := categories[cat]
		if len(skills) == 0 {
			continue
		}

		icon := catIcons[cat]

		catTitle := "  " +
			styles.SectionIconStyle.Render(icon+" ") +
			styles.BoldTextStyle.Render(cat)
		b.WriteString(catTitle)
		b.WriteString("\n")

		// Render skills as monochrome tags
		line := "     "
		lineLen := 5
		for i, skill := range skills {
			tag := styles.SkillTagStyle.Render(skill)
			tagLen := len(skill) + 2
			sep := "  "
			if i == len(skills)-1 {
				sep = ""
			}

			if lineLen+tagLen+2 > width-4 {
				b.WriteString(line)
				b.WriteString("\n")
				line = "     "
				lineLen = 5
			}
			line += tag + sep
			lineLen += tagLen + 2
		}
		b.WriteString(line)
		b.WriteString("\n\n")
	}

	// Total count
	totalBadge := lipgloss.NewStyle().
		Foreground(styles.Zinc500).
		Italic(true).
		Render(fmt.Sprintf("  Total: %d+ technologies and growing", len(m.cv.Skills)))
	b.WriteString(totalBadge)
	b.WriteString("\n")

	return b.String()
}

// ── Page: Achievements ─────────────────────────────────────────────────

func (m Model) renderAchievements(width int) string {
	var b strings.Builder

	b.WriteString("\n")
	b.WriteString("  ")
	b.WriteString(styles.SectionIconStyle.Render("★ "))
	b.WriteString(styles.SectionTitleStyle.Render("Achievements & Awards"))
	b.WriteString("\n\n")

	for _, ach := range m.cv.Achievements {
		var cardContent strings.Builder

		// Icon based on placement
		icon := "◆"
		if strings.Contains(ach.Title, "1st") {
			icon = "★"
		} else if strings.Contains(ach.Title, "4th") || strings.Contains(ach.Title, "Finalist") {
			icon = "◇"
		} else if strings.Contains(ach.Title, "Best") {
			icon = "●"
		}

		cardContent.WriteString(fmt.Sprintf("%s  %s\n",
			styles.SectionIconStyle.Render(icon),
			styles.AchievementStyle.Render(ach.Title),
		))
		cardContent.WriteString(fmt.Sprintf("   %s\n",
			styles.DateStyle.Render(ach.Date),
		))

		if ach.Summary != "" {
			cardContent.WriteString(fmt.Sprintf("\n   %s\n",
				styles.TextStyle.Render(ach.Summary),
			))
		}

		if len(ach.Highlights) > 0 {
			cardContent.WriteString("\n")
			for _, hl := range ach.Highlights {
				wrapped := wordWrap(hl, width-14)
				lines := strings.Split(wrapped, "\n")
				cardContent.WriteString(fmt.Sprintf("   %s %s\n",
					styles.BulletStyle.Render("▸"),
					styles.DimTextStyle.Render(lines[0]),
				))
				for _, line := range lines[1:] {
					cardContent.WriteString(fmt.Sprintf("     %s\n",
						styles.DimTextStyle.Render(line),
					))
				}
			}
		}

		if ach.Stack != "" {
			cardContent.WriteString(fmt.Sprintf("\n   %s\n",
				styles.StackStyle.Render("⚙ "+ach.Stack),
			))
		}

		card := styles.HighlightCardStyle.Width(width - 2).Render(cardContent.String())
		b.WriteString(card)
		b.WriteString("\n\n")
	}

	return b.String()
}

// ── Page: Contact ──────────────────────────────────────────────────────

func (m Model) renderContact(width int) string {
	var b strings.Builder

	b.WriteString("\n")
	b.WriteString("  ")
	b.WriteString(styles.SectionIconStyle.Render("◉ "))
	b.WriteString(styles.SectionTitleStyle.Render("Get In Touch"))
	b.WriteString("\n\n")

	// Contact info card
	contactInfo := fmt.Sprintf(
		"%s  %s\n\n%s  %s\n\n%s  %s\n\n%s  %s",
		styles.BulletStyle.Render("✉"),
		styles.LinkStyle.Render(m.cv.Basics.Email),
		styles.BulletStyle.Render("☎"),
		styles.TextStyle.Render(m.cv.Basics.Phone),
		styles.BulletStyle.Render("◆"),
		styles.LinkStyle.Render(m.cv.Basics.Website),
		styles.BulletStyle.Render("⌂"),
		styles.TextStyle.Render("Indonesia"),
	)

	contactCard := styles.HighlightCardStyle.Width(width - 2).Render(contactInfo)
	b.WriteString(contactCard)
	b.WriteString("\n\n")

	// Social links
	b.WriteString("  ")
	b.WriteString(styles.SectionIconStyle.Render("▸ "))
	b.WriteString(styles.BoldTextStyle.Render("Social Links"))
	b.WriteString("\n\n")

	socialIcons := map[string]string{
		"LinkedIn":  "in",
		"GitHub":    "gh",
		"Blog":      "bg",
		"Instagram": "ig",
		"Medium":    "md",
	}

	for _, social := range m.cv.Socials {
		icon := socialIcons[social.Network]
		if icon == "" {
			icon = "◆"
		}

		badge := lipgloss.NewStyle().
			Foreground(styles.Zinc950).
			Background(styles.Zinc300).
			Bold(true).
			Padding(0, 1).
			Render(icon)

		line := fmt.Sprintf("    %s  %s  %s",
			badge,
			styles.BoldTextStyle.Render(social.Network),
			styles.LinkStyle.Render(social.URL),
		)
		b.WriteString(line)
		b.WriteString("\n\n")
	}

	// Farewell
	b.WriteString("\n")
	farewell := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(styles.Zinc700).
		Padding(1, 3).
		MarginLeft(2).
		Width(width - 2).
		Align(lipgloss.Center).
		Render(
			styles.BoldTextStyle.Render("Thanks for visiting via SSH!") + "\n" +
				styles.DimTextStyle.Render("Let's build something awesome together.") + "\n\n" +
				styles.TextStyle.Render("— Farhan Aulianda"),
		)
	b.WriteString(farewell)
	b.WriteString("\n")

	return b.String()
}

// ── Helpers ────────────────────────────────────────────────────────────

func wordWrap(text string, width int) string {
	if width <= 0 {
		return text
	}

	words := strings.Fields(text)
	if len(words) == 0 {
		return ""
	}

	var lines []string
	currentLine := words[0]

	for _, word := range words[1:] {
		if len(currentLine)+1+len(word) > width {
			lines = append(lines, currentLine)
			currentLine = word
		} else {
			currentLine += " " + word
		}
	}
	lines = append(lines, currentLine)

	return strings.Join(lines, "\n")
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
