package tracer

import (
	"github.com/Lionparcel/dms_internal/shared/logger"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmot"
	"go.elastic.co/apm/transport"
	"os"

	"github.com/opentracing/opentracing-go"
)

const (
	Byte          uint64 = 1
	maxPacketSize        = int(65000 * Byte)
)

var agent string

func InitOpenTracing() error {
	apm.DefaultTracer.Close()
	os.Setenv(`ELASTIC_APM_SERVICE_NAME`, os.Getenv(`ELASTIC_APM_SERVICE_NAME`))
	os.Setenv(`ELASTIC_APM_SERVER_URL`, os.Getenv(`ELASTIC_APM_SERVER_URL`))
	os.Setenv(`ELASTIC_APM_SECRET_TOKEN`, os.Getenv(`ELASTIC_APM_SECRET_TOKEN`))

	if _, err := transport.InitDefault(); err != nil {
		logger.E(err)
		return err
	}

	tracer, err := apm.NewTracer(os.Getenv(`ELASTIC_APM_SERVICE_NAME`), `1`)
	if err != nil {
		logger.E(err)
		return err
	}

	opentracing.SetGlobalTracer(apmot.New(apmot.WithTracer(tracer)))
	agent = os.Getenv(`ELASTIC_APM_SERVER_URL`)

	return nil
}
