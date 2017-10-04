<?php

namespace karrlein\webpage;

use Illuminate\Http\Response;

class IndexResponse
{

	public function returnResponse()
	{
		return new Response(view('index',200));
	}

}
