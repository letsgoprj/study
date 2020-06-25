package main
import "github.com/fogleman/gg"
func main() {
    const S = 256
    dc := gg.NewContext(S, S)
    dc.SetRGB(1, 1, 1)
    dc.Clear()
    if err := dc.LoadFontFace("font.ttf", 36); err != nil {
        panic(err)
    }
    dc.SetRGB(0, 0, 0)
    s := "한글ABC"
    n := 3 // "stroke" size
    for dy := -n; dy <= n; dy++ {
        for dx := -n; dx <= n; dx++ {
            if dx*dx+dy*dy >= n*n {
                // give it rounded corners
                continue
            }
            x := S/2 + float64(dx)
            y := S/2 + float64(dy)
            dc.DrawStringAnchored(s, x, y, 0.5, 0.5)
        }
    }
    dc.SetRGB(1, 1, 0)
    dc.DrawStringAnchored(s, S/2, S/2, 0.5, 0.5)
    dc.SavePNG("out.png")
}
