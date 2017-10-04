<?php declare(strict_types=1)

namespace karrlein\webpage;

class Factory
{

	public static function createIndexResponse(): IndexResponse
	{
		return new IndexResponse();
	}

}
