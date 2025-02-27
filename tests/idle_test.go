package tests

import (
	"testing"
	"time"

	"github.com/ProtonMail/gluon/imap"
)

func TestIDLEExistsUpdates(t *testing.T) {
	runManyToOneTestWithAuth(t, defaultServerOptions(t), []int{1, 2}, func(c map[int]*testConnection, s *testSession) {
		// First client selects in INBOX to receive EXISTS update.
		c[1].C("A006 select INBOX")
		c[1].Se("A006 OK [READ-WRITE] SELECT")

		// First client starts to IDLE.
		c[1].C("A007 IDLE")
		c[1].S("+ Ready")

		// Second client appends to INBOX to generate EXISTS updates.
		// The client is not selected and thus doesn't itself receive responses.
		c[2].doAppend(`INBOX`, `To: 1@pm.me`, `\Seen`).expect("OK")
		c[2].doAppend(`INBOX`, `To: 2@pm.me`, `\Seen`).expect("OK")

		// First client receives the EXISTS and RECENT updates while idling.
		c[1].S(`* 2 EXISTS`, `* 2 RECENT`)

		// First client stops idling.
		c[1].C("DONE")
		c[1].OK(`A007`)

		// Further stuff doesn't trigger any issues.
		c[2].doAppend(`INBOX`, `To: 3@pm.me`, `\Seen`).expect("OK")
	})
}

func TestIDLEPendingUpdates(t *testing.T) {
	runManyToOneTestWithData(t, defaultServerOptions(t), []int{1, 2}, func(c map[int]*testConnection, s *testSession, _ string, _ imap.MailboxID) {
		c[1].C("A001 select INBOX").OK("A001")

		// Generate some pending updates.
		c[2].C("B001 UID MOVE 1,2,3 INBOX").OK("B001")

		// Begin IDLE.
		c[1].C("A002 IDLE").S("+ Ready")

		// Generate some additional updates.
		c[2].C("B002 UID MOVE 4,5,6 INBOX").OK("B002")

		// Pending updates are first flushed.
		c[1].Se(`* 3 EXISTS`)

		// IDLE updates are first second.
		c[1].Se(`* 6 EXISTS`)

		// Stop IDLE.
		c[1].C("DONE").OK("A002")
	})
}

func TestIDLERecentReceivedOnSelectedClient(t *testing.T) {
	runManyToOneTestWithAuth(t, defaultServerOptions(t), []int{1, 2}, func(c map[int]*testConnection, s *testSession) {
		c[1].C("A001 select INBOX").OK("A001")

		// Generate some pending updates.
		c[2].doAppend("INBOX", "To: foo.foo").s.OK("")
		c[2].doAppend("INBOX", "To: bar.bar").s.OK("")

		// Begin IDLE.
		c[1].C("A002 IDLE").S("+ Ready")

		// Pending updates are first flushed. And merged.
		c[1].S(`* 2 EXISTS`, `* 2 RECENT`)

		// And bulked after 500 ms
		c[2].doAppend("INBOX", "To: biz.biz").s.OK("")
		c[2].doAppend("INBOX", "To: gardy.loo").s.OK("")

		// Testing splitting the updates in at least two bulks.
		time.Sleep(510 * time.Millisecond)

		c[2].doAppend("INBOX", "To: wolo.lo").s.OK("")
		c[2].doAppend("INBOX", "To: huga.chaga").s.OK("")
		c[2].C("C2 LOGOUT").OK("C2")

		// First bulk update should have 4 exists and recent.
		c[1].S(`* 4 EXISTS`, `* 4 RECENT`)
		// Depneds on append performance the 6 exist and recent in
		// second or third bulk.
		c[1].Se(`* 6 EXISTS`, `* 6 RECENT`)

		// Stop IDLE.
		c[1].C("DONE").OK("A002")
		c[1].C("A2 LOGOUT").OK("A2")
	})
}
