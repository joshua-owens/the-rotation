<?php

namespace App\Providers;

use App\RecipeScraper;
use App\Scraper;
use Illuminate\Support\ServiceProvider;

class AppServiceProvider extends ServiceProvider
{

    /**
     * All the container singletons that should be registered.
     *
     * @var array
     */
    public $singletons = [
        Scraper::class => RecipeScraper::class,
    ];
}
