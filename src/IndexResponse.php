<?php declare(strict_types=1)

namespace karrlein\webpage;

use Illuminate\Http\Response;

class IndexResponse
{

	public function returnResponse(): Response
	{
		return new Response(view('index'),200);
	}

}
