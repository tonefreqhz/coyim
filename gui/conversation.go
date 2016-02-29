package gui

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/twstrike/coyim/client"
	"github.com/twstrike/coyim/i18n"
	"github.com/twstrike/coyim/ui"
	"github.com/twstrike/gotk3adapter/glibi"
	"github.com/twstrike/gotk3adapter/gtki"
)

var (
	enableWindow  glibi.Signal
	disableWindow glibi.Signal
)

type conversationView interface {
	showIdentityVerificationWarning(u *gtkUI)
	updateSecurityWarning()
	show(userInitiated bool)
	appendStatus(from string, timestamp time.Time, show, showStatus string, gone bool)
	appendMessage(from string, timestamp time.Time, encrypted bool, message []byte, outgoing bool)
	displayNotification(notification string)
	displayNotificationVerifiedOrNot(notificationV, notificationNV string)
	setEnabled(enabled bool)
}

type conversationWindow struct {
	*conversationPane
	win       gtki.Window
	parentWin gtki.Window
}

type conversationPane struct {
	to                 string
	account            *account
	widget             gtki.Box
	menubar            gtki.MenuBar
	entry              gtki.Entry
	history            gtki.TextView
	scrollHistory      gtki.ScrolledWindow
	notificationArea   gtki.Box
	securityWarning    gtki.InfoBar
	fingerprintWarning gtki.InfoBar
	// The window to set dialogs transient for
	transientParent gtki.Window
	sync.Mutex
}

type tags struct {
	table gtki.TextTagTable
}

func (u *gtkUI) getTags() *tags {
	if u.tags == nil {
		u.tags = newTags()
	}
	return u.tags
}

func newTags() *tags {
	t := new(tags)

	t.table, _ = g.gtk.TextTagTableNew()

	outgoingUser, _ := g.gtk.TextTagNew("outgoingUser")
	outgoingUser.SetProperty("foreground", "#3465a4")

	incomingUser, _ := g.gtk.TextTagNew("incomingUser")
	incomingUser.SetProperty("foreground", "#a40000")

	outgoingText, _ := g.gtk.TextTagNew("outgoingText")
	outgoingText.SetProperty("foreground", "#555753")

	incomingText, _ := g.gtk.TextTagNew("incomingText")

	statusText, _ := g.gtk.TextTagNew("statusText")
	statusText.SetProperty("foreground", "#4e9a06")

	t.table.Add(outgoingUser)
	t.table.Add(incomingUser)
	t.table.Add(outgoingText)
	t.table.Add(incomingText)
	t.table.Add(statusText)

	return t
}

func (t *tags) createTextBuffer() gtki.TextBuffer {
	buf, _ := g.gtk.TextBufferNew(t.table)
	return buf
}

func createConversationPane(account *account, uid string, ui *gtkUI, transientParent gtki.Window) *conversationPane {
	builder := builderForDefinition("ConversationPane")

	obj, _ := builder.GetObject("box")
	pane := obj.(gtki.Box)

	obj, _ = builder.GetObject("history")
	history := obj.(gtki.TextView)

	obj, _ = builder.GetObject("historyScroll")
	scrollHistory := obj.(gtki.ScrolledWindow)

	obj, _ = builder.GetObject("message")
	entry := obj.(gtki.Entry)

	obj, _ = builder.GetObject("notification-area")
	notificationArea := obj.(gtki.Box)

	obj, _ = builder.GetObject("security-warning")
	securityWarning := obj.(gtki.InfoBar)

	obj, _ = builder.GetObject("menubar")
	menubar := obj.(gtki.MenuBar)

	cp := &conversationPane{
		to:               uid,
		account:          account,
		history:          history,
		widget:           pane,
		menubar:          menubar,
		entry:            entry,
		scrollHistory:    scrollHistory,
		notificationArea: notificationArea,
		securityWarning:  securityWarning,
		transientParent:  transientParent,
	}

	builder.ConnectSignals(map[string]interface{}{
		"on_send_message_signal": func() {
			entry.SetEditable(false)
			text, _ := entry.GetText()
			entry.SetText("")
			entry.SetEditable(true)
			if text != "" {
				sendError := cp.sendMessage(text)
				if sendError != nil {
					fmt.Printf(i18n.Local("Failed to generate OTR message: %s\n"), sendError.Error())
				}
			}
			entry.GrabFocus()
		},
		// TODO: basically I think this whole menu should be rethought. It's useful for us to have during development
		"on_start_otr_signal": func() {
			//TODO: enable/disable depending on the conversation's encryption state
			session := cp.account.session
			c, _ := session.ConversationManager().EnsureConversationWith(cp.to)
			err := c.StartEncryptedChat(session)
			if err != nil {
				//TODO: notify failure
			}
		},
		"on_end_otr_signal": func() {
			//TODO: errors
			//TODO: enable/disable depending on the conversation's encryption state
			session := cp.account.session
			c, ok := session.ConversationManager().GetConversationWith(cp.to)
			if !ok {
				return
			}

			err := c.EndEncryptedChat(session)
			if err != nil {
				fmt.Printf(i18n.Local("Failed to terminate the encrypted chat: %s\n"), err.Error())
			}
		},
		"on_verify_fp_signal": func() {
			switch verifyFingerprintDialog(cp.account, cp.to, transientParent) {
			case gtki.RESPONSE_YES:
				cp.removeIdentityVerificationWarning()
			}
		},
		"on_connect": func() {
			entry.SetEditable(true)
			entry.SetSensitive(true)
		},
		"on_disconnect": func() {
			entry.SetEditable(false)
			entry.SetSensitive(false)
		},
	})

	cp.history.SetBuffer(ui.getTags().createTextBuffer())

	cp.history.Connect("size-allocate", func() {
		cp.scrollToBottom()
	})

	ui.displaySettings.control(cp.history)
	ui.displaySettings.control(entry)

	return cp
}

func newConversationWindow(account *account, uid string, ui *gtkUI) *conversationWindow {
	builder := builderForDefinition("Conversation")

	obj, _ := builder.GetObject("conversation")
	win := obj.(gtki.Window)
	title := fmt.Sprintf("%s <-> %s", account.session.GetConfig().Account, uid)
	win.SetTitle(title)

	obj, _ = builder.GetObject("box")
	winBox := obj.(gtki.Box)

	cp := createConversationPane(account, uid, ui, win)
	winBox.PackStart(cp.widget, true, true, 0)

	conv := &conversationWindow{
		conversationPane: cp,
		win:              win,
	}

	// Unlike the GTK version, this is not supposed to be used as a callback but
	// it attaches the callback to the widget
	conv.win.HideOnDelete()

	inEventHandler := false
	conv.win.Connect("set-focus", func() {
		if !inEventHandler {
			inEventHandler = true
			conv.entry.GrabFocus()
			inEventHandler = false
		}
	})

	conv.win.Connect("notify::is-active", func() {
		if conv.win.IsActive() {
			inEventHandler = true
			conv.entry.GrabFocus()
			inEventHandler = false
		}
	})

	ui.connectShortcutsChildWindow(conv.win)
	ui.connectShortcutsConversationWindow(conv)
	conv.parentWin = ui.window

	return conv
}

func (conv *conversationPane) addNotification(notification gtki.InfoBar) {
	conv.notificationArea.Add(notification)
}

func (conv *conversationWindow) Hide() {
	conv.win.Hide()
}

func (conv *conversationWindow) tryEnsureCorrectWorkspace() {
	if g.gdk.WorkspaceControlSupported() {
		wi, _ := conv.parentWin.GetWindow()
		parentPlace := wi.GetDesktop()
		cwi, _ := conv.win.GetWindow()
		cwi.MoveToDesktop(parentPlace)
	}
}

func (conv *conversationPane) getConversation() (client.Conversation, bool) {
	return conv.account.session.ConversationManager().GetConversationWith(conv.to)
}

func (conv *conversationPane) isVerified() bool {
	conversation, exists := conv.getConversation()
	if !exists {
		log.Println("Conversation does not exist - this shouldn't happen")
		return false
	}

	fingerprint := conversation.TheirFingerprint()
	conf := conv.account.session.GetConfig()

	p, hasPeer := conf.GetPeer(conv.to)

	if hasPeer {
		p.EnsureHasFingerprint(fingerprint)
	}

	return hasPeer && p.HasTrustedFingerprint(fingerprint)
}

func (conv *conversationPane) showIdentityVerificationWarning(u *gtkUI) {
	conv.Lock()
	defer conv.Unlock()

	if conv.fingerprintWarning != nil {
		log.Println("we are already showing a fingerprint warning, so not doing it again")
		return
	}

	if conv.isVerified() {
		log.Println("We have a peer and a trusted fingerprint already, so no reason to warn")
		return
	}

	conv.fingerprintWarning = buildVerifyIdentityNotification(conv.account, conv.to, conv.transientParent)
	conv.addNotification(conv.fingerprintWarning)
}

func (conv *conversationPane) removeIdentityVerificationWarning() {
	conv.Lock()
	defer conv.Unlock()

	if conv.fingerprintWarning != nil {
		conv.fingerprintWarning.Hide()
		conv.fingerprintWarning.Destroy()
		conv.fingerprintWarning = nil
	}
}

func (conv *conversationPane) updateSecurityWarning() {
	conversation, ok := conv.getConversation()
	if !ok {
		return
	}

	conv.securityWarning.SetVisible(!conversation.IsEncrypted())
}

func (conv *conversationWindow) show(userInitiated bool) {
	conv.updateSecurityWarning()
	conv.win.Show()
	conv.tryEnsureCorrectWorkspace()
}

func (conv *conversationPane) sendMessage(message string) error {
	err := conv.account.session.EncryptAndSendTo(conv.to, message)
	if err != nil {
		return err
	}

	//TODO: review whether it should create a conversation
	//TODO: this should be whether the message was encrypted or not, rather than
	//whether the conversation is encrypted or not
	conversation, _ := conv.account.session.ConversationManager().EnsureConversationWith(conv.to)
	conv.appendMessage(conv.account.session.GetConfig().Account, time.Now(), conversation.IsEncrypted(), ui.StripHTML([]byte(message)), true)

	return nil
}

const timeDisplay = "15:04:05"

// Expects to be called from the GUI thread.
// Expects to be called when conv is already locked
func insertAtEnd(buff gtki.TextBuffer, text string) {
	buff.Insert(buff.GetEndIter(), text)
}

// Expects to be called from the GUI thread.
// Expects to be called when conv is already locked
func insertWithTag(buff gtki.TextBuffer, tagName, text string) {
	charCount := buff.GetCharCount()
	insertAtEnd(buff, text)
	oldEnd := buff.GetIterAtOffset(charCount)
	newEnd := buff.GetEndIter()
	buff.ApplyTagByName(tagName, oldEnd, newEnd)
}

func is(v bool, left, right string) string {
	if v {
		return left
	}
	return right
}

func showForDisplay(show string, gone bool) string {
	switch show {
	case "", "available", "online":
		if gone {
			return ""
		}
		return i18n.Local("Available")
	case "xa":
		return i18n.Local("Not Available")
	case "away":
		return i18n.Local("Away")
	case "dnd":
		return i18n.Local("Busy")
	case "chat":
		return i18n.Local("Free for Chat")
	case "invisible":
		return i18n.Local("Invisible")
	}
	return show
}

func onlineStatus(show, showStatus string) string {
	sshow := showForDisplay(show, false)
	if sshow != "" {
		return sshow + showStatusForDisplay(showStatus)
	}
	return ""
}

func showStatusForDisplay(showStatus string) string {
	if showStatus != "" {
		return " (" + showStatus + ")"
	}
	return ""
}

func extraOfflineStatus(show, showStatus string) string {
	sshow := showForDisplay(show, true)
	if sshow == "" {
		return showStatusForDisplay(showStatus)
	}

	if showStatus != "" {
		return " (" + sshow + ": " + showStatus + ")"
	}
	return " (" + sshow + ")"
}

func createStatusMessage(from string, show, showStatus string, gone bool) string {
	tail := ""
	if gone {
		tail = i18n.Local("Offline") + extraOfflineStatus(show, showStatus)
	} else {
		tail = onlineStatus(show, showStatus)
	}

	if tail != "" {
		return from + i18n.Local(" is now ") + tail
	}
	return ""
}

func (conv *conversationPane) scrollToBottom() {
	adj := conv.scrollHistory.GetVAdjustment()
	adj.SetValue(adj.GetUpper() - adj.GetPageSize())
}

type taggableText struct {
	tag  string
	text string
}

func (conv *conversationPane) appendToHistory(timestamp time.Time, entries ...taggableText) {
	doInUIThread(func() {
		conv.Lock()
		defer conv.Unlock()

		buff, _ := conv.history.GetBuffer()
		if buff.GetCharCount() != 0 {
			insertAtEnd(buff, "\n")
		}

		insertAtEnd(buff, "[")
		insertAtEnd(buff, timestamp.Format(timeDisplay))
		insertAtEnd(buff, "] ")

		for _, entry := range entries {
			if entry.tag != "" {
				insertWithTag(buff, entry.tag, entry.text)
			} else {
				insertAtEnd(buff, entry.text)
			}
		}
	})
}

func (conv *conversationPane) appendStatus(from string, timestamp time.Time, show, showStatus string, gone bool) {
	conv.appendToHistory(timestamp, taggableText{"statusText", createStatusMessage(from, show, showStatus, gone)})
}

func (conv *conversationPane) appendMessage(from string, timestamp time.Time, encrypted bool, message []byte, outgoing bool) {
	conv.appendToHistory(timestamp,
		taggableText{
			is(outgoing, "outgoingUser", "incomingUser"),
			from,
		},
		taggableText{
			text: ":  ",
		},
		taggableText{
			is(outgoing, "outgoingText", "incomingText"),
			string(message),
		})
}

func (conv *conversationPane) displayNotification(notification string) {
	conv.appendToHistory(time.Now(), taggableText{"statusText", notification})
}

func (conv *conversationPane) displayNotificationVerifiedOrNot(notificationV, notificationNV string) {
	if conv.isVerified() {
		conv.displayNotification(notificationV)
	} else {
		conv.displayNotification(notificationNV)
	}
}

func (conv *conversationWindow) setEnabled(enabled bool) {
	if enabled {
		conv.win.Emit("enable")
	} else {
		conv.win.Emit("disable")
	}
}
