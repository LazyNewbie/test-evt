<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Exads\Protobuf\Events;

/**
 */
class EventLogClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * @param \Exads\Protobuf\Events\AdEventRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function Send(\Exads\Protobuf\Events\AdEventRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/exads.schema.events.EventLog/Send',
        $argument,
        ['\Exads\Protobuf\Events\AdEventResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\ClientStreamingCall
     */
    public function SendEvents($metadata = [], $options = []) {
        return $this->_clientStreamRequest('/exads.schema.events.EventLog/SendEvents',
        ['\Exads\Protobuf\Events\AdEventResponse','decode'],
        $metadata, $options);
    }

}
