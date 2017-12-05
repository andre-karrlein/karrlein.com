<?php declare(strict_types=1);

require __DIR__ . '/../vendor/autoload.php';

$app = new Slim\App();
$app->get('/', function($request, $response, $args) {
	$factory = new ak1\karrlein\Factory();
	return $response->write($factory->createMainPage()->getTemplate());
});

$app->run();
