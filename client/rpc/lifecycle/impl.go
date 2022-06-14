package lifecycle

import (
	"context"
	"io"
	"time"

	"github.com/infraboard/keyauth/apps/instance"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewLifecycler(
	client instance.ServiceClient,
	ins *instance.Instance,
) Lifecycler {
	return &manager{
		client: client,
		ins:    ins,
		log:    zap.L().Named("lifecyle"),
	}
}

type manager struct {
	client instance.ServiceClient
	log    logger.Logger
	ins    *instance.Instance
	stream instance.Service_HeartbeatClient
}

func (m *manager) Heartbeat(ctx context.Context) error {
	stream, err := m.client.Heartbeat(ctx)
	if err != nil {
		return err
	}
	m.stream = stream

	go m.sender(ctx)
	go m.receiver(ctx)
	return nil
}

func (m *manager) HeartbeatInterval() time.Duration {
	if m.ins.Config.Heartbeat.Interval == 0 {
		m.ins.Config.Heartbeat.Interval = 5
	}

	return time.Duration(m.ins.Config.Heartbeat.Interval) * time.Second
}

// 发送心跳
func (m *manager) sender(ctx context.Context) {
	defer func() {
		if err := recover(); err != nil {
			m.log.Errorf("sender panic: %s", err)
		}
	}()

	tk := time.NewTicker(m.HeartbeatInterval())
	for {
		select {
		case <-ctx.Done():
			m.log.Infof("heartbeat sender stoped")
			return
		case <-tk.C:
			if err := m.stream.Send(instance.NewHeartbeatRequest(m.ins.Id)); err != nil {
				m.log.Errorf("send heartbeat error, %s", err)
			}
		}
	}
}

// 循环接收服务端返回的数据
func (m *manager) receiver(ctx context.Context) {
	defer func() {
		if err := recover(); err != nil {
			m.log.Error("receiver panic: %s", err)
		}
	}()

	for {
		select {
		case <-ctx.Done():
			m.log.Infof("heartbeat receiver stoped")
			return
		default:
			_, err := m.stream.Recv()
			if err != nil {
				if err == io.EOF {
					m.log.Infof("receive heartbeat response error, server close")
					return
				}
				m.log.Errorf("receive heartbeat receive error, %s", err)
			}
		}
	}
}

// 注销实例
func (m *manager) UnRegistry(ctx context.Context) error {
	ins, err := m.client.UnRegistry(ctx, instance.NewUnregistryRequest(m.ins.Id))
	if err != nil {
		m.log.Errorf("instance %s unregistry error, %s", m.ins.FullName(), err)
	} else {
		m.log.Infof("instance %s unregistry success", ins.FullName())
	}

	return nil
}
