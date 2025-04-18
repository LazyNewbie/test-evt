<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Exads\Common\Grpc\EventLog;

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
     * @param \Exads\Common\Grpc\EventLog\EventLogRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function SendLog(\Exads\Common\Grpc\EventLog\EventLogRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/exads.schema.events.EventLog/SendLog',
        $argument,
        ['\Exads\Common\Grpc\EventLog\SendLogReply', 'decode'],
        $metadata, $options);
    }

}
