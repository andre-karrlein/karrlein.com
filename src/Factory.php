<?php declare(strict_types=1);

namespace ak1\karrlein;

class Factory
{
	public function createMainPage(): HtmlTemplate
	{
		return HtmlTemplate::fromFile(__DIR__ . '/../html/Index.html');
	}
}

