package d2ui

// Widget defines an object that is a UI widget
type Widget interface {
	Drawable
	bindManager(ui *UIManager)
}

// ClickableWidget defines an object that can be clicked
type ClickableWidget interface {
	Widget
	SetEnabled(enabled bool)
	SetPressed(pressed bool)
	GetEnabled() bool
	GetPressed() bool
	OnActivated(callback func())
	Activate()
}

// BaseWidget contains default functionality that all widgets share
type BaseWidget struct {
	manager *UIManager
	x       int
	y       int
	width   int
	height  int
	visible bool
}

// NewBaseWidget creates a new BaseWidget with defaults
func NewBaseWidget(manager *UIManager) *BaseWidget {
	return &BaseWidget{
		manager: manager,
		x:       0,
		y:       0,
		width:   0,
		height:  0,
		visible: true,
	}
}

func (b *BaseWidget) bindManager(manager *UIManager) {
	b.manager = manager
}

// GetSize returns the size of the widget
func (b *BaseWidget) GetSize() (width, height int) {
	return b.width, b.height
}

// SetPosition sets the position of the widget
func (b *BaseWidget) SetPosition(x, y int) {
	b.x, b.y = x, y
}

// OffsetPosition moves the widget by x and y
func (b *BaseWidget) OffsetPosition(x, y int) {
	b.x += x
	b.y += y
}

// GetPosition returns the position of the widget
func (b *BaseWidget) GetPosition() (x, y int) {
	return b.x, b.y
}

// GetVisible returns whether the widget is visible
func (b *BaseWidget) GetVisible() (visible bool) {
	return b.visible
}

// SetVisible make the widget visible, not visible
func (b *BaseWidget) SetVisible(visible bool) {
	b.visible = visible
}
