package utils

// Color - Color Mapping
type Color struct {
	Reset  string
	Red    string
	Green  string
	Yellow string
	Blue   string
	Purple string
	Cyan   string
	Gray   string
	White  string
}

// Init - Constructor for Color
func (c *Color) Init() {
	c.Reset = "\033[0m"
	c.Red = "\033[31m"
	c.Green = "\033[32m"
	c.Yellow = "\033[33m"
	c.Blue = "\033[34m"
	c.Purple = "\033[35m"
	c.Cyan = "\033[36m"
	c.Gray = "\033[37m"
	c.White = "\033[97m"
}
