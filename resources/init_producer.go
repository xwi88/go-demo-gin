// Package resources kafka producers include
// async and sync producer
package resources

import (
	"time"

	"github.com/Shopify/sarama"
	"github.com/kdpujie/log4go"
	"github.com/xwi88/kit4go/json"
	"github.com/xwi88/kit4go/kafka"

	"github.com/xwi88/go-demo-gin/configs"
)

const LayoutDefaultTimestampFormat = "2006-01-02T15:04:05.000+0800"

const (
	ResourceNameProducer      = "producer"       // producer
	ResourceNameAsyncProducer = "async_producer" // async producer
	ResourceNameSyncProducer  = "sync_producer"  // sync producer
)

var (
	serverIP string
	publicIP string
	hostname string
	appEnv   string
)

// Producer kafka producer
type Producer struct {
	ap           *kafka.AsyncProducer
	sp           *kafka.SyncProducer
	producerName string
	async        bool
	topic        string
}

var producer *Producer

// GetProducer get kafka producer
func GetProducer() *Producer {
	if producer == nil {
		var err error
		if producer, err = initProducer(); err != nil {
			log4go.Error(err.Error())
			return nil
		}
	}
	return producer
}

// initProducer create kafka producer instance
func initProducer() (*Producer, error) {
	cfg := configs.GetCfg()
	kfp := cfg.Kafka.Producer
	if !kfp.Enable {
		log4go.Warn("[%v] disabled", ResourceNameProducer)
		return nil, nil
	}

	bufferSize := kfp.BufferSize
	topic := kfp.Topic
	async := kfp.Async
	brokers := kfp.Brokers
	// kafkaVer := sarama.V0_10_0_1
	kafkaVer := sarama.V2_5_0_0 // now 2.5.0, ref https://kafka.apache.org/downloads#2.5.0
	versionStr := kfp.Version
	if kfp.SpecifyVersion {
		if versionStr != "" {
			if kafkaVersion, err := sarama.ParseKafkaVersion(versionStr); err == nil {
				// should be careful set the version, maybe occur EOF error
				kafkaVer = kafkaVersion
			}
		}
	}
	config := sarama.NewConfig()
	// if not specify the version, use the sarama.V0_10_0_1 to guarante the timestamp can be control
	config.Version = kafkaVer
	config.Producer.Return.Successes = kfp.ReturnSuccesses
	config.Producer.Return.Errors = kfp.ReturnSuccesses
	config.Producer.Timeout = kfp.Timeout
	// NewHashPartitioner returns a Partitioner which behaves as follows. If the message's key is nil then a
	// random partition is chosen. Otherwise the FNV-1a hash of the encoded bytes of the message key is used,
	// modulus the number of partitions. This ensures that messages with the same key always end up on the
	// same partition.
	// config.Producer.Partitioner = sarama.NewHashPartitioner
	// config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	// cfg.Producer.Partitioner = sarama.NewReferenceHashPartitioner

	if async {
		ap, err := kafka.NewAsyncProducer(brokers, bufferSize, config)
		if err != nil {
			return nil, err
		}
		producer = &Producer{ap: ap, async: async, topic: topic,
			producerName: ResourceNameAsyncProducer}
		err = Register(ResourceNameAsyncProducer, producer)
		log4go.Debug("[%v] 初始化完成，配置详情: %#v", ResourceNameAsyncProducer, kfp)
		return producer, nil
	} else {
		sp, err := kafka.NewSyncProducer(brokers, bufferSize, config)
		if err != nil {
			return nil, err
		}
		producer = &Producer{sp: sp, async: async, topic: topic,
			producerName: ResourceNameSyncProducer}
		err = Register(ResourceNameSyncProducer, producer)
		log4go.Debug("[%] 初始化完成，配置详情: %#v", ResourceNameSyncProducer, kfp)
		return producer, nil
	}
}

// SendPBMessage send protobuf message
func (f *Producer) SendPBMessage(data []byte) {
	if f == nil || len(data) == 0 {
		return
	}
	topic := f.topic
	if len(topic) == 0 {
		log4go.Error("[%v] SendPBMessage err: blank topic", f.producerName)
		return
	}
	msg := &sarama.ProducerMessage{
		Topic: topic,
		// Timestamp: time.Now(), // auto generate
		Key:   sarama.ByteEncoder(""),
		Value: sarama.ByteEncoder(data),
	}
	f.Send(msg)
}

// SendJsonMessageWithEsIndex write message to kafka with esIndex
func (f *Producer) SendJsonMessageWithEsIndex(data []byte, esIndex string) {
	if f == nil || len(data) == 0 {
		return
	}
	topic := f.topic
	if len(topic) == 0 {
		log4go.Error("[%v] SendJsonMessageWithEsIndex err: blank topic", f.producerName)
		return
	}

	// 给数据添加一个 esIndex 字段，便于 kibana 进行检索
	var dataMap map[string]interface{}
	err := json.Unmarshal(data, &dataMap)
	if err != nil {
		log4go.Error(err.Error())
		return
	}
	kafkaDataMap := make(map[string]interface{})
	if esIndex != "" {
		kafkaDataMap["esIndex"] = esIndex
	}
	// fill timeStamp for kibana search
	kafkaDataMap["app_env"] = appEnv
	kafkaDataMap["timestamp"] = time.Now().Format(LayoutDefaultTimestampFormat)
	kafkaDataMap["message"] = dataMap
	kafkaDataMap["server_ip"] = serverIP
	kafkaDataMap["public_ip"] = publicIP
	kafkaDataMap["hostname"] = hostname

	dataDealtByte, err := json.Marshal(kafkaDataMap)
	if err != nil {
		log4go.Error(err.Error())
		return
	}

	msg := &sarama.ProducerMessage{
		Topic: f.topic,
		// Timestamp: time.Now(), // auto generate
		Key:   sarama.ByteEncoder(""),
		Value: sarama.ByteEncoder(dataDealtByte),
	}
	f.Send(msg)
}

// Send ProducerMessage
func (f *Producer) Send(msg *sarama.ProducerMessage) {
	if msg != nil {
		if f.async {
			f.ap.Send(msg)
		} else {
			f.sp.Send(msg)
		}
	}
}

// Close the producer
func (f *Producer) Close() error {
	if f.async {
		return f.ap.Close()
	}
	return f.sp.Close()
}
