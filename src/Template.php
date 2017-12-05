<?php declare(strict_types=1);

namespace ak1\karrlein;

interface Template
{
	public function fromFile(string $filePath): Template;

	public function getTemplate(): string;
}
