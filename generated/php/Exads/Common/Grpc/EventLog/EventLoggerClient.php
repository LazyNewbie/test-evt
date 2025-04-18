<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Exads\Common\Grpc\EventLog;

/**
 */
class EventLoggerClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * @param \Exads\Common\Grpc\EventLog\AdEventRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function Send(\Exads\Common\Grpc\EventLog\AdEventRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/exads.schema.events.EventLogger/Send',
        $argument,
        ['\Exads\Common\Grpc\EventLog\AdEventReply', 'decode'],
        $metadata, $options);
    }

}
