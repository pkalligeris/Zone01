package model

// Config represents the runtime configuration for the ASCII art generator.
type Config struct {
	// Input is the string to be rendered.
	Input string

	// BannerFile is the name of the banner to use (e.g., "standard", "shadow").
	// Defaults to "standard" if not specified.
	BannerFile string

	// Color is the color to apply to the output (e.g., "red", "blue").
	// If empty, no color is applied.
	Color string

	// ColorSubstr is the specific substring to colorize.
	// If empty and Color is set, the whole string is colored.
	ColorSubstr string

	// OutputFile is the path where the output should be written.
	// If empty, output is written to stdout.
	OutputFile string

	// Align specifies the text alignment ("left", "center", "right", "justify").
	// Defaults to "left".
	Align string
}