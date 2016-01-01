package keyboardemulator

import "github.com/AllenDang/w32"

// Missing virtual keys from w32 package
const (
	Vk0 = 0x30
	Vk1 = 0x31
	Vk2 = 0x32
	Vk3 = 0x33
	Vk4 = 0x34
	Vk5 = 0x35
	Vk6 = 0x36
	Vk7 = 0x37
	Vk8 = 0x38
	Vk9 = 0x39
	VkA = 0x41
	VkB = 0x42
	VkC = 0x43
	VkD = 0x44
	VkE = 0x45
	VkF = 0x46
	VkG = 0x47
	VkH = 0x48
	VkI = 0x49
	VkJ = 0x4A
	VkK = 0x4B
	VkL = 0x4C
	VkM = 0x4D
	VkN = 0x4E
	VkO = 0x4F
	VkP = 0x50
	VkQ = 0x51
	VkR = 0x52
	VkS = 0x53
	VkT = 0x54
	VkU = 0x55
	VkV = 0x56
	VkW = 0x57
	VkX = 0x58
	VkY = 0x59
	VkZ = 0x5A
)

var (
	keymap = map[Key]uint16{
		Key0:                 Vk0,
		Key1:                 Vk1,
		Key2:                 Vk2,
		Key3:                 Vk3,
		Key4:                 Vk4,
		Key5:                 Vk5,
		Key6:                 Vk6,
		Key7:                 Vk7,
		Key8:                 Vk8,
		Key9:                 Vk9,
		KeyA:                 VkA,
		KeyB:                 VkB,
		KeyC:                 VkC,
		KeyD:                 VkD,
		KeyE:                 VkE,
		KeyF:                 VkF,
		KeyG:                 VkG,
		KeyH:                 VkH,
		KeyI:                 VkI,
		KeyJ:                 VkJ,
		KeyK:                 VkK,
		KeyL:                 VkL,
		KeyM:                 VkM,
		KeyN:                 VkN,
		KeyO:                 VkO,
		KeyP:                 VkP,
		KeyQ:                 VkQ,
		KeyR:                 VkR,
		KeyS:                 VkS,
		KeyT:                 VkT,
		KeyU:                 VkU,
		KeyV:                 VkV,
		KeyW:                 VkW,
		KeyX:                 VkX,
		KeyY:                 VkY,
		KeyZ:                 VkZ,
		KeyLbutton:           w32.VK_LBUTTON,
		KeyRbutton:           w32.VK_RBUTTON,
		KeyCancel:            w32.VK_CANCEL,
		KeyMbutton:           w32.VK_MBUTTON,
		KeyXbutton1:          w32.VK_XBUTTON1,
		KeyXbutton2:          w32.VK_XBUTTON2,
		KeyBack:              w32.VK_BACK,
		KeyTab:               w32.VK_TAB,
		KeyClear:             w32.VK_CLEAR,
		KeyReturn:            w32.VK_RETURN,
		KeyShift:             w32.VK_SHIFT,
		KeyControl:           w32.VK_CONTROL,
		KeyAlt:               w32.VK_MENU,
		KeyPause:             w32.VK_PAUSE,
		KeyCapital:           w32.VK_CAPITAL,
		KeyKana:              w32.VK_KANA,
		KeyHangeul:           w32.VK_HANGEUL,
		KeyHangul:            w32.VK_HANGUL,
		KeyJunja:             w32.VK_JUNJA,
		KeyFinal:             w32.VK_FINAL,
		KeyHanja:             w32.VK_HANJA,
		KeyKanji:             w32.VK_KANJI,
		KeyEscape:            w32.VK_ESCAPE,
		KeyConvert:           w32.VK_CONVERT,
		KeyNonconvert:        w32.VK_NONCONVERT,
		KeyAccept:            w32.VK_ACCEPT,
		KeyModechange:        w32.VK_MODECHANGE,
		KeySpace:             w32.VK_SPACE,
		KeyPrior:             w32.VK_PRIOR,
		KeyNext:              w32.VK_NEXT,
		KeyEnd:               w32.VK_END,
		KeyHome:              w32.VK_HOME,
		KeyLeft:              w32.VK_LEFT,
		KeyUp:                w32.VK_UP,
		KeyRight:             w32.VK_RIGHT,
		KeyDown:              w32.VK_DOWN,
		KeySelect:            w32.VK_SELECT,
		KeyPrint:             w32.VK_PRINT,
		KeyExecute:           w32.VK_EXECUTE,
		KeySnapshot:          w32.VK_SNAPSHOT,
		KeyInsert:            w32.VK_INSERT,
		KeyDelete:            w32.VK_DELETE,
		KeyHelp:              w32.VK_HELP,
		KeyLwin:              w32.VK_LWIN,
		KeyRwin:              w32.VK_RWIN,
		KeyApps:              w32.VK_APPS,
		KeySleep:             w32.VK_SLEEP,
		KeyNumpad0:           w32.VK_NUMPAD0,
		KeyNumpad1:           w32.VK_NUMPAD1,
		KeyNumpad2:           w32.VK_NUMPAD2,
		KeyNumpad3:           w32.VK_NUMPAD3,
		KeyNumpad4:           w32.VK_NUMPAD4,
		KeyNumpad5:           w32.VK_NUMPAD5,
		KeyNumpad6:           w32.VK_NUMPAD6,
		KeyNumpad7:           w32.VK_NUMPAD7,
		KeyNumpad8:           w32.VK_NUMPAD8,
		KeyNumpad9:           w32.VK_NUMPAD9,
		KeyMultiply:          w32.VK_MULTIPLY,
		KeyAdd:               w32.VK_ADD,
		KeySeparator:         w32.VK_SEPARATOR,
		KeySubtract:          w32.VK_SUBTRACT,
		KeyDecimal:           w32.VK_DECIMAL,
		KeyDivide:            w32.VK_DIVIDE,
		KeyF1:                w32.VK_F1,
		KeyF2:                w32.VK_F2,
		KeyF3:                w32.VK_F3,
		KeyF4:                w32.VK_F4,
		KeyF5:                w32.VK_F5,
		KeyF6:                w32.VK_F6,
		KeyF7:                w32.VK_F7,
		KeyF8:                w32.VK_F8,
		KeyF9:                w32.VK_F9,
		KeyF10:               w32.VK_F10,
		KeyF11:               w32.VK_F11,
		KeyF12:               w32.VK_F12,
		KeyF13:               w32.VK_F13,
		KeyF14:               w32.VK_F14,
		KeyF15:               w32.VK_F15,
		KeyF16:               w32.VK_F16,
		KeyF17:               w32.VK_F17,
		KeyF18:               w32.VK_F18,
		KeyF19:               w32.VK_F19,
		KeyF20:               w32.VK_F20,
		KeyF21:               w32.VK_F21,
		KeyF22:               w32.VK_F22,
		KeyF23:               w32.VK_F23,
		KeyF24:               w32.VK_F24,
		KeyNumlock:           w32.VK_NUMLOCK,
		KeyScroll:            w32.VK_SCROLL,
		KeyOemNecEqual:       w32.VK_OEM_NEC_EQUAL,
		KeyOemFjJisho:        w32.VK_OEM_FJ_JISHO,
		KeyOemFjMasshou:      w32.VK_OEM_FJ_MASSHOU,
		KeyOemFjTouroku:      w32.VK_OEM_FJ_TOUROKU,
		KeyOemFjLoya:         w32.VK_OEM_FJ_LOYA,
		KeyOemFjRoya:         w32.VK_OEM_FJ_ROYA,
		KeyLshift:            w32.VK_LSHIFT,
		KeyRshift:            w32.VK_RSHIFT,
		KeyLcontrol:          w32.VK_LCONTROL,
		KeyRcontrol:          w32.VK_RCONTROL,
		KeyLmenu:             w32.VK_LMENU,
		KeyRmenu:             w32.VK_RMENU,
		KeyBrowserBack:       w32.VK_BROWSER_BACK,
		KeyBrowserForward:    w32.VK_BROWSER_FORWARD,
		KeyBrowserRefresh:    w32.VK_BROWSER_REFRESH,
		KeyBrowserStop:       w32.VK_BROWSER_STOP,
		KeyBrowserSearch:     w32.VK_BROWSER_SEARCH,
		KeyBrowserFavorites:  w32.VK_BROWSER_FAVORITES,
		KeyBrowserHome:       w32.VK_BROWSER_HOME,
		KeyVolumeMute:        w32.VK_VOLUME_MUTE,
		KeyVolumeDown:        w32.VK_VOLUME_DOWN,
		KeyVolumeUp:          w32.VK_VOLUME_UP,
		KeyMediaNextTrack:    w32.VK_MEDIA_NEXT_TRACK,
		KeyMediaPrevTrack:    w32.VK_MEDIA_PREV_TRACK,
		KeyMediaStop:         w32.VK_MEDIA_STOP,
		KeyMediaPlayPause:    w32.VK_MEDIA_PLAY_PAUSE,
		KeyLaunchMail:        w32.VK_LAUNCH_MAIL,
		KeyLaunchMediaSelect: w32.VK_LAUNCH_MEDIA_SELECT,
		KeyLaunchApp1:        w32.VK_LAUNCH_APP1,
		KeyLaunchApp2:        w32.VK_LAUNCH_APP2,
		KeyOem1:              w32.VK_OEM_1,
		KeyOemPlus:           w32.VK_OEM_PLUS,
		KeyOemComma:          w32.VK_OEM_COMMA,
		KeyOemMinus:          w32.VK_OEM_MINUS,
		KeyOemPeriod:         w32.VK_OEM_PERIOD,
		KeyOem2:              w32.VK_OEM_2,
		KeyOem3:              w32.VK_OEM_3,
		KeyOem4:              w32.VK_OEM_4,
		KeyOem5:              w32.VK_OEM_5,
		KeyOem6:              w32.VK_OEM_6,
		KeyOem7:              w32.VK_OEM_7,
		KeyOem8:              w32.VK_OEM_8,
		KeyOemAx:             w32.VK_OEM_AX,
		KeyOem102:            w32.VK_OEM_102,
		KeyIcoHelp:           w32.VK_ICO_HELP,
		KeyIco00:             w32.VK_ICO_00,
		KeyProcessKey:        w32.VK_PROCESSKEY,
		KeyIcoClear:          w32.VK_ICO_CLEAR,
		KeyOemReset:          w32.VK_OEM_RESET,
		KeyOemJump:           w32.VK_OEM_JUMP,
		KeyOemPa1:            w32.VK_OEM_PA1,
		KeyOemPa2:            w32.VK_OEM_PA2,
		KeyOemPa3:            w32.VK_OEM_PA3,
		KeyOemWsctrl:         w32.VK_OEM_WSCTRL,
		KeyOemCusel:          w32.VK_OEM_CUSEL,
		KeyOemAttn:           w32.VK_OEM_ATTN,
		KeyOemFinish:         w32.VK_OEM_FINISH,
		KeyOemCopy:           w32.VK_OEM_COPY,
		KeyOemAuto:           w32.VK_OEM_AUTO,
		KeyOemEnlw:           w32.VK_OEM_ENLW,
		KeyOemBacktab:        w32.VK_OEM_BACKTAB,
		KeyAttn:              w32.VK_ATTN,
		KeyCrsel:             w32.VK_CRSEL,
		KeyExsel:             w32.VK_EXSEL,
		KeyEreof:             w32.VK_EREOF,
		KeyPlay:              w32.VK_PLAY,
		KeyZoom:              w32.VK_ZOOM,
		KeyNoname:            w32.VK_NONAME,
		KeyPa1:               w32.VK_PA1,
		KeyOemClear:          w32.VK_OEM_CLEAR,
	}
)
