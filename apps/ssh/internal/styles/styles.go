package styles

import (
	"strings"

	"charm.land/lipgloss/v2"
)

// ── Color Palette ── Zinc Monochrome + Emerald (matching website) ──────
var (
	// Zinc scale (main palette)
	Zinc950 = lipgloss.Color("#09090b") // Deepest background
	Zinc900 = lipgloss.Color("#18181b") // Card/panel background
	Zinc800 = lipgloss.Color("#27272a") // Borders, dividers
	Zinc700 = lipgloss.Color("#3f3f46") // Subtle borders
	Zinc600 = lipgloss.Color("#52525b") // Line numbers, muted
	Zinc500 = lipgloss.Color("#71717a") // Dim text, dates
	Zinc400 = lipgloss.Color("#a1a1aa") // Secondary text
	Zinc300 = lipgloss.Color("#d4d4d8") // Body text
	Zinc200 = lipgloss.Color("#e4e4e7") // Bright text
	Zinc100 = lipgloss.Color("#f4f4f5") // Headings
	Zinc50  = lipgloss.Color("#fafafa") // Pure white text

	// Accent — Emerald (sparingly)
	Emerald500 = lipgloss.Color("#10b981")
	Emerald400 = lipgloss.Color("#34d399")
	Emerald300 = lipgloss.Color("#6ee7b7")

	// Danger for glitch
	Red500 = lipgloss.Color("#ef4444")
	Red400 = lipgloss.Color("#f87171")

	// Aliases for backward compat in rendering
	Primary   = Zinc100
	Secondary = Zinc400
	Accent    = Emerald500
	Muted     = Zinc500
	TextLight = Zinc200
	TextDim   = Zinc500
	BgDark    = Zinc950
	BgCard    = Zinc900
	Border    = Zinc800
)

// ── Base Styles ────────────────────────────────────────────────────────
var (
	// Logo / Name — clean white, bold
	LogoStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(Zinc50)

	// Subtitle
	SubtitleStyle = lipgloss.NewStyle().
			Foreground(Zinc400).
			Italic(true)

	// Section title — emerald accent, sparingly
	SectionTitleStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(Zinc100).
				Padding(0, 0, 1, 0)

	// Section icon (the small icon before section titles)
	SectionIconStyle = lipgloss.NewStyle().
				Foreground(Emerald500)

	// Normal text
	TextStyle = lipgloss.NewStyle().
			Foreground(Zinc300)

	// Dimmed text
	DimTextStyle = lipgloss.NewStyle().
			Foreground(Zinc500)

	// Bold text
	BoldTextStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(Zinc200)

	// Link style — emerald underline
	LinkStyle = lipgloss.NewStyle().
			Foreground(Emerald400).
			Underline(true)

	// Card style with border
	CardStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(Zinc800).
			Padding(1, 2).
			MarginLeft(2)

	// Highlighted card — slightly brighter border
	HighlightCardStyle = lipgloss.NewStyle().
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(Zinc700).
				Padding(1, 2).
				MarginLeft(2)

	// Active tab — white on zinc
	ActiveTabStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(Zinc950).
			Background(Zinc200).
			Padding(0, 2)

	// Inactive tab
	InactiveTabStyle = lipgloss.NewStyle().
				Foreground(Zinc500).
				Background(Zinc900).
				Padding(0, 2)

	// Bullet point — emerald
	BulletStyle = lipgloss.NewStyle().
			Foreground(Emerald500)

	// Date range style
	DateStyle = lipgloss.NewStyle().
			Foreground(Zinc500).
			Italic(true)

	// Divider
	DividerStyle = lipgloss.NewStyle().
			Foreground(Zinc800)

	// Scrollbar indicator
	ScrollStyle = lipgloss.NewStyle().
			Foreground(Zinc400)

	// Achievement highlight — bright white
	AchievementStyle = lipgloss.NewStyle().
				Foreground(Zinc100).
				Bold(true)

	// Company name — white, bold
	CompanyStyle = lipgloss.NewStyle().
			Foreground(Zinc100).
			Bold(true)

	// Position title — zinc-400
	PositionStyle = lipgloss.NewStyle().
			Foreground(Zinc400)

	// Project name — emerald accent
	ProjectNameStyle = lipgloss.NewStyle().
				Foreground(Emerald400).
				Bold(true)

	// Stack/tech — zinc-400 italic
	StackStyle = lipgloss.NewStyle().
			Foreground(Zinc400).
			Italic(true)

	// Tag style for skills
	SkillTagStyle = lipgloss.NewStyle().
			Foreground(Zinc300).
			Background(Zinc800).
			Padding(0, 1)

	// Glitch text style
	GlitchStyle = lipgloss.NewStyle().
			Foreground(Red500).
			Bold(true)

	// Status bar key hints
	KeyHintStyle = lipgloss.NewStyle().
			Foreground(Zinc600)

	// Status bar active key
	KeyActiveStyle = lipgloss.NewStyle().
			Foreground(Zinc400).
			Bold(true)
)

// ── Helper Functions ───────────────────────────────────────────────────

func Divider(width int) string {
	return DividerStyle.Render(strings.Repeat("─", width))
}

func DoubleDivider(width int) string {
	return DividerStyle.Render(strings.Repeat("═", width))
}

func DottedDivider(width int) string {
	return DividerStyle.Render(strings.Repeat("·", width))
}

func ThinDivider(width int) string {
	return lipgloss.NewStyle().Foreground(Zinc800).Render(strings.Repeat("╌", width))
}

func RenderTab(label string, active bool) string {
	if active {
		return ActiveTabStyle.Render(label)
	}
	return InactiveTabStyle.Render(label)
}

// GradientBar renders a simple gradient-like bar using block characters
func GradientBar(width int, progress float64) string {
	filled := int(float64(width) * progress)
	if filled > width {
		filled = width
	}

	bar := lipgloss.NewStyle().Foreground(Emerald500).Render(strings.Repeat("█", filled))
	empty := lipgloss.NewStyle().Foreground(Zinc800).Render(strings.Repeat("░", width-filled))
	return bar + empty
}
