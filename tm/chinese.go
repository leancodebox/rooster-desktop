package tm

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"github.com/leancodebox/rooster-desktop/resource"
	"image/color"
)

type MyTheme struct{}

var _ fyne.Theme = (*MyTheme)(nil)

func (m MyTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	//return theme.DefaultTheme().Color(n, v)
	return theme.DefaultTheme().Color(n, v)
}
func (m MyTheme) Icon(name fyne.ThemeIconName) fyne.Resource {

	return theme.DefaultTheme().Icon(name)
}

func (m MyTheme) Font(style fyne.TextStyle) fyne.Resource {
	//return theme.DefaultTheme().Font(style)
	return &fyne.StaticResource{
		StaticName:    "HarmonyOS_Sans_SC_Regular.ttf",
		StaticContent: resource.HMttf,
	}
}

func (m MyTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
