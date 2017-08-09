<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: books.proto

namespace Books;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * <pre>
 * Message for handling only sending the ID
 * </pre>
 *
 * Protobuf type <code>Books.BookIDRequest</code>
 */
class BookIDRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * <code>string id = 1;</code>
     */
    private $id = '';

    public function __construct() {
        \GPBMetadata\Books::initOnce();
        parent::__construct();
    }

    /**
     * <code>string id = 1;</code>
     */
    public function getId()
    {
        return $this->id;
    }

    /**
     * <code>string id = 1;</code>
     */
    public function setId($var)
    {
        GPBUtil::checkString($var, True);
        $this->id = $var;
    }

}

