package gluon

import (
	"crypto/tls"
	"io"
	"os"
	"time"

	"github.com/ProtonMail/gluon/connector"
	"github.com/ProtonMail/gluon/internal/db"
	"github.com/sirupsen/logrus"

	"github.com/ProtonMail/gluon/imap"
	"github.com/ProtonMail/gluon/internal/backend"
	"github.com/ProtonMail/gluon/internal/session"
	"github.com/ProtonMail/gluon/limits"
	"github.com/ProtonMail/gluon/profiling"
	"github.com/ProtonMail/gluon/queue"
	"github.com/ProtonMail/gluon/reporter"
	"github.com/ProtonMail/gluon/store"
	"github.com/ProtonMail/gluon/version"
)

type serverBuilder struct {
	authorizer           connector.Authorizer
	dataDir              string
	databaseDir          string
	delim                string
	loginJailTime        time.Duration
	tlsConfig            *tls.Config
	idleBulkTime         time.Duration
	inLogger             io.Writer
	outLogger            io.Writer
	versionInfo          version.Info
	cmdExecProfBuilder   profiling.CmdProfilerBuilder
	storeBuilder         store.Builder
	reporter             reporter.Reporter
	disableParallelism   bool
	imapLimits           limits.IMAP
	uidValidityGenerator imap.UIDValidityGenerator
}

func newBuilder() (*serverBuilder, error) {
	return &serverBuilder{
		authorizer:           &connector.DummyAuthorizer{},
		delim:                "/",
		cmdExecProfBuilder:   &profiling.NullCmdExecProfilerBuilder{},
		storeBuilder:         &store.OnDiskStoreBuilder{},
		reporter:             &reporter.NullReporter{},
		idleBulkTime:         500 * time.Millisecond,
		imapLimits:           limits.DefaultLimits(),
		uidValidityGenerator: imap.DefaultEpochUIDValidityGenerator(),
	}, nil
}

func (builder *serverBuilder) build() (*Server, error) {
	if builder.dataDir == "" {
		dir, err := os.MkdirTemp("", "gluon-*")
		if err != nil {
			return nil, err
		}

		builder.dataDir = dir
	}

	if err := os.MkdirAll(builder.dataDir, 0o700); err != nil {
		return nil, err
	}

	if builder.databaseDir == "" {
		dir, err := os.MkdirTemp("", "gluon-*")
		if err != nil {
			return nil, err
		}

		builder.databaseDir = dir
	}

	if err := os.MkdirAll(builder.databaseDir, 0o700); err != nil {
		return nil, err
	}

	backend, err := backend.New(
		builder.authorizer,
		builder.dataDir,
		builder.databaseDir,
		builder.storeBuilder,
		builder.delim,
		builder.loginJailTime,
		builder.imapLimits,
	)
	if err != nil {
		return nil, err
	}

	// Defer delete all the previous databases from removed user accounts. This is required since we can't
	// close ent databases on demand.
	if err := db.DeleteDeferredDBFiles(builder.databaseDir); err != nil {
		logrus.WithError(err).Error("Failed to remove old database files")
	}

	return &Server{
		dataDir:              builder.dataDir,
		databaseDir:          builder.databaseDir,
		backend:              backend,
		sessions:             make(map[int]*session.Session),
		serveErrCh:           queue.NewQueuedChannel[error](1, 1),
		serveDoneCh:          make(chan struct{}),
		inLogger:             builder.inLogger,
		outLogger:            builder.outLogger,
		tlsConfig:            builder.tlsConfig,
		idleBulkTime:         builder.idleBulkTime,
		storeBuilder:         builder.storeBuilder,
		cmdExecProfBuilder:   builder.cmdExecProfBuilder,
		versionInfo:          builder.versionInfo,
		reporter:             builder.reporter,
		disableParallelism:   builder.disableParallelism,
		uidValidityGenerator: builder.uidValidityGenerator,
	}, nil
}
