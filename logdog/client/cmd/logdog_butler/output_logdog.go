// Copyright 2015 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package main

import (
	"fmt"
	"time"

	"github.com/luci/luci-go/common/clock/clockflag"
	"github.com/luci/luci-go/common/errors"
	"github.com/luci/luci-go/common/flag/multiflag"
	"github.com/luci/luci-go/logdog/client/butler/output"
	out "github.com/luci/luci-go/logdog/client/butler/output/logdog"
)

func init() {
	registerOutputFactory(new(logdogOutputFactory))
}

// logdogOutputFactory for publishing logs using a LogDog Coordinator host.
type logdogOutputFactory struct {
	service          string
	prefixExpiration clockflag.Duration

	track bool
}

var _ outputFactory = (*logdogOutputFactory)(nil)

func (f *logdogOutputFactory) option() multiflag.Option {
	opt := newOutputOption("logdog", "Output to a LogDog Coordinator instance.", f)

	flags := opt.Flags()
	flags.StringVar(&f.service, "service", "",
		"Optional service within <host> to use. Will be referenced as <service>-dot-<host>.")
	flags.Var(&f.prefixExpiration, "prefix-expiration",
		"Amount of time after registration that the prefix will be active. If omitted, the service "+
			"default will be used. This should exceed the expected lifetime of the job by a fair margin.")

	// TODO(dnj): Default to false when mandatory debugging is finished.
	flags.BoolVar(&f.track, "track", true,
		"Track each sent message and dump at the end. This adds CPU/memory overhead.")

	return opt
}

func (f *logdogOutputFactory) configOutput(a *application) (output.Output, error) {
	auth, err := a.authenticator(a)
	if err != nil {
		return nil, errors.Annotate(err).Reason("failed to instantiate authenticator").Err()
	}

	host := a.coordinatorHost
	if host == "" {
		return nil, errors.New("logdog output requires a Coordinator host (-coordinator-host)")
	}
	if f.service != "" {
		host = fmt.Sprintf("%s-dot-%s", f.service, host)
	}

	cfg := out.Config{
		Auth:             auth,
		Host:             host,
		Project:          a.project,
		Prefix:           a.prefix,
		PrefixExpiration: time.Duration(f.prefixExpiration),
		SourceInfo: []string{
			"LogDog Butler",
		},
		PublishContext: a.ncCtx,
		RPCTimeout:     30 * time.Second,
		Track:          f.track,
	}
	return cfg.Register(a)
}

func (f *logdogOutputFactory) scopes() []string { return out.Scopes() }
