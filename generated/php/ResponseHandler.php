<?php
declare(strict_types=1);

namespace Exads\Common\Grpc;

use stdClass;

class ResponseHandler
{

    private ?array $response;
    private stdClass $status;


    public function __construct(array $grpcResponse)
    {
        list($this->response, $this->status) = $grpcResponse;
    }

}