// +build !ci,!gl

package efl

func oSEngineName() string {
	return "opengl_cocoa"
}

func oSWindowInit(w *window) {
}