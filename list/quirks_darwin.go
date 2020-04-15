// +build darwin

package list

import (
	"github.com/micmonay/keybd_event"
	"golang.org/x/sys/unix"
)

// This method avoid tcell bug https://github.com/gdamore/tcell/issues/194
// Also tracked in https://github.com/ktr0731/go-fuzzyfinder/issues/17
// Aditional EOL event is sent to ensure, consequent events, are correctly handled
func fuzzyquirks() {
	logger.InfoMsg("applying darwin fuzzy quirks")
	termios, err := unix.IoctlGetTermios(0, unix.TIOCGETA)
	if err != nil {
		logger.InfoMsgf("failed to get termios: %s", err)
		return
	}
	newState := *termios
	newState.Lflag &^= unix.ECHO
	if err := unix.IoctlSetTermios(0, unix.TIOCSETA, &newState); err != nil {
		logger.InfoMsgf("failed to set termios: %s", err)
	}
	defer unix.IoctlSetTermios(0, unix.TIOCSETA, termios)

	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		logger.InfoMsgf("failed to get key bonding: %s", err)
		return
	}
	kb.SetKeys(keybd_event.VK_ENTER)
	err = kb.Launching()
	if err != nil {
		logger.InfoMsgf("failed to send event: %s", err)
	}
}
