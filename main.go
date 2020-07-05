package main
import(
    "fmt"
    "github.com/tfriedel6/canvas"
    "github.com/tfriedel6/canvas/sdlcanvas"
    "github.com/tfriedel6/canvas/backend/softwarebackend"
)
func main() {
    // The size i would like my window to be
    screenW := 1600
    screenH := 900
    // The size i would like the draggable square in the center to be.
    squareSize := 128
    wnd, cv, err := sdlcanvas.CreateWindow(screenW, screenH, "Example")
    if err != nil {
        panic(err)
    }
    fmt.Printf("Desired: %v w x %v h\n", screenW, screenH)
    fmt.Printf("Actual: %v w x %v h\n", cv.Width(), cv.Height())
    // Create a square. This square will be displayed in the center
    // of the window and should be draggable.
    backend := softwarebackend.New(squareSize, squareSize)
    cvImg := canvas.New(backend)
    cvImg.SetFillStyle("#FFF")
    cvImg.Rect(0, 0, 64, 64)
    cvImg.Fill()
    img := cvImg.GetImageData(0, 0, squareSize, squareSize)
    // Load the square into the main canvas as an image.
    imgFinal, err := cv.LoadImage(img)
    if err != nil {
        panic(err)
    }
    // Starting position for the square
    xStartPos := (screenW / 2) - (squareSize / 2)
    yStartPos := (screenH / 2) - (squareSize / 2)
    // Active position for the square
    xPos := xStartPos
    yPos := yStartPos
    // Basic properties for tracking the cursor
    followMouse := false
    // When the mouse is released, stop following it and 
    // snap the square back to it's original pos.
    wnd.MouseUp = func(b, x, y int) {
        followMouse = false
        xPos = xStartPos
        yPos = yStartPos
    }
    // Only begin following the cursor if the square is clicked.
    wnd.MouseDown = func(b, x, y int) {
        if x > xStartPos && x < xStartPos + squareSize {
            if y > yStartPos && y < yStartPos + squareSize {
                followMouse = true
            }
        }
    }
    // When the mouse is moved, if we should be following the cursor
    // update the location of the square to match that of the cursor.
    wnd.MouseMove = func(x, y int) {
        if followMouse {
            xPos = x - (squareSize / 2)
            yPos = y - (squareSize / 2)
        }
    }
    defer wnd.Destroy()
    // Start the main loop for rendering.
    wnd.MainLoop(func() {
        // I've been able to determine that for some reason, my canvas
        // is not the same size as the initial values i had originally
        // passed to sdlcanvas.CreateWindow().
        // This seems to properly fill the entire canvas
        w, h := float64(cv.Width()), float64(cv.Height())
        cv.SetFillStyle("#fff")
        cv.FillRect(0, 0, w, h)
        // This does not fill the entire canvas.
        w2, h2 := float64(screenW), float64(screenH)
        cv.SetFillStyle("#000")
        cv.FillRect(0, 0, w2, h2)
        // This should be in the center of the window when it is first
        // spawned but instead is is in the center of the top left
        // quadrant of the window, which happens to be the center of
        // the true canvas.
        cv.DrawImage(imgFinal, float64(xPos), 
                               float64(yPos), 
                               float64(squareSize), 
                               float64(squareSize))
    })
}