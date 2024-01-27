package systray

import "github.com/godbus/dbus/v5"

// SetTemplateIcon sets the systray icon as a template icon (on macOS), falling back
// to a regular icon on other platforms.
// templateIconBytes and iconBytes should be the content of .ico for windows and
// .ico/.jpg/.png for other platforms.
func SetTemplateIcon(templateIconBytes []byte, regularIconBytes []byte) {
	SetIcon(regularIconBytes)
}

// SetRemovalAllowed sets whether a user can remove the systray icon or not.
// This is only supported on macOS.
func SetRemovalAllowed(allowed bool) {
}

// SetIcon sets the icon of a menu item. Only works on macOS and Windows.
// iconBytes should be the content of .ico/.jpg/.png
func (item *MenuItem) SetIcon(iconBytes []byte) {
}

// SetTemplateIcon sets the icon of a menu item as a template icon (on macOS). On Windows, it
// falls back to the regular icon bytes and on Linux it does nothing.
// templateIconBytes and regularIconBytes should be the content of .ico for windows and
// .ico/.jpg/.png for other platforms.
func (item *MenuItem) SetTemplateIcon(templateIconBytes []byte, regularIconBytes []byte) {
}

// ShowMessage shows a notification on the end user's desktop
func ShowMessage(appName, title, msg string) {
	conn, err := dbus.SessionBus()
	if err != nil {
		log.Errorf("Unable to obtain dbus session: %v", err)
		return
	}

	obj := conn.Object(
		"org.freedesktop.Notifications",
		dbus.ObjectPath("/org/freedesktop/Notifications"),
	)

	call := obj.Call(
		"org.freedesktop.Notifications.Notify",
		0,
		appName,
		uint32(0),                 // replaces_id
		"",                        // app_icon
		title,                     // summary
		msg,                       // body
		[]string{},                // actions
		map[string]dbus.Variant{}, // hints
		int32(-1),                 // expire_timeout (msec)
	)
	if call.Err != nil {
		log.Errorf("Unable to send desktop notification: %v", call.Err)
		return
	}

	return
}
