<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Exads\Grpc\Rocinante;

/**
 */
class AdEventHandlerClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * @param \Exads\Grpc\Rocinante\AdEventLog $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function SendLog(\Exads\Grpc\Rocinante\AdEventLog $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/exads.schema.events.AdEventHandler/SendLog',
        $argument,
        ['\Exads\Grpc\Rocinante\SendLogResponse', 'decode'],
        $metadata, $options);
    }

}
