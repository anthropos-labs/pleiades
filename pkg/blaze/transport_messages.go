/*
 * Copyright (c) 2022 Anthropos Labs, Inc.
 *
 * Licensed under the PolyForm Strict License 1.0.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License here:
 *  https://gitlab.com/anthropos-labs/pleiades/-/blob/mainline/LICENSE
 */

package blaze

import (
	"bytes"
	"hash/crc32"
	"io"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/rs/zerolog"
	"google.golang.org/protobuf/proto"
)

func NewMessageStream(stream network.Stream, msg []byte, logger zerolog.Logger) (*MessageStream, error) {
	if len(msg) == 0 || msg == nil {
		err := errors.New("no message to send")
		logger.Error().Err(err).Msg("empty message")
		return nil, err
	}

	return &MessageStream{
		stream: stream,
		body:   msg,
	}, nil
}

type MessageStream struct {
	stream network.Stream
	logger zerolog.Logger
	header requestHeader
	body   []byte
}

func (m *MessageStream) VerifyMagicNumber() error {
	readDeadline := time.Now().Add(magicNumberDuration)
	if err := m.stream.SetReadDeadline(readDeadline); err != nil {
		return err
	}

	magicLen := make([]byte, len(MagicNumber))
	if _, err := io.ReadFull(m.stream, magicLen); err != nil {
		return err
	}

	if !bytes.Equal(magicLen, PoisonNumber[:]) {
		return errPoisonReceived
	}

	if !bytes.Equal(magicLen, MagicNumber[:]) {
		return ErrBadMessage
	}

	return nil
}

func (m *MessageStream) Send(t uint16) error {
	if m.body == nil {
		err := errors.New("no body to send")
		m.logger.Error().Err(err).Msg("body is nil")
		return err
	}

	header := &requestHeader{
		method: t,
		size:   uint64(len(m.body)),
		crc:    crc32.ChecksumIEEE(m.body),
	}

	headerBuf := make([]byte, requestHeaderSize)
	headerBuf = header.encode(headerBuf)

	magicNumDeadline := time.Now().Add(magicNumberDuration)
	if err := m.stream.SetWriteDeadline(magicNumDeadline); err != nil {
		m.logger.Error().Err(err).Msg("failed to set write deadline for magic number")
		return err
	}
	if _, err := m.stream.Write(MagicNumber[:]); err != nil {
		m.logger.Error().Err(err).Msg("failed to write magic number")
		return err
	}

	headerDeadline := time.Now().Add(headerDuration)
	if err := m.stream.SetWriteDeadline(headerDeadline); err != nil {
		m.logger.Error().Err(err).Msg("failed to set write deadline for header")
		return err
	}
	if _, err := m.stream.Write(headerBuf); err != nil {
		m.logger.Error().Err(err).Msg("failed to write header")
		return err
	}

	messageDeadline := time.Now().Add(writeDuration)
	if err := m.stream.SetWriteDeadline(messageDeadline); err != nil {
		m.logger.Error().Err(err).Msg("failed to set write deadline for message")
		return err
	}
	if _, err := m.stream.Write(m.body); err != nil {
		m.logger.Error().Err(err).Msg("failed to write message batch")
		return err
	}

	return nil
}

func (m *MessageStream) Read() (uint16, []byte, error) {
	if err := m.VerifyMagicNumber(); err != nil {
		if err == errPoisonReceived {
			m.logger.Error().Err(err).Msg("poison received")
		}
		if err == ErrBadMessage {
			m.logger.Error().Err(err).Msg("bad message")
			return 0, nil, err
		}
	}

	headerBuf := make([]byte, requestHeaderSize)

	headerDeadline := time.Now().Add(headerDuration)
	if err := m.stream.SetReadDeadline(headerDeadline); err != nil {
		m.logger.Error().Err(err).Msg("failed to set readAndHandle deadline for header")
		return 0, nil, err
	}

	if _, err := io.ReadFull(m.stream, headerBuf); err != nil {
		m.logger.Error().Err(err).Msg("failed to readAndHandle header")
		return 0, nil, err
	}

	header := &requestHeader{}
	if err := header.decode(headerBuf); err != nil {
		m.logger.Error().Err(err).Msg("failed to decode header")
		return 0, nil, err
	}

	if header.size == 0 {
		m.logger.Error().Msg("invalid message size")
		return 0, nil, ErrBadMessage
	}

	buf := make([]byte, header.size)
	messageDeadline := time.Now().Add(readDuration)
	if err := m.stream.SetReadDeadline(messageDeadline); err != nil {
		m.logger.Error().Err(err).Msg("failed to set readAndHandle deadline for message")
		return 0, nil, err
	}

	if _, err := io.ReadFull(m.stream, buf); err != nil {
		m.logger.Error().Err(err).Msg("failed to readAndHandle message")
		return 0, nil, err
	}

	if crc32.ChecksumIEEE(buf) != header.crc {
		err := errors.New("invalid message checksum")
		m.logger.Error().Err(err).Msg("invalid message checksum")
		return 0, nil, err
	}

	return header.method, buf, nil
}

func unmarshal[T proto.Message](payload []byte) (T, error) {
	var t T
	err := proto.Unmarshal(payload, t)
	return t, err
}

func sendFrame(frame *Frame, logger zerolog.Logger, stream network.Stream) {
	respBuf, err := frame.Marshal()
	if err != nil {
		// todo (sienna): add error handling
		logger.Error().Err(err).Msg("error marshaling frame")
	}

	// set the write deadline
	deadline := time.Now().Add(RaftControlRPCWriteTimeout)
	if err := stream.SetWriteDeadline(deadline); err != nil {
		_ = stream.Reset()
	}

	_, err = stream.Write(respBuf)
	if err != nil {
		// todo (sienna): add error handling
		logger.Error().Err(err).Msg("error sending frame")
	}
}
