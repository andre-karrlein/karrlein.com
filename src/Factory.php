<?php

namespace karrlein/webpage;

class Factory {

	public static function createIndexResponse() : IndexResponse
	{
		return new IndexResponse();
	}

}
